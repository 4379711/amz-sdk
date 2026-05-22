package vendor_direct_fulfillment_shipping_20211228

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CreateContainerLabelAPIService CreateContainerLabelAPI service
type CreateContainerLabelAPIService service

type ApiCreateContainerLabelRequest struct {
	ctx        context.Context
	ApiService *CreateContainerLabelAPIService
	body       *CreateContainerLabelRequest
}

// Request body containing the container label data.
func (r ApiCreateContainerLabelRequest) Body(body CreateContainerLabelRequest) ApiCreateContainerLabelRequest {
	r.body = &body
	return r
}

func (r ApiCreateContainerLabelRequest) Execute() (*CreateContainerLabelResponse, *http.Response, error) {
	return r.ApiService.CreateContainerLabelExecute(r)
}

/*
CreateContainerLabel createContainerLabel

Creates a container (pallet) label for the associated shipment package.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may have higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateContainerLabelRequest
*/
func (a *CreateContainerLabelAPIService) CreateContainerLabel(ctx context.Context) ApiCreateContainerLabelRequest {
	return ApiCreateContainerLabelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateContainerLabelResponse
func (a *CreateContainerLabelAPIService) CreateContainerLabelExecute(r ApiCreateContainerLabelRequest) (*CreateContainerLabelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateContainerLabelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreateContainerLabelAPIService.CreateContainerLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor/directFulfillment/shipping/2021-12-28/containerLabel"

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
	localVarHTTPHeaderAccepts := []string{"application/json", "containerLabel"}

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
