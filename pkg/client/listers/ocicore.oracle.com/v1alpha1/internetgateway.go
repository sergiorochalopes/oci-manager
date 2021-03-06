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
package v1alpha1

import (
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocicore.oracle.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InternetGatewayLister helps list InternetGatewaies.
type InternetGatewayLister interface {
	// List lists all InternetGatewaies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.InternetGateway, err error)
	// InternetGatewaies returns an object that can list and get InternetGatewaies.
	InternetGatewaies(namespace string) InternetGatewayNamespaceLister
	InternetGatewayListerExpansion
}

// internetGatewayLister implements the InternetGatewayLister interface.
type internetGatewayLister struct {
	indexer cache.Indexer
}

// NewInternetGatewayLister returns a new InternetGatewayLister.
func NewInternetGatewayLister(indexer cache.Indexer) InternetGatewayLister {
	return &internetGatewayLister{indexer: indexer}
}

// List lists all InternetGatewaies in the indexer.
func (s *internetGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.InternetGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InternetGateway))
	})
	return ret, err
}

// InternetGatewaies returns an object that can list and get InternetGatewaies.
func (s *internetGatewayLister) InternetGatewaies(namespace string) InternetGatewayNamespaceLister {
	return internetGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InternetGatewayNamespaceLister helps list and get InternetGatewaies.
type InternetGatewayNamespaceLister interface {
	// List lists all InternetGatewaies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.InternetGateway, err error)
	// Get retrieves the InternetGateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.InternetGateway, error)
	InternetGatewayNamespaceListerExpansion
}

// internetGatewayNamespaceLister implements the InternetGatewayNamespaceLister
// interface.
type internetGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InternetGatewaies in the indexer for a given namespace.
func (s internetGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InternetGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InternetGateway))
	})
	return ret, err
}

// Get retrieves the InternetGateway from the indexer for a given namespace and name.
func (s internetGatewayNamespaceLister) Get(name string) (*v1alpha1.InternetGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("internetgateway"), name)
	}
	return obj.(*v1alpha1.InternetGateway), nil
}
