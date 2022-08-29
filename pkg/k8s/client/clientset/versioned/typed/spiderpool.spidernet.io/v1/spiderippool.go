// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v1"
	scheme "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SpiderIPPoolsGetter has a method to return a SpiderIPPoolInterface.
// A group's client should implement this interface.
type SpiderIPPoolsGetter interface {
	SpiderIPPools() SpiderIPPoolInterface
}

// SpiderIPPoolInterface has methods to work with SpiderIPPool resources.
type SpiderIPPoolInterface interface {
	Create(ctx context.Context, spiderIPPool *v1.SpiderIPPool, opts metav1.CreateOptions) (*v1.SpiderIPPool, error)
	Update(ctx context.Context, spiderIPPool *v1.SpiderIPPool, opts metav1.UpdateOptions) (*v1.SpiderIPPool, error)
	UpdateStatus(ctx context.Context, spiderIPPool *v1.SpiderIPPool, opts metav1.UpdateOptions) (*v1.SpiderIPPool, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SpiderIPPool, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SpiderIPPoolList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpiderIPPool, err error)
	SpiderIPPoolExpansion
}

// spiderIPPools implements SpiderIPPoolInterface
type spiderIPPools struct {
	client rest.Interface
}

// newSpiderIPPools returns a SpiderIPPools
func newSpiderIPPools(c *SpiderpoolV1Client) *spiderIPPools {
	return &spiderIPPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the spiderIPPool, and returns the corresponding spiderIPPool object, and an error if there is any.
func (c *spiderIPPools) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SpiderIPPool, err error) {
	result = &v1.SpiderIPPool{}
	err = c.client.Get().
		Resource("spiderippools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpiderIPPools that match those selectors.
func (c *spiderIPPools) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SpiderIPPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SpiderIPPoolList{}
	err = c.client.Get().
		Resource("spiderippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spiderIPPools.
func (c *spiderIPPools) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("spiderippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a spiderIPPool and creates it.  Returns the server's representation of the spiderIPPool, and an error, if there is any.
func (c *spiderIPPools) Create(ctx context.Context, spiderIPPool *v1.SpiderIPPool, opts metav1.CreateOptions) (result *v1.SpiderIPPool, err error) {
	result = &v1.SpiderIPPool{}
	err = c.client.Post().
		Resource("spiderippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spiderIPPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a spiderIPPool and updates it. Returns the server's representation of the spiderIPPool, and an error, if there is any.
func (c *spiderIPPools) Update(ctx context.Context, spiderIPPool *v1.SpiderIPPool, opts metav1.UpdateOptions) (result *v1.SpiderIPPool, err error) {
	result = &v1.SpiderIPPool{}
	err = c.client.Put().
		Resource("spiderippools").
		Name(spiderIPPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spiderIPPool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *spiderIPPools) UpdateStatus(ctx context.Context, spiderIPPool *v1.SpiderIPPool, opts metav1.UpdateOptions) (result *v1.SpiderIPPool, err error) {
	result = &v1.SpiderIPPool{}
	err = c.client.Put().
		Resource("spiderippools").
		Name(spiderIPPool.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spiderIPPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the spiderIPPool and deletes it. Returns an error if one occurs.
func (c *spiderIPPools) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("spiderippools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spiderIPPools) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("spiderippools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched spiderIPPool.
func (c *spiderIPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpiderIPPool, err error) {
	result = &v1.SpiderIPPool{}
	err = c.client.Patch(pt).
		Resource("spiderippools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}