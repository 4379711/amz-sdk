package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// BrandSafetyListAPIService BrandSafetyListAPI service
type BrandSafetyListAPIService service

type ApiCreateBrandSafetyDenyListDomainsRequest struct {
	ctx                          context.Context
	ApiService                   *BrandSafetyListAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	brandSafetyPostRequest       *BrandSafetyPostRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateBrandSafetyDenyListDomainsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateBrandSafetyDenyListDomainsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateBrandSafetyDenyListDomainsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateBrandSafetyDenyListDomainsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of Brand Safety List Domain objects. For each object, specify required fields and their values. Maximum length of the array is 10,000 objects.
func (r ApiCreateBrandSafetyDenyListDomainsRequest) BrandSafetyPostRequest(brandSafetyPostRequest BrandSafetyPostRequest) ApiCreateBrandSafetyDenyListDomainsRequest {
	r.brandSafetyPostRequest = &brandSafetyPostRequest
	return r
}

func (r ApiCreateBrandSafetyDenyListDomainsRequest) Execute() (*BrandSafetyUpdateResponse, *http.Response, error) {
	return r.ApiService.CreateBrandSafetyDenyListDomainsExecute(r)
}

func (a *BrandSafetyListAPIService) CreateBrandSafetyDenyListDomains(ctx context.Context) ApiCreateBrandSafetyDenyListDomainsRequest {
	return ApiCreateBrandSafetyDenyListDomainsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandSafetyUpdateResponse
func (a *BrandSafetyListAPIService) CreateBrandSafetyDenyListDomainsExecute(r ApiCreateBrandSafetyDenyListDomainsRequest) (*BrandSafetyUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BrandSafetyUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandSafetyListAPIService.CreateBrandSafetyDenyListDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/brandSafety/deny"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.brandSafetyPostRequest == nil {
		return localVarReturnValue, nil, reportError("brandSafetyPostRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.brandSafetyPostRequest
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

type ApiDeleteBrandSafetyDenyListRequest struct {
	ctx                          context.Context
	ApiService                   *BrandSafetyListAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteBrandSafetyDenyListRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteBrandSafetyDenyListRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteBrandSafetyDenyListRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteBrandSafetyDenyListRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteBrandSafetyDenyListRequest) Execute() (*BrandSafetyUpdateResponse, *http.Response, error) {
	return r.ApiService.DeleteBrandSafetyDenyListExecute(r)
}

func (a *BrandSafetyListAPIService) DeleteBrandSafetyDenyList(ctx context.Context) ApiDeleteBrandSafetyDenyListRequest {
	return ApiDeleteBrandSafetyDenyListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandSafetyUpdateResponse
func (a *BrandSafetyListAPIService) DeleteBrandSafetyDenyListExecute(r ApiDeleteBrandSafetyDenyListRequest) (*BrandSafetyUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BrandSafetyUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandSafetyListAPIService.DeleteBrandSafetyDenyList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/brandSafety/deny"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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

type ApiGetRequestResultsRequest struct {
	ctx                          context.Context
	ApiService                   *BrandSafetyListAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	requestId                    string
	startIndex                   *int32
	count                        *int32
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetRequestResultsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetRequestResultsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetRequestResultsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetRequestResultsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of results. Use in conjunction with the count parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiGetRequestResultsRequest) StartIndex(startIndex int32) ApiGetRequestResultsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of results in the returned array. Use in conjunction with the startIndex parameter to control pagination. For example, to return the first 1000 results set startIndex&#x3D;0 and count&#x3D;1000. To return the next 1000 results, set startIndex&#x3D;1000 and count&#x3D;1000, and so on. Defaults to max page size(1000).
func (r ApiGetRequestResultsRequest) Count(count int32) ApiGetRequestResultsRequest {
	r.count = &count
	return r
}

func (r ApiGetRequestResultsRequest) Execute() (*BrandSafetyRequestResultsResponse, *http.Response, error) {
	return r.ApiService.GetRequestResultsExecute(r)
}

func (a *BrandSafetyListAPIService) GetRequestResults(ctx context.Context, requestId string) ApiGetRequestResultsRequest {
	return ApiGetRequestResultsRequest{
		ApiService: a,
		ctx:        ctx,
		requestId:  requestId,
	}
}

// Execute executes the request
//
//	@return BrandSafetyRequestResultsResponse
func (a *BrandSafetyListAPIService) GetRequestResultsExecute(r ApiGetRequestResultsRequest) (*BrandSafetyRequestResultsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BrandSafetyRequestResultsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandSafetyListAPIService.GetRequestResults")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/brandSafety/{requestId}/results"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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

type ApiGetRequestStatusRequest struct {
	ctx                          context.Context
	ApiService                   *BrandSafetyListAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	requestId                    string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetRequestStatusRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetRequestStatusRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetRequestStatusRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetRequestStatusRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetRequestStatusRequest) Execute() (*BrandSafetyRequestStatusResponse, *http.Response, error) {
	return r.ApiService.GetRequestStatusExecute(r)
}

func (a *BrandSafetyListAPIService) GetRequestStatus(ctx context.Context, requestId string) ApiGetRequestStatusRequest {
	return ApiGetRequestStatusRequest{
		ApiService: a,
		ctx:        ctx,
		requestId:  requestId,
	}
}

// Execute executes the request
//
//	@return BrandSafetyRequestStatusResponse
func (a *BrandSafetyListAPIService) GetRequestStatusExecute(r ApiGetRequestStatusRequest) (*BrandSafetyRequestStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BrandSafetyRequestStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandSafetyListAPIService.GetRequestStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/brandSafety/{requestId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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

type ApiListDomainsRequest struct {
	ctx                          context.Context
	ApiService                   *BrandSafetyListAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListDomainsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListDomainsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListDomainsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListDomainsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of domains. Use in conjunction with the count parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListDomainsRequest) StartIndex(startIndex int32) ApiListDomainsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of domain objects in the returned array. Use in conjunction with the startIndex parameter to control pagination. For example, to return the first 1000 domains set startIndex&#x3D;0 and count&#x3D;1000. To return the next 1000 domains, set startIndex&#x3D;1000 and count&#x3D;1000, and so on. Defaults to max page size(1000).
func (r ApiListDomainsRequest) Count(count int32) ApiListDomainsRequest {
	r.count = &count
	return r
}

func (r ApiListDomainsRequest) Execute() (*BrandSafetyGetResponse, *http.Response, error) {
	return r.ApiService.ListDomainsExecute(r)
}

func (a *BrandSafetyListAPIService) ListDomains(ctx context.Context) ApiListDomainsRequest {
	return ApiListDomainsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandSafetyGetResponse
func (a *BrandSafetyListAPIService) ListDomainsExecute(r ApiListDomainsRequest) (*BrandSafetyGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BrandSafetyGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandSafetyListAPIService.ListDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/brandSafety/deny"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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

type ApiListRequestStatusRequest struct {
	ctx                          context.Context
	ApiService                   *BrandSafetyListAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListRequestStatusRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListRequestStatusRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListRequestStatusRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListRequestStatusRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListRequestStatusRequest) Execute() (*BrandSafetyListRequestStatusResponse, *http.Response, error) {
	return r.ApiService.ListRequestStatusExecute(r)
}

func (a *BrandSafetyListAPIService) ListRequestStatus(ctx context.Context) ApiListRequestStatusRequest {
	return ApiListRequestStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandSafetyListRequestStatusResponse
func (a *BrandSafetyListAPIService) ListRequestStatusExecute(r ApiListRequestStatusRequest) (*BrandSafetyListRequestStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BrandSafetyListRequestStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandSafetyListAPIService.ListRequestStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/brandSafety/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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
