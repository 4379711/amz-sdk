package vendor_direct_fulfillment_payments_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// VendorInvoiceAPIService VendorInvoiceAPI service
type VendorInvoiceAPIService service

type ApiSubmitInvoiceRequest struct {
	ctx        context.Context
	ApiService *VendorInvoiceAPIService
	body       *SubmitInvoiceRequest
}

// The request body containing one or more invoices for vendor orders.
func (r ApiSubmitInvoiceRequest) Body(body SubmitInvoiceRequest) ApiSubmitInvoiceRequest {
	r.body = &body
	return r
}

func (r ApiSubmitInvoiceRequest) Execute() (*SubmitInvoiceResponse, *http.Response, error) {
	return r.ApiService.SubmitInvoiceExecute(r)
}

/*
SubmitInvoice Method for SubmitInvoice

Submits one or more invoices for a vendor's direct fulfillment orders.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSubmitInvoiceRequest
*/
func (a *VendorInvoiceAPIService) SubmitInvoice(ctx context.Context) ApiSubmitInvoiceRequest {
	return ApiSubmitInvoiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubmitInvoiceResponse
func (a *VendorInvoiceAPIService) SubmitInvoiceExecute(r ApiSubmitInvoiceRequest) (*SubmitInvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubmitInvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorInvoiceAPIService.SubmitInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor/directFulfillment/payments/v1/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
