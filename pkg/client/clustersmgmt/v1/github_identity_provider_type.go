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

// GithubIdentityProvider represents the values of the 'github_identity_provider' type.
//
// Details for `github` identity providers.
type GithubIdentityProvider struct {
	ca       *string
	clientID *string
	hostname *string
	teams    []string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *GithubIdentityProvider) Empty() bool {
	return o == nil || (o.ca == nil &&
		o.clientID == nil &&
		o.hostname == nil &&
		o.teams == nil &&
		true)
}

// CA returns the value of the 'CA' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Optional trusted certificate authority bundle to use when making requests tot he server.
func (o *GithubIdentityProvider) CA() string {
	if o != nil && o.ca != nil {
		return *o.ca
	}
	return ""
}

// GetCA returns the value of the 'CA' attribute and
// a flag indicating if the attribute has a value.
//
// Optional trusted certificate authority bundle to use when making requests tot he server.
func (o *GithubIdentityProvider) GetCA() (value string, ok bool) {
	ok = o != nil && o.ca != nil
	if ok {
		value = *o.ca
	}
	return
}

// ClientID returns the value of the 'client_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Client identifier of a registered _GitHub_ OAuth application.
func (o *GithubIdentityProvider) ClientID() string {
	if o != nil && o.clientID != nil {
		return *o.clientID
	}
	return ""
}

// GetClientID returns the value of the 'client_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Client identifier of a registered _GitHub_ OAuth application.
func (o *GithubIdentityProvider) GetClientID() (value string, ok bool) {
	ok = o != nil && o.clientID != nil
	if ok {
		value = *o.clientID
	}
	return
}

// Hostname returns the value of the 'hostname' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// For _GitHub Enterprise_ you must provide the host name of your instance, such as
// `example.com`. This value must match the _GitHub Enterprise_ host name value in the
// `/setup/settings` file and cannot include a port number.
//
// For plain _GitHub_ omit this parameter.
func (o *GithubIdentityProvider) Hostname() string {
	if o != nil && o.hostname != nil {
		return *o.hostname
	}
	return ""
}

// GetHostname returns the value of the 'hostname' attribute and
// a flag indicating if the attribute has a value.
//
// For _GitHub Enterprise_ you must provide the host name of your instance, such as
// `example.com`. This value must match the _GitHub Enterprise_ host name value in the
// `/setup/settings` file and cannot include a port number.
//
// For plain _GitHub_ omit this parameter.
func (o *GithubIdentityProvider) GetHostname() (value string, ok bool) {
	ok = o != nil && o.hostname != nil
	if ok {
		value = *o.hostname
	}
	return
}

// Teams returns the value of the 'teams' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Optional list of temas. Cannot be used in combination with the organizations field.
func (o *GithubIdentityProvider) Teams() []string {
	if o == nil {
		return nil
	}
	return o.teams
}

// GetTeams returns the value of the 'teams' attribute and
// a flag indicating if the attribute has a value.
//
// Optional list of temas. Cannot be used in combination with the organizations field.
func (o *GithubIdentityProvider) GetTeams() (value []string, ok bool) {
	ok = o != nil && o.teams != nil
	if ok {
		value = o.teams
	}
	return
}

// GithubIdentityProviderList is a list of values of the 'github_identity_provider' type.
type GithubIdentityProviderList struct {
	items []*GithubIdentityProvider
}

// Len returns the length of the list.
func (l *GithubIdentityProviderList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *GithubIdentityProviderList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *GithubIdentityProviderList) Get(i int) *GithubIdentityProvider {
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
func (l *GithubIdentityProviderList) Slice() []*GithubIdentityProvider {
	var slice []*GithubIdentityProvider
	if l == nil {
		slice = make([]*GithubIdentityProvider, 0)
	} else {
		slice = make([]*GithubIdentityProvider, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *GithubIdentityProviderList) Each(f func(item *GithubIdentityProvider) bool) {
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
func (l *GithubIdentityProviderList) Range(f func(index int, item *GithubIdentityProvider) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
