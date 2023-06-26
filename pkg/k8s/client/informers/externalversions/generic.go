// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=spiderpool.spidernet.io, Version=v2beta1
	case v2beta1.SchemeGroupVersion.WithResource("spidercoordinators"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Spiderpool().V2beta1().SpiderCoordinators().Informer()}, nil
	case v2beta1.SchemeGroupVersion.WithResource("spiderippools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Spiderpool().V2beta1().SpiderIPPools().Informer()}, nil
	case v2beta1.SchemeGroupVersion.WithResource("spidermultusconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Spiderpool().V2beta1().SpiderMultusConfigs().Informer()}, nil
	case v2beta1.SchemeGroupVersion.WithResource("spidersubnets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Spiderpool().V2beta1().SpiderSubnets().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
