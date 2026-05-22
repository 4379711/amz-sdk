package shipping

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ShippingAPIService ShippingAPI service
type ShippingAPIService service

type ApiCancelShipmentRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	shipmentId string
}

func (r ApiCancelShipmentRequest) Execute() (*CancelShipmentResponse, *http.Response, error) {
	return r.ApiService.CancelShipmentExecute(r)
}

/*
CancelShipment Method for CancelShipment

Cancel a shipment by the given shipmentId.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId Shipment Id to cancel a shipment
	@return ApiCancelShipmentRequest
*/
func (a *ShippingAPIService) CancelShipment(ctx context.Context, shipmentId string) ApiCancelShipmentRequest {
	return ApiCancelShipmentRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return CancelShipmentResponse
func (a *ShippingAPIService) CancelShipmentExecute(r ApiCancelShipmentRequest) (*CancelShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CancelShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.CancelShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/shipments/{shipmentId}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

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

type ApiCreateShipmentRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	body       *CreateShipmentRequest
}

// CreateShipmentRequest Body
func (r ApiCreateShipmentRequest) Body(body CreateShipmentRequest) ApiCreateShipmentRequest {
	r.body = &body
	return r
}

func (r ApiCreateShipmentRequest) Execute() (*CreateShipmentResponse, *http.Response, error) {
	return r.ApiService.CreateShipmentExecute(r)
}

/*
CreateShipment Method for CreateShipment

Create a new shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateShipmentRequest
*/
func (a *ShippingAPIService) CreateShipment(ctx context.Context) ApiCreateShipmentRequest {
	return ApiCreateShipmentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateShipmentResponse
func (a *ShippingAPIService) CreateShipmentExecute(r ApiCreateShipmentRequest) (*CreateShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.CreateShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/shipments"

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

type ApiGetAccountRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
}

func (r ApiGetAccountRequest) Execute() (*GetAccountResponse, *http.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Method for GetAccount

Verify if the current account is valid.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAccountRequest
*/
func (a *ShippingAPIService) GetAccount(ctx context.Context) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAccountResponse
func (a *ShippingAPIService) GetAccountExecute(r ApiGetAccountRequest) (*GetAccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/account"

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

type ApiGetRatesRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	body       *GetRatesRequest
}

// GetRatesRequest body
func (r ApiGetRatesRequest) Body(body GetRatesRequest) ApiGetRatesRequest {
	r.body = &body
	return r
}

func (r ApiGetRatesRequest) Execute() (*GetRatesResponse, *http.Response, error) {
	return r.ApiService.GetRatesExecute(r)
}

/*
GetRates Method for GetRates

Get service rates.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRatesRequest
*/
func (a *ShippingAPIService) GetRates(ctx context.Context) ApiGetRatesRequest {
	return ApiGetRatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRatesResponse
func (a *ShippingAPIService) GetRatesExecute(r ApiGetRatesRequest) (*GetRatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetRatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetRates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/rates"

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

type ApiGetShipmentRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	shipmentId string
}

func (r ApiGetShipmentRequest) Execute() (*GetShipmentResponse, *http.Response, error) {
	return r.ApiService.GetShipmentExecute(r)
}

/*
GetShipment Method for GetShipment

Return the entire shipment object for the shipmentId.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId Shipment id to return the entire shipment object
	@return ApiGetShipmentRequest
*/
func (a *ShippingAPIService) GetShipment(ctx context.Context, shipmentId string) ApiGetShipmentRequest {
	return ApiGetShipmentRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return GetShipmentResponse
func (a *ShippingAPIService) GetShipmentExecute(r ApiGetShipmentRequest) (*GetShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/shipments/{shipmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

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

type ApiGetTrackingInformationRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	trackingId string
}

func (r ApiGetTrackingInformationRequest) Execute() (*GetTrackingInformationResponse, *http.Response, error) {
	return r.ApiService.GetTrackingInformationExecute(r)
}

/*
GetTrackingInformation Method for GetTrackingInformation

Return the tracking information of a shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trackingId Tracking Id
	@return ApiGetTrackingInformationRequest
*/
func (a *ShippingAPIService) GetTrackingInformation(ctx context.Context, trackingId string) ApiGetTrackingInformationRequest {
	return ApiGetTrackingInformationRequest{
		ApiService: a,
		ctx:        ctx,
		trackingId: trackingId,
	}
}

// Execute executes the request
//
//	@return GetTrackingInformationResponse
func (a *ShippingAPIService) GetTrackingInformationExecute(r ApiGetTrackingInformationRequest) (*GetTrackingInformationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTrackingInformationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetTrackingInformation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/tracking/{trackingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"trackingId"+"}", url.PathEscape(parameterValueToString(r.trackingId, "trackingId")), -1)

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

type ApiPurchaseLabelsRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	body       *PurchaseLabelsRequest
	shipmentId string
}

// PurchaseShippingLabelRequest body
func (r ApiPurchaseLabelsRequest) Body(body PurchaseLabelsRequest) ApiPurchaseLabelsRequest {
	r.body = &body
	return r
}

func (r ApiPurchaseLabelsRequest) Execute() (*PurchaseLabelsResponse, *http.Response, error) {
	return r.ApiService.PurchaseLabelsExecute(r)
}

/*
PurchaseLabels Method for PurchaseLabels

Purchase shipping labels based on a given rate.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId Shipment id for purchase shipping label
	@return ApiPurchaseLabelsRequest
*/
func (a *ShippingAPIService) PurchaseLabels(ctx context.Context, shipmentId string) ApiPurchaseLabelsRequest {
	return ApiPurchaseLabelsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return PurchaseLabelsResponse
func (a *ShippingAPIService) PurchaseLabelsExecute(r ApiPurchaseLabelsRequest) (*PurchaseLabelsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PurchaseLabelsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.PurchaseLabels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/shipments/{shipmentId}/purchaseLabels"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

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

type ApiPurchaseShipmentRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	body       *PurchaseShipmentRequest
}

// PurchaseShipmentRequest body
func (r ApiPurchaseShipmentRequest) Body(body PurchaseShipmentRequest) ApiPurchaseShipmentRequest {
	r.body = &body
	return r
}

func (r ApiPurchaseShipmentRequest) Execute() (*PurchaseShipmentResponse, *http.Response, error) {
	return r.ApiService.PurchaseShipmentExecute(r)
}

/*
PurchaseShipment Method for PurchaseShipment

Purchase shipping labels.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPurchaseShipmentRequest
*/
func (a *ShippingAPIService) PurchaseShipment(ctx context.Context) ApiPurchaseShipmentRequest {
	return ApiPurchaseShipmentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PurchaseShipmentResponse
func (a *ShippingAPIService) PurchaseShipmentExecute(r ApiPurchaseShipmentRequest) (*PurchaseShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PurchaseShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.PurchaseShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/purchaseShipment"

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

type ApiRetrieveShippingLabelRequest struct {
	ctx        context.Context
	ApiService *ShippingAPIService
	body       *RetrieveShippingLabelRequest
	shipmentId string
	trackingId string
}

// RetrieveShippingLabelRequest body
func (r ApiRetrieveShippingLabelRequest) Body(body RetrieveShippingLabelRequest) ApiRetrieveShippingLabelRequest {
	r.body = &body
	return r
}

func (r ApiRetrieveShippingLabelRequest) Execute() (*RetrieveShippingLabelResponse, *http.Response, error) {
	return r.ApiService.RetrieveShippingLabelExecute(r)
}

/*
RetrieveShippingLabel Method for RetrieveShippingLabel

Retrieve shipping label based on the shipment id and tracking id.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId Shipment Id to retreive label
	@param trackingId Tracking Id
	@return ApiRetrieveShippingLabelRequest
*/
func (a *ShippingAPIService) RetrieveShippingLabel(ctx context.Context, shipmentId string, trackingId string) ApiRetrieveShippingLabelRequest {
	return ApiRetrieveShippingLabelRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
		trackingId: trackingId,
	}
}

// Execute executes the request
//
//	@return RetrieveShippingLabelResponse
func (a *ShippingAPIService) RetrieveShippingLabelExecute(r ApiRetrieveShippingLabelRequest) (*RetrieveShippingLabelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RetrieveShippingLabelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.RetrieveShippingLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v1/shipments/{shipmentId}/containers/{trackingId}/label"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trackingId"+"}", url.PathEscape(parameterValueToString(r.trackingId, "trackingId")), -1)

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
