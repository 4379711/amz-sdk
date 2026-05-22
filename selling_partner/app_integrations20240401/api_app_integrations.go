package app_integrations20240401

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AppIntegrationsAPIService AppIntegrationsAPI service
type AppIntegrationsAPIService service

type ApiCreateNotificationRequest struct {
	ctx        context.Context
	ApiService *AppIntegrationsAPIService
	body       *CreateNotificationRequest
}

// The request body for the &#x60;createNotification&#x60; operation.
func (r ApiCreateNotificationRequest) Body(body CreateNotificationRequest) ApiCreateNotificationRequest {
	r.body = &body
	return r
}

func (r ApiCreateNotificationRequest) Execute() (*CreateNotificationResponse, *http.Response, error) {
	return r.ApiService.CreateNotificationExecute(r)
}

/*
CreateNotification Method for CreateNotification

Create a notification for sellers in Seller Central.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Sellers whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateNotificationRequest
*/
func (a *AppIntegrationsAPIService) CreateNotification(ctx context.Context) ApiCreateNotificationRequest {
	return ApiCreateNotificationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateNotificationResponse
func (a *AppIntegrationsAPIService) CreateNotificationExecute(r ApiCreateNotificationRequest) (*CreateNotificationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateNotificationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppIntegrationsAPIService.CreateNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appIntegrations/2024-04-01/notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 400 {
		err = &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
	} else {
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	}
	return localVarReturnValue, localVarHTTPResponse, err
}

type ApiDeleteNotificationsRequest struct {
	ctx        context.Context
	ApiService *AppIntegrationsAPIService
	body       *DeleteNotificationsRequest
}

// The request body for the &#x60;deleteNotifications&#x60; operation.
func (r ApiDeleteNotificationsRequest) Body(body DeleteNotificationsRequest) ApiDeleteNotificationsRequest {
	r.body = &body
	return r
}

func (r ApiDeleteNotificationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNotificationsExecute(r)
}

/*
DeleteNotifications Method for DeleteNotifications

Remove your application's notifications from the Appstore notifications dashboard.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Sellers whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeleteNotificationsRequest
*/
func (a *AppIntegrationsAPIService) DeleteNotifications(ctx context.Context) ApiDeleteNotificationsRequest {
	return ApiDeleteNotificationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AppIntegrationsAPIService) DeleteNotificationsExecute(r ApiDeleteNotificationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppIntegrationsAPIService.DeleteNotifications")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appIntegrations/2024-04-01/notifications/deletion"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	return localVarHTTPResponse, nil
}

type ApiRecordActionFeedbackRequest struct {
	ctx            context.Context
	ApiService     *AppIntegrationsAPIService
	body           *RecordActionFeedbackRequest
	notificationId string
}

// The request body for the &#x60;recordActionFeedback&#x60; operation.
func (r ApiRecordActionFeedbackRequest) Body(body RecordActionFeedbackRequest) ApiRecordActionFeedbackRequest {
	r.body = &body
	return r
}

func (r ApiRecordActionFeedbackRequest) Execute() (*http.Response, error) {
	return r.ApiService.RecordActionFeedbackExecute(r)
}

/*
RecordActionFeedback Method for RecordActionFeedback

Records the seller's response to a notification.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Sellers whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationId A `notificationId` uniquely identifies a notification.
	@return ApiRecordActionFeedbackRequest
*/
func (a *AppIntegrationsAPIService) RecordActionFeedback(ctx context.Context, notificationId string) ApiRecordActionFeedbackRequest {
	return ApiRecordActionFeedbackRequest{
		ApiService:     a,
		ctx:            ctx,
		notificationId: notificationId,
	}
}

// Execute executes the request
func (a *AppIntegrationsAPIService) RecordActionFeedbackExecute(r ApiRecordActionFeedbackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppIntegrationsAPIService.RecordActionFeedback")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appIntegrations/2024-04-01/notifications/{notificationId}/feedback"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationId"+"}", url.PathEscape(parameterValueToString(r.notificationId, "notificationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	return localVarHTTPResponse, nil
}
