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

// OrganizationKind is the name of the type used to represent objects
// of type 'organization'.
const OrganizationKind = "Organization"

// OrganizationLinkKind is the name of the type used to represent links
// to objects of type 'organization'.
const OrganizationLinkKind = "OrganizationLink"

// OrganizationNilKind is the name of the type used to nil references
// to objects of type 'organization'.
const OrganizationNilKind = "OrganizationNil"

// Organization represents the values of the 'organization' type.
//
//
type Organization struct {
	id   *string
	href *string
	link bool
	name *string
}

// Kind returns the name of the type of the object.
func (o *Organization) Kind() string {
	if o == nil {
		return OrganizationNilKind
	}
	if o.link {
		return OrganizationLinkKind
	}
	return OrganizationKind
}

// ID returns the identifier of the object.
func (o *Organization) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Organization) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Organization) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Organization) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Organization) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *Organization) Empty() bool {
	return o == nil || (o.id == nil &&
		o.name == nil &&
		true)
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Organization) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Organization) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// OrganizationListKind is the name of the type used to represent list of
// objects of type 'organization'.
const OrganizationListKind = "OrganizationList"

// OrganizationListLinkKind is the name of the type used to represent links
// to list of objects of type 'organization'.
const OrganizationListLinkKind = "OrganizationListLink"

// OrganizationNilKind is the name of the type used to nil lists of
// objects of type 'organization'.
const OrganizationListNilKind = "OrganizationListNil"

// OrganizationList is a list of values of the 'organization' type.
type OrganizationList struct {
	href  *string
	link  bool
	items []*Organization
}

// Kind returns the name of the type of the object.
func (l *OrganizationList) Kind() string {
	if l == nil {
		return OrganizationListNilKind
	}
	if l.link {
		return OrganizationListLinkKind
	}
	return OrganizationListKind
}

// Link returns true iif this is a link.
func (l *OrganizationList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *OrganizationList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *OrganizationList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *OrganizationList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *OrganizationList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *OrganizationList) Get(i int) *Organization {
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
func (l *OrganizationList) Slice() []*Organization {
	var slice []*Organization
	if l == nil {
		slice = make([]*Organization, 0)
	} else {
		slice = make([]*Organization, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *OrganizationList) Each(f func(item *Organization) bool) {
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
func (l *OrganizationList) Range(f func(index int, item *Organization) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
