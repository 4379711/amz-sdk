package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CreativesAPIService CreativesAPI service
type CreativesAPIService service

type ApiCreateCreativesRequest struct {
	ctx                          context.Context
	ApiService                   *CreativesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createCreative               *[]CreateCreative
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateCreativesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateCreativesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateCreativesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateCreativesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of Creative objects to create. Maximum length of the array is 100 objects. Note - when using productAds with landingPageURL of OFF_AMAZON_LINK, STORE, or MOMENT, the following properties are required all together; 1) headline, 2) brandLogo, and 3) rectCustomImage, squareCustomImage.
func (r ApiCreateCreativesRequest) CreateCreative(createCreative []CreateCreative) ApiCreateCreativesRequest {
	r.createCreative = &createCreative
	return r
}

func (r ApiCreateCreativesRequest) Execute() ([]CreativeResponse, *http.Response, error) {
	return r.ApiService.CreateCreativesExecute(r)
}

func (a *CreativesAPIService) CreateCreatives(ctx context.Context) ApiCreateCreativesRequest {
	return ApiCreateCreativesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CreativeResponse
func (a *CreativesAPIService) CreateCreativesExecute(r ApiCreateCreativesRequest) ([]CreativeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CreativeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreativesAPIService.CreateCreatives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/creatives"

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
	localVarPostBody = r.createCreative
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

type ApiListCreativeModerationsRequest struct {
	ctx                          context.Context
	ApiService                   *CreativesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	language                     *Locale
	startIndex                   *int32
	count                        *int32
	adGroupIdFilter              *string
	creativeIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListCreativeModerationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListCreativeModerationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListCreativeModerationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListCreativeModerationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The language of the returned creative moderation metadata.
func (r ApiListCreativeModerationsRequest) Language(language Locale) ApiListCreativeModerationsRequest {
	r.language = &language
	return r
}

// Sets a cursor into the requested set of creative moderations. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListCreativeModerationsRequest) StartIndex(startIndex int32) ApiListCreativeModerationsRequest {
	r.startIndex = &startIndex
	return r
}

// Sets the number of creative objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten creative moderations set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten creative moderations, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListCreativeModerationsRequest) Count(count int32) ApiListCreativeModerationsRequest {
	r.count = &count
	return r
}

// The returned array includes only creative moderations associated with ad group identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;creativeIdFilter&#x60; parameter.
func (r ApiListCreativeModerationsRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListCreativeModerationsRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// The returned array includes only creative moderations with creative identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;adGroupIdFilter&#x60; parameter.
func (r ApiListCreativeModerationsRequest) CreativeIdFilter(creativeIdFilter string) ApiListCreativeModerationsRequest {
	r.creativeIdFilter = &creativeIdFilter
	return r
}

func (r ApiListCreativeModerationsRequest) Execute() ([]CreativeModeration, *http.Response, error) {
	return r.ApiService.ListCreativeModerationsExecute(r)
}

func (a *CreativesAPIService) ListCreativeModerations(ctx context.Context) ApiListCreativeModerationsRequest {
	return ApiListCreativeModerationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CreativeModeration
func (a *CreativesAPIService) ListCreativeModerationsExecute(r ApiListCreativeModerationsRequest) ([]CreativeModeration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CreativeModeration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreativesAPIService.ListCreativeModerations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/moderation/creatives"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.language == nil {
		return localVarReturnValue, nil, reportError("language is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "form", "")
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 100
		r.count = &defaultValue
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
	}
	if r.creativeIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creativeIdFilter", r.creativeIdFilter, "form", "")
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

type ApiListCreativesRequest struct {
	ctx                          context.Context
	ApiService                   *CreativesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	adGroupIdFilter              *string
	creativeIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListCreativesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListCreativesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListCreativesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListCreativesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Sets a cursor into the requested set of creatives. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListCreativesRequest) StartIndex(startIndex int32) ApiListCreativesRequest {
	r.startIndex = &startIndex
	return r
}

// Sets the number of creative objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten creatives set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten creatives, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListCreativesRequest) Count(count int32) ApiListCreativesRequest {
	r.count = &count
	return r
}

// The returned array includes only creatives associated with ad group identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;creativeIdFilter&#x60; parameter.
func (r ApiListCreativesRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListCreativesRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// The returned array includes only creatives with identifiers matching those specified in the comma-delimited string. Cannot be used in conjunction with the &#x60;adGroupIdFilter&#x60; parameter.
func (r ApiListCreativesRequest) CreativeIdFilter(creativeIdFilter string) ApiListCreativesRequest {
	r.creativeIdFilter = &creativeIdFilter
	return r
}

func (r ApiListCreativesRequest) Execute() ([]Creative, *http.Response, error) {
	return r.ApiService.ListCreativesExecute(r)
}

func (a *CreativesAPIService) ListCreatives(ctx context.Context) ApiListCreativesRequest {
	return ApiListCreativesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Creative
func (a *CreativesAPIService) ListCreativesExecute(r ApiListCreativesRequest) ([]Creative, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Creative
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreativesAPIService.ListCreatives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/creatives"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 100
		r.count = &defaultValue
	}
	if r.adGroupIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adGroupIdFilter", r.adGroupIdFilter, "form", "")
	}
	if r.creativeIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creativeIdFilter", r.creativeIdFilter, "form", "")
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

type ApiPostCreativePreviewRequest struct {
	ctx                          context.Context
	ApiService                   *CreativesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	creativePreviewRequest       *CreativePreviewRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiPostCreativePreviewRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiPostCreativePreviewRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiPostCreativePreviewRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiPostCreativePreviewRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiPostCreativePreviewRequest) CreativePreviewRequest(creativePreviewRequest CreativePreviewRequest) ApiPostCreativePreviewRequest {
	r.creativePreviewRequest = &creativePreviewRequest
	return r
}

func (r ApiPostCreativePreviewRequest) Execute() (*CreativePreviewResponse, *http.Response, error) {
	return r.ApiService.PostCreativePreviewExecute(r)
}

func (a *CreativesAPIService) PostCreativePreview(ctx context.Context) ApiPostCreativePreviewRequest {
	return ApiPostCreativePreviewRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativePreviewResponse
func (a *CreativesAPIService) PostCreativePreviewExecute(r ApiPostCreativePreviewRequest) (*CreativePreviewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreativePreviewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreativesAPIService.PostCreativePreview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/creatives/preview"

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
	localVarPostBody = r.creativePreviewRequest
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

type ApiUpdateCreativesRequest struct {
	ctx                          context.Context
	ApiService                   *CreativesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	creativeUpdate               *[]CreativeUpdate
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateCreativesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateCreativesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateCreativesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateCreativesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of creative objects to update. Maximum length of the array is 100 objects.
func (r ApiUpdateCreativesRequest) CreativeUpdate(creativeUpdate []CreativeUpdate) ApiUpdateCreativesRequest {
	r.creativeUpdate = &creativeUpdate
	return r
}

func (r ApiUpdateCreativesRequest) Execute() ([]CreativeResponse, *http.Response, error) {
	return r.ApiService.UpdateCreativesExecute(r)
}

func (a *CreativesAPIService) UpdateCreatives(ctx context.Context) ApiUpdateCreativesRequest {
	return ApiUpdateCreativesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CreativeResponse
func (a *CreativesAPIService) UpdateCreativesExecute(r ApiUpdateCreativesRequest) ([]CreativeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CreativeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreativesAPIService.UpdateCreatives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/creatives"

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
	localVarPostBody = r.creativeUpdate
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
