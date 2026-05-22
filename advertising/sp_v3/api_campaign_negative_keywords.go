package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CampaignNegativeKeywordsAPIService CampaignNegativeKeywordsAPI service
type CampaignNegativeKeywordsAPIService service

type ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest struct {
	ctx                                                                            context.Context
	ApiService                                                                     *CampaignNegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                                   *string
	amazonAdvertisingAPIScope                                                      *string
	sponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent
	prefer                                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest) SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent(sponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent) ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.sponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent = &sponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest) Prefer(prefer string) ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest) Execute() (*SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsCampaignNegativeKeywordsExecute(r)
}

func (a *CampaignNegativeKeywordsAPIService) CreateSponsoredProductsCampaignNegativeKeywords(ctx context.Context) ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest {
	return ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent
func (a *CampaignNegativeKeywordsAPIService) CreateSponsoredProductsCampaignNegativeKeywordsExecute(r ApiCreateSponsoredProductsCampaignNegativeKeywordsRequest) (*SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeKeywordsAPIService.CreateSponsoredProductsCampaignNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeKeywords"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaignNegativeKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsCampaignNegativeKeywordsRequestContent
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

type ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest struct {
	ctx                                                                            context.Context
	ApiService                                                                     *CampaignNegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                                   *string
	amazonAdvertisingAPIScope                                                      *string
	sponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest) SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent(sponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent) ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest {
	r.sponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent = &sponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsCampaignNegativeKeywordsExecute(r)
}

func (a *CampaignNegativeKeywordsAPIService) DeleteSponsoredProductsCampaignNegativeKeywords(ctx context.Context) ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest {
	return ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent
func (a *CampaignNegativeKeywordsAPIService) DeleteSponsoredProductsCampaignNegativeKeywordsExecute(r ApiDeleteSponsoredProductsCampaignNegativeKeywordsRequest) (*SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeKeywordsAPIService.DeleteSponsoredProductsCampaignNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeKeywords/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaignNegativeKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsCampaignNegativeKeywordsRequestContent
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

type ApiListSponsoredProductsCampaignNegativeKeywordsRequest struct {
	ctx                                                                          context.Context
	ApiService                                                                   *CampaignNegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                                 *string
	amazonAdvertisingAPIScope                                                    *string
	sponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsCampaignNegativeKeywordsRequest) SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent(sponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent) ApiListSponsoredProductsCampaignNegativeKeywordsRequest {
	r.sponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent = &sponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent
	return r
}

func (r ApiListSponsoredProductsCampaignNegativeKeywordsRequest) Execute() (*SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsCampaignNegativeKeywordsExecute(r)
}

func (a *CampaignNegativeKeywordsAPIService) ListSponsoredProductsCampaignNegativeKeywords(ctx context.Context) ApiListSponsoredProductsCampaignNegativeKeywordsRequest {
	return ApiListSponsoredProductsCampaignNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent
func (a *CampaignNegativeKeywordsAPIService) ListSponsoredProductsCampaignNegativeKeywordsExecute(r ApiListSponsoredProductsCampaignNegativeKeywordsRequest) (*SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsCampaignNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeKeywordsAPIService.ListSponsoredProductsCampaignNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeKeywords/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spCampaignNegativeKeyword.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsCampaignNegativeKeywordsRequestContent
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

type ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest struct {
	ctx                                                                            context.Context
	ApiService                                                                     *CampaignNegativeKeywordsAPIService
	amazonAdvertisingAPIClientId                                                   *string
	amazonAdvertisingAPIScope                                                      *string
	sponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent
	prefer                                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest) SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent(sponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent) ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.sponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent = &sponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest) Prefer(prefer string) ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsCampaignNegativeKeywordsExecute(r)
}

func (a *CampaignNegativeKeywordsAPIService) UpdateSponsoredProductsCampaignNegativeKeywords(ctx context.Context) ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest {
	return ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent
func (a *CampaignNegativeKeywordsAPIService) UpdateSponsoredProductsCampaignNegativeKeywordsExecute(r ApiUpdateSponsoredProductsCampaignNegativeKeywordsRequest) (*SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeKeywordsAPIService.UpdateSponsoredProductsCampaignNegativeKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeKeywords"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeKeyword.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaignNegativeKeyword.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsCampaignNegativeKeywordsRequestContent
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
