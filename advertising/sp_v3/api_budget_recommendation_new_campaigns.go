package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// BudgetRecommendationNewCampaignsAPIService BudgetRecommendationNewCampaignsAPI service
type BudgetRecommendationNewCampaignsAPIService service

type ApiGetBudgetRecommendationRequest struct {
	ctx                                context.Context
	ApiService                         *BudgetRecommendationNewCampaignsAPIService
	amazonAdvertisingAPIClientId       *string
	amazonAdvertisingAPIScope          *string
	initialBudgetRecommendationRequest *InitialBudgetRecommendationRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetBudgetRecommendationRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetBudgetRecommendationRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiGetBudgetRecommendationRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetBudgetRecommendationRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetBudgetRecommendationRequest) InitialBudgetRecommendationRequest(initialBudgetRecommendationRequest InitialBudgetRecommendationRequest) ApiGetBudgetRecommendationRequest {
	r.initialBudgetRecommendationRequest = &initialBudgetRecommendationRequest
	return r
}

func (r ApiGetBudgetRecommendationRequest) Execute() (*InitialBudgetRecommendationResponse, *http.Response, error) {
	return r.ApiService.GetBudgetRecommendationExecute(r)
}

func (a *BudgetRecommendationNewCampaignsAPIService) GetBudgetRecommendation(ctx context.Context) ApiGetBudgetRecommendationRequest {
	return ApiGetBudgetRecommendationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InitialBudgetRecommendationResponse
func (a *BudgetRecommendationNewCampaignsAPIService) GetBudgetRecommendationExecute(r ApiGetBudgetRecommendationRequest) (*InitialBudgetRecommendationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InitialBudgetRecommendationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRecommendationNewCampaignsAPIService.GetBudgetRecommendation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/initialBudgetRecommendation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spinitialbudgetrecommendation.v3.4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spinitialbudgetrecommendation.v3.4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.initialBudgetRecommendationRequest
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
