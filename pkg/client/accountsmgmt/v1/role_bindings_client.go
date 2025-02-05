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

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// RoleBindingsClient is the client of the 'role_bindings' resource.
//
// Manages the collection of role bindings.
type RoleBindingsClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewRoleBindingsClient creates a new client for the 'role_bindings'
// resource using the given transport to sned the requests and receive the
// responses.
func NewRoleBindingsClient(transport http.RoundTripper, path string, metric string) *RoleBindingsClient {
	client := new(RoleBindingsClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves a list of role bindings.
func (c *RoleBindingsClient) List() *RoleBindingsListRequest {
	request := new(RoleBindingsListRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// Add creates a request for the 'add' method.
//
// Creates a new role binding.
func (c *RoleBindingsClient) Add() *RoleBindingsAddRequest {
	request := new(RoleBindingsAddRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// RoleBinding returns the target 'role_binding' resource for the given identifier.
//
// Reference to the service that manages a specific role binding.
func (c *RoleBindingsClient) RoleBinding(id string) *RoleBindingClient {
	return NewRoleBindingClient(
		c.transport,
		path.Join(c.path, id),
		path.Join(c.metric, "-"),
	)
}

// RoleBindingsListRequest is the request for the 'list' method.
type RoleBindingsListRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	page      *int
	size      *int
	total     *int
}

// Parameter adds a query parameter.
func (r *RoleBindingsListRequest) Parameter(name string, value interface{}) *RoleBindingsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RoleBindingsListRequest) Header(name string, value interface{}) *RoleBindingsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RoleBindingsListRequest) Page(value int) *RoleBindingsListRequest {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RoleBindingsListRequest) Size(value int) *RoleBindingsListRequest {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RoleBindingsListRequest) Total(value int) *RoleBindingsListRequest {
	r.total = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *RoleBindingsListRequest) Send() (result *RoleBindingsListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *RoleBindingsListRequest) SendContext(ctx context.Context) (result *RoleBindingsListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	if r.total != nil {
		helpers.AddValue(&query, "total", *r.total)
	}
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(RoleBindingsListResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// RoleBindingsListResponse is the response for the 'list' method.
type RoleBindingsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *RoleBindingList
}

// Status returns the response status code.
func (r *RoleBindingsListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *RoleBindingsListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *RoleBindingsListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RoleBindingsListResponse) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *RoleBindingsListResponse) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RoleBindingsListResponse) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *RoleBindingsListResponse) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RoleBindingsListResponse) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *RoleBindingsListResponse) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of role bindings.
func (r *RoleBindingsListResponse) Items() *RoleBindingList {
	if r == nil {
		return nil
	}
	return r.items
}

// GetItems returns the value of the 'items' parameter and
// a flag indicating if the parameter has a value.
//
// Retrieved list of role bindings.
func (r *RoleBindingsListResponse) GetItems() (value *RoleBindingList, ok bool) {
	ok = r != nil && r.items != nil
	if ok {
		value = r.items
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *RoleBindingsListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(roleBindingsListResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.page = data.Page
	r.size = data.Size
	r.total = data.Total
	r.items, err = data.Items.unwrap()
	if err != nil {
		return err
	}
	return err
}

// roleBindingsListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type roleBindingsListResponseData struct {
	Page  *int                "json:\"page,omitempty\""
	Size  *int                "json:\"size,omitempty\""
	Total *int                "json:\"total,omitempty\""
	Items roleBindingListData "json:\"items,omitempty\""
}

// RoleBindingsAddRequest is the request for the 'add' method.
type RoleBindingsAddRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	body      *RoleBinding
}

// Parameter adds a query parameter.
func (r *RoleBindingsAddRequest) Parameter(name string, value interface{}) *RoleBindingsAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RoleBindingsAddRequest) Header(name string, value interface{}) *RoleBindingsAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Role binding data.
func (r *RoleBindingsAddRequest) Body(value *RoleBinding) *RoleBindingsAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *RoleBindingsAddRequest) Send() (result *RoleBindingsAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *RoleBindingsAddRequest) SendContext(ctx context.Context) (result *RoleBindingsAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	buffer := new(bytes.Buffer)
	err = r.marshal(buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodPost,
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(RoleBindingsAddResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'add' method.
func (r *RoleBindingsAddRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// RoleBindingsAddResponse is the response for the 'add' method.
type RoleBindingsAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *RoleBinding
}

// Status returns the response status code.
func (r *RoleBindingsAddResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *RoleBindingsAddResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *RoleBindingsAddResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Role binding data.
func (r *RoleBindingsAddResponse) Body() *RoleBinding {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Role binding data.
func (r *RoleBindingsAddResponse) GetBody() (value *RoleBinding, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'add' method.
func (r *RoleBindingsAddResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(roleBindingData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}
