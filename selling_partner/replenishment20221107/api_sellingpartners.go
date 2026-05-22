package replenishment20221107

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// SellingpartnersAPIService SellingpartnersAPI service
type SellingpartnersAPIService service

type ApiGetSellingPartnerMetricsRequest struct {
	ctx        context.Context
	ApiService *SellingpartnersAPIService
	body       *GetSellingPartnerMetricsRequest
}

// The request body for the &#x60;getSellingPartnerMetrics&#x60; operation.
func (r ApiGetSellingPartnerMetricsRequest) Body(body GetSellingPartnerMetricsRequest) ApiGetSellingPartnerMetricsRequest {
	r.body = &body
	return r
}

func (r ApiGetSellingPartnerMetricsRequest) Execute() (*GetSellingPartnerMetricsResponse, *http.Response, error) {
	return r.ApiService.GetSellingPartnerMetricsExecute(r)
}

/*
GetSellingPartnerMetrics Method for GetSellingPartnerMetrics

Returns aggregated replenishment program metrics for a selling partner.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSellingPartnerMetricsRequest
*/
func (a *SellingpartnersAPIService) GetSellingPartnerMetrics(ctx context.Context) ApiGetSellingPartnerMetricsRequest {
	return ApiGetSellingPartnerMetricsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSellingPartnerMetricsResponse
func (a *SellingpartnersAPIService) GetSellingPartnerMetricsExecute(r ApiGetSellingPartnerMetricsRequest) (*GetSellingPartnerMetricsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSellingPartnerMetricsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SellingpartnersAPIService.GetSellingPartnerMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replenishment/2022-11-07/sellingPartners/metrics/search"

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
