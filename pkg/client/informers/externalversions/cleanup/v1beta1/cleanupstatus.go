// Copyright 2019 Antrea Authors
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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	cleanupv1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/cleanup/v1beta1"
	versioned "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/antrea/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/vmware-tanzu/antrea/pkg/client/listers/cleanup/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CleanupStatusInformer provides access to a shared informer and lister for
// CleanupStatuses.
type CleanupStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.CleanupStatusLister
}

type cleanupStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCleanupStatusInformer constructs a new informer for CleanupStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCleanupStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCleanupStatusInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCleanupStatusInformer constructs a new informer for CleanupStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCleanupStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CleanupV1beta1().CleanupStatuses().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CleanupV1beta1().CleanupStatuses().Watch(options)
			},
		},
		&cleanupv1beta1.CleanupStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *cleanupStatusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCleanupStatusInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cleanupStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cleanupv1beta1.CleanupStatus{}, f.defaultInformer)
}

func (f *cleanupStatusInformer) Lister() v1beta1.CleanupStatusLister {
	return v1beta1.NewCleanupStatusLister(f.Informer().GetIndexer())
}
