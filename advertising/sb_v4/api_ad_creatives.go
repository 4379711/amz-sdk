package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// AdCreativesAPIService AdCreativesAPI service
type AdCreativesAPIService service

type ApiCreateBrandVideoCreativeRequest struct {
	ctx                                    context.Context
	ApiService                             *AdCreativesAPIService
	amazonAdvertisingAPIClientId           *string
	amazonAdvertisingAPIScope              *string
	createBrandVideoCreativeRequestContent *CreateBrandVideoCreativeRequestContent
	accept                                 *AcceptHeader
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBrandVideoCreativeRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateBrandVideoCreativeRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateBrandVideoCreativeRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateBrandVideoCreativeRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateBrandVideoCreativeRequest) CreateBrandVideoCreativeRequestContent(createBrandVideoCreativeRequestContent CreateBrandVideoCreativeRequestContent) ApiCreateBrandVideoCreativeRequest {
	r.createBrandVideoCreativeRequestContent = &createBrandVideoCreativeRequestContent
	return r
}

// Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type.
func (r ApiCreateBrandVideoCreativeRequest) Accept(accept AcceptHeader) ApiCreateBrandVideoCreativeRequest {
	r.accept = &accept
	return r
}

func (r ApiCreateBrandVideoCreativeRequest) Execute() (*CreateBrandVideoCreativeResponseContent, *http.Response, error) {
	return r.ApiService.CreateBrandVideoCreativeExecute(r)
}

func (a *AdCreativesAPIService) CreateBrandVideoCreative(ctx context.Context) ApiCreateBrandVideoCreativeRequest {
	return ApiCreateBrandVideoCreativeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateBrandVideoCreativeResponseContent
func (a *AdCreativesAPIService) CreateBrandVideoCreativeExecute(r ApiCreateBrandVideoCreativeRequest) (*CreateBrandVideoCreativeResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateBrandVideoCreativeResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdCreativesAPIService.CreateBrandVideoCreative")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/ads/creatives/brandVideo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createBrandVideoCreativeRequestContent == nil {
		return localVarReturnValue, nil, reportError("createBrandVideoCreativeRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}

	// body params
	localVarPostBody = r.createBrandVideoCreativeRequestContent
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

type ApiCreateExtendedProductCollectionCreativeRequest struct {
	ctx                                                   context.Context
	ApiService                                            *AdCreativesAPIService
	amazonAdvertisingAPIClientId                          *string
	amazonAdvertisingAPIScope                             *string
	createExtendedProductCollectionCreativeRequestContent *CreateExtendedProductCollectionCreativeRequestContent
	accept                                                *AcceptHeader
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateExtendedProductCollectionCreativeRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateExtendedProductCollectionCreativeRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateExtendedProductCollectionCreativeRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateExtendedProductCollectionCreativeRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateExtendedProductCollectionCreativeRequest) CreateExtendedProductCollectionCreativeRequestContent(createExtendedProductCollectionCreativeRequestContent CreateExtendedProductCollectionCreativeRequestContent) ApiCreateExtendedProductCollectionCreativeRequest {
	r.createExtendedProductCollectionCreativeRequestContent = &createExtendedProductCollectionCreativeRequestContent
	return r
}

// Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type.
func (r ApiCreateExtendedProductCollectionCreativeRequest) Accept(accept AcceptHeader) ApiCreateExtendedProductCollectionCreativeRequest {
	r.accept = &accept
	return r
}

func (r ApiCreateExtendedProductCollectionCreativeRequest) Execute() (*CreateExtendedProductCollectionCreativeResponseContent, *http.Response, error) {
	return r.ApiService.CreateExtendedProductCollectionCreativeExecute(r)
}

func (a *AdCreativesAPIService) CreateExtendedProductCollectionCreative(ctx context.Context) ApiCreateExtendedProductCollectionCreativeRequest {
	return ApiCreateExtendedProductCollectionCreativeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateExtendedProductCollectionCreativeResponseContent
func (a *AdCreativesAPIService) CreateExtendedProductCollectionCreativeExecute(r ApiCreateExtendedProductCollectionCreativeRequest) (*CreateExtendedProductCollectionCreativeResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateExtendedProductCollectionCreativeResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdCreativesAPIService.CreateExtendedProductCollectionCreative")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/ads/creatives/productCollectionExtended"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createExtendedProductCollectionCreativeRequestContent == nil {
		return localVarReturnValue, nil, reportError("createExtendedProductCollectionCreativeRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}

	// body params
	localVarPostBody = r.createExtendedProductCollectionCreativeRequestContent
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

type ApiCreateProductCollectionCreativeRequest struct {
	ctx                                           context.Context
	ApiService                                    *AdCreativesAPIService
	amazonAdvertisingAPIClientId                  *string
	amazonAdvertisingAPIScope                     *string
	createProductCollectionCreativeRequestContent *CreateProductCollectionCreativeRequestContent
	accept                                        *AcceptHeader
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateProductCollectionCreativeRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateProductCollectionCreativeRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateProductCollectionCreativeRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateProductCollectionCreativeRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateProductCollectionCreativeRequest) CreateProductCollectionCreativeRequestContent(createProductCollectionCreativeRequestContent CreateProductCollectionCreativeRequestContent) ApiCreateProductCollectionCreativeRequest {
	r.createProductCollectionCreativeRequestContent = &createProductCollectionCreativeRequestContent
	return r
}

// Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type.
func (r ApiCreateProductCollectionCreativeRequest) Accept(accept AcceptHeader) ApiCreateProductCollectionCreativeRequest {
	r.accept = &accept
	return r
}

func (r ApiCreateProductCollectionCreativeRequest) Execute() (*CreateProductCollectionCreativeResponseContent, *http.Response, error) {
	return r.ApiService.CreateProductCollectionCreativeExecute(r)
}

func (a *AdCreativesAPIService) CreateProductCollectionCreative(ctx context.Context) ApiCreateProductCollectionCreativeRequest {
	return ApiCreateProductCollectionCreativeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateProductCollectionCreativeResponseContent
func (a *AdCreativesAPIService) CreateProductCollectionCreativeExecute(r ApiCreateProductCollectionCreativeRequest) (*CreateProductCollectionCreativeResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateProductCollectionCreativeResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdCreativesAPIService.CreateProductCollectionCreative")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/ads/creatives/productCollection"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createProductCollectionCreativeRequestContent == nil {
		return localVarReturnValue, nil, reportError("createProductCollectionCreativeRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}

	// body params
	localVarPostBody = r.createProductCollectionCreativeRequestContent
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

type ApiCreateStoreSpotlightCreativeRequest struct {
	ctx                                        context.Context
	ApiService                                 *AdCreativesAPIService
	amazonAdvertisingAPIClientId               *string
	amazonAdvertisingAPIScope                  *string
	createStoreSpotlightCreativeRequestContent *CreateStoreSpotlightCreativeRequestContent
	accept                                     *AcceptHeader
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateStoreSpotlightCreativeRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateStoreSpotlightCreativeRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateStoreSpotlightCreativeRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateStoreSpotlightCreativeRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateStoreSpotlightCreativeRequest) CreateStoreSpotlightCreativeRequestContent(createStoreSpotlightCreativeRequestContent CreateStoreSpotlightCreativeRequestContent) ApiCreateStoreSpotlightCreativeRequest {
	r.createStoreSpotlightCreativeRequestContent = &createStoreSpotlightCreativeRequestContent
	return r
}

// Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type.
func (r ApiCreateStoreSpotlightCreativeRequest) Accept(accept AcceptHeader) ApiCreateStoreSpotlightCreativeRequest {
	r.accept = &accept
	return r
}

func (r ApiCreateStoreSpotlightCreativeRequest) Execute() (*CreateStoreSpotlightCreativeResponseContent, *http.Response, error) {
	return r.ApiService.CreateStoreSpotlightCreativeExecute(r)
}

func (a *AdCreativesAPIService) CreateStoreSpotlightCreative(ctx context.Context) ApiCreateStoreSpotlightCreativeRequest {
	return ApiCreateStoreSpotlightCreativeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateStoreSpotlightCreativeResponseContent
func (a *AdCreativesAPIService) CreateStoreSpotlightCreativeExecute(r ApiCreateStoreSpotlightCreativeRequest) (*CreateStoreSpotlightCreativeResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateStoreSpotlightCreativeResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdCreativesAPIService.CreateStoreSpotlightCreative")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/ads/creatives/storeSpotlight"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createStoreSpotlightCreativeRequestContent == nil {
		return localVarReturnValue, nil, reportError("createStoreSpotlightCreativeRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}

	// body params
	localVarPostBody = r.createStoreSpotlightCreativeRequestContent
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

type ApiCreateVideoCreativeRequest struct {
	ctx                               context.Context
	ApiService                        *AdCreativesAPIService
	amazonAdvertisingAPIClientId      *string
	amazonAdvertisingAPIScope         *string
	createVideoCreativeRequestContent *CreateVideoCreativeRequestContent
	accept                            *AcceptHeader
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateVideoCreativeRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateVideoCreativeRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiCreateVideoCreativeRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateVideoCreativeRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateVideoCreativeRequest) CreateVideoCreativeRequestContent(createVideoCreativeRequestContent CreateVideoCreativeRequestContent) ApiCreateVideoCreativeRequest {
	r.createVideoCreativeRequestContent = &createVideoCreativeRequestContent
	return r
}

// Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type.
func (r ApiCreateVideoCreativeRequest) Accept(accept AcceptHeader) ApiCreateVideoCreativeRequest {
	r.accept = &accept
	return r
}

func (r ApiCreateVideoCreativeRequest) Execute() (*CreateVideoCreativeResponseContent, *http.Response, error) {
	return r.ApiService.CreateVideoCreativeExecute(r)
}

func (a *AdCreativesAPIService) CreateVideoCreative(ctx context.Context) ApiCreateVideoCreativeRequest {
	return ApiCreateVideoCreativeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateVideoCreativeResponseContent
func (a *AdCreativesAPIService) CreateVideoCreativeExecute(r ApiCreateVideoCreativeRequest) (*CreateVideoCreativeResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateVideoCreativeResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdCreativesAPIService.CreateVideoCreative")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/ads/creatives/video"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createVideoCreativeRequestContent == nil {
		return localVarReturnValue, nil, reportError("createVideoCreativeRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}

	// body params
	localVarPostBody = r.createVideoCreativeRequestContent
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

type ApiListCreativesRequest struct {
	ctx                          context.Context
	ApiService                   *AdCreativesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	listCreativesRequestContent  *ListCreativesRequestContent
	accept                       *AcceptHeader
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiListCreativesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListCreativesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and use profileId from the response to pass as input. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiListCreativesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListCreativesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListCreativesRequest) ListCreativesRequestContent(listCreativesRequestContent ListCreativesRequestContent) ApiListCreativesRequest {
	r.listCreativesRequestContent = &listCreativesRequestContent
	return r
}

// Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type.
func (r ApiListCreativesRequest) Accept(accept AcceptHeader) ApiListCreativesRequest {
	r.accept = &accept
	return r
}

func (r ApiListCreativesRequest) Execute() (*ListCreativesResponseContent, *http.Response, error) {
	return r.ApiService.ListCreativesExecute(r)
}

func (a *AdCreativesAPIService) ListCreatives(ctx context.Context) ApiListCreativesRequest {
	return ApiListCreativesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListCreativesResponseContent
func (a *AdCreativesAPIService) ListCreativesExecute(r ApiListCreativesRequest) (*ListCreativesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListCreativesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdCreativesAPIService.ListCreatives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/ads/creatives/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.listCreativesRequestContent == nil {
		return localVarReturnValue, nil, reportError("listCreativesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbAdCreativeResource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}

	// body params
	localVarPostBody = r.listCreativesRequestContent
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
