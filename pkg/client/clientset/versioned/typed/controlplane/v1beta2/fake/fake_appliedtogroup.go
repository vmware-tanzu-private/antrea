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

	v1beta2 "github.com/vmware-tanzu/antrea/pkg/apis/controlplane/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppliedToGroups implements AppliedToGroupInterface
type FakeAppliedToGroups struct {
	Fake *FakeControlplaneV1beta2
}

var appliedtogroupsResource = schema.GroupVersionResource{Group: "controlplane.antrea.tanzu.vmware.com", Version: "v1beta2", Resource: "appliedtogroups"}

var appliedtogroupsKind = schema.GroupVersionKind{Group: "controlplane.antrea.tanzu.vmware.com", Version: "v1beta2", Kind: "AppliedToGroup"}

// Get takes name of the appliedToGroup, and returns the corresponding appliedToGroup object, and an error if there is any.
func (c *FakeAppliedToGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.AppliedToGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(appliedtogroupsResource, name), &v1beta2.AppliedToGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.AppliedToGroup), err
}

// List takes label and field selectors, and returns the list of AppliedToGroups that match those selectors.
func (c *FakeAppliedToGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.AppliedToGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(appliedtogroupsResource, appliedtogroupsKind, opts), &v1beta2.AppliedToGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.AppliedToGroupList{ListMeta: obj.(*v1beta2.AppliedToGroupList).ListMeta}
	for _, item := range obj.(*v1beta2.AppliedToGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appliedToGroups.
func (c *FakeAppliedToGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(appliedtogroupsResource, opts))
}
