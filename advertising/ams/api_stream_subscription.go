package ams

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// StreamSubscriptionAPIService StreamSubscriptionAPI service
type StreamSubscriptionAPIService service

type ApiCreateStreamSubscriptionRequest struct {
	ctx                                    context.Context
	ApiService                             *StreamSubscriptionAPIService
	amazonAdvertisingAPIClientId           *string
	createStreamSubscriptionRequestContent *CreateStreamSubscriptionRequestContent
	amazonAdsAccountId                     *string
	amazonAdvertisingAPIScope              *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateStreamSubscriptionRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateStreamSubscriptionRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiCreateStreamSubscriptionRequest) CreateStreamSubscriptionRequestContent(createStreamSubscriptionRequestContent CreateStreamSubscriptionRequestContent) ApiCreateStreamSubscriptionRequest {
	r.createStreamSubscriptionRequestContent = &createStreamSubscriptionRequestContent
	return r
}

// The identifier of a DSP advertiser level account
func (r ApiCreateStreamSubscriptionRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiCreateStreamSubscriptionRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateStreamSubscriptionRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateStreamSubscriptionRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateStreamSubscriptionRequest) Execute() (*CreateStreamSubscriptionResponseContent, *http.Response, error) {
	return r.ApiService.CreateStreamSubscriptionExecute(r)
}

func (a *StreamSubscriptionAPIService) CreateStreamSubscription(ctx context.Context) ApiCreateStreamSubscriptionRequest {
	return ApiCreateStreamSubscriptionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateStreamSubscriptionResponseContent
func (a *StreamSubscriptionAPIService) CreateStreamSubscriptionExecute(r ApiCreateStreamSubscriptionRequest) (*CreateStreamSubscriptionResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateStreamSubscriptionResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamSubscriptionAPIService.CreateStreamSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createStreamSubscriptionRequestContent == nil {
		return localVarReturnValue, nil, reportError("createStreamSubscriptionRequestContent is required and must be specified")
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
	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	}

	// body params
	localVarPostBody = r.createStreamSubscriptionRequestContent
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

type ApiGetStreamSubscriptionRequest struct {
	ctx                          context.Context
	ApiService                   *StreamSubscriptionAPIService
	subscriptionId               string
	amazonAdvertisingAPIClientId *string
	amazonAdsAccountId           *string
	amazonAdvertisingAPIScope    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetStreamSubscriptionRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetStreamSubscriptionRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a DSP advertiser level account
func (r ApiGetStreamSubscriptionRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiGetStreamSubscriptionRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetStreamSubscriptionRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetStreamSubscriptionRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetStreamSubscriptionRequest) Execute() (*GetStreamSubscriptionResponseContent, *http.Response, error) {
	return r.ApiService.GetStreamSubscriptionExecute(r)
}

func (a *StreamSubscriptionAPIService) GetStreamSubscription(ctx context.Context, subscriptionId string) ApiGetStreamSubscriptionRequest {
	return ApiGetStreamSubscriptionRequest{
		ApiService:     a,
		ctx:            ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//
//	@return GetStreamSubscriptionResponseContent
func (a *StreamSubscriptionAPIService) GetStreamSubscriptionExecute(r ApiGetStreamSubscriptionRequest) (*GetStreamSubscriptionResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetStreamSubscriptionResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamSubscriptionAPIService.GetStreamSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.subscriptionId) > 255 {
		return localVarReturnValue, nil, reportError("subscriptionId must have less than 255 elements")
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
	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
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

type ApiListStreamSubscriptionsRequest struct {
	ctx                          context.Context
	ApiService                   *StreamSubscriptionAPIService
	amazonAdvertisingAPIClientId *string
	maxResults                   *float32
	startingToken                *string
	amazonAdsAccountId           *string
	amazonAdvertisingAPIScope    *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListStreamSubscriptionsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListStreamSubscriptionsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// desired number of entries in the response, defaults to maximum value
func (r ApiListStreamSubscriptionsRequest) MaxResults(maxResults float32) ApiListStreamSubscriptionsRequest {
	r.maxResults = &maxResults
	return r
}

// Token which can be used to get the next page of results, if more entries exist
func (r ApiListStreamSubscriptionsRequest) StartingToken(startingToken string) ApiListStreamSubscriptionsRequest {
	r.startingToken = &startingToken
	return r
}

// The identifier of a DSP advertiser level account
func (r ApiListStreamSubscriptionsRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiListStreamSubscriptionsRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListStreamSubscriptionsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListStreamSubscriptionsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListStreamSubscriptionsRequest) Execute() (*ListStreamSubscriptionsResponseContent, *http.Response, error) {
	return r.ApiService.ListStreamSubscriptionsExecute(r)
}

func (a *StreamSubscriptionAPIService) ListStreamSubscriptions(ctx context.Context) ApiListStreamSubscriptionsRequest {
	return ApiListStreamSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListStreamSubscriptionsResponseContent
func (a *StreamSubscriptionAPIService) ListStreamSubscriptionsExecute(r ApiListStreamSubscriptionsRequest) (*ListStreamSubscriptionsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListStreamSubscriptionsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamSubscriptionAPIService.ListStreamSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
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

type ApiUpdateStreamSubscriptionRequest struct {
	ctx                                    context.Context
	ApiService                             *StreamSubscriptionAPIService
	subscriptionId                         string
	amazonAdvertisingAPIClientId           *string
	amazonAdsAccountId                     *string
	amazonAdvertisingAPIScope              *string
	updateStreamSubscriptionRequestContent *UpdateStreamSubscriptionRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateStreamSubscriptionRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateStreamSubscriptionRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a DSP advertiser level account
func (r ApiUpdateStreamSubscriptionRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiUpdateStreamSubscriptionRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

// The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateStreamSubscriptionRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateStreamSubscriptionRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateStreamSubscriptionRequest) UpdateStreamSubscriptionRequestContent(updateStreamSubscriptionRequestContent UpdateStreamSubscriptionRequestContent) ApiUpdateStreamSubscriptionRequest {
	r.updateStreamSubscriptionRequestContent = &updateStreamSubscriptionRequestContent
	return r
}

func (r ApiUpdateStreamSubscriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateStreamSubscriptionExecute(r)
}

func (a *StreamSubscriptionAPIService) UpdateStreamSubscription(ctx context.Context, subscriptionId string) ApiUpdateStreamSubscriptionRequest {
	return ApiUpdateStreamSubscriptionRequest{
		ApiService:     a,
		ctx:            ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *StreamSubscriptionAPIService) UpdateStreamSubscriptionExecute(r ApiUpdateStreamSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamSubscriptionAPIService.UpdateStreamSubscription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.subscriptionId) > 255 {
		return nil, reportError("subscriptionId must have less than 255 elements")
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
	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	}

	// body params
	localVarPostBody = r.updateStreamSubscriptionRequestContent
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
