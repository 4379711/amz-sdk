package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// VendorDFSandboxAPIService VendorDFSandboxAPI service
type VendorDFSandboxAPIService service

type ApiGenerateOrderScenariosRequest struct {
	ctx        context.Context
	ApiService *VendorDFSandboxAPIService
	body       *GenerateOrderScenarioRequest
}

// The request payload containing parameters for generating test order data scenarios.
func (r ApiGenerateOrderScenariosRequest) Body(body GenerateOrderScenarioRequest) ApiGenerateOrderScenariosRequest {
	r.body = &body
	return r
}

func (r ApiGenerateOrderScenariosRequest) Execute() (*TransactionReference, *http.Response, error) {
	return r.ApiService.GenerateOrderScenariosExecute(r)
}

/*
GenerateOrderScenarios Method for GenerateOrderScenarios

Submits a request to generate test order data for Vendor Direct Fulfillment API entities.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGenerateOrderScenariosRequest
*/
func (a *VendorDFSandboxAPIService) GenerateOrderScenarios(ctx context.Context) ApiGenerateOrderScenariosRequest {
	return ApiGenerateOrderScenariosRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TransactionReference
func (a *VendorDFSandboxAPIService) GenerateOrderScenariosExecute(r ApiGenerateOrderScenariosRequest) (*TransactionReference, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransactionReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorDFSandboxAPIService.GenerateOrderScenarios")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor/directFulfillment/sandbox/2021-10-28/orders"

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
