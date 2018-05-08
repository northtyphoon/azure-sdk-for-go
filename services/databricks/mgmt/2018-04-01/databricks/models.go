package databricks

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/go.uuid"
	"net/http"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Created ...
	Created ProvisioningState = "Created"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleted ...
	Deleted ProvisioningState = "Deleted"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Ready ...
	Ready ProvisioningState = "Ready"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, Ready, Running, Succeeded, Updating}
}

// ErrorDetail ...
type ErrorDetail struct {
	// Code - The error's code.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
	// Target - Indicates which property in the request is responsible for the error.
	Target *string `json:"target,omitempty"`
}

// ErrorInfo ...
type ErrorInfo struct {
	// Code - A machine readable error code.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
	// Details - error details.
	Details *[]ErrorDetail `json:"details,omitempty"`
	// Innererror - Inner error details if they exist.
	Innererror *string `json:"innererror,omitempty"`
}

// ErrorResponse contains details when the response code indicates an error.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorInfo `json:"error,omitempty"`
}

// Operation REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.ResourceProvider
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list Resource Provider operations. It contains a list of operations
// and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Resource Provider operations supported by the Resource Provider resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Resource the core properties of ARM resources
type Resource struct {
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Sku SKU for the resource.
type Sku struct {
	// Name - The SKU name.
	Name *string `json:"name,omitempty"`
	// Tier - The SKU tier.
	Tier *string `json:"tier,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	if tr.ID != nil {
		objectMap["id"] = tr.ID
	}
	if tr.Name != nil {
		objectMap["name"] = tr.Name
	}
	if tr.Type != nil {
		objectMap["type"] = tr.Type
	}
	return json.Marshal(objectMap)
}

// Workspace information about workspace.
type Workspace struct {
	autorest.Response `json:"-"`
	// WorkspaceProperties - The workspace properties.
	*WorkspaceProperties `json:"properties,omitempty"`
	// Sku - The SKU of the resource.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if w.WorkspaceProperties != nil {
		objectMap["properties"] = w.WorkspaceProperties
	}
	if w.Sku != nil {
		objectMap["sku"] = w.Sku
	}
	if w.Tags != nil {
		objectMap["tags"] = w.Tags
	}
	if w.Location != nil {
		objectMap["location"] = w.Location
	}
	if w.ID != nil {
		objectMap["id"] = w.ID
	}
	if w.Name != nil {
		objectMap["name"] = w.Name
	}
	if w.Type != nil {
		objectMap["type"] = w.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Workspace struct.
func (w *Workspace) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var workspaceProperties WorkspaceProperties
				err = json.Unmarshal(*v, &workspaceProperties)
				if err != nil {
					return err
				}
				w.WorkspaceProperties = &workspaceProperties
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				w.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				w.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				w.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				w.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				w.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				w.Type = &typeVar
			}
		}
	}

	return nil
}

// WorkspaceListResult list of workspaces.
type WorkspaceListResult struct {
	autorest.Response `json:"-"`
	// Value - The array of workspaces.
	Value *[]Workspace `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// WorkspaceListResultIterator provides access to a complete listing of Workspace values.
type WorkspaceListResultIterator struct {
	i    int
	page WorkspaceListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *WorkspaceListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter WorkspaceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter WorkspaceListResultIterator) Response() WorkspaceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter WorkspaceListResultIterator) Value() Workspace {
	if !iter.page.NotDone() {
		return Workspace{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (wlr WorkspaceListResult) IsEmpty() bool {
	return wlr.Value == nil || len(*wlr.Value) == 0
}

// workspaceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (wlr WorkspaceListResult) workspaceListResultPreparer() (*http.Request, error) {
	if wlr.NextLink == nil || len(to.String(wlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(wlr.NextLink)))
}

// WorkspaceListResultPage contains a page of Workspace values.
type WorkspaceListResultPage struct {
	fn  func(WorkspaceListResult) (WorkspaceListResult, error)
	wlr WorkspaceListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *WorkspaceListResultPage) Next() error {
	next, err := page.fn(page.wlr)
	if err != nil {
		return err
	}
	page.wlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page WorkspaceListResultPage) NotDone() bool {
	return !page.wlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page WorkspaceListResultPage) Response() WorkspaceListResult {
	return page.wlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page WorkspaceListResultPage) Values() []Workspace {
	if page.wlr.IsEmpty() {
		return nil
	}
	return *page.wlr.Value
}

// WorkspaceProperties the workspace properties.
type WorkspaceProperties struct {
	// ManagedResourceGroupID - The managed resource group Id.
	ManagedResourceGroupID *string `json:"managedResourceGroupId,omitempty"`
	// Parameters - Name and value pairs that define the workspace parameters.
	Parameters interface{} `json:"parameters,omitempty"`
	// ProvisioningState - The workspace provisioning state. Possible values include: 'Accepted', 'Running', 'Ready', 'Creating', 'Created', 'Deleting', 'Deleted', 'Canceled', 'Failed', 'Succeeded', 'Updating'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// UIDefinitionURI - The blob URI where the UI definition file is located.
	UIDefinitionURI *string `json:"uiDefinitionUri,omitempty"`
	// Authorizations - The workspace provider authorizations.
	Authorizations *[]WorkspaceProviderAuthorization `json:"authorizations,omitempty"`
}

// WorkspaceProviderAuthorization the workspace provider authorization.
type WorkspaceProviderAuthorization struct {
	// PrincipalID - The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the workspace resources.
	PrincipalID *uuid.UUID `json:"principalId,omitempty"`
	// RoleDefinitionID - The provider's role definition identifier. This role will define all the permissions that the provider must have on the workspace's container resource group. This role definition cannot have permission to delete the resource group.
	RoleDefinitionID *uuid.UUID `json:"roleDefinitionId,omitempty"`
}

// WorkspacesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type WorkspacesCreateOrUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WorkspacesCreateOrUpdateFuture) Result(client WorkspacesClient) (w Workspace, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return w, azure.NewAsyncOpIncompleteError("databricks.WorkspacesCreateOrUpdateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		w, err = client.CreateOrUpdateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "databricks.WorkspacesCreateOrUpdateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesCreateOrUpdateFuture", "Result", resp, "Failure sending request")
		return
	}
	w, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesCreateOrUpdateFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// WorkspacesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type WorkspacesDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WorkspacesDeleteFuture) Result(client WorkspacesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return ar, azure.NewAsyncOpIncompleteError("databricks.WorkspacesDeleteFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "databricks.WorkspacesDeleteFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesDeleteFuture", "Result", resp, "Failure sending request")
		return
	}
	ar, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesDeleteFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// WorkspacesUpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type WorkspacesUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WorkspacesUpdateFuture) Result(client WorkspacesClient) (w Workspace, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return w, azure.NewAsyncOpIncompleteError("databricks.WorkspacesUpdateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		w, err = client.UpdateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "databricks.WorkspacesUpdateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesUpdateFuture", "Result", resp, "Failure sending request")
		return
	}
	w, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesUpdateFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// WorkspaceUpdate an update to a workspace.
type WorkspaceUpdate struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for WorkspaceUpdate.
func (wu WorkspaceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if wu.Tags != nil {
		objectMap["tags"] = wu.Tags
	}
	return json.Marshal(objectMap)
}
