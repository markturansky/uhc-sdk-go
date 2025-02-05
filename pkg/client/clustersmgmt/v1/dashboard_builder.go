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

// DashboardBuilder contains the data and logic needed to build 'dashboard' objects.
//
// Collection of metrics intended to render a graphical dashboard.
type DashboardBuilder struct {
	id      *string
	href    *string
	link    bool
	name    *string
	metrics []*MetricBuilder
}

// NewDashboard creates a new builder of 'dashboard' objects.
func NewDashboard() *DashboardBuilder {
	return new(DashboardBuilder)
}

// ID sets the identifier of the object.
func (b *DashboardBuilder) ID(value string) *DashboardBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *DashboardBuilder) HREF(value string) *DashboardBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *DashboardBuilder) Link(value bool) *DashboardBuilder {
	b.link = value
	return b
}

// Name sets the value of the 'name' attribute
// to the given value.
//
//
func (b *DashboardBuilder) Name(value string) *DashboardBuilder {
	b.name = &value
	return b
}

// Metrics sets the value of the 'metrics' attribute
// to the given values.
//
//
func (b *DashboardBuilder) Metrics(values ...*MetricBuilder) *DashboardBuilder {
	b.metrics = make([]*MetricBuilder, len(values))
	copy(b.metrics, values)
	return b
}

// Build creates a 'dashboard' object using the configuration stored in the builder.
func (b *DashboardBuilder) Build() (object *Dashboard, err error) {
	object = new(Dashboard)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.name != nil {
		object.name = b.name
	}
	if b.metrics != nil {
		object.metrics = new(MetricList)
		object.metrics.items = make([]*Metric, len(b.metrics))
		for i, item := range b.metrics {
			object.metrics.items[i], err = item.Build()
			if err != nil {
				return
			}
		}
	}
	return
}
