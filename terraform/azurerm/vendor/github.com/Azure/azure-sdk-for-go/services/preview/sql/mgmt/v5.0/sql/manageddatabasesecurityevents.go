package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ManagedDatabaseSecurityEventsClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ManagedDatabaseSecurityEventsClient struct {
	BaseClient
}

// NewManagedDatabaseSecurityEventsClient creates an instance of the ManagedDatabaseSecurityEventsClient client.
func NewManagedDatabaseSecurityEventsClient(subscriptionID string) ManagedDatabaseSecurityEventsClient {
	return NewManagedDatabaseSecurityEventsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedDatabaseSecurityEventsClientWithBaseURI creates an instance of the ManagedDatabaseSecurityEventsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewManagedDatabaseSecurityEventsClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabaseSecurityEventsClient {
	return ManagedDatabaseSecurityEventsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByDatabase gets a list of security events.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the managed database for which the security events are retrieved.
// filter - an OData filter expression that filters elements in the collection.
// skip - the number of elements in the collection to skip.
// top - the number of elements to return from the collection.
// skiptoken - an opaque token that identifies a starting point in the collection.
func (client ManagedDatabaseSecurityEventsClient) ListByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, filter string, skip *int32, top *int32, skiptoken string) (result SecurityEventCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSecurityEventsClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.sec.Response.Response != nil {
				sc = result.sec.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDatabaseNextResults
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, filter, skip, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSecurityEventsClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.sec.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSecurityEventsClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result.sec, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSecurityEventsClient", "ListByDatabase", resp, "Failure responding to request")
		return
	}
	if result.sec.hasNextLink() && result.sec.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client ManagedDatabaseSecurityEventsClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, filter string, skip *int32, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/securityEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSecurityEventsClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSecurityEventsClient) ListByDatabaseResponder(resp *http.Response) (result SecurityEventCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDatabaseNextResults retrieves the next set of results, if any.
func (client ManagedDatabaseSecurityEventsClient) listByDatabaseNextResults(ctx context.Context, lastResults SecurityEventCollection) (result SecurityEventCollection, err error) {
	req, err := lastResults.securityEventCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseSecurityEventsClient", "listByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseSecurityEventsClient", "listByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSecurityEventsClient", "listByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedDatabaseSecurityEventsClient) ListByDatabaseComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, filter string, skip *int32, top *int32, skiptoken string) (result SecurityEventCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSecurityEventsClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDatabase(ctx, resourceGroupName, managedInstanceName, databaseName, filter, skip, top, skiptoken)
	return
}