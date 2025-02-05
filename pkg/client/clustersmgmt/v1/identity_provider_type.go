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

// IdentityProviderKind is the name of the type used to represent objects
// of type 'identity_provider'.
const IdentityProviderKind = "IdentityProvider"

// IdentityProviderLinkKind is the name of the type used to represent links
// to objects of type 'identity_provider'.
const IdentityProviderLinkKind = "IdentityProviderLink"

// IdentityProviderNilKind is the name of the type used to nil references
// to objects of type 'identity_provider'.
const IdentityProviderNilKind = "IdentityProviderNil"

// IdentityProvider represents the values of the 'identity_provider' type.
//
// Representation of an identity provider.
type IdentityProvider struct {
	id            *string
	href          *string
	link          bool
	type_         *IdentityProviderType
	name          *string
	challenge     *bool
	login         *bool
	mappingMethod *IdentityProviderMappingMethod
	github        *GithubIdentityProvider
	gitlab        *GitlabIdentityProvider
	google        *GoogleIdentityProvider
	ldap          *LdapidentityProvider
	openID        *OpenIdidentityProvider
}

// Kind returns the name of the type of the object.
func (o *IdentityProvider) Kind() string {
	if o == nil {
		return IdentityProviderNilKind
	}
	if o.link {
		return IdentityProviderLinkKind
	}
	return IdentityProviderKind
}

// ID returns the identifier of the object.
func (o *IdentityProvider) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *IdentityProvider) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *IdentityProvider) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *IdentityProvider) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *IdentityProvider) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *IdentityProvider) Empty() bool {
	return o == nil || (o.id == nil &&
		o.type_ == nil &&
		o.name == nil &&
		o.challenge == nil &&
		o.login == nil &&
		o.mappingMethod == nil &&
		o.github == nil &&
		o.gitlab == nil &&
		o.google == nil &&
		o.ldap == nil &&
		o.openID == nil &&
		true)
}

// Type returns the value of the 'type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Type of identity provider. The rest of the attributes will be populated according to this
// value. For example, if the type is `github` then only the `github` attribute will be
// populated.
func (o *IdentityProvider) Type() IdentityProviderType {
	if o != nil && o.type_ != nil {
		return *o.type_
	}
	return IdentityProviderType("")
}

// GetType returns the value of the 'type' attribute and
// a flag indicating if the attribute has a value.
//
// Type of identity provider. The rest of the attributes will be populated according to this
// value. For example, if the type is `github` then only the `github` attribute will be
// populated.
func (o *IdentityProvider) GetType() (value IdentityProviderType, ok bool) {
	ok = o != nil && o.type_ != nil
	if ok {
		value = *o.type_
	}
	return
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The name of the identity provider.
func (o *IdentityProvider) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// The name of the identity provider.
func (o *IdentityProvider) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// Challenge returns the value of the 'challenge' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// When `true` unauthenticated token requests from non-web clients (like the CLI) are sent a
// `WWW-Authenticate` challenge header for this provider.
func (o *IdentityProvider) Challenge() bool {
	if o != nil && o.challenge != nil {
		return *o.challenge
	}
	return false
}

// GetChallenge returns the value of the 'challenge' attribute and
// a flag indicating if the attribute has a value.
//
// When `true` unauthenticated token requests from non-web clients (like the CLI) are sent a
// `WWW-Authenticate` challenge header for this provider.
func (o *IdentityProvider) GetChallenge() (value bool, ok bool) {
	ok = o != nil && o.challenge != nil
	if ok {
		value = *o.challenge
	}
	return
}

// Login returns the value of the 'login' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// When `true` unauthenticated token requests from web clients (like the web console) are
// redirected to the authorize URL to log in.
func (o *IdentityProvider) Login() bool {
	if o != nil && o.login != nil {
		return *o.login
	}
	return false
}

// GetLogin returns the value of the 'login' attribute and
// a flag indicating if the attribute has a value.
//
// When `true` unauthenticated token requests from web clients (like the web console) are
// redirected to the authorize URL to log in.
func (o *IdentityProvider) GetLogin() (value bool, ok bool) {
	ok = o != nil && o.login != nil
	if ok {
		value = *o.login
	}
	return
}

// MappingMethod returns the value of the 'mapping_method' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Controls how mappings are established between this provider's identities and user
// objects.
func (o *IdentityProvider) MappingMethod() IdentityProviderMappingMethod {
	if o != nil && o.mappingMethod != nil {
		return *o.mappingMethod
	}
	return IdentityProviderMappingMethod("")
}

// GetMappingMethod returns the value of the 'mapping_method' attribute and
// a flag indicating if the attribute has a value.
//
// Controls how mappings are established between this provider's identities and user
// objects.
func (o *IdentityProvider) GetMappingMethod() (value IdentityProviderMappingMethod, ok bool) {
	ok = o != nil && o.mappingMethod != nil
	if ok {
		value = *o.mappingMethod
	}
	return
}

// Github returns the value of the 'github' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Details for `github` identity providers.
func (o *IdentityProvider) Github() *GithubIdentityProvider {
	if o == nil {
		return nil
	}
	return o.github
}

// GetGithub returns the value of the 'github' attribute and
// a flag indicating if the attribute has a value.
//
// Details for `github` identity providers.
func (o *IdentityProvider) GetGithub() (value *GithubIdentityProvider, ok bool) {
	ok = o != nil && o.github != nil
	if ok {
		value = o.github
	}
	return
}

// Gitlab returns the value of the 'gitlab' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Details for `gitlab` identity providers.
func (o *IdentityProvider) Gitlab() *GitlabIdentityProvider {
	if o == nil {
		return nil
	}
	return o.gitlab
}

// GetGitlab returns the value of the 'gitlab' attribute and
// a flag indicating if the attribute has a value.
//
// Details for `gitlab` identity providers.
func (o *IdentityProvider) GetGitlab() (value *GitlabIdentityProvider, ok bool) {
	ok = o != nil && o.gitlab != nil
	if ok {
		value = o.gitlab
	}
	return
}

// Google returns the value of the 'google' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Details for `google` identity providers.
func (o *IdentityProvider) Google() *GoogleIdentityProvider {
	if o == nil {
		return nil
	}
	return o.google
}

// GetGoogle returns the value of the 'google' attribute and
// a flag indicating if the attribute has a value.
//
// Details for `google` identity providers.
func (o *IdentityProvider) GetGoogle() (value *GoogleIdentityProvider, ok bool) {
	ok = o != nil && o.google != nil
	if ok {
		value = o.google
	}
	return
}

// LDAP returns the value of the 'LDAP' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Details for `ldap` identity providers.
func (o *IdentityProvider) LDAP() *LdapidentityProvider {
	if o == nil {
		return nil
	}
	return o.ldap
}

// GetLDAP returns the value of the 'LDAP' attribute and
// a flag indicating if the attribute has a value.
//
// Details for `ldap` identity providers.
func (o *IdentityProvider) GetLDAP() (value *LdapidentityProvider, ok bool) {
	ok = o != nil && o.ldap != nil
	if ok {
		value = o.ldap
	}
	return
}

// OpenID returns the value of the 'open_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Details for `openid` identity providers.
func (o *IdentityProvider) OpenID() *OpenIdidentityProvider {
	if o == nil {
		return nil
	}
	return o.openID
}

// GetOpenID returns the value of the 'open_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Details for `openid` identity providers.
func (o *IdentityProvider) GetOpenID() (value *OpenIdidentityProvider, ok bool) {
	ok = o != nil && o.openID != nil
	if ok {
		value = o.openID
	}
	return
}

// IdentityProviderListKind is the name of the type used to represent list of
// objects of type 'identity_provider'.
const IdentityProviderListKind = "IdentityProviderList"

// IdentityProviderListLinkKind is the name of the type used to represent links
// to list of objects of type 'identity_provider'.
const IdentityProviderListLinkKind = "IdentityProviderListLink"

// IdentityProviderNilKind is the name of the type used to nil lists of
// objects of type 'identity_provider'.
const IdentityProviderListNilKind = "IdentityProviderListNil"

// IdentityProviderList is a list of values of the 'identity_provider' type.
type IdentityProviderList struct {
	href  *string
	link  bool
	items []*IdentityProvider
}

// Kind returns the name of the type of the object.
func (l *IdentityProviderList) Kind() string {
	if l == nil {
		return IdentityProviderListNilKind
	}
	if l.link {
		return IdentityProviderListLinkKind
	}
	return IdentityProviderListKind
}

// Link returns true iif this is a link.
func (l *IdentityProviderList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *IdentityProviderList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *IdentityProviderList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *IdentityProviderList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *IdentityProviderList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *IdentityProviderList) Get(i int) *IdentityProvider {
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
func (l *IdentityProviderList) Slice() []*IdentityProvider {
	var slice []*IdentityProvider
	if l == nil {
		slice = make([]*IdentityProvider, 0)
	} else {
		slice = make([]*IdentityProvider, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *IdentityProviderList) Each(f func(item *IdentityProvider) bool) {
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
func (l *IdentityProviderList) Range(f func(index int, item *IdentityProvider) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
