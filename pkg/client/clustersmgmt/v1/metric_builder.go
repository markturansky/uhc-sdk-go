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

// MetricBuilder contains the data and logic needed to build 'metric' objects.
//
// Metric included in a dashboard.
type MetricBuilder struct {
	name   *string
	vector []*SampleBuilder
}

// NewMetric creates a new builder of 'metric' objects.
func NewMetric() *MetricBuilder {
	return new(MetricBuilder)
}

// Name sets the value of the 'name' attribute
// to the given value.
//
//
func (b *MetricBuilder) Name(value string) *MetricBuilder {
	b.name = &value
	return b
}

// Vector sets the value of the 'vector' attribute
// to the given values.
//
//
func (b *MetricBuilder) Vector(values ...*SampleBuilder) *MetricBuilder {
	b.vector = make([]*SampleBuilder, len(values))
	copy(b.vector, values)
	return b
}

// Build creates a 'metric' object using the configuration stored in the builder.
func (b *MetricBuilder) Build() (object *Metric, err error) {
	object = new(Metric)
	if b.name != nil {
		object.name = b.name
	}
	if b.vector != nil {
		object.vector = new(SampleList)
		object.vector.items = make([]*Sample, len(b.vector))
		for i, item := range b.vector {
			object.vector.items[i], err = item.Build()
			if err != nil {
				return
			}
		}
	}
	return
}
