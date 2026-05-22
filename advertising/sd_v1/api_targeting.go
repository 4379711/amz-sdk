package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// TargetingAPIService TargetingAPI service
type TargetingAPIService service

type ApiArchiveTargetingClauseRequest struct {
	ctx                          context.Context
	ApiService                   *TargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	targetId                     int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiArchiveTargetingClauseRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiArchiveTargetingClauseRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiArchiveTargetingClauseRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiArchiveTargetingClauseRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiArchiveTargetingClauseRequest) Execute() (*TargetResponse, *http.Response, error) {
	return r.ApiService.ArchiveTargetingClauseExecute(r)
}

func (a *TargetingAPIService) ArchiveTargetingClause(ctx context.Context, targetId int64) ApiArchiveTargetingClauseRequest {
	return ApiArchiveTargetingClauseRequest{
		ApiService: a,
		ctx:        ctx,
		targetId:   targetId,
	}
}

// Execute executes the request
//
//	@return TargetResponse
func (a *TargetingAPIService) ArchiveTargetingClauseExecute(r ApiArchiveTargetingClauseRequest) (*TargetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TargetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingAPIService.ArchiveTargetingClause")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets/{targetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetId"+"}", url.PathEscape(parameterValueToString(r.targetId, "targetId")), -1)

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

type ApiCreateTargetingClausesRequest struct {
	ctx                          context.Context
	ApiService                   *TargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createTargetingClause        *[]CreateTargetingClause
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of targeting clauses for creation.
func (r ApiCreateTargetingClausesRequest) CreateTargetingClause(createTargetingClause []CreateTargetingClause) ApiCreateTargetingClausesRequest {
	r.createTargetingClause = &createTargetingClause
	return r
}

func (r ApiCreateTargetingClausesRequest) Execute() ([]TargetResponse, *http.Response, error) {
	return r.ApiService.CreateTargetingClausesExecute(r)
}

func (a *TargetingAPIService) CreateTargetingClauses(ctx context.Context) ApiCreateTargetingClausesRequest {
	return ApiCreateTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TargetResponse
func (a *TargetingAPIService) CreateTargetingClausesExecute(r ApiCreateTargetingClausesRequest) ([]TargetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TargetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingAPIService.CreateTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.createTargetingClause
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

type ApiGetTargetsRequest struct {
	ctx                          context.Context
	ApiService                   *TargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	targetId                     int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetTargetsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetTargetsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetTargetsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetTargetsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetTargetsRequest) Execute() (*TargetingClause, *http.Response, error) {
	return r.ApiService.GetTargetsExecute(r)
}

func (a *TargetingAPIService) GetTargets(ctx context.Context, targetId int64) ApiGetTargetsRequest {
	return ApiGetTargetsRequest{
		ApiService: a,
		ctx:        ctx,
		targetId:   targetId,
	}
}

// Execute executes the request
//
//	@return TargetingClause
func (a *TargetingAPIService) GetTargetsExecute(r ApiGetTargetsRequest) (*TargetingClause, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TargetingClause
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingAPIService.GetTargets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets/{targetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetId"+"}", url.PathEscape(parameterValueToString(r.targetId, "targetId")), -1)

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

type ApiGetTargetsExRequest struct {
	ctx                          context.Context
	ApiService                   *TargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	targetId                     int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetTargetsExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetTargetsExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetTargetsExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetTargetsExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetTargetsExRequest) Execute() (*TargetingClauseEx, *http.Response, error) {
	return r.ApiService.GetTargetsExExecute(r)
}

func (a *TargetingAPIService) GetTargetsEx(ctx context.Context, targetId int64) ApiGetTargetsExRequest {
	return ApiGetTargetsExRequest{
		ApiService: a,
		ctx:        ctx,
		targetId:   targetId,
	}
}

// Execute executes the request
//
//	@return TargetingClauseEx
func (a *TargetingAPIService) GetTargetsExExecute(r ApiGetTargetsExRequest) (*TargetingClauseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TargetingClauseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingAPIService.GetTargetsEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets/extended/{targetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetId"+"}", url.PathEscape(parameterValueToString(r.targetId, "targetId")), -1)

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

type ApiListTargetingClausesRequest struct {
	ctx                          context.Context
	ApiService                   *TargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	targetIdFilter               *string
	adGroupIdFilter              *string
	campaignIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. 0-indexed record offset for the result set. Defaults to 0.
func (r ApiListTargetingClausesRequest) StartIndex(startIndex int32) ApiListTargetingClausesRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Number of records to include in the paged response. Defaults to max page size.
func (r ApiListTargetingClausesRequest) Count(count int32) ApiListTargetingClausesRequest {
	r.count = &count
	return r
}

// Optional. Restricts results to those with &#x60;state&#x60; set to values in the specified comma-separated list.
func (r ApiListTargetingClausesRequest) StateFilter(stateFilter string) ApiListTargetingClausesRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. Restricts results to ads with the specified &#x60;tagetId&#x60; specified in comma-separated list
func (r ApiListTargetingClausesRequest) TargetIdFilter(targetIdFilter string) ApiListTargetingClausesRequest {
	r.targetIdFilter = &targetIdFilter
	return r
}

// Optional list of comma separated adGroupIds. Restricts results to targeting clauses with the specified &#x60;adGroupId&#x60;.
func (r ApiListTargetingClausesRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListTargetingClausesRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. Restricts results to targeting clauses within campaigns specified in comma-separated list.
func (r ApiListTargetingClausesRequest) CampaignIdFilter(campaignIdFilter string) ApiListTargetingClausesRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListTargetingClausesRequest) Execute() ([]TargetingClause, *http.Response, error) {
	return r.ApiService.ListTargetingClausesExecute(r)
}

func (a *TargetingAPIService) ListTargetingClauses(ctx context.Context) ApiListTargetingClausesRequest {
	return ApiListTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TargetingClause
func (a *TargetingAPIService) ListTargetingClausesExecute(r ApiListTargetingClausesRequest) ([]TargetingClause, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TargetingClause
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingAPIService.ListTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.stateFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stateFilter", r.stateFilter, "form", "")
	} else {
		var defaultValue string = "enabled, paused, archived"
		r.stateFilter = &defaultValue
	}
	if r.targetIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "targetIdFilter", r.targetIdFilter, "form", "")
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
	}
	if r.campaignIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaignIdFilter", r.campaignIdFilter, "form", "")
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

type ApiListTargetingClausesExRequest struct {
	ctx                          context.Context
	ApiService                   *TargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	targetIdFilter               *string
	adGroupIdFilter              *string
	campaignIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListTargetingClausesExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListTargetingClausesExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListTargetingClausesExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListTargetingClausesExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. 0-indexed record offset for the result set. Defaults to 0.
func (r ApiListTargetingClausesExRequest) StartIndex(startIndex int32) ApiListTargetingClausesExRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Number of records to include in the paged response. Defaults to max page size.
func (r ApiListTargetingClausesExRequest) Count(count int32) ApiListTargetingClausesExRequest {
	r.count = &count
	return r
}

// Optional. Restricts results to keywords with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;, &#x60;paused&#x60;, or &#x60;archived&#x60;. Default behavior is to include enabled, paused, and archived.
func (r ApiListTargetingClausesExRequest) StateFilter(stateFilter string) ApiListTargetingClausesExRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. Restricts results to ads with the specified &#x60;tagetId&#x60; specified in comma-separated list
func (r ApiListTargetingClausesExRequest) TargetIdFilter(targetIdFilter string) ApiListTargetingClausesExRequest {
	r.targetIdFilter = &targetIdFilter
	return r
}

// Optional list of comma separated adGroupIds. Restricts results to targeting clauses with the specified &#x60;adGroupId&#x60;.
func (r ApiListTargetingClausesExRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListTargetingClausesExRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. Restricts results to ads within campaigns specified in comma-separated list.
func (r ApiListTargetingClausesExRequest) CampaignIdFilter(campaignIdFilter string) ApiListTargetingClausesExRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListTargetingClausesExRequest) Execute() ([]TargetingClauseEx, *http.Response, error) {
	return r.ApiService.ListTargetingClausesExExecute(r)
}

func (a *TargetingAPIService) ListTargetingClausesEx(ctx context.Context) ApiListTargetingClausesExRequest {
	return ApiListTargetingClausesExRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TargetingClauseEx
func (a *TargetingAPIService) ListTargetingClausesExExecute(r ApiListTargetingClausesExRequest) ([]TargetingClauseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TargetingClauseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingAPIService.ListTargetingClausesEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets/extended"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.stateFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stateFilter", r.stateFilter, "form", "")
	} else {
		var defaultValue string = "enabled, paused, archived"
		r.stateFilter = &defaultValue
	}
	if r.targetIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "targetIdFilter", r.targetIdFilter, "form", "")
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
	}
	if r.campaignIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaignIdFilter", r.campaignIdFilter, "form", "")
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

type ApiUpdateTargetingClausesRequest struct {
	ctx                          context.Context
	ApiService                   *TargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateTargetingClause        *[]UpdateTargetingClause
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of up to 1000 targeting clauses. Mutable fields: * &#x60;state&#x60; * &#x60;bid&#x60; (only mutable when the targeting clause&#39;s adGroup does not have any enabled optimization rule)
func (r ApiUpdateTargetingClausesRequest) UpdateTargetingClause(updateTargetingClause []UpdateTargetingClause) ApiUpdateTargetingClausesRequest {
	r.updateTargetingClause = &updateTargetingClause
	return r
}

func (r ApiUpdateTargetingClausesRequest) Execute() ([]TargetResponse, *http.Response, error) {
	return r.ApiService.UpdateTargetingClausesExecute(r)
}

func (a *TargetingAPIService) UpdateTargetingClauses(ctx context.Context) ApiUpdateTargetingClausesRequest {
	return ApiUpdateTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TargetResponse
func (a *TargetingAPIService) UpdateTargetingClausesExecute(r ApiUpdateTargetingClausesRequest) ([]TargetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TargetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingAPIService.UpdateTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.updateTargetingClause
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
