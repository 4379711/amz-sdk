package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AdGroupsAPIService AdGroupsAPI service
type AdGroupsAPIService service

type ApiArchiveAdGroupRequest struct {
	ctx                          context.Context
	ApiService                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	adGroupId                    int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiArchiveAdGroupRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiArchiveAdGroupRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiArchiveAdGroupRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiArchiveAdGroupRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiArchiveAdGroupRequest) Execute() (*AdGroupResponse, *http.Response, error) {
	return r.ApiService.ArchiveAdGroupExecute(r)
}

func (a *AdGroupsAPIService) ArchiveAdGroup(ctx context.Context, adGroupId int64) ApiArchiveAdGroupRequest {
	return ApiArchiveAdGroupRequest{
		ApiService: a,
		ctx:        ctx,
		adGroupId:  adGroupId,
	}
}

// Execute executes the request
//
//	@return AdGroupResponse
func (a *AdGroupsAPIService) ArchiveAdGroupExecute(r ApiArchiveAdGroupRequest) (*AdGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.ArchiveAdGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups/{adGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adGroupId"+"}", url.PathEscape(parameterValueToString(r.adGroupId, "adGroupId")), -1)

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

type ApiCreateAdGroupsRequest struct {
	ctx                          context.Context
	ApiService                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createAdGroup                *[]CreateAdGroup
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of AdGroup objects. For each object, specify required fields and their values. Required fields are &#x60;campaignId&#x60;, &#x60;name&#x60;, &#x60;state&#x60;, and &#x60;defaultBid&#x60;. Maximum length of the array is 100 objects. Note - when using landingPageType of OFF_AMAZON_LINK or STORES within productAds, only 1 adGroup is supported.
func (r ApiCreateAdGroupsRequest) CreateAdGroup(createAdGroup []CreateAdGroup) ApiCreateAdGroupsRequest {
	r.createAdGroup = &createAdGroup
	return r
}

func (r ApiCreateAdGroupsRequest) Execute() ([]AdGroupResponse, *http.Response, error) {
	return r.ApiService.CreateAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) CreateAdGroups(ctx context.Context) ApiCreateAdGroupsRequest {
	return ApiCreateAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AdGroupResponse
func (a *AdGroupsAPIService) CreateAdGroupsExecute(r ApiCreateAdGroupsRequest) ([]AdGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AdGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.CreateAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups"

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
	localVarPostBody = r.createAdGroup
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

type ApiGetAdGroupRequest struct {
	ctx                          context.Context
	ApiService                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	adGroupId                    int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetAdGroupRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetAdGroupRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetAdGroupRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetAdGroupRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetAdGroupRequest) Execute() (*AdGroup, *http.Response, error) {
	return r.ApiService.GetAdGroupExecute(r)
}

func (a *AdGroupsAPIService) GetAdGroup(ctx context.Context, adGroupId int64) ApiGetAdGroupRequest {
	return ApiGetAdGroupRequest{
		ApiService: a,
		ctx:        ctx,
		adGroupId:  adGroupId,
	}
}

// Execute executes the request
//
//	@return AdGroup
func (a *AdGroupsAPIService) GetAdGroupExecute(r ApiGetAdGroupRequest) (*AdGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.GetAdGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups/{adGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adGroupId"+"}", url.PathEscape(parameterValueToString(r.adGroupId, "adGroupId")), -1)

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

type ApiGetAdGroupResponseExRequest struct {
	ctx                          context.Context
	ApiService                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	adGroupId                    int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetAdGroupResponseExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetAdGroupResponseExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetAdGroupResponseExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetAdGroupResponseExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetAdGroupResponseExRequest) Execute() (*AdGroupResponseEx, *http.Response, error) {
	return r.ApiService.GetAdGroupResponseExExecute(r)
}

func (a *AdGroupsAPIService) GetAdGroupResponseEx(ctx context.Context, adGroupId int64) ApiGetAdGroupResponseExRequest {
	return ApiGetAdGroupResponseExRequest{
		ApiService: a,
		ctx:        ctx,
		adGroupId:  adGroupId,
	}
}

// Execute executes the request
//
//	@return AdGroupResponseEx
func (a *AdGroupsAPIService) GetAdGroupResponseExExecute(r ApiGetAdGroupResponseExRequest) (*AdGroupResponseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdGroupResponseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.GetAdGroupResponseEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups/extended/{adGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adGroupId"+"}", url.PathEscape(parameterValueToString(r.adGroupId, "adGroupId")), -1)

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

type ApiListAdGroupsRequest struct {
	ctx                          context.Context
	ApiService                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	campaignIdFilter             *string
	adGroupIdFilter              *string
	name                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of campaigns. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListAdGroupsRequest) StartIndex(startIndex int32) ApiListAdGroupsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of AdGroup objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten ad groups set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten ad groups, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListAdGroupsRequest) Count(count int32) ApiListAdGroupsRequest {
	r.count = &count
	return r
}

// Optional. The returned array is filtered to include only ad groups with state set to one of the values in the specified comma-delimited list.
func (r ApiListAdGroupsRequest) StateFilter(stateFilter string) ApiListAdGroupsRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. The returned array is filtered to include only ad groups associated with the campaign identifiers in the specified comma-delimited list.
func (r ApiListAdGroupsRequest) CampaignIdFilter(campaignIdFilter string) ApiListAdGroupsRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

// Optional. The returned array is filtered to include only ad groups with an identifier specified in the comma-delimited list.
func (r ApiListAdGroupsRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListAdGroupsRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. The returned array includes only ad groups with the specified name.
func (r ApiListAdGroupsRequest) Name(name string) ApiListAdGroupsRequest {
	r.name = &name
	return r
}

func (r ApiListAdGroupsRequest) Execute() ([]AdGroup, *http.Response, error) {
	return r.ApiService.ListAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) ListAdGroups(ctx context.Context) ApiListAdGroupsRequest {
	return ApiListAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AdGroup
func (a *AdGroupsAPIService) ListAdGroupsExecute(r ApiListAdGroupsRequest) ([]AdGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AdGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.ListAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups"

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
	if r.campaignIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaignIdFilter", r.campaignIdFilter, "form", "")
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
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

type ApiListAdGroupsExRequest struct {
	ctx                          context.Context
	ApiService                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	campaignIdFilter             *string
	adGroupIdFilter              *string
	name                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListAdGroupsExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListAdGroupsExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListAdGroupsExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListAdGroupsExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of ad groups. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListAdGroupsExRequest) StartIndex(startIndex int32) ApiListAdGroupsExRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of Campaign objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten campaigns set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListAdGroupsExRequest) Count(count int32) ApiListAdGroupsExRequest {
	r.count = &count
	return r
}

// Optional. The returned array is filtered to include only campaigns with state set to one of the values in the comma-delimited list.
func (r ApiListAdGroupsExRequest) StateFilter(stateFilter string) ApiListAdGroupsExRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. The returned array is filtered to include only ad groups associated with the campaign identifiers in the comma-delimited list.
func (r ApiListAdGroupsExRequest) CampaignIdFilter(campaignIdFilter string) ApiListAdGroupsExRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

// Optional. The returned array is filtered to include only ad groups with an identifier specified in the comma-delimited list.
func (r ApiListAdGroupsExRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListAdGroupsExRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. The returned array includes only ad groups with the specified name.
func (r ApiListAdGroupsExRequest) Name(name string) ApiListAdGroupsExRequest {
	r.name = &name
	return r
}

func (r ApiListAdGroupsExRequest) Execute() ([]AdGroupResponseEx, *http.Response, error) {
	return r.ApiService.ListAdGroupsExExecute(r)
}

func (a *AdGroupsAPIService) ListAdGroupsEx(ctx context.Context) ApiListAdGroupsExRequest {
	return ApiListAdGroupsExRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AdGroupResponseEx
func (a *AdGroupsAPIService) ListAdGroupsExExecute(r ApiListAdGroupsExRequest) ([]AdGroupResponseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AdGroupResponseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.ListAdGroupsEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups/extended"

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
	if r.campaignIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaignIdFilter", r.campaignIdFilter, "form", "")
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
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

type ApiUpdateAdGroupsRequest struct {
	ctx                          context.Context
	ApiService                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateAdGroup                *[]UpdateAdGroup
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of AdGroup objects. For each object, specify an ad group identifier and mutable fields with their updated values. The mutable fields are &#39;name&#39;, &#39;defaultBid&#39;, and &#39;state&#39;. Maximum length of the array is 100 objects.
func (r ApiUpdateAdGroupsRequest) UpdateAdGroup(updateAdGroup []UpdateAdGroup) ApiUpdateAdGroupsRequest {
	r.updateAdGroup = &updateAdGroup
	return r
}

func (r ApiUpdateAdGroupsRequest) Execute() ([]AdGroupResponse, *http.Response, error) {
	return r.ApiService.UpdateAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) UpdateAdGroups(ctx context.Context) ApiUpdateAdGroupsRequest {
	return ApiUpdateAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AdGroupResponse
func (a *AdGroupsAPIService) UpdateAdGroupsExecute(r ApiUpdateAdGroupsRequest) ([]AdGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AdGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.UpdateAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups"

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
	localVarPostBody = r.updateAdGroup
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
