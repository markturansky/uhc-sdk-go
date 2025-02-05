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

package accountsmgmt // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt

import (
	"net/http"
	"path"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1"
)

// Client is the client for service 'accounts_mgmt'.
type Client struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewClient creates a new client for the service 'accounts_mgmt' using the
// given transport to send the requests and receive the responses.
func NewClient(transport http.RoundTripper, path string, metric string) *Client {
	client := new(Client)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// V1 returns a reference to a client for version 'v1'.
func (c *Client) V1() *v1.RootClient {
	return v1.NewRootClient(
		c.transport,
		path.Join(c.path, "v1"),
		path.Join(c.metric, "v1"),
	)
}
