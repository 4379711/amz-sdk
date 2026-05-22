package merchant_fulfillment_v0

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MerchantFulfillmentAPIService MerchantFulfillmentAPI service
type MerchantFulfillmentAPIService service

type ApiCancelShipmentRequest struct {
	ctx        context.Context
	ApiService *MerchantFulfillmentAPIService
	shipmentId string
}

func (r ApiCancelShipmentRequest) Execute() (*CancelShipmentResponse, *http.Response, error) {
	return r.ApiService.CancelShipmentExecute(r)
}

/*
CancelShipment Method for CancelShipment

Cancel the shipment indicated by the specified shipment identifier.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the SP-API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId The Amazon-defined shipment identifier for the shipment to cancel.
	@return ApiCancelShipmentRequest
*/
func (a *MerchantFulfillmentAPIService) CancelShipment(ctx context.Context, shipmentId string) ApiCancelShipmentRequest {
	return ApiCancelShipmentRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return CancelShipmentResponse
func (a *MerchantFulfillmentAPIService) CancelShipmentExecute(r ApiCancelShipmentRequest) (*CancelShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CancelShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantFulfillmentAPIService.CancelShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mfn/v0/shipments/{shipmentId}"
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
	ApiService *MerchantFulfillmentAPIService
	body       *CreateShipmentRequest
}

// The request schema for the &#x60;CreateShipment&#x60; operation.
func (r ApiCreateShipmentRequest) Body(body CreateShipmentRequest) ApiCreateShipmentRequest {
	r.body = &body
	return r
}

func (r ApiCreateShipmentRequest) Execute() (*CreateShipmentResponse, *http.Response, error) {
	return r.ApiService.CreateShipmentExecute(r)
}

/*
CreateShipment Method for CreateShipment

Create a shipment with the information provided.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the SP-API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateShipmentRequest
*/
func (a *MerchantFulfillmentAPIService) CreateShipment(ctx context.Context) ApiCreateShipmentRequest {
	return ApiCreateShipmentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateShipmentResponse
func (a *MerchantFulfillmentAPIService) CreateShipmentExecute(r ApiCreateShipmentRequest) (*CreateShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantFulfillmentAPIService.CreateShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mfn/v0/shipments"

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

type ApiGetAdditionalSellerInputsRequest struct {
	ctx        context.Context
	ApiService *MerchantFulfillmentAPIService
	body       *GetAdditionalSellerInputsRequest
}

// The request schema for the &#x60;GetAdditionalSellerInputs&#x60; operation.
func (r ApiGetAdditionalSellerInputsRequest) Body(body GetAdditionalSellerInputsRequest) ApiGetAdditionalSellerInputsRequest {
	r.body = &body
	return r
}

func (r ApiGetAdditionalSellerInputsRequest) Execute() (*GetAdditionalSellerInputsResponse, *http.Response, error) {
	return r.ApiService.GetAdditionalSellerInputsExecute(r)
}

/*
GetAdditionalSellerInputs Method for GetAdditionalSellerInputs

Gets a list of additional seller inputs required for a ship method. This is generally used for international shipping.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the SP-API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAdditionalSellerInputsRequest
*/
func (a *MerchantFulfillmentAPIService) GetAdditionalSellerInputs(ctx context.Context) ApiGetAdditionalSellerInputsRequest {
	return ApiGetAdditionalSellerInputsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAdditionalSellerInputsResponse
func (a *MerchantFulfillmentAPIService) GetAdditionalSellerInputsExecute(r ApiGetAdditionalSellerInputsRequest) (*GetAdditionalSellerInputsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAdditionalSellerInputsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantFulfillmentAPIService.GetAdditionalSellerInputs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mfn/v0/additionalSellerInputs"

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

type ApiGetEligibleShipmentServicesRequest struct {
	ctx        context.Context
	ApiService *MerchantFulfillmentAPIService
	body       *GetEligibleShipmentServicesRequest
}

// The request schema for the &#x60;GetEligibleShipmentServices&#x60; operation.
func (r ApiGetEligibleShipmentServicesRequest) Body(body GetEligibleShipmentServicesRequest) ApiGetEligibleShipmentServicesRequest {
	r.body = &body
	return r
}

func (r ApiGetEligibleShipmentServicesRequest) Execute() (*GetEligibleShipmentServicesResponse, *http.Response, error) {
	return r.ApiService.GetEligibleShipmentServicesExecute(r)
}

/*
GetEligibleShipmentServices Method for GetEligibleShipmentServices

Returns a list of shipping service offers that satisfy the specified shipment request details.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 6 | 12 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the SP-API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEligibleShipmentServicesRequest
*/
func (a *MerchantFulfillmentAPIService) GetEligibleShipmentServices(ctx context.Context) ApiGetEligibleShipmentServicesRequest {
	return ApiGetEligibleShipmentServicesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEligibleShipmentServicesResponse
func (a *MerchantFulfillmentAPIService) GetEligibleShipmentServicesExecute(r ApiGetEligibleShipmentServicesRequest) (*GetEligibleShipmentServicesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEligibleShipmentServicesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantFulfillmentAPIService.GetEligibleShipmentServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mfn/v0/eligibleShippingServices"

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
	ApiService *MerchantFulfillmentAPIService
	shipmentId string
}

func (r ApiGetShipmentRequest) Execute() (*GetShipmentResponse, *http.Response, error) {
	return r.ApiService.GetShipmentExecute(r)
}

/*
GetShipment Method for GetShipment

Returns the shipment information for an existing shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the SP-API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId The Amazon-defined shipment identifier for the shipment.
	@return ApiGetShipmentRequest
*/
func (a *MerchantFulfillmentAPIService) GetShipment(ctx context.Context, shipmentId string) ApiGetShipmentRequest {
	return ApiGetShipmentRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return GetShipmentResponse
func (a *MerchantFulfillmentAPIService) GetShipmentExecute(r ApiGetShipmentRequest) (*GetShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantFulfillmentAPIService.GetShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mfn/v0/shipments/{shipmentId}"
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
