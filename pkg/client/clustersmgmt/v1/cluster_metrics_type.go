/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// ClusterMetrics represents the values of the 'cluster_metrics' type.
//
// Cluster metrics received via telemetry.
type ClusterMetrics struct {
	cpu                *ClusterMetric
	memory             *ClusterMetric
	storage            *ClusterMetric
	computeNodesCPU    *ClusterMetric
	computeNodesMemory *ClusterMetric
	nodes              *ClusterNodes
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterMetrics) Empty() bool {
	return o == nil || (o.cpu == nil &&
		o.memory == nil &&
		o.storage == nil &&
		o.computeNodesCPU == nil &&
		o.computeNodesMemory == nil &&
		o.nodes == nil &&
		true)
}

// CPU returns the value of the 'CPU' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The amount of CPU provisioned and used in the cluster.
func (o *ClusterMetrics) CPU() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.cpu
}

// GetCPU returns the value of the 'CPU' attribute and
// a flag indicating if the attribute has a value.
//
// The amount of CPU provisioned and used in the cluster.
func (o *ClusterMetrics) GetCPU() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.cpu != nil
	if ok {
		value = o.cpu
	}
	return
}

// Memory returns the value of the 'memory' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The amount of memory provisioned and used in the cluster.
func (o *ClusterMetrics) Memory() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.memory
}

// GetMemory returns the value of the 'memory' attribute and
// a flag indicating if the attribute has a value.
//
// The amount of memory provisioned and used in the cluster.
func (o *ClusterMetrics) GetMemory() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.memory != nil
	if ok {
		value = o.memory
	}
	return
}

// Storage returns the value of the 'storage' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The amount of storage provisioned and used in the cluster.
//
// WARNING: This isn't currently populated.
func (o *ClusterMetrics) Storage() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.storage
}

// GetStorage returns the value of the 'storage' attribute and
// a flag indicating if the attribute has a value.
//
// The amount of storage provisioned and used in the cluster.
//
// WARNING: This isn't currently populated.
func (o *ClusterMetrics) GetStorage() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.storage != nil
	if ok {
		value = o.storage
	}
	return
}

// ComputeNodesCPU returns the value of the 'compute_nodes_CPU' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The amount of CPU provisioned and used in the cluster by compute nodes.
func (o *ClusterMetrics) ComputeNodesCPU() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.computeNodesCPU
}

// GetComputeNodesCPU returns the value of the 'compute_nodes_CPU' attribute and
// a flag indicating if the attribute has a value.
//
// The amount of CPU provisioned and used in the cluster by compute nodes.
func (o *ClusterMetrics) GetComputeNodesCPU() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.computeNodesCPU != nil
	if ok {
		value = o.computeNodesCPU
	}
	return
}

// ComputeNodesMemory returns the value of the 'compute_nodes_memory' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The amount of memory provisioned and used in the cluster by compute nodes.
func (o *ClusterMetrics) ComputeNodesMemory() *ClusterMetric {
	if o == nil {
		return nil
	}
	return o.computeNodesMemory
}

// GetComputeNodesMemory returns the value of the 'compute_nodes_memory' attribute and
// a flag indicating if the attribute has a value.
//
// The amount of memory provisioned and used in the cluster by compute nodes.
func (o *ClusterMetrics) GetComputeNodesMemory() (value *ClusterMetric, ok bool) {
	ok = o != nil && o.computeNodesMemory != nil
	if ok {
		value = o.computeNodesMemory
	}
	return
}

// Nodes returns the value of the 'nodes' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The number of nodes provisioned for the cluster.
func (o *ClusterMetrics) Nodes() *ClusterNodes {
	if o == nil {
		return nil
	}
	return o.nodes
}

// GetNodes returns the value of the 'nodes' attribute and
// a flag indicating if the attribute has a value.
//
// The number of nodes provisioned for the cluster.
func (o *ClusterMetrics) GetNodes() (value *ClusterNodes, ok bool) {
	ok = o != nil && o.nodes != nil
	if ok {
		value = o.nodes
	}
	return
}

// ClusterMetricsList is a list of values of the 'cluster_metrics' type.
type ClusterMetricsList struct {
	items []*ClusterMetrics
}

// Len returns the length of the list.
func (l *ClusterMetricsList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterMetricsList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterMetricsList) Get(i int) *ClusterMetrics {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterMetricsList) Slice() []*ClusterMetrics {
	var slice []*ClusterMetrics
	if l == nil {
		slice = make([]*ClusterMetrics, 0)
	} else {
		slice = make([]*ClusterMetrics, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterMetricsList) Each(f func(item *ClusterMetrics) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterMetricsList) Range(f func(index int, item *ClusterMetrics) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
