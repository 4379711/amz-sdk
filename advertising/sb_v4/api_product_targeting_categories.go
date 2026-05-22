package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ProductTargetingCategoriesAPIService ProductTargetingCategoriesAPI service
type ProductTargetingCategoriesAPIService service

type ApiSBTargetingGetRefinementsForCategoryRequest struct {
	ctx                          context.Context
	ApiService                   *ProductTargetingCategoriesAPIService
	categoryRefinementId         string
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	locale                       *SBTargetingLocale
	nextToken                    *string
}

// The identifier of a client associated with a &#x60;Login with Amazon&#x60; account.
func (r ApiSBTargetingGetRefinementsForCategoryRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSBTargetingGetRefinementsForCategoryRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiSBTargetingGetRefinementsForCategoryRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSBTargetingGetRefinementsForCategoryRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The locale to which the caller wishes to translate the targetable categories or refinements to. For example, if the caller wishes to receive the targetable categories in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned tagetable categories will be in the default language of the marketplace.
func (r ApiSBTargetingGetRefinementsForCategoryRequest) Locale(locale SBTargetingLocale) ApiSBTargetingGetRefinementsForCategoryRequest {
	r.locale = &locale
	return r
}

// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results.
func (r ApiSBTargetingGetRefinementsForCategoryRequest) NextToken(nextToken string) ApiSBTargetingGetRefinementsForCategoryRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiSBTargetingGetRefinementsForCategoryRequest) Execute() (*SBTargetingGetRefinementsForCategoryResponseContent, *http.Response, error) {
	return r.ApiService.SBTargetingGetRefinementsForCategoryExecute(r)
}

func (a *ProductTargetingCategoriesAPIService) SBTargetingGetRefinementsForCategory(ctx context.Context, categoryRefinementId string) ApiSBTargetingGetRefinementsForCategoryRequest {
	return ApiSBTargetingGetRefinementsForCategoryRequest{
		ApiService:           a,
		ctx:                  ctx,
		categoryRefinementId: categoryRefinementId,
	}
}

// Execute executes the request
//
//	@return SBTargetingGetRefinementsForCategoryResponseContent
func (a *ProductTargetingCategoriesAPIService) SBTargetingGetRefinementsForCategoryExecute(r ApiSBTargetingGetRefinementsForCategoryRequest) (*SBTargetingGetRefinementsForCategoryResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBTargetingGetRefinementsForCategoryResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingCategoriesAPIService.SBTargetingGetRefinementsForCategory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/targets/categories/{categoryRefinementId}/refinements"
	localVarPath = strings.Replace(localVarPath, "{"+"categoryRefinementId"+"}", url.PathEscape(parameterValueToString(r.categoryRefinementId, "categoryRefinementId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "form", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbtargeting.v4+json"}

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

type ApiSBTargetingGetTargetableASINCountsRequest struct {
	ctx                                              context.Context
	ApiService                                       *ProductTargetingCategoriesAPIService
	amazonAdvertisingAPIClientId                     *string
	amazonAdvertisingAPIScope                        *string
	sBTargetingGetTargetableASINCountsRequestContent *SBTargetingGetTargetableASINCountsRequestContent
}

// The identifier of a client associated with a &#x60;Login with Amazon&#x60; account.
func (r ApiSBTargetingGetTargetableASINCountsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSBTargetingGetTargetableASINCountsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiSBTargetingGetTargetableASINCountsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSBTargetingGetTargetableASINCountsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSBTargetingGetTargetableASINCountsRequest) SBTargetingGetTargetableASINCountsRequestContent(sBTargetingGetTargetableASINCountsRequestContent SBTargetingGetTargetableASINCountsRequestContent) ApiSBTargetingGetTargetableASINCountsRequest {
	r.sBTargetingGetTargetableASINCountsRequestContent = &sBTargetingGetTargetableASINCountsRequestContent
	return r
}

func (r ApiSBTargetingGetTargetableASINCountsRequest) Execute() (*SBTargetingGetTargetableASINCountsResponseContent, *http.Response, error) {
	return r.ApiService.SBTargetingGetTargetableASINCountsExecute(r)
}

func (a *ProductTargetingCategoriesAPIService) SBTargetingGetTargetableASINCounts(ctx context.Context) ApiSBTargetingGetTargetableASINCountsRequest {
	return ApiSBTargetingGetTargetableASINCountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SBTargetingGetTargetableASINCountsResponseContent
func (a *ProductTargetingCategoriesAPIService) SBTargetingGetTargetableASINCountsExecute(r ApiSBTargetingGetTargetableASINCountsRequest) (*SBTargetingGetTargetableASINCountsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBTargetingGetTargetableASINCountsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingCategoriesAPIService.SBTargetingGetTargetableASINCounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/targets/products/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sBTargetingGetTargetableASINCountsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sBTargetingGetTargetableASINCountsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbtargeting.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbtargeting.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sBTargetingGetTargetableASINCountsRequestContent
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

type ApiSBTargetingGetTargetableCategoriesRequest struct {
	ctx                          context.Context
	ApiService                   *ProductTargetingCategoriesAPIService
	supplySource                 *SBTargetingSupplySource
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	locale                       *SBTargetingLocale
	includeOnlyRootCategories    *bool
	parentCategoryRefinementId   *string
	nextToken                    *string
}

// [UPDATE: As of 05/28/2024, &#x60;STREAMING_VIDEO&#x60; is deprecated].   The supply source where the target will be used. Use &#x60;AMAZON&#x60; for placements on Amazon website. Use &#x60;STREAMING_VIDEO&#x60; for off-site video placements such as IMDb TV.
func (r ApiSBTargetingGetTargetableCategoriesRequest) SupplySource(supplySource SBTargetingSupplySource) ApiSBTargetingGetTargetableCategoriesRequest {
	r.supplySource = &supplySource
	return r
}

// The identifier of a client associated with a &#x60;Login with Amazon&#x60; account.
func (r ApiSBTargetingGetTargetableCategoriesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSBTargetingGetTargetableCategoriesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiSBTargetingGetTargetableCategoriesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSBTargetingGetTargetableCategoriesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The locale to which the caller wishes to translate the targetable categories or refinements to. For example, if the caller wishes to receive the targetable categories in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned tagetable categories will be in the default language of the marketplace.
func (r ApiSBTargetingGetTargetableCategoriesRequest) Locale(locale SBTargetingLocale) ApiSBTargetingGetTargetableCategoriesRequest {
	r.locale = &locale
	return r
}

// Indicates whether to only retun root categories or not.
func (r ApiSBTargetingGetTargetableCategoriesRequest) IncludeOnlyRootCategories(includeOnlyRootCategories bool) ApiSBTargetingGetTargetableCategoriesRequest {
	r.includeOnlyRootCategories = &includeOnlyRootCategories
	return r
}

// Returns child categories of category.
func (r ApiSBTargetingGetTargetableCategoriesRequest) ParentCategoryRefinementId(parentCategoryRefinementId string) ApiSBTargetingGetTargetableCategoriesRequest {
	r.parentCategoryRefinementId = &parentCategoryRefinementId
	return r
}

// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results.
func (r ApiSBTargetingGetTargetableCategoriesRequest) NextToken(nextToken string) ApiSBTargetingGetTargetableCategoriesRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiSBTargetingGetTargetableCategoriesRequest) Execute() (*SBTargetingGetTargetableCategoriesResponseContent, *http.Response, error) {
	return r.ApiService.SBTargetingGetTargetableCategoriesExecute(r)
}

func (a *ProductTargetingCategoriesAPIService) SBTargetingGetTargetableCategories(ctx context.Context) ApiSBTargetingGetTargetableCategoriesRequest {
	return ApiSBTargetingGetTargetableCategoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SBTargetingGetTargetableCategoriesResponseContent
func (a *ProductTargetingCategoriesAPIService) SBTargetingGetTargetableCategoriesExecute(r ApiSBTargetingGetTargetableCategoriesRequest) (*SBTargetingGetTargetableCategoriesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBTargetingGetTargetableCategoriesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingCategoriesAPIService.SBTargetingGetTargetableCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/targets/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.supplySource == nil {
		return localVarReturnValue, nil, reportError("supplySource is required and must be specified")
	}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "supplySource", r.supplySource, "form", "")
	if r.includeOnlyRootCategories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeOnlyRootCategories", r.includeOnlyRootCategories, "form", "")
	}
	if r.parentCategoryRefinementId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentCategoryRefinementId", r.parentCategoryRefinementId, "form", "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbtargeting.v4+json"}

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
