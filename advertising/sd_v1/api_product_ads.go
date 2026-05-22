package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ProductAdsAPIService ProductAdsAPI service
type ProductAdsAPIService service

type ApiArchiveProductAdRequest struct {
	ctx                          context.Context
	ApiService                   *ProductAdsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	adId                         int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiArchiveProductAdRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiArchiveProductAdRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiArchiveProductAdRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiArchiveProductAdRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiArchiveProductAdRequest) Execute() (*ProductAdResponse, *http.Response, error) {
	return r.ApiService.ArchiveProductAdExecute(r)
}

func (a *ProductAdsAPIService) ArchiveProductAd(ctx context.Context, adId int64) ApiArchiveProductAdRequest {
	return ApiArchiveProductAdRequest{
		ApiService: a,
		ctx:        ctx,
		adId:       adId,
	}
}

// Execute executes the request
//
//	@return ProductAdResponse
func (a *ProductAdsAPIService) ArchiveProductAdExecute(r ApiArchiveProductAdRequest) (*ProductAdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductAdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.ArchiveProductAd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/productAds/{adId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adId"+"}", url.PathEscape(parameterValueToString(r.adId, "adId")), -1)

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

type ApiCreateProductAdsRequest struct {
	ctx                          context.Context
	ApiService                   *ProductAdsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	createProductAd              *[]CreateProductAd
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateProductAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateProductAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateProductAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateProductAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of ProductAd objects. For each object, specify required fields and their values. Required fields are &#x60;adGroupId&#x60; and &#x60;state&#x60;. Within any campaign, you must pass consistent fields of either &#x60;asin&#x60; (for vendors), &#x60;sku&#x60; (for sellers), or &#x60;landingPageURL&#x60; (when linking to other pages), these cannot be mixed for any given campaign. Maximum length of the array is 100 objects.
func (r ApiCreateProductAdsRequest) CreateProductAd(createProductAd []CreateProductAd) ApiCreateProductAdsRequest {
	r.createProductAd = &createProductAd
	return r
}

func (r ApiCreateProductAdsRequest) Execute() ([]ProductAdResponse, *http.Response, error) {
	return r.ApiService.CreateProductAdsExecute(r)
}

func (a *ProductAdsAPIService) CreateProductAds(ctx context.Context) ApiCreateProductAdsRequest {
	return ApiCreateProductAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ProductAdResponse
func (a *ProductAdsAPIService) CreateProductAdsExecute(r ApiCreateProductAdsRequest) ([]ProductAdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ProductAdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.CreateProductAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/productAds"

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
	localVarPostBody = r.createProductAd
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

type ApiGetProductAdRequest struct {
	ctx                          context.Context
	ApiService                   *ProductAdsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	adId                         int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetProductAdRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetProductAdRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetProductAdRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetProductAdRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetProductAdRequest) Execute() (*ProductAd, *http.Response, error) {
	return r.ApiService.GetProductAdExecute(r)
}

func (a *ProductAdsAPIService) GetProductAd(ctx context.Context, adId int64) ApiGetProductAdRequest {
	return ApiGetProductAdRequest{
		ApiService: a,
		ctx:        ctx,
		adId:       adId,
	}
}

// Execute executes the request
//
//	@return ProductAd
func (a *ProductAdsAPIService) GetProductAdExecute(r ApiGetProductAdRequest) (*ProductAd, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductAd
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.GetProductAd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/productAds/{adId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adId"+"}", url.PathEscape(parameterValueToString(r.adId, "adId")), -1)

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

type ApiGetProductAdResponseExRequest struct {
	ctx                          context.Context
	ApiService                   *ProductAdsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	adId                         int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetProductAdResponseExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetProductAdResponseExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetProductAdResponseExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetProductAdResponseExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetProductAdResponseExRequest) Execute() (*ProductAdResponseEx, *http.Response, error) {
	return r.ApiService.GetProductAdResponseExExecute(r)
}

func (a *ProductAdsAPIService) GetProductAdResponseEx(ctx context.Context, adId int64) ApiGetProductAdResponseExRequest {
	return ApiGetProductAdResponseExRequest{
		ApiService: a,
		ctx:        ctx,
		adId:       adId,
	}
}

// Execute executes the request
//
//	@return ProductAdResponseEx
func (a *ProductAdsAPIService) GetProductAdResponseExExecute(r ApiGetProductAdResponseExRequest) (*ProductAdResponseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductAdResponseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.GetProductAdResponseEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/productAds/extended/{adId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adId"+"}", url.PathEscape(parameterValueToString(r.adId, "adId")), -1)

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

type ApiListProductAdsRequest struct {
	ctx                          context.Context
	ApiService                   *ProductAdsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	adIdFilter                   *string
	adGroupIdFilter              *string
	campaignIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListProductAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListProductAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListProductAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListProductAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of product ads. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListProductAdsRequest) StartIndex(startIndex int32) ApiListProductAdsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of ProductAd objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten product ad set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten product ads, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListProductAdsRequest) Count(count int32) ApiListProductAdsRequest {
	r.count = &count
	return r
}

// Optional. The returned array is filtered to include only products ads associated with campaigns that have state set to one of the values in the comma-delimited list.
func (r ApiListProductAdsRequest) StateFilter(stateFilter string) ApiListProductAdsRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. The returned array includes only product ads with identifiers matching those in the comma-delimited string.
func (r ApiListProductAdsRequest) AdIdFilter(adIdFilter string) ApiListProductAdsRequest {
	r.adIdFilter = &adIdFilter
	return r
}

// Optional. The returned array is filtered to include only products ads associated with ad groups identifiers in the comma-delimited list.
func (r ApiListProductAdsRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListProductAdsRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. The returned array is filtered to include only product ads associated with the campaign identifiers in the comma-delimited list.
func (r ApiListProductAdsRequest) CampaignIdFilter(campaignIdFilter string) ApiListProductAdsRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListProductAdsRequest) Execute() ([]ProductAd, *http.Response, error) {
	return r.ApiService.ListProductAdsExecute(r)
}

func (a *ProductAdsAPIService) ListProductAds(ctx context.Context) ApiListProductAdsRequest {
	return ApiListProductAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ProductAd
func (a *ProductAdsAPIService) ListProductAdsExecute(r ApiListProductAdsRequest) ([]ProductAd, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ProductAd
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.ListProductAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/productAds"

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
		var defaultValue string = "enabled, paused, archived"
		r.stateFilter = &defaultValue
	}
	if r.adIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adIdFilter", r.adIdFilter, "form", "")
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

type ApiListProductAdsExRequest struct {
	ctx                          context.Context
	ApiService                   *ProductAdsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	startIndex                   *int32
	count                        *int32
	stateFilter                  *string
	adIdFilter                   *string
	adGroupIdFilter              *string
	campaignIdFilter             *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListProductAdsExRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListProductAdsExRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiListProductAdsExRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListProductAdsExRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Sets a cursor into the requested set of product ads. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.
func (r ApiListProductAdsExRequest) StartIndex(startIndex int32) ApiListProductAdsExRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. Sets the number of ProduceAdEx objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten product ads set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size.
func (r ApiListProductAdsExRequest) Count(count int32) ApiListProductAdsExRequest {
	r.count = &count
	return r
}

// Optional. The returned array is filtered to include only campaigns with state set to one of the values in the specified comma-delimited list.
func (r ApiListProductAdsExRequest) StateFilter(stateFilter string) ApiListProductAdsExRequest {
	r.stateFilter = &stateFilter
	return r
}

// Optional. The returned array includes only product ads with identifiers matching those in the comma-delimited string.
func (r ApiListProductAdsExRequest) AdIdFilter(adIdFilter string) ApiListProductAdsExRequest {
	r.adIdFilter = &adIdFilter
	return r
}

// Optional. The returned array is filtered to include only products ads associated with ad groups identifiers in the comma-delimited list.
func (r ApiListProductAdsExRequest) AdGroupIdFilter(adGroupIdFilter string) ApiListProductAdsExRequest {
	r.adGroupIdFilter = &adGroupIdFilter
	return r
}

// Optional. The returned array is filtered to include only product ads associated with the campaign identifiers in the comma-delimited list.
func (r ApiListProductAdsExRequest) CampaignIdFilter(campaignIdFilter string) ApiListProductAdsExRequest {
	r.campaignIdFilter = &campaignIdFilter
	return r
}

func (r ApiListProductAdsExRequest) Execute() ([]ProductAdResponseEx, *http.Response, error) {
	return r.ApiService.ListProductAdsExExecute(r)
}

func (a *ProductAdsAPIService) ListProductAdsEx(ctx context.Context) ApiListProductAdsExRequest {
	return ApiListProductAdsExRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ProductAdResponseEx
func (a *ProductAdsAPIService) ListProductAdsExExecute(r ApiListProductAdsExRequest) ([]ProductAdResponseEx, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ProductAdResponseEx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.ListProductAdsEx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/productAds/extended"

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
		var defaultValue string = "enabled, paused, archived"
		r.stateFilter = &defaultValue
	}
	if r.adIdFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adIdFilter", r.adIdFilter, "form", "")
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

type ApiUpdateProductAdsRequest struct {
	ctx                          context.Context
	ApiService                   *ProductAdsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	updateProductAd              *[]UpdateProductAd
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateProductAdsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateProductAdsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiUpdateProductAdsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateProductAdsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// An array of ProductAd objects. For each object, specify a product ad identifier and the only mutable field, &#x60;state&#x60;. Maximum length of the array is 100 objects.
func (r ApiUpdateProductAdsRequest) UpdateProductAd(updateProductAd []UpdateProductAd) ApiUpdateProductAdsRequest {
	r.updateProductAd = &updateProductAd
	return r
}

func (r ApiUpdateProductAdsRequest) Execute() ([]ProductAdResponse, *http.Response, error) {
	return r.ApiService.UpdateProductAdsExecute(r)
}

func (a *ProductAdsAPIService) UpdateProductAds(ctx context.Context) ApiUpdateProductAdsRequest {
	return ApiUpdateProductAdsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ProductAdResponse
func (a *ProductAdsAPIService) UpdateProductAdsExecute(r ApiUpdateProductAdsRequest) ([]ProductAdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ProductAdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAdsAPIService.UpdateProductAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/productAds"

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
	localVarPostBody = r.updateProductAd
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
