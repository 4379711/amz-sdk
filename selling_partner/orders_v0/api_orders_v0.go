package orders_v0

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// OrdersV0APIService OrdersV0API service
type OrdersV0APIService service

type ApiConfirmShipmentRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	payload    *ConfirmShipmentRequest
	orderId    string
}

// Request body of &#x60;confirmShipment&#x60;.
func (r ApiConfirmShipmentRequest) Payload(payload ConfirmShipmentRequest) ApiConfirmShipmentRequest {
	r.payload = &payload
	return r
}

func (r ApiConfirmShipmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.ConfirmShipmentExecute(r)
}

/*
ConfirmShipment Method for ConfirmShipment

Updates the shipment confirmation status for a specified order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiConfirmShipmentRequest
*/
func (a *OrdersV0APIService) ConfirmShipment(ctx context.Context, orderId string) ApiConfirmShipmentRequest {
	return ApiConfirmShipmentRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *OrdersV0APIService) ConfirmShipmentExecute(r ApiConfirmShipmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.ConfirmShipment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/shipmentConfirmation"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return nil, reportError("payload is required and must be specified")
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
	localVarPostBody = r.payload
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

type ApiGetOrderRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	orderId    string
}

func (r ApiGetOrderRequest) Execute() (*GetOrderResponse, *http.Response, error) {
	return r.ApiService.GetOrderExecute(r)
}

/*
GetOrder Method for GetOrder

Returns the order that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiGetOrderRequest
*/
func (a *OrdersV0APIService) GetOrder(ctx context.Context, orderId string) ApiGetOrderRequest {
	return ApiGetOrderRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return GetOrderResponse
func (a *OrdersV0APIService) GetOrderExecute(r ApiGetOrderRequest) (*GetOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.GetOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiGetOrderAddressRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	orderId    string
}

func (r ApiGetOrderAddressRequest) Execute() (*GetOrderAddressResponse, *http.Response, error) {
	return r.ApiService.GetOrderAddressExecute(r)
}

/*
GetOrderAddress Method for GetOrderAddress

Returns the shipping address for the order that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An `orderId` is an Amazon-defined order identifier, in 3-7-7 format.
	@return ApiGetOrderAddressRequest
*/
func (a *OrdersV0APIService) GetOrderAddress(ctx context.Context, orderId string) ApiGetOrderAddressRequest {
	return ApiGetOrderAddressRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return GetOrderAddressResponse
func (a *OrdersV0APIService) GetOrderAddressExecute(r ApiGetOrderAddressRequest) (*GetOrderAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.GetOrderAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/address"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiGetOrderBuyerInfoRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	orderId    string
}

func (r ApiGetOrderBuyerInfoRequest) Execute() (*GetOrderBuyerInfoResponse, *http.Response, error) {
	return r.ApiService.GetOrderBuyerInfoExecute(r)
}

/*
GetOrderBuyerInfo Method for GetOrderBuyerInfo

Returns buyer information for the order that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An `orderId` is an Amazon-defined order identifier, in 3-7-7 format.
	@return ApiGetOrderBuyerInfoRequest
*/
func (a *OrdersV0APIService) GetOrderBuyerInfo(ctx context.Context, orderId string) ApiGetOrderBuyerInfoRequest {
	return ApiGetOrderBuyerInfoRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return GetOrderBuyerInfoResponse
func (a *OrdersV0APIService) GetOrderBuyerInfoExecute(r ApiGetOrderBuyerInfoRequest) (*GetOrderBuyerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderBuyerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.GetOrderBuyerInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/buyerInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiGetOrderItemsRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	orderId    string
	nextToken  *string
}

// A string token returned in the response of your previous request.
func (r ApiGetOrderItemsRequest) NextToken(nextToken string) ApiGetOrderItemsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetOrderItemsRequest) Execute() (*GetOrderItemsResponse, *http.Response, error) {
	return r.ApiService.GetOrderItemsExecute(r)
}

/*
GetOrderItems Method for GetOrderItems

Returns detailed order item information for the order that you specify. If `NextToken` is provided, it's used to retrieve the next page of order items.

__Note__: When an order is in the Pending state (the order has been placed but payment has not been authorized), the getOrderItems operation does not return information about pricing, taxes, shipping charges, gift status or promotions for the order items in the order. After an order leaves the Pending state (this occurs when payment has been authorized) and enters the Unshipped, Partially Shipped, or Shipped state, the getOrderItems operation returns information about pricing, taxes, shipping charges, gift status and promotions for the order items in the order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiGetOrderItemsRequest
*/
func (a *OrdersV0APIService) GetOrderItems(ctx context.Context, orderId string) ApiGetOrderItemsRequest {
	return ApiGetOrderItemsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return GetOrderItemsResponse
func (a *OrdersV0APIService) GetOrderItemsExecute(r ApiGetOrderItemsRequest) (*GetOrderItemsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderItemsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.GetOrderItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/orderItems"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "NextToken", r.nextToken, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetOrderItemsBuyerInfoRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	orderId    string
	nextToken  *string
}

// A string token returned in the response of your previous request.
func (r ApiGetOrderItemsBuyerInfoRequest) NextToken(nextToken string) ApiGetOrderItemsBuyerInfoRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetOrderItemsBuyerInfoRequest) Execute() (*GetOrderItemsBuyerInfoResponse, *http.Response, error) {
	return r.ApiService.GetOrderItemsBuyerInfoExecute(r)
}

/*
GetOrderItemsBuyerInfo Method for GetOrderItemsBuyerInfo

Returns buyer information for the order items in the order that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiGetOrderItemsBuyerInfoRequest
*/
func (a *OrdersV0APIService) GetOrderItemsBuyerInfo(ctx context.Context, orderId string) ApiGetOrderItemsBuyerInfoRequest {
	return ApiGetOrderItemsBuyerInfoRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return GetOrderItemsBuyerInfoResponse
func (a *OrdersV0APIService) GetOrderItemsBuyerInfoExecute(r ApiGetOrderItemsBuyerInfoRequest) (*GetOrderItemsBuyerInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderItemsBuyerInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.GetOrderItemsBuyerInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/orderItems/buyerInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "NextToken", r.nextToken, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetOrderRegulatedInfoRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	orderId    string
}

func (r ApiGetOrderRegulatedInfoRequest) Execute() (*GetOrderRegulatedInfoResponse, *http.Response, error) {
	return r.ApiService.GetOrderRegulatedInfoExecute(r)
}

/*
GetOrderRegulatedInfo Method for GetOrderRegulatedInfo

Returns regulated information for the order that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiGetOrderRegulatedInfoRequest
*/
func (a *OrdersV0APIService) GetOrderRegulatedInfo(ctx context.Context, orderId string) ApiGetOrderRegulatedInfoRequest {
	return ApiGetOrderRegulatedInfoRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return GetOrderRegulatedInfoResponse
func (a *OrdersV0APIService) GetOrderRegulatedInfoExecute(r ApiGetOrderRegulatedInfoRequest) (*GetOrderRegulatedInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderRegulatedInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.GetOrderRegulatedInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/regulatedInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "PendingOrder", "ApprovedOrder", "RejectedOrder"}

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

type ApiGetOrdersRequest struct {
	ctx                             context.Context
	ApiService                      *OrdersV0APIService
	marketplaceIds                  *[]string
	createdAfter                    *string
	createdBefore                   *string
	lastUpdatedAfter                *string
	lastUpdatedBefore               *string
	orderStatuses                   *[]string
	fulfillmentChannels             *[]string
	paymentMethods                  *[]string
	buyerEmail                      *string
	sellerOrderId                   *string
	maxResultsPerPage               *int32
	easyShipShipmentStatuses        *[]string
	electronicInvoiceStatuses       *[]string
	nextToken                       *string
	amazonOrderIds                  *[]string
	actualFulfillmentSupplySourceId *string
	isISPU                          *bool
	storeChainStoreId               *string
	earliestDeliveryDateBefore      *string
	earliestDeliveryDateAfter       *string
	latestDeliveryDateBefore        *string
	latestDeliveryDateAfter         *string
}

// A list of &#x60;MarketplaceId&#x60; values. Used to select orders that were placed in the specified marketplaces.  Refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids) for a complete list of &#x60;marketplaceId&#x60; values.
func (r ApiGetOrdersRequest) MarketplaceIds(marketplaceIds []string) ApiGetOrdersRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// Use this date to select orders created after (or at) a specified time. Only orders placed after the specified time are returned. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: Either the &#x60;CreatedAfter&#x60; parameter or the &#x60;LastUpdatedAfter&#x60; parameter is required. Both cannot be empty. &#x60;LastUpdatedAfter&#x60; and &#x60;LastUpdatedBefore&#x60; cannot be set when &#x60;CreatedAfter&#x60; is set.
func (r ApiGetOrdersRequest) CreatedAfter(createdAfter string) ApiGetOrdersRequest {
	r.createdAfter = &createdAfter
	return r
}

// Use this date to select orders created before (or at) a specified time. Only orders placed before the specified time are returned. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: &#x60;CreatedBefore&#x60; is optional when &#x60;CreatedAfter&#x60; is set. If specified, &#x60;CreatedBefore&#x60; must be equal to or after the &#x60;CreatedAfter&#x60; date and at least two minutes before current time.
func (r ApiGetOrdersRequest) CreatedBefore(createdBefore string) ApiGetOrdersRequest {
	r.createdBefore = &createdBefore
	return r
}

// Use this date to select orders that were last updated after (or at) a specified time. An update is defined as any change in order status, including the creation of a new order. Includes updates made by Amazon and by the seller. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: Either the &#x60;CreatedAfter&#x60; parameter or the &#x60;LastUpdatedAfter&#x60; parameter is required. Both cannot be empty. &#x60;CreatedAfter&#x60; or &#x60;CreatedBefore&#x60; cannot be set when &#x60;LastUpdatedAfter&#x60; is set.
func (r ApiGetOrdersRequest) LastUpdatedAfter(lastUpdatedAfter string) ApiGetOrdersRequest {
	r.lastUpdatedAfter = &lastUpdatedAfter
	return r
}

// Use this date to select orders that were last updated before (or at) a specified time. An update is defined as any change in order status, including the creation of a new order. Includes updates made by Amazon and by the seller. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: &#x60;LastUpdatedBefore&#x60; is optional when &#x60;LastUpdatedAfter&#x60; is set. But if specified, &#x60;LastUpdatedBefore&#x60; must be equal to or after the &#x60;LastUpdatedAfter&#x60; date and at least two minutes before current time.
func (r ApiGetOrdersRequest) LastUpdatedBefore(lastUpdatedBefore string) ApiGetOrdersRequest {
	r.lastUpdatedBefore = &lastUpdatedBefore
	return r
}

// A list of &#x60;OrderStatus&#x60; values used to filter the results.  **Possible values:** - &#x60;PendingAvailability&#x60; (This status is available for pre-orders only. The order has been placed, payment has not been authorized, and the release date of the item is in the future.) - &#x60;Pending&#x60; (The order has been placed but payment has not been authorized.) - &#x60;Unshipped&#x60; (Payment has been authorized and the order is ready for shipment, but no items in the order have been shipped.) - &#x60;PartiallyShipped&#x60; (One or more, but not all, items in the order have been shipped.) - &#x60;Shipped&#x60; (All items in the order have been shipped.) - &#x60;InvoiceUnconfirmed&#x60; (All items in the order have been shipped. The seller has not yet given confirmation to Amazon that the invoice has been shipped to the buyer.) - &#x60;Canceled&#x60; (The order has been canceled.) - &#x60;Unfulfillable&#x60; (The order cannot be fulfilled. This state applies only to Multi-Channel Fulfillment orders.)
func (r ApiGetOrdersRequest) OrderStatuses(orderStatuses []string) ApiGetOrdersRequest {
	r.orderStatuses = &orderStatuses
	return r
}

// A list that indicates how an order was fulfilled. Filters the results by fulfillment channel.   **Possible values**: &#x60;AFN&#x60; (fulfilled by Amazon), &#x60;MFN&#x60; (fulfilled by seller).
func (r ApiGetOrdersRequest) FulfillmentChannels(fulfillmentChannels []string) ApiGetOrdersRequest {
	r.fulfillmentChannels = &fulfillmentChannels
	return r
}

// A list of payment method values. Use this field to select orders that were paid with the specified payment methods.  **Possible values**: &#x60;COD&#x60; (cash on delivery), &#x60;CVS&#x60; (convenience store), &#x60;Other&#x60; (Any payment method other than COD or CVS).
func (r ApiGetOrdersRequest) PaymentMethods(paymentMethods []string) ApiGetOrdersRequest {
	r.paymentMethods = &paymentMethods
	return r
}

// The email address of a buyer. Used to select orders that contain the specified email address.
func (r ApiGetOrdersRequest) BuyerEmail(buyerEmail string) ApiGetOrdersRequest {
	r.buyerEmail = &buyerEmail
	return r
}

// An order identifier that is specified by the seller. Used to select only the orders that match the order identifier. If &#x60;SellerOrderId&#x60; is specified, then &#x60;FulfillmentChannels&#x60;, &#x60;OrderStatuses&#x60;, &#x60;PaymentMethod&#x60;, &#x60;LastUpdatedAfter&#x60;, LastUpdatedBefore, and &#x60;BuyerEmail&#x60; cannot be specified.
func (r ApiGetOrdersRequest) SellerOrderId(sellerOrderId string) ApiGetOrdersRequest {
	r.sellerOrderId = &sellerOrderId
	return r
}

// A number that indicates the maximum number of orders that can be returned per page. Value must be 1 - 100. Default 100.
func (r ApiGetOrdersRequest) MaxResultsPerPage(maxResultsPerPage int32) ApiGetOrdersRequest {
	r.maxResultsPerPage = &maxResultsPerPage
	return r
}

// A list of &#x60;EasyShipShipmentStatus&#x60; values. Used to select Easy Ship orders with statuses that match the specified values. If &#x60;EasyShipShipmentStatus&#x60; is specified, only Amazon Easy Ship orders are returned.  **Possible values:** - &#x60;PendingSchedule&#x60; (The package is awaiting the schedule for pick-up.) - &#x60;PendingPickUp&#x60; (Amazon has not yet picked up the package from the seller.) - &#x60;PendingDropOff&#x60; (The seller will deliver the package to the carrier.) - &#x60;LabelCanceled&#x60; (The seller canceled the pickup.) - &#x60;PickedUp&#x60; (Amazon has picked up the package from the seller.) - &#x60;DroppedOff&#x60; (The package is delivered to the carrier by the seller.) - &#x60;AtOriginFC&#x60; (The packaged is at the origin fulfillment center.) - &#x60;AtDestinationFC&#x60; (The package is at the destination fulfillment center.) - &#x60;Delivered&#x60; (The package has been delivered.) - &#x60;RejectedByBuyer&#x60; (The package has been rejected by the buyer.) - &#x60;Undeliverable&#x60; (The package cannot be delivered.) - &#x60;ReturningToSeller&#x60; (The package was not delivered and is being returned to the seller.) - &#x60;ReturnedToSeller&#x60; (The package was not delivered and was returned to the seller.) - &#x60;Lost&#x60; (The package is lost.) - &#x60;OutForDelivery&#x60; (The package is out for delivery.) - &#x60;Damaged&#x60; (The package was damaged by the carrier.)
func (r ApiGetOrdersRequest) EasyShipShipmentStatuses(easyShipShipmentStatuses []string) ApiGetOrdersRequest {
	r.easyShipShipmentStatuses = &easyShipShipmentStatuses
	return r
}

// A list of &#x60;ElectronicInvoiceStatus&#x60; values. Used to select orders with electronic invoice statuses that match the specified values.  **Possible values:** - &#x60;NotRequired&#x60; (Electronic invoice submission is not required for this order.) - &#x60;NotFound&#x60; (The electronic invoice was not submitted for this order.) - &#x60;Processing&#x60; (The electronic invoice is being processed for this order.) - &#x60;Errored&#x60; (The last submitted electronic invoice was rejected for this order.) - &#x60;Accepted&#x60; (The last submitted electronic invoice was submitted and accepted.)
func (r ApiGetOrdersRequest) ElectronicInvoiceStatuses(electronicInvoiceStatuses []string) ApiGetOrdersRequest {
	r.electronicInvoiceStatuses = &electronicInvoiceStatuses
	return r
}

// A string token returned in the response of your previous request.
func (r ApiGetOrdersRequest) NextToken(nextToken string) ApiGetOrdersRequest {
	r.nextToken = &nextToken
	return r
}

// A list of &#x60;AmazonOrderId&#x60; values. An &#x60;AmazonOrderId&#x60; is an Amazon-defined order identifier, in 3-7-7 format.
func (r ApiGetOrdersRequest) AmazonOrderIds(amazonOrderIds []string) ApiGetOrdersRequest {
	r.amazonOrderIds = &amazonOrderIds
	return r
}

// The &#x60;sourceId&#x60; of the location from where you want the order fulfilled.
func (r ApiGetOrdersRequest) ActualFulfillmentSupplySourceId(actualFulfillmentSupplySourceId string) ApiGetOrdersRequest {
	r.actualFulfillmentSupplySourceId = &actualFulfillmentSupplySourceId
	return r
}

// When true, this order is marked to be picked up from a store rather than delivered.
func (r ApiGetOrdersRequest) IsISPU(isISPU bool) ApiGetOrdersRequest {
	r.isISPU = &isISPU
	return r
}

// The store chain store identifier. Linked to a specific store in a store chain.
func (r ApiGetOrdersRequest) StoreChainStoreId(storeChainStoreId string) ApiGetOrdersRequest {
	r.storeChainStoreId = &storeChainStoreId
	return r
}

// Use this date to select orders with a earliest delivery date before (or at) a specified time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
func (r ApiGetOrdersRequest) EarliestDeliveryDateBefore(earliestDeliveryDateBefore string) ApiGetOrdersRequest {
	r.earliestDeliveryDateBefore = &earliestDeliveryDateBefore
	return r
}

// Use this date to select orders with a earliest delivery date after (or at) a specified time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
func (r ApiGetOrdersRequest) EarliestDeliveryDateAfter(earliestDeliveryDateAfter string) ApiGetOrdersRequest {
	r.earliestDeliveryDateAfter = &earliestDeliveryDateAfter
	return r
}

// Use this date to select orders with a latest delivery date before (or at) a specified time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
func (r ApiGetOrdersRequest) LatestDeliveryDateBefore(latestDeliveryDateBefore string) ApiGetOrdersRequest {
	r.latestDeliveryDateBefore = &latestDeliveryDateBefore
	return r
}

// Use this date to select orders with a latest delivery date after (or at) a specified time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
func (r ApiGetOrdersRequest) LatestDeliveryDateAfter(latestDeliveryDateAfter string) ApiGetOrdersRequest {
	r.latestDeliveryDateAfter = &latestDeliveryDateAfter
	return r
}

func (r ApiGetOrdersRequest) Execute() (*GetOrdersResponse, *http.Response, error) {
	return r.ApiService.GetOrdersExecute(r)
}

/*
GetOrders Method for GetOrders

Returns orders that are created or updated during the specified time period. If you want to return specific types of orders, you can apply filters to your request. `NextToken` doesn't affect any filters that you include in your request; it only impacts the pagination for the filtered orders response.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0167 | 20 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api) in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetOrdersRequest
*/
func (a *OrdersV0APIService) GetOrders(ctx context.Context) ApiGetOrdersRequest {
	return ApiGetOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOrdersResponse
func (a *OrdersV0APIService) GetOrdersExecute(r ApiGetOrdersRequest) (*GetOrdersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrdersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.GetOrders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 50 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 50 elements")
	}

	if r.createdAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "CreatedAfter", r.createdAfter, "", "")
	}
	if r.createdBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "CreatedBefore", r.createdBefore, "", "")
	}
	if r.lastUpdatedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "LastUpdatedAfter", r.lastUpdatedAfter, "", "")
	}
	if r.lastUpdatedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "LastUpdatedBefore", r.lastUpdatedBefore, "", "")
	}
	if r.orderStatuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrderStatuses", r.orderStatuses, "form", "csv")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "MarketplaceIds", r.marketplaceIds, "form", "csv")
	if r.fulfillmentChannels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "FulfillmentChannels", r.fulfillmentChannels, "form", "csv")
	}
	if r.paymentMethods != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PaymentMethods", r.paymentMethods, "form", "csv")
	}
	if r.buyerEmail != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "BuyerEmail", r.buyerEmail, "", "")
	}
	if r.sellerOrderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SellerOrderId", r.sellerOrderId, "", "")
	}
	if r.maxResultsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxResultsPerPage", r.maxResultsPerPage, "", "")
	}
	if r.easyShipShipmentStatuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EasyShipShipmentStatuses", r.easyShipShipmentStatuses, "form", "csv")
	}
	if r.electronicInvoiceStatuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ElectronicInvoiceStatuses", r.electronicInvoiceStatuses, "form", "csv")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "NextToken", r.nextToken, "", "")
	}
	if r.amazonOrderIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "AmazonOrderIds", r.amazonOrderIds, "form", "csv")
	}
	if r.actualFulfillmentSupplySourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ActualFulfillmentSupplySourceId", r.actualFulfillmentSupplySourceId, "", "")
	}
	if r.isISPU != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsISPU", r.isISPU, "", "")
	}
	if r.storeChainStoreId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StoreChainStoreId", r.storeChainStoreId, "", "")
	}
	if r.earliestDeliveryDateBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EarliestDeliveryDateBefore", r.earliestDeliveryDateBefore, "", "")
	}
	if r.earliestDeliveryDateAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EarliestDeliveryDateAfter", r.earliestDeliveryDateAfter, "", "")
	}
	if r.latestDeliveryDateBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "LatestDeliveryDateBefore", r.latestDeliveryDateBefore, "", "")
	}
	if r.latestDeliveryDateAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "LatestDeliveryDateAfter", r.latestDeliveryDateAfter, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiUpdateVerificationStatusRequest struct {
	ctx        context.Context
	ApiService *OrdersV0APIService
	payload    *UpdateVerificationStatusRequest
	orderId    string
}

// The request body for the &#x60;updateVerificationStatus&#x60; operation.
func (r ApiUpdateVerificationStatusRequest) Payload(payload UpdateVerificationStatusRequest) ApiUpdateVerificationStatusRequest {
	r.payload = &payload
	return r
}

func (r ApiUpdateVerificationStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateVerificationStatusExecute(r)
}

/*
UpdateVerificationStatus Method for UpdateVerificationStatus

Updates (approves or rejects) the verification status of an order containing regulated products.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiUpdateVerificationStatusRequest
*/
func (a *OrdersV0APIService) UpdateVerificationStatus(ctx context.Context, orderId string) ApiUpdateVerificationStatusRequest {
	return ApiUpdateVerificationStatusRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *OrdersV0APIService) UpdateVerificationStatusExecute(r ApiUpdateVerificationStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV0APIService.UpdateVerificationStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/regulatedInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return nil, reportError("payload is required and must be specified")
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
	localVarPostBody = r.payload
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
