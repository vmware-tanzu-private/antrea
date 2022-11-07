//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedTo) DeepCopyInto(out *AppliedTo) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalEntitySelector != nil {
		in, out := &in.ExternalEntitySelector, &out.ExternalEntitySelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(NamespacedName)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(NamespacedName)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedTo.
func (in *AppliedTo) DeepCopy() *AppliedTo {
	if in == nil {
		return nil
	}
	out := new(AppliedTo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleExternalNodes) DeepCopyInto(out *BundleExternalNodes) {
	*out = *in
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleExternalNodes.
func (in *BundleExternalNodes) DeepCopy() *BundleExternalNodes {
	if in == nil {
		return nil
	}
	out := new(BundleExternalNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleFileServer) DeepCopyInto(out *BundleFileServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleFileServer.
func (in *BundleFileServer) DeepCopy() *BundleFileServer {
	if in == nil {
		return nil
	}
	out := new(BundleFileServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleNodes) DeepCopyInto(out *BundleNodes) {
	*out = *in
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleNodes.
func (in *BundleNodes) DeepCopy() *BundleNodes {
	if in == nil {
		return nil
	}
	out := new(BundleNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleServerAuthConfiguration) DeepCopyInto(out *BundleServerAuthConfiguration) {
	*out = *in
	if in.AuthSecret != nil {
		in, out := &in.AuthSecret, &out.AuthSecret
		*out = new(corev1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleServerAuthConfiguration.
func (in *BundleServerAuthConfiguration) DeepCopy() *BundleServerAuthConfiguration {
	if in == nil {
		return nil
	}
	out := new(BundleServerAuthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkPolicy) DeepCopyInto(out *ClusterNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkPolicy.
func (in *ClusterNetworkPolicy) DeepCopy() *ClusterNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkPolicyList) DeepCopyInto(out *ClusterNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkPolicyList.
func (in *ClusterNetworkPolicyList) DeepCopy() *ClusterNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkPolicySpec) DeepCopyInto(out *ClusterNetworkPolicySpec) {
	*out = *in
	if in.AppliedTo != nil {
		in, out := &in.AppliedTo, &out.AppliedTo
		*out = make([]AppliedTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkPolicySpec.
func (in *ClusterNetworkPolicySpec) DeepCopy() *ClusterNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNode) DeepCopyInto(out *ExternalNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNode.
func (in *ExternalNode) DeepCopy() *ExternalNode {
	if in == nil {
		return nil
	}
	out := new(ExternalNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNodeList) DeepCopyInto(out *ExternalNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNodeList.
func (in *ExternalNodeList) DeepCopy() *ExternalNodeList {
	if in == nil {
		return nil
	}
	out := new(ExternalNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNodeSpec) DeepCopyInto(out *ExternalNodeSpec) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]NetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNodeSpec.
func (in *ExternalNodeSpec) DeepCopy() *ExternalNodeSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProtocol) DeepCopyInto(out *HTTPProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProtocol.
func (in *HTTPProtocol) DeepCopy() *HTTPProtocol {
	if in == nil {
		return nil
	}
	out := new(HTTPProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPEchoRequestHeader) DeepCopyInto(out *ICMPEchoRequestHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPEchoRequestHeader.
func (in *ICMPEchoRequestHeader) DeepCopy() *ICMPEchoRequestHeader {
	if in == nil {
		return nil
	}
	out := new(ICMPEchoRequestHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPProtocol) DeepCopyInto(out *ICMPProtocol) {
	*out = *in
	if in.ICMPType != nil {
		in, out := &in.ICMPType, &out.ICMPType
		*out = new(int32)
		**out = **in
	}
	if in.ICMPCode != nil {
		in, out := &in.ICMPCode, &out.ICMPCode
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPProtocol.
func (in *ICMPProtocol) DeepCopy() *ICMPProtocol {
	if in == nil {
		return nil
	}
	out := new(ICMPProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IGMPProtocol) DeepCopyInto(out *IGMPProtocol) {
	*out = *in
	if in.IGMPType != nil {
		in, out := &in.IGMPType, &out.IGMPType
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IGMPProtocol.
func (in *IGMPProtocol) DeepCopy() *IGMPProtocol {
	if in == nil {
		return nil
	}
	out := new(IGMPProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPHeader) DeepCopyInto(out *IPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPHeader.
func (in *IPHeader) DeepCopy() *IPHeader {
	if in == nil {
		return nil
	}
	out := new(IPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6Header) DeepCopyInto(out *IPv6Header) {
	*out = *in
	if in.NextHeader != nil {
		in, out := &in.NextHeader, &out.NextHeader
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6Header.
func (in *IPv6Header) DeepCopy() *IPv6Header {
	if in == nil {
		return nil
	}
	out := new(IPv6Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7Protocol) DeepCopyInto(out *L7Protocol) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPProtocol)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Protocol.
func (in *L7Protocol) DeepCopy() *L7Protocol {
	if in == nil {
		return nil
	}
	out := new(L7Protocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterface) DeepCopyInto(out *NetworkInterface) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterface.
func (in *NetworkInterface) DeepCopy() *NetworkInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicy) DeepCopyInto(out *NetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicy.
func (in *NetworkPolicy) DeepCopy() *NetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyCondition) DeepCopyInto(out *NetworkPolicyCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyCondition.
func (in *NetworkPolicyCondition) DeepCopy() *NetworkPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyList) DeepCopyInto(out *NetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyList.
func (in *NetworkPolicyList) DeepCopy() *NetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyPeer) DeepCopyInto(out *NetworkPolicyPeer) {
	*out = *in
	if in.IPBlock != nil {
		in, out := &in.IPBlock, &out.IPBlock
		*out = new(IPBlock)
		**out = **in
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = new(PeerNamespaces)
		**out = **in
	}
	if in.ExternalEntitySelector != nil {
		in, out := &in.ExternalEntitySelector, &out.ExternalEntitySelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(NamespacedName)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyPeer.
func (in *NetworkPolicyPeer) DeepCopy() *NetworkPolicyPeer {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyPort) DeepCopyInto(out *NetworkPolicyPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(corev1.Protocol)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.EndPort != nil {
		in, out := &in.EndPort, &out.EndPort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyPort.
func (in *NetworkPolicyPort) DeepCopy() *NetworkPolicyPort {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyProtocol) DeepCopyInto(out *NetworkPolicyProtocol) {
	*out = *in
	if in.ICMP != nil {
		in, out := &in.ICMP, &out.ICMP
		*out = new(ICMPProtocol)
		(*in).DeepCopyInto(*out)
	}
	if in.IGMP != nil {
		in, out := &in.IGMP, &out.IGMP
		*out = new(IGMPProtocol)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyProtocol.
func (in *NetworkPolicyProtocol) DeepCopy() *NetworkPolicyProtocol {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicySpec) DeepCopyInto(out *NetworkPolicySpec) {
	*out = *in
	if in.AppliedTo != nil {
		in, out := &in.AppliedTo, &out.AppliedTo
		*out = make([]AppliedTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicySpec.
func (in *NetworkPolicySpec) DeepCopy() *NetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyStatus) DeepCopyInto(out *NetworkPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]NetworkPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyStatus.
func (in *NetworkPolicyStatus) DeepCopy() *NetworkPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResult) DeepCopyInto(out *NodeResult) {
	*out = *in
	if in.Observations != nil {
		in, out := &in.Observations, &out.Observations
		*out = make([]Observation, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResult.
func (in *NodeResult) DeepCopy() *NodeResult {
	if in == nil {
		return nil
	}
	out := new(NodeResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observation) DeepCopyInto(out *Observation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observation.
func (in *Observation) DeepCopy() *Observation {
	if in == nil {
		return nil
	}
	out := new(Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Packet) DeepCopyInto(out *Packet) {
	*out = *in
	out.IPHeader = in.IPHeader
	if in.IPv6Header != nil {
		in, out := &in.IPv6Header, &out.IPv6Header
		*out = new(IPv6Header)
		(*in).DeepCopyInto(*out)
	}
	in.TransportHeader.DeepCopyInto(&out.TransportHeader)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Packet.
func (in *Packet) DeepCopy() *Packet {
	if in == nil {
		return nil
	}
	out := new(Packet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerNamespaces) DeepCopyInto(out *PeerNamespaces) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerNamespaces.
func (in *PeerNamespaces) DeepCopy() *PeerNamespaces {
	if in == nil {
		return nil
	}
	out := new(PeerNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(RuleAction)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]NetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]NetworkPolicyProtocol, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.L7Protocols != nil {
		in, out := &in.L7Protocols, &out.L7Protocols
		*out = make([]L7Protocol, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]NetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]NetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToServices != nil {
		in, out := &in.ToServices, &out.ToServices
		*out = make([]NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.AppliedTo != nil {
		in, out := &in.AppliedTo, &out.AppliedTo
		*out = make([]AppliedTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportBundleCollection) DeepCopyInto(out *SupportBundleCollection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportBundleCollection.
func (in *SupportBundleCollection) DeepCopy() *SupportBundleCollection {
	if in == nil {
		return nil
	}
	out := new(SupportBundleCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SupportBundleCollection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportBundleCollectionCondition) DeepCopyInto(out *SupportBundleCollectionCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportBundleCollectionCondition.
func (in *SupportBundleCollectionCondition) DeepCopy() *SupportBundleCollectionCondition {
	if in == nil {
		return nil
	}
	out := new(SupportBundleCollectionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportBundleCollectionList) DeepCopyInto(out *SupportBundleCollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SupportBundleCollection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportBundleCollectionList.
func (in *SupportBundleCollectionList) DeepCopy() *SupportBundleCollectionList {
	if in == nil {
		return nil
	}
	out := new(SupportBundleCollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SupportBundleCollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportBundleCollectionSpec) DeepCopyInto(out *SupportBundleCollectionSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = new(BundleNodes)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalNodes != nil {
		in, out := &in.ExternalNodes, &out.ExternalNodes
		*out = new(BundleExternalNodes)
		(*in).DeepCopyInto(*out)
	}
	out.FileServer = in.FileServer
	in.Authentication.DeepCopyInto(&out.Authentication)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportBundleCollectionSpec.
func (in *SupportBundleCollectionSpec) DeepCopy() *SupportBundleCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(SupportBundleCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportBundleCollectionStatus) DeepCopyInto(out *SupportBundleCollectionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]SupportBundleCollectionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportBundleCollectionStatus.
func (in *SupportBundleCollectionStatus) DeepCopy() *SupportBundleCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(SupportBundleCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPHeader) DeepCopyInto(out *TCPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPHeader.
func (in *TCPHeader) DeepCopy() *TCPHeader {
	if in == nil {
		return nil
	}
	out := new(TCPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tier) DeepCopyInto(out *Tier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tier.
func (in *Tier) DeepCopy() *Tier {
	if in == nil {
		return nil
	}
	out := new(Tier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TierList) DeepCopyInto(out *TierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TierList.
func (in *TierList) DeepCopy() *TierList {
	if in == nil {
		return nil
	}
	out := new(TierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TierSpec) DeepCopyInto(out *TierSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TierSpec.
func (in *TierSpec) DeepCopy() *TierSpec {
	if in == nil {
		return nil
	}
	out := new(TierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Traceflow) DeepCopyInto(out *Traceflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Traceflow.
func (in *Traceflow) DeepCopy() *Traceflow {
	if in == nil {
		return nil
	}
	out := new(Traceflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Traceflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceflowList) DeepCopyInto(out *TraceflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Traceflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceflowList.
func (in *TraceflowList) DeepCopy() *TraceflowList {
	if in == nil {
		return nil
	}
	out := new(TraceflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TraceflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceflowSpec) DeepCopyInto(out *TraceflowSpec) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
	in.Packet.DeepCopyInto(&out.Packet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceflowSpec.
func (in *TraceflowSpec) DeepCopy() *TraceflowSpec {
	if in == nil {
		return nil
	}
	out := new(TraceflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceflowStatus) DeepCopyInto(out *TraceflowStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]NodeResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CapturedPacket != nil {
		in, out := &in.CapturedPacket, &out.CapturedPacket
		*out = new(Packet)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceflowStatus.
func (in *TraceflowStatus) DeepCopy() *TraceflowStatus {
	if in == nil {
		return nil
	}
	out := new(TraceflowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransportHeader) DeepCopyInto(out *TransportHeader) {
	*out = *in
	if in.ICMP != nil {
		in, out := &in.ICMP, &out.ICMP
		*out = new(ICMPEchoRequestHeader)
		**out = **in
	}
	if in.UDP != nil {
		in, out := &in.UDP, &out.UDP
		*out = new(UDPHeader)
		**out = **in
	}
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(TCPHeader)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransportHeader.
func (in *TransportHeader) DeepCopy() *TransportHeader {
	if in == nil {
		return nil
	}
	out := new(TransportHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPHeader) DeepCopyInto(out *UDPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPHeader.
func (in *UDPHeader) DeepCopy() *UDPHeader {
	if in == nil {
		return nil
	}
	out := new(UDPHeader)
	in.DeepCopyInto(out)
	return out
}
