package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SnapshotsAPIService SnapshotsAPI service
type SnapshotsAPIService service

type ApiCreateSnapshotRequest struct {
	ctx                          context.Context
	ApiService                   *SnapshotsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	recordType                   string
	snapshotRequest              *SnapshotRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSnapshotRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSnapshotRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSnapshotRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSnapshotRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Request a snapshot file for all entities of a single record type.
func (r ApiCreateSnapshotRequest) SnapshotRequest(snapshotRequest SnapshotRequest) ApiCreateSnapshotRequest {
	r.snapshotRequest = &snapshotRequest
	return r
}

func (r ApiCreateSnapshotRequest) Execute() (*SnapshotResponse, *http.Response, error) {
	return r.ApiService.CreateSnapshotExecute(r)
}

func (a *SnapshotsAPIService) CreateSnapshot(ctx context.Context, recordType string) ApiCreateSnapshotRequest {
	return ApiCreateSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		recordType: recordType,
	}
}

// Execute executes the request
//
//	@return SnapshotResponse
//
// Deprecated
func (a *SnapshotsAPIService) CreateSnapshotExecute(r ApiCreateSnapshotRequest) (*SnapshotResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SnapshotResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.CreateSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/{recordType}/snapshot"
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
	localVarPostBody = r.snapshotRequest
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

type ApiDownloadSnapshotRequest struct {
	ctx                          context.Context
	ApiService                   *SnapshotsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	snapshotId                   string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDownloadSnapshotRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDownloadSnapshotRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiDownloadSnapshotRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDownloadSnapshotRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDownloadSnapshotRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadSnapshotExecute(r)
}

func (a *SnapshotsAPIService) DownloadSnapshot(ctx context.Context, snapshotId string) ApiDownloadSnapshotRequest {
	return ApiDownloadSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
// Deprecated
func (a *SnapshotsAPIService) DownloadSnapshotExecute(r ApiDownloadSnapshotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.DownloadSnapshot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/snapshots/{snapshotId}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

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
