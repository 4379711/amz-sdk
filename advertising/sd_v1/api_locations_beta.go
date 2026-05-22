package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// LocationsBetaAPIService LocationsBetaAPI service
type LocationsBetaAPIService service

type ApiArchiveLocationsRequest struct {
	ctx                          context.Context
	ApiService                   *LocationsBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	archiveLocationRequest       *ArchiveLocationRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiArchiveLocationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiArchiveLocationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiArchiveLocationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiArchiveLocationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of up to 1000 Location Expression Ids for archival.
func (r ApiArchiveLocationsRequest) ArchiveLocationRequest(archiveLocationRequest ArchiveLocationRequest) ApiArchiveLocationsRequest {
	r.archiveLocationRequest = &archiveLocationRequest
	return r
}

func (r ApiArchiveLocationsRequest) Execute() ([]ArchiveLocationResponse, *http.Response, error) {
	return r.ApiService.ArchiveLocationsExecute(r)
}

func (a *LocationsBetaAPIService) ArchiveLocations(ctx context.Context) ApiArchiveLocationsRequest {
	return ApiArchiveLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ArchiveLocationResponse
func (a *LocationsBetaAPIService) ArchiveLocationsExecute(r ApiArchiveLocationsRequest) ([]ArchiveLocationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ArchiveLocationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsBetaAPIService.ArchiveLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/locations/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.archiveLocationRequest == nil {
		return localVarReturnValue, nil, reportError("archiveLocationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "a-pplication/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.archiveLocationRequest
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

type ApiCreateLocationsRequest struct {
	ctx                          context.Context
	ApiService                   *LocationsBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createLocation               *[]CreateLocation
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateLocationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateLocationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateLocationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateLocationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// A list of up to 100 Locations for creation for each call.  1000 locations can be added for each ad group.
func (r ApiCreateLocationsRequest) CreateLocation(createLocation []CreateLocation) ApiCreateLocationsRequest {
	r.createLocation = &createLocation
	return r
}

func (r ApiCreateLocationsRequest) Execute() ([]Location, *http.Response, error) {
	return r.ApiService.CreateLocationsExecute(r)
}

func (a *LocationsBetaAPIService) CreateLocations(ctx context.Context) ApiCreateLocationsRequest {
	return ApiCreateLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Location
func (a *LocationsBetaAPIService) CreateLocationsExecute(r ApiCreateLocationsRequest) ([]Location, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Location
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsBetaAPIService.CreateLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/locations"

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
	localVarHTTPHeaderAccepts := []string{"application/json", "a-pplication/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createLocation
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

type ApiListLocationsRequest struct {
	ctx                          context.Context
	ApiService                   *LocationsBetaAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	adGroupIdFilter              *string
	campaignIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListLocationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListLocationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListLocationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListLocationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. 0-indexed record offset for the result set. Defaults to 0.
func (r ApiListLocationsRequest) StartIndex(startIndex int32) ApiListLocationsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Number of records to include in the paged response. Defaults to max page size.
func (r ApiListLocationsRequest) Count(count int32) ApiListLocationsRequest {
	r.count = &count
	return r
}

// Optional. Restricts results to those with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;.
func (r ApiListLocationsRequest) StateFilter(stateFilter string) ApiListLocationsRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional list of comma separated adGroupIds. Restricts results to locations with the specified &#x60;adGroupId&#x60;.
func (r ApiListLocationsRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListLocationsRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional list of comma separated campaignIds. Restricts results to locations with the specified &#x60;campaignId&#x60;.
func (r ApiListLocationsRequest) CampaignIdFilter(campaignIdFilter string) ApiListLocationsRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListLocationsRequest) Execute() ([]Location, *http.Response, error) {
	return r.ApiService.ListLocationsExecute(r)
}

func (a *LocationsBetaAPIService) ListLocations(ctx context.Context) ApiListLocationsRequest {
	return ApiListLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Location
func (a *LocationsBetaAPIService) ListLocationsExecute(r ApiListLocationsRequest) ([]Location, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Location
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsBetaAPIService.ListLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.stateFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stateFilter", r.stateFilter, "form", "")
	} else {
		var defaultValue string = "enabled"
		r.stateFilter = &defaultValue
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
	}
	if r.campaignIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaignIdFilter", r.campaignIdFilter, "form", "")
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
