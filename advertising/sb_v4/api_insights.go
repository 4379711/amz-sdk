package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// InsightsAPIService InsightsAPI service
type InsightsAPIService service

type ApiSBInsightsCampaignInsightsRequest struct {
	ctx                                      context.Context
	ApiService                               *InsightsAPIService
	amazonAdvertisingAPIClientId             *string
	amazonAdvertisingAPIScope                *string
	sBInsightsCampaignInsightsRequestContent *SBInsightsCampaignInsightsRequestContent
	nextToken                                *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiSBInsightsCampaignInsightsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSBInsightsCampaignInsightsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiSBInsightsCampaignInsightsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSBInsightsCampaignInsightsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSBInsightsCampaignInsightsRequest) SBInsightsCampaignInsightsRequestContent(sBInsightsCampaignInsightsRequestContent SBInsightsCampaignInsightsRequestContent) ApiSBInsightsCampaignInsightsRequest {
	r.sBInsightsCampaignInsightsRequestContent = &sBInsightsCampaignInsightsRequestContent
	return r
}

// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results.
func (r ApiSBInsightsCampaignInsightsRequest) NextToken(nextToken string) ApiSBInsightsCampaignInsightsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiSBInsightsCampaignInsightsRequest) Execute() (*SBInsightsCampaignInsightsResponseContent, *http.Response, error) {
	return r.ApiService.SBInsightsCampaignInsightsExecute(r)
}

func (a *InsightsAPIService) SBInsightsCampaignInsights(ctx context.Context) ApiSBInsightsCampaignInsightsRequest {
	return ApiSBInsightsCampaignInsightsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SBInsightsCampaignInsightsResponseContent
func (a *InsightsAPIService) SBInsightsCampaignInsightsExecute(r ApiSBInsightsCampaignInsightsRequest) (*SBInsightsCampaignInsightsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBInsightsCampaignInsightsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InsightsAPIService.SBInsightsCampaignInsights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/campaigns/insights"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sBInsightsCampaignInsightsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sBInsightsCampaignInsightsRequestContent is required and must be specified")
	}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbinsights.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbinsights.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sBInsightsCampaignInsightsRequestContent
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
