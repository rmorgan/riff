/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This file was automatically generated by informer-gen

package v1alpha1

import (
	projectriff_io_v1alpha1 "github.com/projectriff/riff/kubernetes-crds/pkg/apis/projectriff.io/v1alpha1"
	versioned "github.com/projectriff/riff/kubernetes-crds/pkg/client/clientset/versioned"
	internalinterfaces "github.com/projectriff/riff/kubernetes-crds/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/projectriff/riff/kubernetes-crds/pkg/client/listers/projectriff/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FunctionInformer provides access to a shared informer and lister for
// Functions.
type FunctionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FunctionLister
}

type functionInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewFunctionInformer constructs a new informer for Function type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFunctionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ProjectriffV1alpha1().Functions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ProjectriffV1alpha1().Functions(namespace).Watch(options)
			},
		},
		&projectriff_io_v1alpha1.Function{},
		resyncPeriod,
		indexers,
	)
}

func defaultFunctionInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFunctionInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *functionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectriff_io_v1alpha1.Function{}, defaultFunctionInformer)
}

func (f *functionInformer) Lister() v1alpha1.FunctionLister {
	return v1alpha1.NewFunctionLister(f.Informer().GetIndexer())
}
