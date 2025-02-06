// Copyright 2022 Antrea Authors
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

package e2esecondary

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"testing"
	"time"

	netattachv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	networkclient "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/client/clientset/versioned"
	logs "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	antreae2e "antrea.io/antrea/test/e2e"
)

type testPodInfo struct {
	podName  string
	nodeName string
	// map from interface name to secondary network name.
	interfaceNetworks map[string]string
	isSRIOV           bool
	secondaryENI      string
}

type testData struct {
	e2eTestData *antreae2e.TestData
	networkType string
	pods        []*testPodInfo
}

const (
	networkTypeSriov = "sriov"
	networkTypeVLAN  = "vlan"

	// Namespace of NetworkAttachmentDefinition CRs.
	attachDefNamespace = "default"

	containerName  = "toolbox"
	podApp         = "secondaryTest"
	osType         = "linux"
	pingCount      = 5
	pingSize       = 40
	defaultTimeout = 10 * time.Second
	sriovReqName   = "intel.com/intel_sriov_netdevice"
	sriovResNum    = 3
)

// formAnnotationStringOfPod forms the annotation string, used in the generation of each Pod YAML file.
func (data *testData) formAnnotationStringOfPod(pod *testPodInfo) string {
	var annotationString = ""
	for i, n := range pod.interfaceNetworks {
		if pod.isSRIOV {
			return n
		}
		podNetworkSpec := fmt.Sprintf("{\"name\": \"%s\", \"namespace\": \"%s\", \"interface\": \"%s\"}",
			n, attachDefNamespace, i)
		if annotationString == "" {
			annotationString = "[" + podNetworkSpec
		} else {
			annotationString = annotationString + ", " + podNetworkSpec
		}
	}
	annotationString = annotationString + "]"
	return annotationString
}

// createPodOnNode creates the Pod for the specific annotations as per the parsed Pod information using the NewPodBuilder API
func (data *testData) createPods(t *testing.T, ns string) error {
	var err error
	for _, pod := range data.pods {
		err := data.createPodForSecondaryNetwork(ns, pod)
		if err != nil {
			return fmt.Errorf("error in creating pods.., err: %v", err)
		}
	}
	return err
}

// The Wrapper function createPodForSecondaryNetwork creates the Pod adding the annotation, arguments, commands, Node, container name,
// resource requests and limits as arguments with the NewPodBuilder API
func (data *testData) createPodForSecondaryNetwork(ns string, pod *testPodInfo) error {
	podBuilder := antreae2e.NewPodBuilder(pod.podName, ns, antreae2e.ToolboxImage).
		OnNode(pod.nodeName).WithContainerName(containerName).
		WithAnnotations(map[string]string{
			"k8s.v1.cni.cncf.io/networks": fmt.Sprintf("%s", data.formAnnotationStringOfPod(pod)),
		}).
		WithLabels(map[string]string{
			"App": fmt.Sprintf("%s", podApp),
		})

	var resNum int64
	resNum = sriovResNum
	if pod.isSRIOV {
		resNum = 1
	}
	if data.networkType == networkTypeSriov {
		computeResources := resource.NewQuantity(resNum, resource.DecimalSI)
		podBuilder = podBuilder.WithResources(corev1.ResourceList{sriovReqName: *computeResources}, corev1.ResourceList{sriovReqName: *computeResources})
	}
	return podBuilder.Create(data.e2eTestData)
}

// listPodIPs returns a map of Pod IPs, indexed by the interface name. All interfaces are included
// and only IPv4 addresses are considered. If an interface is not assigned an IPv4 address, it will
// be included in the map, with a nil value.
func (data *testData) listPodIPs(targetPod *testPodInfo) (map[string]net.IP, error) {
	cmd := []string{"ip", "addr", "show"}
	stdout, _, err := data.e2eTestData.RunCommandFromPod(data.e2eTestData.GetTestNamespace(), targetPod.podName, containerName, cmd)
	if err != nil {
		return nil, fmt.Errorf("error when listing interfaces for %s: %w", targetPod.podName, err)
	}
	result := make(map[string]net.IP)
	var currentInterface string
	lines := strings.Split(stdout, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 && strings.HasSuffix(fields[0], ":") {
			// first field is ifindex, second field is interface name
			currentInterface = strings.Split(strings.TrimSuffix(fields[1], ":"), "@")[0]
			result[currentInterface] = nil
		} else if len(fields) >= 2 && fields[0] == "inet" {
			ipStr := strings.Split(fields[1], "/")[0]
			ip := net.ParseIP(ipStr)
			if ip == nil {
				return nil, fmt.Errorf("failed to parse IP (%s) for interface %s of Pod %s", ipStr, currentInterface, targetPod.podName)
			}
			result[currentInterface] = ip
		}
	}
	return result, nil
}

// pingBetweenInterfaces parses through all the created Pods and pings the other Pod if the two Pods
// both have a secondary network interface on the same network.
func (data *testData) pingBetweenInterfaces(t *testing.T) error {
	e2eTestData := data.e2eTestData
	namespace := e2eTestData.GetTestNamespace()

	type attachment struct {
		network string
		iface   string
		ip      net.IP
	}
	type network struct {
		// maps each Pod to its attachments in this network (typically just one)
		podAttachments map[*testPodInfo][]*attachment
	}
	networks := make(map[string]*network)
	addPodNetworkAttachments := func(pod *testPodInfo, podAttachments []*attachment) {
		for _, pa := range podAttachments {
			if _, ok := networks[pa.network]; !ok {
				networks[pa.network] = &network{
					podAttachments: make(map[*testPodInfo][]*attachment),
				}
			}
			networks[pa.network].podAttachments[pod] = append(networks[pa.network].podAttachments[pod], pa)
		}
	}

	// Collect all secondary network IPs when they are available.
	for _, testPod := range data.pods {
		_, err := e2eTestData.PodWaitFor(defaultTimeout, testPod.podName, namespace, func(pod *corev1.Pod) (bool, error) {
			if pod.Status.Phase != corev1.PodRunning {
				return false, nil
			}
			var podNetworkAttachments []*attachment
			podIPs, err := data.listPodIPs(testPod)
			if err != nil {
				return false, err
			}
			for iface, net := range testPod.interfaceNetworks {
				if podIPs[iface] == nil {
					return false, nil
				}
				podNetworkAttachments = append(podNetworkAttachments, &attachment{
					network: net,
					iface:   iface,
					ip:      podIPs[iface],
				})
			}
			// we found all the expected secondary network interfaces / attachments
			addPodNetworkAttachments(testPod, podNetworkAttachments)
			return true, nil
		})
		if err != nil {
			return fmt.Errorf("error when waiting for secondary IPs for Pod %+v: %v", testPod, err)
		}
	}

	// Run ping-mesh test for each secondary network.
	for _, network := range networks {
		for sourcePod := range network.podAttachments {
			for targetPod, targetPodAttachments := range network.podAttachments {
				if sourcePod == targetPod {
					continue
				}
				for _, targetAttachment := range targetPodAttachments {
					var IPToPing antreae2e.PodIPs
					if targetAttachment.ip.To4() != nil {
						IPToPing = antreae2e.PodIPs{IPv4: &targetAttachment.ip}
					} else {
						IPToPing = antreae2e.PodIPs{IPv6: &targetAttachment.ip}
					}
					if err := e2eTestData.RunPingCommandFromTestPod(antreae2e.PodInfo{Name: sourcePod.podName, OS: osType, NodeName: sourcePod.nodeName, Namespace: namespace},
						namespace, &IPToPing, containerName, pingCount, pingSize, false); err != nil {
						return fmt.Errorf("ping '%s' -> '%s'(Interface: %s, IP Address: %s) failed: %v", sourcePod.podName, targetPod.podName, targetAttachment.iface, targetAttachment.ip, err)
					}
					logs.Infof("ping '%s' -> '%s'( Interface: %s, IP Address: %s): OK", sourcePod.podName, targetPod.podName, targetAttachment.iface, targetAttachment.ip)
				}
			}
		}
	}

	return nil
}

func testSecondaryNetwork(t *testing.T, networkType string, pods []*testPodInfo) {
	e2eTestData, err := antreae2e.SetupTest(t)
	if err != nil {
		t.Fatalf("Error when setting up test: %v", err)
	}
	defer antreae2e.TeardownTest(t, e2eTestData)

	testData := &testData{e2eTestData: e2eTestData, networkType: networkType, pods: pods}

	err = testData.createPods(t, e2eTestData.GetTestNamespace())
	if err != nil {
		t.Fatalf("Error when create test Pods: %v", err)
	}
	err = testData.pingBetweenInterfaces(t)
	if err != nil {
		t.Fatalf("Error when pinging between interfaces: %v", err)
	}
}

func TestVLANNetwork(t *testing.T) {
	if antreae2e.NodeCount() < 2 {
		t.Fatalf("The test requires at least 2 nodes, but the cluster has only %d", antreae2e.NodeCount())
	}
	node1 := antreae2e.NodeName(0)
	node2 := antreae2e.NodeName(1)
	pods := []*testPodInfo{
		{
			podName:           "vlan-pod1",
			nodeName:          node1,
			interfaceNetworks: map[string]string{"eth1": "vlan-net1", "eth2": "vlan-net2"},
		},
		{
			podName:           "vlan-pod2",
			nodeName:          node1,
			interfaceNetworks: map[string]string{"eth1": "vlan-net1", "eth2": "vlan-net3"},
		},
		{
			podName:           "vlan-pod3",
			nodeName:          node2,
			interfaceNetworks: map[string]string{"eth1": "vlan-net2"},
		},
	}
	testSecondaryNetwork(t, networkTypeVLAN, pods)
}

func executeShellScript(script string) (string, error) {
	cmd := exec.Command("bash", "-c", script)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute script: %s, output: %s, %s, error: %v", script, out.String(), stderr.String(), err)
	}

	return out.String(), nil
}

func (data *testData) assignIP() error {
	e2eTestData := data.e2eTestData
	namespace := e2eTestData.GetTestNamespace()

	for _, testPod := range data.pods {
		_, err := e2eTestData.PodWaitFor(defaultTimeout, testPod.podName, namespace, func(pod *corev1.Pod) (bool, error) {
			if pod.Status.Phase != corev1.PodRunning {
				return false, nil
			}
			podIPs, err := data.listPodIPs(testPod)
			if err != nil {
				return false, err
			}
			ip, exists := podIPs["eth1"]
			if !exists || ip == nil {
				logs.Infof("IP not available for interface 'eth1' in Pod %s, retrying...", testPod.podName)
				return false, nil
			}

			cmd := fmt.Sprintf("aws ec2 assign-private-ip-addresses --network-interface-id %s --private-ip-addresses %s", testPod.secondaryENI, ip)
			output, err := executeShellScript(cmd)
			if err != nil {
				return false, err
			}
			logs.Infof("assign private ip addresses: %s", output)
			return true, nil
		})
		if err != nil {
			return fmt.Errorf("error when waiting for secondary IPs for Pod %+v: %v", testPod, err)
		}
	}
	return nil
}

// createNetworkAttachmentDefinition creates a NetworkAttachmentDefinition in the specified namespace.
func createNetworkAttachmentDefinition(client networkclient.Interface, namespace, name, config string) error {
	nad := &netattachv1.NetworkAttachmentDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Annotations: map[string]string{"k8s.v1.cni.cncf.io/resourceName": "intel.com/intel_sriov_netdevice"},
		},
		Spec: netattachv1.NetworkAttachmentDefinitionSpec{
			Config: config,
		},
	}

	_, err := client.K8sCniCncfIoV1().NetworkAttachmentDefinitions(namespace).Create(context.TODO(), nad, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create NetworkAttachmentDefinition: %v", err)
	}

	logs.Infof("NetworkAttachmentDefinition %s created successfully in namespace %s\n", name, namespace)
	return nil
}

func TestSRIOVNetwork(t *testing.T) {
	node0Name := antreae2e.NodeName(0)
	node1Name := antreae2e.NodeName(1)
	pods := []*testPodInfo{
		{
			podName:           "sriov-pod1",
			nodeName:          node0Name,
			interfaceNetworks: map[string]string{"eth1": "sriov-net1"},
			isSRIOV:           true,
			secondaryENI:      antreae2e.NodeENI(0),
		},
		{
			podName:           "sriov-pod2",
			nodeName:          node1Name,
			interfaceNetworks: map[string]string{"eth1": "sriov-net1"},
			isSRIOV:           true,
			secondaryENI:      antreae2e.NodeENI(1),
		},
	}

	e2eTestData, err := antreae2e.SetupTest(t)
	if err != nil {
		t.Fatalf("Error when setting up test: %v", err)
	}
	defer antreae2e.TeardownTest(t, e2eTestData)

	testData := &testData{e2eTestData: e2eTestData, networkType: networkTypeSriov, pods: pods}

	configJSON := `{
		"cniVersion": "0.3.0",
		"type": "antrea",
		"networkType": "sriov",
		"ipam": {
			"type": "antrea",
			"ippools": ["pool1"]
		}
	}`
	err = createNetworkAttachmentDefinition(e2eTestData.GetNetworkClient(), e2eTestData.GetTestNamespace(), "sriov-net1", configJSON)
	require.NoError(t, err)

	err = testData.createPods(t, e2eTestData.GetTestNamespace())
	if err != nil {
		t.Fatalf("Error when create test Pods: %v", err)
	}
	err = testData.assignIP()
	if err != nil {
		t.Fatalf("Error when assign IP to ec2 instance: %v", err)
	}
	err = testData.pingBetweenInterfaces(t)
	if err != nil {
		t.Fatalf("Error when pinging between interfaces: %v", err)
	}
}
