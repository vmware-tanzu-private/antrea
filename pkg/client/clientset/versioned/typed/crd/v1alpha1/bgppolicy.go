// Copyright 2024 Antrea Authors
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

package v1alpha1

import (
	"context"

	v1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BGPPoliciesGetter has a method to return a BGPPolicyInterface.
// A group's client should implement this interface.
type BGPPoliciesGetter interface {
	BGPPolicies() BGPPolicyInterface
}

// BGPPolicyInterface has methods to work with BGPPolicy resources.
type BGPPolicyInterface interface {
	Create(ctx context.Context, bGPPolicy *v1alpha1.BGPPolicy, opts v1.CreateOptions) (*v1alpha1.BGPPolicy, error)
	Update(ctx context.Context, bGPPolicy *v1alpha1.BGPPolicy, opts v1.UpdateOptions) (*v1alpha1.BGPPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BGPPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BGPPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BGPPolicy, err error)
	BGPPolicyExpansion
}

// bGPPolicies implements BGPPolicyInterface
type bGPPolicies struct {
	*gentype.ClientWithList[*v1alpha1.BGPPolicy, *v1alpha1.BGPPolicyList]
}

// newBGPPolicies returns a BGPPolicies
func newBGPPolicies(c *CrdV1alpha1Client) *bGPPolicies {
	return &bGPPolicies{
		gentype.NewClientWithList[*v1alpha1.BGPPolicy, *v1alpha1.BGPPolicyList](
			"bgppolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.BGPPolicy { return &v1alpha1.BGPPolicy{} },
			func() *v1alpha1.BGPPolicyList { return &v1alpha1.BGPPolicyList{} }),
	}
}