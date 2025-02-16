/*
Copyright 2021 The SuperEdge authors.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/superedge/superedge/pkg/site-manager/apis/site/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeUnitLister helps list NodeUnits.
// All objects returned here must be treated as read-only.
type NodeUnitLister interface {
	// List lists all NodeUnits in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NodeUnit, err error)
	// Get retrieves the NodeUnit from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NodeUnit, error)
	NodeUnitListerExpansion
}

// nodeUnitLister implements the NodeUnitLister interface.
type nodeUnitLister struct {
	indexer cache.Indexer
}

// NewNodeUnitLister returns a new NodeUnitLister.
func NewNodeUnitLister(indexer cache.Indexer) NodeUnitLister {
	return &nodeUnitLister{indexer: indexer}
}

// List lists all NodeUnits in the indexer.
func (s *nodeUnitLister) List(selector labels.Selector) (ret []*v1.NodeUnit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeUnit))
	})
	return ret, err
}

// Get retrieves the NodeUnit from the index for a given name.
func (s *nodeUnitLister) Get(name string) (*v1.NodeUnit, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("nodeunit"), name)
	}
	return obj.(*v1.NodeUnit), nil
}
