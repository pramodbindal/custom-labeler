/*
Copyright 2025 Pramod Bindal

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

package v1alpha1

import (
	pramodbindalv1alpha1 "github.com/pramodbindal/my-custom-labeler/pkg/apis/pramodbindal/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// LabelerLister helps list Labelers.
// All objects returned here must be treated as read-only.
type LabelerLister interface {
	// List lists all Labelers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*pramodbindalv1alpha1.Labeler, err error)
	// Labelers returns an object that can list and get Labelers.
	Labelers(namespace string) LabelerNamespaceLister
	LabelerListerExpansion
}

// labelerLister implements the LabelerLister interface.
type labelerLister struct {
	listers.ResourceIndexer[*pramodbindalv1alpha1.Labeler]
}

// NewLabelerLister returns a new LabelerLister.
func NewLabelerLister(indexer cache.Indexer) LabelerLister {
	return &labelerLister{listers.New[*pramodbindalv1alpha1.Labeler](indexer, pramodbindalv1alpha1.Resource("labeler"))}
}

// Labelers returns an object that can list and get Labelers.
func (s *labelerLister) Labelers(namespace string) LabelerNamespaceLister {
	return labelerNamespaceLister{listers.NewNamespaced[*pramodbindalv1alpha1.Labeler](s.ResourceIndexer, namespace)}
}

// LabelerNamespaceLister helps list and get Labelers.
// All objects returned here must be treated as read-only.
type LabelerNamespaceLister interface {
	// List lists all Labelers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*pramodbindalv1alpha1.Labeler, err error)
	// Get retrieves the Labeler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*pramodbindalv1alpha1.Labeler, error)
	LabelerNamespaceListerExpansion
}

// labelerNamespaceLister implements the LabelerNamespaceLister
// interface.
type labelerNamespaceLister struct {
	listers.ResourceIndexer[*pramodbindalv1alpha1.Labeler]
}
