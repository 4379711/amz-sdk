package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// BudgetUsageAPIService BudgetUsageAPI service
type BudgetUsageAPIService service

type ApiSdCampaignsBudgetUsageRequest struct {
	ctx                          context.Context
	ApiService                   *BudgetUsageAPIService
	amazonAdvertisingAPIClientId *interface{}
	amazonAdvertisingAPIScope    *interface{}
	budgetUsageCampaignRequest   *BudgetUsageCampaignRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiSdCampaignsBudgetUsageRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId interface{}) ApiSdCampaignsBudgetUsageRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiSdCampaignsBudgetUsageRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope interface{}) ApiSdCampaignsBudgetUsageRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSdCampaignsBudgetUsageRequest) BudgetUsageCampaignRequest(budgetUsageCampaignRequest BudgetUsageCampaignRequest) ApiSdCampaignsBudgetUsageRequest {
	r.budgetUsageCampaignRequest = &budgetUsageCampaignRequest
	return r
}

func (r ApiSdCampaignsBudgetUsageRequest) Execute() (*BudgetUsageCampaignResponse, *http.Response, error) {
	return r.ApiService.SdCampaignsBudgetUsageExecute(r)
}

func (a *BudgetUsageAPIService) SdCampaignsBudgetUsage(ctx context.Context) ApiSdCampaignsBudgetUsageRequest {
	return ApiSdCampaignsBudgetUsageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BudgetUsageCampaignResponse
func (a *BudgetUsageAPIService) SdCampaignsBudgetUsageExecute(r ApiSdCampaignsBudgetUsageRequest) (*BudgetUsageCampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BudgetUsageCampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetUsageAPIService.SdCampaignsBudgetUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/campaigns/budget/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sdcampaignbudgetusage.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sdcampaignbudgetusage.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.budgetUsageCampaignRequest
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
