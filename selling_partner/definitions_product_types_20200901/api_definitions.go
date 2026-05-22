package definitions_product_types_20200901

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefinitionsAPIService DefinitionsAPI service
type DefinitionsAPIService service

type ApiGetDefinitionsProductTypeRequest struct {
	ctx                  context.Context
	ApiService           *DefinitionsAPIService
	productType          string
	marketplaceIds       *[]string
	sellerId             *string
	productTypeVersion   *string
	requirements         *string
	requirementsEnforced *string
	locale               *string
}

// A comma-delimited list of Amazon marketplace identifiers for the request. Note: This parameter is limited to one marketplaceId at this time.
func (r ApiGetDefinitionsProductTypeRequest) MarketplaceIds(marketplaceIds []string) ApiGetDefinitionsProductTypeRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// A selling partner identifier. When provided, seller-specific requirements and values are populated within the product type definition schema, such as brand names associated with the selling partner.
func (r ApiGetDefinitionsProductTypeRequest) SellerId(sellerId string) ApiGetDefinitionsProductTypeRequest {
	r.sellerId = &sellerId
	return r
}

// The version of the Amazon product type to retrieve. Defaults to \&quot;LATEST\&quot;,. Prerelease versions of product type definitions may be retrieved with \&quot;RELEASE_CANDIDATE\&quot;. If no prerelease version is currently available, the \&quot;LATEST\&quot; live version will be provided.
func (r ApiGetDefinitionsProductTypeRequest) ProductTypeVersion(productTypeVersion string) ApiGetDefinitionsProductTypeRequest {
	r.productTypeVersion = &productTypeVersion
	return r
}

// The name of the requirements set to retrieve requirements for.
func (r ApiGetDefinitionsProductTypeRequest) Requirements(requirements string) ApiGetDefinitionsProductTypeRequest {
	r.requirements = &requirements
	return r
}

// Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all the required attributes being present (such as for partial updates).
func (r ApiGetDefinitionsProductTypeRequest) RequirementsEnforced(requirementsEnforced string) ApiGetDefinitionsProductTypeRequest {
	r.requirementsEnforced = &requirementsEnforced
	return r
}

// Locale for retrieving display labels and other presentation details. Defaults to the default language of the first marketplace in the request.
func (r ApiGetDefinitionsProductTypeRequest) Locale(locale string) ApiGetDefinitionsProductTypeRequest {
	r.locale = &locale
	return r
}

func (r ApiGetDefinitionsProductTypeRequest) Execute() (*ProductTypeDefinition, *http.Response, error) {
	return r.ApiService.GetDefinitionsProductTypeExecute(r)
}

/*
GetDefinitionsProductType Method for GetDefinitionsProductType

Retrieve an Amazon product type definition.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 5 | 10 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param productType The Amazon product type name.
	@return ApiGetDefinitionsProductTypeRequest
*/
func (a *DefinitionsAPIService) GetDefinitionsProductType(ctx context.Context, productType string) ApiGetDefinitionsProductTypeRequest {
	return ApiGetDefinitionsProductTypeRequest{
		ApiService:  a,
		ctx:         ctx,
		productType: productType,
	}
}

// Execute executes the request
//
//	@return ProductTypeDefinition
func (a *DefinitionsAPIService) GetDefinitionsProductTypeExecute(r ApiGetDefinitionsProductTypeRequest) (*ProductTypeDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductTypeDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsAPIService.GetDefinitionsProductType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/definitions/2020-09-01/productTypes/{productType}"
	localVarPath = strings.Replace(localVarPath, "{"+"productType"+"}", url.PathEscape(parameterValueToString(r.productType, "productType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}

	if r.sellerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sellerId", r.sellerId, "", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	if r.productTypeVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productTypeVersion", r.productTypeVersion, "", "")
	}
	if r.requirements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "requirements", r.requirements, "", "")
	}
	if r.requirementsEnforced != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "requirementsEnforced", r.requirementsEnforced, "", "")
	}
	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "", "")
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

type ApiSearchDefinitionsProductTypesRequest struct {
	ctx            context.Context
	ApiService     *DefinitionsAPIService
	marketplaceIds *[]string
	keywords       *[]string
	itemName       *string
	locale         *string
	searchLocale   *string
}

// A comma-delimited list of Amazon marketplace identifiers for the request.
func (r ApiSearchDefinitionsProductTypesRequest) MarketplaceIds(marketplaceIds []string) ApiSearchDefinitionsProductTypesRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// A comma-delimited list of keywords to search product types. **Note:** Cannot be used with &#x60;itemName&#x60;.
func (r ApiSearchDefinitionsProductTypesRequest) Keywords(keywords []string) ApiSearchDefinitionsProductTypesRequest {
	r.keywords = &keywords
	return r
}

// The title of the ASIN to get the product type recommendation. **Note:** Cannot be used with &#x60;keywords&#x60;.
func (r ApiSearchDefinitionsProductTypesRequest) ItemName(itemName string) ApiSearchDefinitionsProductTypesRequest {
	r.itemName = &itemName
	return r
}

// The locale for the display names in the response. Defaults to the primary locale of the marketplace.
func (r ApiSearchDefinitionsProductTypesRequest) Locale(locale string) ApiSearchDefinitionsProductTypesRequest {
	r.locale = &locale
	return r
}

// The locale used for the &#x60;keywords&#x60; and &#x60;itemName&#x60; parameters. Defaults to the primary locale of the marketplace.
func (r ApiSearchDefinitionsProductTypesRequest) SearchLocale(searchLocale string) ApiSearchDefinitionsProductTypesRequest {
	r.searchLocale = &searchLocale
	return r
}

func (r ApiSearchDefinitionsProductTypesRequest) Execute() (*ProductTypeList, *http.Response, error) {
	return r.ApiService.SearchDefinitionsProductTypesExecute(r)
}

/*
SearchDefinitionsProductTypes Method for SearchDefinitionsProductTypes

Search for and return a list of Amazon product types that have definitions available.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 5 | 10 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchDefinitionsProductTypesRequest
*/
func (a *DefinitionsAPIService) SearchDefinitionsProductTypes(ctx context.Context) ApiSearchDefinitionsProductTypesRequest {
	return ApiSearchDefinitionsProductTypesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductTypeList
func (a *DefinitionsAPIService) SearchDefinitionsProductTypesExecute(r ApiSearchDefinitionsProductTypesRequest) (*ProductTypeList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductTypeList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsAPIService.SearchDefinitionsProductTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/definitions/2020-09-01/productTypes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceIds == nil {
		return localVarReturnValue, nil, reportError("marketplaceIds is required and must be specified")
	}

	if r.keywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keywords", r.keywords, "form", "csv")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceIds", r.marketplaceIds, "form", "csv")
	if r.itemName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemName", r.itemName, "", "")
	}
	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "", "")
	}
	if r.searchLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchLocale", r.searchLocale, "", "")
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
