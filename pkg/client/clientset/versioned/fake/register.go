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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clusterinformationv1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/clusterinformation/v1beta1"
	controlplanev1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/controlplane/v1beta1"
	corev1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/core/v1alpha1"
	corev1alpha2 "github.com/vmware-tanzu/antrea/pkg/apis/core/v1alpha2"
	opsv1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/ops/v1alpha1"
	securityv1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/security/v1alpha1"
	statsv1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/stats/v1alpha1"
	systemv1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/system/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	clusterinformationv1beta1.AddToScheme,
	controlplanev1beta1.AddToScheme,
	corev1alpha1.AddToScheme,
	corev1alpha2.AddToScheme,
	opsv1alpha1.AddToScheme,
	securityv1alpha1.AddToScheme,
	statsv1alpha1.AddToScheme,
	systemv1beta1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
