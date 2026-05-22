package ams

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DSPStreamSubscriptionAPIService DSPStreamSubscriptionAPI service
type DSPStreamSubscriptionAPIService service

type ApiCreateDspStreamSubscriptionRequest struct {
	ctx                                       context.Context
	ApiService                                *DSPStreamSubscriptionAPIService
	amazonAdsAccountID                        *string
	amazonAdvertisingAPIClientId              *string
	createDspStreamSubscriptionRequestContent *CreateDspStreamSubscriptionRequestContent
}

// The identifier of a DSP advertiser level account
func (r ApiCreateDspStreamSubscriptionRequest) AmazonAdsAccountID(amazonAdsAccountID string) ApiCreateDspStreamSubscriptionRequest {
	r.amazonAdsAccountID = &amazonAdsAccountID
	return r
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateDspStreamSubscriptionRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateDspStreamSubscriptionRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiCreateDspStreamSubscriptionRequest) CreateDspStreamSubscriptionRequestContent(createDspStreamSubscriptionRequestContent CreateDspStreamSubscriptionRequestContent) ApiCreateDspStreamSubscriptionRequest {
	r.createDspStreamSubscriptionRequestContent = &createDspStreamSubscriptionRequestContent
	return r
}

func (r ApiCreateDspStreamSubscriptionRequest) Execute() (*CreateDspStreamSubscriptionResponseContent, *http.Response, error) {
	return r.ApiService.CreateDspStreamSubscriptionExecute(r)
}

func (a *DSPStreamSubscriptionAPIService) CreateDspStreamSubscription(ctx context.Context) ApiCreateDspStreamSubscriptionRequest {
	return ApiCreateDspStreamSubscriptionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateDspStreamSubscriptionResponseContent
func (a *DSPStreamSubscriptionAPIService) CreateDspStreamSubscriptionExecute(r ApiCreateDspStreamSubscriptionRequest) (*CreateDspStreamSubscriptionResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateDspStreamSubscriptionResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DSPStreamSubscriptionAPIService.CreateDspStreamSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dsp/streams/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.amazonAdsAccountID == nil {
		return localVarReturnValue, nil, reportError("amazonAdsAccountID is required and must be specified")
	}

	if r.createDspStreamSubscriptionRequestContent == nil {
		return localVarReturnValue, nil, reportError("createDspStreamSubscriptionRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.amazonmarketingstreamsubscriptions.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.amazonmarketingstreamsubscriptions.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-Account-ID", r.amazonAdsAccountID, "simple", "")

	// body params
	localVarPostBody = r.createDspStreamSubscriptionRequestContent
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

type ApiGetDspStreamSubscriptionRequest struct {
	ctx                          context.Context
	ApiService                   *DSPStreamSubscriptionAPIService
	subscriptionId               string
	amazonAdsAccountID           *string
	amazonAdvertisingAPIClientId *string
}

// The identifier of a DSP advertiser level account
func (r ApiGetDspStreamSubscriptionRequest) AmazonAdsAccountID(amazonAdsAccountID string) ApiGetDspStreamSubscriptionRequest {
	r.amazonAdsAccountID = &amazonAdsAccountID
	return r
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetDspStreamSubscriptionRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetDspStreamSubscriptionRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiGetDspStreamSubscriptionRequest) Execute() (*GetDspStreamSubscriptionResponseContent, *http.Response, error) {
	return r.ApiService.GetDspStreamSubscriptionExecute(r)
}

func (a *DSPStreamSubscriptionAPIService) GetDspStreamSubscription(ctx context.Context, subscriptionId string) ApiGetDspStreamSubscriptionRequest {
	return ApiGetDspStreamSubscriptionRequest{
		ApiService:     a,
		ctx:            ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//
//	@return GetDspStreamSubscriptionResponseContent
func (a *DSPStreamSubscriptionAPIService) GetDspStreamSubscriptionExecute(r ApiGetDspStreamSubscriptionRequest) (*GetDspStreamSubscriptionResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetDspStreamSubscriptionResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DSPStreamSubscriptionAPIService.GetDspStreamSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dsp/streams/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.subscriptionId) > 255 {
		return localVarReturnValue, nil, reportError("subscriptionId must have less than 255 elements")
	}
	if r.amazonAdsAccountID == nil {
		return localVarReturnValue, nil, reportError("amazonAdsAccountID is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.amazonmarketingstreamsubscriptions.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-Account-ID", r.amazonAdsAccountID, "simple", "")

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

type ApiListDspStreamSubscriptionsRequest struct {
	ctx                          context.Context
	ApiService                   *DSPStreamSubscriptionAPIService
	amazonAdsAccountID           *string
	amazonAdvertisingAPIClientId *string
	maxResults                   *float32
	startingToken                *string
}

// The identifier of a DSP advertiser level account
func (r ApiListDspStreamSubscriptionsRequest) AmazonAdsAccountID(amazonAdsAccountID string) ApiListDspStreamSubscriptionsRequest {
	r.amazonAdsAccountID = &amazonAdsAccountID
	return r
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListDspStreamSubscriptionsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListDspStreamSubscriptionsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// desired number of entries in the response, defaults to maximum value
func (r ApiListDspStreamSubscriptionsRequest) MaxResults(maxResults float32) ApiListDspStreamSubscriptionsRequest {
	r.maxResults = &maxResults
	return r
}

// Token which can be used to get the next page of results, if more entries exist
func (r ApiListDspStreamSubscriptionsRequest) StartingToken(startingToken string) ApiListDspStreamSubscriptionsRequest {
	r.startingToken = &startingToken
	return r
}

func (r ApiListDspStreamSubscriptionsRequest) Execute() (*ListDspStreamSubscriptionsResponseContent, *http.Response, error) {
	return r.ApiService.ListDspStreamSubscriptionsExecute(r)
}

func (a *DSPStreamSubscriptionAPIService) ListDspStreamSubscriptions(ctx context.Context) ApiListDspStreamSubscriptionsRequest {
	return ApiListDspStreamSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDspStreamSubscriptionsResponseContent
func (a *DSPStreamSubscriptionAPIService) ListDspStreamSubscriptionsExecute(r ApiListDspStreamSubscriptionsRequest) (*ListDspStreamSubscriptionsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListDspStreamSubscriptionsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DSPStreamSubscriptionAPIService.ListDspStreamSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dsp/streams/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.amazonAdsAccountID == nil {
		return localVarReturnValue, nil, reportError("amazonAdsAccountID is required and must be specified")
	}

	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxResults", r.maxResults, "form", "")
	}
	if r.startingToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingToken", r.startingToken, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.amazonmarketingstreamsubscriptions.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-Account-ID", r.amazonAdsAccountID, "simple", "")

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

type ApiUpdateDspStreamSubscriptionRequest struct {
	ctx                                       context.Context
	ApiService                                *DSPStreamSubscriptionAPIService
	subscriptionId                            string
	amazonAdsAccountID                        *string
	amazonAdvertisingAPIClientId              *string
	updateDspStreamSubscriptionRequestContent *UpdateDspStreamSubscriptionRequestContent
}

// The identifier of a DSP advertiser level account
func (r ApiUpdateDspStreamSubscriptionRequest) AmazonAdsAccountID(amazonAdsAccountID string) ApiUpdateDspStreamSubscriptionRequest {
	r.amazonAdsAccountID = &amazonAdsAccountID
	return r
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateDspStreamSubscriptionRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateDspStreamSubscriptionRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiUpdateDspStreamSubscriptionRequest) UpdateDspStreamSubscriptionRequestContent(updateDspStreamSubscriptionRequestContent UpdateDspStreamSubscriptionRequestContent) ApiUpdateDspStreamSubscriptionRequest {
	r.updateDspStreamSubscriptionRequestContent = &updateDspStreamSubscriptionRequestContent
	return r
}

func (r ApiUpdateDspStreamSubscriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateDspStreamSubscriptionExecute(r)
}

func (a *DSPStreamSubscriptionAPIService) UpdateDspStreamSubscription(ctx context.Context, subscriptionId string) ApiUpdateDspStreamSubscriptionRequest {
	return ApiUpdateDspStreamSubscriptionRequest{
		ApiService:     a,
		ctx:            ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *DSPStreamSubscriptionAPIService) UpdateDspStreamSubscriptionExecute(r ApiUpdateDspStreamSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DSPStreamSubscriptionAPIService.UpdateDspStreamSubscription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dsp/streams/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.subscriptionId) > 255 {
		return nil, reportError("subscriptionId must have less than 255 elements")
	}
	if r.amazonAdsAccountID == nil {
		return nil, reportError("amazonAdsAccountID is required and must be specified")
	}
	if r.amazonAdvertisingAPIClientId == nil {
		return nil, reportError("amazonAdvertisingAPIClientId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.amazonmarketingstreamsubscriptions.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.amazonmarketingstreamsubscriptions.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-Account-ID", r.amazonAdsAccountID, "simple", "")

	// body params
	localVarPostBody = r.updateDspStreamSubscriptionRequestContent
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
