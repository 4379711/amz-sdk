package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// RecommendationsAPIService RecommendationsAPI service
type RecommendationsAPIService service

type ApiGetBudgetRecommendationsRequest struct {
	ctx                                    context.Context
	ApiService                             *RecommendationsAPIService
	amazonAdvertisingAPIClientId           *string
	getBudgetRecommendationsRequestContent *GetBudgetRecommendationsRequestContent
	amazonAdvertisingAPIScope              *string
	amazonAdsAccountId                     *string
}

// The identifier of a client associated with a &#x60;Login with Amazon&#x60; account.
func (r ApiGetBudgetRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetBudgetRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiGetBudgetRecommendationsRequest) GetBudgetRecommendationsRequestContent(getBudgetRecommendationsRequestContent GetBudgetRecommendationsRequestContent) ApiGetBudgetRecommendationsRequest {
	r.getBudgetRecommendationsRequestContent = &getBudgetRecommendationsRequestContent
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetBudgetRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetBudgetRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The identifier of an account. Account must be a global advertising account.
func (r ApiGetBudgetRecommendationsRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiGetBudgetRecommendationsRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

func (r ApiGetBudgetRecommendationsRequest) Execute() (*GetBudgetRecommendationsResponseContent, *http.Response, error) {
	return r.ApiService.GetBudgetRecommendationsExecute(r)
}

func (a *RecommendationsAPIService) GetBudgetRecommendations(ctx context.Context) ApiGetBudgetRecommendationsRequest {
	return ApiGetBudgetRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBudgetRecommendationsResponseContent
func (a *RecommendationsAPIService) GetBudgetRecommendationsExecute(r ApiGetBudgetRecommendationsRequest) (*GetBudgetRecommendationsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBudgetRecommendationsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecommendationsAPIService.GetBudgetRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/campaigns/budgetRecommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.getBudgetRecommendationsRequestContent == nil {
		return localVarReturnValue, nil, reportError("getBudgetRecommendationsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbbudgetrecommendation.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbbudgetrecommendation.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	}
	// body params
	localVarPostBody = r.getBudgetRecommendationsRequestContent
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

type ApiGetHeadlineRecommendationsRequest struct {
	ctx                          context.Context
	ApiService                   *RecommendationsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	headlineSuggestionRequest    *HeadlineSuggestionRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetHeadlineRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetHeadlineRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetHeadlineRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetHeadlineRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetHeadlineRecommendationsRequest) HeadlineSuggestionRequest(headlineSuggestionRequest HeadlineSuggestionRequest) ApiGetHeadlineRecommendationsRequest {
	r.headlineSuggestionRequest = &headlineSuggestionRequest
	return r
}

func (r ApiGetHeadlineRecommendationsRequest) Execute() (*HeadlineSuggestionResponse, *http.Response, error) {
	return r.ApiService.GetHeadlineRecommendationsExecute(r)
}

func (a *RecommendationsAPIService) GetHeadlineRecommendations(ctx context.Context) ApiGetHeadlineRecommendationsRequest {
	return ApiGetHeadlineRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HeadlineSuggestionResponse
func (a *RecommendationsAPIService) GetHeadlineRecommendationsExecute(r ApiGetHeadlineRecommendationsRequest) (*HeadlineSuggestionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HeadlineSuggestionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecommendationsAPIService.GetHeadlineRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/recommendations/creative/headline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.headlineSuggestionRequest == nil {
		return localVarReturnValue, nil, reportError("headlineSuggestionRequest is required and must be specified")
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
	localVarPostBody = r.headlineSuggestionRequest
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

type ApiSBOptimizationRecommendationRequest struct {
	ctx                                        context.Context
	ApiService                                 *RecommendationsAPIService
	amazonAdvertisingAPIClientId               *string
	amazonAdvertisingAPIScope                  *string
	sBOptimizationRecommendationRequestContent *SBOptimizationRecommendationRequestContent
}

// The identifier of a client associated with a &#x60;Login with Amazon&#x60; account.
func (r ApiSBOptimizationRecommendationRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSBOptimizationRecommendationRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiSBOptimizationRecommendationRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSBOptimizationRecommendationRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSBOptimizationRecommendationRequest) SBOptimizationRecommendationRequestContent(sBOptimizationRecommendationRequestContent SBOptimizationRecommendationRequestContent) ApiSBOptimizationRecommendationRequest {
	r.sBOptimizationRecommendationRequestContent = &sBOptimizationRecommendationRequestContent
	return r
}

func (r ApiSBOptimizationRecommendationRequest) Execute() (*SBOptimizationRecommendationResponseContent, *http.Response, error) {
	return r.ApiService.SBOptimizationRecommendationExecute(r)
}

func (a *RecommendationsAPIService) SBOptimizationRecommendation(ctx context.Context) ApiSBOptimizationRecommendationRequest {
	return ApiSBOptimizationRecommendationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SBOptimizationRecommendationResponseContent
func (a *RecommendationsAPIService) SBOptimizationRecommendationExecute(r ApiSBOptimizationRecommendationRequest) (*SBOptimizationRecommendationResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBOptimizationRecommendationResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecommendationsAPIService.SBOptimizationRecommendation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/recommendations/optimization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sBOptimizationRecommendationRequestContent == nil {
		return localVarReturnValue, nil, reportError("sBOptimizationRecommendationRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sboptimizationrecommendationresource.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.sboptimizationrecommendationresource.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sBOptimizationRecommendationRequestContent
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

type ApiSBTargetingGetNegativeBrandsRequest struct {
	ctx                          context.Context
	ApiService                   *RecommendationsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	nextToken                    *string
}

// The identifier of a client associated with a &#x60;Login with Amazon&#x60; account.
func (r ApiSBTargetingGetNegativeBrandsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSBTargetingGetNegativeBrandsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiSBTargetingGetNegativeBrandsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSBTargetingGetNegativeBrandsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results.
func (r ApiSBTargetingGetNegativeBrandsRequest) NextToken(nextToken string) ApiSBTargetingGetNegativeBrandsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiSBTargetingGetNegativeBrandsRequest) Execute() (*SBTargetingGetNegativeBrandsResponseContent, *http.Response, error) {
	return r.ApiService.SBTargetingGetNegativeBrandsExecute(r)
}

func (a *RecommendationsAPIService) SBTargetingGetNegativeBrands(ctx context.Context) ApiSBTargetingGetNegativeBrandsRequest {
	return ApiSBTargetingGetNegativeBrandsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SBTargetingGetNegativeBrandsResponseContent
func (a *RecommendationsAPIService) SBTargetingGetNegativeBrandsExecute(r ApiSBTargetingGetNegativeBrandsRequest) (*SBTargetingGetNegativeBrandsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBTargetingGetNegativeBrandsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecommendationsAPIService.SBTargetingGetNegativeBrands")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/negativeTargets/brands/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbtargeting.v4+json"}

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
