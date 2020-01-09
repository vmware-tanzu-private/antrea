// Copyright 2019 Antrea Authors
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
	v1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/cleanup/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCleanupStatuses implements CleanupStatusInterface
type FakeCleanupStatuses struct {
	Fake *FakeCleanupV1beta1
}

var cleanupstatusesResource = schema.GroupVersionResource{Group: "cleanup.antrea.tanzu.vmware.com", Version: "v1beta1", Resource: "cleanupstatuses"}

var cleanupstatusesKind = schema.GroupVersionKind{Group: "cleanup.antrea.tanzu.vmware.com", Version: "v1beta1", Kind: "CleanupStatus"}

// Get takes name of the cleanupStatus, and returns the corresponding cleanupStatus object, and an error if there is any.
func (c *FakeCleanupStatuses) Get(name string, options v1.GetOptions) (result *v1beta1.CleanupStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cleanupstatusesResource, name), &v1beta1.CleanupStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CleanupStatus), err
}

// List takes label and field selectors, and returns the list of CleanupStatuses that match those selectors.
func (c *FakeCleanupStatuses) List(opts v1.ListOptions) (result *v1beta1.CleanupStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cleanupstatusesResource, cleanupstatusesKind, opts), &v1beta1.CleanupStatusList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CleanupStatusList{ListMeta: obj.(*v1beta1.CleanupStatusList).ListMeta}
	for _, item := range obj.(*v1beta1.CleanupStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cleanupStatuses.
func (c *FakeCleanupStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cleanupstatusesResource, opts))
}

// Create takes the representation of a cleanupStatus and creates it.  Returns the server's representation of the cleanupStatus, and an error, if there is any.
func (c *FakeCleanupStatuses) Create(cleanupStatus *v1beta1.CleanupStatus) (result *v1beta1.CleanupStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cleanupstatusesResource, cleanupStatus), &v1beta1.CleanupStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CleanupStatus), err
}

// Update takes the representation of a cleanupStatus and updates it. Returns the server's representation of the cleanupStatus, and an error, if there is any.
func (c *FakeCleanupStatuses) Update(cleanupStatus *v1beta1.CleanupStatus) (result *v1beta1.CleanupStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cleanupstatusesResource, cleanupStatus), &v1beta1.CleanupStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CleanupStatus), err
}

// Delete takes name of the cleanupStatus and deletes it. Returns an error if one occurs.
func (c *FakeCleanupStatuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cleanupstatusesResource, name), &v1beta1.CleanupStatus{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCleanupStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cleanupstatusesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.CleanupStatusList{})
	return err
}

// Patch applies the patch and returns the patched cleanupStatus.
func (c *FakeCleanupStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CleanupStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cleanupstatusesResource, name, pt, data, subresources...), &v1beta1.CleanupStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CleanupStatus), err
}
