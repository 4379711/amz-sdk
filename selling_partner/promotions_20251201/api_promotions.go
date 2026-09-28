package promotions_20251201

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PromotionsAPIService PromotionsAPI service for the Promotions v2025-12-01 API
type PromotionsAPIService service

type ApiSearchPromotionsRequest struct {
	ctx              context.Context
	ApiService       *PromotionsAPIService
	marketplaceIds   *[]string
	locale           *string
	statuses         *[]string
	asins            *[]string
	skus             *[]string
	promotionTypes   *[]string
	startDateBefore  *time.Time
	startDateAfter   *time.Time
	endDateBefore    *time.Time
	endDateAfter     *time.Time
	updateDateAfter  *time.Time
	updateDateBefore *time.Time
	paginationToken  *string
	revision         *string
	limit            *int64
	includedData     *[]string
}

// The Amazon stores from which to retrieve promotions. Refer to [Store Identifiers](https://developer-docs.amazon/sp-api/docs/store-identifiers) for a list of Amazon store values.
func (r ApiSearchPromotionsRequest) MarketplaceIds(marketplaceIds []string) ApiSearchPromotionsRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

// The locale from which to retrieve promotions. Formatted as an ISO 639 language code, followed by an underscore, followed by an ISO 3166-1 alpha-2 country code.
func (r ApiSearchPromotionsRequest) Locale(locale string) ApiSearchPromotionsRequest {
	r.locale = &locale
	return r
}

// The statuses of promotions to retrieve, formatted as a comma-delimited list.
func (r ApiSearchPromotionsRequest) Statuses(statuses []string) ApiSearchPromotionsRequest {
	r.statuses = &statuses
	return r
}

// The ASINs to which promotions apply, formatted as a comma-delimited list.
func (r ApiSearchPromotionsRequest) Asins(asins []string) ApiSearchPromotionsRequest {
	r.asins = &asins
	return r
}

// The SKUs to which promotions apply, formatted as a comma-delimited list.
func (r ApiSearchPromotionsRequest) Skus(skus []string) ApiSearchPromotionsRequest {
	r.skus = &skus
	return r
}

// The promotion types to which promotions apply, formatted as a comma-delimited list.
func (r ApiSearchPromotionsRequest) PromotionTypes(promotionTypes []string) ApiSearchPromotionsRequest {
	r.promotionTypes = &promotionTypes
	return r
}

// Promotions that start before this date are returned. Formatted in ISO 8601 format, including the timezone. For example: &#x60;1970-01-01T00:00:00-07:00&#x60;.
func (r ApiSearchPromotionsRequest) StartDateBefore(startDateBefore time.Time) ApiSearchPromotionsRequest {
	r.startDateBefore = &startDateBefore
	return r
}

// Promotions that start after this date are returned. Formatted in ISO 8601 format, including the timezone. For example: &#x60;1970-01-01T00:00:00-07:00&#x60;.
func (r ApiSearchPromotionsRequest) StartDateAfter(startDateAfter time.Time) ApiSearchPromotionsRequest {
	r.startDateAfter = &startDateAfter
	return r
}

// Promotions that end before this date are returned. Formatted in ISO 8601 format, including the timezone. For example: &#x60;1970-01-01T00:00:00-07:00&#x60;.
func (r ApiSearchPromotionsRequest) EndDateBefore(endDateBefore time.Time) ApiSearchPromotionsRequest {
	r.endDateBefore = &endDateBefore
	return r
}

// Promotions that end after this date are returned. Formatted in ISO 8601 format, including the timezone. For example: &#x60;1970-01-01T00:00:00-07:00&#x60;.
func (r ApiSearchPromotionsRequest) EndDateAfter(endDateAfter time.Time) ApiSearchPromotionsRequest {
	r.endDateAfter = &endDateAfter
	return r
}

// Filter promotions that were last modified after this timestamp. Zoned Datetime in ISO 8601 format (e.g., 1970-01-01T00:00:00-07:00).
func (r ApiSearchPromotionsRequest) UpdateDateAfter(updateDateAfter time.Time) ApiSearchPromotionsRequest {
	r.updateDateAfter = &updateDateAfter
	return r
}

// Filter promotions that were last modified before this timestamp. Zoned Datetime in ISO 8601 format (e.g., 1970-01-01T00:00:00-07:00).
func (r ApiSearchPromotionsRequest) UpdateDateBefore(updateDateBefore time.Time) ApiSearchPromotionsRequest {
	r.updateDateBefore = &updateDateBefore
	return r
}

// A token that you use to retrieve the next page of results. The response includes &#x60;paginationToken&#x60; when the number of results exceeds the specified &#x60;limit&#x60; value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until &#x60;paginationToken&#x60; is null. Note that this operation can return empty pages.
func (r ApiSearchPromotionsRequest) PaginationToken(paginationToken string) ApiSearchPromotionsRequest {
	r.paginationToken = &paginationToken
	return r
}

// Specifies which promotion revision or revisions to match against when filtering. This controls which promotions are included in search results, not the shape of the response. The response always returns the published revision in the main body, with &#x60;latestRevision&#x60; included when the latest revision diverges.
func (r ApiSearchPromotionsRequest) Revision(revision string) ApiSearchPromotionsRequest {
	r.revision = &revision
	return r
}

// The maximum number of response results per page.
func (r ApiSearchPromotionsRequest) Limit(limit int64) ApiSearchPromotionsRequest {
	r.limit = &limit
	return r
}

// A comma-delimited list of datasets to include in the response.
func (r ApiSearchPromotionsRequest) IncludedData(includedData []string) ApiSearchPromotionsRequest {
	r.includedData = &includedData
	return r
}

func (r ApiSearchPromotionsRequest) Execute() (*SearchPromotionsResponse, *http.Response, error) {
	return r.ApiService.SearchPromotionsExecute(r)
}

/*
SearchPromotions Method for SearchPromotions

Search and filter promotions based on various criteria. Returns a paginated list of promotion summaries.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchPromotionsRequest
*/
func (a *PromotionsAPIService) SearchPromotions(ctx context.Context) ApiSearchPromotionsRequest {
	return ApiSearchPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchPromotionsResponse
func (a *PromotionsAPIService) SearchPromotionsExecute(r ApiSearchPromotionsRequest) (*SearchPromotionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchPromotionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PromotionsAPIService.SearchPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/promotions/2025-12-01/promotions"

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
	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "", "")
	} else {
		var defaultValue string = "en_US"
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", defaultValue, "", "")
		r.locale = &defaultValue
	}
	if r.statuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", r.statuses, "form", "csv")
	}
	if r.asins != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asins", r.asins, "form", "csv")
	}
	if r.skus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skus", r.skus, "form", "csv")
	}
	if r.promotionTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotionTypes", r.promotionTypes, "form", "csv")
	}
	if r.startDateBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDateBefore", r.startDateBefore, "", "")
	}
	if r.startDateAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDateAfter", r.startDateAfter, "", "")
	}
	if r.endDateBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDateBefore", r.endDateBefore, "", "")
	}
	if r.endDateAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDateAfter", r.endDateAfter, "", "")
	}
	if r.updateDateAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updateDateAfter", r.updateDateAfter, "", "")
	}
	if r.updateDateBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updateDateBefore", r.updateDateBefore, "", "")
	}
	if r.paginationToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paginationToken", r.paginationToken, "", "")
	}
	if r.revision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "revision", r.revision, "", "")
	} else {
		var defaultValue string = "PUBLISHED"
		parameterAddToHeaderOrQuery(localVarQueryParams, "revision", defaultValue, "", "")
		r.revision = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
		var defaultValue int64 = 20
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", defaultValue, "", "")
		r.limit = &defaultValue
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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPromotionRequest struct {
	ctx          context.Context
	ApiService   *PromotionsAPIService
	promotionId  string
	includedData *[]string
	locale       *string
}

// A comma-delimited list of datasets to include in the response.
func (r ApiGetPromotionRequest) IncludedData(includedData []string) ApiGetPromotionRequest {
	r.includedData = &includedData
	return r
}

// The locale of the promotion. Formatted as an ISO 639 language code, followed by an underscore, followed by an ISO 3166-1 alpha-2 country code.
func (r ApiGetPromotionRequest) Locale(locale string) ApiGetPromotionRequest {
	r.locale = &locale
	return r
}

func (r ApiGetPromotionRequest) Execute() (*GetPromotionResponse, *http.Response, error) {
	return r.ApiService.GetPromotionExecute(r)
}

/*
GetPromotion Method for GetPromotion

Retrieve details of a specified promotion.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param promotionId The ID of the promotion.
	@return ApiGetPromotionRequest
*/
func (a *PromotionsAPIService) GetPromotion(ctx context.Context, promotionId string) ApiGetPromotionRequest {
	return ApiGetPromotionRequest{
		ApiService:  a,
		ctx:         ctx,
		promotionId: promotionId,
	}
}

// Execute executes the request
//
//	@return GetPromotionResponse
func (a *PromotionsAPIService) GetPromotionExecute(r ApiGetPromotionRequest) (*GetPromotionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPromotionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PromotionsAPIService.GetPromotion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/promotions/2025-12-01/promotions/{promotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"promotionId"+"}", url.PathEscape(parameterValueToString(r.promotionId, "promotionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includedData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includedData", r.includedData, "form", "csv")
	}
	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "", "")
	} else {
		var defaultValue string = "en_US"
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", defaultValue, "", "")
		r.locale = &defaultValue
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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSelectionRequest struct {
	ctx             context.Context
	ApiService      *PromotionsAPIService
	promotionId     string
	selectionId     string
	revisionId      *int32
	locale          *string
	paginationToken *string
	limit           *int64
	includedData    *[]string
}

// The revision identifier for the selection. Use the &#x60;revisionId&#x60; from the &#x60;getPromotion&#x60; response. A promotion may have multiple selection revisions when an update is in progress. Passing the correct &#x60;revisionId&#x60; ensures you retrieve the expected data.
func (r ApiGetSelectionRequest) RevisionId(revisionId int32) ApiGetSelectionRequest {
	r.revisionId = &revisionId
	return r
}

// The locale of the promotion. Formatted as an ISO 639 language code, followed by an underscore, followed by an ISO 3166-1 alpha-2 country code.
func (r ApiGetSelectionRequest) Locale(locale string) ApiGetSelectionRequest {
	r.locale = &locale
	return r
}

// A token that you use to retrieve the next page of results. The response includes &#x60;paginationToken&#x60; when the number of results exceeds the specified &#x60;limit&#x60; value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until &#x60;paginationToken&#x60; is null. Note that this operation can return empty pages.
func (r ApiGetSelectionRequest) PaginationToken(paginationToken string) ApiGetSelectionRequest {
	r.paginationToken = &paginationToken
	return r
}

// The maximum number of response results per page.
func (r ApiGetSelectionRequest) Limit(limit int64) ApiGetSelectionRequest {
	r.limit = &limit
	return r
}

// A comma-delimited list of datasets to include in the response.
func (r ApiGetSelectionRequest) IncludedData(includedData []string) ApiGetSelectionRequest {
	r.includedData = &includedData
	return r
}

func (r ApiGetSelectionRequest) Execute() (*GetSelectionResponse, *http.Response, error) {
	return r.ApiService.GetSelectionExecute(r)
}

/*
GetSelection Method for GetSelection

Retrieve up to 100 product items that are associated with a specified promotion. This operation only supports items found in `SelectionType.ITEMS`. Items found in `SelectionType.CATALOG` are not supported. Selection objects always include `selectionDetails` with item information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param promotionId The ID of the promotion.
	@param selectionId The ID of the selection.
	@return ApiGetSelectionRequest
*/
func (a *PromotionsAPIService) GetSelection(ctx context.Context, promotionId string, selectionId string) ApiGetSelectionRequest {
	return ApiGetSelectionRequest{
		ApiService:  a,
		ctx:         ctx,
		promotionId: promotionId,
		selectionId: selectionId,
	}
}

// Execute executes the request
//
//	@return GetSelectionResponse
func (a *PromotionsAPIService) GetSelectionExecute(r ApiGetSelectionRequest) (*GetSelectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSelectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PromotionsAPIService.GetSelection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/promotions/2025-12-01/promotions/{promotionId}/selections/{selectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"promotionId"+"}", url.PathEscape(parameterValueToString(r.promotionId, "promotionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"selectionId"+"}", url.PathEscape(parameterValueToString(r.selectionId, "selectionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.revisionId == nil {
		return localVarReturnValue, nil, reportError("revisionId is required and must be specified")
	}
	if *r.revisionId < 1 {
		return localVarReturnValue, nil, reportError("revisionId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "revisionId", r.revisionId, "", "")
	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "", "")
	} else {
		var defaultValue string = "en_US"
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", defaultValue, "", "")
		r.locale = &defaultValue
	}
	if r.paginationToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "paginationToken", r.paginationToken, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
		var defaultValue int64 = 20
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", defaultValue, "", "")
		r.limit = &defaultValue
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

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
