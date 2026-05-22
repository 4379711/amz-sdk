package seller_wallet_20240301

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// TransferPreviewAPIService TransferPreviewAPI service
type TransferPreviewAPIService service

type ApiGetTransferPreviewRequest struct {
	ctx                     context.Context
	ApiService              *TransferPreviewAPIService
	sourceCountryCode       *string
	sourceCurrencyCode      *string
	destinationCountryCode  *string
	destinationCurrencyCode *string
	baseAmount              *float32
}

// Country code of the source transaction account in ISO 3166 format.
func (r ApiGetTransferPreviewRequest) SourceCountryCode(sourceCountryCode string) ApiGetTransferPreviewRequest {
	r.sourceCountryCode = &sourceCountryCode
	return r
}

// Currency code of the source transaction country in ISO 4217 format.
func (r ApiGetTransferPreviewRequest) SourceCurrencyCode(sourceCurrencyCode string) ApiGetTransferPreviewRequest {
	r.sourceCurrencyCode = &sourceCurrencyCode
	return r
}

// Country code of the destination transaction account in ISO 3166 format.
func (r ApiGetTransferPreviewRequest) DestinationCountryCode(destinationCountryCode string) ApiGetTransferPreviewRequest {
	r.destinationCountryCode = &destinationCountryCode
	return r
}

// Currency code of the destination transaction country in ISO 4217 format.
func (r ApiGetTransferPreviewRequest) DestinationCurrencyCode(destinationCurrencyCode string) ApiGetTransferPreviewRequest {
	r.destinationCurrencyCode = &destinationCurrencyCode
	return r
}

// The base transaction amount without any markup fees.
func (r ApiGetTransferPreviewRequest) BaseAmount(baseAmount float32) ApiGetTransferPreviewRequest {
	r.baseAmount = &baseAmount
	return r
}

func (r ApiGetTransferPreviewRequest) Execute() (*TransferRatePreview, *http.Response, error) {
	return r.ApiService.GetTransferPreviewExecute(r)
}

/*
GetTransferPreview Fetch potential fees that could be applied on a transaction on the basis of the source and destination country currency code

Retrieve a list of potential fees on a transaction.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTransferPreviewRequest
*/
func (a *TransferPreviewAPIService) GetTransferPreview(ctx context.Context) ApiGetTransferPreviewRequest {
	return ApiGetTransferPreviewRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TransferRatePreview
func (a *TransferPreviewAPIService) GetTransferPreviewExecute(r ApiGetTransferPreviewRequest) (*TransferRatePreview, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransferRatePreview
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransferPreviewAPIService.GetTransferPreview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/transfers/wallet/2024-03-01/transferPreview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sourceCountryCode == nil {
		return localVarReturnValue, nil, reportError("sourceCountryCode is required and must be specified")
	}
	if r.sourceCurrencyCode == nil {
		return localVarReturnValue, nil, reportError("sourceCurrencyCode is required and must be specified")
	}
	if r.destinationCountryCode == nil {
		return localVarReturnValue, nil, reportError("destinationCountryCode is required and must be specified")
	}
	if r.destinationCurrencyCode == nil {
		return localVarReturnValue, nil, reportError("destinationCurrencyCode is required and must be specified")
	}
	if r.baseAmount == nil {
		return localVarReturnValue, nil, reportError("baseAmount is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "sourceCountryCode", r.sourceCountryCode, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "sourceCurrencyCode", r.sourceCurrencyCode, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "destinationCountryCode", r.destinationCountryCode, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "destinationCurrencyCode", r.destinationCurrencyCode, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "baseAmount", r.baseAmount, "", "")
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
