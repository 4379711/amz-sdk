package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CampaignNegativeTargetingClausesAPIService CampaignNegativeTargetingClausesAPI service
type CampaignNegativeTargetingClausesAPIService service

type ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest struct {
	ctx                                                                                    context.Context
	ApiService                                                                             *CampaignNegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                           *string
	amazonAdvertisingAPIScope                                                              *string
	sponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	prefer                                                                                 *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest) SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent(sponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.sponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent = &sponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest) Prefer(prefer string) ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest) Execute() (*SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsCampaignNegativeTargetingClausesExecute(r)
}

func (a *CampaignNegativeTargetingClausesAPIService) CreateSponsoredProductsCampaignNegativeTargetingClauses(ctx context.Context) ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	return ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
func (a *CampaignNegativeTargetingClausesAPIService) CreateSponsoredProductsCampaignNegativeTargetingClausesExecute(r ApiCreateSponsoredProductsCampaignNegativeTargetingClausesRequest) (*SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeTargetingClausesAPIService.CreateSponsoredProductsCampaignNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeTargets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaignNegativeTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
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

type ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest struct {
	ctx                                                                                    context.Context
	ApiService                                                                             *CampaignNegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                           *string
	amazonAdvertisingAPIScope                                                              *string
	sponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest) SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent(sponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.sponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent = &sponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsCampaignNegativeTargetingClausesExecute(r)
}

func (a *CampaignNegativeTargetingClausesAPIService) DeleteSponsoredProductsCampaignNegativeTargetingClauses(ctx context.Context) ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest {
	return ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent
func (a *CampaignNegativeTargetingClausesAPIService) DeleteSponsoredProductsCampaignNegativeTargetingClausesExecute(r ApiDeleteSponsoredProductsCampaignNegativeTargetingClausesRequest) (*SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeTargetingClausesAPIService.DeleteSponsoredProductsCampaignNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeTargets/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaignNegativeTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsCampaignNegativeTargetingClausesRequestContent
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

type ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest struct {
	ctx                                                                                  context.Context
	ApiService                                                                           *CampaignNegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                         *string
	amazonAdvertisingAPIScope                                                            *string
	sponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest) SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent(sponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.sponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent = &sponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	return r
}

func (r ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest) Execute() (*SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsCampaignNegativeTargetingClausesExecute(r)
}

func (a *CampaignNegativeTargetingClausesAPIService) ListSponsoredProductsCampaignNegativeTargetingClauses(ctx context.Context) ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest {
	return ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent
func (a *CampaignNegativeTargetingClausesAPIService) ListSponsoredProductsCampaignNegativeTargetingClausesExecute(r ApiListSponsoredProductsCampaignNegativeTargetingClausesRequest) (*SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeTargetingClausesAPIService.ListSponsoredProductsCampaignNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeTargets/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spCampaignNegativeTargetingClause.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsCampaignNegativeTargetingClausesRequestContent
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

type ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest struct {
	ctx                                                                                    context.Context
	ApiService                                                                             *CampaignNegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                           *string
	amazonAdvertisingAPIScope                                                              *string
	sponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	prefer                                                                                 *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest) SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent(sponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent) ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.sponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent = &sponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest) Prefer(prefer string) ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsCampaignNegativeTargetingClausesExecute(r)
}

func (a *CampaignNegativeTargetingClausesAPIService) UpdateSponsoredProductsCampaignNegativeTargetingClauses(ctx context.Context) ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest {
	return ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
func (a *CampaignNegativeTargetingClausesAPIService) UpdateSponsoredProductsCampaignNegativeTargetingClausesExecute(r ApiUpdateSponsoredProductsCampaignNegativeTargetingClausesRequest) (*SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignNegativeTargetingClausesAPIService.UpdateSponsoredProductsCampaignNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaignNegativeTargets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaignNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaignNegativeTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsCampaignNegativeTargetingClausesRequestContent
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
