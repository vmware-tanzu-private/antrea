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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta2 "antrea.io/antrea/pkg/apis/controlplane/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSupportBundleCollections implements SupportBundleCollectionInterface
type FakeSupportBundleCollections struct {
	Fake *FakeControlplaneV1beta2
}

var supportbundlecollectionsResource = schema.GroupVersionResource{Group: "controlplane.antrea.io", Version: "v1beta2", Resource: "supportbundlecollections"}

var supportbundlecollectionsKind = schema.GroupVersionKind{Group: "controlplane.antrea.io", Version: "v1beta2", Kind: "SupportBundleCollection"}

// Get takes name of the supportBundleCollection, and returns the corresponding supportBundleCollection object, and an error if there is any.
func (c *FakeSupportBundleCollections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.SupportBundleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(supportbundlecollectionsResource, name), &v1beta2.SupportBundleCollection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SupportBundleCollection), err
}

// List takes label and field selectors, and returns the list of SupportBundleCollections that match those selectors.
func (c *FakeSupportBundleCollections) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.SupportBundleCollectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(supportbundlecollectionsResource, supportbundlecollectionsKind, opts), &v1beta2.SupportBundleCollectionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.SupportBundleCollectionList{ListMeta: obj.(*v1beta2.SupportBundleCollectionList).ListMeta}
	for _, item := range obj.(*v1beta2.SupportBundleCollectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested supportBundleCollections.
func (c *FakeSupportBundleCollections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(supportbundlecollectionsResource, opts))
}
