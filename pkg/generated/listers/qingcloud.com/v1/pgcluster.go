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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/radondb/postgres-operator/pkg/apis/radondb.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PgclusterLister helps list Pgclusters.
// All objects returned here must be treated as read-only.
type PgclusterLister interface {
	// List lists all Pgclusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Pgcluster, err error)
	// Pgclusters returns an object that can list and get Pgclusters.
	Pgclusters(namespace string) PgclusterNamespaceLister
	PgclusterListerExpansion
}

// pgclusterLister implements the PgclusterLister interface.
type pgclusterLister struct {
	indexer cache.Indexer
}

// NewPgclusterLister returns a new PgclusterLister.
func NewPgclusterLister(indexer cache.Indexer) PgclusterLister {
	return &pgclusterLister{indexer: indexer}
}

// List lists all Pgclusters in the indexer.
func (s *pgclusterLister) List(selector labels.Selector) (ret []*v1.Pgcluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pgcluster))
	})
	return ret, err
}

// Pgclusters returns an object that can list and get Pgclusters.
func (s *pgclusterLister) Pgclusters(namespace string) PgclusterNamespaceLister {
	return pgclusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PgclusterNamespaceLister helps list and get Pgclusters.
// All objects returned here must be treated as read-only.
type PgclusterNamespaceLister interface {
	// List lists all Pgclusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Pgcluster, err error)
	// Get retrieves the Pgcluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Pgcluster, error)
	PgclusterNamespaceListerExpansion
}

// pgclusterNamespaceLister implements the PgclusterNamespaceLister
// interface.
type pgclusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Pgclusters in the indexer for a given namespace.
func (s pgclusterNamespaceLister) List(selector labels.Selector) (ret []*v1.Pgcluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pgcluster))
	})
	return ret, err
}

// Get retrieves the Pgcluster from the indexer for a given namespace and name.
func (s pgclusterNamespaceLister) Get(name string) (*v1.Pgcluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pgcluster"), name)
	}
	return obj.(*v1.Pgcluster), nil
}
