package finances_v0

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiListFinancialEventGroupsRequest struct {
	ctx                              context.Context
	ApiService                       *DefaultAPIService
	maxResultsPerPage                *int32
	financialEventGroupStartedBefore *time.Time
	financialEventGroupStartedAfter  *time.Time
	nextToken                        *string
}

// The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#39;InvalidInput&#39;.
func (r ApiListFinancialEventGroupsRequest) MaxResultsPerPage(maxResultsPerPage int32) ApiListFinancialEventGroupsRequest {
	r.maxResultsPerPage = &maxResultsPerPage
	return r
}

// A date used for selecting financial event groups that opened before (but not at) a specified date and time, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. The date-time  must be later than FinancialEventGroupStartedAfter and no later than two minutes before the request was submitted. If FinancialEventGroupStartedAfter and FinancialEventGroupStartedBefore are more than 180 days apart, no financial event groups are returned.
func (r ApiListFinancialEventGroupsRequest) FinancialEventGroupStartedBefore(financialEventGroupStartedBefore time.Time) ApiListFinancialEventGroupsRequest {
	r.financialEventGroupStartedBefore = &financialEventGroupStartedBefore
	return r
}

// A date used for selecting financial event groups that opened after (or at) a specified date and time, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format. The date-time must be no later than two minutes before the request was submitted.
func (r ApiListFinancialEventGroupsRequest) FinancialEventGroupStartedAfter(financialEventGroupStartedAfter time.Time) ApiListFinancialEventGroupsRequest {
	r.financialEventGroupStartedAfter = &financialEventGroupStartedAfter
	return r
}

// A string token returned in the response of your previous request.
func (r ApiListFinancialEventGroupsRequest) NextToken(nextToken string) ApiListFinancialEventGroupsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListFinancialEventGroupsRequest) Execute() (*ListFinancialEventGroupsResponse, *http.Response, error) {
	return r.ApiService.ListFinancialEventGroupsExecute(r)
}

/*
ListFinancialEventGroups Method for ListFinancialEventGroups

Returns financial event groups for a given date range. It may take up to 48 hours for orders to appear in your financial events.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListFinancialEventGroupsRequest
*/
func (a *DefaultAPIService) ListFinancialEventGroups(ctx context.Context) ApiListFinancialEventGroupsRequest {
	return ApiListFinancialEventGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListFinancialEventGroupsResponse
func (a *DefaultAPIService) ListFinancialEventGroupsExecute(r ApiListFinancialEventGroupsRequest) (*ListFinancialEventGroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListFinancialEventGroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ListFinancialEventGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/v0/financialEventGroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxResultsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxResultsPerPage", r.maxResultsPerPage, "", "")
	}
	if r.financialEventGroupStartedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "FinancialEventGroupStartedBefore", r.financialEventGroupStartedBefore, "", "")
	}
	if r.financialEventGroupStartedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "FinancialEventGroupStartedAfter", r.financialEventGroupStartedAfter, "", "")
	}
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

type ApiListFinancialEventsRequest struct {
	ctx               context.Context
	ApiService        *DefaultAPIService
	maxResultsPerPage *int32
	postedAfter       *time.Time
	postedBefore      *time.Time
	nextToken         *string
}

// The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#39;InvalidInput&#39;.
func (r ApiListFinancialEventsRequest) MaxResultsPerPage(maxResultsPerPage int32) ApiListFinancialEventsRequest {
	r.maxResultsPerPage = &maxResultsPerPage
	return r
}

// A date used for selecting financial events posted after (or at) a specified time. The date-time must be no later than two minutes before the request was submitted, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date time format.
func (r ApiListFinancialEventsRequest) PostedAfter(postedAfter time.Time) ApiListFinancialEventsRequest {
	r.postedAfter = &postedAfter
	return r
}

// A date used for selecting financial events posted before (but not at) a specified time. The date-time must be later than PostedAfter and no later than two minutes before the request was submitted, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date time format. If PostedAfter and PostedBefore are more than 180 days apart, no financial events are returned. You must specify the PostedAfter parameter if you specify the PostedBefore parameter. Default: Now minus two minutes.
func (r ApiListFinancialEventsRequest) PostedBefore(postedBefore time.Time) ApiListFinancialEventsRequest {
	r.postedBefore = &postedBefore
	return r
}

// A string token returned in the response of your previous request.
func (r ApiListFinancialEventsRequest) NextToken(nextToken string) ApiListFinancialEventsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListFinancialEventsRequest) Execute() (*ListFinancialEventsResponse, *http.Response, error) {
	return r.ApiService.ListFinancialEventsExecute(r)
}

/*
ListFinancialEvents Method for ListFinancialEvents

Returns financial events for the specified data range. It may take up to 48 hours for orders to appear in your financial events. **Note:** in `ListFinancialEvents`, deferred events don't show up in responses until in they are released.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListFinancialEventsRequest
*/
func (a *DefaultAPIService) ListFinancialEvents(ctx context.Context) ApiListFinancialEventsRequest {
	return ApiListFinancialEventsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListFinancialEventsResponse
func (a *DefaultAPIService) ListFinancialEventsExecute(r ApiListFinancialEventsRequest) (*ListFinancialEventsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListFinancialEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ListFinancialEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/v0/financialEvents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxResultsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxResultsPerPage", r.maxResultsPerPage, "", "")
	}
	if r.postedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PostedAfter", r.postedAfter, "", "")
	}
	if r.postedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PostedBefore", r.postedBefore, "", "")
	}
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

type ApiListFinancialEventsByGroupIdRequest struct {
	ctx               context.Context
	ApiService        *DefaultAPIService
	eventGroupId      string
	maxResultsPerPage *int32
	postedAfter       *time.Time
	postedBefore      *time.Time
	nextToken         *string
}

// The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#39;InvalidInput&#39;.
func (r ApiListFinancialEventsByGroupIdRequest) MaxResultsPerPage(maxResultsPerPage int32) ApiListFinancialEventsByGroupIdRequest {
	r.maxResultsPerPage = &maxResultsPerPage
	return r
}

// A date used for selecting financial events posted after (or at) a specified time. The date-time **must** be more than two minutes before the time of the request, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date time format.
func (r ApiListFinancialEventsByGroupIdRequest) PostedAfter(postedAfter time.Time) ApiListFinancialEventsByGroupIdRequest {
	r.postedAfter = &postedAfter
	return r
}

// A date used for selecting financial events posted before (but not at) a specified time. The date-time must be later than &#x60;PostedAfter&#x60; and no later than two minutes before the request was submitted, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date time format. If &#x60;PostedAfter&#x60; and &#x60;PostedBefore&#x60; are more than 180 days apart, no financial events are returned. You must specify the &#x60;PostedAfter&#x60; parameter if you specify the &#x60;PostedBefore&#x60; parameter. Default: Now minus two minutes.
func (r ApiListFinancialEventsByGroupIdRequest) PostedBefore(postedBefore time.Time) ApiListFinancialEventsByGroupIdRequest {
	r.postedBefore = &postedBefore
	return r
}

// A string token returned in the response of your previous request.
func (r ApiListFinancialEventsByGroupIdRequest) NextToken(nextToken string) ApiListFinancialEventsByGroupIdRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListFinancialEventsByGroupIdRequest) Execute() (*ListFinancialEventsResponse, *http.Response, error) {
	return r.ApiService.ListFinancialEventsByGroupIdExecute(r)
}

/*
ListFinancialEventsByGroupId Method for ListFinancialEventsByGroupId

Returns all financial events for the specified financial event group. It may take up to 48 hours for orders to appear in your financial events.

**Note:** This operation will only retrieve group's data for the past two years. If a request is submitted for data spanning more than two years, an empty response is returned.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventGroupId The identifier of the financial event group to which the events belong.
	@return ApiListFinancialEventsByGroupIdRequest
*/
func (a *DefaultAPIService) ListFinancialEventsByGroupId(ctx context.Context, eventGroupId string) ApiListFinancialEventsByGroupIdRequest {
	return ApiListFinancialEventsByGroupIdRequest{
		ApiService:   a,
		ctx:          ctx,
		eventGroupId: eventGroupId,
	}
}

// Execute executes the request
//
//	@return ListFinancialEventsResponse
func (a *DefaultAPIService) ListFinancialEventsByGroupIdExecute(r ApiListFinancialEventsByGroupIdRequest) (*ListFinancialEventsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListFinancialEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ListFinancialEventsByGroupId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/v0/financialEventGroups/{eventGroupId}/financialEvents"
	localVarPath = strings.Replace(localVarPath, "{"+"eventGroupId"+"}", url.PathEscape(parameterValueToString(r.eventGroupId, "eventGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxResultsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxResultsPerPage", r.maxResultsPerPage, "", "")
	}
	if r.postedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PostedAfter", r.postedAfter, "", "")
	}
	if r.postedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PostedBefore", r.postedBefore, "", "")
	}
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

type ApiListFinancialEventsByOrderIdRequest struct {
	ctx               context.Context
	ApiService        *DefaultAPIService
	orderId           string
	maxResultsPerPage *int32
	nextToken         *string
}

// The maximum number of results to return per page. If the response exceeds the maximum number of transactions or 10 MB, the API responds with &#39;InvalidInput&#39;.
func (r ApiListFinancialEventsByOrderIdRequest) MaxResultsPerPage(maxResultsPerPage int32) ApiListFinancialEventsByOrderIdRequest {
	r.maxResultsPerPage = &maxResultsPerPage
	return r
}

// A string token returned in the response of your previous request.
func (r ApiListFinancialEventsByOrderIdRequest) NextToken(nextToken string) ApiListFinancialEventsByOrderIdRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListFinancialEventsByOrderIdRequest) Execute() (*ListFinancialEventsResponse, *http.Response, error) {
	return r.ApiService.ListFinancialEventsByOrderIdExecute(r)
}

/*
ListFinancialEventsByOrderId Method for ListFinancialEventsByOrderId

Returns all financial events for the specified order. It may take up to 48 hours for orders to appear in your financial events.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiListFinancialEventsByOrderIdRequest
*/
func (a *DefaultAPIService) ListFinancialEventsByOrderId(ctx context.Context, orderId string) ApiListFinancialEventsByOrderIdRequest {
	return ApiListFinancialEventsByOrderIdRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
//
//	@return ListFinancialEventsResponse
func (a *DefaultAPIService) ListFinancialEventsByOrderIdExecute(r ApiListFinancialEventsByOrderIdRequest) (*ListFinancialEventsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListFinancialEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ListFinancialEventsByOrderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/v0/orders/{orderId}/financialEvents"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxResultsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxResultsPerPage", r.maxResultsPerPage, "", "")
	}
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
