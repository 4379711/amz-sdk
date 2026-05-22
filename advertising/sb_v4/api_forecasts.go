package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ForecastsAPIService ForecastsAPI service
type ForecastsAPIService service

type ApiSBCampaignPerformanceForecastsRequest struct {
	ctx                                          context.Context
	ApiService                                   *ForecastsAPIService
	amazonAdvertisingAPIClientId                 *string
	amazonAdvertisingAPIScope                    *string
	sBCampaignPerformanceForecastsRequestContent *SBCampaignPerformanceForecastsRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiSBCampaignPerformanceForecastsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiSBCampaignPerformanceForecastsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiSBCampaignPerformanceForecastsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiSBCampaignPerformanceForecastsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiSBCampaignPerformanceForecastsRequest) SBCampaignPerformanceForecastsRequestContent(sBCampaignPerformanceForecastsRequestContent SBCampaignPerformanceForecastsRequestContent) ApiSBCampaignPerformanceForecastsRequest {
	r.sBCampaignPerformanceForecastsRequestContent = &sBCampaignPerformanceForecastsRequestContent
	return r
}

func (r ApiSBCampaignPerformanceForecastsRequest) Execute() (*SBCampaignPerformanceForecastsResponseContent, *http.Response, error) {
	return r.ApiService.SBCampaignPerformanceForecastsExecute(r)
}

func (a *ForecastsAPIService) SBCampaignPerformanceForecasts(ctx context.Context) ApiSBCampaignPerformanceForecastsRequest {
	return ApiSBCampaignPerformanceForecastsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SBCampaignPerformanceForecastsResponseContent
func (a *ForecastsAPIService) SBCampaignPerformanceForecastsExecute(r ApiSBCampaignPerformanceForecastsRequest) (*SBCampaignPerformanceForecastsResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SBCampaignPerformanceForecastsResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ForecastsAPIService.SBCampaignPerformanceForecasts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/forecasts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sBCampaignPerformanceForecastsRequestContent == nil {
		return localVarReturnValue, nil, reportError("sBCampaignPerformanceForecastsRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbforecasting.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbforecasting.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sBCampaignPerformanceForecastsRequestContent
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
