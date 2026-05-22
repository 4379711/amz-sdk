package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CampaignsAPIService CampaignsAPI service
type CampaignsAPIService service

type ApiCreateSponsoredProductsCampaignsRequest struct {
	ctx                                                             context.Context
	ApiService                                                      *CampaignsAPIService
	amazonAdvertisingAPIClientId                                    *string
	amazonAdvertisingAPIScope                                       *string
	sponsoredProductsCreateSponsoredProductsCampaignsRequestContent *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent
	prefer                                                          *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsCampaignsRequest) SponsoredProductsCreateSponsoredProductsCampaignsRequestContent(sponsoredProductsCreateSponsoredProductsCampaignsRequestContent SponsoredProductsCreateSponsoredProductsCampaignsRequestContent) ApiCreateSponsoredProductsCampaignsRequest {
	r.sponsoredProductsCreateSponsoredProductsCampaignsRequestContent = &sponsoredProductsCreateSponsoredProductsCampaignsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsCampaignsRequest) Prefer(prefer string) ApiCreateSponsoredProductsCampaignsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsCampaignsRequest) Execute() (*SponsoredProductsCreateSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsCampaignsExecute(r)
}

func (a *CampaignsAPIService) CreateSponsoredProductsCampaigns(ctx context.Context) ApiCreateSponsoredProductsCampaignsRequest {
	return ApiCreateSponsoredProductsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsCampaignsResponseContent
func (a *CampaignsAPIService) CreateSponsoredProductsCampaignsExecute(r ApiCreateSponsoredProductsCampaignsRequest) (*SponsoredProductsCreateSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.CreateSponsoredProductsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsCampaignsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsCampaignsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaign.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaign.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsCampaignsRequestContent
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

type ApiDeleteSponsoredProductsCampaignsRequest struct {
	ctx                                                             context.Context
	ApiService                                                      *CampaignsAPIService
	amazonAdvertisingAPIClientId                                    *string
	amazonAdvertisingAPIScope                                       *string
	sponsoredProductsDeleteSponsoredProductsCampaignsRequestContent *SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsCampaignsRequest) SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent(sponsoredProductsDeleteSponsoredProductsCampaignsRequestContent SponsoredProductsDeleteSponsoredProductsCampaignsRequestContent) ApiDeleteSponsoredProductsCampaignsRequest {
	r.sponsoredProductsDeleteSponsoredProductsCampaignsRequestContent = &sponsoredProductsDeleteSponsoredProductsCampaignsRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsCampaignsRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsCampaignsExecute(r)
}

func (a *CampaignsAPIService) DeleteSponsoredProductsCampaigns(ctx context.Context) ApiDeleteSponsoredProductsCampaignsRequest {
	return ApiDeleteSponsoredProductsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent
func (a *CampaignsAPIService) DeleteSponsoredProductsCampaignsExecute(r ApiDeleteSponsoredProductsCampaignsRequest) (*SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.DeleteSponsoredProductsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsCampaignsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsCampaignsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaign.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaign.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsCampaignsRequestContent
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

type ApiListSponsoredProductsCampaignsRequest struct {
	ctx                                                           context.Context
	ApiService                                                    *CampaignsAPIService
	amazonAdvertisingAPIClientId                                  *string
	amazonAdvertisingAPIScope                                     *string
	sponsoredProductsListSponsoredProductsCampaignsRequestContent *SponsoredProductsListSponsoredProductsCampaignsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsCampaignsRequest) SponsoredProductsListSponsoredProductsCampaignsRequestContent(sponsoredProductsListSponsoredProductsCampaignsRequestContent SponsoredProductsListSponsoredProductsCampaignsRequestContent) ApiListSponsoredProductsCampaignsRequest {
	r.sponsoredProductsListSponsoredProductsCampaignsRequestContent = &sponsoredProductsListSponsoredProductsCampaignsRequestContent
	return r
}

func (r ApiListSponsoredProductsCampaignsRequest) Execute() (*SponsoredProductsListSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsCampaignsExecute(r)
}

func (a *CampaignsAPIService) ListSponsoredProductsCampaigns(ctx context.Context) ApiListSponsoredProductsCampaignsRequest {
	return ApiListSponsoredProductsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsCampaignsResponseContent
func (a *CampaignsAPIService) ListSponsoredProductsCampaignsExecute(r ApiListSponsoredProductsCampaignsRequest) (*SponsoredProductsListSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListSponsoredProductsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaign.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spCampaign.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsCampaignsRequestContent
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

type ApiUpdateSponsoredProductsCampaignsRequest struct {
	ctx                                                             context.Context
	ApiService                                                      *CampaignsAPIService
	amazonAdvertisingAPIClientId                                    *string
	amazonAdvertisingAPIScope                                       *string
	sponsoredProductsUpdateSponsoredProductsCampaignsRequestContent *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent
	prefer                                                          *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsCampaignsRequest) SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent(sponsoredProductsUpdateSponsoredProductsCampaignsRequestContent SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) ApiUpdateSponsoredProductsCampaignsRequest {
	r.sponsoredProductsUpdateSponsoredProductsCampaignsRequestContent = &sponsoredProductsUpdateSponsoredProductsCampaignsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsCampaignsRequest) Prefer(prefer string) ApiUpdateSponsoredProductsCampaignsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsCampaignsRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsCampaignsExecute(r)
}

func (a *CampaignsAPIService) UpdateSponsoredProductsCampaigns(ctx context.Context) ApiUpdateSponsoredProductsCampaignsRequest {
	return ApiUpdateSponsoredProductsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent
func (a *CampaignsAPIService) UpdateSponsoredProductsCampaignsExecute(r ApiUpdateSponsoredProductsCampaignsRequest) (*SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.UpdateSponsoredProductsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsCampaignsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsCampaignsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spCampaign.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spCampaign.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsCampaignsRequestContent
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
