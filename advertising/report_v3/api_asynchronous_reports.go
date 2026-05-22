package report_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AsynchronousReportsAPIService AsynchronousReportsAPI service
type AsynchronousReportsAPIService service

type ApiCreateAsyncReportRequest struct {
	ctx                          context.Context
	ApiService                   *AsynchronousReportsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	amazonAdsAccountId           *string
	createAsyncReportRequest     *CreateAsyncReportRequest
}

// The client identifier of the customer making the request.
func (r ApiCreateAsyncReportRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateAsyncReportRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateAsyncReportRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateAsyncReportRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The identifier of a DSP advertiser account.
func (r ApiCreateAsyncReportRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiCreateAsyncReportRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

func (r ApiCreateAsyncReportRequest) CreateAsyncReportRequest(createAsyncReportRequest CreateAsyncReportRequest) ApiCreateAsyncReportRequest {
	r.createAsyncReportRequest = &createAsyncReportRequest
	return r
}

func (r ApiCreateAsyncReportRequest) Execute() (*AsyncReport, *http.Response, error) {
	return r.ApiService.CreateAsyncReportExecute(r)
}

func (a *AsynchronousReportsAPIService) CreateAsyncReport(ctx context.Context) ApiCreateAsyncReportRequest {
	return ApiCreateAsyncReportRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AsyncReport
func (a *AsynchronousReportsAPIService) CreateAsyncReportExecute(r ApiCreateAsyncReportRequest) (*AsyncReport, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AsyncReport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AsynchronousReportsAPIService.CreateAsyncReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting/reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.createasyncreportrequest.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.createasyncreportresponse.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createAsyncReportRequest
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

type ApiDeleteAsyncReportRequest struct {
	ctx                          context.Context
	ApiService                   *AsynchronousReportsAPIService
	reportId                     string
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteAsyncReportRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteAsyncReportRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDeleteAsyncReportRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteAsyncReportRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteAsyncReportRequest) Execute() (*DeleteAsyncReportResponse, *http.Response, error) {
	return r.ApiService.DeleteAsyncReportExecute(r)
}

func (a *AsynchronousReportsAPIService) DeleteAsyncReport(ctx context.Context, reportId string) ApiDeleteAsyncReportRequest {
	return ApiDeleteAsyncReportRequest{
		ApiService: a,
		ctx:        ctx,
		reportId:   reportId,
	}
}

// Execute executes the request
//
//	@return DeleteAsyncReportResponse
func (a *AsynchronousReportsAPIService) DeleteAsyncReportExecute(r ApiDeleteAsyncReportRequest) (*DeleteAsyncReportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteAsyncReportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AsynchronousReportsAPIService.DeleteAsyncReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting/reports/{reportId}"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.deleteasyncreportresponse.v3+json", "application/json"}

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

type ApiGetAsyncReportRequest struct {
	ctx                          context.Context
	ApiService                   *AsynchronousReportsAPIService
	reportId                     string
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetAsyncReportRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetAsyncReportRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetAsyncReportRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetAsyncReportRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetAsyncReportRequest) Execute() (*AsyncReport, *http.Response, error) {
	return r.ApiService.GetAsyncReportExecute(r)
}

func (a *AsynchronousReportsAPIService) GetAsyncReport(ctx context.Context, reportId string) ApiGetAsyncReportRequest {
	return ApiGetAsyncReportRequest{
		ApiService: a,
		ctx:        ctx,
		reportId:   reportId,
	}
}

// Execute executes the request
//
//	@return AsyncReport
func (a *AsynchronousReportsAPIService) GetAsyncReportExecute(r ApiGetAsyncReportRequest) (*AsyncReport, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AsyncReport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AsynchronousReportsAPIService.GetAsyncReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reporting/reports/{reportId}"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.getasyncreportresponse.v3+json", "application/json"}

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
