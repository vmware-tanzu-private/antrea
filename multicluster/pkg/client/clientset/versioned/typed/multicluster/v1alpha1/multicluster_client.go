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

package v1alpha1

import (
	v1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	"antrea.io/antrea/multicluster/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type MulticlusterV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClusterClaimsGetter
	ClusterSetsGetter
	MemberClusterAnnouncesGetter
	ResourceExportsGetter
	ResourceExportFiltersGetter
	ResourceImportsGetter
	ResourceImportFiltersGetter
}

// MulticlusterV1alpha1Client is used to interact with features provided by the multicluster.crd.antrea.io group.
type MulticlusterV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MulticlusterV1alpha1Client) ClusterClaims(namespace string) ClusterClaimInterface {
	return newClusterClaims(c, namespace)
}

func (c *MulticlusterV1alpha1Client) ClusterSets(namespace string) ClusterSetInterface {
	return newClusterSets(c, namespace)
}

func (c *MulticlusterV1alpha1Client) MemberClusterAnnounces(namespace string) MemberClusterAnnounceInterface {
	return newMemberClusterAnnounces(c, namespace)
}

func (c *MulticlusterV1alpha1Client) ResourceExports(namespace string) ResourceExportInterface {
	return newResourceExports(c, namespace)
}

func (c *MulticlusterV1alpha1Client) ResourceExportFilters(namespace string) ResourceExportFilterInterface {
	return newResourceExportFilters(c, namespace)
}

func (c *MulticlusterV1alpha1Client) ResourceImports(namespace string) ResourceImportInterface {
	return newResourceImports(c, namespace)
}

func (c *MulticlusterV1alpha1Client) ResourceImportFilters(namespace string) ResourceImportFilterInterface {
	return newResourceImportFilters(c, namespace)
}

// NewForConfig creates a new MulticlusterV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MulticlusterV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MulticlusterV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MulticlusterV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MulticlusterV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MulticlusterV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MulticlusterV1alpha1Client {
	return &MulticlusterV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MulticlusterV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
