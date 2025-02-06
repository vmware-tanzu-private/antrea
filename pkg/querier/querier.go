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

package querier

import (
	"context"

	v1 "k8s.io/api/core/v1"
	apitypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/pkg/agent/apis"
	"antrea.io/antrea/pkg/agent/bgp"
	bgpcontroller "antrea.io/antrea/pkg/agent/controller/bgp"
	"antrea.io/antrea/pkg/agent/interfacestore"
	"antrea.io/antrea/pkg/agent/multicast"
	"antrea.io/antrea/pkg/agent/types"
	cpv1beta "antrea.io/antrea/pkg/apis/controlplane/v1beta2"
	"antrea.io/antrea/pkg/util/env"
	"antrea.io/antrea/pkg/version"
)

type NetworkPolicyInfoQuerier interface {
	GetNetworkPolicyNum() int
	GetAddressGroupNum() int
	GetAppliedToGroupNum() int
}

type AgentNetworkPolicyInfoQuerier interface {
	NetworkPolicyInfoQuerier
	GetControllerConnectionStatus() bool
	GetNetworkPolicies(npFilter *NetworkPolicyQueryFilter) []cpv1beta.NetworkPolicy
	GetAddressGroups() []cpv1beta.AddressGroup
	GetAppliedToGroups() []cpv1beta.AppliedToGroup
	GetAppliedNetworkPolicies(pod, namespace string, npFilter *NetworkPolicyQueryFilter) []cpv1beta.NetworkPolicy
	GetNetworkPolicyByRuleFlowID(ruleFlowID uint32) *cpv1beta.NetworkPolicyReference
	GetRuleByFlowID(ruleFlowID uint32) *types.PolicyRule
	GetFQDNCache(fqdnFilter FQDNCacheFilter) []types.DnsCacheEntry
}

type AgentMulticastInfoQuerier interface {
	// CollectIGMPReportNPStats gets statistics generated by NetworkPolicies that block or allow
	// IGMP reports message. The statistics returned are incremental and will be reset after each call.
	CollectIGMPReportNPStats() (annpStats, acnpStats map[apitypes.UID]map[string]*types.RuleMetric)
	// GetGroupPods gets a map that saves the local Pod members of multicast groups on the Node.
	GetGroupPods() map[string][]cpv1beta.PodReference
	// GetAllPodsStats gets multicast traffic statistics of all local Pods.
	GetAllPodsStats() map[*interfacestore.InterfaceConfig]*multicast.PodTrafficStats
	// GetPodStats gets multicast traffic statistics of a local Pod, specified by podName and podNamespace.
	GetPodStats(podName string, podNamespace string) *multicast.PodTrafficStats
}

type ControllerNetworkPolicyInfoQuerier interface {
	NetworkPolicyInfoQuerier
	GetConnectedAgentNum() int
}

type EgressQuerier interface {
	GetEgressIPByMark(mark uint32) (string, error)
	GetEgress(podNamespace, podName string) (string, string, string, error)
}

// GetSelfPod gets current pod.
func GetSelfPod() v1.ObjectReference {
	podName := env.GetPodName()
	podNamespace := env.GetPodNamespace()
	if podName == "" || podNamespace == "" {
		return v1.ObjectReference{}
	}
	return v1.ObjectReference{Kind: "Pod", Name: podName, Namespace: podNamespace}
}

// GetSelfNode gets current node.
func GetSelfNode(isAgent bool, node string) v1.ObjectReference {
	if isAgent {
		if node == "" {
			return v1.ObjectReference{}
		}
		return v1.ObjectReference{Kind: "Node", Name: node}
	}
	nodeName, _ := env.GetNodeName()
	if nodeName == "" {
		return v1.ObjectReference{}
	}
	return v1.ObjectReference{Kind: "Node", Name: nodeName}
}

// GetVersion gets current version.
func GetVersion() string {
	return version.GetFullVersion()
}

// FQDNCacheFilter is used to filter the result while retrieving FQDN cache
type FQDNCacheFilter struct {
	// The Name or wildcard matching expression of the domain that is being filtered
	DomainName string
}

// NetworkPolicyQueryFilter is used to filter the result while retrieve network policy
// An empty attribute, which won't be used as a condition, means match all.
// e.g SourceType = "" means all type network policy will be retrieved
// Can have more attributes in future if more args are required
type NetworkPolicyQueryFilter struct {
	// The Name of the controlplane network policy. If this field is set then
	// none of the other fields can be.
	Name string
	// The Name of the original network policy.
	SourceName string
	// The namespace of the original Namespace that the internal NetworkPolicy is created for.
	Namespace string
	// The type of the original NetworkPolicy that the internal NetworkPolicy is created for.(K8sNP, ACNP, ANNP, ANP and BANP)
	SourceType cpv1beta.NetworkPolicyType
}

// From user shorthand input to cpv1beta1.NetworkPolicyType
var NetworkPolicyTypeMap = map[string]cpv1beta.NetworkPolicyType{
	"K8SNP": cpv1beta.K8sNetworkPolicy,
	"ACNP":  cpv1beta.AntreaClusterNetworkPolicy,
	"ANNP":  cpv1beta.AntreaNetworkPolicy,
	"ANP":   cpv1beta.AdminNetworkPolicy,
	"BANP":  cpv1beta.BaselineAdminNetworkPolicy,
}

func GetNetworkPolicyTypeShorthands() []string {
	validTypes := make([]string, 0, len(NetworkPolicyTypeMap))
	for k := range NetworkPolicyTypeMap {
		validTypes = append(validTypes, k)
	}
	return validTypes
}

var NamespaceScopedPolicyTypes = sets.New[string]("ANNP", "K8SNP")

// ServiceExternalIPStatusQuerier queries the Service external IP status for debugging purposes.
// Ideally, every Node should have consistent results eventually. This should only be used when
// ServiceExternalIP feature is enabled.
type ServiceExternalIPStatusQuerier interface {
	GetServiceExternalIPStatus() []apis.ServiceExternalIPInfo
}

type AgentBGPPolicyInfoQuerier interface {
	// GetBGPPolicyInfo returns Name, RouterID, LocalASN and ListenPort of effective BGP Policy applied on the Node.
	GetBGPPolicyInfo() (string, string, int32, int32)
	// GetBGPPeerStatus returns current status of BGP Peers of effective BGP Policy applied on the Node.
	GetBGPPeerStatus(ctx context.Context) ([]bgp.PeerStatus, error)
	// GetBGPRoutes returns the advertised BGP routes.
	GetBGPRoutes(ctx context.Context) (map[bgp.Route]bgpcontroller.RouteMetadata, error)
}
