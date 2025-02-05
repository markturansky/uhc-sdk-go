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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

// ClusterAuthorizationResponse represents the values of the 'cluster_authorization_response' type.
//
//
type ClusterAuthorizationResponse struct {
	allowed         *bool
	excessResources *ReservedResourceList
	subscription    *Subscription
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterAuthorizationResponse) Empty() bool {
	return o == nil || (o.allowed == nil &&
		o.excessResources.Empty() &&
		o.subscription == nil &&
		true)
}

// Allowed returns the value of the 'allowed' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ClusterAuthorizationResponse) Allowed() bool {
	if o != nil && o.allowed != nil {
		return *o.allowed
	}
	return false
}

// GetAllowed returns the value of the 'allowed' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ClusterAuthorizationResponse) GetAllowed() (value bool, ok bool) {
	ok = o != nil && o.allowed != nil
	if ok {
		value = *o.allowed
	}
	return
}

// ExcessResources returns the value of the 'excess_resources' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ClusterAuthorizationResponse) ExcessResources() *ReservedResourceList {
	if o == nil {
		return nil
	}
	return o.excessResources
}

// GetExcessResources returns the value of the 'excess_resources' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ClusterAuthorizationResponse) GetExcessResources() (value *ReservedResourceList, ok bool) {
	ok = o != nil && o.excessResources != nil
	if ok {
		value = o.excessResources
	}
	return
}

// Subscription returns the value of the 'subscription' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ClusterAuthorizationResponse) Subscription() *Subscription {
	if o == nil {
		return nil
	}
	return o.subscription
}

// GetSubscription returns the value of the 'subscription' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ClusterAuthorizationResponse) GetSubscription() (value *Subscription, ok bool) {
	ok = o != nil && o.subscription != nil
	if ok {
		value = o.subscription
	}
	return
}

// ClusterAuthorizationResponseList is a list of values of the 'cluster_authorization_response' type.
type ClusterAuthorizationResponseList struct {
	items []*ClusterAuthorizationResponse
}

// Len returns the length of the list.
func (l *ClusterAuthorizationResponseList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterAuthorizationResponseList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterAuthorizationResponseList) Get(i int) *ClusterAuthorizationResponse {
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
func (l *ClusterAuthorizationResponseList) Slice() []*ClusterAuthorizationResponse {
	var slice []*ClusterAuthorizationResponse
	if l == nil {
		slice = make([]*ClusterAuthorizationResponse, 0)
	} else {
		slice = make([]*ClusterAuthorizationResponse, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterAuthorizationResponseList) Each(f func(item *ClusterAuthorizationResponse) bool) {
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
func (l *ClusterAuthorizationResponseList) Range(f func(index int, item *ClusterAuthorizationResponse) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
