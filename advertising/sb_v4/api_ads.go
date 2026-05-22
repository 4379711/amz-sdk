package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// AdsAPIService AdsAPI service
type AdsAPIService service

type ApiCreateSponsoredBrandStoreSpotlightAdsRequest struct {
	ctx                                                 context.Context
	ApiService                                          *AdsAPIService
	amazonAdvertisingAPIClientId                        *string
	amazonAdvertisingAPIScope                           *string
	createSponsoredBrandStoreSpotlightAdsRequestContent *CreateSponsoredBrandStoreSpotlightAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandStoreSpotlightAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandStoreSpotlightAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandStoreSpotlightAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandStoreSpotlightAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandStoreSpotlightAdsRequest) CreateSponsoredBrandStoreSpotlightAdsRequestContent(createSponsoredBrandStoreSpotlightAdsRequestContent CreateSponsoredBrandStoreSpotlightAdsRequestContent) ApiCreateSponsoredBrandStoreSpotlightAdsRequest {
	r.createSponsoredBrandStoreSpotlightAdsRequestContent = &createSponsoredBrandStoreSpotlightAdsRequestContent
	return r
}

func (r ApiCreateSponsoredBrandStoreSpotlightAdsRequest) Execute() (*CreateSponsoredBrandStoreSpotlightAdsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandStoreSpotlightAdsExecute(r)
}

func (a *AdsAPIService) CreateSponsoredBrandStoreSpotlightAds(ctx context.Context) ApiCreateSponsoredBrandStoreSpotlightAdsRequest {
	return ApiCreateSponsoredBrandStoreSpotlightAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandStoreSpotlightAdsResponseContent
func (a *AdsAPIService) CreateSponsoredBrandStoreSpotlightAdsExecute(r ApiCreateSponsoredBrandStoreSpotlightAdsRequest) (*CreateSponsoredBrandStoreSpotlightAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandStoreSpotlightAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.CreateSponsoredBrandStoreSpotlightAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads/storeSpotlight"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandStoreSpotlightAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandStoreSpotlightAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandStoreSpotlightAdsRequestContent
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

type ApiCreateSponsoredBrandsBrandVideoAdsRequest struct {
	ctx                                              context.Context
	ApiService                                       *AdsAPIService
	amazonAdvertisingAPIClientId                     *string
	amazonAdvertisingAPIScope                        *string
	createSponsoredBrandsBrandVideoAdsRequestContent *CreateSponsoredBrandsBrandVideoAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandsBrandVideoAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandsBrandVideoAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandsBrandVideoAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandsBrandVideoAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandsBrandVideoAdsRequest) CreateSponsoredBrandsBrandVideoAdsRequestContent(createSponsoredBrandsBrandVideoAdsRequestContent CreateSponsoredBrandsBrandVideoAdsRequestContent) ApiCreateSponsoredBrandsBrandVideoAdsRequest {
	r.createSponsoredBrandsBrandVideoAdsRequestContent = &createSponsoredBrandsBrandVideoAdsRequestContent
	return r
}

func (r ApiCreateSponsoredBrandsBrandVideoAdsRequest) Execute() (*CreateSponsoredBrandsBrandVideoAdsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandsBrandVideoAdsExecute(r)
}

func (a *AdsAPIService) CreateSponsoredBrandsBrandVideoAds(ctx context.Context) ApiCreateSponsoredBrandsBrandVideoAdsRequest {
	return ApiCreateSponsoredBrandsBrandVideoAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandsBrandVideoAdsResponseContent
func (a *AdsAPIService) CreateSponsoredBrandsBrandVideoAdsExecute(r ApiCreateSponsoredBrandsBrandVideoAdsRequest) (*CreateSponsoredBrandsBrandVideoAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandsBrandVideoAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.CreateSponsoredBrandsBrandVideoAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads/brandVideo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandsBrandVideoAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandsBrandVideoAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandsBrandVideoAdsRequestContent
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

type ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest struct {
	ctx                                                             context.Context
	ApiService                                                      *AdsAPIService
	amazonAdvertisingAPIClientId                                    *string
	amazonAdvertisingAPIScope                                       *string
	createSponsoredBrandsExtendedProductCollectionAdsRequestContent *CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest) CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent(createSponsoredBrandsExtendedProductCollectionAdsRequestContent CreateSponsoredBrandsExtendedProductCollectionAdsRequestContent) ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest {
	r.createSponsoredBrandsExtendedProductCollectionAdsRequestContent = &createSponsoredBrandsExtendedProductCollectionAdsRequestContent
	return r
}

func (r ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest) Execute() (*CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandsExtendedProductCollectionAdsExecute(r)
}

func (a *AdsAPIService) CreateSponsoredBrandsExtendedProductCollectionAds(ctx context.Context) ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest {
	return ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent
func (a *AdsAPIService) CreateSponsoredBrandsExtendedProductCollectionAdsExecute(r ApiCreateSponsoredBrandsExtendedProductCollectionAdsRequest) (*CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandsExtendedProductCollectionAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.CreateSponsoredBrandsExtendedProductCollectionAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads/productCollectionExtended"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandsExtendedProductCollectionAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandsExtendedProductCollectionAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandsExtendedProductCollectionAdsRequestContent
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

type ApiCreateSponsoredBrandsProductCollectionAdsRequest struct {
	ctx                                                     context.Context
	ApiService                                              *AdsAPIService
	amazonAdvertisingAPIClientId                            *string
	amazonAdvertisingAPIScope                               *string
	createSponsoredBrandsProductCollectionAdsRequestContent *CreateSponsoredBrandsProductCollectionAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandsProductCollectionAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandsProductCollectionAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandsProductCollectionAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandsProductCollectionAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandsProductCollectionAdsRequest) CreateSponsoredBrandsProductCollectionAdsRequestContent(createSponsoredBrandsProductCollectionAdsRequestContent CreateSponsoredBrandsProductCollectionAdsRequestContent) ApiCreateSponsoredBrandsProductCollectionAdsRequest {
	r.createSponsoredBrandsProductCollectionAdsRequestContent = &createSponsoredBrandsProductCollectionAdsRequestContent
	return r
}

func (r ApiCreateSponsoredBrandsProductCollectionAdsRequest) Execute() (*CreateSponsoredBrandsProductCollectionAdsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandsProductCollectionAdsExecute(r)
}

func (a *AdsAPIService) CreateSponsoredBrandsProductCollectionAds(ctx context.Context) ApiCreateSponsoredBrandsProductCollectionAdsRequest {
	return ApiCreateSponsoredBrandsProductCollectionAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandsProductCollectionAdsResponseContent
func (a *AdsAPIService) CreateSponsoredBrandsProductCollectionAdsExecute(r ApiCreateSponsoredBrandsProductCollectionAdsRequest) (*CreateSponsoredBrandsProductCollectionAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandsProductCollectionAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.CreateSponsoredBrandsProductCollectionAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads/productCollection"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandsProductCollectionAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandsProductCollectionAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandsProductCollectionAdsRequestContent
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

type ApiCreateSponsoredBrandsVideoAdsRequest struct {
	ctx                                         context.Context
	ApiService                                  *AdsAPIService
	amazonAdvertisingAPIClientId                *string
	amazonAdvertisingAPIScope                   *string
	createSponsoredBrandsVideoAdsRequestContent *CreateSponsoredBrandsVideoAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandsVideoAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandsVideoAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandsVideoAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandsVideoAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandsVideoAdsRequest) CreateSponsoredBrandsVideoAdsRequestContent(createSponsoredBrandsVideoAdsRequestContent CreateSponsoredBrandsVideoAdsRequestContent) ApiCreateSponsoredBrandsVideoAdsRequest {
	r.createSponsoredBrandsVideoAdsRequestContent = &createSponsoredBrandsVideoAdsRequestContent
	return r
}

func (r ApiCreateSponsoredBrandsVideoAdsRequest) Execute() (*CreateSponsoredBrandsVideoAdsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandsVideoAdsExecute(r)
}

func (a *AdsAPIService) CreateSponsoredBrandsVideoAds(ctx context.Context) ApiCreateSponsoredBrandsVideoAdsRequest {
	return ApiCreateSponsoredBrandsVideoAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandsVideoAdsResponseContent
func (a *AdsAPIService) CreateSponsoredBrandsVideoAdsExecute(r ApiCreateSponsoredBrandsVideoAdsRequest) (*CreateSponsoredBrandsVideoAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandsVideoAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.CreateSponsoredBrandsVideoAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads/video"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandsVideoAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandsVideoAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandsVideoAdsRequestContent
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

type ApiDeleteSponsoredBrandsAdsRequest struct {
	ctx                                    context.Context
	ApiService                             *AdsAPIService
	amazonAdvertisingAPIClientId           *string
	amazonAdvertisingAPIScope              *string
	deleteSponsoredBrandsAdsRequestContent *DeleteSponsoredBrandsAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredBrandsAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredBrandsAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiDeleteSponsoredBrandsAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredBrandsAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredBrandsAdsRequest) DeleteSponsoredBrandsAdsRequestContent(deleteSponsoredBrandsAdsRequestContent DeleteSponsoredBrandsAdsRequestContent) ApiDeleteSponsoredBrandsAdsRequest {
	r.deleteSponsoredBrandsAdsRequestContent = &deleteSponsoredBrandsAdsRequestContent
	return r
}

func (r ApiDeleteSponsoredBrandsAdsRequest) Execute() (*DeleteSponsoredBrandsAdsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredBrandsAdsExecute(r)
}

func (a *AdsAPIService) DeleteSponsoredBrandsAds(ctx context.Context) ApiDeleteSponsoredBrandsAdsRequest {
	return ApiDeleteSponsoredBrandsAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeleteSponsoredBrandsAdsResponseContent
func (a *AdsAPIService) DeleteSponsoredBrandsAdsExecute(r ApiDeleteSponsoredBrandsAdsRequest) (*DeleteSponsoredBrandsAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteSponsoredBrandsAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.DeleteSponsoredBrandsAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.deleteSponsoredBrandsAdsRequestContent
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

type ApiListSponsoredBrandsAdsRequest struct {
	ctx                                  context.Context
	ApiService                           *AdsAPIService
	amazonAdvertisingAPIClientId         *string
	amazonAdvertisingAPIScope            *string
	listSponsoredBrandsAdsRequestContent *ListSponsoredBrandsAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredBrandsAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredBrandsAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiListSponsoredBrandsAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredBrandsAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredBrandsAdsRequest) ListSponsoredBrandsAdsRequestContent(listSponsoredBrandsAdsRequestContent ListSponsoredBrandsAdsRequestContent) ApiListSponsoredBrandsAdsRequest {
	r.listSponsoredBrandsAdsRequestContent = &listSponsoredBrandsAdsRequestContent
	return r
}

func (r ApiListSponsoredBrandsAdsRequest) Execute() (*ListSponsoredBrandsAdsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredBrandsAdsExecute(r)
}

func (a *AdsAPIService) ListSponsoredBrandsAds(ctx context.Context) ApiListSponsoredBrandsAdsRequest {
	return ApiListSponsoredBrandsAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListSponsoredBrandsAdsResponseContent
func (a *AdsAPIService) ListSponsoredBrandsAdsExecute(r ApiListSponsoredBrandsAdsRequest) (*ListSponsoredBrandsAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSponsoredBrandsAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.ListSponsoredBrandsAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.listSponsoredBrandsAdsRequestContent
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

type ApiUpdateSponsoredBrandsAdsRequest struct {
	ctx                                    context.Context
	ApiService                             *AdsAPIService
	amazonAdvertisingAPIClientId           *string
	amazonAdvertisingAPIScope              *string
	updateSponsoredBrandsAdsRequestContent *UpdateSponsoredBrandsAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredBrandsAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredBrandsAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiUpdateSponsoredBrandsAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredBrandsAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredBrandsAdsRequest) UpdateSponsoredBrandsAdsRequestContent(updateSponsoredBrandsAdsRequestContent UpdateSponsoredBrandsAdsRequestContent) ApiUpdateSponsoredBrandsAdsRequest {
	r.updateSponsoredBrandsAdsRequestContent = &updateSponsoredBrandsAdsRequestContent
	return r
}

func (r ApiUpdateSponsoredBrandsAdsRequest) Execute() (*UpdateSponsoredBrandsAdsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredBrandsAdsExecute(r)
}

func (a *AdsAPIService) UpdateSponsoredBrandsAds(ctx context.Context) ApiUpdateSponsoredBrandsAdsRequest {
	return ApiUpdateSponsoredBrandsAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateSponsoredBrandsAdsResponseContent
func (a *AdsAPIService) UpdateSponsoredBrandsAdsExecute(r ApiUpdateSponsoredBrandsAdsRequest) (*UpdateSponsoredBrandsAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSponsoredBrandsAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdsAPIService.UpdateSponsoredBrandsAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/ads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSponsoredBrandsAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("updateSponsoredBrandsAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.updateSponsoredBrandsAdsRequestContent
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
