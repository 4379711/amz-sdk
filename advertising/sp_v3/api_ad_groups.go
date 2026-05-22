package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// AdGroupsAPIService AdGroupsAPI service
type AdGroupsAPIService service

type ApiCreateSponsoredProductsAdGroupsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *AdGroupsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsCreateSponsoredProductsAdGroupsRequestContent *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent
	prefer                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsAdGroupsRequest) SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent(sponsoredProductsCreateSponsoredProductsAdGroupsRequestContent SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) ApiCreateSponsoredProductsAdGroupsRequest {
	r.sponsoredProductsCreateSponsoredProductsAdGroupsRequestContent = &sponsoredProductsCreateSponsoredProductsAdGroupsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsAdGroupsRequest) Prefer(prefer string) ApiCreateSponsoredProductsAdGroupsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsAdGroupsRequest) Execute() (*SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) CreateSponsoredProductsAdGroups(ctx context.Context) ApiCreateSponsoredProductsAdGroupsRequest {
	return ApiCreateSponsoredProductsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent
func (a *AdGroupsAPIService) CreateSponsoredProductsAdGroupsExecute(r ApiCreateSponsoredProductsAdGroupsRequest) (*SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.CreateSponsoredProductsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/adGroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsAdGroupsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsAdGroupsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spAdGroup.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spAdGroup.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsAdGroupsRequestContent
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

type ApiDeleteSponsoredProductsAdGroupsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *AdGroupsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsAdGroupsRequest) SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent(sponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) ApiDeleteSponsoredProductsAdGroupsRequest {
	r.sponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent = &sponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsAdGroupsRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) DeleteSponsoredProductsAdGroups(ctx context.Context) ApiDeleteSponsoredProductsAdGroupsRequest {
	return ApiDeleteSponsoredProductsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent
func (a *AdGroupsAPIService) DeleteSponsoredProductsAdGroupsExecute(r ApiDeleteSponsoredProductsAdGroupsRequest) (*SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.DeleteSponsoredProductsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/adGroups/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spAdGroup.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spAdGroup.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent
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

type ApiListSponsoredProductsAdGroupsRequest struct {
	ctx                                                          context.Context
	ApiService                                                   *AdGroupsAPIService
	amazonAdvertisingAPIClientId                                 *string
	amazonAdvertisingAPIScope                                    *string
	sponsoredProductsListSponsoredProductsAdGroupsRequestContent *SponsoredProductsListSponsoredProductsAdGroupsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsAdGroupsRequest) SponsoredProductsListSponsoredProductsAdGroupsRequestContent(sponsoredProductsListSponsoredProductsAdGroupsRequestContent SponsoredProductsListSponsoredProductsAdGroupsRequestContent) ApiListSponsoredProductsAdGroupsRequest {
	r.sponsoredProductsListSponsoredProductsAdGroupsRequestContent = &sponsoredProductsListSponsoredProductsAdGroupsRequestContent
	return r
}

func (r ApiListSponsoredProductsAdGroupsRequest) Execute() (*SponsoredProductsListSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) ListSponsoredProductsAdGroups(ctx context.Context) ApiListSponsoredProductsAdGroupsRequest {
	return ApiListSponsoredProductsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsAdGroupsResponseContent
func (a *AdGroupsAPIService) ListSponsoredProductsAdGroupsExecute(r ApiListSponsoredProductsAdGroupsRequest) (*SponsoredProductsListSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.ListSponsoredProductsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/adGroups/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spAdGroup.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spAdGroup.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsAdGroupsRequestContent
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

type ApiUpdateSponsoredProductsAdGroupsRequest struct {
	ctx                                                            context.Context
	ApiService                                                     *AdGroupsAPIService
	amazonAdvertisingAPIClientId                                   *string
	amazonAdvertisingAPIScope                                      *string
	sponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent
	prefer                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsAdGroupsRequest) SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent(sponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) ApiUpdateSponsoredProductsAdGroupsRequest {
	r.sponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent = &sponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsAdGroupsRequest) Prefer(prefer string) ApiUpdateSponsoredProductsAdGroupsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsAdGroupsRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) UpdateSponsoredProductsAdGroups(ctx context.Context) ApiUpdateSponsoredProductsAdGroupsRequest {
	return ApiUpdateSponsoredProductsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent
func (a *AdGroupsAPIService) UpdateSponsoredProductsAdGroupsExecute(r ApiUpdateSponsoredProductsAdGroupsRequest) (*SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.UpdateSponsoredProductsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/adGroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spAdGroup.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spAdGroup.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent
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
