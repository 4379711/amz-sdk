package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ConsolidatedRecommendationsAPIService ConsolidatedRecommendationsAPI service
type ConsolidatedRecommendationsAPIService service

type ApiGetCampaignRecommendationsRequest struct {
	ctx                          context.Context
	ApiService                   *ConsolidatedRecommendationsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	nextToken                    *string
	maxResults                   *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetCampaignRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetCampaignRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiGetCampaignRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetCampaignRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Optional. Token to retrieve subsequent page of results.
func (r ApiGetCampaignRecommendationsRequest) NextToken(nextToken string) ApiGetCampaignRecommendationsRequest {
	r.nextToken = &nextToken
	return r
}

// Optional. Limits the number of items to return in the response.
func (r ApiGetCampaignRecommendationsRequest) MaxResults(maxResults string) ApiGetCampaignRecommendationsRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiGetCampaignRecommendationsRequest) Execute() (*GetCampaignRecommendationsResponse, *http.Response, error) {
	return r.ApiService.GetCampaignRecommendationsExecute(r)
}

func (a *ConsolidatedRecommendationsAPIService) GetCampaignRecommendations(ctx context.Context) ApiGetCampaignRecommendationsRequest {
	return ApiGetCampaignRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCampaignRecommendationsResponse
func (a *ConsolidatedRecommendationsAPIService) GetCampaignRecommendationsExecute(r ApiGetCampaignRecommendationsRequest) (*GetCampaignRecommendationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCampaignRecommendationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsolidatedRecommendationsAPIService.GetCampaignRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/campaign/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxResults", r.maxResults, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spgetcampaignrecommendationsresponse.v1+json"}

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
