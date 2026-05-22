package vendor_direct_fulfillment_shipping_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VendorShippingLabelsAPIService VendorShippingLabelsAPI service
type VendorShippingLabelsAPIService service

type ApiGetShippingLabelRequest struct {
	ctx                 context.Context
	ApiService          *VendorShippingLabelsAPIService
	purchaseOrderNumber string
}

func (r ApiGetShippingLabelRequest) Execute() (*GetShippingLabelResponse, *http.Response, error) {
	return r.ApiService.GetShippingLabelExecute(r)
}

/*
GetShippingLabel Method for GetShippingLabel

Returns a shipping label for the purchaseOrderNumber that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param purchaseOrderNumber The purchase order number for which you want to return the shipping label. It should be the same purchaseOrderNumber as received in the order.
	@return ApiGetShippingLabelRequest
*/
func (a *VendorShippingLabelsAPIService) GetShippingLabel(ctx context.Context, purchaseOrderNumber string) ApiGetShippingLabelRequest {
	return ApiGetShippingLabelRequest{
		ApiService:          a,
		ctx:                 ctx,
		purchaseOrderNumber: purchaseOrderNumber,
	}
}

// Execute executes the request
//
//	@return GetShippingLabelResponse
func (a *VendorShippingLabelsAPIService) GetShippingLabelExecute(r ApiGetShippingLabelRequest) (*GetShippingLabelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetShippingLabelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorShippingLabelsAPIService.GetShippingLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor/directFulfillment/shipping/v1/shippingLabels/{purchaseOrderNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"purchaseOrderNumber"+"}", url.PathEscape(parameterValueToString(r.purchaseOrderNumber, "purchaseOrderNumber")), -1)

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

type ApiGetShippingLabelsRequest struct {
	ctx             context.Context
	ApiService      *VendorShippingLabelsAPIService
	createdAfter    *time.Time
	createdBefore   *time.Time
	shipFromPartyId *string
	limit           *int32
	sortOrder       *string
	nextToken       *string
}

// Shipping labels that became available after this date and time will be included in the result. Must be in ISO-8601 date/time format.
func (r ApiGetShippingLabelsRequest) CreatedAfter(createdAfter time.Time) ApiGetShippingLabelsRequest {
	r.createdAfter = &createdAfter
	return r
}

// Shipping labels that became available before this date and time will be included in the result. Must be in ISO-8601 date/time format.
func (r ApiGetShippingLabelsRequest) CreatedBefore(createdBefore time.Time) ApiGetShippingLabelsRequest {
	r.createdBefore = &createdBefore
	return r
}

// The vendor warehouseId for order fulfillment. If not specified, the result will contain orders for all warehouses.
func (r ApiGetShippingLabelsRequest) ShipFromPartyId(shipFromPartyId string) ApiGetShippingLabelsRequest {
	r.shipFromPartyId = &shipFromPartyId
	return r
}

// The limit to the number of records returned.
func (r ApiGetShippingLabelsRequest) Limit(limit int32) ApiGetShippingLabelsRequest {
	r.limit = &limit
	return r
}

// Sort ASC or DESC by order creation date.
func (r ApiGetShippingLabelsRequest) SortOrder(sortOrder string) ApiGetShippingLabelsRequest {
	r.sortOrder = &sortOrder
	return r
}

// Used for pagination when there are more ship labels than the specified result size limit. The token value is returned in the previous API call.
func (r ApiGetShippingLabelsRequest) NextToken(nextToken string) ApiGetShippingLabelsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetShippingLabelsRequest) Execute() (*GetShippingLabelListResponse, *http.Response, error) {
	return r.ApiService.GetShippingLabelsExecute(r)
}

/*
GetShippingLabels Method for GetShippingLabels

Returns a list of shipping labels created during the time frame that you specify. You define that time frame using the createdAfter and createdBefore parameters. You must use both of these parameters. The date range to search must not be more than 7 days.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetShippingLabelsRequest
*/
func (a *VendorShippingLabelsAPIService) GetShippingLabels(ctx context.Context) ApiGetShippingLabelsRequest {
	return ApiGetShippingLabelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetShippingLabelListResponse
func (a *VendorShippingLabelsAPIService) GetShippingLabelsExecute(r ApiGetShippingLabelsRequest) (*GetShippingLabelListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetShippingLabelListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorShippingLabelsAPIService.GetShippingLabels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor/directFulfillment/shipping/v1/shippingLabels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createdAfter == nil {
		return localVarReturnValue, nil, reportError("createdAfter is required and must be specified")
	}
	if r.createdBefore == nil {
		return localVarReturnValue, nil, reportError("createdBefore is required and must be specified")
	}

	if r.shipFromPartyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shipFromPartyId", r.shipFromPartyId, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "createdAfter", r.createdAfter, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "createdBefore", r.createdBefore, "", "")
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", r.sortOrder, "", "")
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
	localVarHTTPHeaderAccepts := []string{"application/json", "payload"}

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

type ApiSubmitShippingLabelRequestRequest struct {
	ctx        context.Context
	ApiService *VendorShippingLabelsAPIService
	body       *SubmitShippingLabelsRequest
}

// Request body containing one or more shipping labels data.
func (r ApiSubmitShippingLabelRequestRequest) Body(body SubmitShippingLabelsRequest) ApiSubmitShippingLabelRequestRequest {
	r.body = &body
	return r
}

func (r ApiSubmitShippingLabelRequestRequest) Execute() (*SubmitShippingLabelsResponse, *http.Response, error) {
	return r.ApiService.SubmitShippingLabelRequestExecute(r)
}

/*
SubmitShippingLabelRequest Method for SubmitShippingLabelRequest

Creates a shipping label for a purchase order and returns a transactionId for reference.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSubmitShippingLabelRequestRequest
*/
func (a *VendorShippingLabelsAPIService) SubmitShippingLabelRequest(ctx context.Context) ApiSubmitShippingLabelRequestRequest {
	return ApiSubmitShippingLabelRequestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubmitShippingLabelsResponse
func (a *VendorShippingLabelsAPIService) SubmitShippingLabelRequestExecute(r ApiSubmitShippingLabelRequestRequest) (*SubmitShippingLabelsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubmitShippingLabelsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorShippingLabelsAPIService.SubmitShippingLabelRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor/directFulfillment/shipping/v1/shippingLabels"

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
