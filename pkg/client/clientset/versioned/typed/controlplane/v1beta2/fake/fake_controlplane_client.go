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
	v1beta2 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeControlplaneV1beta2 struct {
	*testing.Fake
}

func (c *FakeControlplaneV1beta2) AddressGroups() v1beta2.AddressGroupInterface {
	return &FakeAddressGroups{c}
}

func (c *FakeControlplaneV1beta2) AppliedToGroups() v1beta2.AppliedToGroupInterface {
	return &FakeAppliedToGroups{c}
}

func (c *FakeControlplaneV1beta2) ClusterGroupMembers() v1beta2.ClusterGroupMembersInterface {
	return &FakeClusterGroupMembers{c}
}

func (c *FakeControlplaneV1beta2) EgressGroups() v1beta2.EgressGroupInterface {
	return &FakeEgressGroups{c}
}

func (c *FakeControlplaneV1beta2) GroupAssociations(namespace string) v1beta2.GroupAssociationInterface {
	return &FakeGroupAssociations{c, namespace}
}

func (c *FakeControlplaneV1beta2) NetworkPolicies() v1beta2.NetworkPolicyInterface {
	return &FakeNetworkPolicies{c}
}

func (c *FakeControlplaneV1beta2) NodeStatsSummaries() v1beta2.NodeStatsSummaryInterface {
	return &FakeNodeStatsSummaries{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeControlplaneV1beta2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
