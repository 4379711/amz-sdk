package orders_20260101

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// OrdersAPIService OrdersAPI service for the Orders v2026-01-01 API
type OrdersAPIService service

type ApiSearchOrdersRequest struct {
	ctx                 context.Context
	ApiService          *OrdersAPIService
	createdAfter        *time.Time
	createdBefore       *time.Time
	lastUpdatedAfter    *time.Time
	lastUpdatedBefore   *time.Time
	fulfillmentStatuses *[]string
	marketplaceIds      *[]string
	fulfilledBy         *[]string
	maxResultsPerPage   *int32
	paginationToken     *string
	includedData        *[]string
}

// The response includes orders created at or after this time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: You must provide exactly one of &#x60;createdAfter&#x60; and &#x60;lastUpdatedAfter&#x60; in your request. If &#x60;createdAfter&#x60; is provided, neither &#x60;lastUpdatedAfter&#x60; nor &#x60;lastUpdatedBefore&#x60; may be provided.
func (r ApiSearchOrdersRequest) CreatedAfter(createdAfter time.Time) ApiSearchOrdersRequest {
	r.createdAfter = &createdAfter
	return r
}

// The response includes orders created at or before this time. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: If you include &#x60;createdAfter&#x60; in the request, &#x60;createdBefore&#x60; is optional, and if provided must be equal to or after the &#x60;createdAfter&#x60; date and at least two minutes before the time of the request. If &#x60;createdBefore&#x60; is provided, neither &#x60;lastUpdatedAfter&#x60; nor &#x60;lastUpdatedBefore&#x60; may be provided.
func (r ApiSearchOrdersRequest) CreatedBefore(createdBefore time.Time) ApiSearchOrdersRequest {
	r.createdBefore = &createdBefore
	return r
}

// The response includes orders updated at or after this time. An update is any change made by Amazon or the seller, including changes to order status. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: You must provide exactly one of &#x60;createdAfter&#x60; and &#x60;lastUpdatedAfter&#x60;. If &#x60;lastUpdatedAfter&#x60; is provided, neither &#x60;createdAfter&#x60; nor &#x60;createdBefore&#x60; may be provided.
func (r ApiSearchOrdersRequest) LastUpdatedAfter(lastUpdatedAfter time.Time) ApiSearchOrdersRequest {
	r.lastUpdatedAfter = &lastUpdatedAfter
	return r
}

// The response includes orders updated at or before this time. An update is any change made by Amazon or the seller, including changes to order status. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.  **Note**: If you include &#x60;lastUpdatedAfter&#x60; in the request, &#x60;lastUpdatedBefore&#x60; is optional, and if provided must be equal to or after the &#x60;lastUpdatedAfter&#x60; date and at least two minutes before the time of the request. If &#x60;lastUpdatedBefore&#x60; is provided, neither &#x60;createdAfter&#x60; nor &#x60;createdBefore&#x60; may be provided.
func (r ApiSearchOrdersRequest) LastUpdatedBefore(lastUpdatedBefore time.Time) ApiSearchOrdersRequest {
	r.lastUpdatedBefore = &lastUpdatedBefore
	return r
}

// A list of &#x60;FulfillmentStatus&#x60; values you can use to filter the results.
func (r ApiSearchOrdersRequest) FulfillmentStatuses(fulfillmentStatuses []string) ApiSearchOrdersRequest {
	r.fulfillmentStatuses = &fulfillmentStatuses
	return r
}

// The response includes orders that were placed in marketplaces you include in this list.  Refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids) for a complete list of &#x60;marketplaceId&#x60; values.
func (r ApiSearchOrdersRequest) MarketplaceIds(marketplaceIds []string) ApiSearchOrdersRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// The response includes orders that are fulfilled by the parties that you include in this list.
func (r ApiSearchOrdersRequest) FulfilledBy(fulfilledBy []string) ApiSearchOrdersRequest {
	r.fulfilledBy = &fulfilledBy
	return r
}

// The maximum number of orders that can be returned per page. The value must be between 1 and 100. **Default:** 100.
func (r ApiSearchOrdersRequest) MaxResultsPerPage(maxResultsPerPage int32) ApiSearchOrdersRequest {
	r.maxResultsPerPage = &maxResultsPerPage
	return r
}

// Pagination occurs when a request produces a response that exceeds the &#x60;maxResultsPerPage&#x60;. This means that the response is divided into individual pages. To retrieve the next page, you must pass the &#x60;nextToken&#x60; value as the &#x60;paginationToken&#x60; query parameter in the next request. You will not receive a &#x60;nextToken&#x60; value on the last page.
func (r ApiSearchOrdersRequest) PaginationToken(paginationToken string) ApiSearchOrdersRequest {
	r.paginationToken = &paginationToken
	return r
}

// A list of datasets to include in the response.
func (r ApiSearchOrdersRequest) IncludedData(includedData []string) ApiSearchOrdersRequest {
	r.includedData = &includedData
	return r
}

func (r ApiSearchOrdersRequest) Execute() (*SearchOrdersResponse, *http.Response, error) {
	return r.ApiService.SearchOrdersExecute(r)
}

/*
SearchOrders Method for SearchOrders

Returns orders created or updated during the time period that you specify. You can filter the response for specific types of orders.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0056 | 20 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api) in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchOrdersRequest
*/
func (a *OrdersAPIService) SearchOrders(ctx context.Context) ApiSearchOrdersRequest {
	return ApiSearchOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchOrdersResponse
func (a *OrdersAPIService) SearchOrdersExecute(r ApiSearchOrdersRequest) (*SearchOrdersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchOrdersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersAPIService.SearchOrders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/2026-01-01/orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createdAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdAfter", r.createdAfter, "", "")
	}
	if r.createdBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdBefore", r.createdBefore, "", "")
	}
	if r.lastUpdatedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastUpdatedAfter", r.lastUpdatedAfter, "", "")
	}
	if r.lastUpdatedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastUpdatedBefore", r.lastUpdatedBefore, "", "")
	}
	if r.fulfillmentStatuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fulfillmentStatuses", r.fulfillmentStatuses, "form", "csv")
	}
	if r.marketplaceIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	}
	if r.fulfilledBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fulfilledBy", r.fulfilledBy, "form", "csv")
	}
	if r.maxResultsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxResultsPerPage", r.maxResultsPerPage, "", "")
	}
	if r.paginationToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paginationToken", r.paginationToken, "", "")
	}
	if r.includedData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includedData", r.includedData, "form", "csv")
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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrderRequest struct {
	ctx          context.Context
	ApiService   *OrdersAPIService
	orderId      string
	includedData *[]string
}

// A list of datasets to include in the response.
func (r ApiGetOrderRequest) IncludedData(includedData []string) ApiGetOrderRequest {
	r.includedData = &includedData
	return r
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

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api) in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier.
	@return ApiGetOrderRequest
*/
func (a *OrdersAPIService) GetOrder(ctx context.Context, orderId string) ApiGetOrderRequest {
	return ApiGetOrderRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return GetOrderResponse
func (a *OrdersAPIService) GetOrderExecute(r ApiGetOrderRequest) (*GetOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersAPIService.GetOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/2026-01-01/orders/{orderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includedData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includedData", r.includedData, "form", "csv")
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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
