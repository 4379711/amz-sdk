package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// BudgetRulesRecommendationAPIService BudgetRulesRecommendationAPI service
type BudgetRulesRecommendationAPIService service

type ApiSPGetAllRuleEventsRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetRulesRecommendationAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	body                         *map[string]interface{}
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiSPGetAllRuleEventsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSPGetAllRuleEventsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiSPGetAllRuleEventsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSPGetAllRuleEventsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSPGetAllRuleEventsRequest) Body(body map[string]interface{}) ApiSPGetAllRuleEventsRequest {
	r.body = &body
	return r
}

func (r ApiSPGetAllRuleEventsRequest) Execute() (*SPGetAllRuleEventResponse, *http.Response, error) {
	return r.ApiService.SPGetAllRuleEventsExecute(r)
}

func (a *BudgetRulesRecommendationAPIService) SPGetAllRuleEvents(ctx context.Context) ApiSPGetAllRuleEventsRequest {
	return ApiSPGetAllRuleEventsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SPGetAllRuleEventResponse
func (a *BudgetRulesRecommendationAPIService) SPGetAllRuleEventsExecute(r ApiSPGetAllRuleEventsRequest) (*SPGetAllRuleEventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SPGetAllRuleEventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesRecommendationAPIService.SPGetAllRuleEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/v1/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiSPGetBudgetRulesRecommendationRequest struct {
	ctx                                   context.Context
	ApiService                            *BudgetRulesRecommendationAPIService
	amazonAdvertisingAPIClientId          *string
	amazonAdvertisingAPIScope             *string
	sPGetBudgetRulesRecommendationRequest *SPGetBudgetRulesRecommendationRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiSPGetBudgetRulesRecommendationRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSPGetBudgetRulesRecommendationRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiSPGetBudgetRulesRecommendationRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSPGetBudgetRulesRecommendationRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSPGetBudgetRulesRecommendationRequest) SPGetBudgetRulesRecommendationRequest(sPGetBudgetRulesRecommendationRequest SPGetBudgetRulesRecommendationRequest) ApiSPGetBudgetRulesRecommendationRequest {
	r.sPGetBudgetRulesRecommendationRequest = &sPGetBudgetRulesRecommendationRequest
	return r
}

func (r ApiSPGetBudgetRulesRecommendationRequest) Execute() (*SPBudgetRulesRecommendationEventResponse, *http.Response, error) {
	return r.ApiService.SPGetBudgetRulesRecommendationExecute(r)
}

func (a *BudgetRulesRecommendationAPIService) SPGetBudgetRulesRecommendation(ctx context.Context) ApiSPGetBudgetRulesRecommendationRequest {
	return ApiSPGetBudgetRulesRecommendationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SPBudgetRulesRecommendationEventResponse
func (a *BudgetRulesRecommendationAPIService) SPGetBudgetRulesRecommendationExecute(r ApiSPGetBudgetRulesRecommendationRequest) (*SPBudgetRulesRecommendationEventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SPBudgetRulesRecommendationEventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetRulesRecommendationAPIService.SPGetBudgetRulesRecommendation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaigns/budgetRules/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spbudgetrulesrecommendation.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spbudgetrulesrecommendation.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sPGetBudgetRulesRecommendationRequest
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
