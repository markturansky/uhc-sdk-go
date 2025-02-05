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

// ResourceQuotaKind is the name of the type used to represent objects
// of type 'resource_quota'.
const ResourceQuotaKind = "ResourceQuota"

// ResourceQuotaLinkKind is the name of the type used to represent links
// to objects of type 'resource_quota'.
const ResourceQuotaLinkKind = "ResourceQuotaLink"

// ResourceQuotaNilKind is the name of the type used to nil references
// to objects of type 'resource_quota'.
const ResourceQuotaNilKind = "ResourceQuotaNil"

// ResourceQuota represents the values of the 'resource_quota' type.
//
//
type ResourceQuota struct {
	id                   *string
	href                 *string
	link                 bool
	organizationID       *string
	sku                  *string
	resourceName         *string
	resourceType         *string
	byoc                 *bool
	availabilityZoneType *string
	allowed              *int
	reserved             *int
}

// Kind returns the name of the type of the object.
func (o *ResourceQuota) Kind() string {
	if o == nil {
		return ResourceQuotaNilKind
	}
	if o.link {
		return ResourceQuotaLinkKind
	}
	return ResourceQuotaKind
}

// ID returns the identifier of the object.
func (o *ResourceQuota) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *ResourceQuota) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *ResourceQuota) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *ResourceQuota) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *ResourceQuota) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ResourceQuota) Empty() bool {
	return o == nil || (o.id == nil &&
		o.organizationID == nil &&
		o.sku == nil &&
		o.resourceName == nil &&
		o.resourceType == nil &&
		o.byoc == nil &&
		o.availabilityZoneType == nil &&
		o.allowed == nil &&
		o.reserved == nil &&
		true)
}

// OrganizationID returns the value of the 'organization_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) OrganizationID() string {
	if o != nil && o.organizationID != nil {
		return *o.organizationID
	}
	return ""
}

// GetOrganizationID returns the value of the 'organization_ID' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetOrganizationID() (value string, ok bool) {
	ok = o != nil && o.organizationID != nil
	if ok {
		value = *o.organizationID
	}
	return
}

// SKU returns the value of the 'SKU' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) SKU() string {
	if o != nil && o.sku != nil {
		return *o.sku
	}
	return ""
}

// GetSKU returns the value of the 'SKU' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetSKU() (value string, ok bool) {
	ok = o != nil && o.sku != nil
	if ok {
		value = *o.sku
	}
	return
}

// ResourceName returns the value of the 'resource_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) ResourceName() string {
	if o != nil && o.resourceName != nil {
		return *o.resourceName
	}
	return ""
}

// GetResourceName returns the value of the 'resource_name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetResourceName() (value string, ok bool) {
	ok = o != nil && o.resourceName != nil
	if ok {
		value = *o.resourceName
	}
	return
}

// ResourceType returns the value of the 'resource_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) ResourceType() string {
	if o != nil && o.resourceType != nil {
		return *o.resourceType
	}
	return ""
}

// GetResourceType returns the value of the 'resource_type' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetResourceType() (value string, ok bool) {
	ok = o != nil && o.resourceType != nil
	if ok {
		value = *o.resourceType
	}
	return
}

// BYOC returns the value of the 'BYOC' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) BYOC() bool {
	if o != nil && o.byoc != nil {
		return *o.byoc
	}
	return false
}

// GetBYOC returns the value of the 'BYOC' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetBYOC() (value bool, ok bool) {
	ok = o != nil && o.byoc != nil
	if ok {
		value = *o.byoc
	}
	return
}

// AvailabilityZoneType returns the value of the 'availability_zone_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) AvailabilityZoneType() string {
	if o != nil && o.availabilityZoneType != nil {
		return *o.availabilityZoneType
	}
	return ""
}

// GetAvailabilityZoneType returns the value of the 'availability_zone_type' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetAvailabilityZoneType() (value string, ok bool) {
	ok = o != nil && o.availabilityZoneType != nil
	if ok {
		value = *o.availabilityZoneType
	}
	return
}

// Allowed returns the value of the 'allowed' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) Allowed() int {
	if o != nil && o.allowed != nil {
		return *o.allowed
	}
	return 0
}

// GetAllowed returns the value of the 'allowed' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetAllowed() (value int, ok bool) {
	ok = o != nil && o.allowed != nil
	if ok {
		value = *o.allowed
	}
	return
}

// Reserved returns the value of the 'reserved' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ResourceQuota) Reserved() int {
	if o != nil && o.reserved != nil {
		return *o.reserved
	}
	return 0
}

// GetReserved returns the value of the 'reserved' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ResourceQuota) GetReserved() (value int, ok bool) {
	ok = o != nil && o.reserved != nil
	if ok {
		value = *o.reserved
	}
	return
}

// ResourceQuotaListKind is the name of the type used to represent list of
// objects of type 'resource_quota'.
const ResourceQuotaListKind = "ResourceQuotaList"

// ResourceQuotaListLinkKind is the name of the type used to represent links
// to list of objects of type 'resource_quota'.
const ResourceQuotaListLinkKind = "ResourceQuotaListLink"

// ResourceQuotaNilKind is the name of the type used to nil lists of
// objects of type 'resource_quota'.
const ResourceQuotaListNilKind = "ResourceQuotaListNil"

// ResourceQuotaList is a list of values of the 'resource_quota' type.
type ResourceQuotaList struct {
	href  *string
	link  bool
	items []*ResourceQuota
}

// Kind returns the name of the type of the object.
func (l *ResourceQuotaList) Kind() string {
	if l == nil {
		return ResourceQuotaListNilKind
	}
	if l.link {
		return ResourceQuotaListLinkKind
	}
	return ResourceQuotaListKind
}

// Link returns true iif this is a link.
func (l *ResourceQuotaList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ResourceQuotaList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ResourceQuotaList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *ResourceQuotaList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ResourceQuotaList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ResourceQuotaList) Get(i int) *ResourceQuota {
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
func (l *ResourceQuotaList) Slice() []*ResourceQuota {
	var slice []*ResourceQuota
	if l == nil {
		slice = make([]*ResourceQuota, 0)
	} else {
		slice = make([]*ResourceQuota, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ResourceQuotaList) Each(f func(item *ResourceQuota) bool) {
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
func (l *ResourceQuotaList) Range(f func(index int, item *ResourceQuota) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
