//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2021 Antrea Authors
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	crdv1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
	"antrea.io/antrea/pkg/apis/crd/v1alpha2"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apisv1alpha1 "sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaim) DeepCopyInto(out *ClusterClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaim.
func (in *ClusterClaim) DeepCopy() *ClusterClaim {
	if in == nil {
		return nil
	}
	out := new(ClusterClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaimList) DeepCopyInto(out *ClusterClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaimList.
func (in *ClusterClaimList) DeepCopy() *ClusterClaimList {
	if in == nil {
		return nil
	}
	out := new(ClusterClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	if in.GatewayInfos != nil {
		in, out := &in.GatewayInfos, &out.GatewayInfos
		*out = make([]GatewayInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoImport) DeepCopyInto(out *ClusterInfoImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoImport.
func (in *ClusterInfoImport) DeepCopy() *ClusterInfoImport {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInfoImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoImportList) DeepCopyInto(out *ClusterInfoImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterInfoImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoImportList.
func (in *ClusterInfoImportList) DeepCopy() *ClusterInfoImportList {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInfoImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoImportStatus) DeepCopyInto(out *ClusterInfoImportStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoImportStatus.
func (in *ClusterInfoImportStatus) DeepCopy() *ClusterInfoImportStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSet) DeepCopyInto(out *ClusterSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSet.
func (in *ClusterSet) DeepCopy() *ClusterSet {
	if in == nil {
		return nil
	}
	out := new(ClusterSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetCondition) DeepCopyInto(out *ClusterSetCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetCondition.
func (in *ClusterSetCondition) DeepCopy() *ClusterSetCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterSetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetList) DeepCopyInto(out *ClusterSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetList.
func (in *ClusterSetList) DeepCopy() *ClusterSetList {
	if in == nil {
		return nil
	}
	out := new(ClusterSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetSpec) DeepCopyInto(out *ClusterSetSpec) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]MemberCluster, len(*in))
		copy(*out, *in)
	}
	if in.Leaders != nil {
		in, out := &in.Leaders, &out.Leaders
		*out = make([]MemberCluster, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetSpec.
func (in *ClusterSetSpec) DeepCopy() *ClusterSetSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetStatus) DeepCopyInto(out *ClusterSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterSetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterStatuses != nil {
		in, out := &in.ClusterStatuses, &out.ClusterStatuses
		*out = make([]ClusterStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetStatus.
func (in *ClusterSetStatus) DeepCopy() *ClusterSetStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointsExport) DeepCopyInto(out *EndpointsExport) {
	*out = *in
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]v1.EndpointSubset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointsExport.
func (in *EndpointsExport) DeepCopy() *EndpointsExport {
	if in == nil {
		return nil
	}
	out := new(EndpointsExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointsImport) DeepCopyInto(out *EndpointsImport) {
	*out = *in
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]v1.EndpointSubset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointsImport.
func (in *EndpointsImport) DeepCopy() *EndpointsImport {
	if in == nil {
		return nil
	}
	out := new(EndpointsImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEntityExport) DeepCopyInto(out *ExternalEntityExport) {
	*out = *in
	in.ExternalEntitySpec.DeepCopyInto(&out.ExternalEntitySpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEntityExport.
func (in *ExternalEntityExport) DeepCopy() *ExternalEntityExport {
	if in == nil {
		return nil
	}
	out := new(ExternalEntityExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEntityImport) DeepCopyInto(out *ExternalEntityImport) {
	*out = *in
	if in.ExternalEntitySpec != nil {
		in, out := &in.ExternalEntitySpec, &out.ExternalEntitySpec
		*out = new(v1alpha2.ExternalEntitySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEntityImport.
func (in *ExternalEntityImport) DeepCopy() *ExternalEntityImport {
	if in == nil {
		return nil
	}
	out := new(ExternalEntityImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayInfo) DeepCopyInto(out *GatewayInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayInfo.
func (in *GatewayInfo) DeepCopy() *GatewayInfo {
	if in == nil {
		return nil
	}
	out := new(GatewayInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayList) DeepCopyInto(out *GatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayList.
func (in *GatewayList) DeepCopy() *GatewayList {
	if in == nil {
		return nil
	}
	out := new(GatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberCluster) DeepCopyInto(out *MemberCluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberCluster.
func (in *MemberCluster) DeepCopy() *MemberCluster {
	if in == nil {
		return nil
	}
	out := new(MemberCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberClusterAnnounce) DeepCopyInto(out *MemberClusterAnnounce) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberClusterAnnounce.
func (in *MemberClusterAnnounce) DeepCopy() *MemberClusterAnnounce {
	if in == nil {
		return nil
	}
	out := new(MemberClusterAnnounce)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemberClusterAnnounce) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberClusterAnnounceList) DeepCopyInto(out *MemberClusterAnnounceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MemberClusterAnnounce, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberClusterAnnounceList.
func (in *MemberClusterAnnounceList) DeepCopy() *MemberClusterAnnounceList {
	if in == nil {
		return nil
	}
	out := new(MemberClusterAnnounceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemberClusterAnnounceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterConfig) DeepCopyInto(out *MultiClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterConfig.
func (in *MultiClusterConfig) DeepCopy() *MultiClusterConfig {
	if in == nil {
		return nil
	}
	out := new(MultiClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawResourceExport) DeepCopyInto(out *RawResourceExport) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawResourceExport.
func (in *RawResourceExport) DeepCopy() *RawResourceExport {
	if in == nil {
		return nil
	}
	out := new(RawResourceExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawResourceImport) DeepCopyInto(out *RawResourceImport) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawResourceImport.
func (in *RawResourceImport) DeepCopy() *RawResourceImport {
	if in == nil {
		return nil
	}
	out := new(RawResourceImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceCondition) DeepCopyInto(out *ResourceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCondition.
func (in *ResourceCondition) DeepCopy() *ResourceCondition {
	if in == nil {
		return nil
	}
	out := new(ResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExport) DeepCopyInto(out *ResourceExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExport.
func (in *ResourceExport) DeepCopy() *ResourceExport {
	if in == nil {
		return nil
	}
	out := new(ResourceExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExportCondition) DeepCopyInto(out *ResourceExportCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExportCondition.
func (in *ResourceExportCondition) DeepCopy() *ResourceExportCondition {
	if in == nil {
		return nil
	}
	out := new(ResourceExportCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExportList) DeepCopyInto(out *ResourceExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExportList.
func (in *ResourceExportList) DeepCopy() *ResourceExportList {
	if in == nil {
		return nil
	}
	out := new(ResourceExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExportSpec) DeepCopyInto(out *ResourceExportSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceExport)
		(*in).DeepCopyInto(*out)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = new(EndpointsExport)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = new(ClusterInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalEntity != nil {
		in, out := &in.ExternalEntity, &out.ExternalEntity
		*out = new(ExternalEntityExport)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterNetworkPolicy != nil {
		in, out := &in.ClusterNetworkPolicy, &out.ClusterNetworkPolicy
		*out = new(crdv1alpha1.ClusterNetworkPolicySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = new(RawResourceExport)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExportSpec.
func (in *ResourceExportSpec) DeepCopy() *ResourceExportSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExportStatus) DeepCopyInto(out *ResourceExportStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ResourceExportCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExportStatus.
func (in *ResourceExportStatus) DeepCopy() *ResourceExportStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceImport) DeepCopyInto(out *ResourceImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceImport.
func (in *ResourceImport) DeepCopy() *ResourceImport {
	if in == nil {
		return nil
	}
	out := new(ResourceImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceImportClusterStatus) DeepCopyInto(out *ResourceImportClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ResourceImportCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceImportClusterStatus.
func (in *ResourceImportClusterStatus) DeepCopy() *ResourceImportClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceImportClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceImportCondition) DeepCopyInto(out *ResourceImportCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceImportCondition.
func (in *ResourceImportCondition) DeepCopy() *ResourceImportCondition {
	if in == nil {
		return nil
	}
	out := new(ResourceImportCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceImportList) DeepCopyInto(out *ResourceImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceImportList.
func (in *ResourceImportList) DeepCopy() *ResourceImportList {
	if in == nil {
		return nil
	}
	out := new(ResourceImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceImportSpec) DeepCopyInto(out *ResourceImportSpec) {
	*out = *in
	if in.ClusterIDs != nil {
		in, out := &in.ClusterIDs, &out.ClusterIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceImport != nil {
		in, out := &in.ServiceImport, &out.ServiceImport
		*out = new(apisv1alpha1.ServiceImport)
		(*in).DeepCopyInto(*out)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = new(EndpointsImport)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = new(ClusterInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalEntity != nil {
		in, out := &in.ExternalEntity, &out.ExternalEntity
		*out = new(ExternalEntityImport)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterNetworkPolicy != nil {
		in, out := &in.ClusterNetworkPolicy, &out.ClusterNetworkPolicy
		*out = new(crdv1alpha1.ClusterNetworkPolicySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = new(RawResourceImport)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceImportSpec.
func (in *ResourceImportSpec) DeepCopy() *ResourceImportSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceImportStatus) DeepCopyInto(out *ResourceImportStatus) {
	*out = *in
	if in.ClusterStatuses != nil {
		in, out := &in.ClusterStatuses, &out.ClusterStatuses
		*out = make([]ResourceImportClusterStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceImportStatus.
func (in *ResourceImportStatus) DeepCopy() *ResourceImportStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExport) DeepCopyInto(out *ServiceExport) {
	*out = *in
	in.ServiceSpec.DeepCopyInto(&out.ServiceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExport.
func (in *ServiceExport) DeepCopy() *ServiceExport {
	if in == nil {
		return nil
	}
	out := new(ServiceExport)
	in.DeepCopyInto(out)
	return out
}
