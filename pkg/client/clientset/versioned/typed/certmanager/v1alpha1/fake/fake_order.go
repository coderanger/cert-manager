/*
Copyright 2018 The Jetstack cert-manager contributors.

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

package fake

import (
	v1alpha1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOrders implements OrderInterface
type FakeOrders struct {
	Fake *FakeCertmanagerV1alpha1
	ns   string
}

var ordersResource = schema.GroupVersionResource{Group: "certmanager.k8s.io", Version: "v1alpha1", Resource: "orders"}

var ordersKind = schema.GroupVersionKind{Group: "certmanager.k8s.io", Version: "v1alpha1", Kind: "Order"}

// Get takes name of the order, and returns the corresponding order object, and an error if there is any.
func (c *FakeOrders) Get(name string, options v1.GetOptions) (result *v1alpha1.Order, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ordersResource, c.ns, name), &v1alpha1.Order{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Order), err
}

// List takes label and field selectors, and returns the list of Orders that match those selectors.
func (c *FakeOrders) List(opts v1.ListOptions) (result *v1alpha1.OrderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ordersResource, ordersKind, c.ns, opts), &v1alpha1.OrderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrderList{}
	for _, item := range obj.(*v1alpha1.OrderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested orders.
func (c *FakeOrders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ordersResource, c.ns, opts))

}

// Create takes the representation of a order and creates it.  Returns the server's representation of the order, and an error, if there is any.
func (c *FakeOrders) Create(order *v1alpha1.Order) (result *v1alpha1.Order, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ordersResource, c.ns, order), &v1alpha1.Order{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Order), err
}

// Update takes the representation of a order and updates it. Returns the server's representation of the order, and an error, if there is any.
func (c *FakeOrders) Update(order *v1alpha1.Order) (result *v1alpha1.Order, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ordersResource, c.ns, order), &v1alpha1.Order{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Order), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrders) UpdateStatus(order *v1alpha1.Order) (*v1alpha1.Order, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ordersResource, "status", c.ns, order), &v1alpha1.Order{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Order), err
}

// Delete takes name of the order and deletes it. Returns an error if one occurs.
func (c *FakeOrders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ordersResource, c.ns, name), &v1alpha1.Order{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ordersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrderList{})
	return err
}

// Patch applies the patch and returns the patched order.
func (c *FakeOrders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Order, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ordersResource, c.ns, name, data, subresources...), &v1alpha1.Order{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Order), err
}