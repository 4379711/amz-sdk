package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// NegativeKeywordsAPIService NegativeKeywordsAPI service
type NegativeKeywordsAPIService service

type ApiCreateSponsoredProductsNegativeKeywordsRequest struct {
	ctx                                                                    context.Context
	ApiService                                                             *NegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                           *string
	amazonAdvertisingAPIScope                                              *string
	sponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent
	prefer                                                                 *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsNegativeKeywordsRequest) SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent(sponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) ApiCreateSponsoredProductsNegativeKeywordsRequest {
	r.sponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent = &sponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsNegativeKeywordsRequest) Prefer(prefer string) ApiCreateSponsoredProductsNegativeKeywordsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsNegativeKeywordsRequest) Execute() (*SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsNegativeKeywordsExecute(r)
}

func (a *NegativeKeywordsAPIService) CreateSponsoredProductsNegativeKeywords(ctx context.Context) ApiCreateSponsoredProductsNegativeKeywordsRequest {
	return ApiCreateSponsoredProductsNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent
func (a *NegativeKeywordsAPIService) CreateSponsoredProductsNegativeKeywordsExecute(r ApiCreateSponsoredProductsNegativeKeywordsRequest) (*SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeKeywordsAPIService.CreateSponsoredProductsNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeKeywords"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spNegativeKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent
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

type ApiDeleteSponsoredProductsNegativeKeywordsRequest struct {
	ctx                                                                    context.Context
	ApiService                                                             *NegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                           *string
	amazonAdvertisingAPIScope                                              *string
	sponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsNegativeKeywordsRequest) SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent(sponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) ApiDeleteSponsoredProductsNegativeKeywordsRequest {
	r.sponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent = &sponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsNegativeKeywordsRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsNegativeKeywordsExecute(r)
}

func (a *NegativeKeywordsAPIService) DeleteSponsoredProductsNegativeKeywords(ctx context.Context) ApiDeleteSponsoredProductsNegativeKeywordsRequest {
	return ApiDeleteSponsoredProductsNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent
func (a *NegativeKeywordsAPIService) DeleteSponsoredProductsNegativeKeywordsExecute(r ApiDeleteSponsoredProductsNegativeKeywordsRequest) (*SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeKeywordsAPIService.DeleteSponsoredProductsNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeKeywords/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spNegativeKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent
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

type ApiListSponsoredProductsNegativeKeywordsRequest struct {
	ctx                                                                  context.Context
	ApiService                                                           *NegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                         *string
	amazonAdvertisingAPIScope                                            *string
	sponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent *SponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsNegativeKeywordsRequest) SponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent(sponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent SponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent) ApiListSponsoredProductsNegativeKeywordsRequest {
	r.sponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent = &sponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent
	return r
}

func (r ApiListSponsoredProductsNegativeKeywordsRequest) Execute() (*SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsNegativeKeywordsExecute(r)
}

func (a *NegativeKeywordsAPIService) ListSponsoredProductsNegativeKeywords(ctx context.Context) ApiListSponsoredProductsNegativeKeywordsRequest {
	return ApiListSponsoredProductsNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent
func (a *NegativeKeywordsAPIService) ListSponsoredProductsNegativeKeywordsExecute(r ApiListSponsoredProductsNegativeKeywordsRequest) (*SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeKeywordsAPIService.ListSponsoredProductsNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeKeywords/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spNegativeKeyword.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsNegativeKeywordsRequestContent
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

type ApiUpdateSponsoredProductsNegativeKeywordsRequest struct {
	ctx                                                                    context.Context
	ApiService                                                             *NegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                           *string
	amazonAdvertisingAPIScope                                              *string
	sponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent
	prefer                                                                 *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsNegativeKeywordsRequest) SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent(sponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) ApiUpdateSponsoredProductsNegativeKeywordsRequest {
	r.sponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent = &sponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsNegativeKeywordsRequest) Prefer(prefer string) ApiUpdateSponsoredProductsNegativeKeywordsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsNegativeKeywordsRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsNegativeKeywordsExecute(r)
}

func (a *NegativeKeywordsAPIService) UpdateSponsoredProductsNegativeKeywords(ctx context.Context) ApiUpdateSponsoredProductsNegativeKeywordsRequest {
	return ApiUpdateSponsoredProductsNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent
func (a *NegativeKeywordsAPIService) UpdateSponsoredProductsNegativeKeywordsExecute(r ApiUpdateSponsoredProductsNegativeKeywordsRequest) (*SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeKeywordsAPIService.UpdateSponsoredProductsNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeKeywords"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spNegativeKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent
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
