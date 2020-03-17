// Copyright 2020 Antrea Authors
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

package route

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"

	"github.com/vmware-tanzu/antrea/pkg/agent/config"
	"github.com/vmware-tanzu/antrea/pkg/agent/openflow"
	"github.com/vmware-tanzu/antrea/pkg/agent/util/ipset"
	"github.com/vmware-tanzu/antrea/pkg/agent/util/iptables"
)

const (
	// AntreaServiceTable is route table name for Antrea service traffic.
	AntreaServiceTable = "Antrea-service"
	// AntreaServiceTableIdx is route table index for Antrea service traffic.
	AntreaServiceTableIdx = 300
	mainTable             = "main"
	mainTableIdx          = 254

	routeTableConfigPath = "/etc/iproute2/rt_tables"
	// AntreaIPRulePriority is Antrea IP rule priority
	AntreaIPRulePriority = 300

	// Antrea managed ipset.
	// antreaPodIPSet contains all Pod CIDRs of this cluster.
	antreaPodIPSet = "ANTREA-POD-IP"

	// Antrea managed iptables chains.
	antreaForwardChain     = "ANTREA-FORWARD"
	antreaPostRoutingChain = "ANTREA-POSTROUTING"
	antreaMangleChain      = "ANTREA-MANGLE"
	antreaRawChain         = "ANTREA-RAW"
)

var (
	// RtTblSelectorValue selects which route table to use to forward service traffic back to host gateway gw0.
	RtTblSelectorValue = 1 << 11
	// the type of ipRule.Mask is int so 0xffffffff cannot be used on 32-bit
	// architectures (e.g. arm/v7). Using -1 does not seem to work (mismatch
	// when reading the route back, we get 0xffffffff).
	// TODO: it should be fixed in netlink package, see https://github.com/vishvananda/netlink/issues/528
	// RtTblSelectorMask  = 0xffffffff
	// RtTblSelectorMask = -1
	RtTblSelectorMask = 0x8fffffff
	rtTblSelectorMark = fmt.Sprintf("%#x/%#x", RtTblSelectorValue, RtTblSelectorValue)
)

// Client implements Interface.
var _ Interface = &Client{}

// Client takes care of routing container packets in host network, coordinating ip route, ip rule, iptables and ipset.
type Client struct {
	nodeConfig  *config.NodeConfig
	encapMode   config.TrafficEncapModeType
	hostGateway string
	serviceCIDR *net.IPNet
	ipt         *iptables.Client
	// serviceRtTable contains Antrea service route table information.
	serviceRtTable *serviceRtTableConfig
	// nodeRoutes caches ip routes to remote Pods. It's a map of podCIDR to routes.
	nodeRoutes sync.Map
}

type serviceRtTableConfig struct {
	Idx  int
	Name string
}

func (s *serviceRtTableConfig) String() string {
	return fmt.Sprintf("%s: idx %d", s.Name, s.Idx)
}

func (s *serviceRtTableConfig) IsMainTable() bool {
	return s.Name == "main"
}

// NewClient returns a route client.
func NewClient(hostGateway string, serviceCIDR *net.IPNet, encapMode config.TrafficEncapModeType) (*Client, error) {
	ipt, err := iptables.New()
	if err != nil {
		return nil, fmt.Errorf("error creating IPTables instance: %v", err)
	}

	serviceRtTable := &serviceRtTableConfig{Idx: mainTableIdx, Name: mainTable}
	if encapMode.SupportsNoEncap() {
		serviceRtTable.Idx = AntreaServiceTableIdx
		serviceRtTable.Name = AntreaServiceTable
	}

	return &Client{
		hostGateway:    hostGateway,
		serviceCIDR:    serviceCIDR,
		encapMode:      encapMode,
		ipt:            ipt,
		serviceRtTable: serviceRtTable,
	}, nil
}

// Initialize initializes all infrastructures required to route container packets in host network.
// It is idempotent and can be safely called on every startup.
func (c *Client) Initialize(nodeConfig *config.NodeConfig) error {
	c.nodeConfig = nodeConfig

	// Sets up the ipset that will be used in iptables.
	if err := c.initIPSet(); err != nil {
		return fmt.Errorf("failed to initialize ipset: %v", err)
	}

	// Sets up the iptables infrastructure required to route packets in host network.
	if err := c.initIPTables(); err != nil {
		return fmt.Errorf("failed to initialize iptables: %v", err)
	}

	// Sets up the IP routes and IP rule required to route packets in host network.
	if err := c.initIPRoutes(); err != nil {
		return fmt.Errorf("failed to initialize ip routes: %v", err)
	}

	// send_redirects must be disabled because packets from hostGateway are
	// routed back to it. Otherwise redirect packets will be sent to source
	// Pods.
	// send_redirects for the interface will be enabled if at least one of
	// conf/{all,interface}/send_redirects is set to TRUE, so "all" and the
	// interface must be disabled together.
	// See https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt.
	if err := disableICMPSendRedirects("all"); err != nil {
		return err
	}
	if err := disableICMPSendRedirects(c.hostGateway); err != nil {
		return err
	}
	return nil
}

// initIPSet ensures that the required ipset exists and it has the initial members.
func (c *Client) initIPSet() error {
	if err := ipset.CreateIPSet(antreaPodIPSet, ipset.HashNet); err != nil {
		return err
	}
	// Ensure its own PodCIDR is in it.
	if err := ipset.AddEntry(antreaPodIPSet, c.nodeConfig.PodCIDR.String()); err != nil {
		return err
	}
	return nil
}

// initIPTables ensure that the iptables infrastructure we use is set up.
// It's idempotent and can safely be called on every startup.
func (c *Client) initIPTables() error {
	// Create the antrea managed chains and link them to built-in chains.
	// We cannot use iptables-restore for these jump rules because there
	// are non antrea managed rules in built-in chains.
	jumpRules := []struct{ table, srcChain, dstChain, comment string }{
		{iptables.FilterTable, iptables.ForwardChain, antreaForwardChain, "Antrea: jump to Antrea forwarding rules"},
		{iptables.NATTable, iptables.PostRoutingChain, antreaPostRoutingChain, "Antrea: jump to Antrea postrouting rules"},
		{iptables.MangleTable, iptables.PreRoutingChain, antreaMangleChain, "Antrea: jump to Antrea mangle rules"},
		{iptables.RawTable, iptables.PreRoutingChain, antreaRawChain, "Antrea: jump to Antrea raw rules"},
	}
	for _, rule := range jumpRules {
		if err := c.ipt.EnsureChain(rule.table, rule.dstChain); err != nil {
			return err
		}
		ruleSpec := []string{"-j", rule.dstChain, "-m", "comment", "--comment", rule.comment}
		if err := c.ipt.EnsureRule(rule.table, rule.srcChain, ruleSpec); err != nil {
			return err
		}
	}

	// Create required rules in the antrea chains.
	// Use iptables-restore as it flushes the involved chains and creates the desired rules
	// with a single call, instead of string matching to clean up stale rules.
	iptablesData := bytes.NewBuffer(nil)
	// Write head lines anyway so the undesired rules can be deleted when noEncap -> encap.
	writeLine(iptablesData, "*mangle")
	writeLine(iptablesData, iptables.MakeChainLine(antreaMangleChain))
	if c.encapMode.SupportsNoEncap() {
		writeLine(iptablesData, []string{
			"-A", antreaMangleChain,
			"-m", "comment", "--comment", `"Antrea: mark pod to service packets"`,
			"-i", c.hostGateway, "-d", c.serviceCIDR.String(),
			"-j", iptables.MarkTarget, "--set-xmark", rtTblSelectorMark,
		}...)
	}
	writeLine(iptablesData, "COMMIT")

	writeLine(iptablesData, "*filter")
	writeLine(iptablesData, iptables.MakeChainLine(antreaForwardChain))
	writeLine(iptablesData, []string{
		"-A", antreaForwardChain,
		"-m", "comment", "--comment", `"Antrea: accept packets from local pods"`,
		"-i", c.hostGateway,
		"-j", iptables.AcceptTarget,
	}...)
	writeLine(iptablesData, []string{
		"-A", antreaForwardChain,
		"-m", "comment", "--comment", `"Antrea: accept packets to local pods"`,
		"-o", c.hostGateway,
		"-j", iptables.AcceptTarget,
	}...)
	writeLine(iptablesData, "COMMIT")

	writeLine(iptablesData, "*nat")
	writeLine(iptablesData, iptables.MakeChainLine(antreaPostRoutingChain))
	writeLine(iptablesData, []string{
		"-A", antreaPostRoutingChain,
		"-m", "comment", "--comment", `"Antrea: masquerade pod to external packets"`,
		"-s", c.nodeConfig.PodCIDR.String(), "-m", "set", "!", "--match-set", antreaPodIPSet, "dst",
		"-j", iptables.MasqueradeTarget,
	}...)
	writeLine(iptablesData, "COMMIT")

	writeLine(iptablesData, "*raw")
	writeLine(iptablesData, iptables.MakeChainLine(antreaRawChain))
	if c.encapMode.SupportsNoEncap() {
		writeLine(iptablesData, []string{
			"-A", antreaRawChain,
			"-m", "comment", "--comment", `"Antrea: reentry pod traffic skip conntrack"`,
			"-i", c.hostGateway, "-m", "mac", "--mac-source", openflow.ReentranceMAC.String(),
			"-j", iptables.ConnTrackTarget, "--notrack",
		}...)
	}
	writeLine(iptablesData, "COMMIT")

	// Setting --noflush to keep the previous contents (i.e. non antrea managed chains) of the tables.
	if err := c.ipt.Restore(iptablesData.Bytes(), false); err != nil {
		return err
	}
	return nil
}

func (c *Client) initIPRoutes() error {
	if c.serviceRtTable.IsMainTable() {
		_ = c.removeServiceRouting()
		return nil
	}
	return c.addServiceRouting()
}

// Reconcile removes orphaned podCIDRs from ipset and removes routes to orphaned podCIDRs
// based on the desired podCIDRs.
func (c *Client) Reconcile(podCIDRs []string) error {
	desiredPodCIDRs := sets.NewString(podCIDRs...)

	// Remove orphaned podCIDRs from antreaPodIPSet.
	entries, err := ipset.ListEntries(antreaPodIPSet)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if desiredPodCIDRs.Has(entry) {
			continue
		}
		klog.V(4).Infof("Deleting orphaned ip %s from ipset", entry)
		if err := ipset.DelEntry(antreaPodIPSet, entry); err != nil {
			return err
		}
	}

	// Remove orphaned routes from host network.
	actualRouteMap, err := c.listIPRoutes()
	if err != nil {
		return fmt.Errorf("error listing ip routes: %v", err)
	}
	for podCIDR, actualRoutes := range actualRouteMap {
		if desiredPodCIDRs.Has(podCIDR) {
			continue
		}
		for _, route := range actualRoutes {
			klog.V(4).Infof("Deleting orphaned route %v", route)
			if err := netlink.RouteDel(route); err != nil && err != unix.ESRCH {
				return err
			}
		}
	}
	return nil
}

// AddRoutes adds routes to a new podCIDR. It overrides the routes if they already exist.
func (c *Client) AddRoutes(podCIDR *net.IPNet, nodeIP, nodeGwIP net.IP) error {
	podCIDRStr := podCIDR.String()
	// Add this podCIDR to antreaPodIPSet so that packets to them won't be masqueraded when they leave the host.
	if err := ipset.AddEntry(antreaPodIPSet, podCIDRStr); err != nil {
		return err
	}

	// Install routes to this Node.
	routes := []*netlink.Route{
		{
			Dst:       podCIDR,
			Flags:     int(netlink.FLAG_ONLINK),
			LinkIndex: c.nodeConfig.GatewayConfig.LinkIndex,
			Gw:        nodeGwIP,
			Table:     c.serviceRtTable.Idx,
		},
	}

	// If service route table and main route table is not the same , add
	// peer CIDR to main route table too (i.e in NoEncap and hybrid mode)
	if !c.serviceRtTable.IsMainTable() {
		if c.encapMode.NeedsEncapToPeer(nodeIP, c.nodeConfig.NodeIPAddr) {
			// need overlay tunnel
			routes = append(routes, &netlink.Route{
				Dst:       podCIDR,
				Flags:     int(netlink.FLAG_ONLINK),
				LinkIndex: c.nodeConfig.GatewayConfig.LinkIndex,
				Gw:        nodeGwIP,
			})
		} else if !c.encapMode.NeedsRoutingToPeer(nodeIP, c.nodeConfig.NodeIPAddr) {
			routes = append(routes, &netlink.Route{
				Dst: podCIDR,
				Gw:  nodeIP,
			})
		}
		// If Pod traffic needs underlying routing support, it is handled by host default route.
	}

	// clean up function if any route add failed
	deleteRtFn := func() {
		for _, route := range routes {
			_ = netlink.RouteDel(route)
		}
	}

	for _, route := range routes {
		if err := netlink.RouteReplace(route); err != nil {
			deleteRtFn()
			return fmt.Errorf("failed to install route to peer %s with netlink: %v", nodeIP, err)
		}
	}
	c.nodeRoutes.Store(podCIDRStr, routes)
	return nil
}

// DeleteRoutes deletes routes to a PodCIDR. It does nothing if the routes doesn't exist.
func (c *Client) DeleteRoutes(podCIDR *net.IPNet) error {
	podCIDRStr := podCIDR.String()
	// Delete this podCIDR from antreaPodIPSet as the CIDR is no longer for Pods.
	if err := ipset.DelEntry(antreaPodIPSet, podCIDRStr); err != nil {
		return err
	}

	routes, exists := c.nodeRoutes.Load(podCIDRStr)
	if !exists {
		return nil
	}
	for _, r := range routes.([]*netlink.Route) {
		klog.V(4).Infof("Deleting route %v", r)
		if err := netlink.RouteDel(r); err != nil && err != unix.ESRCH {
			return err
		}
	}
	c.nodeRoutes.Delete(podCIDRStr)
	return nil
}

// listIPRoutes returns list of routes from peer and local CIDRs
func (c *Client) listIPRoutes() (map[string][]*netlink.Route, error) {
	// get all routes on gw0 from service table.
	filter := &netlink.Route{
		Table:     c.serviceRtTable.Idx,
		LinkIndex: c.nodeConfig.GatewayConfig.LinkIndex}
	routes, err := netlink.RouteListFiltered(netlink.FAMILY_V4, filter, netlink.RT_FILTER_TABLE|netlink.RT_FILTER_OIF)
	if err != nil {
		return nil, err
	}

	rtMap := make(map[string][]*netlink.Route)
	for _, rt := range routes {
		// rt is reference to actual data, as it changes,
		// it cannot be used for assignment
		tmpRt := rt
		rtMap[rt.Dst.String()] = append(rtMap[rt.Dst.String()], &tmpRt)
	}

	if !c.serviceRtTable.IsMainTable() {
		// get all routes on gw0 from main table.
		filter.Table = 0
		routes, err := netlink.RouteListFiltered(netlink.FAMILY_V4, filter, netlink.RT_FILTER_OIF)
		if err != nil {
			return nil, err
		}
		for _, rt := range routes {
			// rt is reference to actual data, as it changes,
			// it cannot be used for assignment
			tmpRt := rt
			rtMap[rt.Dst.String()] = append(rtMap[rt.Dst.String()], &tmpRt)
		}

		// now get all routes gw0 on other interfaces from main table.
		routes, err = netlink.RouteListFiltered(netlink.FAMILY_V4, nil, 0)
		if err != nil {
			return nil, err
		}
		for _, rt := range routes {
			if rt.Dst == nil {
				continue
			}
			// insert the route if it is CIDR route and has not been added already.
			// routes with same dst are different if table or linkIndex differs.
			if rl, ok := rtMap[rt.Dst.String()]; ok && (rl[len(rl)-1].LinkIndex != rt.LinkIndex || rl[len(rl)-1].Table != rt.Table) {
				tmpRt := rt
				rtMap[rt.Dst.String()] = append(rl, &tmpRt)
			}
		}
	}
	return rtMap, nil
}

func (c *Client) addServiceRouting() error {
	f, err := os.OpenFile(routeTableConfigPath, os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		return fmt.Errorf("unable to create service route table(open): %v", err)
	}
	defer f.Close()

	oldTablesRaw := make([]byte, 1024)
	bLen, err := f.Read(oldTablesRaw)
	if err != nil {
		return fmt.Errorf("unable to create service route table(read): %v", err)
	}
	oldTables := string(oldTablesRaw[:bLen])
	newTable := fmt.Sprintf("%d %s", c.serviceRtTable.Idx, c.serviceRtTable.Name)

	if strings.Index(oldTables, newTable) == -1 {
		if _, err := f.WriteString(newTable); err != nil {
			return fmt.Errorf("failed to add antrea service route table: %v", err)
		}
	}

	gwConfig := c.nodeConfig.GatewayConfig
	if gwConfig != nil && c.nodeConfig.PodCIDR != nil {
		// Add local podCIDR if applicable to service rt table.
		route := &netlink.Route{
			LinkIndex: gwConfig.LinkIndex,
			Scope:     netlink.SCOPE_LINK,
			Dst:       c.nodeConfig.PodCIDR,
			Table:     c.serviceRtTable.Idx,
		}
		if err := netlink.RouteReplace(route); err != nil {
			return fmt.Errorf("failed to add link route to service table: %v", err)
		}
	}

	// create ip rule to select route table
	ipRule := netlink.NewRule()
	ipRule.IifName = c.nodeConfig.GatewayConfig.Name
	ipRule.Mark = RtTblSelectorValue
	ipRule.Mask = RtTblSelectorMask
	ipRule.Table = c.serviceRtTable.Idx
	ipRule.Priority = AntreaIPRulePriority

	ruleList, err := netlink.RuleList(netlink.FAMILY_V4)
	if err != nil {
		return fmt.Errorf("failed to get ip rule: %v", err)
	}
	// Check for ip rule presence.
	for _, rule := range ruleList {
		if rule == *ipRule {
			return nil
		}
	}
	err = netlink.RuleAdd(ipRule)
	if err != nil {
		return fmt.Errorf("failed to create ip rule for service route table: %v", err)
	}
	return nil
}

func (c *Client) readRtTable() (string, error) {
	f, err := os.OpenFile(routeTableConfigPath, os.O_RDONLY, 0)
	if err != nil {
		return "", fmt.Errorf("route table(open): %w", err)
	}
	defer f.Close()

	tablesRaw := make([]byte, 1024)
	bLen, err := f.Read(tablesRaw)
	if err != nil {
		return "", fmt.Errorf("route table(read): %w", err)
	}
	return string(tablesRaw[:bLen]), nil
}

// removeServiceRouting removes service routing setup.
func (c *Client) removeServiceRouting() error {
	// remove service table
	tables, err := c.readRtTable()
	if err != nil {
		return err
	}
	newTable := fmt.Sprintf("%d %s", AntreaServiceTableIdx, AntreaServiceTable)
	if strings.Index(tables, newTable) != -1 {
		tables = strings.Replace(tables, newTable, "", -1)
		f, err := os.OpenFile(routeTableConfigPath, os.O_WRONLY|os.O_TRUNC, 0)
		if err != nil {
			return fmt.Errorf("route table(open): %w", err)
		}
		defer f.Close()
		if _, err = f.WriteString(tables); err != nil {
			return fmt.Errorf("route table(write): %w", err)
		}
	}

	// flush service table
	filter := &netlink.Route{
		Table:     AntreaServiceTableIdx,
		LinkIndex: c.nodeConfig.GatewayConfig.LinkIndex}
	routes, err := netlink.RouteListFiltered(netlink.FAMILY_V4, filter, netlink.RT_FILTER_TABLE|netlink.RT_FILTER_OIF)
	if err != nil {
		return fmt.Errorf("route table(list): %w", err)
	}
	for _, route := range routes {
		if err = netlink.RouteDel(&route); err != nil {
			return fmt.Errorf("route delete: %w", err)
		}
	}

	// delete ip rule for service table
	ipRule := netlink.NewRule()
	ipRule.IifName = c.nodeConfig.GatewayConfig.Name
	ipRule.Mark = RtTblSelectorValue
	ipRule.Mask = RtTblSelectorMask
	ipRule.Table = AntreaServiceTableIdx
	ipRule.Priority = AntreaIPRulePriority
	if err = netlink.RuleDel(ipRule); err != nil {
		return fmt.Errorf("ip rule delete: %w", err)
	}
	return nil
}

// Join all words with spaces, terminate with newline and write to buf.
func writeLine(buf *bytes.Buffer, words ...string) {
	// We avoid strings.Join for performance reasons.
	for i := range words {
		buf.WriteString(words[i])
		if i < len(words)-1 {
			buf.WriteByte(' ')
		} else {
			buf.WriteByte('\n')
		}
	}
}

func disableICMPSendRedirects(intfName string) error {
	cmdStr := fmt.Sprintf("echo 0 > /proc/sys/net/ipv4/conf/%s/send_redirects", intfName)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	if err := cmd.Run(); err != nil {
		klog.Errorf("Failed to disable send_redirect for interface %s: %v", intfName, err)
		return err
	}
	return nil
}
