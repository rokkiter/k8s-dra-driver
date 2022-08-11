/*
 * Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/NVIDIA/k8s-dra-driver/pkg/crd/nvidia/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GpuParameterSetLister helps list GpuParameterSets.
// All objects returned here must be treated as read-only.
type GpuParameterSetLister interface {
	// List lists all GpuParameterSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GpuParameterSet, err error)
	// GpuParameterSets returns an object that can list and get GpuParameterSets.
	GpuParameterSets(namespace string) GpuParameterSetNamespaceLister
	GpuParameterSetListerExpansion
}

// gpuParameterSetLister implements the GpuParameterSetLister interface.
type gpuParameterSetLister struct {
	indexer cache.Indexer
}

// NewGpuParameterSetLister returns a new GpuParameterSetLister.
func NewGpuParameterSetLister(indexer cache.Indexer) GpuParameterSetLister {
	return &gpuParameterSetLister{indexer: indexer}
}

// List lists all GpuParameterSets in the indexer.
func (s *gpuParameterSetLister) List(selector labels.Selector) (ret []*v1.GpuParameterSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GpuParameterSet))
	})
	return ret, err
}

// GpuParameterSets returns an object that can list and get GpuParameterSets.
func (s *gpuParameterSetLister) GpuParameterSets(namespace string) GpuParameterSetNamespaceLister {
	return gpuParameterSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GpuParameterSetNamespaceLister helps list and get GpuParameterSets.
// All objects returned here must be treated as read-only.
type GpuParameterSetNamespaceLister interface {
	// List lists all GpuParameterSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GpuParameterSet, err error)
	// Get retrieves the GpuParameterSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GpuParameterSet, error)
	GpuParameterSetNamespaceListerExpansion
}

// gpuParameterSetNamespaceLister implements the GpuParameterSetNamespaceLister
// interface.
type gpuParameterSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GpuParameterSets in the indexer for a given namespace.
func (s gpuParameterSetNamespaceLister) List(selector labels.Selector) (ret []*v1.GpuParameterSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GpuParameterSet))
	})
	return ret, err
}

// Get retrieves the GpuParameterSet from the indexer for a given namespace and name.
func (s gpuParameterSetNamespaceLister) Get(name string) (*v1.GpuParameterSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gpuparameterset"), name)
	}
	return obj.(*v1.GpuParameterSet), nil
}