package feeds_20210630

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FeedsAPIService FeedsAPI service
type FeedsAPIService service

type ApiCancelFeedRequest struct {
	ctx        context.Context
	ApiService *FeedsAPIService
	feedId     string
}

func (r ApiCancelFeedRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelFeedExecute(r)
}

/*
CancelFeed Method for CancelFeed

Cancels the feed that you specify. Only feeds with `processingStatus=IN_QUEUE` can be cancelled. Cancelled feeds are returned in subsequent calls to the [`getFeed`](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference#getfeed) and [`getFeeds`](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference#getfeeds) operations.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param feedId The identifier for the feed. This identifier is unique only in combination with a seller ID.
	@return ApiCancelFeedRequest
*/
func (a *FeedsAPIService) CancelFeed(ctx context.Context, feedId string) ApiCancelFeedRequest {
	return ApiCancelFeedRequest{
		ApiService: a,
		ctx:        ctx,
		feedId:     feedId,
	}
}

// Execute executes the request
func (a *FeedsAPIService) CancelFeedExecute(r ApiCancelFeedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeedsAPIService.CancelFeed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/2021-06-30/feeds/{feedId}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedId"+"}", url.PathEscape(parameterValueToString(r.feedId, "feedId")), -1)

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

type ApiCreateFeedRequest struct {
	ctx        context.Context
	ApiService *FeedsAPIService
	body       *CreateFeedSpecification
}

// Information required to create the feed.
func (r ApiCreateFeedRequest) Body(body CreateFeedSpecification) ApiCreateFeedRequest {
	r.body = &body
	return r
}

func (r ApiCreateFeedRequest) Execute() (*CreateFeedResponse, *http.Response, error) {
	return r.ApiService.CreateFeedExecute(r)
}

/*
CreateFeed Method for CreateFeed

Creates a feed. Upload the contents of the feed document before calling this operation.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0083 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

The rate limit for the [`JSON_LISTINGS_FEED`](https://developer-docs.amazon.com/sp-api/docs/listings-feed-type-values#listings-feed) feed type differs from the rate limit for the [`createFeed`](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference#post-feeds2021-06-30feeds) operation. For more information, refer to the [Building Listings Management Workflows Guide](https://developer-docs.amazon.com/sp-api/docs/building-listings-management-workflows-guide#should-i-submit-in-bulk-using-the-json_listings_feed-or-individually-with-the-listings-items-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateFeedRequest
*/
func (a *FeedsAPIService) CreateFeed(ctx context.Context) ApiCreateFeedRequest {
	return ApiCreateFeedRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateFeedResponse
func (a *FeedsAPIService) CreateFeedExecute(r ApiCreateFeedRequest) (*CreateFeedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateFeedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeedsAPIService.CreateFeed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/2021-06-30/feeds"

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

type ApiCreateFeedDocumentRequest struct {
	ctx        context.Context
	ApiService *FeedsAPIService
	body       *CreateFeedDocumentSpecification
}

// Specifies the content type for the createFeedDocument operation.
func (r ApiCreateFeedDocumentRequest) Body(body CreateFeedDocumentSpecification) ApiCreateFeedDocumentRequest {
	r.body = &body
	return r
}

func (r ApiCreateFeedDocumentRequest) Execute() (*CreateFeedDocumentResponse, *http.Response, error) {
	return r.ApiService.CreateFeedDocumentExecute(r)
}

/*
CreateFeedDocument Method for CreateFeedDocument

Creates a feed document for the feed type that you specify. This operation returns a presigned URL for uploading the feed document contents. It also returns a `feedDocumentId` value that you can pass in with a subsequent call to the [`createFeed`](https://developer-docs.amazon.com/sp-api/docs/feeds-api-v2021-06-30-reference#createfeed) operation.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateFeedDocumentRequest
*/
func (a *FeedsAPIService) CreateFeedDocument(ctx context.Context) ApiCreateFeedDocumentRequest {
	return ApiCreateFeedDocumentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateFeedDocumentResponse
func (a *FeedsAPIService) CreateFeedDocumentExecute(r ApiCreateFeedDocumentRequest) (*CreateFeedDocumentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateFeedDocumentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeedsAPIService.CreateFeedDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/2021-06-30/documents"

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

type ApiGetFeedRequest struct {
	ctx        context.Context
	ApiService *FeedsAPIService
	feedId     string
}

func (r ApiGetFeedRequest) Execute() (*Feed, *http.Response, error) {
	return r.ApiService.GetFeedExecute(r)
}

/*
GetFeed Method for GetFeed

Returns feed details (including the `resultDocumentId`, if available) for the feed that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param feedId The identifier for the feed. This identifier is unique only in combination with a seller ID.
	@return ApiGetFeedRequest
*/
func (a *FeedsAPIService) GetFeed(ctx context.Context, feedId string) ApiGetFeedRequest {
	return ApiGetFeedRequest{
		ApiService: a,
		ctx:        ctx,
		feedId:     feedId,
	}
}

// Execute executes the request
//
//	@return Feed
func (a *FeedsAPIService) GetFeedExecute(r ApiGetFeedRequest) (*Feed, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Feed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeedsAPIService.GetFeed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/2021-06-30/feeds/{feedId}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedId"+"}", url.PathEscape(parameterValueToString(r.feedId, "feedId")), -1)

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

type ApiGetFeedDocumentRequest struct {
	ctx            context.Context
	ApiService     *FeedsAPIService
	feedDocumentId string
}

func (r ApiGetFeedDocumentRequest) Execute() (*FeedDocument, *http.Response, error) {
	return r.ApiService.GetFeedDocumentExecute(r)
}

/*
GetFeedDocument Method for GetFeedDocument

Returns the information required for retrieving a feed document's contents.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param feedDocumentId The identifier of the feed document.
	@return ApiGetFeedDocumentRequest
*/
func (a *FeedsAPIService) GetFeedDocument(ctx context.Context, feedDocumentId string) ApiGetFeedDocumentRequest {
	return ApiGetFeedDocumentRequest{
		ApiService:     a,
		ctx:            ctx,
		feedDocumentId: feedDocumentId,
	}
}

// Execute executes the request
//
//	@return FeedDocument
func (a *FeedsAPIService) GetFeedDocumentExecute(r ApiGetFeedDocumentRequest) (*FeedDocument, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FeedDocument
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeedsAPIService.GetFeedDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/2021-06-30/documents/{feedDocumentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedDocumentId"+"}", url.PathEscape(parameterValueToString(r.feedDocumentId, "feedDocumentId")), -1)

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

type ApiGetFeedsRequest struct {
	ctx                context.Context
	ApiService         *FeedsAPIService
	feedTypes          *[]string
	marketplaceIds     *[]string
	pageSize           *int32
	processingStatuses *[]string
	createdSince       *time.Time
	createdUntil       *time.Time
	nextToken          *string
}

// A list of feed types used to filter feeds. When feedTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either feedTypes or nextToken is required.
func (r ApiGetFeedsRequest) FeedTypes(feedTypes []string) ApiGetFeedsRequest {
	r.feedTypes = &feedTypes
	return r
}

// A list of marketplace identifiers used to filter feeds. The feeds returned will match at least one of the marketplaces that you specify.
func (r ApiGetFeedsRequest) MarketplaceIds(marketplaceIds []string) ApiGetFeedsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// The maximum number of feeds to return in a single call.
func (r ApiGetFeedsRequest) PageSize(pageSize int32) ApiGetFeedsRequest {
	r.pageSize = &pageSize
	return r
}

// A list of processing statuses used to filter feeds.
func (r ApiGetFeedsRequest) ProcessingStatuses(processingStatuses []string) ApiGetFeedsRequest {
	r.processingStatuses = &processingStatuses
	return r
}

// The earliest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is 90 days ago. Feeds are retained for a maximum of 90 days.
func (r ApiGetFeedsRequest) CreatedSince(createdSince time.Time) ApiGetFeedsRequest {
	r.createdSince = &createdSince
	return r
}

// The latest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is now.
func (r ApiGetFeedsRequest) CreatedUntil(createdUntil time.Time) ApiGetFeedsRequest {
	r.createdUntil = &createdUntil
	return r
}

// A string token returned in the response to your previous request. nextToken is returned when the number of results exceeds the specified pageSize value. To get the next page of results, call the getFeeds operation and include this token as the only parameter. Specifying nextToken with any other parameters will cause the request to fail.
func (r ApiGetFeedsRequest) NextToken(nextToken string) ApiGetFeedsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetFeedsRequest) Execute() (*GetFeedsResponse, *http.Response, error) {
	return r.ApiService.GetFeedsExecute(r)
}

/*
GetFeeds Method for GetFeeds

Returns feed details for the feeds that match the filters that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetFeedsRequest
*/
func (a *FeedsAPIService) GetFeeds(ctx context.Context) ApiGetFeedsRequest {
	return ApiGetFeedsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFeedsResponse
func (a *FeedsAPIService) GetFeedsExecute(r ApiGetFeedsRequest) (*GetFeedsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetFeedsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeedsAPIService.GetFeeds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/2021-06-30/feeds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.feedTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "feedTypes", r.feedTypes, "form", "csv")
	}
	if r.marketplaceIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "", "")
	}
	if r.processingStatuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processingStatuses", r.processingStatuses, "form", "csv")
	}
	if r.createdSince != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdSince", r.createdSince, "", "")
	}
	if r.createdUntil != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdUntil", r.createdUntil, "", "")
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
