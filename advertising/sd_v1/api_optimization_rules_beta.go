package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// OptimizationRulesBetaAPIService OptimizationRulesBetaAPI service
type OptimizationRulesBetaAPIService service

type ApiAssociateOptimizationRulesWithAdGroupRequest struct {
	ctx                                      context.Context
	ApiService                               *OptimizationRulesBetaAPIService
	amazonAdvertisingAPIClientId             *string
	amazonAdvertisingAPIScope                *string
	adGroupId                                int64
	createAssociatedOptimizationRulesRequest *CreateAssociatedOptimizationRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiAssociateOptimizationRulesWithAdGroupRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiAssociateOptimizationRulesWithAdGroupRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiAssociateOptimizationRulesWithAdGroupRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiAssociateOptimizationRulesWithAdGroupRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of optimization rule identifiers. Only one optimization rule identifier is currently supported per request. This note will be removed when multiple rule identifiers are supported.  For each ad group, only one optimization rule metric name is supported, based on the ad group&#39;s &#x60;bidOptimization&#x60; type. Refer to the following table for the metric names supported for each type. |  AdGroup.bidOptimization |     Supported OptimizationRule.metricName       | |------------------|--------------------| |   reach       | COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS  | |   clicks      | COST_PER_CLICK          | |  conversions  | COST_PER_ORDER          |
func (r ApiAssociateOptimizationRulesWithAdGroupRequest) CreateAssociatedOptimizationRulesRequest(createAssociatedOptimizationRulesRequest CreateAssociatedOptimizationRulesRequest) ApiAssociateOptimizationRulesWithAdGroupRequest {
	r.createAssociatedOptimizationRulesRequest = &createAssociatedOptimizationRulesRequest
	return r
}

func (r ApiAssociateOptimizationRulesWithAdGroupRequest) Execute() ([]OptimizationRuleResponse, *http.Response, error) {
	return r.ApiService.AssociateOptimizationRulesWithAdGroupExecute(r)
}

func (a *OptimizationRulesBetaAPIService) AssociateOptimizationRulesWithAdGroup(ctx context.Context, adGroupId int64) ApiAssociateOptimizationRulesWithAdGroupRequest {
	return ApiAssociateOptimizationRulesWithAdGroupRequest{
		ApiService: a,
		ctx:        ctx,
		adGroupId:  adGroupId,
	}
}

// Execute executes the request
//
//	@return []OptimizationRuleResponse
func (a *OptimizationRulesBetaAPIService) AssociateOptimizationRulesWithAdGroupExecute(r ApiAssociateOptimizationRulesWithAdGroupRequest) ([]OptimizationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []OptimizationRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesBetaAPIService.AssociateOptimizationRulesWithAdGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups/{adGroupId}/optimizationRules"
	localVarPath = strings.Replace(localVarPath, "{"+"adGroupId"+"}", url.PathEscape(parameterValueToString(r.adGroupId, "adGroupId")), -1)

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
	localVarPostBody = r.createAssociatedOptimizationRulesRequest
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
	ctx                          context.Context
	ApiService                   *OptimizationRulesBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createOptimizationRule       *[]CreateOptimizationRule
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of OptimizationRule objects. For each object, specify required fields and their values. Required fields are &#x60;state&#x60; and &#x60;ruleConditions&#x60;.
func (r ApiCreateOptimizationRulesRequest) CreateOptimizationRule(createOptimizationRule []CreateOptimizationRule) ApiCreateOptimizationRulesRequest {
	r.createOptimizationRule = &createOptimizationRule
	return r
}

func (r ApiCreateOptimizationRulesRequest) Execute() ([]OptimizationRuleResponse, *http.Response, error) {
	return r.ApiService.CreateOptimizationRulesExecute(r)
}

func (a *OptimizationRulesBetaAPIService) CreateOptimizationRules(ctx context.Context) ApiCreateOptimizationRulesRequest {
	return ApiCreateOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []OptimizationRuleResponse
func (a *OptimizationRulesBetaAPIService) CreateOptimizationRulesExecute(r ApiCreateOptimizationRulesRequest) ([]OptimizationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []OptimizationRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesBetaAPIService.CreateOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/optimizationRules"

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
	localVarPostBody = r.createOptimizationRule
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

type ApiDisassociateOptimizationRulesFromAdGroupRequest struct {
	ctx                                      context.Context
	ApiService                               *OptimizationRulesBetaAPIService
	amazonAdvertisingAPIClientId             *string
	amazonAdvertisingAPIScope                *string
	adGroupId                                int64
	createAssociatedOptimizationRulesRequest *CreateAssociatedOptimizationRulesRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDisassociateOptimizationRulesFromAdGroupRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDisassociateOptimizationRulesFromAdGroupRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDisassociateOptimizationRulesFromAdGroupRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDisassociateOptimizationRulesFromAdGroupRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of optimization rule identifiers. Only one optimization rule identifier is currently supported per request. This note will be removed when multiple rule identifiers are supported.
func (r ApiDisassociateOptimizationRulesFromAdGroupRequest) CreateAssociatedOptimizationRulesRequest(createAssociatedOptimizationRulesRequest CreateAssociatedOptimizationRulesRequest) ApiDisassociateOptimizationRulesFromAdGroupRequest {
	r.createAssociatedOptimizationRulesRequest = &createAssociatedOptimizationRulesRequest
	return r
}

func (r ApiDisassociateOptimizationRulesFromAdGroupRequest) Execute() (*OptimizationRuleAssociationResponse, *http.Response, error) {
	return r.ApiService.DisassociateOptimizationRulesFromAdGroupExecute(r)
}

func (a *OptimizationRulesBetaAPIService) DisassociateOptimizationRulesFromAdGroup(ctx context.Context, adGroupId int64) ApiDisassociateOptimizationRulesFromAdGroupRequest {
	return ApiDisassociateOptimizationRulesFromAdGroupRequest{
		ApiService: a,
		ctx:        ctx,
		adGroupId:  adGroupId,
	}
}

// Execute executes the request
//
//	@return OptimizationRuleAssociationResponse
func (a *OptimizationRulesBetaAPIService) DisassociateOptimizationRulesFromAdGroupExecute(r ApiDisassociateOptimizationRulesFromAdGroupRequest) (*OptimizationRuleAssociationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OptimizationRuleAssociationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesBetaAPIService.DisassociateOptimizationRulesFromAdGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups/{adGroupId}/optimizationRules/disassociate"
	localVarPath = strings.Replace(localVarPath, "{"+"adGroupId"+"}", url.PathEscape(parameterValueToString(r.adGroupId, "adGroupId")), -1)

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
	localVarPostBody = r.createAssociatedOptimizationRulesRequest
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

type ApiListOptimizationRulesRequest struct {
	ctx                          context.Context
	ApiService                   *OptimizationRulesBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	name                         *string
	optimizationRuleIdFilter     *string
	adGroupIdFilter              *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of optimization rules. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListOptimizationRulesRequest) StartIndex(startIndex int32) ApiListOptimizationRulesRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of OptimizationRule objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten optimization rules set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten optimization rules, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListOptimizationRulesRequest) Count(count int32) ApiListOptimizationRulesRequest {
	r.count = &count
	return r
}

// Optional. The returned array is filtered to include only optimization rules with state set to one of the values in the specified comma-delimited list. Available values:   - enabled   - paused [COMING LATER]   - enabled, paused [COMING LATER]
func (r ApiListOptimizationRulesRequest) StateFilter(stateFilter string) ApiListOptimizationRulesRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. The returned array includes only optimization rules with the specified name using an exact string match.
func (r ApiListOptimizationRulesRequest) Name(name string) ApiListOptimizationRulesRequest {
	r.name = &name
	return r
}

// Optional. The returned array is filtered to include only optimization rules associated with the optimization rule identifiers in the specified comma-delimited list.  Maximum size limit 50.
func (r ApiListOptimizationRulesRequest) OptimizationRuleIdFilter(optimizationRuleIdFilter string) ApiListOptimizationRulesRequest {
	r.optimizationRuleIdFilter = &optimizationRuleIdFilter
	return r
}

// Optional. The returned array is filtered to include only optimization rules associated with the ad group identifiers in the comma-delimited list.  Maximum size limit 50.
func (r ApiListOptimizationRulesRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListOptimizationRulesRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

func (r ApiListOptimizationRulesRequest) Execute() ([]OptimizationRule, *http.Response, error) {
	return r.ApiService.ListOptimizationRulesExecute(r)
}

func (a *OptimizationRulesBetaAPIService) ListOptimizationRules(ctx context.Context) ApiListOptimizationRulesRequest {
	return ApiListOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []OptimizationRule
func (a *OptimizationRulesBetaAPIService) ListOptimizationRulesExecute(r ApiListOptimizationRulesRequest) ([]OptimizationRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []OptimizationRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesBetaAPIService.ListOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/optimizationRules"

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
		var defaultValue string = "enabled"
		r.stateFilter = &defaultValue
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.optimizationRuleIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "optimizationRuleIdFilter", r.optimizationRuleIdFilter, "form", "")
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
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

type ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest struct {
	ctx                          context.Context
	ApiService                   *OptimizationRulesBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	adGroupId                    int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest) Execute() ([]OptimizationRule, *http.Response, error) {
	return r.ApiService.SdAdGroupsAdGroupIdOptimizationRulesGetExecute(r)
}

func (a *OptimizationRulesBetaAPIService) SdAdGroupsAdGroupIdOptimizationRulesGet(ctx context.Context, adGroupId int64) ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest {
	return ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest{
		ApiService: a,
		ctx:        ctx,
		adGroupId:  adGroupId,
	}
}

// Execute executes the request
//
//	@return []OptimizationRule
func (a *OptimizationRulesBetaAPIService) SdAdGroupsAdGroupIdOptimizationRulesGetExecute(r ApiSdAdGroupsAdGroupIdOptimizationRulesGetRequest) ([]OptimizationRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []OptimizationRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesBetaAPIService.SdAdGroupsAdGroupIdOptimizationRulesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/adGroups/{adGroupId}/optimizationRules"
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

type ApiSdOptimizationRulesOptimizationRuleIdGetRequest struct {
	ctx                          context.Context
	ApiService                   *OptimizationRulesBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	optimizationRuleId           string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiSdOptimizationRulesOptimizationRuleIdGetRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSdOptimizationRulesOptimizationRuleIdGetRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiSdOptimizationRulesOptimizationRuleIdGetRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSdOptimizationRulesOptimizationRuleIdGetRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSdOptimizationRulesOptimizationRuleIdGetRequest) Execute() (*OptimizationRule, *http.Response, error) {
	return r.ApiService.SdOptimizationRulesOptimizationRuleIdGetExecute(r)
}

func (a *OptimizationRulesBetaAPIService) SdOptimizationRulesOptimizationRuleIdGet(ctx context.Context, optimizationRuleId string) ApiSdOptimizationRulesOptimizationRuleIdGetRequest {
	return ApiSdOptimizationRulesOptimizationRuleIdGetRequest{
		ApiService:         a,
		ctx:                ctx,
		optimizationRuleId: optimizationRuleId,
	}
}

// Execute executes the request
//
//	@return OptimizationRule
func (a *OptimizationRulesBetaAPIService) SdOptimizationRulesOptimizationRuleIdGetExecute(r ApiSdOptimizationRulesOptimizationRuleIdGetRequest) (*OptimizationRule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OptimizationRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesBetaAPIService.SdOptimizationRulesOptimizationRuleIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/optimizationRules/{optimizationRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationRuleId"+"}", url.PathEscape(parameterValueToString(r.optimizationRuleId, "optimizationRuleId")), -1)

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

type ApiUpdateOptimizationRulesRequest struct {
	ctx                          context.Context
	ApiService                   *OptimizationRulesBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateOptimizationRule       *[]UpdateOptimizationRule
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of OptimizationRule objects. For each object, specify an optimization rule identifier and mutable fields with their updated values. The mutable fields are &#x60;ruleName&#x60;, &#x60;state&#x60;, and &#x60;ruleConditions&#x60;.
func (r ApiUpdateOptimizationRulesRequest) UpdateOptimizationRule(updateOptimizationRule []UpdateOptimizationRule) ApiUpdateOptimizationRulesRequest {
	r.updateOptimizationRule = &updateOptimizationRule
	return r
}

func (r ApiUpdateOptimizationRulesRequest) Execute() ([]OptimizationRuleResponse, *http.Response, error) {
	return r.ApiService.UpdateOptimizationRulesExecute(r)
}

func (a *OptimizationRulesBetaAPIService) UpdateOptimizationRules(ctx context.Context) ApiUpdateOptimizationRulesRequest {
	return ApiUpdateOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []OptimizationRuleResponse
func (a *OptimizationRulesBetaAPIService) UpdateOptimizationRulesExecute(r ApiUpdateOptimizationRulesRequest) ([]OptimizationRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []OptimizationRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesBetaAPIService.UpdateOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/optimizationRules"

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
	localVarPostBody = r.updateOptimizationRule
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
