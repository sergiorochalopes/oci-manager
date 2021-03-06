/*
Copyright 2018 Oracle and/or its affiliates. All rights reserved.

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
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocicore.oracle.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubnets implements SubnetInterface
type FakeSubnets struct {
	Fake *FakeOcicoreV1alpha1
	ns   string
}

var subnetsResource = schema.GroupVersionResource{Group: "ocicore.oracle.com", Version: "v1alpha1", Resource: "subnets"}

var subnetsKind = schema.GroupVersionKind{Group: "ocicore.oracle.com", Version: "v1alpha1", Kind: "Subnet"}

// Get takes name of the subnet, and returns the corresponding subnet object, and an error if there is any.
func (c *FakeSubnets) Get(name string, options v1.GetOptions) (result *v1alpha1.Subnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subnetsResource, c.ns, name), &v1alpha1.Subnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subnet), err
}

// List takes label and field selectors, and returns the list of Subnets that match those selectors.
func (c *FakeSubnets) List(opts v1.ListOptions) (result *v1alpha1.SubnetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subnetsResource, subnetsKind, c.ns, opts), &v1alpha1.SubnetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubnetList{ListMeta: obj.(*v1alpha1.SubnetList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubnetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subnets.
func (c *FakeSubnets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subnetsResource, c.ns, opts))

}

// Create takes the representation of a subnet and creates it.  Returns the server's representation of the subnet, and an error, if there is any.
func (c *FakeSubnets) Create(subnet *v1alpha1.Subnet) (result *v1alpha1.Subnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subnetsResource, c.ns, subnet), &v1alpha1.Subnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subnet), err
}

// Update takes the representation of a subnet and updates it. Returns the server's representation of the subnet, and an error, if there is any.
func (c *FakeSubnets) Update(subnet *v1alpha1.Subnet) (result *v1alpha1.Subnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subnetsResource, c.ns, subnet), &v1alpha1.Subnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subnet), err
}

// Delete takes name of the subnet and deletes it. Returns an error if one occurs.
func (c *FakeSubnets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(subnetsResource, c.ns, name), &v1alpha1.Subnet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubnets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subnetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubnetList{})
	return err
}

// Patch applies the patch and returns the patched subnet.
func (c *FakeSubnets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Subnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subnetsResource, c.ns, name, data, subresources...), &v1alpha1.Subnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subnet), err
}
