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
	v1alpha2 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCrdV1alpha2 struct {
	*testing.Fake
}

func (c *FakeCrdV1alpha2) ClusterGroups() v1alpha2.ClusterGroupInterface {
	return &FakeClusterGroups{c}
}

func (c *FakeCrdV1alpha2) ClusterNetworkPolicies() v1alpha2.ClusterNetworkPolicyInterface {
	return &FakeClusterNetworkPolicies{c}
}

func (c *FakeCrdV1alpha2) Egresses() v1alpha2.EgressInterface {
	return &FakeEgresses{c}
}

func (c *FakeCrdV1alpha2) ExternalEntities(namespace string) v1alpha2.ExternalEntityInterface {
	return &FakeExternalEntities{c, namespace}
}

func (c *FakeCrdV1alpha2) ExternalIPPools() v1alpha2.ExternalIPPoolInterface {
	return &FakeExternalIPPools{c}
}

func (c *FakeCrdV1alpha2) IPPools() v1alpha2.IPPoolInterface {
	return &FakeIPPools{c}
}

func (c *FakeCrdV1alpha2) NetworkPolicies(namespace string) v1alpha2.NetworkPolicyInterface {
	return &FakeNetworkPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCrdV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
