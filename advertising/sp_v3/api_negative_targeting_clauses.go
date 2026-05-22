package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// NegativeTargetingClausesAPIService NegativeTargetingClausesAPI service
type NegativeTargetingClausesAPIService service

type ApiCreateSponsoredProductsNegativeTargetingClausesRequest struct {
	ctx                                                                            context.Context
	ApiService                                                                     *NegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                   *string
	amazonAdvertisingAPIScope                                                      *string
	sponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent
	prefer                                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredProductsNegativeTargetingClausesRequest) SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent(sponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent) ApiCreateSponsoredProductsNegativeTargetingClausesRequest {
	r.sponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent = &sponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiCreateSponsoredProductsNegativeTargetingClausesRequest) Prefer(prefer string) ApiCreateSponsoredProductsNegativeTargetingClausesRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreateSponsoredProductsNegativeTargetingClausesRequest) Execute() (*SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredProductsNegativeTargetingClausesExecute(r)
}

func (a *NegativeTargetingClausesAPIService) CreateSponsoredProductsNegativeTargetingClauses(ctx context.Context) ApiCreateSponsoredProductsNegativeTargetingClausesRequest {
	return ApiCreateSponsoredProductsNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent
func (a *NegativeTargetingClausesAPIService) CreateSponsoredProductsNegativeTargetingClausesExecute(r ApiCreateSponsoredProductsNegativeTargetingClausesRequest) (*SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsCreateSponsoredProductsNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingClausesAPIService.CreateSponsoredProductsNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeTargets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spNegativeTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsCreateSponsoredProductsNegativeTargetingClausesRequestContent
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

type ApiDeleteSponsoredProductsNegativeTargetingClausesRequest struct {
	ctx                                                                            context.Context
	ApiService                                                                     *NegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                   *string
	amazonAdvertisingAPIScope                                                      *string
	sponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredProductsNegativeTargetingClausesRequest) SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent(sponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) ApiDeleteSponsoredProductsNegativeTargetingClausesRequest {
	r.sponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent = &sponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent
	return r
}

func (r ApiDeleteSponsoredProductsNegativeTargetingClausesRequest) Execute() (*SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredProductsNegativeTargetingClausesExecute(r)
}

func (a *NegativeTargetingClausesAPIService) DeleteSponsoredProductsNegativeTargetingClauses(ctx context.Context) ApiDeleteSponsoredProductsNegativeTargetingClausesRequest {
	return ApiDeleteSponsoredProductsNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent
func (a *NegativeTargetingClausesAPIService) DeleteSponsoredProductsNegativeTargetingClausesExecute(r ApiDeleteSponsoredProductsNegativeTargetingClausesRequest) (*SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingClausesAPIService.DeleteSponsoredProductsNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeTargets/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spNegativeTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent
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

type ApiListSponsoredProductsNegativeTargetingClausesRequest struct {
	ctx                                                                          context.Context
	ApiService                                                                   *NegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                 *string
	amazonAdvertisingAPIScope                                                    *string
	sponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent *SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredProductsNegativeTargetingClausesRequest) SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent(sponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent SponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent) ApiListSponsoredProductsNegativeTargetingClausesRequest {
	r.sponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent = &sponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent
	return r
}

func (r ApiListSponsoredProductsNegativeTargetingClausesRequest) Execute() (*SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredProductsNegativeTargetingClausesExecute(r)
}

func (a *NegativeTargetingClausesAPIService) ListSponsoredProductsNegativeTargetingClauses(ctx context.Context) ApiListSponsoredProductsNegativeTargetingClausesRequest {
	return ApiListSponsoredProductsNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent
func (a *NegativeTargetingClausesAPIService) ListSponsoredProductsNegativeTargetingClausesExecute(r ApiListSponsoredProductsNegativeTargetingClausesRequest) (*SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingClausesAPIService.ListSponsoredProductsNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeTargets/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spNegativeTargetingClause.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sponsoredProductsListSponsoredProductsNegativeTargetingClausesRequestContent
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

type ApiUpdateSponsoredProductsNegativeTargetingClausesRequest struct {
	ctx                                                                            context.Context
	ApiService                                                                     *NegativeTargetingClausesAPIService
	amazonAdvertisingAPIClientId                                                   *string
	amazonAdvertisingAPIScope                                                      *string
	sponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent
	prefer                                                                         *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateSponsoredProductsNegativeTargetingClausesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredProductsNegativeTargetingClausesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredProductsNegativeTargetingClausesRequest) SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent(sponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent) ApiUpdateSponsoredProductsNegativeTargetingClausesRequest {
	r.sponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent = &sponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids. Please note that the extendedData field will be part of the full object for /list endpoints only.
func (r ApiUpdateSponsoredProductsNegativeTargetingClausesRequest) Prefer(prefer string) ApiUpdateSponsoredProductsNegativeTargetingClausesRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdateSponsoredProductsNegativeTargetingClausesRequest) Execute() (*SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredProductsNegativeTargetingClausesExecute(r)
}

func (a *NegativeTargetingClausesAPIService) UpdateSponsoredProductsNegativeTargetingClauses(ctx context.Context) ApiUpdateSponsoredProductsNegativeTargetingClausesRequest {
	return ApiUpdateSponsoredProductsNegativeTargetingClausesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent
func (a *NegativeTargetingClausesAPIService) UpdateSponsoredProductsNegativeTargetingClausesExecute(r ApiUpdateSponsoredProductsNegativeTargetingClausesRequest) (*SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegativeTargetingClausesAPIService.UpdateSponsoredProductsNegativeTargetingClauses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeTargets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent == nil {
		return localVarReturnValue, nil, reportError("sponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spNegativeTargetingClause.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spNegativeTargetingClause.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.sponsoredProductsUpdateSponsoredProductsNegativeTargetingClausesRequestContent
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
