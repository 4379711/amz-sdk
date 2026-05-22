package seller_wallet_20240301

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AccountsAPIService AccountsAPI service
type AccountsAPIService service

type ApiGetAccountRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	accountId  string
}

func (r ApiGetAccountRequest) Execute() (*BankAccount, *http.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Find particular Amazon Seller Wallet account by Amazon account identifier

Retrieve an Amazon Seller Wallet bank account by Amazon account identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId The ID of the Amazon Seller Wallet account.
	@return ApiGetAccountRequest
*/
func (a *AccountsAPIService) GetAccount(ctx context.Context, accountId string) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return BankAccount
func (a *AccountsAPIService) GetAccountExecute(r ApiGetAccountRequest) (*BankAccount, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BankAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/transfers/wallet/2024-03-01/accounts/{accountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

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

type ApiListAccountBalancesRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	accountId  string
}

func (r ApiListAccountBalancesRequest) Execute() (*BalanceListing, *http.Response, error) {
	return r.ApiService.ListAccountBalancesExecute(r)
}

/*
ListAccountBalances Find balance in particular Amazon Seller Wallet account by Amazon account identifier

Retrieve the balance in a given Amazon Seller Wallet bank account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId The ID of the Amazon Seller Wallet account.
	@return ApiListAccountBalancesRequest
*/
func (a *AccountsAPIService) ListAccountBalances(ctx context.Context, accountId string) ApiListAccountBalancesRequest {
	return ApiListAccountBalancesRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return BalanceListing
func (a *AccountsAPIService) ListAccountBalancesExecute(r ApiListAccountBalancesRequest) (*BalanceListing, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BalanceListing
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ListAccountBalances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/transfers/wallet/2024-03-01/accounts/{accountId}/balance"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

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

type ApiListAccountsRequest struct {
	ctx           context.Context
	ApiService    *AccountsAPIService
	marketplaceId *string
}

// The marketplace for which items are returned. The marketplace ID is the globally unique identifier of a marketplace. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
func (r ApiListAccountsRequest) MarketplaceId(marketplaceId string) ApiListAccountsRequest {
	r.marketplaceId = &marketplaceId
	return r
}

func (r ApiListAccountsRequest) Execute() (*BankAccountListing, *http.Response, error) {
	return r.ApiService.ListAccountsExecute(r)
}

/*
ListAccounts Get all Amazon Seller Wallet accounts for the seller

Get all Seller Wallet accounts for a given seller.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAccountsRequest
*/
func (a *AccountsAPIService) ListAccounts(ctx context.Context) ApiListAccountsRequest {
	return ApiListAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BankAccountListing
func (a *AccountsAPIService) ListAccountsExecute(r ApiListAccountsRequest) (*BankAccountListing, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BankAccountListing
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ListAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finances/transfers/wallet/2024-03-01/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceId == nil {
		return localVarReturnValue, nil, reportError("marketplaceId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "marketplaceId", r.marketplaceId, "", "")
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
