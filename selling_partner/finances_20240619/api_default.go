package finances_20240619

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiListTransactionsRequest struct {
	ctx           context.Context
	ApiService    *DefaultAPIService
	postedAfter   *time.Time
	postedBefore  *time.Time
	marketplaceId *string
	nextToken     *string
}

// A date used for selecting transactions posted after (or at) a specified time. The date-time must be no later than two minutes before the request was submitted, in ISO 8601 date time format.
func (r ApiListTransactionsRequest) PostedAfter(postedAfter time.Time) ApiListTransactionsRequest {
	r.postedAfter = &postedAfter
	return r
}

// A date used for selecting transactions posted before (but not at) a specified time. The date-time must be later than PostedAfter and no later than two minutes before the request was submitted, in ISO 8601 date time format. If PostedAfter and PostedBefore are more than 180 days apart, no transactions are returned. You must specify the PostedAfter parameter if you specify the PostedBefore parameter. Default: Now minus two minutes.
func (r ApiListTransactionsRequest) PostedBefore(postedBefore time.Time) ApiListTransactionsRequest {
	r.postedBefore = &postedBefore
	return r
}

// A string token used to select Marketplace ID.
func (r ApiListTransactionsRequest) MarketplaceId(marketplaceId string) ApiListTransactionsRequest {
	r.marketplaceId = &marketplaceId
	return r
}

// A string token returned in the response of your previous request.
func (r ApiListTransactionsRequest) NextToken(nextToken string) ApiListTransactionsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListTransactionsRequest) Execute() (*ListTransactionsResponseWrapper, *http.Response, error) {
	return r.ApiService.ListTransactionsExecute(r)
}

/*
ListTransactions Method for ListTransactions

Returns transactions for the given parameters. It may take up to 48 hours for transactions to appear in your transaction events.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListTransactionsRequest
*/
func (a *DefaultAPIService) ListTransactions(ctx context.Context) ApiListTransactionsRequest {
	return ApiListTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListTransactionsResponse
func (a *DefaultAPIService) ListTransactionsExecute(r ApiListTransactionsRequest) (*ListTransactionsResponseWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListTransactionsResponseWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ListTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/2024-06-19/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postedAfter == nil {
		return localVarReturnValue, nil, reportError("postedAfter is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "postedAfter", r.postedAfter, "", "")
	if r.postedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "postedBefore", r.postedBefore, "", "")
	}
	if r.marketplaceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceId", r.marketplaceId, "", "")
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
