package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// VendorDFSandboxtransactionstatusAPIService VendorDFSandboxtransactionstatusAPI service
type VendorDFSandboxtransactionstatusAPIService service

type ApiGetOrderScenariosRequest struct {
	ctx           context.Context
	ApiService    *VendorDFSandboxtransactionstatusAPIService
	transactionId string
}

func (r ApiGetOrderScenariosRequest) Execute() (*TransactionStatus, *http.Response, error) {
	return r.ApiService.GetOrderScenariosExecute(r)
}

/*
GetOrderScenarios Method for GetOrderScenarios

Returns the status of the transaction indicated by the specified transactionId. If the transaction was successful, also returns the requested test order data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId The transaction identifier returned in the response to the generateOrderScenarios operation.
	@return ApiGetOrderScenariosRequest
*/
func (a *VendorDFSandboxtransactionstatusAPIService) GetOrderScenarios(ctx context.Context, transactionId string) ApiGetOrderScenariosRequest {
	return ApiGetOrderScenariosRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//
//	@return TransactionStatus
func (a *VendorDFSandboxtransactionstatusAPIService) GetOrderScenariosExecute(r ApiGetOrderScenariosRequest) (*TransactionStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransactionStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VendorDFSandboxtransactionstatusAPIService.GetOrderScenarios")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vendor/directFulfillment/sandbox/2021-10-28/transactions/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

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
