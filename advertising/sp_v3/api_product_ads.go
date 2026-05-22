package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ProductAdsAPIService ProductAdsAPI service
type ProductAdsAPIService service

type ApiCreateSponsoredProductsProductAdsRequest struct {
	ctx                                                              context.Context
	ApiService                                                       *ProductAdsAPIService
	amazonAdvertisingAPIClientId                                     *string
	amazonAdvertisingAPIScope                                        *string
	sponsoredProductsCreateSponsoredProductsProductAdsRequestContent *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent
	prefer                                                           *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsProductAdsRequest) SponsoredProductsCreateSponsoredProductsProductAdsRequestContent(sponsoredProductsCreateSponsoredProductsProductAdsRequestContent SponsoredProductsCreateSponsoredProductsProductAdsRequestContent) ApiCreateSponsoredProductsProductAdsRequest {
	r.sponsoredProductsCreateSponsoredProductsProductAdsRequestContent = &sponsoredProductsCreateSponsoredProductsProductAdsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsProductAdsRequest) Prefer(prefer string) ApiCreateSponsoredProductsProductAdsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsProductAdsRequest) Execute() (*SponsoredProductsCreateSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsProductAdsExecute(r)
}

func (a *ProductAdsAPIService) CreateSponsoredProductsProductAds(ctx context.Context) ApiCreateSponsoredProductsProductAdsRequest {
	return ApiCreateSponsoredProductsProductAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsProductAdsResponseContent
func (a *ProductAdsAPIService) CreateSponsoredProductsProductAdsExecute(r ApiCreateSponsoredProductsProductAdsRequest) (*SponsoredProductsCreateSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.CreateSponsoredProductsProductAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/productAds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsProductAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsProductAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spProductAd.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spProductAd.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsProductAdsRequestContent
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

type ApiDeleteSponsoredProductsProductAdsRequest struct {
	ctx                                                              context.Context
	ApiService                                                       *ProductAdsAPIService
	amazonAdvertisingAPIClientId                                     *string
	amazonAdvertisingAPIScope                                        *string
	sponsoredProductsDeleteSponsoredProductsProductAdsRequestContent *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsProductAdsRequest) SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent(sponsoredProductsDeleteSponsoredProductsProductAdsRequestContent SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) ApiDeleteSponsoredProductsProductAdsRequest {
	r.sponsoredProductsDeleteSponsoredProductsProductAdsRequestContent = &sponsoredProductsDeleteSponsoredProductsProductAdsRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsProductAdsRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsProductAdsExecute(r)
}

func (a *ProductAdsAPIService) DeleteSponsoredProductsProductAds(ctx context.Context) ApiDeleteSponsoredProductsProductAdsRequest {
	return ApiDeleteSponsoredProductsProductAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent
func (a *ProductAdsAPIService) DeleteSponsoredProductsProductAdsExecute(r ApiDeleteSponsoredProductsProductAdsRequest) (*SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.DeleteSponsoredProductsProductAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/productAds/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsProductAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsProductAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spProductAd.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spProductAd.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsProductAdsRequestContent
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

type ApiListSponsoredProductsProductAdsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *ProductAdsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsListSponsoredProductsProductAdsRequestContent *SponsoredProductsListSponsoredProductsProductAdsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsProductAdsRequest) SponsoredProductsListSponsoredProductsProductAdsRequestContent(sponsoredProductsListSponsoredProductsProductAdsRequestContent SponsoredProductsListSponsoredProductsProductAdsRequestContent) ApiListSponsoredProductsProductAdsRequest {
	r.sponsoredProductsListSponsoredProductsProductAdsRequestContent = &sponsoredProductsListSponsoredProductsProductAdsRequestContent
	return r
}

func (r ApiListSponsoredProductsProductAdsRequest) Execute() (*SponsoredProductsListSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsProductAdsExecute(r)
}

func (a *ProductAdsAPIService) ListSponsoredProductsProductAds(ctx context.Context) ApiListSponsoredProductsProductAdsRequest {
	return ApiListSponsoredProductsProductAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsProductAdsResponseContent
func (a *ProductAdsAPIService) ListSponsoredProductsProductAdsExecute(r ApiListSponsoredProductsProductAdsRequest) (*SponsoredProductsListSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsProductAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.ListSponsoredProductsProductAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/productAds/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spProductAd.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spProductAd.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsProductAdsRequestContent
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

type ApiUpdateSponsoredProductsProductAdsRequest struct {
	ctx                                                              context.Context
	ApiService                                                       *ProductAdsAPIService
	amazonAdvertisingAPIClientId                                     *string
	amazonAdvertisingAPIScope                                        *string
	sponsoredProductsUpdateSponsoredProductsProductAdsRequestContent *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent
	prefer                                                           *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsProductAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsProductAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsProductAdsRequest) SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent(sponsoredProductsUpdateSponsoredProductsProductAdsRequestContent SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) ApiUpdateSponsoredProductsProductAdsRequest {
	r.sponsoredProductsUpdateSponsoredProductsProductAdsRequestContent = &sponsoredProductsUpdateSponsoredProductsProductAdsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsProductAdsRequest) Prefer(prefer string) ApiUpdateSponsoredProductsProductAdsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsProductAdsRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsProductAdsExecute(r)
}

func (a *ProductAdsAPIService) UpdateSponsoredProductsProductAds(ctx context.Context) ApiUpdateSponsoredProductsProductAdsRequest {
	return ApiUpdateSponsoredProductsProductAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent
func (a *ProductAdsAPIService) UpdateSponsoredProductsProductAdsExecute(r ApiUpdateSponsoredProductsProductAdsRequest) (*SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.UpdateSponsoredProductsProductAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/productAds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsProductAdsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsProductAdsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spProductAd.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spProductAd.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsProductAdsRequestContent
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
