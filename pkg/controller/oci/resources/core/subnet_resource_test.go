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

package core

import (
	"testing"

	"github.com/oracle/oci-manager/pkg/client/clientset/versioned"
	fakeoci "github.com/oracle/oci-manager/pkg/controller/oci/resources/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocicore.oracle.com/v1alpha1"
	fakeclient "github.com/oracle/oci-manager/pkg/client/clientset/versioned/fake"
	resourcescommon "github.com/oracle/oci-manager/pkg/controller/oci/resources/common"
)

var (
	subnettest1 = corev1alpha1.Subnet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "subnet.test1",
			Namespace: fakeNs,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "ocicore.oracle.com/v1alpha1",
			Kind:       corev1alpha1.SubnetKind,
		},
		Spec: corev1alpha1.SubnetSpec{
			DisplayName:        "bla",
			CompartmentRef:     "compartment.test1",
			VcnRef:             "vcn.test1",
			AvailabilityDomain: "ad.test1",
			CidrBlock:          "10.0.0.0/24",
			DNSLabel:           "test1",
			RouteTableRef:      "routeTable.test1",
		},
	}
)

func TestSubnetResourceBasic(t *testing.T) {

	t.Log("Testing subnet_resource")
	clientset := fakeclient.NewSimpleClientset()
	vcnClient := fakeoci.NewVcnClient()

	subnetAdapter := SubnetAdapter{}
	subnetAdapter.clientset = clientset
	subnetAdapter.vcnClient = vcnClient

	comp, err := clientset.OciidentityV1alpha1().Compartments(fakeNs).Create(&compartment)
	if err != nil {
		t.Errorf("Got compartment error %v", err)
	}
	t.Logf("Created compartment object %v", comp)

	vcn, err := clientset.OcicoreV1alpha1().Vcns(fakeNs).Create(&vcntest1)
	if err != nil {
		t.Errorf("Got vcn error %v", err)
	}
	t.Logf("Created vcn object %v", vcn)

	rt, err := clientset.OcicoreV1alpha1().RouteTables(fakeNs).Create(&routeTabletest1)
	if err != nil {
		t.Errorf("Got rt error %v", err)
	}
	t.Logf("Created rt object %v", rt)

	newSubnet, err := subnetAdapter.CreateObject(&subnettest1)
	if err != nil {
		t.Errorf("Got subnet error %v", err)
	}
	t.Logf("Created Subnet object %v", newSubnet)

	if subnetAdapter.IsExpectedType(newSubnet) {
		t.Logf("Checked Subnet type")
	}

	subnetWithResource, err := subnetAdapter.Create(newSubnet)

	if err != nil {
		t.Errorf("Got error %v", err)
	}
	t.Logf("Created Subnet resource %v", subnetWithResource)

	subnetWithResource, err = subnetAdapter.Get(newSubnet)
	subnetWithResource, err = subnetAdapter.Update(newSubnet)
	subnetWithResource, err = subnetAdapter.Delete(newSubnet)
	eq := subnetAdapter.Equivalent(newSubnet, newSubnet)
	if !eq {
		t.Errorf("should be equal")
	}
	_, err = subnetAdapter.DependsOnRefs(newSubnet)
}

func NewFakeSubnetAdapter(clientset versioned.Interface) resourcescommon.ResourceTypeAdapter {
	vcnClient := fakeoci.NewVcnClient()
	subnetAdapter := SubnetAdapter{}
	subnetAdapter.clientset = clientset
	subnetAdapter.vcnClient = vcnClient
	return &subnetAdapter
}
