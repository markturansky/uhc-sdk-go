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

// ClusterRegistrationResponseBuilder contains the data and logic needed to build 'cluster_registration_response' objects.
//
//
type ClusterRegistrationResponseBuilder struct {
	clusterID          *string
	authorizationToken *string
	accountID          *string
	expiresAt          *int64
}

// NewClusterRegistrationResponse creates a new builder of 'cluster_registration_response' objects.
func NewClusterRegistrationResponse() *ClusterRegistrationResponseBuilder {
	return new(ClusterRegistrationResponseBuilder)
}

// ClusterID sets the value of the 'cluster_ID' attribute
// to the given value.
//
//
func (b *ClusterRegistrationResponseBuilder) ClusterID(value string) *ClusterRegistrationResponseBuilder {
	b.clusterID = &value
	return b
}

// AuthorizationToken sets the value of the 'authorization_token' attribute
// to the given value.
//
//
func (b *ClusterRegistrationResponseBuilder) AuthorizationToken(value string) *ClusterRegistrationResponseBuilder {
	b.authorizationToken = &value
	return b
}

// AccountID sets the value of the 'account_ID' attribute
// to the given value.
//
//
func (b *ClusterRegistrationResponseBuilder) AccountID(value string) *ClusterRegistrationResponseBuilder {
	b.accountID = &value
	return b
}

// ExpiresAt sets the value of the 'expires_at' attribute
// to the given value.
//
//
func (b *ClusterRegistrationResponseBuilder) ExpiresAt(value int64) *ClusterRegistrationResponseBuilder {
	b.expiresAt = &value
	return b
}

// Build creates a 'cluster_registration_response' object using the configuration stored in the builder.
func (b *ClusterRegistrationResponseBuilder) Build() (object *ClusterRegistrationResponse, err error) {
	object = new(ClusterRegistrationResponse)
	if b.clusterID != nil {
		object.clusterID = b.clusterID
	}
	if b.authorizationToken != nil {
		object.authorizationToken = b.authorizationToken
	}
	if b.accountID != nil {
		object.accountID = b.accountID
	}
	if b.expiresAt != nil {
		object.expiresAt = b.expiresAt
	}
	return
}
