package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ProductTargetingAPIService ProductTargetingAPI service
type ProductTargetingAPIService service

type ApiGetCategoryRecommendationsForASINsRequest struct {
	ctx                                       context.Context
	ApiService                                *ProductTargetingAPIService
	amazonAdvertisingAPIClientId              *string
	amazonAdvertisingAPIScope                 *string
	prefer                                    *string
	locale                                    *string
	getCategoryRecommendationsForAsinsRequest *GetCategoryRecommendationsForAsinsRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCategoryRecommendationsForASINsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetCategoryRecommendationsForASINsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetCategoryRecommendationsForASINsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetCategoryRecommendationsForASINsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Used to indicate the behavior preferred by the client but is not required for successful completion of the request. Supported values will be updated in the future.
func (r ApiGetCategoryRecommendationsForASINsRequest) Prefer(prefer string) ApiGetCategoryRecommendationsForASINsRequest {
	r.prefer = &prefer
	return r
}

// The locale to which the caller wishes to translate the list of category recommendations to. For example, if the caller wishes to receive a list of category recommendations in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned list of category recommendations will be in the default language of the marketplace.
func (r ApiGetCategoryRecommendationsForASINsRequest) Locale(locale string) ApiGetCategoryRecommendationsForASINsRequest {
	r.locale = &locale
	return r
}

func (r ApiGetCategoryRecommendationsForASINsRequest) GetCategoryRecommendationsForAsinsRequest(getCategoryRecommendationsForAsinsRequest GetCategoryRecommendationsForAsinsRequest) ApiGetCategoryRecommendationsForASINsRequest {
	r.getCategoryRecommendationsForAsinsRequest = &getCategoryRecommendationsForAsinsRequest
	return r
}

func (r ApiGetCategoryRecommendationsForASINsRequest) Execute() (*CategoryRecommendations, *http.Response, error) {
	return r.ApiService.GetCategoryRecommendationsForASINsExecute(r)
}

func (a *ProductTargetingAPIService) GetCategoryRecommendationsForASINs(ctx context.Context) ApiGetCategoryRecommendationsForASINsRequest {
	return ApiGetCategoryRecommendationsForASINsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CategoryRecommendations
func (a *ProductTargetingAPIService) GetCategoryRecommendationsForASINsExecute(r ApiGetCategoryRecommendationsForASINsRequest) (*CategoryRecommendations, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CategoryRecommendations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingAPIService.GetCategoryRecommendationsForASINs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/categories/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spproducttargeting.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spproducttargetingresponse.v3+json", "application/vnd.spproducttargetingresponse.v4+json", "application/vnd.spproducttargetingresponse.v5+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.getCategoryRecommendationsForAsinsRequest
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

type ApiGetNegativeBrandsRequest struct {
	ctx                          context.Context
	ApiService                   *ProductTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	prefer                       *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetNegativeBrandsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetNegativeBrandsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetNegativeBrandsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetNegativeBrandsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Used to indicate the behavior preferred by the client but is not required for successful completion of the request. Supported values will be updated in the future.
func (r ApiGetNegativeBrandsRequest) Prefer(prefer string) ApiGetNegativeBrandsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiGetNegativeBrandsRequest) Execute() ([]Brand, *http.Response, error) {
	return r.ApiService.GetNegativeBrandsExecute(r)
}

func (a *ProductTargetingAPIService) GetNegativeBrands(ctx context.Context) ApiGetNegativeBrandsRequest {
	return ApiGetNegativeBrandsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Brand
func (a *ProductTargetingAPIService) GetNegativeBrandsExecute(r ApiGetNegativeBrandsRequest) ([]Brand, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Brand
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingAPIService.GetNegativeBrands")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeTargets/brands/recommendations"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.spproducttargetingresponse.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
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

type ApiGetRefinementsForCategoryRequest struct {
	ctx                          context.Context
	ApiService                   *ProductTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	categoryId                   string
	prefer                       *string
	locale                       *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetRefinementsForCategoryRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetRefinementsForCategoryRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetRefinementsForCategoryRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetRefinementsForCategoryRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Used to indicate the behavior preferred by the client but is not required for successful completion of the request. Supported values will be updated in the future.
func (r ApiGetRefinementsForCategoryRequest) Prefer(prefer string) ApiGetRefinementsForCategoryRequest {
	r.prefer = &prefer
	return r
}

// The locale to which the caller wishes to translate the refinements to. For example, if the caller wishes to receive the refinements in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the refinements will be in the default language of the marketplace.
func (r ApiGetRefinementsForCategoryRequest) Locale(locale string) ApiGetRefinementsForCategoryRequest {
	r.locale = &locale
	return r
}

func (r ApiGetRefinementsForCategoryRequest) Execute() (*Refinements, *http.Response, error) {
	return r.ApiService.GetRefinementsForCategoryExecute(r)
}

func (a *ProductTargetingAPIService) GetRefinementsForCategory(ctx context.Context, categoryId string) ApiGetRefinementsForCategoryRequest {
	return ApiGetRefinementsForCategoryRequest{
		ApiService: a,
		ctx:        ctx,
		categoryId: categoryId,
	}
}

// Execute executes the request
//
//	@return Refinements
func (a *ProductTargetingAPIService) GetRefinementsForCategoryExecute(r ApiGetRefinementsForCategoryRequest) (*Refinements, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Refinements
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingAPIService.GetRefinementsForCategory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/category/{categoryId}/refinements"
	localVarPath = strings.Replace(localVarPath, "{"+"categoryId"+"}", url.PathEscape(parameterValueToString(r.categoryId, "categoryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spproducttargetingresponse.v3+json", "application/vnd.spproducttargetingresponse.v4+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
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

type ApiGetTargetableASINCountsRequest struct {
	ctx                            context.Context
	ApiService                     *ProductTargetingAPIService
	amazonAdvertisingAPIClientId   *string
	amazonAdvertisingAPIScope      *string
	prefer                         *string
	getTargetableAsinCountsRequest *GetTargetableAsinCountsRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetTargetableASINCountsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetTargetableASINCountsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetTargetableASINCountsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetTargetableASINCountsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Used to indicate the behavior preferred by the client but is not required for successful completion of the request. Supported values will be updated in the future.
func (r ApiGetTargetableASINCountsRequest) Prefer(prefer string) ApiGetTargetableASINCountsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiGetTargetableASINCountsRequest) GetTargetableAsinCountsRequest(getTargetableAsinCountsRequest GetTargetableAsinCountsRequest) ApiGetTargetableASINCountsRequest {
	r.getTargetableAsinCountsRequest = &getTargetableAsinCountsRequest
	return r
}

func (r ApiGetTargetableASINCountsRequest) Execute() (*TargetableAsinCounts, *http.Response, error) {
	return r.ApiService.GetTargetableASINCountsExecute(r)
}

func (a *ProductTargetingAPIService) GetTargetableASINCounts(ctx context.Context) ApiGetTargetableASINCountsRequest {
	return ApiGetTargetableASINCountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TargetableAsinCounts
func (a *ProductTargetingAPIService) GetTargetableASINCountsExecute(r ApiGetTargetableASINCountsRequest) (*TargetableAsinCounts, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TargetableAsinCounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingAPIService.GetTargetableASINCounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/products/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spproducttargeting.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spproducttargetingresponse.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.getTargetableAsinCountsRequest
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

type ApiGetTargetableCategoriesRequest struct {
	ctx                          context.Context
	ApiService                   *ProductTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	prefer                       *string
	locale                       *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetTargetableCategoriesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetTargetableCategoriesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetTargetableCategoriesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetTargetableCategoriesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Used to indicate the behavior preferred by the client but is not required for successful completion of the request. Supported values will be updated in the future.
func (r ApiGetTargetableCategoriesRequest) Prefer(prefer string) ApiGetTargetableCategoriesRequest {
	r.prefer = &prefer
	return r
}

// The locale to which the caller wishes to translate the targetable categories to. For example, if the caller wishes to receive the targetable categories in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned targetable categories will be in the default language of the marketplace.
func (r ApiGetTargetableCategoriesRequest) Locale(locale string) ApiGetTargetableCategoriesRequest {
	r.locale = &locale
	return r
}

func (r ApiGetTargetableCategoriesRequest) Execute() (*TargetableCategories, *http.Response, error) {
	return r.ApiService.GetTargetableCategoriesExecute(r)
}

func (a *ProductTargetingAPIService) GetTargetableCategories(ctx context.Context) ApiGetTargetableCategoriesRequest {
	return ApiGetTargetableCategoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TargetableCategories
func (a *ProductTargetingAPIService) GetTargetableCategoriesExecute(r ApiGetTargetableCategoriesRequest) (*TargetableCategories, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TargetableCategories
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingAPIService.GetTargetableCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spproducttargetingresponse.v3+json", "application/vnd.spproducttargetingresponse.v4+json", "application/vnd.spproducttargetingresponse.v5+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
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

type ApiSearchBrandsRequest struct {
	ctx                          context.Context
	ApiService                   *ProductTargetingAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	prefer                       *string
	searchBrandsRequest          *SearchBrandsRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiSearchBrandsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSearchBrandsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiSearchBrandsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSearchBrandsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Used to indicate the behavior preferred by the client but is not required for successful completion of the request. Supported values will be updated in the future.
func (r ApiSearchBrandsRequest) Prefer(prefer string) ApiSearchBrandsRequest {
	r.prefer = &prefer
	return r
}

func (r ApiSearchBrandsRequest) SearchBrandsRequest(searchBrandsRequest SearchBrandsRequest) ApiSearchBrandsRequest {
	r.searchBrandsRequest = &searchBrandsRequest
	return r
}

func (r ApiSearchBrandsRequest) Execute() ([]Brand, *http.Response, error) {
	return r.ApiService.SearchBrandsExecute(r)
}

func (a *ProductTargetingAPIService) SearchBrands(ctx context.Context) ApiSearchBrandsRequest {
	return ApiSearchBrandsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Brand
func (a *ProductTargetingAPIService) SearchBrandsExecute(r ApiSearchBrandsRequest) ([]Brand, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Brand
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTargetingAPIService.SearchBrands")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/negativeTargets/brands/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spproducttargeting.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spproducttargetingresponse.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.prefer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Prefer", r.prefer, "simple", "")
	}
	// body params
	localVarPostBody = r.searchBrandsRequest
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
