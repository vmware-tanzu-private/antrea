/*
Copyright 2021 Antrea Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceExports implements ResourceExportInterface
type FakeResourceExports struct {
	Fake *FakeMulticlusterV1alpha1
	ns   string
}

var resourceexportsResource = schema.GroupVersionResource{Group: "multicluster.crd.antrea.io", Version: "v1alpha1", Resource: "resourceexports"}

var resourceexportsKind = schema.GroupVersionKind{Group: "multicluster.crd.antrea.io", Version: "v1alpha1", Kind: "ResourceExport"}

// Get takes name of the resourceExport, and returns the corresponding resourceExport object, and an error if there is any.
func (c *FakeResourceExports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourceexportsResource, c.ns, name), &v1alpha1.ResourceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceExport), err
}

// List takes label and field selectors, and returns the list of ResourceExports that match those selectors.
func (c *FakeResourceExports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceExportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourceexportsResource, resourceexportsKind, c.ns, opts), &v1alpha1.ResourceExportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceExportList{ListMeta: obj.(*v1alpha1.ResourceExportList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceExportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceExports.
func (c *FakeResourceExports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourceexportsResource, c.ns, opts))

}

// Create takes the representation of a resourceExport and creates it.  Returns the server's representation of the resourceExport, and an error, if there is any.
func (c *FakeResourceExports) Create(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.CreateOptions) (result *v1alpha1.ResourceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourceexportsResource, c.ns, resourceExport), &v1alpha1.ResourceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceExport), err
}

// Update takes the representation of a resourceExport and updates it. Returns the server's representation of the resourceExport, and an error, if there is any.
func (c *FakeResourceExports) Update(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.UpdateOptions) (result *v1alpha1.ResourceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourceexportsResource, c.ns, resourceExport), &v1alpha1.ResourceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceExport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceExports) UpdateStatus(ctx context.Context, resourceExport *v1alpha1.ResourceExport, opts v1.UpdateOptions) (*v1alpha1.ResourceExport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourceexportsResource, "status", c.ns, resourceExport), &v1alpha1.ResourceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceExport), err
}

// Delete takes name of the resourceExport and deletes it. Returns an error if one occurs.
func (c *FakeResourceExports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourceexportsResource, c.ns, name, opts), &v1alpha1.ResourceExport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceExports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourceexportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceExportList{})
	return err
}

// Patch applies the patch and returns the patched resourceExport.
func (c *FakeResourceExports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourceexportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourceExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceExport), err
}
