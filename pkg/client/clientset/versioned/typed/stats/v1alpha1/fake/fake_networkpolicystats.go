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
	"context"

	v1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/stats/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkPolicyStats implements NetworkPolicyStatsInterface
type FakeNetworkPolicyStats struct {
	Fake *FakeStatsV1alpha1
	ns   string
}

var networkpolicystatsResource = schema.GroupVersionResource{Group: "stats.antrea.tanzu.vmware.com", Version: "v1alpha1", Resource: "networkpolicystats"}

var networkpolicystatsKind = schema.GroupVersionKind{Group: "stats.antrea.tanzu.vmware.com", Version: "v1alpha1", Kind: "NetworkPolicyStats"}

// Get takes name of the networkPolicyStats, and returns the corresponding networkPolicyStats object, and an error if there is any.
func (c *FakeNetworkPolicyStats) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkPolicyStats, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkpolicystatsResource, c.ns, name), &v1alpha1.NetworkPolicyStats{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPolicyStats), err
}

// List takes label and field selectors, and returns the list of NetworkPolicyStats that match those selectors.
func (c *FakeNetworkPolicyStats) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkPolicyStatsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkpolicystatsResource, networkpolicystatsKind, c.ns, opts), &v1alpha1.NetworkPolicyStatsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkPolicyStatsList{ListMeta: obj.(*v1alpha1.NetworkPolicyStatsList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkPolicyStatsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPolicyStats.
func (c *FakeNetworkPolicyStats) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkpolicystatsResource, c.ns, opts))

}
