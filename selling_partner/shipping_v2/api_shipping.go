package shipping_v2

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
	ctx                     context.Context
	ApiService              *ShippingAPIService
	shipmentId              string
	xAmznShippingBusinessId *string
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiCancelShipmentRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiCancelShipmentRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiCancelShipmentRequest) Execute() (*CancelShipmentResponse, *http.Response, error) {
	return r.ApiService.CancelShipmentExecute(r)
}

/*
CancelShipment Method for CancelShipment

Cancels a purchased shipment. Returns an empty object if the shipment is successfully cancelled.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId The shipment identifier originally returned by the purchaseShipment operation.
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
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CancelShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.CancelShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/shipments/{shipmentId}/cancel"
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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiCreateClaimRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *CreateClaimRequest
	xAmznShippingBusinessId *string
}

// Request body for the createClaim operation
func (r ApiCreateClaimRequest) Body(body CreateClaimRequest) ApiCreateClaimRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiCreateClaimRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiCreateClaimRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiCreateClaimRequest) Execute() (*CreateClaimResponse, *http.Response, error) {
	return r.ApiService.CreateClaimExecute(r)
}

/*
CreateClaim Method for CreateClaim

This API will be used to create claim for single eligible shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateClaimRequest
*/
func (a *ShippingAPIService) CreateClaim(ctx context.Context) ApiCreateClaimRequest {
	return ApiCreateClaimRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateClaimResponse
func (a *ShippingAPIService) CreateClaimExecute(r ApiCreateClaimRequest) (*CreateClaimResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateClaimResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.CreateClaim")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/claims"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiDirectPurchaseShipmentRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *DirectPurchaseRequest
	xAmznIdempotencyKey     *string
	locale                  *string
	xAmznShippingBusinessId *string
}

// DirectPurchaseRequest body
func (r ApiDirectPurchaseShipmentRequest) Body(body DirectPurchaseRequest) ApiDirectPurchaseShipmentRequest {
	r.body = &body
	return r
}

// A unique value which the server uses to recognize subsequent retries of the same request.
func (r ApiDirectPurchaseShipmentRequest) XAmznIdempotencyKey(xAmznIdempotencyKey string) ApiDirectPurchaseShipmentRequest {
	r.xAmznIdempotencyKey = &xAmznIdempotencyKey
	return r
}

// The IETF Language Tag. Note that this only supports the primary language subtag with one secondary language subtag (i.e. en-US, fr-CA). The secondary language subtag is almost always a regional designation. This does not support additional subtags beyond the primary and secondary language subtags.
func (r ApiDirectPurchaseShipmentRequest) Locale(locale string) ApiDirectPurchaseShipmentRequest {
	r.locale = &locale
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiDirectPurchaseShipmentRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiDirectPurchaseShipmentRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiDirectPurchaseShipmentRequest) Execute() (*DirectPurchaseResponse, *http.Response, error) {
	return r.ApiService.DirectPurchaseShipmentExecute(r)
}

/*
DirectPurchaseShipment Method for DirectPurchaseShipment

Purchases the shipping service for a shipment using the best fit service offering. Returns purchase related details and documents.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDirectPurchaseShipmentRequest
*/
func (a *ShippingAPIService) DirectPurchaseShipment(ctx context.Context) ApiDirectPurchaseShipmentRequest {
	return ApiDirectPurchaseShipmentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DirectPurchaseResponse
func (a *ShippingAPIService) DirectPurchaseShipmentExecute(r ApiDirectPurchaseShipmentRequest) (*DirectPurchaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DirectPurchaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.DirectPurchaseShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/shipments/directPurchase"

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
	if r.xAmznIdempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-IdempotencyKey", r.xAmznIdempotencyKey, "", "")
	}
	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "locale", r.locale, "", "")
	}
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGenerateCollectionFormRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *GenerateCollectionFormRequest
	xAmznIdempotencyKey     *string
	xAmznShippingBusinessId *string
}

// GenerateCollectionFormRequest body
func (r ApiGenerateCollectionFormRequest) Body(body GenerateCollectionFormRequest) ApiGenerateCollectionFormRequest {
	r.body = &body
	return r
}

// A unique value which the server uses to recognize subsequent retries of the same request.
func (r ApiGenerateCollectionFormRequest) XAmznIdempotencyKey(xAmznIdempotencyKey string) ApiGenerateCollectionFormRequest {
	r.xAmznIdempotencyKey = &xAmznIdempotencyKey
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGenerateCollectionFormRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGenerateCollectionFormRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGenerateCollectionFormRequest) Execute() (*GenerateCollectionFormResponse, *http.Response, error) {
	return r.ApiService.GenerateCollectionFormExecute(r)
}

/*
GenerateCollectionForm Method for GenerateCollectionForm

This API  Call to generate the collection form.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGenerateCollectionFormRequest
*/
func (a *ShippingAPIService) GenerateCollectionForm(ctx context.Context) ApiGenerateCollectionFormRequest {
	return ApiGenerateCollectionFormRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GenerateCollectionFormResponse
func (a *ShippingAPIService) GenerateCollectionFormExecute(r ApiGenerateCollectionFormRequest) (*GenerateCollectionFormResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GenerateCollectionFormResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GenerateCollectionForm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/collectionForms"

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
	if r.xAmznIdempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-IdempotencyKey", r.xAmznIdempotencyKey, "", "")
	}
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetAccessPointsRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	accessPointTypes        *[]string
	countryCode             *string
	postalCode              *string
	xAmznShippingBusinessId *string
}

// Access point types
func (r ApiGetAccessPointsRequest) AccessPointTypes(accessPointTypes []string) ApiGetAccessPointsRequest {
	r.accessPointTypes = &accessPointTypes
	return r
}

// Country code for access point
func (r ApiGetAccessPointsRequest) CountryCode(countryCode string) ApiGetAccessPointsRequest {
	r.countryCode = &countryCode
	return r
}

// postal code for access point
func (r ApiGetAccessPointsRequest) PostalCode(postalCode string) ApiGetAccessPointsRequest {
	r.postalCode = &postalCode
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetAccessPointsRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetAccessPointsRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetAccessPointsRequest) Execute() (*GetAccessPointsResponse, *http.Response, error) {
	return r.ApiService.GetAccessPointsExecute(r)
}

/*
GetAccessPoints Method for GetAccessPoints

Returns a list of access points in proximity of input postal code.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAccessPointsRequest
*/
func (a *ShippingAPIService) GetAccessPoints(ctx context.Context) ApiGetAccessPointsRequest {
	return ApiGetAccessPointsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAccessPointsResponse
func (a *ShippingAPIService) GetAccessPointsExecute(r ApiGetAccessPointsRequest) (*GetAccessPointsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAccessPointsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetAccessPoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/accessPoints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessPointTypes == nil {
		return localVarReturnValue, nil, reportError("accessPointTypes is required and must be specified")
	}
	if r.countryCode == nil {
		return localVarReturnValue, nil, reportError("countryCode is required and must be specified")
	}
	if r.postalCode == nil {
		return localVarReturnValue, nil, reportError("postalCode is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "accessPointTypes", r.accessPointTypes, "form", "csv")
	parameterAddToHeaderOrQuery(localVarQueryParams, "countryCode", r.countryCode, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "postalCode", r.postalCode, "", "")
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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetAdditionalInputsRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	requestToken            *string
	rateId                  *string
	xAmznShippingBusinessId *string
}

// The request token returned in the response to the getRates operation.
func (r ApiGetAdditionalInputsRequest) RequestToken(requestToken string) ApiGetAdditionalInputsRequest {
	r.requestToken = &requestToken
	return r
}

// The rate identifier for the shipping offering (rate) returned in the response to the getRates operation.
func (r ApiGetAdditionalInputsRequest) RateId(rateId string) ApiGetAdditionalInputsRequest {
	r.rateId = &rateId
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetAdditionalInputsRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetAdditionalInputsRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetAdditionalInputsRequest) Execute() (*GetAdditionalInputsResponse, *http.Response, error) {
	return r.ApiService.GetAdditionalInputsExecute(r)
}

/*
GetAdditionalInputs Method for GetAdditionalInputs

Returns the JSON schema to use for providing additional inputs when needed to purchase a shipping offering. Call the getAdditionalInputs operation when the response to a previous call to the getRates operation indicates that additional inputs are required for the rate (shipping offering) that you want to purchase.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAdditionalInputsRequest
*/
func (a *ShippingAPIService) GetAdditionalInputs(ctx context.Context) ApiGetAdditionalInputsRequest {
	return ApiGetAdditionalInputsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAdditionalInputsResponse
func (a *ShippingAPIService) GetAdditionalInputsExecute(r ApiGetAdditionalInputsRequest) (*GetAdditionalInputsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAdditionalInputsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetAdditionalInputs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/shipments/additionalInputs/schema"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestToken == nil {
		return localVarReturnValue, nil, reportError("requestToken is required and must be specified")
	}
	if r.rateId == nil {
		return localVarReturnValue, nil, reportError("rateId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "requestToken", r.requestToken, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "rateId", r.rateId, "", "")
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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetCarrierAccountFormInputsRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	xAmznShippingBusinessId *string
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetCarrierAccountFormInputsRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetCarrierAccountFormInputsRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetCarrierAccountFormInputsRequest) Execute() (*GetCarrierAccountFormInputsResponse, *http.Response, error) {
	return r.ApiService.GetCarrierAccountFormInputsExecute(r)
}

/*
GetCarrierAccountFormInputs Method for GetCarrierAccountFormInputs

This API will return a list of input schema required to register a shipper account with the carrier.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCarrierAccountFormInputsRequest
*/
func (a *ShippingAPIService) GetCarrierAccountFormInputs(ctx context.Context) ApiGetCarrierAccountFormInputsRequest {
	return ApiGetCarrierAccountFormInputsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCarrierAccountFormInputsResponse
func (a *ShippingAPIService) GetCarrierAccountFormInputsExecute(r ApiGetCarrierAccountFormInputsRequest) (*GetCarrierAccountFormInputsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCarrierAccountFormInputsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetCarrierAccountFormInputs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/carrierAccountFormInputs"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetCarrierAccountsRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *GetCarrierAccountsRequest
	xAmznShippingBusinessId *string
}

// GetCarrierAccountsRequest body
func (r ApiGetCarrierAccountsRequest) Body(body GetCarrierAccountsRequest) ApiGetCarrierAccountsRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetCarrierAccountsRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetCarrierAccountsRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetCarrierAccountsRequest) Execute() (*GetCarrierAccountsResponse, *http.Response, error) {
	return r.ApiService.GetCarrierAccountsExecute(r)
}

/*
GetCarrierAccounts Method for GetCarrierAccounts

This API will return Get all carrier accounts for a merchant.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCarrierAccountsRequest
*/
func (a *ShippingAPIService) GetCarrierAccounts(ctx context.Context) ApiGetCarrierAccountsRequest {
	return ApiGetCarrierAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCarrierAccountsResponse
func (a *ShippingAPIService) GetCarrierAccountsExecute(r ApiGetCarrierAccountsRequest) (*GetCarrierAccountsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCarrierAccountsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetCarrierAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/carrierAccounts"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetCollectionFormRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	collectionFormId        string
	xAmznShippingBusinessId *string
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetCollectionFormRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetCollectionFormRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetCollectionFormRequest) Execute() (*GetCollectionFormResponse, *http.Response, error) {
	return r.ApiService.GetCollectionFormExecute(r)
}

/*
GetCollectionForm Method for GetCollectionForm

This API reprint a collection form.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param collectionFormId collection form Id to reprint a collection.
	@return ApiGetCollectionFormRequest
*/
func (a *ShippingAPIService) GetCollectionForm(ctx context.Context, collectionFormId string) ApiGetCollectionFormRequest {
	return ApiGetCollectionFormRequest{
		ApiService:       a,
		ctx:              ctx,
		collectionFormId: collectionFormId,
	}
}

// Execute executes the request
//
//	@return GetCollectionFormResponse
func (a *ShippingAPIService) GetCollectionFormExecute(r ApiGetCollectionFormRequest) (*GetCollectionFormResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCollectionFormResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetCollectionForm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/collectionForms/{collectionFormId}"
	localVarPath = strings.Replace(localVarPath, "{"+"collectionFormId"+"}", url.PathEscape(parameterValueToString(r.collectionFormId, "collectionFormId")), -1)

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetCollectionFormHistoryRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *GetCollectionFormHistoryRequest
	xAmznShippingBusinessId *string
}

// GetCollectionFormHistoryRequest body
func (r ApiGetCollectionFormHistoryRequest) Body(body GetCollectionFormHistoryRequest) ApiGetCollectionFormHistoryRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetCollectionFormHistoryRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetCollectionFormHistoryRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetCollectionFormHistoryRequest) Execute() (*GetCollectionFormHistoryResponse, *http.Response, error) {
	return r.ApiService.GetCollectionFormHistoryExecute(r)
}

/*
GetCollectionFormHistory Method for GetCollectionFormHistory

This API Call to get the history of the previously generated collection forms.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCollectionFormHistoryRequest
*/
func (a *ShippingAPIService) GetCollectionFormHistory(ctx context.Context) ApiGetCollectionFormHistoryRequest {
	return ApiGetCollectionFormHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCollectionFormHistoryResponse
func (a *ShippingAPIService) GetCollectionFormHistoryExecute(r ApiGetCollectionFormHistoryRequest) (*GetCollectionFormHistoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCollectionFormHistoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetCollectionFormHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/collectionForms/history"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetRatesRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *GetRatesRequest
	xAmznShippingBusinessId *string
}

// GetRatesRequest body
func (r ApiGetRatesRequest) Body(body GetRatesRequest) ApiGetRatesRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetRatesRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetRatesRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetRatesRequest) Execute() (*GetRatesResponse, *http.Response, error) {
	return r.ApiService.GetRatesExecute(r)
}

/*
GetRates Method for GetRates

Returns the available shipping service offerings.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

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

	localVarPath := localBasePath + "/shipping/v2/shipments/rates"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetShipmentDocumentsRequest struct {
	ctx                      context.Context
	ApiService               *ShippingAPIService
	shipmentId               string
	packageClientReferenceId *string
	format                   *string
	dpi                      *float32
	xAmznShippingBusinessId  *string
}

// The package client reference identifier originally provided in the request body parameter for the getRates operation.
func (r ApiGetShipmentDocumentsRequest) PackageClientReferenceId(packageClientReferenceId string) ApiGetShipmentDocumentsRequest {
	r.packageClientReferenceId = &packageClientReferenceId
	return r
}

// The file format of the document. Must be one of the supported formats returned by the getRates operation.
func (r ApiGetShipmentDocumentsRequest) Format(format string) ApiGetShipmentDocumentsRequest {
	r.format = &format
	return r
}

// The resolution of the document (for example, 300 means 300 dots per inch). Must be one of the supported resolutions returned in the response to the getRates operation.
func (r ApiGetShipmentDocumentsRequest) Dpi(dpi float32) ApiGetShipmentDocumentsRequest {
	r.dpi = &dpi
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetShipmentDocumentsRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetShipmentDocumentsRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetShipmentDocumentsRequest) Execute() (*GetShipmentDocumentsResponse, *http.Response, error) {
	return r.ApiService.GetShipmentDocumentsExecute(r)
}

/*
GetShipmentDocuments Method for GetShipmentDocuments

Returns the shipping documents associated with a package in a shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId The shipment identifier originally returned by the purchaseShipment operation.
	@return ApiGetShipmentDocumentsRequest
*/
func (a *ShippingAPIService) GetShipmentDocuments(ctx context.Context, shipmentId string) ApiGetShipmentDocumentsRequest {
	return ApiGetShipmentDocumentsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
//
//	@return GetShipmentDocumentsResponse
func (a *ShippingAPIService) GetShipmentDocumentsExecute(r ApiGetShipmentDocumentsRequest) (*GetShipmentDocumentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetShipmentDocumentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetShipmentDocuments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/shipments/{shipmentId}/documents"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.packageClientReferenceId == nil {
		return localVarReturnValue, nil, reportError("packageClientReferenceId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "packageClientReferenceId", r.packageClientReferenceId, "", "")
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "", "")
	}
	if r.dpi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dpi", r.dpi, "", "")
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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetTrackingRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	trackingId              *string
	carrierId               *string
	xAmznShippingBusinessId *string
}

// A carrier-generated tracking identifier originally returned by the purchaseShipment operation.
func (r ApiGetTrackingRequest) TrackingId(trackingId string) ApiGetTrackingRequest {
	r.trackingId = &trackingId
	return r
}

// A carrier identifier originally returned by the getRates operation for the selected rate.
func (r ApiGetTrackingRequest) CarrierId(carrierId string) ApiGetTrackingRequest {
	r.carrierId = &carrierId
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetTrackingRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetTrackingRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetTrackingRequest) Execute() (*GetTrackingResponse, *http.Response, error) {
	return r.ApiService.GetTrackingExecute(r)
}

/*
GetTracking Method for GetTracking

Returns tracking information for a purchased shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTrackingRequest
*/
func (a *ShippingAPIService) GetTracking(ctx context.Context) ApiGetTrackingRequest {
	return ApiGetTrackingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTrackingResponse
func (a *ShippingAPIService) GetTrackingExecute(r ApiGetTrackingRequest) (*GetTrackingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTrackingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetTracking")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/tracking"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.trackingId == nil {
		return localVarReturnValue, nil, reportError("trackingId is required and must be specified")
	}
	if r.carrierId == nil {
		return localVarReturnValue, nil, reportError("carrierId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "trackingId", r.trackingId, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "carrierId", r.carrierId, "", "")
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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiGetUnmanifestedShipmentsRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *GetUnmanifestedShipmentsRequest
	xAmznShippingBusinessId *string
}

// GetUmanifestedShipmentsRequest body
func (r ApiGetUnmanifestedShipmentsRequest) Body(body GetUnmanifestedShipmentsRequest) ApiGetUnmanifestedShipmentsRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiGetUnmanifestedShipmentsRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiGetUnmanifestedShipmentsRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiGetUnmanifestedShipmentsRequest) Execute() (*GetUnmanifestedShipmentsResponse, *http.Response, error) {
	return r.ApiService.GetUnmanifestedShipmentsExecute(r)
}

/*
GetUnmanifestedShipments Method for GetUnmanifestedShipments

This API Get all unmanifested carriers with shipment locations. Any locations which has unmanifested shipments

	with an eligible carrier for manifesting shall be returned.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetUnmanifestedShipmentsRequest
*/
func (a *ShippingAPIService) GetUnmanifestedShipments(ctx context.Context) ApiGetUnmanifestedShipmentsRequest {
	return ApiGetUnmanifestedShipmentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUnmanifestedShipmentsResponse
func (a *ShippingAPIService) GetUnmanifestedShipmentsExecute(r ApiGetUnmanifestedShipmentsRequest) (*GetUnmanifestedShipmentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUnmanifestedShipmentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.GetUnmanifestedShipments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/unmanifestedShipments"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiLinkCarrierAccountRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *LinkCarrierAccountRequest
	carrierId               string
	xAmznShippingBusinessId *string
}

// LinkCarrierAccountRequest body
func (r ApiLinkCarrierAccountRequest) Body(body LinkCarrierAccountRequest) ApiLinkCarrierAccountRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiLinkCarrierAccountRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiLinkCarrierAccountRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiLinkCarrierAccountRequest) Execute() (*LinkCarrierAccountResponse, *http.Response, error) {
	return r.ApiService.LinkCarrierAccountExecute(r)
}

/*
LinkCarrierAccount Method for LinkCarrierAccount

This API associates/links the specified carrier account with the merchant.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param carrierId An identifier for the carrier with which the seller's account is being linked.
	@return ApiLinkCarrierAccountRequest
*/
func (a *ShippingAPIService) LinkCarrierAccount(ctx context.Context, carrierId string) ApiLinkCarrierAccountRequest {
	return ApiLinkCarrierAccountRequest{
		ApiService: a,
		ctx:        ctx,
		carrierId:  carrierId,
	}
}

// Execute executes the request
//
//	@return LinkCarrierAccountResponse
func (a *ShippingAPIService) LinkCarrierAccountExecute(r ApiLinkCarrierAccountRequest) (*LinkCarrierAccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LinkCarrierAccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.LinkCarrierAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/carrierAccounts/{carrierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"carrierId"+"}", url.PathEscape(parameterValueToString(r.carrierId, "carrierId")), -1)

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiLinkCarrierAccount_0Request struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *LinkCarrierAccountRequest
	carrierId               string
	xAmznShippingBusinessId *string
}

// LinkCarrierAccountRequest body
func (r ApiLinkCarrierAccount_0Request) Body(body LinkCarrierAccountRequest) ApiLinkCarrierAccount_0Request {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiLinkCarrierAccount_0Request) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiLinkCarrierAccount_0Request {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiLinkCarrierAccount_0Request) Execute() (*LinkCarrierAccountResponse, *http.Response, error) {
	return r.ApiService.LinkCarrierAccount_1Execute(r)
}

/*
LinkCarrierAccount_0 Method for LinkCarrierAccount_0

This API associates/links the specified carrier account with the merchant.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param carrierId An identifier for the carrier with which the seller's account is being linked.
	@return ApiLinkCarrierAccount_0Request
*/
func (a *ShippingAPIService) LinkCarrierAccount_1(ctx context.Context, carrierId string) ApiLinkCarrierAccount_0Request {
	return ApiLinkCarrierAccount_0Request{
		ApiService: a,
		ctx:        ctx,
		carrierId:  carrierId,
	}
}

// Execute executes the request
//
//	@return LinkCarrierAccountResponse
func (a *ShippingAPIService) LinkCarrierAccount_1Execute(r ApiLinkCarrierAccount_0Request) (*LinkCarrierAccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LinkCarrierAccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.LinkCarrierAccount_1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/carrierAccounts/{carrierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"carrierId"+"}", url.PathEscape(parameterValueToString(r.carrierId, "carrierId")), -1)

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiOneClickShipmentRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *OneClickShipmentRequest
	xAmznShippingBusinessId *string
}

// OneClickShipmentRequest body
func (r ApiOneClickShipmentRequest) Body(body OneClickShipmentRequest) ApiOneClickShipmentRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiOneClickShipmentRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiOneClickShipmentRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiOneClickShipmentRequest) Execute() (*OneClickShipmentResponse, *http.Response, error) {
	return r.ApiService.OneClickShipmentExecute(r)
}

/*
OneClickShipment Method for OneClickShipment

Purchases a shipping service identifier and returns purchase-related details and documents.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOneClickShipmentRequest
*/
func (a *ShippingAPIService) OneClickShipment(ctx context.Context) ApiOneClickShipmentRequest {
	return ApiOneClickShipmentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OneClickShipmentResponse
func (a *ShippingAPIService) OneClickShipmentExecute(r ApiOneClickShipmentRequest) (*OneClickShipmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OneClickShipmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.OneClickShipment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/oneClickShipment"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *PurchaseShipmentRequest
	xAmznIdempotencyKey     *string
	xAmznShippingBusinessId *string
}

// PurchaseShipmentRequest body
func (r ApiPurchaseShipmentRequest) Body(body PurchaseShipmentRequest) ApiPurchaseShipmentRequest {
	r.body = &body
	return r
}

// A unique value which the server uses to recognize subsequent retries of the same request.
func (r ApiPurchaseShipmentRequest) XAmznIdempotencyKey(xAmznIdempotencyKey string) ApiPurchaseShipmentRequest {
	r.xAmznIdempotencyKey = &xAmznIdempotencyKey
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiPurchaseShipmentRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiPurchaseShipmentRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiPurchaseShipmentRequest) Execute() (*PurchaseShipmentResponse, *http.Response, error) {
	return r.ApiService.PurchaseShipmentExecute(r)
}

/*
PurchaseShipment Method for PurchaseShipment

Purchases a shipping service and returns purchase related details and documents.

Note: You must complete the purchase within 10 minutes of rate creation by the shipping service provider. If you make the request after the 10 minutes have expired, you will receive an error response with the error code equal to "TOKEN_EXPIRED". If you receive this error response, you must get the rates for the shipment again.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

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

	localVarPath := localBasePath + "/shipping/v2/shipments"

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
	if r.xAmznIdempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-IdempotencyKey", r.xAmznIdempotencyKey, "", "")
	}
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiSubmitNdrFeedbackRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *SubmitNdrFeedbackRequest
	xAmznShippingBusinessId *string
}

// Request body for ndrFeedback operation
func (r ApiSubmitNdrFeedbackRequest) Body(body SubmitNdrFeedbackRequest) ApiSubmitNdrFeedbackRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiSubmitNdrFeedbackRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiSubmitNdrFeedbackRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiSubmitNdrFeedbackRequest) Execute() (*http.Response, error) {
	return r.ApiService.SubmitNdrFeedbackExecute(r)
}

/*
SubmitNdrFeedback Method for SubmitNdrFeedback

This API submits the NDR (Non-delivery Report) Feedback for any eligible shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSubmitNdrFeedbackRequest
*/
func (a *ShippingAPIService) SubmitNdrFeedback(ctx context.Context) ApiSubmitNdrFeedbackRequest {
	return ApiSubmitNdrFeedbackRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ShippingAPIService) SubmitNdrFeedbackExecute(r ApiSubmitNdrFeedbackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.SubmitNdrFeedback")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/ndrFeedback"

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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

type ApiUnlinkCarrierAccountRequest struct {
	ctx                     context.Context
	ApiService              *ShippingAPIService
	body                    *UnlinkCarrierAccountRequest
	carrierId               string
	xAmznShippingBusinessId *string
}

// UnlinkCarrierAccountRequest body
func (r ApiUnlinkCarrierAccountRequest) Body(body UnlinkCarrierAccountRequest) ApiUnlinkCarrierAccountRequest {
	r.body = &body
	return r
}

// Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
func (r ApiUnlinkCarrierAccountRequest) XAmznShippingBusinessId(xAmznShippingBusinessId string) ApiUnlinkCarrierAccountRequest {
	r.xAmznShippingBusinessId = &xAmznShippingBusinessId
	return r
}

func (r ApiUnlinkCarrierAccountRequest) Execute() (*UnlinkCarrierAccountResponse, *http.Response, error) {
	return r.ApiService.UnlinkCarrierAccountExecute(r)
}

/*
UnlinkCarrierAccount Method for UnlinkCarrierAccount

This API Unlink the specified carrier account with the merchant.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 80 | 100 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param carrierId carrier Id to unlink with merchant.
	@return ApiUnlinkCarrierAccountRequest
*/
func (a *ShippingAPIService) UnlinkCarrierAccount(ctx context.Context, carrierId string) ApiUnlinkCarrierAccountRequest {
	return ApiUnlinkCarrierAccountRequest{
		ApiService: a,
		ctx:        ctx,
		carrierId:  carrierId,
	}
}

// Execute executes the request
//
//	@return UnlinkCarrierAccountResponse
func (a *ShippingAPIService) UnlinkCarrierAccountExecute(r ApiUnlinkCarrierAccountRequest) (*UnlinkCarrierAccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UnlinkCarrierAccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingAPIService.UnlinkCarrierAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping/v2/carrierAccounts/{carrierId}/unlink"
	localVarPath = strings.Replace(localVarPath, "{"+"carrierId"+"}", url.PathEscape(parameterValueToString(r.carrierId, "carrierId")), -1)

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
	if r.xAmznShippingBusinessId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amzn-shipping-business-id", r.xAmznShippingBusinessId, "", "")
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
