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

// ClusterMetricsBuilder contains the data and logic needed to build 'cluster_metrics' objects.
//
// Cluster metrics received via telemetry.
type ClusterMetricsBuilder struct {
	cpu                *ClusterMetricBuilder
	memory             *ClusterMetricBuilder
	storage            *ClusterMetricBuilder
	computeNodesCPU    *ClusterMetricBuilder
	computeNodesMemory *ClusterMetricBuilder
	nodes              *ClusterNodesBuilder
}

// NewClusterMetrics creates a new builder of 'cluster_metrics' objects.
func NewClusterMetrics() *ClusterMetricsBuilder {
	return new(ClusterMetricsBuilder)
}

// CPU sets the value of the 'CPU' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterMetricsBuilder) CPU(value *ClusterMetricBuilder) *ClusterMetricsBuilder {
	b.cpu = value
	return b
}

// Memory sets the value of the 'memory' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterMetricsBuilder) Memory(value *ClusterMetricBuilder) *ClusterMetricsBuilder {
	b.memory = value
	return b
}

// Storage sets the value of the 'storage' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterMetricsBuilder) Storage(value *ClusterMetricBuilder) *ClusterMetricsBuilder {
	b.storage = value
	return b
}

// ComputeNodesCPU sets the value of the 'compute_nodes_CPU' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterMetricsBuilder) ComputeNodesCPU(value *ClusterMetricBuilder) *ClusterMetricsBuilder {
	b.computeNodesCPU = value
	return b
}

// ComputeNodesMemory sets the value of the 'compute_nodes_memory' attribute
// to the given value.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
func (b *ClusterMetricsBuilder) ComputeNodesMemory(value *ClusterMetricBuilder) *ClusterMetricsBuilder {
	b.computeNodesMemory = value
	return b
}

// Nodes sets the value of the 'nodes' attribute
// to the given value.
//
// Counts of different classes of nodes inside a cluster.
func (b *ClusterMetricsBuilder) Nodes(value *ClusterNodesBuilder) *ClusterMetricsBuilder {
	b.nodes = value
	return b
}

// Build creates a 'cluster_metrics' object using the configuration stored in the builder.
func (b *ClusterMetricsBuilder) Build() (object *ClusterMetrics, err error) {
	object = new(ClusterMetrics)
	if b.cpu != nil {
		object.cpu, err = b.cpu.Build()
		if err != nil {
			return
		}
	}
	if b.memory != nil {
		object.memory, err = b.memory.Build()
		if err != nil {
			return
		}
	}
	if b.storage != nil {
		object.storage, err = b.storage.Build()
		if err != nil {
			return
		}
	}
	if b.computeNodesCPU != nil {
		object.computeNodesCPU, err = b.computeNodesCPU.Build()
		if err != nil {
			return
		}
	}
	if b.computeNodesMemory != nil {
		object.computeNodesMemory, err = b.computeNodesMemory.Build()
		if err != nil {
			return
		}
	}
	if b.nodes != nil {
		object.nodes, err = b.nodes.Build()
		if err != nil {
			return
		}
	}
	return
}
