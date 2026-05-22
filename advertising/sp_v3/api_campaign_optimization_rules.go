package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CampaignOptimizationRulesAPIService CampaignOptimizationRulesAPI service
type CampaignOptimizationRulesAPIService service

type ApiCreateOptimizationRuleRequest struct {
	ctx                                      context.Context
	ApiService                               *CampaignOptimizationRulesAPIService
	amazonAdvertisingAPIClientId             *string
	amazonAdvertisingAPIScope                *string
	createSPCampaignOptimizationRulesRequest *CreateSPCampaignOptimizationRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateOptimizationRuleRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateOptimizationRuleRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateOptimizationRuleRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateOptimizationRuleRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateOptimizationRuleRequest) CreateSPCampaignOptimizationRulesRequest(createSPCampaignOptimizationRulesRequest CreateSPCampaignOptimizationRulesRequest) ApiCreateOptimizationRuleRequest {
	r.createSPCampaignOptimizationRulesRequest = &createSPCampaignOptimizationRulesRequest
	return r
}

func (r ApiCreateOptimizationRuleRequest) Execute() (*CreateSPCampaignOptimizationRulesResponse, *http.Response, error) {
	return r.ApiService.CreateOptimizationRuleExecute(r)
}

func (a *CampaignOptimizationRulesAPIService) CreateOptimizationRule(ctx context.Context) ApiCreateOptimizationRuleRequest {
	return ApiCreateOptimizationRuleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSPCampaignOptimizationRulesResponse
func (a *CampaignOptimizationRulesAPIService) CreateOptimizationRuleExecute(r ApiCreateOptimizationRuleRequest) (*CreateSPCampaignOptimizationRulesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSPCampaignOptimizationRulesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignOptimizationRulesAPIService.CreateOptimizationRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/campaignOptimization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSPCampaignOptimizationRulesRequest == nil {
		return localVarReturnValue, nil, reportError("createSPCampaignOptimizationRulesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.optimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.optimizationrules.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSPCampaignOptimizationRulesRequest
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

type ApiDeleteCampaignOptimizationRuleRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignOptimizationRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignOptimizationId       string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiDeleteCampaignOptimizationRuleRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteCampaignOptimizationRuleRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiDeleteCampaignOptimizationRuleRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteCampaignOptimizationRuleRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteCampaignOptimizationRuleRequest) Execute() (*DeleteSPCampaignOptimizationRuleResponse, *http.Response, error) {
	return r.ApiService.DeleteCampaignOptimizationRuleExecute(r)
}

func (a *CampaignOptimizationRulesAPIService) DeleteCampaignOptimizationRule(ctx context.Context, campaignOptimizationId string) ApiDeleteCampaignOptimizationRuleRequest {
	return ApiDeleteCampaignOptimizationRuleRequest{
		ApiService:             a,
		ctx:                    ctx,
		campaignOptimizationId: campaignOptimizationId,
	}
}

// Execute executes the request
//
//	@return DeleteSPCampaignOptimizationRuleResponse
func (a *CampaignOptimizationRulesAPIService) DeleteCampaignOptimizationRuleExecute(r ApiDeleteCampaignOptimizationRuleRequest) (*DeleteSPCampaignOptimizationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteSPCampaignOptimizationRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignOptimizationRulesAPIService.DeleteCampaignOptimizationRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/campaignOptimization/{campaignOptimizationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignOptimizationId"+"}", url.PathEscape(parameterValueToString(r.campaignOptimizationId, "campaignOptimizationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.optimizationrules.v1+json", "application/json"}

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

type ApiGetCampaignOptimizationRuleRequest struct {
	ctx                          context.Context
	ApiService                   *CampaignOptimizationRulesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	campaignOptimizationId       string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCampaignOptimizationRuleRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetCampaignOptimizationRuleRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCampaignOptimizationRuleRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetCampaignOptimizationRuleRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetCampaignOptimizationRuleRequest) Execute() (*GetSPCampaignOptimizationRuleResponse, *http.Response, error) {
	return r.ApiService.GetCampaignOptimizationRuleExecute(r)
}

func (a *CampaignOptimizationRulesAPIService) GetCampaignOptimizationRule(ctx context.Context, campaignOptimizationId string) ApiGetCampaignOptimizationRuleRequest {
	return ApiGetCampaignOptimizationRuleRequest{
		ApiService:             a,
		ctx:                    ctx,
		campaignOptimizationId: campaignOptimizationId,
	}
}

// Execute executes the request
//
//	@return GetSPCampaignOptimizationRuleResponse
func (a *CampaignOptimizationRulesAPIService) GetCampaignOptimizationRuleExecute(r ApiGetCampaignOptimizationRuleRequest) (*GetSPCampaignOptimizationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSPCampaignOptimizationRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignOptimizationRulesAPIService.GetCampaignOptimizationRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/campaignOptimization/{campaignOptimizationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignOptimizationId"+"}", url.PathEscape(parameterValueToString(r.campaignOptimizationId, "campaignOptimizationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.optimizationrules.v1+json", "application/json"}

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

type ApiGetOptimizationRuleEligibilityRequest struct {
	ctx                                             context.Context
	ApiService                                      *CampaignOptimizationRulesAPIService
	amazonAdvertisingAPIClientId                    *string
	amazonAdvertisingAPIScope                       *string
	sPCampaignOptimizationRecommendationsAPIRequest *SPCampaignOptimizationRecommendationsAPIRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetOptimizationRuleEligibilityRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetOptimizationRuleEligibilityRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetOptimizationRuleEligibilityRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetOptimizationRuleEligibilityRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetOptimizationRuleEligibilityRequest) SPCampaignOptimizationRecommendationsAPIRequest(sPCampaignOptimizationRecommendationsAPIRequest SPCampaignOptimizationRecommendationsAPIRequest) ApiGetOptimizationRuleEligibilityRequest {
	r.sPCampaignOptimizationRecommendationsAPIRequest = &sPCampaignOptimizationRecommendationsAPIRequest
	return r
}

func (r ApiGetOptimizationRuleEligibilityRequest) Execute() (*SPCampaignOptimizationRecommendationAPIResponse, *http.Response, error) {
	return r.ApiService.GetOptimizationRuleEligibilityExecute(r)
}

func (a *CampaignOptimizationRulesAPIService) GetOptimizationRuleEligibility(ctx context.Context) ApiGetOptimizationRuleEligibilityRequest {
	return ApiGetOptimizationRuleEligibilityRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SPCampaignOptimizationRecommendationAPIResponse
func (a *CampaignOptimizationRulesAPIService) GetOptimizationRuleEligibilityExecute(r ApiGetOptimizationRuleEligibilityRequest) (*SPCampaignOptimizationRecommendationAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SPCampaignOptimizationRecommendationAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignOptimizationRulesAPIService.GetOptimizationRuleEligibility")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/campaignOptimization/eligibility"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sPCampaignOptimizationRecommendationsAPIRequest == nil {
		return localVarReturnValue, nil, reportError("sPCampaignOptimizationRecommendationsAPIRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.optimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.optimizationrules.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sPCampaignOptimizationRecommendationsAPIRequest
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

type ApiGetRuleNotificationRequest struct {
	ctx                                          context.Context
	ApiService                                   *CampaignOptimizationRulesAPIService
	amazonAdvertisingAPIClientId                 *string
	amazonAdvertisingAPIScope                    *string
	sPCampaignOptimizationNotificationAPIRequest *SPCampaignOptimizationNotificationAPIRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetRuleNotificationRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetRuleNotificationRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetRuleNotificationRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetRuleNotificationRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetRuleNotificationRequest) SPCampaignOptimizationNotificationAPIRequest(sPCampaignOptimizationNotificationAPIRequest SPCampaignOptimizationNotificationAPIRequest) ApiGetRuleNotificationRequest {
	r.sPCampaignOptimizationNotificationAPIRequest = &sPCampaignOptimizationNotificationAPIRequest
	return r
}

func (r ApiGetRuleNotificationRequest) Execute() (*SPCampaignOptimizationNotificationAPIResponse, *http.Response, error) {
	return r.ApiService.GetRuleNotificationExecute(r)
}

func (a *CampaignOptimizationRulesAPIService) GetRuleNotification(ctx context.Context) ApiGetRuleNotificationRequest {
	return ApiGetRuleNotificationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SPCampaignOptimizationNotificationAPIResponse
func (a *CampaignOptimizationRulesAPIService) GetRuleNotificationExecute(r ApiGetRuleNotificationRequest) (*SPCampaignOptimizationNotificationAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SPCampaignOptimizationNotificationAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignOptimizationRulesAPIService.GetRuleNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/campaignOptimization/state"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sPCampaignOptimizationNotificationAPIRequest == nil {
		return localVarReturnValue, nil, reportError("sPCampaignOptimizationNotificationAPIRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.optimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.optimizationrules.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sPCampaignOptimizationNotificationAPIRequest
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

type ApiUpdateOptimizationRuleRequest struct {
	ctx                                      context.Context
	ApiService                               *CampaignOptimizationRulesAPIService
	amazonAdvertisingAPIClientId             *string
	amazonAdvertisingAPIScope                *string
	updateSPCampaignOptimizationRulesRequest *UpdateSPCampaignOptimizationRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateOptimizationRuleRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateOptimizationRuleRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiUpdateOptimizationRuleRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateOptimizationRuleRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateOptimizationRuleRequest) UpdateSPCampaignOptimizationRulesRequest(updateSPCampaignOptimizationRulesRequest UpdateSPCampaignOptimizationRulesRequest) ApiUpdateOptimizationRuleRequest {
	r.updateSPCampaignOptimizationRulesRequest = &updateSPCampaignOptimizationRulesRequest
	return r
}

func (r ApiUpdateOptimizationRuleRequest) Execute() (*UpdateSPCampaignOptimizationRuleResponse, *http.Response, error) {
	return r.ApiService.UpdateOptimizationRuleExecute(r)
}

func (a *CampaignOptimizationRulesAPIService) UpdateOptimizationRule(ctx context.Context) ApiUpdateOptimizationRuleRequest {
	return ApiUpdateOptimizationRuleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateSPCampaignOptimizationRuleResponse
func (a *CampaignOptimizationRulesAPIService) UpdateOptimizationRuleExecute(r ApiUpdateOptimizationRuleRequest) (*UpdateSPCampaignOptimizationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSPCampaignOptimizationRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignOptimizationRulesAPIService.UpdateOptimizationRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/rules/campaignOptimization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSPCampaignOptimizationRulesRequest == nil {
		return localVarReturnValue, nil, reportError("updateSPCampaignOptimizationRulesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.optimizationrules.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.optimizationrules.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.updateSPCampaignOptimizationRulesRequest
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
