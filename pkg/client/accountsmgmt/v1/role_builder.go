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

// RoleBuilder contains the data and logic needed to build 'role' objects.
//
//
type RoleBuilder struct {
	id          *string
	href        *string
	link        bool
	name        *string
	permissions []*PermissionBuilder
}

// NewRole creates a new builder of 'role' objects.
func NewRole() *RoleBuilder {
	return new(RoleBuilder)
}

// ID sets the identifier of the object.
func (b *RoleBuilder) ID(value string) *RoleBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *RoleBuilder) HREF(value string) *RoleBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *RoleBuilder) Link(value bool) *RoleBuilder {
	b.link = value
	return b
}

// Name sets the value of the 'name' attribute
// to the given value.
//
//
func (b *RoleBuilder) Name(value string) *RoleBuilder {
	b.name = &value
	return b
}

// Permissions sets the value of the 'permissions' attribute
// to the given values.
//
//
func (b *RoleBuilder) Permissions(values ...*PermissionBuilder) *RoleBuilder {
	b.permissions = make([]*PermissionBuilder, len(values))
	copy(b.permissions, values)
	return b
}

// Build creates a 'role' object using the configuration stored in the builder.
func (b *RoleBuilder) Build() (object *Role, err error) {
	object = new(Role)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.name != nil {
		object.name = b.name
	}
	if b.permissions != nil {
		object.permissions = new(PermissionList)
		object.permissions.items = make([]*Permission, len(b.permissions))
		for i, item := range b.permissions {
			object.permissions.items[i], err = item.Build()
			if err != nil {
				return
			}
		}
	}
	return
}
