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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned"
	clusterinformationv1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/clusterinformation/v1beta1"
	fakeclusterinformationv1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/clusterinformation/v1beta1/fake"
	controlplanev1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta1"
	fakecontrolplanev1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta1/fake"
	controlplanev1beta2 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta2"
	fakecontrolplanev1beta2 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta2/fake"
	corev1alpha2 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/core/v1alpha2"
	fakecorev1alpha2 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/core/v1alpha2/fake"
	egressv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/egress/v1alpha1"
	fakeegressv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/egress/v1alpha1/fake"
	opsv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/ops/v1alpha1"
	fakeopsv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/ops/v1alpha1/fake"
	securityv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/security/v1alpha1"
	fakesecurityv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/security/v1alpha1/fake"
	statsv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/stats/v1alpha1"
	fakestatsv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/stats/v1alpha1/fake"
	systemv1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/system/v1beta1"
	fakesystemv1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/system/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ClusterinformationV1beta1 retrieves the ClusterinformationV1beta1Client
func (c *Clientset) ClusterinformationV1beta1() clusterinformationv1beta1.ClusterinformationV1beta1Interface {
	return &fakeclusterinformationv1beta1.FakeClusterinformationV1beta1{Fake: &c.Fake}
}

// ControlplaneV1beta1 retrieves the ControlplaneV1beta1Client
func (c *Clientset) ControlplaneV1beta1() controlplanev1beta1.ControlplaneV1beta1Interface {
	return &fakecontrolplanev1beta1.FakeControlplaneV1beta1{Fake: &c.Fake}
}

// ControlplaneV1beta2 retrieves the ControlplaneV1beta2Client
func (c *Clientset) ControlplaneV1beta2() controlplanev1beta2.ControlplaneV1beta2Interface {
	return &fakecontrolplanev1beta2.FakeControlplaneV1beta2{Fake: &c.Fake}
}

// CoreV1alpha2 retrieves the CoreV1alpha2Client
func (c *Clientset) CoreV1alpha2() corev1alpha2.CoreV1alpha2Interface {
	return &fakecorev1alpha2.FakeCoreV1alpha2{Fake: &c.Fake}
}

// EgressV1alpha1 retrieves the EgressV1alpha1Client
func (c *Clientset) EgressV1alpha1() egressv1alpha1.EgressV1alpha1Interface {
	return &fakeegressv1alpha1.FakeEgressV1alpha1{Fake: &c.Fake}
}

// OpsV1alpha1 retrieves the OpsV1alpha1Client
func (c *Clientset) OpsV1alpha1() opsv1alpha1.OpsV1alpha1Interface {
	return &fakeopsv1alpha1.FakeOpsV1alpha1{Fake: &c.Fake}
}

// SecurityV1alpha1 retrieves the SecurityV1alpha1Client
func (c *Clientset) SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface {
	return &fakesecurityv1alpha1.FakeSecurityV1alpha1{Fake: &c.Fake}
}

// StatsV1alpha1 retrieves the StatsV1alpha1Client
func (c *Clientset) StatsV1alpha1() statsv1alpha1.StatsV1alpha1Interface {
	return &fakestatsv1alpha1.FakeStatsV1alpha1{Fake: &c.Fake}
}

// SystemV1beta1 retrieves the SystemV1beta1Client
func (c *Clientset) SystemV1beta1() systemv1beta1.SystemV1beta1Interface {
	return &fakesystemv1beta1.FakeSystemV1beta1{Fake: &c.Fake}
}
