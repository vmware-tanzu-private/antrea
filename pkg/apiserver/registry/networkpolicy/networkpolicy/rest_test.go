// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkpolicy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/antrea/pkg/apis/controlplane"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/vmware-tanzu/antrea/pkg/controller/networkpolicy/store"
	"github.com/vmware-tanzu/antrea/pkg/controller/types"
)

func TestRESTList(t *testing.T) {
	tests := []struct {
		name            string
		networkPolicies []*types.NetworkPolicy
		labelSelector   labels.Selector
		expectedObj     runtime.Object
	}{
		{
			name: "label selector selecting nothing",
			networkPolicies: []*types.NetworkPolicy{
				{
					Name: "foo",
				},
			},
			labelSelector: labels.Nothing(),
			expectedObj:   &controlplane.NetworkPolicyList{},
		},
		{
			name: "label selector selecting everything",
			networkPolicies: []*types.NetworkPolicy{
				{
					Name: "foo",
				},
			},
			labelSelector: labels.Everything(),
			expectedObj: &controlplane.NetworkPolicyList{
				Items: []controlplane.NetworkPolicy{
					{
						ObjectMeta: v1.ObjectMeta{
							Name: "foo",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := store.NewNetworkPolicyStore()
			for _, obj := range tt.networkPolicies {
				storage.Create(obj)
			}
			r := NewREST(storage)
			actualObj, err := r.List(context.TODO(), &internalversion.ListOptions{LabelSelector: tt.labelSelector})
			assert.NoError(t, err)
			assert.ElementsMatch(t, tt.expectedObj.(*controlplane.NetworkPolicyList).Items, actualObj.(*controlplane.NetworkPolicyList).Items)
		})
	}
}
