package supply_sources_20200701

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SupplySourcesAPIService SupplySourcesAPI service
type SupplySourcesAPIService service

type ApiArchiveSupplySourceRequest struct {
	ctx            context.Context
	ApiService     *SupplySourcesAPIService
	supplySourceId string
}

func (r ApiArchiveSupplySourceRequest) Execute() (*ErrorList, *http.Response, error) {
	return r.ApiService.ArchiveSupplySourceExecute(r)
}

/*
ArchiveSupplySource Method for ArchiveSupplySource

Archive a supply source, making it inactive. Cannot be undone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param supplySourceId The unique identifier of a supply source.
	@return ApiArchiveSupplySourceRequest
*/
func (a *SupplySourcesAPIService) ArchiveSupplySource(ctx context.Context, supplySourceId string) ApiArchiveSupplySourceRequest {
	return ApiArchiveSupplySourceRequest{
		ApiService:     a,
		ctx:            ctx,
		supplySourceId: supplySourceId,
	}
}

// Execute executes the request
//
//	@return ErrorList
func (a *SupplySourcesAPIService) ArchiveSupplySourceExecute(r ApiArchiveSupplySourceRequest) (*ErrorList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupplySourcesAPIService.ArchiveSupplySource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/supplySources/2020-07-01/supplySources/{supplySourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"supplySourceId"+"}", url.PathEscape(parameterValueToString(r.supplySourceId, "supplySourceId")), -1)

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

type ApiCreateSupplySourceRequest struct {
	ctx        context.Context
	ApiService *SupplySourcesAPIService
	payload    *CreateSupplySourceRequest
}

// A request to create a supply source.
func (r ApiCreateSupplySourceRequest) Payload(payload CreateSupplySourceRequest) ApiCreateSupplySourceRequest {
	r.payload = &payload
	return r
}

func (r ApiCreateSupplySourceRequest) Execute() (*CreateSupplySourceResponse, *http.Response, error) {
	return r.ApiService.CreateSupplySourceExecute(r)
}

/*
CreateSupplySource Method for CreateSupplySource

Create a new supply source.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateSupplySourceRequest
*/
func (a *SupplySourcesAPIService) CreateSupplySource(ctx context.Context) ApiCreateSupplySourceRequest {
	return ApiCreateSupplySourceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSupplySourceResponse
func (a *SupplySourcesAPIService) CreateSupplySourceExecute(r ApiCreateSupplySourceRequest) (*CreateSupplySourceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSupplySourceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupplySourcesAPIService.CreateSupplySource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/supplySources/2020-07-01/supplySources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return localVarReturnValue, nil, reportError("payload is required and must be specified")
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
	localVarPostBody = r.payload
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

type ApiGetSupplySourceRequest struct {
	ctx            context.Context
	ApiService     *SupplySourcesAPIService
	supplySourceId string
}

func (r ApiGetSupplySourceRequest) Execute() (*SupplySource, *http.Response, error) {
	return r.ApiService.GetSupplySourceExecute(r)
}

/*
GetSupplySource Method for GetSupplySource

Retrieve a supply source.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param supplySourceId The unique identifier of a supply source.
	@return ApiGetSupplySourceRequest
*/
func (a *SupplySourcesAPIService) GetSupplySource(ctx context.Context, supplySourceId string) ApiGetSupplySourceRequest {
	return ApiGetSupplySourceRequest{
		ApiService:     a,
		ctx:            ctx,
		supplySourceId: supplySourceId,
	}
}

// Execute executes the request
//
//	@return SupplySource
func (a *SupplySourcesAPIService) GetSupplySourceExecute(r ApiGetSupplySourceRequest) (*SupplySource, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SupplySource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupplySourcesAPIService.GetSupplySource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/supplySources/2020-07-01/supplySources/{supplySourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"supplySourceId"+"}", url.PathEscape(parameterValueToString(r.supplySourceId, "supplySourceId")), -1)

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

type ApiGetSupplySourcesRequest struct {
	ctx           context.Context
	ApiService    *SupplySourcesAPIService
	nextPageToken *string
	pageSize      *float32
}

// The pagination token to retrieve a specific page of results.
func (r ApiGetSupplySourcesRequest) NextPageToken(nextPageToken string) ApiGetSupplySourcesRequest {
	r.nextPageToken = &nextPageToken
	return r
}

// The number of supply sources to return per paginated request.
func (r ApiGetSupplySourcesRequest) PageSize(pageSize float32) ApiGetSupplySourcesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetSupplySourcesRequest) Execute() (*GetSupplySourcesResponse, *http.Response, error) {
	return r.ApiService.GetSupplySourcesExecute(r)
}

/*
GetSupplySources Method for GetSupplySources

The path to retrieve paginated supply sources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSupplySourcesRequest
*/
func (a *SupplySourcesAPIService) GetSupplySources(ctx context.Context) ApiGetSupplySourcesRequest {
	return ApiGetSupplySourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSupplySourcesResponse
func (a *SupplySourcesAPIService) GetSupplySourcesExecute(r ApiGetSupplySourcesRequest) (*GetSupplySourcesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSupplySourcesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupplySourcesAPIService.GetSupplySources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/supplySources/2020-07-01/supplySources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextPageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextPageToken", r.nextPageToken, "", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "", "")
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

type ApiUpdateSupplySourceRequest struct {
	ctx            context.Context
	ApiService     *SupplySourcesAPIService
	supplySourceId string
	payload        *UpdateSupplySourceRequest
}

func (r ApiUpdateSupplySourceRequest) Payload(payload UpdateSupplySourceRequest) ApiUpdateSupplySourceRequest {
	r.payload = &payload
	return r
}

func (r ApiUpdateSupplySourceRequest) Execute() (*ErrorList, *http.Response, error) {
	return r.ApiService.UpdateSupplySourceExecute(r)
}

/*
UpdateSupplySource Method for UpdateSupplySource

Update the configuration and capabilities of a supply source.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param supplySourceId The unique identitier of a supply source.
	@return ApiUpdateSupplySourceRequest
*/
func (a *SupplySourcesAPIService) UpdateSupplySource(ctx context.Context, supplySourceId string) ApiUpdateSupplySourceRequest {
	return ApiUpdateSupplySourceRequest{
		ApiService:     a,
		ctx:            ctx,
		supplySourceId: supplySourceId,
	}
}

// Execute executes the request
//
//	@return ErrorList
func (a *SupplySourcesAPIService) UpdateSupplySourceExecute(r ApiUpdateSupplySourceRequest) (*ErrorList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupplySourcesAPIService.UpdateSupplySource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/supplySources/2020-07-01/supplySources/{supplySourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"supplySourceId"+"}", url.PathEscape(parameterValueToString(r.supplySourceId, "supplySourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.payload
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

type ApiUpdateSupplySourceStatusRequest struct {
	ctx            context.Context
	ApiService     *SupplySourcesAPIService
	supplySourceId string
	payload        *UpdateSupplySourceStatusRequest
}

func (r ApiUpdateSupplySourceStatusRequest) Payload(payload UpdateSupplySourceStatusRequest) ApiUpdateSupplySourceStatusRequest {
	r.payload = &payload
	return r
}

func (r ApiUpdateSupplySourceStatusRequest) Execute() (*ErrorList, *http.Response, error) {
	return r.ApiService.UpdateSupplySourceStatusExecute(r)
}

/*
UpdateSupplySourceStatus Method for UpdateSupplySourceStatus

Update the status of a supply source.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param supplySourceId The unique identifier of a supply source.
	@return ApiUpdateSupplySourceStatusRequest
*/
func (a *SupplySourcesAPIService) UpdateSupplySourceStatus(ctx context.Context, supplySourceId string) ApiUpdateSupplySourceStatusRequest {
	return ApiUpdateSupplySourceStatusRequest{
		ApiService:     a,
		ctx:            ctx,
		supplySourceId: supplySourceId,
	}
}

// Execute executes the request
//
//	@return ErrorList
func (a *SupplySourcesAPIService) UpdateSupplySourceStatusExecute(r ApiUpdateSupplySourceStatusRequest) (*ErrorList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupplySourcesAPIService.UpdateSupplySourceStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/supplySources/2020-07-01/supplySources/{supplySourceId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"supplySourceId"+"}", url.PathEscape(parameterValueToString(r.supplySourceId, "supplySourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.payload
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
