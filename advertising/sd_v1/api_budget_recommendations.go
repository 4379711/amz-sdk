package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// BudgetRecommendationsAPIService BudgetRecommendationsAPI service
type BudgetRecommendationsAPIService service

type ApiGetSDBudgetRecommendationsRequest struct {
	ctx                            context.Context
	ApiService                     *BudgetRecommendationsAPIService
	amazonAdvertisingAPIClientId   *string
	amazonAdvertisingAPIScope      *string
	amazonAdsAccountId             *string
	sDBudgetRecommendationsRequest *SDBudgetRecommendationsRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetSDBudgetRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetSDBudgetRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetSDBudgetRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetSDBudgetRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The identifier of an account. Account must be a global advertising account.
func (r ApiGetSDBudgetRecommendationsRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiGetSDBudgetRecommendationsRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

func (r ApiGetSDBudgetRecommendationsRequest) SDBudgetRecommendationsRequest(sDBudgetRecommendationsRequest SDBudgetRecommendationsRequest) ApiGetSDBudgetRecommendationsRequest {
	r.sDBudgetRecommendationsRequest = &sDBudgetRecommendationsRequest
	return r
}

func (r ApiGetSDBudgetRecommendationsRequest) Execute() (*SDBudgetRecommendationsResponse, *http.Response, error) {
	return r.ApiService.GetSDBudgetRecommendationsExecute(r)
}

func (a *BudgetRecommendationsAPIService) GetSDBudgetRecommendations(ctx context.Context) ApiGetSDBudgetRecommendationsRequest {
	return ApiGetSDBudgetRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SDBudgetRecommendationsResponse
func (a *BudgetRecommendationsAPIService) GetSDBudgetRecommendationsExecute(r ApiGetSDBudgetRecommendationsRequest) (*SDBudgetRecommendationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SDBudgetRecommendationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRecommendationsAPIService.GetSDBudgetRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns/budgetRecommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sdbudgetrecommendations.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sdbudgetrecommendations.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	}
	// body params
	localVarPostBody = r.sDBudgetRecommendationsRequest
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
