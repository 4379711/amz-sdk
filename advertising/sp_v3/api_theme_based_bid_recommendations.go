package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ThemeBasedBidRecommendationsAPIService ThemeBasedBidRecommendationsAPI service
type ThemeBasedBidRecommendationsAPIService service

type ApiGetThemeBasedBidRecommendationForAdGroupV1Request struct {
	ctx                                               context.Context
	ApiService                                        *ThemeBasedBidRecommendationsAPIService
	amazonAdvertisingAPIClientId                      *string
	amazonAdvertisingAPIScope                         *string
	getThemeBasedBidRecommendationForAdGroupV1Request *GetThemeBasedBidRecommendationForAdGroupV1Request
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetThemeBasedBidRecommendationForAdGroupV1Request) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetThemeBasedBidRecommendationForAdGroupV1Request {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiGetThemeBasedBidRecommendationForAdGroupV1Request) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetThemeBasedBidRecommendationForAdGroupV1Request {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetThemeBasedBidRecommendationForAdGroupV1Request) GetThemeBasedBidRecommendationForAdGroupV1Request(getThemeBasedBidRecommendationForAdGroupV1Request GetThemeBasedBidRecommendationForAdGroupV1Request) ApiGetThemeBasedBidRecommendationForAdGroupV1Request {
	r.getThemeBasedBidRecommendationForAdGroupV1Request = &getThemeBasedBidRecommendationForAdGroupV1Request
	return r
}

func (r ApiGetThemeBasedBidRecommendationForAdGroupV1Request) Execute() (*ThemeBasedBidRecommendationResponseV4, *http.Response, error) {
	return r.ApiService.GetThemeBasedBidRecommendationForAdGroupV1Execute(r)
}

func (a *ThemeBasedBidRecommendationsAPIService) GetThemeBasedBidRecommendationForAdGroupV1(ctx context.Context) ApiGetThemeBasedBidRecommendationForAdGroupV1Request {
	return ApiGetThemeBasedBidRecommendationForAdGroupV1Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ThemeBasedBidRecommendationResponseV4
func (a *ThemeBasedBidRecommendationsAPIService) GetThemeBasedBidRecommendationForAdGroupV1Execute(r ApiGetThemeBasedBidRecommendationForAdGroupV1Request) (*ThemeBasedBidRecommendationResponseV4, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ThemeBasedBidRecommendationResponseV4
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemeBasedBidRecommendationsAPIService.GetThemeBasedBidRecommendationForAdGroupV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/bid/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spthemebasedbidrecommendation.v4+json", "application/vnd.spthemebasedbidrecommendation.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spthemebasedbidrecommendation.v4+json", "application/vnd.spthemebasedbidrecommendation.v3+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.getThemeBasedBidRecommendationForAdGroupV1Request
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
