package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CampaignsAPIService CampaignsAPI service
type CampaignsAPIService service

type ApiArchiveCampaignRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignId                   int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiArchiveCampaignRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiArchiveCampaignRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiArchiveCampaignRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiArchiveCampaignRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiArchiveCampaignRequest) Execute() (*CampaignResponse, *http.Response, error) {
	return r.ApiService.ArchiveCampaignExecute(r)
}

func (a *CampaignsAPIService) ArchiveCampaign(ctx context.Context, campaignId int64) ApiArchiveCampaignRequest {
	return ApiArchiveCampaignRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return CampaignResponse
func (a *CampaignsAPIService) ArchiveCampaignExecute(r ApiArchiveCampaignRequest) (*CampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ArchiveCampaign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns/{campaignId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

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

type ApiCreateCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createCampaign               *[]CreateCampaign
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of Campaign objects. For each object, specify required fields and their values. Required fields are &#x60;name&#x60;, &#x60;tactic&#x60;, &#x60;state&#x60;, and &#x60;startDate&#x60;. Maximum length of the array is 100 objects. If you don&#39;t specify a &#x60;budget&#x60;, it will be set as the [default budget for your region](https://advertising.amazon.com/API/docs/en-us/concepts/limits#default-budgets). Campaign names must be unique across SD, SB, and SP.   If you are using Optimization rules, the following campaign budget must be at least:   - 5x the value of any COST_PER_ORDER threshold.   - 10x the value of any COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS threshold.   - 20x the value of any COST_PER_CLICK threshold.
func (r ApiCreateCampaignsRequest) CreateCampaign(createCampaign []CreateCampaign) ApiCreateCampaignsRequest {
	r.createCampaign = &createCampaign
	return r
}

func (r ApiCreateCampaignsRequest) Execute() ([]CampaignResponse, *http.Response, error) {
	return r.ApiService.CreateCampaignsExecute(r)
}

func (a *CampaignsAPIService) CreateCampaigns(ctx context.Context) ApiCreateCampaignsRequest {
	return ApiCreateCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CampaignResponse
func (a *CampaignsAPIService) CreateCampaignsExecute(r ApiCreateCampaignsRequest) ([]CampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.CreateCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns"

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
	localVarPostBody = r.createCampaign
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

type ApiGetCampaignRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignId                   int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetCampaignRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetCampaignRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetCampaignRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetCampaignRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetCampaignRequest) Execute() (*Campaign, *http.Response, error) {
	return r.ApiService.GetCampaignExecute(r)
}

func (a *CampaignsAPIService) GetCampaign(ctx context.Context, campaignId int64) ApiGetCampaignRequest {
	return ApiGetCampaignRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return Campaign
func (a *CampaignsAPIService) GetCampaignExecute(r ApiGetCampaignRequest) (*Campaign, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Campaign
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.GetCampaign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns/{campaignId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

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

type ApiGetCampaignResponseExRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignId                   int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetCampaignResponseExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetCampaignResponseExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetCampaignResponseExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetCampaignResponseExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetCampaignResponseExRequest) Execute() (*CampaignResponseEx, *http.Response, error) {
	return r.ApiService.GetCampaignResponseExExecute(r)
}

func (a *CampaignsAPIService) GetCampaignResponseEx(ctx context.Context, campaignId int64) ApiGetCampaignResponseExRequest {
	return ApiGetCampaignResponseExRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return CampaignResponseEx
func (a *CampaignsAPIService) GetCampaignResponseExExecute(r ApiGetCampaignResponseExRequest) (*CampaignResponseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CampaignResponseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.GetCampaignResponseEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns/extended/{campaignId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

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

type ApiListCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	name                         *string
	campaignIdFilter             *string
	portfolioIdFilter            *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of campaigns. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListCampaignsRequest) StartIndex(startIndex int32) ApiListCampaignsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of Campaign objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten campaigns set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListCampaignsRequest) Count(count int32) ApiListCampaignsRequest {
	r.count = &count
	return r
}

// Optional. The returned array is filtered to include only campaigns with state set to one of the values in the specified comma-delimited list.
func (r ApiListCampaignsRequest) StateFilter(stateFilter string) ApiListCampaignsRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. The returned array includes only campaign with the specified name using an exact string match.
func (r ApiListCampaignsRequest) Name(name string) ApiListCampaignsRequest {
	r.name = &name
	return r
}

// Optional. The returned array includes only campaigns with identifiers matching those specified in the comma-delimited string.
func (r ApiListCampaignsRequest) CampaignIdFilter(campaignIdFilter string) ApiListCampaignsRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

// Optional. The returned array includes only campaigns associated with Portfolio identifiers matching those specified in the comma-delimited string.
func (r ApiListCampaignsRequest) PortfolioIdFilter(portfolioIdFilter string) ApiListCampaignsRequest {
	r.portfolioIdFilter = &portfolioIdFilter
	return r
}

func (r ApiListCampaignsRequest) Execute() ([]Campaign, *http.Response, error) {
	return r.ApiService.ListCampaignsExecute(r)
}

func (a *CampaignsAPIService) ListCampaigns(ctx context.Context) ApiListCampaignsRequest {
	return ApiListCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Campaign
func (a *CampaignsAPIService) ListCampaignsExecute(r ApiListCampaignsRequest) ([]Campaign, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Campaign
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.campaignIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaignIdFilter", r.campaignIdFilter, "form", "")
	}
	if r.portfolioIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "portfolioIdFilter", r.portfolioIdFilter, "form", "")
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

type ApiListCampaignsExRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	name                         *string
	campaignIdFilter             *string
	portfolioIdFilter            *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListCampaignsExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListCampaignsExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListCampaignsExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListCampaignsExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of campaigns. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListCampaignsExRequest) StartIndex(startIndex int32) ApiListCampaignsExRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of Campaign objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten campaigns set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListCampaignsExRequest) Count(count int32) ApiListCampaignsExRequest {
	r.count = &count
	return r
}

// Optional. The returned array is filtered to include only campaigns with state set to one of the values in the specified comma-delimited list.
func (r ApiListCampaignsExRequest) StateFilter(stateFilter string) ApiListCampaignsExRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. The returned array includes only campaign with the specified name using an exact string match.
func (r ApiListCampaignsExRequest) Name(name string) ApiListCampaignsExRequest {
	r.name = &name
	return r
}

// Optional. The returned array includes only campaigns with identifiers matching those specified in the comma-delimited string.
func (r ApiListCampaignsExRequest) CampaignIdFilter(campaignIdFilter string) ApiListCampaignsExRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

// Optional. The returned array includes only campaigns associated with Portfolio identifiers matching those specified in the comma-delimited string.
func (r ApiListCampaignsExRequest) PortfolioIdFilter(portfolioIdFilter string) ApiListCampaignsExRequest {
	r.portfolioIdFilter = &portfolioIdFilter
	return r
}

func (r ApiListCampaignsExRequest) Execute() ([]CampaignResponseEx, *http.Response, error) {
	return r.ApiService.ListCampaignsExExecute(r)
}

func (a *CampaignsAPIService) ListCampaignsEx(ctx context.Context) ApiListCampaignsExRequest {
	return ApiListCampaignsExRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CampaignResponseEx
func (a *CampaignsAPIService) ListCampaignsExExecute(r ApiListCampaignsExRequest) ([]CampaignResponseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CampaignResponseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListCampaignsEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns/extended"

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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.campaignIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaignIdFilter", r.campaignIdFilter, "form", "")
	}
	if r.portfolioIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "portfolioIdFilter", r.portfolioIdFilter, "form", "")
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

type ApiUpdateCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateCampaign               *[]UpdateCampaign
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of Campaign objects. For each object, specify a campaign identifier and mutable fields with their updated values. The mutable fields are &#x60;name&#x60;, &#x60;state&#x60;, &#x60;budget&#x60;, &#x60;startDate&#x60;, and &#x60;endDate&#x60;. Maximum length of the array is 100 objects.
func (r ApiUpdateCampaignsRequest) UpdateCampaign(updateCampaign []UpdateCampaign) ApiUpdateCampaignsRequest {
	r.updateCampaign = &updateCampaign
	return r
}

func (r ApiUpdateCampaignsRequest) Execute() ([]CampaignResponse, *http.Response, error) {
	return r.ApiService.UpdateCampaignsExecute(r)
}

func (a *CampaignsAPIService) UpdateCampaigns(ctx context.Context) ApiUpdateCampaignsRequest {
	return ApiUpdateCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CampaignResponse
func (a *CampaignsAPIService) UpdateCampaignsExecute(r ApiUpdateCampaignsRequest) ([]CampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.UpdateCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns"

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
	localVarPostBody = r.updateCampaign
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
