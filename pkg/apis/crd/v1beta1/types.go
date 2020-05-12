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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterNetworkPolicy struct {
	metav1.TypeMeta `json:",inline"`
	// Standard metadata of the object.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of ClusterNetworkPolicy.
	Spec ClusterNetworkPolicySpec `json:"spec"`
}

// ClusterNetworkPolicySpec defines the desired state for ClusterNetworkPolicy.
type ClusterNetworkPolicySpec struct {
	// Priority specfies the order of the ClusterNetworkPolicy relative to
	// other ClusterNetworkPolicies.
	Priority int32 `json:"priority"`
	// Select workloads on which the rules will be applied to.
	AppliedTo []NetworkPolicyPeer `json:"appliedTo"`
	// Set of ingress rules evaluated based on the order in which they are set.
	// +optional
	Ingress []IngressRule `json:"ingress"`
	// Set of egress rules evaluated based on the order in which they are set.
	// +optional
	Egress []EgressRule `json:"egress"`
}

// IngressRule describes the traffic allowed to the workloads selected by
// Spec.AppliedTo. Based on the action specified in the rule, traffic is either
// allowed or denied which exactly match the specified ports and protocol.
type IngressRule struct {
	// Action specifies the action to be applied on the rule. Defaults to
	// ALLOW action.
	// +optional
	Action *RuleAction `json:"action"`
	// Set of port and protocol allowed/denied by the rule. If this field is unset
	// or empty, this rule matches all ports.
	// +optional
	Ports []NetworkPolicyPort `json:"ports"`
	// Rule is matched if traffic originates from workloads selected by
	// this field. If this field is empty or missing, this rule matches all
	// sources.
	// +optional
	From []NetworkPolicyPeer `json:"from"`
}

// EgressRule describes the traffic allowed from the workloads selected by
// Spec.AppliedTo. Based on the action specified in the rule, traffic is either
// allowed or denied which exactly match the specified ports and protocol.
type EgressRule struct {
	// Action specifies the action to be applied on the rule. Defaults to
	// ALLOW action.
	// +optional
	Action *RuleAction `json:"action"`
	// Set of port and protocol allowed/denied by the rule. If this field is unset
	// or empty, this rule matches all ports.
	// +optional
	Ports []NetworkPolicyPort `json:"ports"`
	// Rule is matched if traffic is intended for workloads selected by
	// this field. If this field is empty or missing, this rule matches all
	// destinations.
	// +optional
	To []NetworkPolicyPeer `json:"to"`
}

// NetworkPolicyPeer describes the grouping selector of workloads.
type NetworkPolicyPeer struct {
	// IPBlock describes the IPAddresses/IPBlocks that is matched in to/from.
	// IPBlock cannot be set as part of the AppliedTo field
	// Cannot be set with any other selector.
	// +optional
	IPBlock *IPBlock `json:"ipBlock,omitempty"`
	// Select Pods from all Namespaces as workloads in AppliedTo/To/From
	// fields. If set with NamespaceSelector, Pods are matched from Namespaces
	// matched by the NamespaceSelector.
	// Cannot be set with any other selector except NamespaceSelector.
	// +optional
	PodSelector *metav1.LabelSelector `json:"podSelector,omitempty"`
	// Select all Pods from Namespaces matched by this selector, as
	// workloads in AppliedTo/To/From fields. If set with PodSelector,
	// Pods are matched from Namespaces matched by the NamespaceSelector.
	// Cannot be set with any other selector except PodSelector.
	// +optional
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
}

// IPBlock describes a particular CIDR (Ex. "192.168.1.1/24") that is allowed
// or denied to/from the workloads matched by a Spec.AppliedTo.
type IPBlock struct {
	// CIDR is a string representing the IP Block
	// Valid examples are "192.168.1.1/24".
	CIDR string `json:"cidr"`
}

// NetworkPolicyPort describes the port and protocol to match in a rule.
type NetworkPolicyPort struct {
	// The protocol (TCP, UDP, or SCTP) which traffic must match.
	// If not specified, this field defaults to TCP.
	// +optional
	Protocol *v1.Protocol `json:"protocol"`
	// The port on the given protocol. This can either be a numerical
	// or named port on a Pod. If this field is not provided, this
	// matches all port names and numbers.
	// TODO: extend it to include Port Range.
	// +optional
	Port *intstr.IntOrString `json:"port"`
}

// RuleAction describes the action to be applied on traffic matching a rule.
type RuleAction string

const (
	// RuleActionAllow describes that rule matching traffic must be allowed.
	RuleActionAllow RuleAction = "ALLOW"
	// RuleActionDrop describes that rule matching traffic must be dropped.
	RuleActionDrop RuleAction = "DROP"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ClusterNetworkPolicy `json:"items"`
}
