package messaging

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MessagingAPIService MessagingAPI service
type MessagingAPIService service

type ApiConfirmCustomizationDetailsRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateConfirmCustomizationDetailsRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiConfirmCustomizationDetailsRequest) Body(body CreateConfirmCustomizationDetailsRequest) ApiConfirmCustomizationDetailsRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiConfirmCustomizationDetailsRequest) MarketplaceIds(marketplaceIds []string) ApiConfirmCustomizationDetailsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiConfirmCustomizationDetailsRequest) Execute() (*CreateConfirmCustomizationDetailsResponse, *http.Response, error) {
	return r.ApiService.ConfirmCustomizationDetailsExecute(r)
}

/*
ConfirmCustomizationDetails Method for ConfirmCustomizationDetails

Sends a message asking a buyer to provide or verify customization details such as name spelling, images, initials, etc.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiConfirmCustomizationDetailsRequest
*/
func (a *MessagingAPIService) ConfirmCustomizationDetails(ctx context.Context, amazonOrderId string) ApiConfirmCustomizationDetailsRequest {
	return ApiConfirmCustomizationDetailsRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateConfirmCustomizationDetailsResponse
func (a *MessagingAPIService) ConfirmCustomizationDetailsExecute(r ApiConfirmCustomizationDetailsRequest) (*CreateConfirmCustomizationDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateConfirmCustomizationDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.ConfirmCustomizationDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateAmazonMotorsRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateAmazonMotorsRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateAmazonMotorsRequest) Body(body CreateAmazonMotorsRequest) ApiCreateAmazonMotorsRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateAmazonMotorsRequest) MarketplaceIds(marketplaceIds []string) ApiCreateAmazonMotorsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateAmazonMotorsRequest) Execute() (*CreateAmazonMotorsResponse, *http.Response, error) {
	return r.ApiService.CreateAmazonMotorsExecute(r)
}

/*
CreateAmazonMotors Method for CreateAmazonMotors

Sends a message to a buyer to provide details about an Amazon Motors order. This message can only be sent by Amazon Motors sellers.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateAmazonMotorsRequest
*/
func (a *MessagingAPIService) CreateAmazonMotors(ctx context.Context, amazonOrderId string) ApiCreateAmazonMotorsRequest {
	return ApiCreateAmazonMotorsRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateAmazonMotorsResponse
func (a *MessagingAPIService) CreateAmazonMotorsExecute(r ApiCreateAmazonMotorsRequest) (*CreateAmazonMotorsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateAmazonMotorsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateAmazonMotors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/amazonMotors"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateConfirmDeliveryDetailsRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateConfirmDeliveryDetailsRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateConfirmDeliveryDetailsRequest) Body(body CreateConfirmDeliveryDetailsRequest) ApiCreateConfirmDeliveryDetailsRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateConfirmDeliveryDetailsRequest) MarketplaceIds(marketplaceIds []string) ApiCreateConfirmDeliveryDetailsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateConfirmDeliveryDetailsRequest) Execute() (*CreateConfirmDeliveryDetailsResponse, *http.Response, error) {
	return r.ApiService.CreateConfirmDeliveryDetailsExecute(r)
}

/*
CreateConfirmDeliveryDetails Method for CreateConfirmDeliveryDetails

Sends a message to a buyer to arrange a delivery or to confirm contact information for making a delivery.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateConfirmDeliveryDetailsRequest
*/
func (a *MessagingAPIService) CreateConfirmDeliveryDetails(ctx context.Context, amazonOrderId string) ApiCreateConfirmDeliveryDetailsRequest {
	return ApiCreateConfirmDeliveryDetailsRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateConfirmDeliveryDetailsResponse
func (a *MessagingAPIService) CreateConfirmDeliveryDetailsExecute(r ApiCreateConfirmDeliveryDetailsRequest) (*CreateConfirmDeliveryDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateConfirmDeliveryDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateConfirmDeliveryDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateConfirmOrderDetailsRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateConfirmOrderDetailsRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateConfirmOrderDetailsRequest) Body(body CreateConfirmOrderDetailsRequest) ApiCreateConfirmOrderDetailsRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateConfirmOrderDetailsRequest) MarketplaceIds(marketplaceIds []string) ApiCreateConfirmOrderDetailsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateConfirmOrderDetailsRequest) Execute() (*CreateConfirmOrderDetailsResponse, *http.Response, error) {
	return r.ApiService.CreateConfirmOrderDetailsExecute(r)
}

/*
CreateConfirmOrderDetails Method for CreateConfirmOrderDetails

Sends a message to ask a buyer an order-related question prior to shipping their order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateConfirmOrderDetailsRequest
*/
func (a *MessagingAPIService) CreateConfirmOrderDetails(ctx context.Context, amazonOrderId string) ApiCreateConfirmOrderDetailsRequest {
	return ApiCreateConfirmOrderDetailsRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateConfirmOrderDetailsResponse
func (a *MessagingAPIService) CreateConfirmOrderDetailsExecute(r ApiCreateConfirmOrderDetailsRequest) (*CreateConfirmOrderDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateConfirmOrderDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateConfirmOrderDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateConfirmServiceDetailsRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateConfirmServiceDetailsRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateConfirmServiceDetailsRequest) Body(body CreateConfirmServiceDetailsRequest) ApiCreateConfirmServiceDetailsRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateConfirmServiceDetailsRequest) MarketplaceIds(marketplaceIds []string) ApiCreateConfirmServiceDetailsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateConfirmServiceDetailsRequest) Execute() (*CreateConfirmServiceDetailsResponse, *http.Response, error) {
	return r.ApiService.CreateConfirmServiceDetailsExecute(r)
}

/*
CreateConfirmServiceDetails Method for CreateConfirmServiceDetails

Sends a message to contact a Home Service customer to arrange a service call or to gather information prior to a service call.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateConfirmServiceDetailsRequest
*/
func (a *MessagingAPIService) CreateConfirmServiceDetails(ctx context.Context, amazonOrderId string) ApiCreateConfirmServiceDetailsRequest {
	return ApiCreateConfirmServiceDetailsRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateConfirmServiceDetailsResponse
func (a *MessagingAPIService) CreateConfirmServiceDetailsExecute(r ApiCreateConfirmServiceDetailsRequest) (*CreateConfirmServiceDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateConfirmServiceDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateConfirmServiceDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/confirmServiceDetails"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateDigitalAccessKeyRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateDigitalAccessKeyRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateDigitalAccessKeyRequest) Body(body CreateDigitalAccessKeyRequest) ApiCreateDigitalAccessKeyRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateDigitalAccessKeyRequest) MarketplaceIds(marketplaceIds []string) ApiCreateDigitalAccessKeyRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateDigitalAccessKeyRequest) Execute() (*CreateDigitalAccessKeyResponse, *http.Response, error) {
	return r.ApiService.CreateDigitalAccessKeyExecute(r)
}

/*
CreateDigitalAccessKey Method for CreateDigitalAccessKey

Sends a buyer a message to share a digital access key that is required to utilize digital content in their order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateDigitalAccessKeyRequest
*/
func (a *MessagingAPIService) CreateDigitalAccessKey(ctx context.Context, amazonOrderId string) ApiCreateDigitalAccessKeyRequest {
	return ApiCreateDigitalAccessKeyRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateDigitalAccessKeyResponse
func (a *MessagingAPIService) CreateDigitalAccessKeyExecute(r ApiCreateDigitalAccessKeyRequest) (*CreateDigitalAccessKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateDigitalAccessKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateDigitalAccessKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateLegalDisclosureRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateLegalDisclosureRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateLegalDisclosureRequest) Body(body CreateLegalDisclosureRequest) ApiCreateLegalDisclosureRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateLegalDisclosureRequest) MarketplaceIds(marketplaceIds []string) ApiCreateLegalDisclosureRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateLegalDisclosureRequest) Execute() (*CreateLegalDisclosureResponse, *http.Response, error) {
	return r.ApiService.CreateLegalDisclosureExecute(r)
}

/*
CreateLegalDisclosure Method for CreateLegalDisclosure

Sends a critical message that contains documents that a seller is legally obligated to provide to the buyer. This message should only be used to deliver documents that are required by law.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateLegalDisclosureRequest
*/
func (a *MessagingAPIService) CreateLegalDisclosure(ctx context.Context, amazonOrderId string) ApiCreateLegalDisclosureRequest {
	return ApiCreateLegalDisclosureRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateLegalDisclosureResponse
func (a *MessagingAPIService) CreateLegalDisclosureExecute(r ApiCreateLegalDisclosureRequest) (*CreateLegalDisclosureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateLegalDisclosureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateLegalDisclosure")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateNegativeFeedbackRemovalRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	amazonOrderId  string
	marketplaceIds *[]string
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateNegativeFeedbackRemovalRequest) MarketplaceIds(marketplaceIds []string) ApiCreateNegativeFeedbackRemovalRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateNegativeFeedbackRemovalRequest) Execute() (*CreateNegativeFeedbackRemovalResponse, *http.Response, error) {
	return r.ApiService.CreateNegativeFeedbackRemovalExecute(r)
}

/*
CreateNegativeFeedbackRemoval Method for CreateNegativeFeedbackRemoval

Sends a non-critical message that asks a buyer to remove their negative feedback. This message should only be sent after the seller has resolved the buyer's problem.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateNegativeFeedbackRemovalRequest
*/
func (a *MessagingAPIService) CreateNegativeFeedbackRemoval(ctx context.Context, amazonOrderId string) ApiCreateNegativeFeedbackRemovalRequest {
	return ApiCreateNegativeFeedbackRemovalRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateNegativeFeedbackRemovalResponse
func (a *MessagingAPIService) CreateNegativeFeedbackRemovalExecute(r ApiCreateNegativeFeedbackRemovalRequest) (*CreateNegativeFeedbackRemovalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateNegativeFeedbackRemovalResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateNegativeFeedbackRemoval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateUnexpectedProblemRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateUnexpectedProblemRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateUnexpectedProblemRequest) Body(body CreateUnexpectedProblemRequest) ApiCreateUnexpectedProblemRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateUnexpectedProblemRequest) MarketplaceIds(marketplaceIds []string) ApiCreateUnexpectedProblemRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateUnexpectedProblemRequest) Execute() (*CreateUnexpectedProblemResponse, *http.Response, error) {
	return r.ApiService.CreateUnexpectedProblemExecute(r)
}

/*
CreateUnexpectedProblem Method for CreateUnexpectedProblem

Sends a critical message to a buyer that an unexpected problem was encountered affecting the completion of the order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateUnexpectedProblemRequest
*/
func (a *MessagingAPIService) CreateUnexpectedProblem(ctx context.Context, amazonOrderId string) ApiCreateUnexpectedProblemRequest {
	return ApiCreateUnexpectedProblemRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateUnexpectedProblemResponse
func (a *MessagingAPIService) CreateUnexpectedProblemExecute(r ApiCreateUnexpectedProblemRequest) (*CreateUnexpectedProblemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateUnexpectedProblemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateUnexpectedProblem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiCreateWarrantyRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *CreateWarrantyRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiCreateWarrantyRequest) Body(body CreateWarrantyRequest) ApiCreateWarrantyRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiCreateWarrantyRequest) MarketplaceIds(marketplaceIds []string) ApiCreateWarrantyRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateWarrantyRequest) Execute() (*CreateWarrantyResponse, *http.Response, error) {
	return r.ApiService.CreateWarrantyExecute(r)
}

/*
CreateWarranty Method for CreateWarranty

Sends a message to a buyer to provide details about warranty information on a purchase in their order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiCreateWarrantyRequest
*/
func (a *MessagingAPIService) CreateWarranty(ctx context.Context, amazonOrderId string) ApiCreateWarrantyRequest {
	return ApiCreateWarrantyRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateWarrantyResponse
func (a *MessagingAPIService) CreateWarrantyExecute(r ApiCreateWarrantyRequest) (*CreateWarrantyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateWarrantyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.CreateWarranty")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/warranty"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiGetAttributesRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	amazonOrderId  string
	marketplaceIds *[]string
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiGetAttributesRequest) MarketplaceIds(marketplaceIds []string) ApiGetAttributesRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiGetAttributesRequest) Execute() (*GetAttributesResponse, *http.Response, error) {
	return r.ApiService.GetAttributesExecute(r)
}

/*
GetAttributes Method for GetAttributes

Returns a response containing attributes related to an order. This includes buyer preferences.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiGetAttributesRequest
*/
func (a *MessagingAPIService) GetAttributes(ctx context.Context, amazonOrderId string) ApiGetAttributesRequest {
	return ApiGetAttributesRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return GetAttributesResponse
func (a *MessagingAPIService) GetAttributesExecute(r ApiGetAttributesRequest) (*GetAttributesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAttributesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.GetAttributes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiGetMessagingActionsForOrderRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	amazonOrderId  string
	marketplaceIds *[]string
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiGetMessagingActionsForOrderRequest) MarketplaceIds(marketplaceIds []string) ApiGetMessagingActionsForOrderRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiGetMessagingActionsForOrderRequest) Execute() (*GetMessagingActionsForOrderResponse, *http.Response, error) {
	return r.ApiService.GetMessagingActionsForOrderExecute(r)
}

/*
GetMessagingActionsForOrder Method for GetMessagingActionsForOrder

Returns a list of message types that are available for an order that you specify. A message type is represented by an actions object, which contains a path and query parameter(s). You can use the path and parameter(s) to call an operation that sends a message.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header contains the usage plan rate limits for the operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This specifies the order for which you want a list of available message types.
	@return ApiGetMessagingActionsForOrderRequest
*/
func (a *MessagingAPIService) GetMessagingActionsForOrder(ctx context.Context, amazonOrderId string) ApiGetMessagingActionsForOrderRequest {
	return ApiGetMessagingActionsForOrderRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return GetMessagingActionsForOrderResponse
func (a *MessagingAPIService) GetMessagingActionsForOrderExecute(r ApiGetMessagingActionsForOrderRequest) (*GetMessagingActionsForOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMessagingActionsForOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.GetMessagingActionsForOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiSendInvoiceRequest struct {
	ctx            context.Context
	ApiService     *MessagingAPIService
	body           *InvoiceRequest
	amazonOrderId  string
	marketplaceIds *[]string
}

// This contains the message body for a message.
func (r ApiSendInvoiceRequest) Body(body InvoiceRequest) ApiSendInvoiceRequest {
	r.body = &body
	return r
}

// A marketplace identifier. This identifies the marketplace in which the order was placed. You can only specify one marketplace.
func (r ApiSendInvoiceRequest) MarketplaceIds(marketplaceIds []string) ApiSendInvoiceRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiSendInvoiceRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.SendInvoiceExecute(r)
}

/*
SendInvoice Method for SendInvoice

Sends a message providing the buyer an invoice

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This identifies the order for which a message is sent.
	@return ApiSendInvoiceRequest
*/
func (a *MessagingAPIService) SendInvoice(ctx context.Context, amazonOrderId string) ApiSendInvoiceRequest {
	return ApiSendInvoiceRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return InvoiceResponse
func (a *MessagingAPIService) SendInvoiceExecute(r ApiSendInvoiceRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagingAPIService.SendInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messaging/v1/orders/{amazonOrderId}/messages/invoice"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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
