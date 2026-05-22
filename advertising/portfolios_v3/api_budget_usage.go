package portfolios_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// BudgetUsageAPIService BudgetUsageAPI service
type BudgetUsageAPIService service

type ApiPortfolioBudgetUsageRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetUsageAPIService
	amazonAdvertisingAPIClientId *interface{}
	amazonAdvertisingAPIScope    *interface{}
	budgetUsagePortfolioRequest  *BudgetUsagePortfolioRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiPortfolioBudgetUsageRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId interface{}) ApiPortfolioBudgetUsageRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiPortfolioBudgetUsageRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope interface{}) ApiPortfolioBudgetUsageRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiPortfolioBudgetUsageRequest) BudgetUsagePortfolioRequest(budgetUsagePortfolioRequest BudgetUsagePortfolioRequest) ApiPortfolioBudgetUsageRequest {
	r.budgetUsagePortfolioRequest = &budgetUsagePortfolioRequest
	return r
}

func (r ApiPortfolioBudgetUsageRequest) Execute() (*BudgetUsagePortfolioResponse, *http.Response, error) {
	return r.ApiService.PortfolioBudgetUsageExecute(r)
}

func (a *BudgetUsageAPIService) PortfolioBudgetUsage(ctx context.Context) ApiPortfolioBudgetUsageRequest {
	return ApiPortfolioBudgetUsageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BudgetUsagePortfolioResponse
func (a *BudgetUsageAPIService) PortfolioBudgetUsageExecute(r ApiPortfolioBudgetUsageRequest) (*BudgetUsagePortfolioResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BudgetUsagePortfolioResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetUsageAPIService.PortfolioBudgetUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/portfolios/budget/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.portfoliobudgetusage.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.portfoliobudgetusage.v1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.budgetUsagePortfolioRequest
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
