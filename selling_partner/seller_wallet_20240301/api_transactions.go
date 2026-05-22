package seller_wallet_20240301

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// TransactionsAPIService TransactionsAPI service
type TransactionsAPIService service

type ApiCreateTransactionRequest struct {
	ctx                         context.Context
	ApiService                  *TransactionsAPIService
	body                        *TransactionInitiationRequest
	destAccountDigitalSignature *string
	amountDigitalSignature      *string
}

// The payload of the request
func (r ApiCreateTransactionRequest) Body(body TransactionInitiationRequest) ApiCreateTransactionRequest {
	r.body = &body
	return r
}

// Digital signature for the destination bank account details.
func (r ApiCreateTransactionRequest) DestAccountDigitalSignature(destAccountDigitalSignature string) ApiCreateTransactionRequest {
	r.destAccountDigitalSignature = &destAccountDigitalSignature
	return r
}

// Digital signature for the source currency transaction amount.
func (r ApiCreateTransactionRequest) AmountDigitalSignature(amountDigitalSignature string) ApiCreateTransactionRequest {
	r.amountDigitalSignature = &amountDigitalSignature
	return r
}

func (r ApiCreateTransactionRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.CreateTransactionExecute(r)
}

/*
CreateTransaction Create a transaction request from Amazon Seller Wallet account to another customer-provided account

Create a transaction request from an Amazon Seller Wallet account to another customer-provided account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTransactionRequest
*/
func (a *TransactionsAPIService) CreateTransaction(ctx context.Context) ApiCreateTransactionRequest {
	return ApiCreateTransactionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Transaction
func (a *TransactionsAPIService) CreateTransactionExecute(r ApiCreateTransactionRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.CreateTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/transfers/wallet/2024-03-01/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.destAccountDigitalSignature == nil {
		return localVarReturnValue, nil, reportError("destAccountDigitalSignature is required and must be specified")
	}
	if r.amountDigitalSignature == nil {
		return localVarReturnValue, nil, reportError("amountDigitalSignature is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "destAccountDigitalSignature", r.destAccountDigitalSignature, "", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "amountDigitalSignature", r.amountDigitalSignature, "", "")
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

type ApiGetTransactionRequest struct {
	ctx           context.Context
	ApiService    *TransactionsAPIService
	transactionId string
}

func (r ApiGetTransactionRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.GetTransactionExecute(r)
}

/*
GetTransaction Find particular Amazon Seller Wallet account transaction by Amazon transaction identifier

Find a transaction by the Amazon transaction identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId The ID of the Amazon Seller Wallet transaction.
	@return ApiGetTransactionRequest
*/
func (a *TransactionsAPIService) GetTransaction(ctx context.Context, transactionId string) ApiGetTransactionRequest {
	return ApiGetTransactionRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//
//	@return Transaction
func (a *TransactionsAPIService) GetTransactionExecute(r ApiGetTransactionRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/transfers/wallet/2024-03-01/transactions/{transactionId}"
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

type ApiListAccountTransactionsRequest struct {
	ctx           context.Context
	ApiService    *TransactionsAPIService
	accountId     *string
	nextPageToken *string
}

// The ID of the Amazon Seller Wallet account.
func (r ApiListAccountTransactionsRequest) AccountId(accountId string) ApiListAccountTransactionsRequest {
	r.accountId = &accountId
	return r
}

// A token that you use to retrieve the next page of results. The response includes &#x60;nextPageToken&#x60; when the number of results exceeds 100. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until &#x60;nextPageToken&#x60; is null. Note that this operation can return empty pages.
func (r ApiListAccountTransactionsRequest) NextPageToken(nextPageToken string) ApiListAccountTransactionsRequest {
	r.nextPageToken = &nextPageToken
	return r
}

func (r ApiListAccountTransactionsRequest) Execute() (*TransactionListing, *http.Response, error) {
	return r.ApiService.ListAccountTransactionsExecute(r)
}

/*
ListAccountTransactions The API will return all the transactions for a given Amazon Seller Wallet account sorted by the transaction request date

Retrieve a list of transactions for a given Amazon Seller Wallet bank account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAccountTransactionsRequest
*/
func (a *TransactionsAPIService) ListAccountTransactions(ctx context.Context) ApiListAccountTransactionsRequest {
	return ApiListAccountTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TransactionListing
func (a *TransactionsAPIService) ListAccountTransactionsExecute(r ApiListAccountTransactionsRequest) (*TransactionListing, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransactionListing
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.ListAccountTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/transfers/wallet/2024-03-01/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, reportError("accountId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "", "")
	if r.nextPageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextPageToken", r.nextPageToken, "", "")
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
