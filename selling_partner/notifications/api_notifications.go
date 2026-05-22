package notifications

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// NotificationsAPIService NotificationsAPI service
type NotificationsAPIService service

type ApiCreateDestinationRequest struct {
	ctx        context.Context
	ApiService *NotificationsAPIService
	body       *CreateDestinationRequest
}

func (r ApiCreateDestinationRequest) Body(body CreateDestinationRequest) ApiCreateDestinationRequest {
	r.body = &body
	return r
}

func (r ApiCreateDestinationRequest) Execute() (*CreateDestinationResponse, *http.Response, error) {
	return r.ApiService.CreateDestinationExecute(r)
}

/*
CreateDestination Method for CreateDestination

Creates a destination resource to receive notifications. The `createDestination` operation is grantless. For more information, refer to [Grantless operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateDestinationRequest
*/
func (a *NotificationsAPIService) CreateDestination(ctx context.Context) ApiCreateDestinationRequest {
	return ApiCreateDestinationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateDestinationResponse
func (a *NotificationsAPIService) CreateDestinationExecute(r ApiCreateDestinationRequest) (*CreateDestinationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateDestinationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.CreateDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/destinations"

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
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Response"}

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

type ApiCreateSubscriptionRequest struct {
	ctx              context.Context
	ApiService       *NotificationsAPIService
	body             *CreateSubscriptionRequest
	notificationType string
}

func (r ApiCreateSubscriptionRequest) Body(body CreateSubscriptionRequest) ApiCreateSubscriptionRequest {
	r.body = &body
	return r
}

func (r ApiCreateSubscriptionRequest) Execute() (*CreateSubscriptionResponse, *http.Response, error) {
	return r.ApiService.CreateSubscriptionExecute(r)
}

/*
CreateSubscription Method for CreateSubscription

Creates a subscription for the specified notification type to be delivered to the specified destination. Before you can subscribe, you must first create the destination by calling the `createDestination` operation. In cases where the specified notification type supports multiple payload versions, you can utilize this API to subscribe to a different payload version if you already have an existing subscription for a different payload version.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationType The type of notification.   For more information about notification types, refer to [Notification Type Values](https://developer-docs.amazon.com/sp-api/docs/notification-type-values).
	@return ApiCreateSubscriptionRequest
*/
func (a *NotificationsAPIService) CreateSubscription(ctx context.Context, notificationType string) ApiCreateSubscriptionRequest {
	return ApiCreateSubscriptionRequest{
		ApiService:       a,
		ctx:              ctx,
		notificationType: notificationType,
	}
}

// Execute executes the request
//
//	@return CreateSubscriptionResponse
func (a *NotificationsAPIService) CreateSubscriptionExecute(r ApiCreateSubscriptionRequest) (*CreateSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.CreateSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/subscriptions/{notificationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterValueToString(r.notificationType, "notificationType")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Response"}

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

type ApiDeleteDestinationRequest struct {
	ctx           context.Context
	ApiService    *NotificationsAPIService
	destinationId string
}

func (r ApiDeleteDestinationRequest) Execute() (*DeleteDestinationResponse, *http.Response, error) {
	return r.ApiService.DeleteDestinationExecute(r)
}

/*
DeleteDestination Method for DeleteDestination

Deletes the destination that you specify. The `deleteDestination` operation is grantless. For more information, refer to [Grantless operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param destinationId The identifier for the destination that you want to delete.
	@return ApiDeleteDestinationRequest
*/
func (a *NotificationsAPIService) DeleteDestination(ctx context.Context, destinationId string) ApiDeleteDestinationRequest {
	return ApiDeleteDestinationRequest{
		ApiService:    a,
		ctx:           ctx,
		destinationId: destinationId,
	}
}

// Execute executes the request
//
//	@return DeleteDestinationResponse
func (a *NotificationsAPIService) DeleteDestinationExecute(r ApiDeleteDestinationRequest) (*DeleteDestinationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteDestinationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.DeleteDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/destinations/{destinationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"destinationId"+"}", url.PathEscape(parameterValueToString(r.destinationId, "destinationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Response"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiDeleteSubscriptionByIdRequest struct {
	ctx              context.Context
	ApiService       *NotificationsAPIService
	subscriptionId   string
	notificationType string
}

func (r ApiDeleteSubscriptionByIdRequest) Execute() (*DeleteSubscriptionByIdResponse, *http.Response, error) {
	return r.ApiService.DeleteSubscriptionByIdExecute(r)
}

/*
DeleteSubscriptionById Method for DeleteSubscriptionById

Deletes the subscription indicated by the subscription identifier and notification type that you specify. The subscription identifier can be for any subscription associated with your application. After you successfully call this operation, notifications will stop being sent for the associated subscription. The `deleteSubscriptionById` operation is grantless. For more information, refer to [Grantless operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subscriptionId The identifier for the subscription that you want to delete.
	@param notificationType The type of notification.   For more information about notification types, refer to [Notification Type Values](https://developer-docs.amazon.com/sp-api/docs/notification-type-values).
	@return ApiDeleteSubscriptionByIdRequest
*/
func (a *NotificationsAPIService) DeleteSubscriptionById(ctx context.Context, subscriptionId string, notificationType string) ApiDeleteSubscriptionByIdRequest {
	return ApiDeleteSubscriptionByIdRequest{
		ApiService:       a,
		ctx:              ctx,
		subscriptionId:   subscriptionId,
		notificationType: notificationType,
	}
}

// Execute executes the request
//
//	@return DeleteSubscriptionByIdResponse
func (a *NotificationsAPIService) DeleteSubscriptionByIdExecute(r ApiDeleteSubscriptionByIdRequest) (*DeleteSubscriptionByIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteSubscriptionByIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.DeleteSubscriptionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/subscriptions/{notificationType}/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterValueToString(r.notificationType, "notificationType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Operation Response"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiGetDestinationRequest struct {
	ctx           context.Context
	ApiService    *NotificationsAPIService
	destinationId string
}

func (r ApiGetDestinationRequest) Execute() (*GetDestinationResponse, *http.Response, error) {
	return r.ApiService.GetDestinationExecute(r)
}

/*
GetDestination Method for GetDestination

Returns information about the destination that you specify. The `getDestination` operation is grantless. For more information, refer to [Grantless operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param destinationId The identifier generated when you created the destination.
	@return ApiGetDestinationRequest
*/
func (a *NotificationsAPIService) GetDestination(ctx context.Context, destinationId string) ApiGetDestinationRequest {
	return ApiGetDestinationRequest{
		ApiService:    a,
		ctx:           ctx,
		destinationId: destinationId,
	}
}

// Execute executes the request
//
//	@return GetDestinationResponse
func (a *NotificationsAPIService) GetDestinationExecute(r ApiGetDestinationRequest) (*GetDestinationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDestinationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.GetDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/destinations/{destinationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"destinationId"+"}", url.PathEscape(parameterValueToString(r.destinationId, "destinationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Response"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiGetDestinationsRequest struct {
	ctx        context.Context
	ApiService *NotificationsAPIService
}

func (r ApiGetDestinationsRequest) Execute() (*GetDestinationsResponse, *http.Response, error) {
	return r.ApiService.GetDestinationsExecute(r)
}

/*
GetDestinations Method for GetDestinations

Returns information about all destinations. The `getDestinations` operation is grantless. For more information, refer to [Grantless operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDestinationsRequest
*/
func (a *NotificationsAPIService) GetDestinations(ctx context.Context) ApiGetDestinationsRequest {
	return ApiGetDestinationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDestinationsResponse
func (a *NotificationsAPIService) GetDestinationsExecute(r ApiGetDestinationsRequest) (*GetDestinationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDestinationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.GetDestinations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/destinations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Response"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiGetSubscriptionRequest struct {
	ctx              context.Context
	ApiService       *NotificationsAPIService
	notificationType string
	payloadVersion   *string
}

// The version of the payload object to be used in the notification.
func (r ApiGetSubscriptionRequest) PayloadVersion(payloadVersion string) ApiGetSubscriptionRequest {
	r.payloadVersion = &payloadVersion
	return r
}

func (r ApiGetSubscriptionRequest) Execute() (*GetSubscriptionResponse, *http.Response, error) {
	return r.ApiService.GetSubscriptionExecute(r)
}

/*
GetSubscription Method for GetSubscription

Returns information about subscription of the specified notification type and payload version. `payloadVersion` is an optional parameter. When `payloadVersion` is not provided, it will return latest payload version subscription's information. You can use this API to get subscription information when you do not have a subscription identifier.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationType The type of notification.   For more information about notification types, refer to [Notification Type Values](https://developer-docs.amazon.com/sp-api/docs/notification-type-values).
	@return ApiGetSubscriptionRequest
*/
func (a *NotificationsAPIService) GetSubscription(ctx context.Context, notificationType string) ApiGetSubscriptionRequest {
	return ApiGetSubscriptionRequest{
		ApiService:       a,
		ctx:              ctx,
		notificationType: notificationType,
	}
}

// Execute executes the request
//
//	@return GetSubscriptionResponse
func (a *NotificationsAPIService) GetSubscriptionExecute(r ApiGetSubscriptionRequest) (*GetSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.GetSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/subscriptions/{notificationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterValueToString(r.notificationType, "notificationType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.payloadVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payloadVersion", r.payloadVersion, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Response"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiGetSubscriptionByIdRequest struct {
	ctx              context.Context
	ApiService       *NotificationsAPIService
	subscriptionId   string
	notificationType string
}

func (r ApiGetSubscriptionByIdRequest) Execute() (*GetSubscriptionByIdResponse, *http.Response, error) {
	return r.ApiService.GetSubscriptionByIdExecute(r)
}

/*
GetSubscriptionById Method for GetSubscriptionById

Returns information about a subscription for the specified notification type. The `getSubscriptionById` operation is grantless. For more information, refer to [Grantless operations](https://developer-docs.amazon.com/sp-api/docs/grantless-operations).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may observe higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subscriptionId The identifier for the subscription that you want to get.
	@param notificationType The type of notification.   For more information about notification types, refer to [Notification Type Values](https://developer-docs.amazon.com/sp-api/docs/notification-type-values).
	@return ApiGetSubscriptionByIdRequest
*/
func (a *NotificationsAPIService) GetSubscriptionById(ctx context.Context, subscriptionId string, notificationType string) ApiGetSubscriptionByIdRequest {
	return ApiGetSubscriptionByIdRequest{
		ApiService:       a,
		ctx:              ctx,
		subscriptionId:   subscriptionId,
		notificationType: notificationType,
	}
}

// Execute executes the request
//
//	@return GetSubscriptionByIdResponse
func (a *NotificationsAPIService) GetSubscriptionByIdExecute(r ApiGetSubscriptionByIdRequest) (*GetSubscriptionByIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSubscriptionByIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationsAPIService.GetSubscriptionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/v1/subscriptions/{notificationType}/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterValueToString(r.notificationType, "notificationType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "Successful Response"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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
