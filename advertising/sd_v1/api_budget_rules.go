package sd_v1

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

type ApiCreateBudgetRulesForSDCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createSDBudgetRulesRequest   *CreateSDBudgetRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBudgetRulesForSDCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateBudgetRulesForSDCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBudgetRulesForSDCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateBudgetRulesForSDCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateBudgetRulesForSDCampaignsRequest) CreateSDBudgetRulesRequest(createSDBudgetRulesRequest CreateSDBudgetRulesRequest) ApiCreateBudgetRulesForSDCampaignsRequest {
	r.createSDBudgetRulesRequest = &createSDBudgetRulesRequest
	return r
}

func (r ApiCreateBudgetRulesForSDCampaignsRequest) Execute() (*CreateBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.CreateBudgetRulesForSDCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) CreateBudgetRulesForSDCampaigns(ctx context.Context) ApiCreateBudgetRulesForSDCampaignsRequest {
	return ApiCreateBudgetRulesForSDCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateBudgetRulesResponse
func (a *BudgetRulesAPIService) CreateBudgetRulesForSDCampaignsExecute(r ApiCreateBudgetRulesForSDCampaignsRequest) (*CreateBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.CreateBudgetRulesForSDCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/budgetRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSDBudgetRulesRequest == nil {
		return localVarReturnValue, nil, reportError("createSDBudgetRulesRequest is required and must be specified")
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
	localVarPostBody = r.createSDBudgetRulesRequest
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

type ApiGetBudgetRuleByRuleIdForSDCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	budgetRuleId                 string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetBudgetRuleByRuleIdForSDCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetBudgetRuleByRuleIdForSDCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetBudgetRuleByRuleIdForSDCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetBudgetRuleByRuleIdForSDCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetBudgetRuleByRuleIdForSDCampaignsRequest) Execute() (*GetSDBudgetRuleResponse, *http.Response, error) {
	return r.ApiService.GetBudgetRuleByRuleIdForSDCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) GetBudgetRuleByRuleIdForSDCampaigns(ctx context.Context, budgetRuleId string) ApiGetBudgetRuleByRuleIdForSDCampaignsRequest {
	return ApiGetBudgetRuleByRuleIdForSDCampaignsRequest{
		ApiService:   a,
		ctx:          ctx,
		budgetRuleId: budgetRuleId,
	}
}

// Execute executes the request
//
//	@return GetSDBudgetRuleResponse
func (a *BudgetRulesAPIService) GetBudgetRuleByRuleIdForSDCampaignsExecute(r ApiGetBudgetRuleByRuleIdForSDCampaignsRequest) (*GetSDBudgetRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSDBudgetRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetBudgetRuleByRuleIdForSDCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/budgetRules/{budgetRuleId}"
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

type ApiGetSDBudgetRulesForAdvertiserRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	pageSize                     *float32
	nextToken                    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetSDBudgetRulesForAdvertiserRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetSDBudgetRulesForAdvertiserRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetSDBudgetRulesForAdvertiserRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetSDBudgetRulesForAdvertiserRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30.
func (r ApiGetSDBudgetRulesForAdvertiserRequest) PageSize(pageSize float32) ApiGetSDBudgetRulesForAdvertiserRequest {
	r.pageSize = &pageSize
	return r
}

// To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results.
func (r ApiGetSDBudgetRulesForAdvertiserRequest) NextToken(nextToken string) ApiGetSDBudgetRulesForAdvertiserRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetSDBudgetRulesForAdvertiserRequest) Execute() (*GetSDBudgetRulesForAdvertiserResponse, *http.Response, error) {
	return r.ApiService.GetSDBudgetRulesForAdvertiserExecute(r)
}

func (a *BudgetRulesAPIService) GetSDBudgetRulesForAdvertiser(ctx context.Context) ApiGetSDBudgetRulesForAdvertiserRequest {
	return ApiGetSDBudgetRulesForAdvertiserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSDBudgetRulesForAdvertiserResponse
func (a *BudgetRulesAPIService) GetSDBudgetRulesForAdvertiserExecute(r ApiGetSDBudgetRulesForAdvertiserRequest) (*GetSDBudgetRulesForAdvertiserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSDBudgetRulesForAdvertiserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.GetSDBudgetRulesForAdvertiser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/budgetRules"

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

type ApiUpdateBudgetRulesForSDCampaignsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateSDBudgetRulesRequest   *UpdateSDBudgetRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateBudgetRulesForSDCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateBudgetRulesForSDCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateBudgetRulesForSDCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateBudgetRulesForSDCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateBudgetRulesForSDCampaignsRequest) UpdateSDBudgetRulesRequest(updateSDBudgetRulesRequest UpdateSDBudgetRulesRequest) ApiUpdateBudgetRulesForSDCampaignsRequest {
	r.updateSDBudgetRulesRequest = &updateSDBudgetRulesRequest
	return r
}

func (r ApiUpdateBudgetRulesForSDCampaignsRequest) Execute() (*UpdateBudgetRulesResponse, *http.Response, error) {
	return r.ApiService.UpdateBudgetRulesForSDCampaignsExecute(r)
}

func (a *BudgetRulesAPIService) UpdateBudgetRulesForSDCampaigns(ctx context.Context) ApiUpdateBudgetRulesForSDCampaignsRequest {
	return ApiUpdateBudgetRulesForSDCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateBudgetRulesResponse
func (a *BudgetRulesAPIService) UpdateBudgetRulesForSDCampaignsExecute(r ApiUpdateBudgetRulesForSDCampaignsRequest) (*UpdateBudgetRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateBudgetRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesAPIService.UpdateBudgetRulesForSDCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/budgetRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSDBudgetRulesRequest == nil {
		return localVarReturnValue, nil, reportError("updateSDBudgetRulesRequest is required and must be specified")
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
	localVarPostBody = r.updateSDBudgetRulesRequest
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
