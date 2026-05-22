package awd_20240509

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// AwdAPIService AwdAPI service
type AwdAPIService service

type ApiCancelInboundRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	orderId    string
}

func (r ApiCancelInboundRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelInboundExecute(r)
}

/*
CancelInbound Method for CancelInbound

Cancels an AWD Inbound order and its associated shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The ID of the inbound order you want to cancel.
	@return ApiCancelInboundRequest
*/
func (a *AwdAPIService) CancelInbound(ctx context.Context, orderId string) ApiCancelInboundRequest {
	return ApiCancelInboundRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *AwdAPIService) CancelInboundExecute(r ApiCancelInboundRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.CancelInbound")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundOrders/{orderId}/cancellation"
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

type ApiCheckInboundEligibilityRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	body       *InboundPackages
}

// Represents the packages you want to inbound.
func (r ApiCheckInboundEligibilityRequest) Body(body InboundPackages) ApiCheckInboundEligibilityRequest {
	r.body = &body
	return r
}

func (r ApiCheckInboundEligibilityRequest) Execute() (*InboundEligibility, *http.Response, error) {
	return r.ApiService.CheckInboundEligibilityExecute(r)
}

/*
CheckInboundEligibility Method for CheckInboundEligibility

Determines if the packages you specify are eligible for an AWD inbound order and contains error details for ineligible packages.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCheckInboundEligibilityRequest
*/
func (a *AwdAPIService) CheckInboundEligibility(ctx context.Context) ApiCheckInboundEligibilityRequest {
	return ApiCheckInboundEligibilityRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InboundEligibility
func (a *AwdAPIService) CheckInboundEligibilityExecute(r ApiCheckInboundEligibilityRequest) (*InboundEligibility, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InboundEligibility
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.CheckInboundEligibility")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundEligibility"

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

type ApiConfirmInboundRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	orderId    string
}

func (r ApiConfirmInboundRequest) Execute() (*http.Response, error) {
	return r.ApiService.ConfirmInboundExecute(r)
}

/*
ConfirmInbound Method for ConfirmInbound

Confirms an AWD inbound order in `DRAFT` status.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The ID of the inbound order that you want to confirm.
	@return ApiConfirmInboundRequest
*/
func (a *AwdAPIService) ConfirmInbound(ctx context.Context, orderId string) ApiConfirmInboundRequest {
	return ApiConfirmInboundRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *AwdAPIService) ConfirmInboundExecute(r ApiConfirmInboundRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.ConfirmInbound")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundOrders/{orderId}/confirmation"
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

type ApiCreateInboundRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	body       *InboundOrderCreationData
}

// Payload for creating an inbound order.
func (r ApiCreateInboundRequest) Body(body InboundOrderCreationData) ApiCreateInboundRequest {
	r.body = &body
	return r
}

func (r ApiCreateInboundRequest) Execute() (*InboundOrderReference, *http.Response, error) {
	return r.ApiService.CreateInboundExecute(r)
}

/*
CreateInbound Method for CreateInbound

Creates a draft AWD inbound order with a list of packages for inbound shipment. The operation creates one shipment per order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateInboundRequest
*/
func (a *AwdAPIService) CreateInbound(ctx context.Context) ApiCreateInboundRequest {
	return ApiCreateInboundRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InboundOrderReference
func (a *AwdAPIService) CreateInboundExecute(r ApiCreateInboundRequest) (*InboundOrderReference, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InboundOrderReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.CreateInbound")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundOrders"

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

type ApiGetInboundRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	orderId    string
}

func (r ApiGetInboundRequest) Execute() (*InboundOrder, *http.Response, error) {
	return r.ApiService.GetInboundExecute(r)
}

/*
GetInbound Method for GetInbound

Retrieves an AWD inbound order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The ID of the inbound order that you want to retrieve.
	@return ApiGetInboundRequest
*/
func (a *AwdAPIService) GetInbound(ctx context.Context, orderId string) ApiGetInboundRequest {
	return ApiGetInboundRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return InboundOrder
func (a *AwdAPIService) GetInboundExecute(r ApiGetInboundRequest) (*InboundOrder, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InboundOrder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.GetInbound")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundOrders/{orderId}"
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

type ApiGetInboundShipmentRequest struct {
	ctx           context.Context
	ApiService    *AwdAPIService
	shipmentId    string
	skuQuantities *string
}

// If equal to &#x60;SHOW&#x60;, the response includes the shipment SKU quantity details.  Defaults to &#x60;HIDE&#x60;, in which case the response does not contain SKU quantities
func (r ApiGetInboundShipmentRequest) SkuQuantities(skuQuantities string) ApiGetInboundShipmentRequest {
	r.skuQuantities = &skuQuantities
	return r
}

func (r ApiGetInboundShipmentRequest) Execute() (*InboundShipment, *http.Response, error) {
	return r.ApiService.GetInboundShipmentExecute(r)
}

/*
GetInboundShipment Method for GetInboundShipment

Retrieves an AWD inbound shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId ID for the shipment. A shipment contains the cases being inbounded.
	@return ApiGetInboundShipmentRequest
*/
func (a *AwdAPIService) GetInboundShipment(ctx context.Context, shipmentId string) ApiGetInboundShipmentRequest {
	return ApiGetInboundShipmentRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return InboundShipment
func (a *AwdAPIService) GetInboundShipmentExecute(r ApiGetInboundShipmentRequest) (*InboundShipment, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InboundShipment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.GetInboundShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundShipments/{shipmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.shipmentId) < 1 {
		return localVarReturnValue, nil, reportError("shipmentId must have at least 1 elements")
	}

	if r.skuQuantities != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skuQuantities", r.skuQuantities, "", "")
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

type ApiGetInboundShipmentLabelsRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	shipmentId string
	pageType   *string
	formatType *string
}

// Page type for the generated labels. The default is &#x60;PLAIN_PAPER&#x60;.
func (r ApiGetInboundShipmentLabelsRequest) PageType(pageType string) ApiGetInboundShipmentLabelsRequest {
	r.pageType = &pageType
	return r
}

// The format type of the output file that contains your labels. The default format type is &#x60;PDF&#x60;.
func (r ApiGetInboundShipmentLabelsRequest) FormatType(formatType string) ApiGetInboundShipmentLabelsRequest {
	r.formatType = &formatType
	return r
}

func (r ApiGetInboundShipmentLabelsRequest) Execute() (*ShipmentLabels, *http.Response, error) {
	return r.ApiService.GetInboundShipmentLabelsExecute(r)
}

/*
GetInboundShipmentLabels Method for GetInboundShipmentLabels

Retrieves the box labels for a shipment ID that you specify. This is an asynchronous operation. If the label status is `GENERATED`, then the label URL is available.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId ID for the shipment.
	@return ApiGetInboundShipmentLabelsRequest
*/
func (a *AwdAPIService) GetInboundShipmentLabels(ctx context.Context, shipmentId string) ApiGetInboundShipmentLabelsRequest {
	return ApiGetInboundShipmentLabelsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return ShipmentLabels
func (a *AwdAPIService) GetInboundShipmentLabelsExecute(r ApiGetInboundShipmentLabelsRequest) (*ShipmentLabels, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShipmentLabels
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.GetInboundShipmentLabels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundShipments/{shipmentId}/labels"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.shipmentId) < 1 {
		return localVarReturnValue, nil, reportError("shipmentId must have at least 1 elements")
	}

	if r.pageType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageType", r.pageType, "", "")
	}
	if r.formatType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "formatType", r.formatType, "", "")
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

type ApiListInboundShipmentsRequest struct {
	ctx            context.Context
	ApiService     *AwdAPIService
	sortBy         *string
	sortOrder      *string
	shipmentStatus *string
	updatedAfter   *time.Time
	updatedBefore  *time.Time
	maxResults     *int32
	nextToken      *string
}

// Field to sort results by. By default, the response will be sorted by UPDATED_AT.
func (r ApiListInboundShipmentsRequest) SortBy(sortBy string) ApiListInboundShipmentsRequest {
	r.sortBy = &sortBy
	return r
}

// Sort the response in ASCENDING or DESCENDING order. By default, the response will be sorted in DESCENDING order.
func (r ApiListInboundShipmentsRequest) SortOrder(sortOrder string) ApiListInboundShipmentsRequest {
	r.sortOrder = &sortOrder
	return r
}

// Filter by inbound shipment status.
func (r ApiListInboundShipmentsRequest) ShipmentStatus(shipmentStatus string) ApiListInboundShipmentsRequest {
	r.shipmentStatus = &shipmentStatus
	return r
}

// List the inbound shipments that were updated after a certain time (inclusive). The date must be in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; format.
func (r ApiListInboundShipmentsRequest) UpdatedAfter(updatedAfter time.Time) ApiListInboundShipmentsRequest {
	r.updatedAfter = &updatedAfter
	return r
}

// List the inbound shipments that were updated before a certain time (inclusive). The date must be in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; format.
func (r ApiListInboundShipmentsRequest) UpdatedBefore(updatedBefore time.Time) ApiListInboundShipmentsRequest {
	r.updatedBefore = &updatedBefore
	return r
}

// Maximum number of results to return.
func (r ApiListInboundShipmentsRequest) MaxResults(maxResults int32) ApiListInboundShipmentsRequest {
	r.maxResults = &maxResults
	return r
}

// A token that is used to retrieve the next page of results. The response includes &#x60;nextToken&#x60; when the number of results exceeds the specified &#x60;maxResults&#x60; value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until &#x60;nextToken&#x60; is null. Note that this operation can return empty pages.
func (r ApiListInboundShipmentsRequest) NextToken(nextToken string) ApiListInboundShipmentsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListInboundShipmentsRequest) Execute() (*ShipmentListing, *http.Response, error) {
	return r.ApiService.ListInboundShipmentsExecute(r)
}

/*
ListInboundShipments Method for ListInboundShipments

Retrieves a summary of all the inbound AWD shipments associated with a merchant, with the ability to apply optional filters.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListInboundShipmentsRequest
*/
func (a *AwdAPIService) ListInboundShipments(ctx context.Context) ApiListInboundShipmentsRequest {
	return ApiListInboundShipmentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ShipmentListing
func (a *AwdAPIService) ListInboundShipmentsExecute(r ApiListInboundShipmentsRequest) (*ShipmentListing, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShipmentListing
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.ListInboundShipments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundShipments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "", "")
	}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", r.sortOrder, "", "")
	}
	if r.shipmentStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shipmentStatus", r.shipmentStatus, "", "")
	}
	if r.updatedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updatedAfter", r.updatedAfter, "", "")
	}
	if r.updatedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updatedBefore", r.updatedBefore, "", "")
	}
	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxResults", r.maxResults, "", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "", "")
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

type ApiListInventoryRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	sku        *string
	sortOrder  *string
	details    *string
	nextToken  *string
	maxResults *int32
}

// Filter by seller or merchant SKU for the item.
func (r ApiListInventoryRequest) Sku(sku string) ApiListInventoryRequest {
	r.sku = &sku
	return r
}

// Sort the response in &#x60;ASCENDING&#x60; or &#x60;DESCENDING&#x60; order.
func (r ApiListInventoryRequest) SortOrder(sortOrder string) ApiListInventoryRequest {
	r.sortOrder = &sortOrder
	return r
}

// Set to &#x60;SHOW&#x60; to return summaries with additional inventory details. Defaults to &#x60;HIDE,&#x60; which returns only inventory summary totals.
func (r ApiListInventoryRequest) Details(details string) ApiListInventoryRequest {
	r.details = &details
	return r
}

// A token that is used to retrieve the next page of results. The response includes &#x60;nextToken&#x60; when the number of results exceeds the specified &#x60;maxResults&#x60; value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until &#x60;nextToken&#x60; is null. Note that this operation can return empty pages.
func (r ApiListInventoryRequest) NextToken(nextToken string) ApiListInventoryRequest {
	r.nextToken = &nextToken
	return r
}

// Maximum number of results to return.
func (r ApiListInventoryRequest) MaxResults(maxResults int32) ApiListInventoryRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiListInventoryRequest) Execute() (*InventoryListing, *http.Response, error) {
	return r.ApiService.ListInventoryExecute(r)
}

/*
ListInventory Method for ListInventory

Lists AWD inventory associated with a merchant with the ability to apply optional filters.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListInventoryRequest
*/
func (a *AwdAPIService) ListInventory(ctx context.Context) ApiListInventoryRequest {
	return ApiListInventoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InventoryListing
func (a *AwdAPIService) ListInventoryExecute(r ApiListInventoryRequest) (*InventoryListing, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InventoryListing
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.ListInventory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inventory"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sku != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sku", r.sku, "", "")
	}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", r.sortOrder, "", "")
	}
	if r.details != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "details", r.details, "", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "", "")
	}
	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxResults", r.maxResults, "", "")
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

type ApiUpdateInboundRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	body       *InboundOrder
	orderId    string
}

// Represents an AWD inbound order.
func (r ApiUpdateInboundRequest) Body(body InboundOrder) ApiUpdateInboundRequest {
	r.body = &body
	return r
}

func (r ApiUpdateInboundRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateInboundExecute(r)
}

/*
UpdateInbound Method for UpdateInbound

Updates an AWD inbound order that is in `DRAFT` status and not yet confirmed. Use this operation to update the `packagesToInbound`, `originAddress` and `preferences` attributes.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The ID of the inbound order that you want to update.
	@return ApiUpdateInboundRequest
*/
func (a *AwdAPIService) UpdateInbound(ctx context.Context, orderId string) ApiUpdateInboundRequest {
	return ApiUpdateInboundRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *AwdAPIService) UpdateInboundExecute(r ApiUpdateInboundRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.UpdateInbound")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundOrders/{orderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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

type ApiUpdateInboundShipmentTransportDetailsRequest struct {
	ctx        context.Context
	ApiService *AwdAPIService
	body       *TransportationDetails
	shipmentId string
}

// Transportation details for the shipment.
func (r ApiUpdateInboundShipmentTransportDetailsRequest) Body(body TransportationDetails) ApiUpdateInboundShipmentTransportDetailsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateInboundShipmentTransportDetailsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateInboundShipmentTransportDetailsExecute(r)
}

/*
UpdateInboundShipmentTransportDetails Method for UpdateInboundShipmentTransportDetails

Updates transport details for an AWD shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId The shipment ID.
	@return ApiUpdateInboundShipmentTransportDetailsRequest
*/
func (a *AwdAPIService) UpdateInboundShipmentTransportDetails(ctx context.Context, shipmentId string) ApiUpdateInboundShipmentTransportDetailsRequest {
	return ApiUpdateInboundShipmentTransportDetailsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *AwdAPIService) UpdateInboundShipmentTransportDetailsExecute(r ApiUpdateInboundShipmentTransportDetailsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AwdAPIService.UpdateInboundShipmentTransportDetails")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/awd/2024-05-09/inboundShipments/{shipmentId}/transport"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}
	if strlen(r.shipmentId) < 1 {
		return nil, reportError("shipmentId must have at least 1 elements")
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
