package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// AdGroupsAPIService AdGroupsAPI service
type AdGroupsAPIService service

type ApiCreateSponsoredBrandsAdGroupsRequest struct {
	ctx                                         context.Context
	ApiService                                  *AdGroupsAPIService
	amazonAdvertisingAPIClientId                *string
	amazonAdvertisingAPIScope                   *string
	createSponsoredBrandsAdGroupsRequestContent *CreateSponsoredBrandsAdGroupsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandsAdGroupsRequest) CreateSponsoredBrandsAdGroupsRequestContent(createSponsoredBrandsAdGroupsRequestContent CreateSponsoredBrandsAdGroupsRequestContent) ApiCreateSponsoredBrandsAdGroupsRequest {
	r.createSponsoredBrandsAdGroupsRequestContent = &createSponsoredBrandsAdGroupsRequestContent
	return r
}

func (r ApiCreateSponsoredBrandsAdGroupsRequest) Execute() (*CreateSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) CreateSponsoredBrandsAdGroups(ctx context.Context) ApiCreateSponsoredBrandsAdGroupsRequest {
	return ApiCreateSponsoredBrandsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandsAdGroupsResponseContent
func (a *AdGroupsAPIService) CreateSponsoredBrandsAdGroupsExecute(r ApiCreateSponsoredBrandsAdGroupsRequest) (*CreateSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.CreateSponsoredBrandsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/adGroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandsAdGroupsRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandsAdGroupsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandsAdGroupsRequestContent
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

type ApiDeleteSponsoredBrandsAdGroupsRequest struct {
	ctx                                         context.Context
	ApiService                                  *AdGroupsAPIService
	amazonAdvertisingAPIClientId                *string
	amazonAdvertisingAPIScope                   *string
	deleteSponsoredBrandsAdGroupsRequestContent *DeleteSponsoredBrandsAdGroupsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiDeleteSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredBrandsAdGroupsRequest) DeleteSponsoredBrandsAdGroupsRequestContent(deleteSponsoredBrandsAdGroupsRequestContent DeleteSponsoredBrandsAdGroupsRequestContent) ApiDeleteSponsoredBrandsAdGroupsRequest {
	r.deleteSponsoredBrandsAdGroupsRequestContent = &deleteSponsoredBrandsAdGroupsRequestContent
	return r
}

func (r ApiDeleteSponsoredBrandsAdGroupsRequest) Execute() (*DeleteSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredBrandsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) DeleteSponsoredBrandsAdGroups(ctx context.Context) ApiDeleteSponsoredBrandsAdGroupsRequest {
	return ApiDeleteSponsoredBrandsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeleteSponsoredBrandsAdGroupsResponseContent
func (a *AdGroupsAPIService) DeleteSponsoredBrandsAdGroupsExecute(r ApiDeleteSponsoredBrandsAdGroupsRequest) (*DeleteSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteSponsoredBrandsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.DeleteSponsoredBrandsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/adGroups/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.deleteSponsoredBrandsAdGroupsRequestContent
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

type ApiListSponsoredBrandsAdGroupsRequest struct {
	ctx                                       context.Context
	ApiService                                *AdGroupsAPIService
	amazonAdvertisingAPIClientId              *string
	amazonAdvertisingAPIScope                 *string
	listSponsoredBrandsAdGroupsRequestContent *ListSponsoredBrandsAdGroupsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiListSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredBrandsAdGroupsRequest) ListSponsoredBrandsAdGroupsRequestContent(listSponsoredBrandsAdGroupsRequestContent ListSponsoredBrandsAdGroupsRequestContent) ApiListSponsoredBrandsAdGroupsRequest {
	r.listSponsoredBrandsAdGroupsRequestContent = &listSponsoredBrandsAdGroupsRequestContent
	return r
}

func (r ApiListSponsoredBrandsAdGroupsRequest) Execute() (*ListSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredBrandsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) ListSponsoredBrandsAdGroups(ctx context.Context) ApiListSponsoredBrandsAdGroupsRequest {
	return ApiListSponsoredBrandsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListSponsoredBrandsAdGroupsResponseContent
func (a *AdGroupsAPIService) ListSponsoredBrandsAdGroupsExecute(r ApiListSponsoredBrandsAdGroupsRequest) (*ListSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSponsoredBrandsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.ListSponsoredBrandsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/adGroups/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.listSponsoredBrandsAdGroupsRequestContent
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

type ApiUpdateSponsoredBrandsAdGroupsRequest struct {
	ctx                                         context.Context
	ApiService                                  *AdGroupsAPIService
	amazonAdvertisingAPIClientId                *string
	amazonAdvertisingAPIScope                   *string
	updateSponsoredBrandsAdGroupsRequestContent *UpdateSponsoredBrandsAdGroupsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiUpdateSponsoredBrandsAdGroupsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredBrandsAdGroupsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredBrandsAdGroupsRequest) UpdateSponsoredBrandsAdGroupsRequestContent(updateSponsoredBrandsAdGroupsRequestContent UpdateSponsoredBrandsAdGroupsRequestContent) ApiUpdateSponsoredBrandsAdGroupsRequest {
	r.updateSponsoredBrandsAdGroupsRequestContent = &updateSponsoredBrandsAdGroupsRequestContent
	return r
}

func (r ApiUpdateSponsoredBrandsAdGroupsRequest) Execute() (*UpdateSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredBrandsAdGroupsExecute(r)
}

func (a *AdGroupsAPIService) UpdateSponsoredBrandsAdGroups(ctx context.Context) ApiUpdateSponsoredBrandsAdGroupsRequest {
	return ApiUpdateSponsoredBrandsAdGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateSponsoredBrandsAdGroupsResponseContent
func (a *AdGroupsAPIService) UpdateSponsoredBrandsAdGroupsExecute(r ApiUpdateSponsoredBrandsAdGroupsRequest) (*UpdateSponsoredBrandsAdGroupsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSponsoredBrandsAdGroupsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdGroupsAPIService.UpdateSponsoredBrandsAdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/adGroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSponsoredBrandsAdGroupsRequestContent == nil {
		return localVarReturnValue, nil, reportError("updateSponsoredBrandsAdGroupsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbadgroupresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.updateSponsoredBrandsAdGroupsRequestContent
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
