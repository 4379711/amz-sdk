package listings_items_20210801

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ListingsAPIService ListingsAPI service
type ListingsAPIService service

type ApiDeleteListingsItemRequest struct {
	ctx            context.Context
	ApiService     *ListingsAPIService
	sellerId       string
	sku            string
	marketplaceIds *[]string
	issueLocale    *string
}

// A comma-delimited list of Amazon marketplace identifiers for the request.
func (r ApiDeleteListingsItemRequest) MarketplaceIds(marketplaceIds []string) ApiDeleteListingsItemRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when a localization is not available in the specified locale.
func (r ApiDeleteListingsItemRequest) IssueLocale(issueLocale string) ApiDeleteListingsItemRequest {
	r.issueLocale = &issueLocale
	return r
}

func (r ApiDeleteListingsItemRequest) Execute() (*ListingsItemSubmissionResponse, *http.Response, error) {
	return r.ApiService.DeleteListingsItemExecute(r)
}

/*
DeleteListingsItem Method for DeleteListingsItem

Delete a listings item for a selling partner.

**Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sellerId A selling partner identifier, such as a merchant account or vendor code.
	@param sku A selling partner provided identifier for an Amazon listing.
	@return ApiDeleteListingsItemRequest
*/
func (a *ListingsAPIService) DeleteListingsItem(ctx context.Context, sellerId string, sku string) ApiDeleteListingsItemRequest {
	return ApiDeleteListingsItemRequest{
		ApiService: a,
		ctx:        ctx,
		sellerId:   sellerId,
		sku:        sku,
	}
}

// Execute executes the request
//
//	@return ListingsItemSubmissionResponse
func (a *ListingsAPIService) DeleteListingsItemExecute(r ApiDeleteListingsItemRequest) (*ListingsItemSubmissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListingsItemSubmissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListingsAPIService.DeleteListingsItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listings/2021-08-01/items/{sellerId}/{sku}"
	localVarPath = strings.Replace(localVarPath, "{"+"sellerId"+"}", url.PathEscape(parameterValueToString(r.sellerId, "sellerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	if r.issueLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issueLocale", r.issueLocale, "", "")
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

type ApiGetListingsItemRequest struct {
	ctx            context.Context
	ApiService     *ListingsAPIService
	sellerId       string
	sku            string
	marketplaceIds *[]string
	issueLocale    *string
	includedData   *[]string
}

// A comma-delimited list of Amazon marketplace identifiers for the request.
func (r ApiGetListingsItemRequest) MarketplaceIds(marketplaceIds []string) ApiGetListingsItemRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when a localization is not available in the specified locale.
func (r ApiGetListingsItemRequest) IssueLocale(issueLocale string) ApiGetListingsItemRequest {
	r.issueLocale = &issueLocale
	return r
}

// A comma-delimited list of data sets to include in the response. Default: &#x60;summaries&#x60;.
func (r ApiGetListingsItemRequest) IncludedData(includedData []string) ApiGetListingsItemRequest {
	r.includedData = &includedData
	return r
}

func (r ApiGetListingsItemRequest) Execute() (*Item, *http.Response, error) {
	return r.ApiService.GetListingsItemExecute(r)
}

/*
GetListingsItem Method for GetListingsItem

Returns details about a listings item for a selling partner.

**Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sellerId A selling partner identifier, such as a merchant account or vendor code.
	@param sku A selling partner provided identifier for an Amazon listing.
	@return ApiGetListingsItemRequest
*/
func (a *ListingsAPIService) GetListingsItem(ctx context.Context, sellerId string, sku string) ApiGetListingsItemRequest {
	return ApiGetListingsItemRequest{
		ApiService: a,
		ctx:        ctx,
		sellerId:   sellerId,
		sku:        sku,
	}
}

// Execute executes the request
//
//	@return Item
func (a *ListingsAPIService) GetListingsItemExecute(r ApiGetListingsItemRequest) (*Item, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Item
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListingsAPIService.GetListingsItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listings/2021-08-01/items/{sellerId}/{sku}"
	localVarPath = strings.Replace(localVarPath, "{"+"sellerId"+"}", url.PathEscape(parameterValueToString(r.sellerId, "sellerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	if r.issueLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issueLocale", r.issueLocale, "", "")
	}
	if r.includedData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includedData", r.includedData, "form", "csv")
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

type ApiPatchListingsItemRequest struct {
	ctx            context.Context
	ApiService     *ListingsAPIService
	body           *ListingsItemPatchRequest
	sellerId       string
	sku            string
	marketplaceIds *[]string
	includedData   *[]string
	mode           *string
	issueLocale    *string
}

// The request body schema for the &#x60;patchListingsItem&#x60; operation.
func (r ApiPatchListingsItemRequest) Body(body ListingsItemPatchRequest) ApiPatchListingsItemRequest {
	r.body = &body
	return r
}

// A comma-delimited list of Amazon marketplace identifiers for the request.
func (r ApiPatchListingsItemRequest) MarketplaceIds(marketplaceIds []string) ApiPatchListingsItemRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// A comma-delimited list of data sets to include in the response. Default: &#x60;issues&#x60;.
func (r ApiPatchListingsItemRequest) IncludedData(includedData []string) ApiPatchListingsItemRequest {
	r.includedData = &includedData
	return r
}

// The mode of operation for the request.
func (r ApiPatchListingsItemRequest) Mode(mode string) ApiPatchListingsItemRequest {
	r.mode = &mode
	return r
}

// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when a localization is not available in the specified locale.
func (r ApiPatchListingsItemRequest) IssueLocale(issueLocale string) ApiPatchListingsItemRequest {
	r.issueLocale = &issueLocale
	return r
}

func (r ApiPatchListingsItemRequest) Execute() (*ListingsItemSubmissionResponse, *http.Response, error) {
	return r.ApiService.PatchListingsItemExecute(r)
}

/*
PatchListingsItem Method for PatchListingsItem

Partially update (patch) a listings item for a selling partner. Only top-level listings item attributes can be patched. Patching nested attributes is not supported.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput can receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api) in the Selling Partner API documentation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sellerId A selling partner identifier, such as a merchant account or vendor code.
	@param sku A selling partner provided identifier for an Amazon listing.
	@return ApiPatchListingsItemRequest
*/
func (a *ListingsAPIService) PatchListingsItem(ctx context.Context, sellerId string, sku string) ApiPatchListingsItemRequest {
	return ApiPatchListingsItemRequest{
		ApiService: a,
		ctx:        ctx,
		sellerId:   sellerId,
		sku:        sku,
	}
}

// Execute executes the request
//
//	@return ListingsItemSubmissionResponse
func (a *ListingsAPIService) PatchListingsItemExecute(r ApiPatchListingsItemRequest) (*ListingsItemSubmissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListingsItemSubmissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListingsAPIService.PatchListingsItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listings/2021-08-01/items/{sellerId}/{sku}"
	localVarPath = strings.Replace(localVarPath, "{"+"sellerId"+"}", url.PathEscape(parameterValueToString(r.sellerId, "sellerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	if r.includedData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includedData", r.includedData, "form", "csv")
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "", "")
	}
	if r.issueLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issueLocale", r.issueLocale, "", "")
	}
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
	localVarPostBody = r.body
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

type ApiPutListingsItemRequest struct {
	ctx            context.Context
	ApiService     *ListingsAPIService
	body           *ListingsItemPutRequest
	sellerId       string
	sku            string
	marketplaceIds *[]string
	includedData   *[]string
	mode           *string
	issueLocale    *string
}

// The request body schema for the &#x60;putListingsItem&#x60; operation.
func (r ApiPutListingsItemRequest) Body(body ListingsItemPutRequest) ApiPutListingsItemRequest {
	r.body = &body
	return r
}

// A comma-delimited list of Amazon marketplace identifiers for the request.
func (r ApiPutListingsItemRequest) MarketplaceIds(marketplaceIds []string) ApiPutListingsItemRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// A comma-delimited list of data sets to include in the response. Default: &#x60;issues&#x60;.
func (r ApiPutListingsItemRequest) IncludedData(includedData []string) ApiPutListingsItemRequest {
	r.includedData = &includedData
	return r
}

// The mode of operation for the request.
func (r ApiPutListingsItemRequest) Mode(mode string) ApiPutListingsItemRequest {
	r.mode = &mode
	return r
}

// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: &#x60;en_US&#x60;, &#x60;fr_CA&#x60;, &#x60;fr_FR&#x60;. Localized messages default to &#x60;en_US&#x60; when a localization is not available in the specified locale.
func (r ApiPutListingsItemRequest) IssueLocale(issueLocale string) ApiPutListingsItemRequest {
	r.issueLocale = &issueLocale
	return r
}

func (r ApiPutListingsItemRequest) Execute() (*ListingsItemSubmissionResponse, *http.Response, error) {
	return r.ApiService.PutListingsItemExecute(r)
}

/*
PutListingsItem Method for PutListingsItem

Creates or fully updates an existing listings item for a selling partner.

**Note:** This operation has a throttling rate of one request per second when `mode` is `VALIDATION_PREVIEW`.

**Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sellerId A selling partner identifier, such as a merchant account or vendor code.
	@param sku A selling partner provided identifier for an Amazon listing.
	@return ApiPutListingsItemRequest
*/
func (a *ListingsAPIService) PutListingsItem(ctx context.Context, sellerId string, sku string) ApiPutListingsItemRequest {
	return ApiPutListingsItemRequest{
		ApiService: a,
		ctx:        ctx,
		sellerId:   sellerId,
		sku:        sku,
	}
}

// Execute executes the request
//
//	@return ListingsItemSubmissionResponse
func (a *ListingsAPIService) PutListingsItemExecute(r ApiPutListingsItemRequest) (*ListingsItemSubmissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListingsItemSubmissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListingsAPIService.PutListingsItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listings/2021-08-01/items/{sellerId}/{sku}"
	localVarPath = strings.Replace(localVarPath, "{"+"sellerId"+"}", url.PathEscape(parameterValueToString(r.sellerId, "sellerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	if r.includedData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includedData", r.includedData, "form", "csv")
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "", "")
	}
	if r.issueLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issueLocale", r.issueLocale, "", "")
	}
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
	localVarPostBody = r.body
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

type ApiSearchListingsItemsRequest struct {
	ctx                 context.Context
	ApiService          *ListingsAPIService
	sellerId            string
	marketplaceIds      *[]string
	issueLocale         *string
	includedData        *[]string
	identifiers         *[]string
	identifiersType     *string
	variationParentSku  *string
	packageHierarchySku *string
	createdAfter        *time.Time
	createdBefore       *time.Time
	lastUpdatedAfter    *time.Time
	lastUpdatedBefore   *time.Time
	withIssueSeverity   *[]string
	withStatus          *[]string
	withoutStatus       *[]string
	sortBy              *string
	sortOrder           *string
	pageSize            *int32
	pageToken           *string
}

// A comma-delimited list of Amazon marketplace identifiers for the request.
func (r ApiSearchListingsItemsRequest) MarketplaceIds(marketplaceIds []string) ApiSearchListingsItemsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// A locale that is used to localize issues. When not provided, the default language code of the first marketplace is used. Examples: \&quot;en_US\&quot;, \&quot;fr_CA\&quot;, \&quot;fr_FR\&quot;. When a localization is not available in the specified locale, localized messages default to \&quot;en_US\&quot;.
func (r ApiSearchListingsItemsRequest) IssueLocale(issueLocale string) ApiSearchListingsItemsRequest {
	r.issueLocale = &issueLocale
	return r
}

// A comma-delimited list of datasets that you want to include in the response. Default: &#x60;summaries&#x60;.
func (r ApiSearchListingsItemsRequest) IncludedData(includedData []string) ApiSearchListingsItemsRequest {
	r.includedData = &includedData
	return r
}

// A comma-delimited list of product identifiers that you can use to search for listings items.   **Note**:  1. This is required when you specify &#x60;identifiersType&#x60;. 2. You cannot use &#39;identifiers&#39; if you specify &#x60;variationParentSku&#x60; or &#x60;packageHierarchySku&#x60;.
func (r ApiSearchListingsItemsRequest) Identifiers(identifiers []string) ApiSearchListingsItemsRequest {
	r.identifiers = &identifiers
	return r
}

// A type of product identifiers that you can use to search for listings items.   **Note**:  This is required when &#x60;identifiers&#x60; is provided.
func (r ApiSearchListingsItemsRequest) IdentifiersType(identifiersType string) ApiSearchListingsItemsRequest {
	r.identifiersType = &identifiersType
	return r
}

// Filters results to include listing items that are variation children of the specified SKU.   **Note**: You cannot use &#x60;variationParentSku&#x60; if you include &#x60;identifiers&#x60; or &#x60;packageHierarchySku&#x60; in your request.
func (r ApiSearchListingsItemsRequest) VariationParentSku(variationParentSku string) ApiSearchListingsItemsRequest {
	r.variationParentSku = &variationParentSku
	return r
}

// Filter results to include listing items that contain or are contained by the specified SKU.   **Note**: You cannot use &#x60;packageHierarchySku&#x60; if you include &#x60;identifiers&#x60; or &#x60;variationParentSku&#x60; in your request.
func (r ApiSearchListingsItemsRequest) PackageHierarchySku(packageHierarchySku string) ApiSearchListingsItemsRequest {
	r.packageHierarchySku = &packageHierarchySku
	return r
}

// A date-time that is used to filter listing items. The response includes listings items that were created at or after this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format.
func (r ApiSearchListingsItemsRequest) CreatedAfter(createdAfter time.Time) ApiSearchListingsItemsRequest {
	r.createdAfter = &createdAfter
	return r
}

// A date-time that is used to filter listing items. The response includes listings items that were created at or before this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format.
func (r ApiSearchListingsItemsRequest) CreatedBefore(createdBefore time.Time) ApiSearchListingsItemsRequest {
	r.createdBefore = &createdBefore
	return r
}

// A date-time that is used to filter listing items. The response includes listings items that were last updated at or after this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format.
func (r ApiSearchListingsItemsRequest) LastUpdatedAfter(lastUpdatedAfter time.Time) ApiSearchListingsItemsRequest {
	r.lastUpdatedAfter = &lastUpdatedAfter
	return r
}

// A date-time that is used to filter listing items. The response includes listings items that were last updated at or before this time. Values are in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format.
func (r ApiSearchListingsItemsRequest) LastUpdatedBefore(lastUpdatedBefore time.Time) ApiSearchListingsItemsRequest {
	r.lastUpdatedBefore = &lastUpdatedBefore
	return r
}

// Filter results to include only listing items that have issues that match one or more of the specified severity levels.
func (r ApiSearchListingsItemsRequest) WithIssueSeverity(withIssueSeverity []string) ApiSearchListingsItemsRequest {
	r.withIssueSeverity = &withIssueSeverity
	return r
}

// Filter results to include only listing items that have the specified status.
func (r ApiSearchListingsItemsRequest) WithStatus(withStatus []string) ApiSearchListingsItemsRequest {
	r.withStatus = &withStatus
	return r
}

// Filter results to include only listing items that don&#39;t contain the specified statuses.
func (r ApiSearchListingsItemsRequest) WithoutStatus(withoutStatus []string) ApiSearchListingsItemsRequest {
	r.withoutStatus = &withoutStatus
	return r
}

// An attribute by which to sort the returned listing items.
func (r ApiSearchListingsItemsRequest) SortBy(sortBy string) ApiSearchListingsItemsRequest {
	r.sortBy = &sortBy
	return r
}

// The order in which to sort the result items.
func (r ApiSearchListingsItemsRequest) SortOrder(sortOrder string) ApiSearchListingsItemsRequest {
	r.sortOrder = &sortOrder
	return r
}

// The number of results that you want to include on each page.
func (r ApiSearchListingsItemsRequest) PageSize(pageSize int32) ApiSearchListingsItemsRequest {
	r.pageSize = &pageSize
	return r
}

// A token that you can use to fetch a specific page when there are multiple pages of results.
func (r ApiSearchListingsItemsRequest) PageToken(pageToken string) ApiSearchListingsItemsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiSearchListingsItemsRequest) Execute() (*ItemSearchResults, *http.Response, error) {
	return r.ApiService.SearchListingsItemsExecute(r)
}

/*
SearchListingsItems Method for SearchListingsItems

Search for and return a list of selling partner listings items and their respective details.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that are applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput might have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sellerId A selling partner identifier, such as a merchant account or vendor code.
	@return ApiSearchListingsItemsRequest
*/
func (a *ListingsAPIService) SearchListingsItems(ctx context.Context, sellerId string) ApiSearchListingsItemsRequest {
	return ApiSearchListingsItemsRequest{
		ApiService: a,
		ctx:        ctx,
		sellerId:   sellerId,
	}
}

// Execute executes the request
//
//	@return ItemSearchResults
func (a *ListingsAPIService) SearchListingsItemsExecute(r ApiSearchListingsItemsRequest) (*ItemSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ItemSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListingsAPIService.SearchListingsItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listings/2021-08-01/items/{sellerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sellerId"+"}", url.PathEscape(parameterValueToString(r.sellerId, "sellerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}
	if len(*r.marketplaceIds) > 1 {
		return localVarReturnValue, nil, reportError("marketplaceIds must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	if r.issueLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issueLocale", r.issueLocale, "", "")
	}
	if r.includedData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includedData", r.includedData, "form", "csv")
	}
	if r.identifiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identifiers", r.identifiers, "form", "csv")
	}
	if r.identifiersType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identifiersType", r.identifiersType, "", "")
	}
	if r.variationParentSku != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "variationParentSku", r.variationParentSku, "", "")
	}
	if r.packageHierarchySku != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageHierarchySku", r.packageHierarchySku, "", "")
	}
	if r.createdAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdAfter", r.createdAfter, "", "")
	}
	if r.createdBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdBefore", r.createdBefore, "", "")
	}
	if r.lastUpdatedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastUpdatedAfter", r.lastUpdatedAfter, "", "")
	}
	if r.lastUpdatedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastUpdatedBefore", r.lastUpdatedBefore, "", "")
	}
	if r.withIssueSeverity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withIssueSeverity", r.withIssueSeverity, "form", "csv")
	}
	if r.withStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withStatus", r.withStatus, "form", "csv")
	}
	if r.withoutStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withoutStatus", r.withoutStatus, "form", "csv")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "", "")
	}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", r.sortOrder, "", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "", "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "", "")
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
