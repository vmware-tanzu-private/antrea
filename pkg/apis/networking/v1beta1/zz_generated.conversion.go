// +build !ignore_autogenerated

// Copyright 2020 Antrea Authors
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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	networking "github.com/vmware-tanzu/antrea/pkg/apis/networking"
	v1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/security/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AddressGroup)(nil), (*networking.AddressGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AddressGroup_To_networking_AddressGroup(a.(*AddressGroup), b.(*networking.AddressGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.AddressGroup)(nil), (*AddressGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_AddressGroup_To_v1beta1_AddressGroup(a.(*networking.AddressGroup), b.(*AddressGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AddressGroupList)(nil), (*networking.AddressGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AddressGroupList_To_networking_AddressGroupList(a.(*AddressGroupList), b.(*networking.AddressGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.AddressGroupList)(nil), (*AddressGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_AddressGroupList_To_v1beta1_AddressGroupList(a.(*networking.AddressGroupList), b.(*AddressGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AddressGroupPatch)(nil), (*networking.AddressGroupPatch)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AddressGroupPatch_To_networking_AddressGroupPatch(a.(*AddressGroupPatch), b.(*networking.AddressGroupPatch), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.AddressGroupPatch)(nil), (*AddressGroupPatch)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_AddressGroupPatch_To_v1beta1_AddressGroupPatch(a.(*networking.AddressGroupPatch), b.(*AddressGroupPatch), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppliedToGroup)(nil), (*networking.AppliedToGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AppliedToGroup_To_networking_AppliedToGroup(a.(*AppliedToGroup), b.(*networking.AppliedToGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.AppliedToGroup)(nil), (*AppliedToGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_AppliedToGroup_To_v1beta1_AppliedToGroup(a.(*networking.AppliedToGroup), b.(*AppliedToGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppliedToGroupList)(nil), (*networking.AppliedToGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AppliedToGroupList_To_networking_AppliedToGroupList(a.(*AppliedToGroupList), b.(*networking.AppliedToGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.AppliedToGroupList)(nil), (*AppliedToGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_AppliedToGroupList_To_v1beta1_AppliedToGroupList(a.(*networking.AppliedToGroupList), b.(*AppliedToGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AppliedToGroupPatch)(nil), (*networking.AppliedToGroupPatch)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AppliedToGroupPatch_To_networking_AppliedToGroupPatch(a.(*AppliedToGroupPatch), b.(*networking.AppliedToGroupPatch), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.AppliedToGroupPatch)(nil), (*AppliedToGroupPatch)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_AppliedToGroupPatch_To_v1beta1_AppliedToGroupPatch(a.(*networking.AppliedToGroupPatch), b.(*AppliedToGroupPatch), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*GroupMemberPod)(nil), (*networking.GroupMemberPod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_GroupMemberPod_To_networking_GroupMemberPod(a.(*GroupMemberPod), b.(*networking.GroupMemberPod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.GroupMemberPod)(nil), (*GroupMemberPod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_GroupMemberPod_To_v1beta1_GroupMemberPod(a.(*networking.GroupMemberPod), b.(*GroupMemberPod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPBlock)(nil), (*networking.IPBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IPBlock_To_networking_IPBlock(a.(*IPBlock), b.(*networking.IPBlock), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IPBlock)(nil), (*IPBlock)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IPBlock_To_v1beta1_IPBlock(a.(*networking.IPBlock), b.(*IPBlock), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPNet)(nil), (*networking.IPNet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IPNet_To_networking_IPNet(a.(*IPNet), b.(*networking.IPNet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IPNet)(nil), (*IPNet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IPNet_To_v1beta1_IPNet(a.(*networking.IPNet), b.(*IPNet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NamedPort)(nil), (*networking.NamedPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NamedPort_To_networking_NamedPort(a.(*NamedPort), b.(*networking.NamedPort), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.NamedPort)(nil), (*NamedPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_NamedPort_To_v1beta1_NamedPort(a.(*networking.NamedPort), b.(*NamedPort), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkPolicy)(nil), (*networking.NetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NetworkPolicy_To_networking_NetworkPolicy(a.(*NetworkPolicy), b.(*networking.NetworkPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicy)(nil), (*NetworkPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_NetworkPolicy_To_v1beta1_NetworkPolicy(a.(*networking.NetworkPolicy), b.(*NetworkPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkPolicyList)(nil), (*networking.NetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NetworkPolicyList_To_networking_NetworkPolicyList(a.(*NetworkPolicyList), b.(*networking.NetworkPolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyList)(nil), (*NetworkPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_NetworkPolicyList_To_v1beta1_NetworkPolicyList(a.(*networking.NetworkPolicyList), b.(*NetworkPolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkPolicyPeer)(nil), (*networking.NetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(a.(*NetworkPolicyPeer), b.(*networking.NetworkPolicyPeer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyPeer)(nil), (*NetworkPolicyPeer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer(a.(*networking.NetworkPolicyPeer), b.(*NetworkPolicyPeer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkPolicyRule)(nil), (*networking.NetworkPolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NetworkPolicyRule_To_networking_NetworkPolicyRule(a.(*NetworkPolicyRule), b.(*networking.NetworkPolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.NetworkPolicyRule)(nil), (*NetworkPolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_NetworkPolicyRule_To_v1beta1_NetworkPolicyRule(a.(*networking.NetworkPolicyRule), b.(*NetworkPolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodReference)(nil), (*networking.PodReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodReference_To_networking_PodReference(a.(*PodReference), b.(*networking.PodReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.PodReference)(nil), (*PodReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_PodReference_To_v1beta1_PodReference(a.(*networking.PodReference), b.(*PodReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Service)(nil), (*networking.Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Service_To_networking_Service(a.(*Service), b.(*networking.Service), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.Service)(nil), (*Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_Service_To_v1beta1_Service(a.(*networking.Service), b.(*Service), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_AddressGroup_To_networking_AddressGroup(in *AddressGroup, out *networking.AddressGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Pods = *(*[]networking.GroupMemberPod)(unsafe.Pointer(&in.Pods))
	return nil
}

// Convert_v1beta1_AddressGroup_To_networking_AddressGroup is an autogenerated conversion function.
func Convert_v1beta1_AddressGroup_To_networking_AddressGroup(in *AddressGroup, out *networking.AddressGroup, s conversion.Scope) error {
	return autoConvert_v1beta1_AddressGroup_To_networking_AddressGroup(in, out, s)
}

func autoConvert_networking_AddressGroup_To_v1beta1_AddressGroup(in *networking.AddressGroup, out *AddressGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Pods = *(*[]GroupMemberPod)(unsafe.Pointer(&in.Pods))
	return nil
}

// Convert_networking_AddressGroup_To_v1beta1_AddressGroup is an autogenerated conversion function.
func Convert_networking_AddressGroup_To_v1beta1_AddressGroup(in *networking.AddressGroup, out *AddressGroup, s conversion.Scope) error {
	return autoConvert_networking_AddressGroup_To_v1beta1_AddressGroup(in, out, s)
}

func autoConvert_v1beta1_AddressGroupList_To_networking_AddressGroupList(in *AddressGroupList, out *networking.AddressGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.AddressGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_AddressGroupList_To_networking_AddressGroupList is an autogenerated conversion function.
func Convert_v1beta1_AddressGroupList_To_networking_AddressGroupList(in *AddressGroupList, out *networking.AddressGroupList, s conversion.Scope) error {
	return autoConvert_v1beta1_AddressGroupList_To_networking_AddressGroupList(in, out, s)
}

func autoConvert_networking_AddressGroupList_To_v1beta1_AddressGroupList(in *networking.AddressGroupList, out *AddressGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]AddressGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_networking_AddressGroupList_To_v1beta1_AddressGroupList is an autogenerated conversion function.
func Convert_networking_AddressGroupList_To_v1beta1_AddressGroupList(in *networking.AddressGroupList, out *AddressGroupList, s conversion.Scope) error {
	return autoConvert_networking_AddressGroupList_To_v1beta1_AddressGroupList(in, out, s)
}

func autoConvert_v1beta1_AddressGroupPatch_To_networking_AddressGroupPatch(in *AddressGroupPatch, out *networking.AddressGroupPatch, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.AddedPods = *(*[]networking.GroupMemberPod)(unsafe.Pointer(&in.AddedPods))
	out.RemovedPods = *(*[]networking.GroupMemberPod)(unsafe.Pointer(&in.RemovedPods))
	return nil
}

// Convert_v1beta1_AddressGroupPatch_To_networking_AddressGroupPatch is an autogenerated conversion function.
func Convert_v1beta1_AddressGroupPatch_To_networking_AddressGroupPatch(in *AddressGroupPatch, out *networking.AddressGroupPatch, s conversion.Scope) error {
	return autoConvert_v1beta1_AddressGroupPatch_To_networking_AddressGroupPatch(in, out, s)
}

func autoConvert_networking_AddressGroupPatch_To_v1beta1_AddressGroupPatch(in *networking.AddressGroupPatch, out *AddressGroupPatch, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.AddedPods = *(*[]GroupMemberPod)(unsafe.Pointer(&in.AddedPods))
	out.RemovedPods = *(*[]GroupMemberPod)(unsafe.Pointer(&in.RemovedPods))
	return nil
}

// Convert_networking_AddressGroupPatch_To_v1beta1_AddressGroupPatch is an autogenerated conversion function.
func Convert_networking_AddressGroupPatch_To_v1beta1_AddressGroupPatch(in *networking.AddressGroupPatch, out *AddressGroupPatch, s conversion.Scope) error {
	return autoConvert_networking_AddressGroupPatch_To_v1beta1_AddressGroupPatch(in, out, s)
}

func autoConvert_v1beta1_AppliedToGroup_To_networking_AppliedToGroup(in *AppliedToGroup, out *networking.AppliedToGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Pods = *(*[]networking.GroupMemberPod)(unsafe.Pointer(&in.Pods))
	return nil
}

// Convert_v1beta1_AppliedToGroup_To_networking_AppliedToGroup is an autogenerated conversion function.
func Convert_v1beta1_AppliedToGroup_To_networking_AppliedToGroup(in *AppliedToGroup, out *networking.AppliedToGroup, s conversion.Scope) error {
	return autoConvert_v1beta1_AppliedToGroup_To_networking_AppliedToGroup(in, out, s)
}

func autoConvert_networking_AppliedToGroup_To_v1beta1_AppliedToGroup(in *networking.AppliedToGroup, out *AppliedToGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Pods = *(*[]GroupMemberPod)(unsafe.Pointer(&in.Pods))
	return nil
}

// Convert_networking_AppliedToGroup_To_v1beta1_AppliedToGroup is an autogenerated conversion function.
func Convert_networking_AppliedToGroup_To_v1beta1_AppliedToGroup(in *networking.AppliedToGroup, out *AppliedToGroup, s conversion.Scope) error {
	return autoConvert_networking_AppliedToGroup_To_v1beta1_AppliedToGroup(in, out, s)
}

func autoConvert_v1beta1_AppliedToGroupList_To_networking_AppliedToGroupList(in *AppliedToGroupList, out *networking.AppliedToGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.AppliedToGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_AppliedToGroupList_To_networking_AppliedToGroupList is an autogenerated conversion function.
func Convert_v1beta1_AppliedToGroupList_To_networking_AppliedToGroupList(in *AppliedToGroupList, out *networking.AppliedToGroupList, s conversion.Scope) error {
	return autoConvert_v1beta1_AppliedToGroupList_To_networking_AppliedToGroupList(in, out, s)
}

func autoConvert_networking_AppliedToGroupList_To_v1beta1_AppliedToGroupList(in *networking.AppliedToGroupList, out *AppliedToGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]AppliedToGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_networking_AppliedToGroupList_To_v1beta1_AppliedToGroupList is an autogenerated conversion function.
func Convert_networking_AppliedToGroupList_To_v1beta1_AppliedToGroupList(in *networking.AppliedToGroupList, out *AppliedToGroupList, s conversion.Scope) error {
	return autoConvert_networking_AppliedToGroupList_To_v1beta1_AppliedToGroupList(in, out, s)
}

func autoConvert_v1beta1_AppliedToGroupPatch_To_networking_AppliedToGroupPatch(in *AppliedToGroupPatch, out *networking.AppliedToGroupPatch, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.AddedPods = *(*[]networking.GroupMemberPod)(unsafe.Pointer(&in.AddedPods))
	out.RemovedPods = *(*[]networking.GroupMemberPod)(unsafe.Pointer(&in.RemovedPods))
	return nil
}

// Convert_v1beta1_AppliedToGroupPatch_To_networking_AppliedToGroupPatch is an autogenerated conversion function.
func Convert_v1beta1_AppliedToGroupPatch_To_networking_AppliedToGroupPatch(in *AppliedToGroupPatch, out *networking.AppliedToGroupPatch, s conversion.Scope) error {
	return autoConvert_v1beta1_AppliedToGroupPatch_To_networking_AppliedToGroupPatch(in, out, s)
}

func autoConvert_networking_AppliedToGroupPatch_To_v1beta1_AppliedToGroupPatch(in *networking.AppliedToGroupPatch, out *AppliedToGroupPatch, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.AddedPods = *(*[]GroupMemberPod)(unsafe.Pointer(&in.AddedPods))
	out.RemovedPods = *(*[]GroupMemberPod)(unsafe.Pointer(&in.RemovedPods))
	return nil
}

// Convert_networking_AppliedToGroupPatch_To_v1beta1_AppliedToGroupPatch is an autogenerated conversion function.
func Convert_networking_AppliedToGroupPatch_To_v1beta1_AppliedToGroupPatch(in *networking.AppliedToGroupPatch, out *AppliedToGroupPatch, s conversion.Scope) error {
	return autoConvert_networking_AppliedToGroupPatch_To_v1beta1_AppliedToGroupPatch(in, out, s)
}

func autoConvert_v1beta1_GroupMemberPod_To_networking_GroupMemberPod(in *GroupMemberPod, out *networking.GroupMemberPod, s conversion.Scope) error {
	out.Pod = (*networking.PodReference)(unsafe.Pointer(in.Pod))
	out.IP = *(*networking.IPAddress)(unsafe.Pointer(&in.IP))
	out.Ports = *(*[]networking.NamedPort)(unsafe.Pointer(&in.Ports))
	return nil
}

// Convert_v1beta1_GroupMemberPod_To_networking_GroupMemberPod is an autogenerated conversion function.
func Convert_v1beta1_GroupMemberPod_To_networking_GroupMemberPod(in *GroupMemberPod, out *networking.GroupMemberPod, s conversion.Scope) error {
	return autoConvert_v1beta1_GroupMemberPod_To_networking_GroupMemberPod(in, out, s)
}

func autoConvert_networking_GroupMemberPod_To_v1beta1_GroupMemberPod(in *networking.GroupMemberPod, out *GroupMemberPod, s conversion.Scope) error {
	out.Pod = (*PodReference)(unsafe.Pointer(in.Pod))
	out.IP = *(*IPAddress)(unsafe.Pointer(&in.IP))
	out.Ports = *(*[]NamedPort)(unsafe.Pointer(&in.Ports))
	return nil
}

// Convert_networking_GroupMemberPod_To_v1beta1_GroupMemberPod is an autogenerated conversion function.
func Convert_networking_GroupMemberPod_To_v1beta1_GroupMemberPod(in *networking.GroupMemberPod, out *GroupMemberPod, s conversion.Scope) error {
	return autoConvert_networking_GroupMemberPod_To_v1beta1_GroupMemberPod(in, out, s)
}

func autoConvert_v1beta1_IPBlock_To_networking_IPBlock(in *IPBlock, out *networking.IPBlock, s conversion.Scope) error {
	if err := Convert_v1beta1_IPNet_To_networking_IPNet(&in.CIDR, &out.CIDR, s); err != nil {
		return err
	}
	out.Except = *(*[]networking.IPNet)(unsafe.Pointer(&in.Except))
	return nil
}

// Convert_v1beta1_IPBlock_To_networking_IPBlock is an autogenerated conversion function.
func Convert_v1beta1_IPBlock_To_networking_IPBlock(in *IPBlock, out *networking.IPBlock, s conversion.Scope) error {
	return autoConvert_v1beta1_IPBlock_To_networking_IPBlock(in, out, s)
}

func autoConvert_networking_IPBlock_To_v1beta1_IPBlock(in *networking.IPBlock, out *IPBlock, s conversion.Scope) error {
	if err := Convert_networking_IPNet_To_v1beta1_IPNet(&in.CIDR, &out.CIDR, s); err != nil {
		return err
	}
	out.Except = *(*[]IPNet)(unsafe.Pointer(&in.Except))
	return nil
}

// Convert_networking_IPBlock_To_v1beta1_IPBlock is an autogenerated conversion function.
func Convert_networking_IPBlock_To_v1beta1_IPBlock(in *networking.IPBlock, out *IPBlock, s conversion.Scope) error {
	return autoConvert_networking_IPBlock_To_v1beta1_IPBlock(in, out, s)
}

func autoConvert_v1beta1_IPNet_To_networking_IPNet(in *IPNet, out *networking.IPNet, s conversion.Scope) error {
	out.IP = *(*networking.IPAddress)(unsafe.Pointer(&in.IP))
	out.PrefixLength = in.PrefixLength
	return nil
}

// Convert_v1beta1_IPNet_To_networking_IPNet is an autogenerated conversion function.
func Convert_v1beta1_IPNet_To_networking_IPNet(in *IPNet, out *networking.IPNet, s conversion.Scope) error {
	return autoConvert_v1beta1_IPNet_To_networking_IPNet(in, out, s)
}

func autoConvert_networking_IPNet_To_v1beta1_IPNet(in *networking.IPNet, out *IPNet, s conversion.Scope) error {
	out.IP = *(*IPAddress)(unsafe.Pointer(&in.IP))
	out.PrefixLength = in.PrefixLength
	return nil
}

// Convert_networking_IPNet_To_v1beta1_IPNet is an autogenerated conversion function.
func Convert_networking_IPNet_To_v1beta1_IPNet(in *networking.IPNet, out *IPNet, s conversion.Scope) error {
	return autoConvert_networking_IPNet_To_v1beta1_IPNet(in, out, s)
}

func autoConvert_v1beta1_NamedPort_To_networking_NamedPort(in *NamedPort, out *networking.NamedPort, s conversion.Scope) error {
	out.Port = in.Port
	out.Name = in.Name
	out.Protocol = networking.Protocol(in.Protocol)
	return nil
}

// Convert_v1beta1_NamedPort_To_networking_NamedPort is an autogenerated conversion function.
func Convert_v1beta1_NamedPort_To_networking_NamedPort(in *NamedPort, out *networking.NamedPort, s conversion.Scope) error {
	return autoConvert_v1beta1_NamedPort_To_networking_NamedPort(in, out, s)
}

func autoConvert_networking_NamedPort_To_v1beta1_NamedPort(in *networking.NamedPort, out *NamedPort, s conversion.Scope) error {
	out.Port = in.Port
	out.Name = in.Name
	out.Protocol = Protocol(in.Protocol)
	return nil
}

// Convert_networking_NamedPort_To_v1beta1_NamedPort is an autogenerated conversion function.
func Convert_networking_NamedPort_To_v1beta1_NamedPort(in *networking.NamedPort, out *NamedPort, s conversion.Scope) error {
	return autoConvert_networking_NamedPort_To_v1beta1_NamedPort(in, out, s)
}

func autoConvert_v1beta1_NetworkPolicy_To_networking_NetworkPolicy(in *NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]networking.NetworkPolicyRule)(unsafe.Pointer(&in.Rules))
	out.AppliedToGroups = *(*[]string)(unsafe.Pointer(&in.AppliedToGroups))
	out.Priority = (*float64)(unsafe.Pointer(in.Priority))
	return nil
}

// Convert_v1beta1_NetworkPolicy_To_networking_NetworkPolicy is an autogenerated conversion function.
func Convert_v1beta1_NetworkPolicy_To_networking_NetworkPolicy(in *NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkPolicy_To_networking_NetworkPolicy(in, out, s)
}

func autoConvert_networking_NetworkPolicy_To_v1beta1_NetworkPolicy(in *networking.NetworkPolicy, out *NetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]NetworkPolicyRule)(unsafe.Pointer(&in.Rules))
	out.AppliedToGroups = *(*[]string)(unsafe.Pointer(&in.AppliedToGroups))
	out.Priority = (*float64)(unsafe.Pointer(in.Priority))
	return nil
}

// Convert_networking_NetworkPolicy_To_v1beta1_NetworkPolicy is an autogenerated conversion function.
func Convert_networking_NetworkPolicy_To_v1beta1_NetworkPolicy(in *networking.NetworkPolicy, out *NetworkPolicy, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicy_To_v1beta1_NetworkPolicy(in, out, s)
}

func autoConvert_v1beta1_NetworkPolicyList_To_networking_NetworkPolicyList(in *NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.NetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_NetworkPolicyList_To_networking_NetworkPolicyList is an autogenerated conversion function.
func Convert_v1beta1_NetworkPolicyList_To_networking_NetworkPolicyList(in *NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkPolicyList_To_networking_NetworkPolicyList(in, out, s)
}

func autoConvert_networking_NetworkPolicyList_To_v1beta1_NetworkPolicyList(in *networking.NetworkPolicyList, out *NetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]NetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_networking_NetworkPolicyList_To_v1beta1_NetworkPolicyList is an autogenerated conversion function.
func Convert_networking_NetworkPolicyList_To_v1beta1_NetworkPolicyList(in *networking.NetworkPolicyList, out *NetworkPolicyList, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyList_To_v1beta1_NetworkPolicyList(in, out, s)
}

func autoConvert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
	out.AddressGroups = *(*[]string)(unsafe.Pointer(&in.AddressGroups))
	out.IPBlocks = *(*[]networking.IPBlock)(unsafe.Pointer(&in.IPBlocks))
	return nil
}

// Convert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer is an autogenerated conversion function.
func Convert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in, out, s)
}

func autoConvert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *NetworkPolicyPeer, s conversion.Scope) error {
	out.AddressGroups = *(*[]string)(unsafe.Pointer(&in.AddressGroups))
	out.IPBlocks = *(*[]IPBlock)(unsafe.Pointer(&in.IPBlocks))
	return nil
}

// Convert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer is an autogenerated conversion function.
func Convert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *NetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer(in, out, s)
}

func autoConvert_v1beta1_NetworkPolicyRule_To_networking_NetworkPolicyRule(in *NetworkPolicyRule, out *networking.NetworkPolicyRule, s conversion.Scope) error {
	out.Direction = networking.Direction(in.Direction)
	if err := Convert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(&in.From, &out.From, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(&in.To, &out.To, s); err != nil {
		return err
	}
	out.Services = *(*[]networking.Service)(unsafe.Pointer(&in.Services))
	out.Priority = in.Priority
	out.Action = (*v1alpha1.RuleAction)(unsafe.Pointer(in.Action))
	return nil
}

// Convert_v1beta1_NetworkPolicyRule_To_networking_NetworkPolicyRule is an autogenerated conversion function.
func Convert_v1beta1_NetworkPolicyRule_To_networking_NetworkPolicyRule(in *NetworkPolicyRule, out *networking.NetworkPolicyRule, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkPolicyRule_To_networking_NetworkPolicyRule(in, out, s)
}

func autoConvert_networking_NetworkPolicyRule_To_v1beta1_NetworkPolicyRule(in *networking.NetworkPolicyRule, out *NetworkPolicyRule, s conversion.Scope) error {
	out.Direction = Direction(in.Direction)
	if err := Convert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer(&in.From, &out.From, s); err != nil {
		return err
	}
	if err := Convert_networking_NetworkPolicyPeer_To_v1beta1_NetworkPolicyPeer(&in.To, &out.To, s); err != nil {
		return err
	}
	out.Services = *(*[]Service)(unsafe.Pointer(&in.Services))
	out.Priority = in.Priority
	out.Action = (*v1alpha1.RuleAction)(unsafe.Pointer(in.Action))
	return nil
}

// Convert_networking_NetworkPolicyRule_To_v1beta1_NetworkPolicyRule is an autogenerated conversion function.
func Convert_networking_NetworkPolicyRule_To_v1beta1_NetworkPolicyRule(in *networking.NetworkPolicyRule, out *NetworkPolicyRule, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyRule_To_v1beta1_NetworkPolicyRule(in, out, s)
}

func autoConvert_v1beta1_PodReference_To_networking_PodReference(in *PodReference, out *networking.PodReference, s conversion.Scope) error {
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_v1beta1_PodReference_To_networking_PodReference is an autogenerated conversion function.
func Convert_v1beta1_PodReference_To_networking_PodReference(in *PodReference, out *networking.PodReference, s conversion.Scope) error {
	return autoConvert_v1beta1_PodReference_To_networking_PodReference(in, out, s)
}

func autoConvert_networking_PodReference_To_v1beta1_PodReference(in *networking.PodReference, out *PodReference, s conversion.Scope) error {
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_networking_PodReference_To_v1beta1_PodReference is an autogenerated conversion function.
func Convert_networking_PodReference_To_v1beta1_PodReference(in *networking.PodReference, out *PodReference, s conversion.Scope) error {
	return autoConvert_networking_PodReference_To_v1beta1_PodReference(in, out, s)
}

func autoConvert_v1beta1_Service_To_networking_Service(in *Service, out *networking.Service, s conversion.Scope) error {
	out.Protocol = (*networking.Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
	return nil
}

// Convert_v1beta1_Service_To_networking_Service is an autogenerated conversion function.
func Convert_v1beta1_Service_To_networking_Service(in *Service, out *networking.Service, s conversion.Scope) error {
	return autoConvert_v1beta1_Service_To_networking_Service(in, out, s)
}

func autoConvert_networking_Service_To_v1beta1_Service(in *networking.Service, out *Service, s conversion.Scope) error {
	out.Protocol = (*Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
	return nil
}

// Convert_networking_Service_To_v1beta1_Service is an autogenerated conversion function.
func Convert_networking_Service_To_v1beta1_Service(in *networking.Service, out *Service, s conversion.Scope) error {
	return autoConvert_networking_Service_To_v1beta1_Service(in, out, s)
}
