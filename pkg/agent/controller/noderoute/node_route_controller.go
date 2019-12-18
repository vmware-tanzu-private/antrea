// Copyright 2019 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package noderoute

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/vishvananda/netlink"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"

	"github.com/vmware-tanzu/antrea/pkg/agent/interfacestore"
	"github.com/vmware-tanzu/antrea/pkg/agent/openflow"
	"github.com/vmware-tanzu/antrea/pkg/agent/route"
	"github.com/vmware-tanzu/antrea/pkg/agent/types"
	"github.com/vmware-tanzu/antrea/pkg/agent/util"
	"github.com/vmware-tanzu/antrea/pkg/ovs/ovsconfig"
)

const (
	controllerName = "AntreaAgentNodeRouteController"
	// Interval of synchronizing node status from apiserver
	nodeSyncPeriod = 60 * time.Second
	// How long to wait before retrying the processing of a node change
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing a node change
	defaultWorkers = 4

	ovsExternalIDNodeName = "node-name"
)

// Controller is responsible for setting up necessary IP routes and Openflow entries for inter-node traffic.
type Controller struct {
	kubeClient       clientset.Interface
	nodeInformer     coreinformers.NodeInformer
	nodeLister       corelisters.NodeLister
	nodeListerSynced cache.InformerSynced
	queue            workqueue.RateLimitingInterface
	ofClient         openflow.Client
	ovsBridgeClient  ovsconfig.OVSBridgeClient
	interfaceStore   interfacestore.InterfaceStore
	nodeConfig       *types.NodeConfig
	tunnelType       ovsconfig.TunnelType
	// Pre-shared key for IPSec IKE authentication. If not empty IPSec tunnels will
	// be enabled.
	ipsecPSK    string
	gatewayLink netlink.Link
	routeClient *route.Client
	// installedNodes records routes and flows installation states of Nodes.
	// The key is the host name of the Node, the value is the route to the Node.
	// If the flows of the Node are installed, the installedNodes must contains a key which is the host name.
	// If the route of the Node are installed, the flows of the Node must be installed first and the value of host name
	// key must not be nil.
	// TODO: handle agent restart cases.
	installedNodes *sync.Map
}

// NewNodeRouteController instantiates a new Controller object which will process Node events
// and ensure connectivity between different Nodes.
func NewNodeRouteController(
	kubeClient clientset.Interface,
	informerFactory informers.SharedInformerFactory,
	client openflow.Client,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	interfaceStore interfacestore.InterfaceStore,
	config *types.NodeConfig,
	tunnelType ovsconfig.TunnelType,
	ipsecPSK string,
	routeClient *route.Client,
) *Controller {
	nodeInformer := informerFactory.Core().V1().Nodes()
	link, _ := netlink.LinkByName(config.GatewayConfig.Name)

	controller := &Controller{
		kubeClient:       kubeClient,
		nodeInformer:     nodeInformer,
		nodeLister:       nodeInformer.Lister(),
		nodeListerSynced: nodeInformer.Informer().HasSynced,
		queue:            workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(minRetryDelay, maxRetryDelay), "noderoute"),
		ofClient:         client,
		ovsBridgeClient:  ovsBridgeClient,
		interfaceStore:   interfaceStore,
		nodeConfig:       config,
		gatewayLink:      link,
		installedNodes:   &sync.Map{},
		tunnelType:       tunnelType,
		ipsecPSK:         ipsecPSK,
		routeClient:      routeClient,
	}
	nodeInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(cur interface{}) {
				controller.enqueueNode(cur)
			},
			UpdateFunc: func(old, cur interface{}) {
				controller.enqueueNode(cur)
			},
			DeleteFunc: func(old interface{}) {
				controller.enqueueNode(old)
			},
		},
		nodeSyncPeriod,
	)
	return controller
}

// enqueueNode adds an object to the controller work queue
// obj could be an *v1.Node, or a DeletionFinalStateUnknown item.
func (c *Controller) enqueueNode(obj interface{}) {
	node, isNode := obj.(*v1.Node)
	if !isNode {
		deletedState, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			klog.Errorf("Received unexpected object: %v", obj)
			return
		}
		node, ok = deletedState.Obj.(*v1.Node)
		if !ok {
			klog.Errorf("DeletedFinalStateUnknown contains non-Node object: %v", deletedState.Obj)
			return
		}
	}

	// Ignore notifications for this Node, no need to establish connectivity to itself.
	if node.Name != c.nodeConfig.Name {
		c.queue.Add(node.Name)
	}
}

// listRoutes reads all the routes for the gateway and returns them as a map with the destination
// subnet as the key.
func (c *Controller) listRoutes() (routes map[string]*netlink.Route, err error) {
	routes = make(map[string]*netlink.Route)
	routeList, err := netlink.RouteList(c.gatewayLink, netlink.FAMILY_V4)
	if err != nil {
		return routes, err
	}
	for idx := 0; idx < len(routeList); idx++ {
		route := &routeList[idx]
		routes[route.Dst.String()] = route
	}
	return routes, nil
}

// removeStaleGatewayRoutes removes all the gateway routes which no longer correspond to a Node in
// the cluster. If the antrea agent restarts and Nodes have left the cluster, this function will
// take care of removing routes which are no longer valid.
func (c *Controller) removeStaleGatewayRoutes() error {
	nodes, err := c.nodeLister.List(labels.Everything())
	if err != nil {
		return fmt.Errorf("error when listing Nodes: %v", err)
	}

	routes, err := c.listRoutes()
	if err != nil {
		return fmt.Errorf("error when listing existing gateway routes: %v", err)
	}

	for _, node := range nodes {
		_, podCIDR, err := net.ParseCIDR(node.Spec.PodCIDR)
		if err != nil {
			return fmt.Errorf("error when parsing Pod CIDR for Node %s: %v", node.Name, err)
		}
		podCIDRStr := podCIDR.String()
		// remove from map of routes: at the end of the loop the remaining routes will be
		// the stale ones (the ones we want to delete).
		// we iterate over all current Nodes, including the Node on which this agent is
		// running, so the route to local Pods will be removed from the map as well.
		delete(routes, podCIDRStr)
	}

	// remove all remaining routes in the map: these routes do not match any Node in the
	// cluster.
	klog.V(4).Infof("Need to delete %d routes", len(routes))
	for _, route := range routes {
		if err := netlink.RouteDel(route); err != nil {
			klog.Errorf("error when deleting route '%v': %v", route, err)
		}
	}

	return nil
}

// removeStaleTunnelPorts removes all the tunnel ports which no longer correspond to a Node in the
// cluster. If the antrea agent restarts and Nodes have left the cluster, this function will take
// care of removing tunnel ports which are no longer valid. If the tunnel port configuration has
// changed, the tunnel port will also be deleted (the controller loop will later take care of
// re-creating the port with the correct configuration).
func (c *Controller) removeStaleTunnelPorts() error {
	nodes, err := c.nodeLister.List(labels.Everything())
	if err != nil {
		return fmt.Errorf("error when listing Nodes: %v", err)
	}
	// desiredInterfaces is the set of interfaces we wish to have, based on the current list of
	// Nodes. If a tunnel port corresponds to a valid Node but its configuration is wrong, we
	// will not include it in the set.
	desiredInterfaces := make(map[string]bool)
	// knownInterfaces is the list of interfaces currently in the local cache.
	knownInterfaces := c.interfaceStore.GetInterfaceKeysByType(interfacestore.TunnelInterface)

	if c.ipsecPSK != "" {
		for _, node := range nodes {
			interfaceConfig, found := c.interfaceStore.GetNodeTunnelInterface(node.Name)
			if !found {
				// Tunnel port not created for this Node, nothing to do.
				continue
			}

			peerNodeIP, err := getNodeAddr(node)
			if err != nil {
				klog.Errorf("Failed to retrieve IP address of Node %s: %v", node.Name, err)
				continue
			}

			ifaceID := util.GenerateNodeTunnelInterfaceKey(node.Name)
			validConfiguration := interfaceConfig.PSK == c.ipsecPSK &&
				interfaceConfig.RemoteIP.Equal(peerNodeIP) &&
				interfaceConfig.TunnelInterfaceConfig.Type == c.tunnelType
			if validConfiguration {
				desiredInterfaces[ifaceID] = true
			}
		}
	}

	// remove all ports which are no longer needed or for which the configuration is no longer
	// valid.
	for _, ifaceID := range knownInterfaces {
		if _, found := desiredInterfaces[ifaceID]; found {
			// this interface matches an existing Node, nothing to do.
			continue
		}
		interfaceConfig, found := c.interfaceStore.GetInterface(ifaceID)
		if !found {
			// should not happen, nothing should have concurrent access to the interface
			// store for tunnel interfaces.
			klog.Errorf("Interface %s can no longer be found in the interface store", ifaceID)
			continue
		}
		if interfaceConfig.InterfaceName == types.DefaultTunPortName {
			continue
		}
		if err := c.ovsBridgeClient.DeletePort(interfaceConfig.PortUUID); err != nil {
			klog.Errorf("Failed to delete OVS tunnel port %s: %v", interfaceConfig.InterfaceName, err)
		} else {
			c.interfaceStore.DeleteInterface(interfaceConfig)
		}
	}

	return nil
}

func (c *Controller) reconcile() error {
	klog.Infof("Reconciliation for %s", controllerName)
	// reconciliation consists of removing stale routes and stale / invalid tunnel ports:
	// missing routes and tunnel ports will be added normally by processNextWorkItem, which will
	// also take care of updating incorrect routes.
	if err := c.removeStaleGatewayRoutes(); err != nil {
		return fmt.Errorf("error when removing stale routes: %v", err)
	}
	if err := c.removeStaleTunnelPorts(); err != nil {
		return fmt.Errorf("error when removing stale tunnel ports: %v", err)
	}
	return nil
}

// Run will create defaultWorkers workers (go routines) which will process the Node events from the
// workqueue.
func (c *Controller) Run(stopCh <-chan struct{}) {
	defer c.queue.ShutDown()

	klog.Infof("Starting %s", controllerName)
	defer klog.Infof("Shutting down %s", controllerName)

	klog.Infof("Waiting for caches to sync for %s", controllerName)
	if !cache.WaitForCacheSync(stopCh, c.nodeListerSynced) {
		klog.Errorf("Unable to sync caches for %s", controllerName)
		return
	}
	klog.Infof("Caches are synced for %s", controllerName)

	if err := c.reconcile(); err != nil {
		klog.Errorf("Error during %s reconciliation", controllerName)
	}

	for i := 0; i < defaultWorkers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}
	<-stopCh
}

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the workqueue.
func (c *Controller) worker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem processes an item in the "node" work queue, by calling syncNodeRoute after
// casting the item to a string (Node name). If syncNodeRoute returns an error, this function
// handles it by requeueing the item so that it can be processed again later. If syncNodeRoute is
// successful, the Node is removed from the queue until we get notified of a new change. This
// function returns false if and only if the work queue was shutdown (no more items will be
// processed).
func (c *Controller) processNextWorkItem() bool {
	obj, quit := c.queue.Get()
	if quit {
		return false
	}
	// We call Done here so the workqueue knows we have finished processing this item. We also
	// must remember to call Forget if we do not want this work item being re-queued. For
	// example, we do not call Forget if a transient error occurs, instead the item is put back
	// on the workqueue and attempted again after a back-off period.
	defer c.queue.Done(obj)

	// We expect strings (Node name) to come off the workqueue.
	if key, ok := obj.(string); !ok {
		// As the item in the workqueue is actually invalid, we call Forget here else we'd
		// go into a loop of attempting to process a work item that is invalid.
		// This should not happen: enqueueNode only enqueues strings.
		c.queue.Forget(obj)
		klog.Errorf("Expected string in work queue but got %#v", obj)
		return true
	} else if err := c.syncNodeRoute(key); err == nil {
		// If no error occurs we Forget this item so it does not get queued again until
		// another change happens.
		c.queue.Forget(key)
	} else {
		// Put the item back on the workqueue to handle any transient errors.
		c.queue.AddRateLimited(key)
		klog.Errorf("Error syncing Node %s, requeuing. Error: %v", key, err)
	}
	return true
}

// syncNode manages connectivity to "peer" Node with name nodeName
// If we have not established connectivity to the Node yet:
//   * we install the appropriate Linux route:
// Destination     Gateway         Use Iface
// peerPodCIDR     peerGatewayIP   localGatewayIface (e.g gw0)
//   * we install the appropriate OpenFlow flows to ensure that all the traffic destined to
//   peerPodCIDR goes through the correct L3 tunnel.
// If the Node no longer exists (cannot be retrieved by name from nodeLister) we delete the route
// and OpenFlow flows associated with it.
func (c *Controller) syncNodeRoute(nodeName string) error {
	startTime := time.Now()
	defer func() {
		klog.V(4).Infof("Finished syncing Node Route for %s. (%v)", nodeName, time.Since(startTime))
	}()

	// The work queue guarantees that concurrent goroutines cannot call syncNodeRoute on the
	// same Node, which is required by the InstallNodeFlows / UninstallNodeFlows OF Client
	// methods.

	if node, err := c.nodeLister.Get(nodeName); err != nil {
		return c.deleteNodeRoute(nodeName, node)
	} else {
		return c.addNodeRoute(nodeName, node)
	}
}

func (c *Controller) deleteNodeRoute(nodeName string, node *v1.Node) error {
	klog.Infof("Deleting routes and flows to Node %s", nodeName)

	entry, flowsAreInstalled := c.installedNodes.Load(nodeName)
	routes := entry.([]*netlink.Route)
	for _, r := range routes {
		if err := netlink.RouteDel(r); err != nil {
			return fmt.Errorf("failed to delete the route to Node %s: %v", nodeName, err)
		}
	}
	c.installedNodes.Store(nodeName, nil)

	if flowsAreInstalled {
		if err := c.ofClient.UninstallNodeFlows(nodeName); err != nil {
			return fmt.Errorf("failed to uninstall flows to Node %s: %v", nodeName, err)
		}
		c.installedNodes.Delete(nodeName)
	} else if entry, flowsAreInstalled = c.installedNodes.Load(nodeName); entry == nil {
		klog.Infof("Adding routes and flows to Node %s, podCIDR: %s, addresses: %v",
			nodeName, node.Spec.PodCIDR, node.Status.Addresses)
		if node.Spec.PodCIDR == "" {
			klog.V(1).Infof("PodCIDR is empty for peer node %s", nodeName)
		}
	}

	if c.ipsecPSK != "" {
		interfaceConfig, ok := c.interfaceStore.GetNodeTunnelInterface(nodeName)
		if !ok {
			// Tunnel port not created for this Node.
			return nil
		}
		if err := c.ovsBridgeClient.DeletePort(interfaceConfig.PortUUID); err != nil {
			klog.Errorf("Failed to delete OVS tunnel port %s for Node %s: %v",
				interfaceConfig.InterfaceName, nodeName, err)
			return fmt.Errorf("failed to delete OVS tunnel port for Node %s", nodeName)
		}
		c.interfaceStore.DeleteInterface(interfaceConfig)
	}
	return nil
}

func (c *Controller) addNodeRoute(nodeName string, node *v1.Node) error {
	entry, flowsAreInstalled := c.installedNodes.Load(nodeName)
	if entry != nil {
		// Route is already added for this Node.
		return nil
	}

	klog.Infof("Adding routes and flows to Node %s, podCIDR: %s, addresses: %v",
		nodeName, node.Spec.PodCIDR, node.Status.Addresses)

	if node.Spec.PodCIDR == "" {
		klog.Errorf("PodCIDR is empty for Node %s", nodeName)
		// Does not help to return an error and trigger controller retries.
		return nil
	}
	peerPodCIDRAddr, peerPodCIDR, err := net.ParseCIDR(node.Spec.PodCIDR)
	if err != nil {
		klog.Errorf("Failed to parse PodCIDR %s for Node %s", node.Spec.PodCIDR, nodeName)
		return nil
	}
	peerNodeIP, err := getNodeAddr(node)
	if err != nil {
		klog.Errorf("Failed to retrieve IP address of Node %s: %v", nodeName, err)
		return nil
	}
	peerGatewayIP := ip.NextIP(peerPodCIDRAddr)

	var tunOFPort int32
	remoteIP := peerNodeIP
	if c.nodeConfig.PodEncapMode.UseTunnel(peerNodeIP, c.nodeConfig.NodeIPAddr) {
		if c.ipsecPSK != "" {
			// Create a separate tunnel port for the Node, as OVS does not support flow
			// based tunnel for IPSec.
			if tunOFPort, err = c.createIPSecTunnelPort(nodeName, peerNodeIP); err != nil {
				return err
			}
			remoteIP = nil
		} else {
			// Use the default tunnel port.
			tunOFPort = types.DefaultTunOFPort
		}
	}

	if !flowsAreInstalled { // then install flows
		err = c.ofClient.InstallNodeFlows(
			nodeName,
			c.nodeConfig.GatewayConfig.MAC,
			peerGatewayIP,
			*peerPodCIDR,
			remoteIP,
			uint32(tunOFPort))
		if err != nil {
			return fmt.Errorf("failed to install flows to Node %s: %v", nodeName, err)
		}
		c.installedNodes.Store(nodeName, nil)
	}

	routes, err := c.routeClient.AddPeerCIDRRoute(peerPodCIDR, c.gatewayLink.Attrs().Index, peerNodeIP, peerGatewayIP)
	c.installedNodes.Store(nodeName, routes)
	return err
}

// createIPSecTunnelPort creates an IPSec tunnel port for the remote Node if the
// tunnel does not exist, and returns the ofport number.
func (c *Controller) createIPSecTunnelPort(nodeName string, nodeIP net.IP) (int32, error) {
	interfaceConfig, ok := c.interfaceStore.GetNodeTunnelInterface(nodeName)
	if ok {
		// TODO: check if Node IP, PSK, or tunnel type changes. This can
		// happen if removeStaleTunnelPorts fails to remove a "stale"
		// tunnel port for which the configuration has changed.
		if interfaceConfig.OFPort != 0 {
			return interfaceConfig.OFPort, nil
		}
	} else {
		portName := util.GenerateNodeTunnelInterfaceName(nodeName)
		ovsExternalIDs := map[string]interface{}{ovsExternalIDNodeName: nodeName}
		portUUID, err := c.ovsBridgeClient.CreateTunnelPortExt(
			portName,
			c.tunnelType,
			0, // ofPortRequest - let OVS allocate OFPort number.
			nodeIP.String(),
			c.ipsecPSK,
			ovsExternalIDs)
		if err != nil {
			klog.Errorf("Failed to create OVS IPSec tunnel port for Node %s: %v", nodeName, err)
			return 0, fmt.Errorf("failed to create IPSec tunnel port for Node %s", nodeName)
		}
		klog.Infof("Created IPSec tunnel port %s for Node %s", portName, nodeName)

		ovsPortConfig := &interfacestore.OVSPortConfig{PortUUID: portUUID}
		interfaceConfig = interfacestore.NewIPSecTunnelInterface(
			portName,
			c.tunnelType,
			nodeName,
			nodeIP,
			c.ipsecPSK)
		interfaceConfig.OVSPortConfig = ovsPortConfig
		c.interfaceStore.AddInterface(interfaceConfig)
	}

	// GetOFPort will wait for up to 1 second for OVSDB to report the OFPort number.
	ofPort, err := c.ovsBridgeClient.GetOFPort(interfaceConfig.InterfaceName)
	if err != nil {
		// Could be a temporary OVSDB connection failure or timeout.
		// Let NodeRouteController retry at errors.
		klog.Errorf("Failed to get of_port of the tunnel port for Node %s: %v", nodeName, err)
		return 0, fmt.Errorf("failed to get of_port of IPSec tunnel port for Node %s", nodeName)
	}
	interfaceConfig.OFPort = ofPort
	return ofPort, nil
}

// ParseTunnelInterfaceConfig initializes and returns an InterfaceConfig struct
// for a tunnel interface. It reads tunnel type, remote IP, IPSec PSK from the
// OVS interface options, and NodeName from the OVS port external_ids.
// nil is returned, if the OVS port and interface configurations are not valid
// for a tunnel interface.
func ParseTunnelInterfaceConfig(
	portData *ovsconfig.OVSPortData,
	portConfig *interfacestore.OVSPortConfig) *interfacestore.InterfaceConfig {
	if portData.Options == nil {
		klog.V(2).Infof("OVS port %s has no options", portData.Name)
		return nil
	}
	remoteIP, psk := ovsconfig.ParseTunnelInterfaceOptions(portData)

	var interfaceConfig *interfacestore.InterfaceConfig
	var nodeName string
	if portData.ExternalIDs != nil {
		nodeName = portData.ExternalIDs[ovsExternalIDNodeName]
	}
	if psk != "" {
		interfaceConfig = interfacestore.NewIPSecTunnelInterface(
			portData.Name,
			ovsconfig.TunnelType(portData.IFType),
			nodeName,
			remoteIP,
			psk)
	} else {
		interfaceConfig = interfacestore.NewTunnelInterface(portData.Name, ovsconfig.TunnelType(portData.IFType))
	}
	interfaceConfig.OVSPortConfig = portConfig
	return interfaceConfig
}

// getNodeAddr gets the available IP address of a Node. getNodeAddr will first try to get the
// NodeInternalIP, then try to get the NodeExternalIP.
func getNodeAddr(node *v1.Node) (net.IP, error) {
	addresses := make(map[v1.NodeAddressType]string)
	for _, addr := range node.Status.Addresses {
		addresses[addr.Type] = addr.Address
	}
	var ipAddrStr string
	if internalIp, ok := addresses[v1.NodeInternalIP]; ok {
		ipAddrStr = internalIp
	} else if externalIp, ok := addresses[v1.NodeExternalIP]; ok {
		ipAddrStr = externalIp
	} else {
		return nil, fmt.Errorf("Node %s has neither external ip nor internal ip", node.Name)
	}
	ipAddr := net.ParseIP(ipAddrStr)
	if ipAddr == nil {
		return nil, fmt.Errorf("<%v> is not a valid ip address", ipAddrStr)
	}
	return ipAddr, nil
}
