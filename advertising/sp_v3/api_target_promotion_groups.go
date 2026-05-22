package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// TargetPromotionGroupsAPIService TargetPromotionGroupsAPI service
type TargetPromotionGroupsAPIService service

type ApiCreateTargetPromotionGroupTargetsRequest struct {
	ctx                                                              context.Context
	ApiService                                                       *TargetPromotionGroupsAPIService
	amazonAdvertisingAPIClientId                                     *string
	amazonAdvertisingAPIScope                                        *string
	sponsoredProductsCreateTargetPromotionGroupTargetsRequestContent *SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent
}

// The identifier of a client associated with a &#39;Login with Amazon&#39; account. This is a required     header for advertisers and integrators using the Advertising API.
func (r ApiCreateTargetPromotionGroupTargetsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateTargetPromotionGroupTargetsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles     resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a     required header for advertisers and integrators using the Advertising API.
func (r ApiCreateTargetPromotionGroupTargetsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateTargetPromotionGroupTargetsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateTargetPromotionGroupTargetsRequest) SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent(sponsoredProductsCreateTargetPromotionGroupTargetsRequestContent SponsoredProductsCreateTargetPromotionGroupTargetsRequestContent) ApiCreateTargetPromotionGroupTargetsRequest {
	r.sponsoredProductsCreateTargetPromotionGroupTargetsRequestContent = &sponsoredProductsCreateTargetPromotionGroupTargetsRequestContent
	return r
}

func (r ApiCreateTargetPromotionGroupTargetsRequest) Execute() (*SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent, *http.Response, error) {
	return r.ApiService.CreateTargetPromotionGroupTargetsExecute(r)
}

func (a *TargetPromotionGroupsAPIService) CreateTargetPromotionGroupTargets(ctx context.Context) ApiCreateTargetPromotionGroupTargetsRequest {
	return ApiCreateTargetPromotionGroupTargetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent
func (a *TargetPromotionGroupsAPIService) CreateTargetPromotionGroupTargetsExecute(r ApiCreateTargetPromotionGroupTargetsRequest) (*SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetPromotionGroupsAPIService.CreateTargetPromotionGroupTargets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targetPromotionGroups/targets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateTargetPromotionGroupTargetsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateTargetPromotionGroupTargetsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetPromotionGroupTarget.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spTargetPromotionGroupTarget.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsCreateTargetPromotionGroupTargetsRequestContent
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

type ApiCreateTargetPromotionGroupsRequest struct {
	ctx                                                        context.Context
	ApiService                                                 *TargetPromotionGroupsAPIService
	amazonAdvertisingAPIClientId                               *string
	amazonAdvertisingAPIScope                                  *string
	sponsoredProductsCreateTargetPromotionGroupsRequestContent *SponsoredProductsCreateTargetPromotionGroupsRequestContent
}

// The identifier of a client associated with a &#39;Login with Amazon&#39; account. This is a required     header for advertisers and integrators using the Advertising API.
func (r ApiCreateTargetPromotionGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateTargetPromotionGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles     resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a     required header for advertisers and integrators using the Advertising API.
func (r ApiCreateTargetPromotionGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateTargetPromotionGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateTargetPromotionGroupsRequest) SponsoredProductsCreateTargetPromotionGroupsRequestContent(sponsoredProductsCreateTargetPromotionGroupsRequestContent SponsoredProductsCreateTargetPromotionGroupsRequestContent) ApiCreateTargetPromotionGroupsRequest {
	r.sponsoredProductsCreateTargetPromotionGroupsRequestContent = &sponsoredProductsCreateTargetPromotionGroupsRequestContent
	return r
}

func (r ApiCreateTargetPromotionGroupsRequest) Execute() (*SponsoredProductsCreateTargetPromotionGroupsResponseContent, *http.Response, error) {
	return r.ApiService.CreateTargetPromotionGroupsExecute(r)
}

func (a *TargetPromotionGroupsAPIService) CreateTargetPromotionGroups(ctx context.Context) ApiCreateTargetPromotionGroupsRequest {
	return ApiCreateTargetPromotionGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateTargetPromotionGroupsResponseContent
func (a *TargetPromotionGroupsAPIService) CreateTargetPromotionGroupsExecute(r ApiCreateTargetPromotionGroupsRequest) (*SponsoredProductsCreateTargetPromotionGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateTargetPromotionGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetPromotionGroupsAPIService.CreateTargetPromotionGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targetPromotionGroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateTargetPromotionGroupsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateTargetPromotionGroupsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetPromotionGroup.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spTargetPromotionGroup.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsCreateTargetPromotionGroupsRequestContent
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

type ApiGetTargetPromotionGroupsRecommendationsRequest struct {
	ctx                                                                    context.Context
	ApiService                                                             *TargetPromotionGroupsAPIService
	amazonAdvertisingAPIClientId                                           *string
	amazonAdvertisingAPIScope                                              *string
	sponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent *SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent
}

// The identifier of a client associated with a &#39;Login with Amazon&#39; account.
func (r ApiGetTargetPromotionGroupsRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetTargetPromotionGroupsRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetTargetPromotionGroupsRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetTargetPromotionGroupsRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetTargetPromotionGroupsRecommendationsRequest) SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent(sponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent SponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent) ApiGetTargetPromotionGroupsRecommendationsRequest {
	r.sponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent = &sponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent
	return r
}

func (r ApiGetTargetPromotionGroupsRecommendationsRequest) Execute() (*SponsoredProductsGetTargetPromotionGroupsRecommendationsResponseContent, *http.Response, error) {
	return r.ApiService.GetTargetPromotionGroupsRecommendationsExecute(r)
}

func (a *TargetPromotionGroupsAPIService) GetTargetPromotionGroupsRecommendations(ctx context.Context) ApiGetTargetPromotionGroupsRecommendationsRequest {
	return ApiGetTargetPromotionGroupsRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsGetTargetPromotionGroupsRecommendationsResponseContent
func (a *TargetPromotionGroupsAPIService) GetTargetPromotionGroupsRecommendationsExecute(r ApiGetTargetPromotionGroupsRecommendationsRequest) (*SponsoredProductsGetTargetPromotionGroupsRecommendationsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsGetTargetPromotionGroupsRecommendationsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetPromotionGroupsAPIService.GetTargetPromotionGroupsRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targetPromotionGroups/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetPromotionGroupsRecommendations.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spTargetPromotionGroupsRecommendations.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsGetTargetPromotionGroupsRecommendationsRequestContent
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

type ApiListTargetPromotionGroupTargetsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *TargetPromotionGroupsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsListTargetPromotionGroupTargetsRequestContent *SponsoredProductsListTargetPromotionGroupTargetsRequestContent
}

// The identifier of a client associated with a &#39;Login with Amazon&#39; account. This is a required     header for advertisers and integrators using the Advertising API.
func (r ApiListTargetPromotionGroupTargetsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListTargetPromotionGroupTargetsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles     resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a     required header for advertisers and integrators using the Advertising API.
func (r ApiListTargetPromotionGroupTargetsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListTargetPromotionGroupTargetsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListTargetPromotionGroupTargetsRequest) SponsoredProductsListTargetPromotionGroupTargetsRequestContent(sponsoredProductsListTargetPromotionGroupTargetsRequestContent SponsoredProductsListTargetPromotionGroupTargetsRequestContent) ApiListTargetPromotionGroupTargetsRequest {
	r.sponsoredProductsListTargetPromotionGroupTargetsRequestContent = &sponsoredProductsListTargetPromotionGroupTargetsRequestContent
	return r
}

func (r ApiListTargetPromotionGroupTargetsRequest) Execute() (*SponsoredProductsListTargetPromotionGroupTargetsResponseContent, *http.Response, error) {
	return r.ApiService.ListTargetPromotionGroupTargetsExecute(r)
}

func (a *TargetPromotionGroupsAPIService) ListTargetPromotionGroupTargets(ctx context.Context) ApiListTargetPromotionGroupTargetsRequest {
	return ApiListTargetPromotionGroupTargetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListTargetPromotionGroupTargetsResponseContent
func (a *TargetPromotionGroupsAPIService) ListTargetPromotionGroupTargetsExecute(r ApiListTargetPromotionGroupTargetsRequest) (*SponsoredProductsListTargetPromotionGroupTargetsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListTargetPromotionGroupTargetsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetPromotionGroupsAPIService.ListTargetPromotionGroupTargets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targetPromotionGroups/targets/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetPromotionGroupTarget.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spTargetPromotionGroupTarget.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListTargetPromotionGroupTargetsRequestContent
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

type ApiListTargetPromotionGroupsRequest struct {
	ctx                                                      context.Context
	ApiService                                               *TargetPromotionGroupsAPIService
	amazonAdvertisingAPIClientId                             *string
	amazonAdvertisingAPIScope                                *string
	sponsoredProductsListTargetPromotionGroupsRequestContent *SponsoredProductsListTargetPromotionGroupsRequestContent
}

// The identifier of a client associated with a &#39;Login with Amazon&#39; account. This is a required     header for advertisers and integrators using the Advertising API.
func (r ApiListTargetPromotionGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListTargetPromotionGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles     resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a     required header for advertisers and integrators using the Advertising API.
func (r ApiListTargetPromotionGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListTargetPromotionGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListTargetPromotionGroupsRequest) SponsoredProductsListTargetPromotionGroupsRequestContent(sponsoredProductsListTargetPromotionGroupsRequestContent SponsoredProductsListTargetPromotionGroupsRequestContent) ApiListTargetPromotionGroupsRequest {
	r.sponsoredProductsListTargetPromotionGroupsRequestContent = &sponsoredProductsListTargetPromotionGroupsRequestContent
	return r
}

func (r ApiListTargetPromotionGroupsRequest) Execute() (*SponsoredProductsListTargetPromotionGroupsResponseContent, *http.Response, error) {
	return r.ApiService.ListTargetPromotionGroupsExecute(r)
}

func (a *TargetPromotionGroupsAPIService) ListTargetPromotionGroups(ctx context.Context) ApiListTargetPromotionGroupsRequest {
	return ApiListTargetPromotionGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListTargetPromotionGroupsResponseContent
func (a *TargetPromotionGroupsAPIService) ListTargetPromotionGroupsExecute(r ApiListTargetPromotionGroupsRequest) (*SponsoredProductsListTargetPromotionGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListTargetPromotionGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetPromotionGroupsAPIService.ListTargetPromotionGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targetPromotionGroups/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetPromotionGroup.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spTargetPromotionGroup.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListTargetPromotionGroupsRequestContent
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
