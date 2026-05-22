package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// OptimizationRulesAPIService OptimizationRulesAPI service
type OptimizationRulesAPIService service

type ApiAssociateOptimizationRulesToCampaignRequest struct {
	ctx                                                             context.Context
	ApiService                                                      *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                                    *string
	amazonAdvertisingAPIScope                                       *string
	campaignId                                                      string
	optimizationRulesAPIAssociateOptimizationRulesToCampaignRequest *OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiAssociateOptimizationRulesToCampaignRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiAssociateOptimizationRulesToCampaignRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiAssociateOptimizationRulesToCampaignRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiAssociateOptimizationRulesToCampaignRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiAssociateOptimizationRulesToCampaignRequest) OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest(optimizationRulesAPIAssociateOptimizationRulesToCampaignRequest OptimizationRulesAPIAssociateOptimizationRulesToCampaignRequest) ApiAssociateOptimizationRulesToCampaignRequest {
	r.optimizationRulesAPIAssociateOptimizationRulesToCampaignRequest = &optimizationRulesAPIAssociateOptimizationRulesToCampaignRequest
	return r
}

func (r ApiAssociateOptimizationRulesToCampaignRequest) Execute() (*OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse, *http.Response, error) {
	return r.ApiService.AssociateOptimizationRulesToCampaignExecute(r)
}

func (a *OptimizationRulesAPIService) AssociateOptimizationRulesToCampaign(ctx context.Context, campaignId string) ApiAssociateOptimizationRulesToCampaignRequest {
	return ApiAssociateOptimizationRulesToCampaignRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
	}
}

// Execute executes the request
//
//	@return OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse
func (a *OptimizationRulesAPIService) AssociateOptimizationRulesToCampaignExecute(r ApiAssociateOptimizationRulesToCampaignRequest) (*OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OptimizationRulesAPIAssociateOptimizationRulesToCampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.AssociateOptimizationRulesToCampaign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/{campaignId}/optimizationRules"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterValueToString(r.campaignId, "campaignId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spoptimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spoptimizationrules.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.optimizationRulesAPIAssociateOptimizationRulesToCampaignRequest
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

type ApiCreateOptimizationRulesRequest struct {
	ctx                                                context.Context
	ApiService                                         *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                       *string
	amazonAdvertisingAPIScope                          *string
	optimizationRulesAPICreateOptimizationRulesRequest *OptimizationRulesAPICreateOptimizationRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateOptimizationRulesRequest) OptimizationRulesAPICreateOptimizationRulesRequest(optimizationRulesAPICreateOptimizationRulesRequest OptimizationRulesAPICreateOptimizationRulesRequest) ApiCreateOptimizationRulesRequest {
	r.optimizationRulesAPICreateOptimizationRulesRequest = &optimizationRulesAPICreateOptimizationRulesRequest
	return r
}

func (r ApiCreateOptimizationRulesRequest) Execute() (*OptimizationRulesAPIOptimizationRulesResponse, *http.Response, error) {
	return r.ApiService.CreateOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) CreateOptimizationRules(ctx context.Context) ApiCreateOptimizationRulesRequest {
	return ApiCreateOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptimizationRulesAPIOptimizationRulesResponse
func (a *OptimizationRulesAPIService) CreateOptimizationRulesExecute(r ApiCreateOptimizationRulesRequest) (*OptimizationRulesAPIOptimizationRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OptimizationRulesAPIOptimizationRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.CreateOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/optimization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spoptimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spoptimizationrules.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.optimizationRulesAPICreateOptimizationRulesRequest
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

type ApiSearchOptimizationRulesRequest struct {
	ctx                                                context.Context
	ApiService                                         *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                       *string
	amazonAdvertisingAPIScope                          *string
	optimizationRulesAPISearchOptimizationRulesRequest *OptimizationRulesAPISearchOptimizationRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiSearchOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSearchOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiSearchOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSearchOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSearchOptimizationRulesRequest) OptimizationRulesAPISearchOptimizationRulesRequest(optimizationRulesAPISearchOptimizationRulesRequest OptimizationRulesAPISearchOptimizationRulesRequest) ApiSearchOptimizationRulesRequest {
	r.optimizationRulesAPISearchOptimizationRulesRequest = &optimizationRulesAPISearchOptimizationRulesRequest
	return r
}

func (r ApiSearchOptimizationRulesRequest) Execute() (*OptimizationRulesAPISearchOptimizationRulesResponse, *http.Response, error) {
	return r.ApiService.SearchOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) SearchOptimizationRules(ctx context.Context) ApiSearchOptimizationRulesRequest {
	return ApiSearchOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptimizationRulesAPISearchOptimizationRulesResponse
func (a *OptimizationRulesAPIService) SearchOptimizationRulesExecute(r ApiSearchOptimizationRulesRequest) (*OptimizationRulesAPISearchOptimizationRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OptimizationRulesAPISearchOptimizationRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.SearchOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/optimization/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spoptimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spoptimizationrules.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.optimizationRulesAPISearchOptimizationRulesRequest
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

type ApiUpdateOptimizationRulesRequest struct {
	ctx                                                context.Context
	ApiService                                         *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                       *string
	amazonAdvertisingAPIScope                          *string
	optimizationRulesAPIUpdateOptimizationRulesRequest *OptimizationRulesAPIUpdateOptimizationRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateOptimizationRulesRequest) OptimizationRulesAPIUpdateOptimizationRulesRequest(optimizationRulesAPIUpdateOptimizationRulesRequest OptimizationRulesAPIUpdateOptimizationRulesRequest) ApiUpdateOptimizationRulesRequest {
	r.optimizationRulesAPIUpdateOptimizationRulesRequest = &optimizationRulesAPIUpdateOptimizationRulesRequest
	return r
}

func (r ApiUpdateOptimizationRulesRequest) Execute() (*OptimizationRulesAPIOptimizationRulesResponse, *http.Response, error) {
	return r.ApiService.UpdateOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) UpdateOptimizationRules(ctx context.Context) ApiUpdateOptimizationRulesRequest {
	return ApiUpdateOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptimizationRulesAPIOptimizationRulesResponse
func (a *OptimizationRulesAPIService) UpdateOptimizationRulesExecute(r ApiUpdateOptimizationRulesRequest) (*OptimizationRulesAPIOptimizationRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OptimizationRulesAPIOptimizationRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.UpdateOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/optimization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spoptimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spoptimizationrules.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.optimizationRulesAPIUpdateOptimizationRulesRequest
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
