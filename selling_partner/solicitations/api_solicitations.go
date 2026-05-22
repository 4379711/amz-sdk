package solicitations

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SolicitationsAPIService SolicitationsAPI service
type SolicitationsAPIService service

type ApiCreateProductReviewAndSellerFeedbackSolicitationRequest struct {
	ctx            context.Context
	ApiService     *SolicitationsAPIService
	amazonOrderId  string
	marketplaceIds *[]string
}

// A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
func (r ApiCreateProductReviewAndSellerFeedbackSolicitationRequest) MarketplaceIds(marketplaceIds []string) ApiCreateProductReviewAndSellerFeedbackSolicitationRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiCreateProductReviewAndSellerFeedbackSolicitationRequest) Execute() (*CreateProductReviewAndSellerFeedbackSolicitationResponse, *http.Response, error) {
	return r.ApiService.CreateProductReviewAndSellerFeedbackSolicitationExecute(r)
}

/*
CreateProductReviewAndSellerFeedbackSolicitation Method for CreateProductReviewAndSellerFeedbackSolicitation

Sends a solicitation to a buyer asking for seller feedback and a product review for the specified order. Send only one productReviewAndSellerFeedback or free form proactive message per order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This specifies the order for which a solicitation is sent.
	@return ApiCreateProductReviewAndSellerFeedbackSolicitationRequest
*/
func (a *SolicitationsAPIService) CreateProductReviewAndSellerFeedbackSolicitation(ctx context.Context, amazonOrderId string) ApiCreateProductReviewAndSellerFeedbackSolicitationRequest {
	return ApiCreateProductReviewAndSellerFeedbackSolicitationRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return CreateProductReviewAndSellerFeedbackSolicitationResponse
func (a *SolicitationsAPIService) CreateProductReviewAndSellerFeedbackSolicitationExecute(r ApiCreateProductReviewAndSellerFeedbackSolicitationRequest) (*CreateProductReviewAndSellerFeedbackSolicitationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateProductReviewAndSellerFeedbackSolicitationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolicitationsAPIService.CreateProductReviewAndSellerFeedbackSolicitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solicitations/v1/orders/{amazonOrderId}/solicitations/productReviewAndSellerFeedback"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

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
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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

type ApiGetSolicitationActionsForOrderRequest struct {
	ctx            context.Context
	ApiService     *SolicitationsAPIService
	amazonOrderId  string
	marketplaceIds *[]string
}

// A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
func (r ApiGetSolicitationActionsForOrderRequest) MarketplaceIds(marketplaceIds []string) ApiGetSolicitationActionsForOrderRequest {
	r.marketplaceIds = &marketplaceIds
	return r
}

func (r ApiGetSolicitationActionsForOrderRequest) Execute() (*GetSolicitationActionsForOrderResponse, *http.Response, error) {
	return r.ApiService.GetSolicitationActionsForOrderExecute(r)
}

/*
GetSolicitationActionsForOrder Method for GetSolicitationActionsForOrder

Returns a list of solicitation types that are available for an order that you specify. A solicitation type is represented by an actions object, which contains a path and query parameter(s). You can use the path and parameter(s) to call an operation that sends a solicitation. Currently only the productReviewAndSellerFeedbackSolicitation solicitation type is available.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param amazonOrderId An Amazon order identifier. This specifies the order for which you want a list of available solicitation types.
	@return ApiGetSolicitationActionsForOrderRequest
*/
func (a *SolicitationsAPIService) GetSolicitationActionsForOrder(ctx context.Context, amazonOrderId string) ApiGetSolicitationActionsForOrderRequest {
	return ApiGetSolicitationActionsForOrderRequest{
		ApiService:    a,
		ctx:           ctx,
		amazonOrderId: amazonOrderId,
	}
}

// Execute executes the request
//
//	@return GetSolicitationActionsForOrderResponse
func (a *SolicitationsAPIService) GetSolicitationActionsForOrderExecute(r ApiGetSolicitationActionsForOrderRequest) (*GetSolicitationActionsForOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSolicitationActionsForOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SolicitationsAPIService.GetSolicitationActionsForOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/solicitations/v1/orders/{amazonOrderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"amazonOrderId"+"}", url.PathEscape(parameterValueToString(r.amazonOrderId, "amazonOrderId")), -1)

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
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json"}

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
