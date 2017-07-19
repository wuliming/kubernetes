/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1beta2

import (
	apps_v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/internalinterfaces"
	v1beta2 "k8s.io/kubernetes/pkg/client/listers/apps/v1beta2"
	time "time"
)

// DeploymentInformer provides access to a shared informer and lister for
// Deployments.
type DeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.DeploymentLister
}

type deploymentInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newDeploymentInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.AppsV1beta2().Deployments(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.AppsV1beta2().Deployments(v1.NamespaceAll).Watch(options)
			},
		},
		&apps_v1beta2.Deployment{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *deploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apps_v1beta2.Deployment{}, newDeploymentInformer)
}

func (f *deploymentInformer) Lister() v1beta2.DeploymentLister {
	return v1beta2.NewDeploymentLister(f.Informer().GetIndexer())
}
