/*
Copyright 2020 - 2021 Radondb Data Solutions, Inc.
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

package fake

import (
	"context"

	randondbcomv1 "github.com/randondb/postgres-operator/pkg/apis/randondb.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePgpolicies implements PgpolicyInterface
type FakePgpolicies struct {
	Fake *FakeRadondbV1
	ns   string
}

var pgpoliciesResource = schema.GroupVersionResource{Group: "randondb.com", Version: "v1", Resource: "pgpolicies"}

var pgpoliciesKind = schema.GroupVersionKind{Group: "randondb.com", Version: "v1", Kind: "Pgpolicy"}

// Get takes name of the pgpolicy, and returns the corresponding pgpolicy object, and an error if there is any.
func (c *FakePgpolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *randondbcomv1.Pgpolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pgpoliciesResource, c.ns, name), &randondbcomv1.Pgpolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*randondbcomv1.Pgpolicy), err
}

// List takes label and field selectors, and returns the list of Pgpolicies that match those selectors.
func (c *FakePgpolicies) List(ctx context.Context, opts v1.ListOptions) (result *randondbcomv1.PgpolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pgpoliciesResource, pgpoliciesKind, c.ns, opts), &randondbcomv1.PgpolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &randondbcomv1.PgpolicyList{ListMeta: obj.(*randondb.comcomv1.PgpolicyList).ListMeta}
	for _, item := range obj.(*randondbcomv1.PgpolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pgpolicies.
func (c *FakePgpolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pgpoliciesResource, c.ns, opts))

}

// Create takes the representation of a pgpolicy and creates it.  Returns the server's representation of the pgpolicy, and an error, if there is any.
func (c *FakePgpolicies) Create(ctx context.Context, pgpolicy *randondbcomv1.Pgpolicy, opts v1.CreateOptions) (result *randondb.comcomv1.Pgpolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pgpoliciesResource, c.ns, pgpolicy), &randondbcomv1.Pgpolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*randondbcomv1.Pgpolicy), err
}

// Update takes the representation of a pgpolicy and updates it. Returns the server's representation of the pgpolicy, and an error, if there is any.
func (c *FakePgpolicies) Update(ctx context.Context, pgpolicy *randondbcomv1.Pgpolicy, opts v1.UpdateOptions) (result *randondb.comcomv1.Pgpolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pgpoliciesResource, c.ns, pgpolicy), &randondbcomv1.Pgpolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*randondbcomv1.Pgpolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePgpolicies) UpdateStatus(ctx context.Context, pgpolicy *randondbcomv1.Pgpolicy, opts v1.UpdateOptions) (*randondb.comcomv1.Pgpolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pgpoliciesResource, "status", c.ns, pgpolicy), &randondbcomv1.Pgpolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*randondbcomv1.Pgpolicy), err
}

// Delete takes name of the pgpolicy and deletes it. Returns an error if one occurs.
func (c *FakePgpolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pgpoliciesResource, c.ns, name), &randondbcomv1.Pgpolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePgpolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pgpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &randondbcomv1.PgpolicyList{})
	return err
}

// Patch applies the patch and returns the patched pgpolicy.
func (c *FakePgpolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *randondbcomv1.Pgpolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pgpoliciesResource, c.ns, name, pt, data, subresources...), &randondbcomv1.Pgpolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*randondbcomv1.Pgpolicy), err
}
