package sp_v3

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

type ApiCreateAssociatedBudgetRulesForSPCampaignsRequest struct {
	ctx                                context.Context
	ApiService                         *BudgetRulesAPIService
	amazonAdvertisingAPIClientId       *string
	amazonAdvertisingAPIScope          *string
	campaignId                         float32
	createAssociatedBudgetRulesRequest *CreateAssociatedBudgetRulesRequest
}

// The identifier of a client associated with a Login with Amazon account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateAssociatedBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateAssociatedBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateAssociatedBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateAssociatedBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateAssociatedBudgetRulesForSPCampaignsRequest) CreateAssociatedBudgetRulesRequest(createAssociatedBudgetRulesRequest CreateAssociatedBudgetRulesRequest) ApiCreateAssociatedBudgetRulesForSPCampaignsRequest {
	r.createAssociatedBudgetRulesRequest = &createAssociatedBudgetRulesRequest
	return r
}

func (r ApiCreateAssociatedBudgetRulesForSPCampaignsRequest) Execute() (*CreateAssociatedBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.CreateAssociatedBudgetRulesForSPCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) CreateAssociatedBudgetRulesForSPCampaigns(ctx context.Context, campaignId float32) ApiCreateAssociatedBudgetRulesForSPCampaignsRequest {
	return ApiCreateAssociatedBudgetRulesForSPCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return CreateAssociatedBudgetRulesResponse
func (a *BudgetRulesAPIService) CreateAssociatedBudgetRulesForSPCampaignsExecute(r ApiCreateAssociatedBudgetRulesForSPCampaignsRequest) (*CreateAssociatedBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateAssociatedBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.CreateAssociatedBudgetRulesForSPCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/{campaignId}/budgetRules"
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

type ApiCreateBudgetRulesForSPCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createSPBudgetRulesRequest   *CreateSPBudgetRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateBudgetRulesForSPCampaignsRequest) CreateSPBudgetRulesRequest(createSPBudgetRulesRequest CreateSPBudgetRulesRequest) ApiCreateBudgetRulesForSPCampaignsRequest {
	r.createSPBudgetRulesRequest = &createSPBudgetRulesRequest
	return r
}

func (r ApiCreateBudgetRulesForSPCampaignsRequest) Execute() (*CreateBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.CreateBudgetRulesForSPCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) CreateBudgetRulesForSPCampaigns(ctx context.Context) ApiCreateBudgetRulesForSPCampaignsRequest {
	return ApiCreateBudgetRulesForSPCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateBudgetRulesResponse
func (a *BudgetRulesAPIService) CreateBudgetRulesForSPCampaignsExecute(r ApiCreateBudgetRulesForSPCampaignsRequest) (*CreateBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.CreateBudgetRulesForSPCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/budgetRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSPBudgetRulesRequest == nil {
		return localVarReturnValue, nil, reportError("createSPBudgetRulesRequest is required and must be specified")
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
	localVarPostBody = r.createSPBudgetRulesRequest
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

type ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignId                   float32
	budgetRuleId                 string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DisassociateAssociatedBudgetRuleForSPCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) DisassociateAssociatedBudgetRuleForSPCampaigns(ctx context.Context, campaignId float32, budgetRuleId string) ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest {
	return ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest{
		ApiService:   a,
		ctx:          ctx,
		campaignId:   campaignId,
		budgetRuleId: budgetRuleId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *BudgetRulesAPIService) DisassociateAssociatedBudgetRuleForSPCampaignsExecute(r ApiDisassociateAssociatedBudgetRuleForSPCampaignsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.DisassociateAssociatedBudgetRuleForSPCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/{campaignId}/budgetRules/{budgetRuleId}"
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

type ApiGetBudgetRuleByRuleIdForSPCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	budgetRuleId                 string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetBudgetRuleByRuleIdForSPCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetBudgetRuleByRuleIdForSPCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetBudgetRuleByRuleIdForSPCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetBudgetRuleByRuleIdForSPCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetBudgetRuleByRuleIdForSPCampaignsRequest) Execute() (*GetSPBudgetRuleResponse, *http.Response, error) {
	return r.ApiService.GetBudgetRuleByRuleIdForSPCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) GetBudgetRuleByRuleIdForSPCampaigns(ctx context.Context, budgetRuleId string) ApiGetBudgetRuleByRuleIdForSPCampaignsRequest {
	return ApiGetBudgetRuleByRuleIdForSPCampaignsRequest{
		ApiService:   a,
		ctx:          ctx,
		budgetRuleId: budgetRuleId,
	}
}

// Execute executes the request
//
//	@return GetSPBudgetRuleResponse
func (a *BudgetRulesAPIService) GetBudgetRuleByRuleIdForSPCampaignsExecute(r ApiGetBudgetRuleByRuleIdForSPCampaignsRequest) (*GetSPBudgetRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSPBudgetRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetBudgetRuleByRuleIdForSPCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/budgetRules/{budgetRuleId}"
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

type ApiGetCampaignsAssociatedWithSPBudgetRuleRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	budgetRuleId                 string
	pageSize                     *float32
	nextToken                    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCampaignsAssociatedWithSPBudgetRuleRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetCampaignsAssociatedWithSPBudgetRuleRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCampaignsAssociatedWithSPBudgetRuleRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetCampaignsAssociatedWithSPBudgetRuleRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30.
func (r ApiGetCampaignsAssociatedWithSPBudgetRuleRequest) PageSize(pageSize float32) ApiGetCampaignsAssociatedWithSPBudgetRuleRequest {
	r.pageSize = &pageSize
	return r
}

// To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results.
func (r ApiGetCampaignsAssociatedWithSPBudgetRuleRequest) NextToken(nextToken string) ApiGetCampaignsAssociatedWithSPBudgetRuleRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetCampaignsAssociatedWithSPBudgetRuleRequest) Execute() (*SPGetAssociatedCampaignsResponse, *http.Response, error) {
	return r.ApiService.GetCampaignsAssociatedWithSPBudgetRuleExecute(r)
}

func (a *BudgetRulesAPIService) GetCampaignsAssociatedWithSPBudgetRule(ctx context.Context, budgetRuleId string) ApiGetCampaignsAssociatedWithSPBudgetRuleRequest {
	return ApiGetCampaignsAssociatedWithSPBudgetRuleRequest{
		ApiService:   a,
		ctx:          ctx,
		budgetRuleId: budgetRuleId,
	}
}

// Execute executes the request
//
//	@return SPGetAssociatedCampaignsResponse
func (a *BudgetRulesAPIService) GetCampaignsAssociatedWithSPBudgetRuleExecute(r ApiGetCampaignsAssociatedWithSPBudgetRuleRequest) (*SPGetAssociatedCampaignsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SPGetAssociatedCampaignsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetCampaignsAssociatedWithSPBudgetRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/budgetRules/{budgetRuleId}/campaigns"
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

type ApiGetSPBudgetRulesForAdvertiserRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	pageSize                     *float32
	nextToken                    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetSPBudgetRulesForAdvertiserRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetSPBudgetRulesForAdvertiserRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetSPBudgetRulesForAdvertiserRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetSPBudgetRulesForAdvertiserRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30.
func (r ApiGetSPBudgetRulesForAdvertiserRequest) PageSize(pageSize float32) ApiGetSPBudgetRulesForAdvertiserRequest {
	r.pageSize = &pageSize
	return r
}

// To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results.
func (r ApiGetSPBudgetRulesForAdvertiserRequest) NextToken(nextToken string) ApiGetSPBudgetRulesForAdvertiserRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetSPBudgetRulesForAdvertiserRequest) Execute() (*GetSPBudgetRulesForAdvertiserResponse, *http.Response, error) {
	return r.ApiService.GetSPBudgetRulesForAdvertiserExecute(r)
}

func (a *BudgetRulesAPIService) GetSPBudgetRulesForAdvertiser(ctx context.Context) ApiGetSPBudgetRulesForAdvertiserRequest {
	return ApiGetSPBudgetRulesForAdvertiserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSPBudgetRulesForAdvertiserResponse
func (a *BudgetRulesAPIService) GetSPBudgetRulesForAdvertiserExecute(r ApiGetSPBudgetRulesForAdvertiserRequest) (*GetSPBudgetRulesForAdvertiserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSPBudgetRulesForAdvertiserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetSPBudgetRulesForAdvertiser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/budgetRules"

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

type ApiListAssociatedBudgetRulesForSPCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignId                   float32
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiListAssociatedBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListAssociatedBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiListAssociatedBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListAssociatedBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListAssociatedBudgetRulesForSPCampaignsRequest) Execute() (*SPListAssociatedBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.ListAssociatedBudgetRulesForSPCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) ListAssociatedBudgetRulesForSPCampaigns(ctx context.Context, campaignId float32) ApiListAssociatedBudgetRulesForSPCampaignsRequest {
	return ApiListAssociatedBudgetRulesForSPCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return SPListAssociatedBudgetRulesResponse
func (a *BudgetRulesAPIService) ListAssociatedBudgetRulesForSPCampaignsExecute(r ApiListAssociatedBudgetRulesForSPCampaignsRequest) (*SPListAssociatedBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SPListAssociatedBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.ListAssociatedBudgetRulesForSPCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/{campaignId}/budgetRules"
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

type ApiUpdateBudgetRulesForSPCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateSPBudgetRulesRequest   *UpdateSPBudgetRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateBudgetRulesForSPCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateBudgetRulesForSPCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateBudgetRulesForSPCampaignsRequest) UpdateSPBudgetRulesRequest(updateSPBudgetRulesRequest UpdateSPBudgetRulesRequest) ApiUpdateBudgetRulesForSPCampaignsRequest {
	r.updateSPBudgetRulesRequest = &updateSPBudgetRulesRequest
	return r
}

func (r ApiUpdateBudgetRulesForSPCampaignsRequest) Execute() (*UpdateBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.UpdateBudgetRulesForSPCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) UpdateBudgetRulesForSPCampaigns(ctx context.Context) ApiUpdateBudgetRulesForSPCampaignsRequest {
	return ApiUpdateBudgetRulesForSPCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateBudgetRulesResponse
func (a *BudgetRulesAPIService) UpdateBudgetRulesForSPCampaignsExecute(r ApiUpdateBudgetRulesForSPCampaignsRequest) (*UpdateBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.UpdateBudgetRulesForSPCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/budgetRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSPBudgetRulesRequest == nil {
		return localVarReturnValue, nil, reportError("updateSPBudgetRulesRequest is required and must be specified")
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
	localVarPostBody = r.updateSPBudgetRulesRequest
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
