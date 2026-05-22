package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CampaignsAPIService CampaignsAPI service
type CampaignsAPIService service

type ApiCreateSponsoredBrandsCampaignsRequest struct {
	ctx                                          context.Context
	ApiService                                   *CampaignsAPIService
	amazonAdvertisingAPIClientId                 *string
	amazonAdvertisingAPIScope                    *string
	createSponsoredBrandsCampaignsRequestContent *CreateSponsoredBrandsCampaignsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandsCampaignsRequest) CreateSponsoredBrandsCampaignsRequestContent(createSponsoredBrandsCampaignsRequestContent CreateSponsoredBrandsCampaignsRequestContent) ApiCreateSponsoredBrandsCampaignsRequest {
	r.createSponsoredBrandsCampaignsRequestContent = &createSponsoredBrandsCampaignsRequestContent
	return r
}

func (r ApiCreateSponsoredBrandsCampaignsRequest) Execute() (*CreateSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandsCampaignsExecute(r)
}

func (a *CampaignsAPIService) CreateSponsoredBrandsCampaigns(ctx context.Context) ApiCreateSponsoredBrandsCampaignsRequest {
	return ApiCreateSponsoredBrandsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandsCampaignsResponseContent
func (a *CampaignsAPIService) CreateSponsoredBrandsCampaignsExecute(r ApiCreateSponsoredBrandsCampaignsRequest) (*CreateSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.CreateSponsoredBrandsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandsCampaignsRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandsCampaignsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandsCampaignsRequestContent
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

type ApiDeleteSponsoredBrandsCampaignsRequest struct {
	ctx                                          context.Context
	ApiService                                   *CampaignsAPIService
	amazonAdvertisingAPIClientId                 *string
	amazonAdvertisingAPIScope                    *string
	deleteSponsoredBrandsCampaignsRequestContent *DeleteSponsoredBrandsCampaignsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDeleteSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDeleteSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiDeleteSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDeleteSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDeleteSponsoredBrandsCampaignsRequest) DeleteSponsoredBrandsCampaignsRequestContent(deleteSponsoredBrandsCampaignsRequestContent DeleteSponsoredBrandsCampaignsRequestContent) ApiDeleteSponsoredBrandsCampaignsRequest {
	r.deleteSponsoredBrandsCampaignsRequestContent = &deleteSponsoredBrandsCampaignsRequestContent
	return r
}

func (r ApiDeleteSponsoredBrandsCampaignsRequest) Execute() (*DeleteSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.DeleteSponsoredBrandsCampaignsExecute(r)
}

func (a *CampaignsAPIService) DeleteSponsoredBrandsCampaigns(ctx context.Context) ApiDeleteSponsoredBrandsCampaignsRequest {
	return ApiDeleteSponsoredBrandsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeleteSponsoredBrandsCampaignsResponseContent
func (a *CampaignsAPIService) DeleteSponsoredBrandsCampaignsExecute(r ApiDeleteSponsoredBrandsCampaignsRequest) (*DeleteSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteSponsoredBrandsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.DeleteSponsoredBrandsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/campaigns/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.deleteSponsoredBrandsCampaignsRequestContent
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

type ApiListSponsoredBrandsCampaignsRequest struct {
	ctx                                        context.Context
	ApiService                                 *CampaignsAPIService
	amazonAdvertisingAPIClientId               *string
	amazonAdvertisingAPIScope                  *string
	listSponsoredBrandsCampaignsRequestContent *ListSponsoredBrandsCampaignsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiListSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredBrandsCampaignsRequest) ListSponsoredBrandsCampaignsRequestContent(listSponsoredBrandsCampaignsRequestContent ListSponsoredBrandsCampaignsRequestContent) ApiListSponsoredBrandsCampaignsRequest {
	r.listSponsoredBrandsCampaignsRequestContent = &listSponsoredBrandsCampaignsRequestContent
	return r
}

func (r ApiListSponsoredBrandsCampaignsRequest) Execute() (*ListSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredBrandsCampaignsExecute(r)
}

func (a *CampaignsAPIService) ListSponsoredBrandsCampaigns(ctx context.Context) ApiListSponsoredBrandsCampaignsRequest {
	return ApiListSponsoredBrandsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListSponsoredBrandsCampaignsResponseContent
func (a *CampaignsAPIService) ListSponsoredBrandsCampaignsExecute(r ApiListSponsoredBrandsCampaignsRequest) (*ListSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSponsoredBrandsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.ListSponsoredBrandsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/campaigns/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.listSponsoredBrandsCampaignsRequestContent
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

type ApiUpdateSponsoredBrandsCampaignsRequest struct {
	ctx                                          context.Context
	ApiService                                   *CampaignsAPIService
	amazonAdvertisingAPIClientId                 *string
	amazonAdvertisingAPIScope                    *string
	updateSponsoredBrandsCampaignsRequestContent *UpdateSponsoredBrandsCampaignsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiUpdateSponsoredBrandsCampaignsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredBrandsCampaignsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredBrandsCampaignsRequest) UpdateSponsoredBrandsCampaignsRequestContent(updateSponsoredBrandsCampaignsRequestContent UpdateSponsoredBrandsCampaignsRequestContent) ApiUpdateSponsoredBrandsCampaignsRequest {
	r.updateSponsoredBrandsCampaignsRequestContent = &updateSponsoredBrandsCampaignsRequestContent
	return r
}

func (r ApiUpdateSponsoredBrandsCampaignsRequest) Execute() (*UpdateSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredBrandsCampaignsExecute(r)
}

func (a *CampaignsAPIService) UpdateSponsoredBrandsCampaigns(ctx context.Context) ApiUpdateSponsoredBrandsCampaignsRequest {
	return ApiUpdateSponsoredBrandsCampaignsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateSponsoredBrandsCampaignsResponseContent
func (a *CampaignsAPIService) UpdateSponsoredBrandsCampaignsExecute(r ApiUpdateSponsoredBrandsCampaignsRequest) (*UpdateSponsoredBrandsCampaignsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSponsoredBrandsCampaignsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignsAPIService.UpdateSponsoredBrandsCampaigns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/v4/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSponsoredBrandsCampaignsRequestContent == nil {
		return localVarReturnValue, nil, reportError("updateSponsoredBrandsCampaignsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbcampaignresource.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.updateSponsoredBrandsCampaignsRequestContent
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
