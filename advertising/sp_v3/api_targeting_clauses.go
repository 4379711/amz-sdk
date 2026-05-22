package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// TargetingClausesAPIService TargetingClausesAPI service
type TargetingClausesAPIService service

type ApiCreateSponsoredProductsTargetingClausesRequest struct {
	ctx                                                                    context.Context
	ApiService                                                             *TargetingClausesAPIService
	amazonAdvertisingAPIClientId                                           *string
	amazonAdvertisingAPIScope                                              *string
	sponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent *SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent
	prefer                                                                 *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsTargetingClausesRequest) SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent(sponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent) ApiCreateSponsoredProductsTargetingClausesRequest {
	r.sponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent = &sponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsTargetingClausesRequest) Prefer(prefer string) ApiCreateSponsoredProductsTargetingClausesRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsTargetingClausesRequest) Execute() (*SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsTargetingClausesExecute(r)
}

func (a *TargetingClausesAPIService) CreateSponsoredProductsTargetingClauses(ctx context.Context) ApiCreateSponsoredProductsTargetingClausesRequest {
	return ApiCreateSponsoredProductsTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent
func (a *TargetingClausesAPIService) CreateSponsoredProductsTargetingClausesExecute(r ApiCreateSponsoredProductsTargetingClausesRequest) (*SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingClausesAPIService.CreateSponsoredProductsTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsTargetingClausesRequestContent
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

type ApiDeleteSponsoredProductsTargetingClausesRequest struct {
	ctx                                                                    context.Context
	ApiService                                                             *TargetingClausesAPIService
	amazonAdvertisingAPIClientId                                           *string
	amazonAdvertisingAPIScope                                              *string
	sponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsTargetingClausesRequest) SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent(sponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) ApiDeleteSponsoredProductsTargetingClausesRequest {
	r.sponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent = &sponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsTargetingClausesRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsTargetingClausesExecute(r)
}

func (a *TargetingClausesAPIService) DeleteSponsoredProductsTargetingClauses(ctx context.Context) ApiDeleteSponsoredProductsTargetingClausesRequest {
	return ApiDeleteSponsoredProductsTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent
func (a *TargetingClausesAPIService) DeleteSponsoredProductsTargetingClausesExecute(r ApiDeleteSponsoredProductsTargetingClausesRequest) (*SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingClausesAPIService.DeleteSponsoredProductsTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent
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

type ApiListSponsoredProductsTargetingClausesRequest struct {
	ctx                                                                  context.Context
	ApiService                                                           *TargetingClausesAPIService
	amazonAdvertisingAPIClientId                                         *string
	amazonAdvertisingAPIScope                                            *string
	sponsoredProductsListSponsoredProductsTargetingClausesRequestContent *SponsoredProductsListSponsoredProductsTargetingClausesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsTargetingClausesRequest) SponsoredProductsListSponsoredProductsTargetingClausesRequestContent(sponsoredProductsListSponsoredProductsTargetingClausesRequestContent SponsoredProductsListSponsoredProductsTargetingClausesRequestContent) ApiListSponsoredProductsTargetingClausesRequest {
	r.sponsoredProductsListSponsoredProductsTargetingClausesRequestContent = &sponsoredProductsListSponsoredProductsTargetingClausesRequestContent
	return r
}

func (r ApiListSponsoredProductsTargetingClausesRequest) Execute() (*SponsoredProductsListSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsTargetingClausesExecute(r)
}

func (a *TargetingClausesAPIService) ListSponsoredProductsTargetingClauses(ctx context.Context) ApiListSponsoredProductsTargetingClausesRequest {
	return ApiListSponsoredProductsTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsTargetingClausesResponseContent
func (a *TargetingClausesAPIService) ListSponsoredProductsTargetingClausesExecute(r ApiListSponsoredProductsTargetingClausesRequest) (*SponsoredProductsListSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingClausesAPIService.ListSponsoredProductsTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spTargetingClause.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsTargetingClausesRequestContent
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

type ApiUpdateSponsoredProductsTargetingClausesRequest struct {
	ctx                                                                    context.Context
	ApiService                                                             *TargetingClausesAPIService
	amazonAdvertisingAPIClientId                                           *string
	amazonAdvertisingAPIScope                                              *string
	sponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent *SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent
	prefer                                                                 *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsTargetingClausesRequest) SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent(sponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent) ApiUpdateSponsoredProductsTargetingClausesRequest {
	r.sponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent = &sponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsTargetingClausesRequest) Prefer(prefer string) ApiUpdateSponsoredProductsTargetingClausesRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsTargetingClausesRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsTargetingClausesExecute(r)
}

func (a *TargetingClausesAPIService) UpdateSponsoredProductsTargetingClauses(ctx context.Context) ApiUpdateSponsoredProductsTargetingClausesRequest {
	return ApiUpdateSponsoredProductsTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent
func (a *TargetingClausesAPIService) UpdateSponsoredProductsTargetingClausesExecute(r ApiUpdateSponsoredProductsTargetingClausesRequest) (*SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingClausesAPIService.UpdateSponsoredProductsTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsTargetingClausesRequestContent
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
