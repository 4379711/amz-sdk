package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// BudgetRulesAPIService BudgetRulesAPI service
type BudgetRulesAPIService service

type ApiCreateAssociatedBudgetRulesForSBCampaignsRequest struct {
	ctx                                context.Context
	ApiService                         *BudgetRulesAPIService
	amazonAdvertisingAPIClientId       *string
	amazonAdvertisingAPIScope          *string
	campaignId                         float32
	createAssociatedBudgetRulesRequest *CreateAssociatedBudgetRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateAssociatedBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateAssociatedBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateAssociatedBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateAssociatedBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateAssociatedBudgetRulesForSBCampaignsRequest) CreateAssociatedBudgetRulesRequest(createAssociatedBudgetRulesRequest CreateAssociatedBudgetRulesRequest) ApiCreateAssociatedBudgetRulesForSBCampaignsRequest {
	r.createAssociatedBudgetRulesRequest = &createAssociatedBudgetRulesRequest
	return r
}

func (r ApiCreateAssociatedBudgetRulesForSBCampaignsRequest) Execute() (*CreateAssociatedBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.CreateAssociatedBudgetRulesForSBCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) CreateAssociatedBudgetRulesForSBCampaigns(ctx context.Context, campaignId float32) ApiCreateAssociatedBudgetRulesForSBCampaignsRequest {
	return ApiCreateAssociatedBudgetRulesForSBCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return CreateAssociatedBudgetRulesResponse
func (a *BudgetRulesAPIService) CreateAssociatedBudgetRulesForSBCampaignsExecute(r ApiCreateAssociatedBudgetRulesForSBCampaignsRequest) (*CreateAssociatedBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateAssociatedBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.CreateAssociatedBudgetRulesForSBCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/campaigns/{campaignId}/budgetRules"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createAssociatedBudgetRulesRequest == nil {
		return localVarReturnValue, nil, reportError("createAssociatedBudgetRulesRequest is required and must be specified")
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
	localVarPostBody = r.createAssociatedBudgetRulesRequest
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

type ApiCreateBudgetRulesForSBCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createSBBudgetRulesRequest   *CreateSBBudgetRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateBudgetRulesForSBCampaignsRequest) CreateSBBudgetRulesRequest(createSBBudgetRulesRequest CreateSBBudgetRulesRequest) ApiCreateBudgetRulesForSBCampaignsRequest {
	r.createSBBudgetRulesRequest = &createSBBudgetRulesRequest
	return r
}

func (r ApiCreateBudgetRulesForSBCampaignsRequest) Execute() (*CreateBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.CreateBudgetRulesForSBCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) CreateBudgetRulesForSBCampaigns(ctx context.Context) ApiCreateBudgetRulesForSBCampaignsRequest {
	return ApiCreateBudgetRulesForSBCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateBudgetRulesResponse
func (a *BudgetRulesAPIService) CreateBudgetRulesForSBCampaignsExecute(r ApiCreateBudgetRulesForSBCampaignsRequest) (*CreateBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.CreateBudgetRulesForSBCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/budgetRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSBBudgetRulesRequest == nil {
		return localVarReturnValue, nil, reportError("createSBBudgetRulesRequest is required and must be specified")
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
	localVarPostBody = r.createSBBudgetRulesRequest
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

type ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignId                   float32
	budgetRuleId                 string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DisassociateAssociatedBudgetRuleForSBCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) DisassociateAssociatedBudgetRuleForSBCampaigns(ctx context.Context, campaignId float32, budgetRuleId string) ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest {
	return ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest{
		ApiService:   a,
		ctx:          ctx,
		campaignId:   campaignId,
		budgetRuleId: budgetRuleId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *BudgetRulesAPIService) DisassociateAssociatedBudgetRuleForSBCampaignsExecute(r ApiDisassociateAssociatedBudgetRuleForSBCampaignsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.DisassociateAssociatedBudgetRuleForSBCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/campaigns/{campaignId}/budgetRules/{budgetRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"budgetRuleId"+"}", url.PathEscape(parameterValueToString(r.budgetRuleId, "budgetRuleId")), -1)

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

type ApiGetBudgetRuleByRuleIdForSBCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	budgetRuleId                 string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetBudgetRuleByRuleIdForSBCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetBudgetRuleByRuleIdForSBCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetBudgetRuleByRuleIdForSBCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetBudgetRuleByRuleIdForSBCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetBudgetRuleByRuleIdForSBCampaignsRequest) Execute() (*GetSBBudgetRuleResponse, *http.Response, error) {
	return r.ApiService.GetBudgetRuleByRuleIdForSBCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) GetBudgetRuleByRuleIdForSBCampaigns(ctx context.Context, budgetRuleId string) ApiGetBudgetRuleByRuleIdForSBCampaignsRequest {
	return ApiGetBudgetRuleByRuleIdForSBCampaignsRequest{
		ApiService:   a,
		ctx:          ctx,
		budgetRuleId: budgetRuleId,
	}
}

// Execute executes the request
//
//	@return GetSBBudgetRuleResponse
func (a *BudgetRulesAPIService) GetBudgetRuleByRuleIdForSBCampaignsExecute(r ApiGetBudgetRuleByRuleIdForSBCampaignsRequest) (*GetSBBudgetRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSBBudgetRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetBudgetRuleByRuleIdForSBCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/budgetRules/{budgetRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"budgetRuleId"+"}", url.PathEscape(parameterValueToString(r.budgetRuleId, "budgetRuleId")), -1)

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

type ApiGetCampaignsAssociatedWithSBBudgetRuleRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	budgetRuleId                 string
	pageSize                     *float32
	nextToken                    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCampaignsAssociatedWithSBBudgetRuleRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetCampaignsAssociatedWithSBBudgetRuleRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCampaignsAssociatedWithSBBudgetRuleRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetCampaignsAssociatedWithSBBudgetRuleRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30.
func (r ApiGetCampaignsAssociatedWithSBBudgetRuleRequest) PageSize(pageSize float32) ApiGetCampaignsAssociatedWithSBBudgetRuleRequest {
	r.pageSize = &pageSize
	return r
}

// To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results.
func (r ApiGetCampaignsAssociatedWithSBBudgetRuleRequest) NextToken(nextToken string) ApiGetCampaignsAssociatedWithSBBudgetRuleRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetCampaignsAssociatedWithSBBudgetRuleRequest) Execute() (*SBGetAssociatedCampaignsResponse, *http.Response, error) {
	return r.ApiService.GetCampaignsAssociatedWithSBBudgetRuleExecute(r)
}

func (a *BudgetRulesAPIService) GetCampaignsAssociatedWithSBBudgetRule(ctx context.Context, budgetRuleId string) ApiGetCampaignsAssociatedWithSBBudgetRuleRequest {
	return ApiGetCampaignsAssociatedWithSBBudgetRuleRequest{
		ApiService:   a,
		ctx:          ctx,
		budgetRuleId: budgetRuleId,
	}
}

// Execute executes the request
//
//	@return SBGetAssociatedCampaignsResponse
func (a *BudgetRulesAPIService) GetCampaignsAssociatedWithSBBudgetRuleExecute(r ApiGetCampaignsAssociatedWithSBBudgetRuleRequest) (*SBGetAssociatedCampaignsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBGetAssociatedCampaignsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetCampaignsAssociatedWithSBBudgetRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/budgetRules/{budgetRuleId}/campaigns"
	localVarPath = strings.Replace(localVarPath, "{"+"budgetRuleId"+"}", url.PathEscape(parameterValueToString(r.budgetRuleId, "budgetRuleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize == nil {
		return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
	}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
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

type ApiGetSBBudgetRulesForAdvertiserRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	pageSize                     *float32
	nextToken                    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetSBBudgetRulesForAdvertiserRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetSBBudgetRulesForAdvertiserRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetSBBudgetRulesForAdvertiserRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetSBBudgetRulesForAdvertiserRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30.
func (r ApiGetSBBudgetRulesForAdvertiserRequest) PageSize(pageSize float32) ApiGetSBBudgetRulesForAdvertiserRequest {
	r.pageSize = &pageSize
	return r
}

// To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results.
func (r ApiGetSBBudgetRulesForAdvertiserRequest) NextToken(nextToken string) ApiGetSBBudgetRulesForAdvertiserRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetSBBudgetRulesForAdvertiserRequest) Execute() (*GetSBBudgetRulesForAdvertiserResponse, *http.Response, error) {
	return r.ApiService.GetSBBudgetRulesForAdvertiserExecute(r)
}

func (a *BudgetRulesAPIService) GetSBBudgetRulesForAdvertiser(ctx context.Context) ApiGetSBBudgetRulesForAdvertiserRequest {
	return ApiGetSBBudgetRulesForAdvertiserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSBBudgetRulesForAdvertiserResponse
func (a *BudgetRulesAPIService) GetSBBudgetRulesForAdvertiserExecute(r ApiGetSBBudgetRulesForAdvertiserRequest) (*GetSBBudgetRulesForAdvertiserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSBBudgetRulesForAdvertiserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetSBBudgetRulesForAdvertiser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/budgetRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize == nil {
		return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
	}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
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

type ApiListAssociatedBudgetRulesForSBCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignId                   float32
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiListAssociatedBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListAssociatedBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiListAssociatedBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListAssociatedBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListAssociatedBudgetRulesForSBCampaignsRequest) Execute() (*SBListAssociatedBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.ListAssociatedBudgetRulesForSBCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) ListAssociatedBudgetRulesForSBCampaigns(ctx context.Context, campaignId float32) ApiListAssociatedBudgetRulesForSBCampaignsRequest {
	return ApiListAssociatedBudgetRulesForSBCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return SBListAssociatedBudgetRulesResponse
func (a *BudgetRulesAPIService) ListAssociatedBudgetRulesForSBCampaignsExecute(r ApiListAssociatedBudgetRulesForSBCampaignsRequest) (*SBListAssociatedBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBListAssociatedBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.ListAssociatedBudgetRulesForSBCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/campaigns/{campaignId}/budgetRules"
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

type ApiUpdateBudgetRulesForSBCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateSBBudgetRulesRequest   *UpdateSBBudgetRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateBudgetRulesForSBCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateBudgetRulesForSBCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateBudgetRulesForSBCampaignsRequest) UpdateSBBudgetRulesRequest(updateSBBudgetRulesRequest UpdateSBBudgetRulesRequest) ApiUpdateBudgetRulesForSBCampaignsRequest {
	r.updateSBBudgetRulesRequest = &updateSBBudgetRulesRequest
	return r
}

func (r ApiUpdateBudgetRulesForSBCampaignsRequest) Execute() (*UpdateBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.UpdateBudgetRulesForSBCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) UpdateBudgetRulesForSBCampaigns(ctx context.Context) ApiUpdateBudgetRulesForSBCampaignsRequest {
	return ApiUpdateBudgetRulesForSBCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateBudgetRulesResponse
func (a *BudgetRulesAPIService) UpdateBudgetRulesForSBCampaignsExecute(r ApiUpdateBudgetRulesForSBCampaignsRequest) (*UpdateBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.UpdateBudgetRulesForSBCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/budgetRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSBBudgetRulesRequest == nil {
		return localVarReturnValue, nil, reportError("updateSBBudgetRulesRequest is required and must be specified")
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
	localVarPostBody = r.updateSBBudgetRulesRequest
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
