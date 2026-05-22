package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ReportsAPIService ReportsAPI service
type ReportsAPIService service

type ApiDownloadReportRequest struct {
	ctx                          context.Context
	ApiService                   *ReportsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	reportId                     string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDownloadReportRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDownloadReportRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDownloadReportRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDownloadReportRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDownloadReportRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadReportExecute(r)
}

func (a *ReportsAPIService) DownloadReport(ctx context.Context, reportId string) ApiDownloadReportRequest {
	return ApiDownloadReportRequest{
		ApiService: a,
		ctx:        ctx,
		reportId:   reportId,
	}
}

// Execute executes the request
func (a *ReportsAPIService) DownloadReportExecute(r ApiDownloadReportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.DownloadReport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/reports/{reportId}/download"
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

	if localVarHTTPResponse.StatusCode >= 400 {
		err = &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
	}
	return localVarHTTPResponse, err
}

type ApiRequestReportRequest struct {
	ctx                          context.Context
	ApiService                   *ReportsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	recordType                   string
	reportRequest                *ReportRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiRequestReportRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiRequestReportRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiRequestReportRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiRequestReportRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiRequestReportRequest) ReportRequest(reportRequest ReportRequest) ApiRequestReportRequest {
	r.reportRequest = &reportRequest
	return r
}

func (r ApiRequestReportRequest) Execute() (*ReportResponse, *http.Response, error) {
	return r.ApiService.RequestReportExecute(r)
}

func (a *ReportsAPIService) RequestReport(ctx context.Context, recordType string) ApiRequestReportRequest {
	return ApiRequestReportRequest{
		ApiService: a,
		ctx:        ctx,
		recordType: recordType,
	}
}

// Execute executes the request
//
//	@return ReportResponse
//
// Deprecated
func (a *ReportsAPIService) RequestReportExecute(r ApiRequestReportRequest) (*ReportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.RequestReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/{recordType}/report"
	localVarPath = strings.Replace(localVarPath, "{"+"recordType"+"}", url.PathEscape(parameterValueToString(r.recordType, "recordType")), -1)

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
	localVarPostBody = r.reportRequest
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
