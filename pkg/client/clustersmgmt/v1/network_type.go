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

// Network represents the values of the 'network' type.
//
// Network configuration of a cluster.
type Network struct {
	podCIDR     *string
	machineCIDR *string
	serviceCIDR *string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *Network) Empty() bool {
	return o == nil || (o.podCIDR == nil &&
		o.machineCIDR == nil &&
		o.serviceCIDR == nil &&
		true)
}

// PodCIDR returns the value of the 'pod_CIDR' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// IP address block from which to assign pod IP addresses, for example `10.128.0.0/14`.
func (o *Network) PodCIDR() string {
	if o != nil && o.podCIDR != nil {
		return *o.podCIDR
	}
	return ""
}

// GetPodCIDR returns the value of the 'pod_CIDR' attribute and
// a flag indicating if the attribute has a value.
//
// IP address block from which to assign pod IP addresses, for example `10.128.0.0/14`.
func (o *Network) GetPodCIDR() (value string, ok bool) {
	ok = o != nil && o.podCIDR != nil
	if ok {
		value = *o.podCIDR
	}
	return
}

// MachineCIDR returns the value of the 'machine_CIDR' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// IP address block from which to assign machine IP addresses, for example `10.0.0.0/16`.
func (o *Network) MachineCIDR() string {
	if o != nil && o.machineCIDR != nil {
		return *o.machineCIDR
	}
	return ""
}

// GetMachineCIDR returns the value of the 'machine_CIDR' attribute and
// a flag indicating if the attribute has a value.
//
// IP address block from which to assign machine IP addresses, for example `10.0.0.0/16`.
func (o *Network) GetMachineCIDR() (value string, ok bool) {
	ok = o != nil && o.machineCIDR != nil
	if ok {
		value = *o.machineCIDR
	}
	return
}

// ServiceCIDR returns the value of the 'service_CIDR' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// IP address block from which to assign service IP addresses, for example `172.30.0.0/16`.
func (o *Network) ServiceCIDR() string {
	if o != nil && o.serviceCIDR != nil {
		return *o.serviceCIDR
	}
	return ""
}

// GetServiceCIDR returns the value of the 'service_CIDR' attribute and
// a flag indicating if the attribute has a value.
//
// IP address block from which to assign service IP addresses, for example `172.30.0.0/16`.
func (o *Network) GetServiceCIDR() (value string, ok bool) {
	ok = o != nil && o.serviceCIDR != nil
	if ok {
		value = *o.serviceCIDR
	}
	return
}

// NetworkList is a list of values of the 'network' type.
type NetworkList struct {
	items []*Network
}

// Len returns the length of the list.
func (l *NetworkList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *NetworkList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *NetworkList) Get(i int) *Network {
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
func (l *NetworkList) Slice() []*Network {
	var slice []*Network
	if l == nil {
		slice = make([]*Network, 0)
	} else {
		slice = make([]*Network, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *NetworkList) Each(f func(item *Network) bool) {
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
func (l *NetworkList) Range(f func(index int, item *Network) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
