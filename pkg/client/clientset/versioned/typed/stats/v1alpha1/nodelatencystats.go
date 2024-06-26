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
	"time"

	v1alpha1 "antrea.io/antrea/pkg/apis/stats/v1alpha1"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// NodeLatencyStatsGetter has a method to return a NodeLatencyStatsInterface.
// A group's client should implement this interface.
type NodeLatencyStatsGetter interface {
	NodeLatencyStats() NodeLatencyStatsInterface
}

// NodeLatencyStatsInterface has methods to work with NodeLatencyStats resources.
type NodeLatencyStatsInterface interface {
	Create(ctx context.Context, nodeLatencyStats *v1alpha1.NodeLatencyStats, opts v1.CreateOptions) (*v1alpha1.NodeLatencyStats, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NodeLatencyStats, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodeLatencyStatsList, error)
	NodeLatencyStatsExpansion
}

// nodeLatencyStats implements NodeLatencyStatsInterface
type nodeLatencyStats struct {
	client rest.Interface
}

// newNodeLatencyStats returns a NodeLatencyStats
func newNodeLatencyStats(c *StatsV1alpha1Client) *nodeLatencyStats {
	return &nodeLatencyStats{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeLatencyStats, and returns the corresponding nodeLatencyStats object, and an error if there is any.
func (c *nodeLatencyStats) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeLatencyStats, err error) {
	result = &v1alpha1.NodeLatencyStats{}
	err = c.client.Get().
		Resource("nodelatencystats").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeLatencyStats that match those selectors.
func (c *nodeLatencyStats) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeLatencyStatsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodeLatencyStatsList{}
	err = c.client.Get().
		Resource("nodelatencystats").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Create takes the representation of a nodeLatencyStats and creates it.  Returns the server's representation of the nodeLatencyStats, and an error, if there is any.
func (c *nodeLatencyStats) Create(ctx context.Context, nodeLatencyStats *v1alpha1.NodeLatencyStats, opts v1.CreateOptions) (result *v1alpha1.NodeLatencyStats, err error) {
	result = &v1alpha1.NodeLatencyStats{}
	err = c.client.Post().
		Resource("nodelatencystats").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeLatencyStats).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodeLatencyStats and deletes it. Returns an error if one occurs.
func (c *nodeLatencyStats) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodelatencystats").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}
