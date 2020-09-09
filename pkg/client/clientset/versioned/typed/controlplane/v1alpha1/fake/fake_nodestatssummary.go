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

	v1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/controlplane/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeNodeStatsSummaries implements NodeStatsSummaryInterface
type FakeNodeStatsSummaries struct {
	Fake *FakeControlplaneV1alpha1
}

var nodestatssummariesResource = schema.GroupVersionResource{Group: "controlplane.antrea.tanzu.vmware.com", Version: "v1alpha1", Resource: "nodestatssummaries"}

var nodestatssummariesKind = schema.GroupVersionKind{Group: "controlplane.antrea.tanzu.vmware.com", Version: "v1alpha1", Kind: "NodeStatsSummary"}

// Create takes the representation of a nodeStatsSummary and creates it.  Returns the server's representation of the nodeStatsSummary, and an error, if there is any.
func (c *FakeNodeStatsSummaries) Create(ctx context.Context, nodeStatsSummary *v1alpha1.NodeStatsSummary, opts v1.CreateOptions) (result *v1alpha1.NodeStatsSummary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodestatssummariesResource, nodeStatsSummary), &v1alpha1.NodeStatsSummary{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeStatsSummary), err
}
