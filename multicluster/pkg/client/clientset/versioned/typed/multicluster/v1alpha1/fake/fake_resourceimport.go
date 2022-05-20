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

// FakeResourceImports implements ResourceImportInterface
type FakeResourceImports struct {
	Fake *FakeMulticlusterV1alpha1
	ns   string
}

var resourceimportsResource = schema.GroupVersionResource{Group: "multicluster.crd.antrea.io", Version: "v1alpha1", Resource: "resourceimports"}

var resourceimportsKind = schema.GroupVersionKind{Group: "multicluster.crd.antrea.io", Version: "v1alpha1", Kind: "ResourceImport"}

// Get takes name of the resourceImport, and returns the corresponding resourceImport object, and an error if there is any.
func (c *FakeResourceImports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceImport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourceimportsResource, c.ns, name), &v1alpha1.ResourceImport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceImport), err
}

// List takes label and field selectors, and returns the list of ResourceImports that match those selectors.
func (c *FakeResourceImports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceImportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourceimportsResource, resourceimportsKind, c.ns, opts), &v1alpha1.ResourceImportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceImportList{ListMeta: obj.(*v1alpha1.ResourceImportList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceImportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceImports.
func (c *FakeResourceImports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourceimportsResource, c.ns, opts))

}

// Create takes the representation of a resourceImport and creates it.  Returns the server's representation of the resourceImport, and an error, if there is any.
func (c *FakeResourceImports) Create(ctx context.Context, resourceImport *v1alpha1.ResourceImport, opts v1.CreateOptions) (result *v1alpha1.ResourceImport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourceimportsResource, c.ns, resourceImport), &v1alpha1.ResourceImport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceImport), err
}

// Update takes the representation of a resourceImport and updates it. Returns the server's representation of the resourceImport, and an error, if there is any.
func (c *FakeResourceImports) Update(ctx context.Context, resourceImport *v1alpha1.ResourceImport, opts v1.UpdateOptions) (result *v1alpha1.ResourceImport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourceimportsResource, c.ns, resourceImport), &v1alpha1.ResourceImport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceImport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceImports) UpdateStatus(ctx context.Context, resourceImport *v1alpha1.ResourceImport, opts v1.UpdateOptions) (*v1alpha1.ResourceImport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourceimportsResource, "status", c.ns, resourceImport), &v1alpha1.ResourceImport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceImport), err
}

// Delete takes name of the resourceImport and deletes it. Returns an error if one occurs.
func (c *FakeResourceImports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourceimportsResource, c.ns, name, opts), &v1alpha1.ResourceImport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceImports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourceimportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceImportList{})
	return err
}

// Patch applies the patch and returns the patched resourceImport.
func (c *FakeResourceImports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceImport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourceimportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourceImport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceImport), err
}
