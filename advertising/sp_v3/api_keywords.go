package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// KeywordsAPIService KeywordsAPI service
type KeywordsAPIService service

type ApiCreateSponsoredProductsKeywordsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *KeywordsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsCreateSponsoredProductsKeywordsRequestContent *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent
	prefer                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsKeywordsRequest) SponsoredProductsCreateSponsoredProductsKeywordsRequestContent(sponsoredProductsCreateSponsoredProductsKeywordsRequestContent SponsoredProductsCreateSponsoredProductsKeywordsRequestContent) ApiCreateSponsoredProductsKeywordsRequest {
	r.sponsoredProductsCreateSponsoredProductsKeywordsRequestContent = &sponsoredProductsCreateSponsoredProductsKeywordsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsKeywordsRequest) Prefer(prefer string) ApiCreateSponsoredProductsKeywordsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsKeywordsRequest) Execute() (*SponsoredProductsCreateSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsKeywordsExecute(r)
}

func (a *KeywordsAPIService) CreateSponsoredProductsKeywords(ctx context.Context) ApiCreateSponsoredProductsKeywordsRequest {
	return ApiCreateSponsoredProductsKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsKeywordsResponseContent
func (a *KeywordsAPIService) CreateSponsoredProductsKeywordsExecute(r ApiCreateSponsoredProductsKeywordsRequest) (*SponsoredProductsCreateSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordsAPIService.CreateSponsoredProductsKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/keywords"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsKeywordsRequestContent
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

type ApiDeleteSponsoredProductsKeywordsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *KeywordsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsDeleteSponsoredProductsKeywordsRequestContent *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsKeywordsRequest) SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent(sponsoredProductsDeleteSponsoredProductsKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) ApiDeleteSponsoredProductsKeywordsRequest {
	r.sponsoredProductsDeleteSponsoredProductsKeywordsRequestContent = &sponsoredProductsDeleteSponsoredProductsKeywordsRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsKeywordsRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsKeywordsExecute(r)
}

func (a *KeywordsAPIService) DeleteSponsoredProductsKeywords(ctx context.Context) ApiDeleteSponsoredProductsKeywordsRequest {
	return ApiDeleteSponsoredProductsKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent
func (a *KeywordsAPIService) DeleteSponsoredProductsKeywordsExecute(r ApiDeleteSponsoredProductsKeywordsRequest) (*SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordsAPIService.DeleteSponsoredProductsKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/keywords/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsKeywordsRequestContent
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

type ApiListSponsoredProductsKeywordsRequest struct {
	ctx                                                          context.Context
	ApiService                                                   *KeywordsAPIService
	amazonAdvertisingAPIClientId                                 *string
	amazonAdvertisingAPIScope                                    *string
	sponsoredProductsListSponsoredProductsKeywordsRequestContent *SponsoredProductsListSponsoredProductsKeywordsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsKeywordsRequest) SponsoredProductsListSponsoredProductsKeywordsRequestContent(sponsoredProductsListSponsoredProductsKeywordsRequestContent SponsoredProductsListSponsoredProductsKeywordsRequestContent) ApiListSponsoredProductsKeywordsRequest {
	r.sponsoredProductsListSponsoredProductsKeywordsRequestContent = &sponsoredProductsListSponsoredProductsKeywordsRequestContent
	return r
}

func (r ApiListSponsoredProductsKeywordsRequest) Execute() (*SponsoredProductsListSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsKeywordsExecute(r)
}

func (a *KeywordsAPIService) ListSponsoredProductsKeywords(ctx context.Context) ApiListSponsoredProductsKeywordsRequest {
	return ApiListSponsoredProductsKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsKeywordsResponseContent
func (a *KeywordsAPIService) ListSponsoredProductsKeywordsExecute(r ApiListSponsoredProductsKeywordsRequest) (*SponsoredProductsListSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordsAPIService.ListSponsoredProductsKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/keywords/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spKeyword.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsKeywordsRequestContent
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

type ApiUpdateSponsoredProductsKeywordsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *KeywordsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsUpdateSponsoredProductsKeywordsRequestContent *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent
	prefer                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsKeywordsRequest) SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent(sponsoredProductsUpdateSponsoredProductsKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) ApiUpdateSponsoredProductsKeywordsRequest {
	r.sponsoredProductsUpdateSponsoredProductsKeywordsRequestContent = &sponsoredProductsUpdateSponsoredProductsKeywordsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsKeywordsRequest) Prefer(prefer string) ApiUpdateSponsoredProductsKeywordsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsKeywordsRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsKeywordsExecute(r)
}

func (a *KeywordsAPIService) UpdateSponsoredProductsKeywords(ctx context.Context) ApiUpdateSponsoredProductsKeywordsRequest {
	return ApiUpdateSponsoredProductsKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent
func (a *KeywordsAPIService) UpdateSponsoredProductsKeywordsExecute(r ApiUpdateSponsoredProductsKeywordsRequest) (*SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordsAPIService.UpdateSponsoredProductsKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/keywords"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsKeywordsRequestContent
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
