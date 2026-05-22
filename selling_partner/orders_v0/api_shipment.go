package orders_v0

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ShipmentAPIService ShipmentAPI service
type ShipmentAPIService service

type ApiUpdateShipmentStatusRequest struct {
	ctx        context.Context
	ApiService *ShipmentAPIService
	payload    *UpdateShipmentStatusRequest
	orderId    string
}

// The request body for the &#x60;updateShipmentStatus&#x60; operation.
func (r ApiUpdateShipmentStatusRequest) Payload(payload UpdateShipmentStatusRequest) ApiUpdateShipmentStatusRequest {
	r.payload = &payload
	return r
}

func (r ApiUpdateShipmentStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateShipmentStatusExecute(r)
}

/*
UpdateShipmentStatus Method for UpdateShipmentStatus

Update the shipment status for an order that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The preceding table contains the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may receive higher rate and burst values then those shown here. For more information, refer to [Usage Plans and Rate Limits](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId An Amazon-defined order identifier, in 3-7-7 format.
	@return ApiUpdateShipmentStatusRequest
*/
func (a *ShipmentAPIService) UpdateShipmentStatus(ctx context.Context, orderId string) ApiUpdateShipmentStatusRequest {
	return ApiUpdateShipmentStatusRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *ShipmentAPIService) UpdateShipmentStatusExecute(r ApiUpdateShipmentStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentAPIService.UpdateShipmentStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/v0/orders/{orderId}/shipment"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return nil, reportError("payload is required and must be specified")
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
	localVarPostBody = r.payload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	return localVarHTTPResponse, nil
}
