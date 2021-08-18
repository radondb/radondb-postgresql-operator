/*
Copyright 2020 - 2021 Crunchy Data Solutions, Inc.
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

package v1

import (
	"context"
	"time"

	v1 "github.com/radondb/postgres-operator/pkg/apis/RadonDB.com/v1"
	scheme "github.com/radondb/postgres-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PgtasksGetter has a method to return a PgtaskInterface.
// A group's client should implement this interface.
type PgtasksGetter interface {
	Pgtasks(namespace string) PgtaskInterface
}

// PgtaskInterface has methods to work with Pgtask resources.
type PgtaskInterface interface {
	Create(ctx context.Context, pgtask *v1.Pgtask, opts metav1.CreateOptions) (*v1.Pgtask, error)
	Update(ctx context.Context, pgtask *v1.Pgtask, opts metav1.UpdateOptions) (*v1.Pgtask, error)
	UpdateStatus(ctx context.Context, pgtask *v1.Pgtask, opts metav1.UpdateOptions) (*v1.Pgtask, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Pgtask, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PgtaskList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Pgtask, err error)
	PgtaskExpansion
}

// pgtasks implements PgtaskInterface
type pgtasks struct {
	client rest.Interface
	ns     string
}

// newPgtasks returns a Pgtasks
func newPgtasks(c *RadonDBV1Client, namespace string) *pgtasks {
	return &pgtasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pgtask, and returns the corresponding pgtask object, and an error if there is any.
func (c *pgtasks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Pgtask, err error) {
	result = &v1.Pgtask{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pgtasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Pgtasks that match those selectors.
func (c *pgtasks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PgtaskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PgtaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pgtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pgtasks.
func (c *pgtasks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pgtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a pgtask and creates it.  Returns the server's representation of the pgtask, and an error, if there is any.
func (c *pgtasks) Create(ctx context.Context, pgtask *v1.Pgtask, opts metav1.CreateOptions) (result *v1.Pgtask, err error) {
	result = &v1.Pgtask{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pgtasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pgtask).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a pgtask and updates it. Returns the server's representation of the pgtask, and an error, if there is any.
func (c *pgtasks) Update(ctx context.Context, pgtask *v1.Pgtask, opts metav1.UpdateOptions) (result *v1.Pgtask, err error) {
	result = &v1.Pgtask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pgtasks").
		Name(pgtask.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pgtask).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *pgtasks) UpdateStatus(ctx context.Context, pgtask *v1.Pgtask, opts metav1.UpdateOptions) (result *v1.Pgtask, err error) {
	result = &v1.Pgtask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pgtasks").
		Name(pgtask.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pgtask).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the pgtask and deletes it. Returns an error if one occurs.
func (c *pgtasks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pgtasks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pgtasks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pgtasks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched pgtask.
func (c *pgtasks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Pgtask, err error) {
	result = &v1.Pgtask{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pgtasks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
