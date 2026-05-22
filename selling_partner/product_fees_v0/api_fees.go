package product_fees_v0

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// FeesAPIService FeesAPI service
type FeesAPIService service

type ApiGetMyFeesEstimateForASINRequest struct {
	ctx        context.Context
	ApiService *FeesAPIService
	body       *GetMyFeesEstimateRequest
	asin       string
}

func (r ApiGetMyFeesEstimateForASINRequest) Body(body GetMyFeesEstimateRequest) ApiGetMyFeesEstimateForASINRequest {
	r.body = &body
	return r
}

func (r ApiGetMyFeesEstimateForASINRequest) Execute() (*GetMyFeesEstimateResponse, *http.Response, error) {
	return r.ApiService.GetMyFeesEstimateForASINExecute(r)
}

/*
GetMyFeesEstimateForASIN Method for GetMyFeesEstimateForASIN

Returns the estimated fees for the item indicated by the specified ASIN in the marketplace specified in the request body.

You can call `getMyFeesEstimateForASIN` for an item on behalf of a selling partner before the selling partner sets the item's price. The selling partner can then take estimated fees into account. Each fees request must include an original identifier. This identifier is included in the fees estimate so you can correlate a fees estimate with the original request.

**Note:** This identifier value is used to identify an estimate. Actual costs may vary. Search "fees" in [Seller Central](https://sellercentral.amazon.com/) and consult the store-specific fee schedule for the most up-to-date information.

**Note:** When using the `getMyFeesEstimateForASIN` operation with an ASIN, the fee estimates might be different. This is because these estimates use the item's catalog size, which might not always match the actual size of the item sent to Amazon.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param asin The Amazon Standard Identification Number (ASIN) of the item.
	@return ApiGetMyFeesEstimateForASINRequest
*/
func (a *FeesAPIService) GetMyFeesEstimateForASIN(ctx context.Context, asin string) ApiGetMyFeesEstimateForASINRequest {
	return ApiGetMyFeesEstimateForASINRequest{
		ApiService: a,
		ctx:        ctx,
		asin:       asin,
	}
}

// Execute executes the request
//
//	@return GetMyFeesEstimateResponse
func (a *FeesAPIService) GetMyFeesEstimateForASINExecute(r ApiGetMyFeesEstimateForASINRequest) (*GetMyFeesEstimateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMyFeesEstimateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.GetMyFeesEstimateForASIN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/fees/v0/items/{Asin}/feesEstimate"
	localVarPath = strings.Replace(localVarPath, "{"+"Asin"+"}", url.PathEscape(parameterValueToString(r.asin, "asin")), -1)

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

type ApiGetMyFeesEstimateForSKURequest struct {
	ctx        context.Context
	ApiService *FeesAPIService
	body       *GetMyFeesEstimateRequest
	sellerSKU  string
}

func (r ApiGetMyFeesEstimateForSKURequest) Body(body GetMyFeesEstimateRequest) ApiGetMyFeesEstimateForSKURequest {
	r.body = &body
	return r
}

func (r ApiGetMyFeesEstimateForSKURequest) Execute() (*GetMyFeesEstimateResponse, *http.Response, error) {
	return r.ApiService.GetMyFeesEstimateForSKUExecute(r)
}

/*
GetMyFeesEstimateForSKU Method for GetMyFeesEstimateForSKU

Returns the estimated fees for the item indicated by the specified seller SKU in the marketplace specified in the request body.

**Note:** The parameters associated with this operation may contain special characters that require URL encoding to call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

You can call `getMyFeesEstimateForSKU` for an item on behalf of a selling partner before the selling partner sets the item's price. The selling partner can then take any estimated fees into account. Each fees estimate request must include an original identifier. This identifier is included in the fees estimate so that you can correlate a fees estimate with the original request.

**Note:** This identifier value is used to identify an estimate. Actual costs may vary. Search "fees" in [Seller Central](https://sellercentral.amazon.com/) and consult the store-specific fee schedule for the most up-to-date information.

**Note:** When sellers use the `getMyFeesEstimateForSKU` operation with their `SellerSKU`, they get accurate fees based on real item measurements, but only after they've sent their items to Amazon.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sellerSKU Used to identify an item in the given marketplace. SellerSKU is qualified by the seller's SellerId, which is included with every operation that you submit.
	@return ApiGetMyFeesEstimateForSKURequest
*/
func (a *FeesAPIService) GetMyFeesEstimateForSKU(ctx context.Context, sellerSKU string) ApiGetMyFeesEstimateForSKURequest {
	return ApiGetMyFeesEstimateForSKURequest{
		ApiService: a,
		ctx:        ctx,
		sellerSKU:  sellerSKU,
	}
}

// Execute executes the request
//
//	@return GetMyFeesEstimateResponse
func (a *FeesAPIService) GetMyFeesEstimateForSKUExecute(r ApiGetMyFeesEstimateForSKURequest) (*GetMyFeesEstimateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMyFeesEstimateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.GetMyFeesEstimateForSKU")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/fees/v0/listings/{SellerSKU}/feesEstimate"
	localVarPath = strings.Replace(localVarPath, "{"+"SellerSKU"+"}", url.PathEscape(parameterValueToString(r.sellerSKU, "sellerSKU")), -1)

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

type ApiGetMyFeesEstimatesRequest struct {
	ctx        context.Context
	ApiService *FeesAPIService
	body       *[]FeesEstimateByIdRequest
}

func (r ApiGetMyFeesEstimatesRequest) Body(body []FeesEstimateByIdRequest) ApiGetMyFeesEstimatesRequest {
	r.body = &body
	return r
}

func (r ApiGetMyFeesEstimatesRequest) Execute() ([]FeesEstimateResult, *http.Response, error) {
	return r.ApiService.GetMyFeesEstimatesExecute(r)
}

/*
GetMyFeesEstimates Method for GetMyFeesEstimates

Returns the estimated fees for a list of products.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMyFeesEstimatesRequest
*/
func (a *FeesAPIService) GetMyFeesEstimates(ctx context.Context) ApiGetMyFeesEstimatesRequest {
	return ApiGetMyFeesEstimatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []FeesEstimateResult
func (a *FeesAPIService) GetMyFeesEstimatesExecute(r ApiGetMyFeesEstimatesRequest) ([]FeesEstimateResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []FeesEstimateResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.GetMyFeesEstimates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/fees/v0/feesEstimate"

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
