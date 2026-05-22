package reports_20210630

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ReportsAPIService ReportsAPI service
type ReportsAPIService service

type ApiCancelReportRequest struct {
	ctx        context.Context
	ApiService *ReportsAPIService
	reportId   string
}

func (r ApiCancelReportRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelReportExecute(r)
}

/*
CancelReport Method for CancelReport

Cancels the report that you specify. Only reports with `processingStatus=IN_QUEUE` can be cancelled. Cancelled reports are returned in subsequent calls to the `getReport` and `getReports` operations.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reportId The identifier for the report. This identifier is unique only in combination with a seller ID.
	@return ApiCancelReportRequest
*/
func (a *ReportsAPIService) CancelReport(ctx context.Context, reportId string) ApiCancelReportRequest {
	return ApiCancelReportRequest{
		ApiService: a,
		ctx:        ctx,
		reportId:   reportId,
	}
}

// Execute executes the request
func (a *ReportsAPIService) CancelReportExecute(r ApiCancelReportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.CancelReport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/reports/{reportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reportId"+"}", url.PathEscape(parameterValueToString(r.reportId, "reportId")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	return localVarHTTPResponse, nil
}

type ApiCancelReportScheduleRequest struct {
	ctx              context.Context
	ApiService       *ReportsAPIService
	reportScheduleId string
}

func (r ApiCancelReportScheduleRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelReportScheduleExecute(r)
}

/*
CancelReportSchedule Method for CancelReportSchedule

Cancels the report schedule that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reportScheduleId The identifier for the report schedule. This identifier is unique only in combination with a seller ID.
	@return ApiCancelReportScheduleRequest
*/
func (a *ReportsAPIService) CancelReportSchedule(ctx context.Context, reportScheduleId string) ApiCancelReportScheduleRequest {
	return ApiCancelReportScheduleRequest{
		ApiService:       a,
		ctx:              ctx,
		reportScheduleId: reportScheduleId,
	}
}

// Execute executes the request
func (a *ReportsAPIService) CancelReportScheduleExecute(r ApiCancelReportScheduleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.CancelReportSchedule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/schedules/{reportScheduleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reportScheduleId"+"}", url.PathEscape(parameterValueToString(r.reportScheduleId, "reportScheduleId")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	return localVarHTTPResponse, nil
}

type ApiCreateReportRequest struct {
	ctx        context.Context
	ApiService *ReportsAPIService
	body       *CreateReportSpecification
}

// Information required to create the report.
func (r ApiCreateReportRequest) Body(body CreateReportSpecification) ApiCreateReportRequest {
	r.body = &body
	return r
}

func (r ApiCreateReportRequest) Execute() (*CreateReportResponse, *http.Response, error) {
	return r.ApiService.CreateReportExecute(r)
}

/*
CreateReport Method for CreateReport

Creates a report.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0167 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateReportRequest
*/
func (a *ReportsAPIService) CreateReport(ctx context.Context) ApiCreateReportRequest {
	return ApiCreateReportRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateReportResponse
func (a *ReportsAPIService) CreateReportExecute(r ApiCreateReportRequest) (*CreateReportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateReportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.CreateReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type ApiCreateReportScheduleRequest struct {
	ctx        context.Context
	ApiService *ReportsAPIService
	body       *CreateReportScheduleSpecification
}

// Information required to create the report schedule.
func (r ApiCreateReportScheduleRequest) Body(body CreateReportScheduleSpecification) ApiCreateReportScheduleRequest {
	r.body = &body
	return r
}

func (r ApiCreateReportScheduleRequest) Execute() (*CreateReportScheduleResponse, *http.Response, error) {
	return r.ApiService.CreateReportScheduleExecute(r)
}

/*
CreateReportSchedule Method for CreateReportSchedule

Creates a report schedule. If a report schedule with the same report type and marketplace IDs already exists, it will be cancelled and replaced with this one.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateReportScheduleRequest
*/
func (a *ReportsAPIService) CreateReportSchedule(ctx context.Context) ApiCreateReportScheduleRequest {
	return ApiCreateReportScheduleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateReportScheduleResponse
func (a *ReportsAPIService) CreateReportScheduleExecute(r ApiCreateReportScheduleRequest) (*CreateReportScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateReportScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.CreateReportSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/schedules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type ApiGetReportRequest struct {
	ctx        context.Context
	ApiService *ReportsAPIService
	reportId   string
}

func (r ApiGetReportRequest) Execute() (*Report, *http.Response, error) {
	return r.ApiService.GetReportExecute(r)
}

/*
GetReport Method for GetReport

Returns report details (including the `reportDocumentId`, if available) for the report that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reportId The identifier for the report. This identifier is unique only in combination with a seller ID.
	@return ApiGetReportRequest
*/
func (a *ReportsAPIService) GetReport(ctx context.Context, reportId string) ApiGetReportRequest {
	return ApiGetReportRequest{
		ApiService: a,
		ctx:        ctx,
		reportId:   reportId,
	}
}

// Execute executes the request
//
//	@return Report
func (a *ReportsAPIService) GetReportExecute(r ApiGetReportRequest) (*Report, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Report
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/reports/{reportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reportId"+"}", url.PathEscape(parameterValueToString(r.reportId, "reportId")), -1)

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

type ApiGetReportDocumentRequest struct {
	ctx              context.Context
	ApiService       *ReportsAPIService
	reportDocumentId string
}

func (r ApiGetReportDocumentRequest) Execute() (*ReportDocument, *http.Response, error) {
	return r.ApiService.GetReportDocumentExecute(r)
}

/*
GetReportDocument Method for GetReportDocument

Returns the information required for retrieving a report document's contents.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0167 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reportDocumentId The identifier for the report document.
	@return ApiGetReportDocumentRequest
*/
func (a *ReportsAPIService) GetReportDocument(ctx context.Context, reportDocumentId string) ApiGetReportDocumentRequest {
	return ApiGetReportDocumentRequest{
		ApiService:       a,
		ctx:              ctx,
		reportDocumentId: reportDocumentId,
	}
}

// Execute executes the request
//
//	@return ReportDocument
func (a *ReportsAPIService) GetReportDocumentExecute(r ApiGetReportDocumentRequest) (*ReportDocument, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReportDocument
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReportDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/documents/{reportDocumentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reportDocumentId"+"}", url.PathEscape(parameterValueToString(r.reportDocumentId, "reportDocumentId")), -1)

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

type ApiGetReportScheduleRequest struct {
	ctx              context.Context
	ApiService       *ReportsAPIService
	reportScheduleId string
}

func (r ApiGetReportScheduleRequest) Execute() (*ReportSchedule, *http.Response, error) {
	return r.ApiService.GetReportScheduleExecute(r)
}

/*
GetReportSchedule Method for GetReportSchedule

Returns report schedule details for the report schedule that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reportScheduleId The identifier for the report schedule. This identifier is unique only in combination with a seller ID.
	@return ApiGetReportScheduleRequest
*/
func (a *ReportsAPIService) GetReportSchedule(ctx context.Context, reportScheduleId string) ApiGetReportScheduleRequest {
	return ApiGetReportScheduleRequest{
		ApiService:       a,
		ctx:              ctx,
		reportScheduleId: reportScheduleId,
	}
}

// Execute executes the request
//
//	@return ReportSchedule
func (a *ReportsAPIService) GetReportScheduleExecute(r ApiGetReportScheduleRequest) (*ReportSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReportSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReportSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/schedules/{reportScheduleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reportScheduleId"+"}", url.PathEscape(parameterValueToString(r.reportScheduleId, "reportScheduleId")), -1)

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

type ApiGetReportSchedulesRequest struct {
	ctx         context.Context
	ApiService  *ReportsAPIService
	reportTypes *[]string
}

// A list of report types used to filter report schedules. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information.
func (r ApiGetReportSchedulesRequest) ReportTypes(reportTypes []string) ApiGetReportSchedulesRequest {
	r.reportTypes = &reportTypes
	return r
}

func (r ApiGetReportSchedulesRequest) Execute() (*ReportScheduleList, *http.Response, error) {
	return r.ApiService.GetReportSchedulesExecute(r)
}

/*
GetReportSchedules Method for GetReportSchedules

Returns report schedule details that match the filters that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetReportSchedulesRequest
*/
func (a *ReportsAPIService) GetReportSchedules(ctx context.Context) ApiGetReportSchedulesRequest {
	return ApiGetReportSchedulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportScheduleList
func (a *ReportsAPIService) GetReportSchedulesExecute(r ApiGetReportSchedulesRequest) (*ReportScheduleList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReportScheduleList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReportSchedules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/schedules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reportTypes == nil {
		return localVarReturnValue, nil, reportError("reportTypes is required and must be specified")
	}
	if len(*r.reportTypes) < 1 {
		return localVarReturnValue, nil, reportError("reportTypes must have at least 1 elements")
	}
	if len(*r.reportTypes) > 10 {
		return localVarReturnValue, nil, reportError("reportTypes must have less than 10 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "reportTypes", r.reportTypes, "form", "csv")
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

type ApiGetReportsRequest struct {
	ctx                context.Context
	ApiService         *ReportsAPIService
	reportTypes        *[]string
	processingStatuses *[]string
	marketplaceIds     *[]string
	pageSize           *int32
	createdSince       *time.Time
	createdUntil       *time.Time
	nextToken          *string
}

// A list of report types used to filter reports. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. When reportTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either reportTypes or nextToken is required.
func (r ApiGetReportsRequest) ReportTypes(reportTypes []string) ApiGetReportsRequest {
	r.reportTypes = &reportTypes
	return r
}

// A list of processing statuses used to filter reports.
func (r ApiGetReportsRequest) ProcessingStatuses(processingStatuses []string) ApiGetReportsRequest {
	r.processingStatuses = &processingStatuses
	return r
}

// A list of marketplace identifiers used to filter reports. The reports returned will match at least one of the marketplaces that you specify.
func (r ApiGetReportsRequest) MarketplaceIds(marketplaceIds []string) ApiGetReportsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// The maximum number of reports to return in a single call.
func (r ApiGetReportsRequest) PageSize(pageSize int32) ApiGetReportsRequest {
	r.pageSize = &pageSize
	return r
}

// The earliest report creation date and time for reports to include in the response, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. The default is 90 days ago. Reports are retained for a maximum of 90 days.
func (r ApiGetReportsRequest) CreatedSince(createdSince time.Time) ApiGetReportsRequest {
	r.createdSince = &createdSince
	return r
}

// The latest report creation date and time for reports to include in the response, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. The default is now.
func (r ApiGetReportsRequest) CreatedUntil(createdUntil time.Time) ApiGetReportsRequest {
	r.createdUntil = &createdUntil
	return r
}

// A string token returned in the response to your previous request. &#x60;nextToken&#x60; is returned when the number of results exceeds the specified &#x60;pageSize&#x60; value. To get the next page of results, call the &#x60;getReports&#x60; operation and include this token as the only parameter. Specifying &#x60;nextToken&#x60; with any other parameters will cause the request to fail.
func (r ApiGetReportsRequest) NextToken(nextToken string) ApiGetReportsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetReportsRequest) Execute() (*GetReportsResponse, *http.Response, error) {
	return r.ApiService.GetReportsExecute(r)
}

/*
GetReports Method for GetReports

Returns report details for the reports that match the filters that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetReportsRequest
*/
func (a *ReportsAPIService) GetReports(ctx context.Context) ApiGetReportsRequest {
	return ApiGetReportsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetReportsResponse
func (a *ReportsAPIService) GetReportsExecute(r ApiGetReportsRequest) (*GetReportsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetReportsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/2021-06-30/reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.reportTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reportTypes", r.reportTypes, "form", "csv")
	}
	if r.processingStatuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processingStatuses", r.processingStatuses, "form", "csv")
	}
	if r.marketplaceIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "", "")
	}
	if r.createdSince != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdSince", r.createdSince, "", "")
	}
	if r.createdUntil != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdUntil", r.createdUntil, "", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "", "")
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
