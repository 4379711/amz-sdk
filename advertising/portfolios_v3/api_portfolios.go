package portfolios_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// PortfoliosAPIService PortfoliosAPI service
type PortfoliosAPIService service

type ApiCreatePortfoliosRequest struct {
	ctx                            context.Context
	ApiService                     *PortfoliosAPIService
	createPortfoliosRequestContent *CreatePortfoliosRequestContent
	amazonAdvertisingAPIClientId   *string
	amazonAdvertisingAPIScope      *string
	prefer                         *string
}

func (r ApiCreatePortfoliosRequest) CreatePortfoliosRequestContent(createPortfoliosRequestContent CreatePortfoliosRequestContent) ApiCreatePortfoliosRequest {
	r.createPortfoliosRequestContent = &createPortfoliosRequestContent
	return r
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreatePortfoliosRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreatePortfoliosRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreatePortfoliosRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreatePortfoliosRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;canonicalObjectId - (default) - returns objectIds encoded using canonical representation return&#x3D;obfuscatedObjectId - returns obfuscated objectIds return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids
func (r ApiCreatePortfoliosRequest) Prefer(prefer string) ApiCreatePortfoliosRequest {
	r.prefer = &prefer
	return r
}

func (r ApiCreatePortfoliosRequest) Execute() (*CreatePortfoliosResponseContent, *http.Response, error) {
	return r.ApiService.CreatePortfoliosExecute(r)
}

func (a *PortfoliosAPIService) CreatePortfolios(ctx context.Context) ApiCreatePortfoliosRequest {
	return ApiCreatePortfoliosRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreatePortfoliosResponseContent
func (a *PortfoliosAPIService) CreatePortfoliosExecute(r ApiCreatePortfoliosRequest) (*CreatePortfoliosResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreatePortfoliosResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfoliosAPIService.CreatePortfolios")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/portfolios"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createPortfoliosRequestContent == nil {
		return localVarReturnValue, nil, reportError("createPortfoliosRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spPortfolio.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spPortfolio.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.createPortfoliosRequestContent
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

type ApiListPortfoliosRequest struct {
	ctx                          context.Context
	ApiService                   *PortfoliosAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	prefer                       *string
	listPortfoliosRequestContent *ListPortfoliosRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListPortfoliosRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListPortfoliosRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListPortfoliosRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListPortfoliosRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;canonicalObjectId - (default) - returns objectIds encoded using canonical representation return&#x3D;obfuscatedObjectId - returns obfuscated objectIds return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids
func (r ApiListPortfoliosRequest) Prefer(prefer string) ApiListPortfoliosRequest {
	r.prefer = &prefer
	return r
}

func (r ApiListPortfoliosRequest) ListPortfoliosRequestContent(listPortfoliosRequestContent ListPortfoliosRequestContent) ApiListPortfoliosRequest {
	r.listPortfoliosRequestContent = &listPortfoliosRequestContent
	return r
}

func (r ApiListPortfoliosRequest) Execute() (*ListPortfoliosResponseContent, *http.Response, error) {
	return r.ApiService.ListPortfoliosExecute(r)
}

func (a *PortfoliosAPIService) ListPortfolios(ctx context.Context) ApiListPortfoliosRequest {
	return ApiListPortfoliosRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListPortfoliosResponseContent
func (a *PortfoliosAPIService) ListPortfoliosExecute(r ApiListPortfoliosRequest) (*ListPortfoliosResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListPortfoliosResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfoliosAPIService.ListPortfolios")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/portfolios/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spPortfolio.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spPortfolio.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.listPortfoliosRequestContent
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

type ApiUpdatePortfoliosRequest struct {
	ctx                            context.Context
	ApiService                     *PortfoliosAPIService
	updatePortfoliosRequestContent *UpdatePortfoliosRequestContent
	amazonAdvertisingAPIClientId   *string
	amazonAdvertisingAPIScope      *string
	prefer                         *string
}

func (r ApiUpdatePortfoliosRequest) UpdatePortfoliosRequestContent(updatePortfoliosRequestContent UpdatePortfoliosRequestContent) ApiUpdatePortfoliosRequest {
	r.updatePortfoliosRequestContent = &updatePortfoliosRequestContent
	return r
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdatePortfoliosRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdatePortfoliosRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdatePortfoliosRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdatePortfoliosRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The \&quot;Prefer\&quot; header, as defined in [RFC7240], allows clients to request certain behavior from the service. The service ignores preference values that are either not supported or not known by the service. Either multiple Prefer headers are passed or single one with comma separated values, both forms are equivalent Supported preferences: return&#x3D;canonicalObjectId - (default) - returns objectIds encoded using canonical representation return&#x3D;obfuscatedObjectId - returns obfuscated objectIds return&#x3D;representation - return the full object when doing create/update/delete operations instead of ids
func (r ApiUpdatePortfoliosRequest) Prefer(prefer string) ApiUpdatePortfoliosRequest {
	r.prefer = &prefer
	return r
}

func (r ApiUpdatePortfoliosRequest) Execute() (*UpdatePortfoliosResponseContent, *http.Response, error) {
	return r.ApiService.UpdatePortfoliosExecute(r)
}

func (a *PortfoliosAPIService) UpdatePortfolios(ctx context.Context) ApiUpdatePortfoliosRequest {
	return ApiUpdatePortfoliosRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdatePortfoliosResponseContent
func (a *PortfoliosAPIService) UpdatePortfoliosExecute(r ApiUpdatePortfoliosRequest) (*UpdatePortfoliosResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdatePortfoliosResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfoliosAPIService.UpdatePortfolios")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/portfolios"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updatePortfoliosRequestContent == nil {
		return localVarReturnValue, nil, reportError("updatePortfoliosRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spPortfolio.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.spPortfolio.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.updatePortfoliosRequestContent
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
