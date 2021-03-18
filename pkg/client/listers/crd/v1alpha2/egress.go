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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/vmware-tanzu/antrea/pkg/apis/crd/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EgressLister helps list Egresses.
// All objects returned here must be treated as read-only.
type EgressLister interface {
	// List lists all Egresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.Egress, err error)
	// Get retrieves the Egress from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.Egress, error)
	EgressListerExpansion
}

// egressLister implements the EgressLister interface.
type egressLister struct {
	indexer cache.Indexer
}

// NewEgressLister returns a new EgressLister.
func NewEgressLister(indexer cache.Indexer) EgressLister {
	return &egressLister{indexer: indexer}
}

// List lists all Egresses in the indexer.
func (s *egressLister) List(selector labels.Selector) (ret []*v1alpha2.Egress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Egress))
	})
	return ret, err
}

// Get retrieves the Egress from the index for a given name.
func (s *egressLister) Get(name string) (*v1alpha2.Egress, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("egress"), name)
	}
	return obj.(*v1alpha2.Egress), nil
}
