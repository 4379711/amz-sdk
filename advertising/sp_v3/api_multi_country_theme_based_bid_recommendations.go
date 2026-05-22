package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// MultiCountryThemeBasedBidRecommendationsAPIService MultiCountryThemeBasedBidRecommendationsAPI service
type MultiCountryThemeBasedBidRecommendationsAPIService service

type ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request struct {
	ctx                                                           context.Context
	ApiService                                                    *MultiCountryThemeBasedBidRecommendationsAPIService
	amazonAdvertisingAPIClientId                                  *string
	amazonAdsAccountId                                            *string
	getMultiCountryThemeBasedBidRecommendationForAdGroupV1Request *GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// Global Account Id Identifier.
func (r ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) AmazonAdsAccountId(amazonAdsAccountId string) ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

func (r ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request(getMultiCountryThemeBasedBidRecommendationForAdGroupV1Request GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	r.getMultiCountryThemeBasedBidRecommendationForAdGroupV1Request = &getMultiCountryThemeBasedBidRecommendationForAdGroupV1Request
	return r
}

func (r ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) Execute() (*MultiCountryThemeBasedBidRecommendationResponse, *http.Response, error) {
	return r.ApiService.GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Execute(r)
}

func (a *MultiCountryThemeBasedBidRecommendationsAPIService) GetMultiCountryThemeBasedBidRecommendationForAdGroupV1(ctx context.Context) ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request {
	return ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MultiCountryThemeBasedBidRecommendationResponse
func (a *MultiCountryThemeBasedBidRecommendationsAPIService) GetMultiCountryThemeBasedBidRecommendationForAdGroupV1Execute(r ApiGetMultiCountryThemeBasedBidRecommendationForAdGroupV1Request) (*MultiCountryThemeBasedBidRecommendationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MultiCountryThemeBasedBidRecommendationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MultiCountryThemeBasedBidRecommendationsAPIService.GetMultiCountryThemeBasedBidRecommendationForAdGroupV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/global/targets/bid/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.amazonAdsAccountId == nil {
		return localVarReturnValue, nil, reportError("amazonAdsAccountId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spthemebasedglobalbidrecommendation.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	// body params
	localVarPostBody = r.getMultiCountryThemeBasedBidRecommendationForAdGroupV1Request
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
