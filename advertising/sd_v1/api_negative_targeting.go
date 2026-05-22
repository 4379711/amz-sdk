package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// NegativeTargetingAPIService NegativeTargetingAPI service
type NegativeTargetingAPIService service

type ApiArchiveNegativeTargetingClauseRequest struct {
	ctx                          context.Context
	ApiService                   *NegativeTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	negativeTargetId             int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiArchiveNegativeTargetingClauseRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiArchiveNegativeTargetingClauseRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiArchiveNegativeTargetingClauseRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiArchiveNegativeTargetingClauseRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiArchiveNegativeTargetingClauseRequest) Execute() (*TargetResponse, *http.Response, error) {
	return r.ApiService.ArchiveNegativeTargetingClauseExecute(r)
}

func (a *NegativeTargetingAPIService) ArchiveNegativeTargetingClause(ctx context.Context, negativeTargetId int64) ApiArchiveNegativeTargetingClauseRequest {
	return ApiArchiveNegativeTargetingClauseRequest{
		ApiService:       a,
		ctx:              ctx,
		negativeTargetId: negativeTargetId,
	}
}

// Execute executes the request
//
//	@return TargetResponse
func (a *NegativeTargetingAPIService) ArchiveNegativeTargetingClauseExecute(r ApiArchiveNegativeTargetingClauseRequest) (*TargetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TargetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingAPIService.ArchiveNegativeTargetingClause")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/negativeTargets/{negativeTargetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"negativeTargetId"+"}", url.PathEscape(parameterValueToString(r.negativeTargetId, "negativeTargetId")), -1)

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

type ApiCreateNegativeTargetingClausesRequest struct {
	ctx                           context.Context
	ApiService                    *NegativeTargetingAPIService
	amazonAdvertisingAPIClientId  *string
	amazonAdvertisingAPIScope     *string
	createNegativeTargetingClause *[]CreateNegativeTargetingClause
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of up to 100 negative targeting clauses for creation.
func (r ApiCreateNegativeTargetingClausesRequest) CreateNegativeTargetingClause(createNegativeTargetingClause []CreateNegativeTargetingClause) ApiCreateNegativeTargetingClausesRequest {
	r.createNegativeTargetingClause = &createNegativeTargetingClause
	return r
}

func (r ApiCreateNegativeTargetingClausesRequest) Execute() ([]TargetResponse, *http.Response, error) {
	return r.ApiService.CreateNegativeTargetingClausesExecute(r)
}

func (a *NegativeTargetingAPIService) CreateNegativeTargetingClauses(ctx context.Context) ApiCreateNegativeTargetingClausesRequest {
	return ApiCreateNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TargetResponse
func (a *NegativeTargetingAPIService) CreateNegativeTargetingClausesExecute(r ApiCreateNegativeTargetingClausesRequest) ([]TargetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TargetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingAPIService.CreateNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/negativeTargets"

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
	localVarPostBody = r.createNegativeTargetingClause
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

type ApiGetNegativeTargetsRequest struct {
	ctx                          context.Context
	ApiService                   *NegativeTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	negativeTargetId             int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetNegativeTargetsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetNegativeTargetsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetNegativeTargetsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetNegativeTargetsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetNegativeTargetsRequest) Execute() (*NegativeTargetingClause, *http.Response, error) {
	return r.ApiService.GetNegativeTargetsExecute(r)
}

func (a *NegativeTargetingAPIService) GetNegativeTargets(ctx context.Context, negativeTargetId int64) ApiGetNegativeTargetsRequest {
	return ApiGetNegativeTargetsRequest{
		ApiService:       a,
		ctx:              ctx,
		negativeTargetId: negativeTargetId,
	}
}

// Execute executes the request
//
//	@return NegativeTargetingClause
func (a *NegativeTargetingAPIService) GetNegativeTargetsExecute(r ApiGetNegativeTargetsRequest) (*NegativeTargetingClause, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *NegativeTargetingClause
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingAPIService.GetNegativeTargets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/negativeTargets/{negativeTargetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"negativeTargetId"+"}", url.PathEscape(parameterValueToString(r.negativeTargetId, "negativeTargetId")), -1)

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

type ApiGetNegativeTargetsExRequest struct {
	ctx                          context.Context
	ApiService                   *NegativeTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	negativeTargetId             int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetNegativeTargetsExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetNegativeTargetsExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetNegativeTargetsExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetNegativeTargetsExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetNegativeTargetsExRequest) Execute() (*NegativeTargetingClauseEx, *http.Response, error) {
	return r.ApiService.GetNegativeTargetsExExecute(r)
}

func (a *NegativeTargetingAPIService) GetNegativeTargetsEx(ctx context.Context, negativeTargetId int64) ApiGetNegativeTargetsExRequest {
	return ApiGetNegativeTargetsExRequest{
		ApiService:       a,
		ctx:              ctx,
		negativeTargetId: negativeTargetId,
	}
}

// Execute executes the request
//
//	@return NegativeTargetingClauseEx
func (a *NegativeTargetingAPIService) GetNegativeTargetsExExecute(r ApiGetNegativeTargetsExRequest) (*NegativeTargetingClauseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *NegativeTargetingClauseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingAPIService.GetNegativeTargetsEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/negativeTargets/extended/{negativeTargetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"negativeTargetId"+"}", url.PathEscape(parameterValueToString(r.negativeTargetId, "negativeTargetId")), -1)

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

type ApiListNegativeTargetingClausesRequest struct {
	ctx                          context.Context
	ApiService                   *NegativeTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	adGroupIdFilter              *string
	campaignIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. 0-indexed record offset for the result set. Defaults to 0.
func (r ApiListNegativeTargetingClausesRequest) StartIndex(startIndex int32) ApiListNegativeTargetingClausesRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Number of records to include in the paged response. Defaults to max page size.
func (r ApiListNegativeTargetingClausesRequest) Count(count int32) ApiListNegativeTargetingClausesRequest {
	r.count = &count
	return r
}

// Optional. Restricts results to those with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;, &#x60;paused&#x60;, or &#x60;archived&#x60;. Default behavior is to include enabled, paused, and archived.
func (r ApiListNegativeTargetingClausesRequest) StateFilter(stateFilter string) ApiListNegativeTargetingClausesRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional list of comma separated adGroupIds. Restricts results to negative targeting clauses with the specified &#x60;adGroupId&#x60;.
func (r ApiListNegativeTargetingClausesRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListNegativeTargetingClausesRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. Restricts results to targeting clauses within campaigns specified in comma-separated list.
func (r ApiListNegativeTargetingClausesRequest) CampaignIdFilter(campaignIdFilter string) ApiListNegativeTargetingClausesRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListNegativeTargetingClausesRequest) Execute() ([]NegativeTargetingClause, *http.Response, error) {
	return r.ApiService.ListNegativeTargetingClausesExecute(r)
}

func (a *NegativeTargetingAPIService) ListNegativeTargetingClauses(ctx context.Context) ApiListNegativeTargetingClausesRequest {
	return ApiListNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []NegativeTargetingClause
func (a *NegativeTargetingAPIService) ListNegativeTargetingClausesExecute(r ApiListNegativeTargetingClausesRequest) ([]NegativeTargetingClause, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []NegativeTargetingClause
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingAPIService.ListNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/negativeTargets"

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

type ApiListNegativeTargetingClausesExRequest struct {
	ctx                          context.Context
	ApiService                   *NegativeTargetingAPIService
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
func (r ApiListNegativeTargetingClausesExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListNegativeTargetingClausesExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListNegativeTargetingClausesExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListNegativeTargetingClausesExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. 0-indexed record offset for the result set. Defaults to 0.
func (r ApiListNegativeTargetingClausesExRequest) StartIndex(startIndex int32) ApiListNegativeTargetingClausesExRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Number of records to include in the paged response. Defaults to max page size.
func (r ApiListNegativeTargetingClausesExRequest) Count(count int32) ApiListNegativeTargetingClausesExRequest {
	r.count = &count
	return r
}

// Optional. Restricts results to keywords with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;, &#x60;paused&#x60;, or &#x60;archived&#x60;. Default behavior is to include &#x60;enabled&#x60;, &#x60;paused&#x60;, and &#x60;archived&#x60;.
func (r ApiListNegativeTargetingClausesExRequest) StateFilter(stateFilter string) ApiListNegativeTargetingClausesExRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. Restricts results to ads with the specified &#x60;tagetId&#x60; specified in comma-separated list
func (r ApiListNegativeTargetingClausesExRequest) TargetIdFilter(targetIdFilter string) ApiListNegativeTargetingClausesExRequest {
	r.targetIdFilter = &targetIdFilter
	return r
}

// Optional list of comma separated adGroupIds. Restricts results to negative targeting clauses with the specified &#x60;adGroupId&#x60;.
func (r ApiListNegativeTargetingClausesExRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListNegativeTargetingClausesExRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. Restricts results to ads within campaigns specified in the comma-separated list.
func (r ApiListNegativeTargetingClausesExRequest) CampaignIdFilter(campaignIdFilter string) ApiListNegativeTargetingClausesExRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListNegativeTargetingClausesExRequest) Execute() ([]NegativeTargetingClauseEx, *http.Response, error) {
	return r.ApiService.ListNegativeTargetingClausesExExecute(r)
}

func (a *NegativeTargetingAPIService) ListNegativeTargetingClausesEx(ctx context.Context) ApiListNegativeTargetingClausesExRequest {
	return ApiListNegativeTargetingClausesExRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []NegativeTargetingClauseEx
func (a *NegativeTargetingAPIService) ListNegativeTargetingClausesExExecute(r ApiListNegativeTargetingClausesExRequest) ([]NegativeTargetingClauseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []NegativeTargetingClauseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingAPIService.ListNegativeTargetingClausesEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/negativeTargets/extended"

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

type ApiUpdateNegativeTargetingClausesRequest struct {
	ctx                           context.Context
	ApiService                    *NegativeTargetingAPIService
	amazonAdvertisingAPIClientId  *string
	amazonAdvertisingAPIScope     *string
	updateNegativeTargetingClause *[]UpdateNegativeTargetingClause
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of up to 100 negative targeting clauses. Note that the only mutable field is &#x60;state&#x60;.
func (r ApiUpdateNegativeTargetingClausesRequest) UpdateNegativeTargetingClause(updateNegativeTargetingClause []UpdateNegativeTargetingClause) ApiUpdateNegativeTargetingClausesRequest {
	r.updateNegativeTargetingClause = &updateNegativeTargetingClause
	return r
}

func (r ApiUpdateNegativeTargetingClausesRequest) Execute() ([]TargetResponse, *http.Response, error) {
	return r.ApiService.UpdateNegativeTargetingClausesExecute(r)
}

func (a *NegativeTargetingAPIService) UpdateNegativeTargetingClauses(ctx context.Context) ApiUpdateNegativeTargetingClausesRequest {
	return ApiUpdateNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []TargetResponse
func (a *NegativeTargetingAPIService) UpdateNegativeTargetingClausesExecute(r ApiUpdateNegativeTargetingClausesRequest) ([]TargetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []TargetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingAPIService.UpdateNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/negativeTargets"

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
	localVarPostBody = r.updateNegativeTargetingClause
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
