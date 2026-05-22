package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// BudgetRecommendationsAndMissedOpportunitiesAPIService BudgetRecommendationsAndMissedOpportunitiesAPI service
type BudgetRecommendationsAndMissedOpportunitiesAPIService service

type ApiGetBudgetRecommendationsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRecommendationsAndMissedOpportunitiesAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	budgetRecommendationRequest  *BudgetRecommendationRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetBudgetRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetBudgetRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetBudgetRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetBudgetRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetBudgetRecommendationsRequest) BudgetRecommendationRequest(budgetRecommendationRequest BudgetRecommendationRequest) ApiGetBudgetRecommendationsRequest {
	r.budgetRecommendationRequest = &budgetRecommendationRequest
	return r
}

func (r ApiGetBudgetRecommendationsRequest) Execute() (*BudgetRecommendationResponse, *http.Response, error) {
	return r.ApiService.GetBudgetRecommendationsExecute(r)
}

func (a *BudgetRecommendationsAndMissedOpportunitiesAPIService) GetBudgetRecommendations(ctx context.Context) ApiGetBudgetRecommendationsRequest {
	return ApiGetBudgetRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BudgetRecommendationResponse
func (a *BudgetRecommendationsAndMissedOpportunitiesAPIService) GetBudgetRecommendationsExecute(r ApiGetBudgetRecommendationsRequest) (*BudgetRecommendationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BudgetRecommendationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRecommendationsAndMissedOpportunitiesAPIService.GetBudgetRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/budgetRecommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.budgetRecommendationRequest == nil {
		return localVarReturnValue, nil, reportError("budgetRecommendationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.budgetrecommendation.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.budgetrecommendation.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.budgetRecommendationRequest
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
