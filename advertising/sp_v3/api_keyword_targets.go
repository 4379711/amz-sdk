package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// KeywordTargetsAPIService KeywordTargetsAPI service
type KeywordTargetsAPIService service

type ApiGetGlobalRankedKeywordRecommendationRequest struct {
	ctx                                         context.Context
	ApiService                                  *KeywordTargetsAPIService
	amazonAdvertisingAPIClientId                *string
	amazonAdvertisingAPIScope                   *string
	amazonAdsAccountId                          *string
	amazonAdvertisingAPIMarketplaceId           *string
	amazonAdvertisingAPIAdvertiserId            *string
	getGlobalRankedKeywordRecommendationRequest *GetGlobalRankedKeywordRecommendationRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetGlobalRankedKeywordRecommendationRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetGlobalRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetGlobalRankedKeywordRecommendationRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetGlobalRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The identifier of a profile associated with the advertiser account. Used for authentication of Global Account.
func (r ApiGetGlobalRankedKeywordRecommendationRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiGetGlobalRankedKeywordRecommendationRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

// The advertiser&#39;s Marketplace ID associated with the advertiser account to support single marketplace request. Will not be used if global account id is provided.
func (r ApiGetGlobalRankedKeywordRecommendationRequest) AmazonAdvertisingAPIMarketplaceId(amazonAdvertisingAPIMarketplaceId string) ApiGetGlobalRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIMarketplaceId = &amazonAdvertisingAPIMarketplaceId
	return r
}

// The advertiser&#39;s ID associated with the advertiser account to support single marketplace request. Will not be used if global account id is provided.
func (r ApiGetGlobalRankedKeywordRecommendationRequest) AmazonAdvertisingAPIAdvertiserId(amazonAdvertisingAPIAdvertiserId string) ApiGetGlobalRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIAdvertiserId = &amazonAdvertisingAPIAdvertiserId
	return r
}

func (r ApiGetGlobalRankedKeywordRecommendationRequest) GetGlobalRankedKeywordRecommendationRequest(getGlobalRankedKeywordRecommendationRequest GetGlobalRankedKeywordRecommendationRequest) ApiGetGlobalRankedKeywordRecommendationRequest {
	r.getGlobalRankedKeywordRecommendationRequest = &getGlobalRankedKeywordRecommendationRequest
	return r
}

func (r ApiGetGlobalRankedKeywordRecommendationRequest) Execute() (*GlobalRankedTargetWithThemedBidsResponse, *http.Response, error) {
	return r.ApiService.GetGlobalRankedKeywordRecommendationExecute(r)
}

func (a *KeywordTargetsAPIService) GetGlobalRankedKeywordRecommendation(ctx context.Context) ApiGetGlobalRankedKeywordRecommendationRequest {
	return ApiGetGlobalRankedKeywordRecommendationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GlobalRankedTargetWithThemedBidsResponse
func (a *KeywordTargetsAPIService) GetGlobalRankedKeywordRecommendationExecute(r ApiGetGlobalRankedKeywordRecommendationRequest) (*GlobalRankedTargetWithThemedBidsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GlobalRankedTargetWithThemedBidsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordTargetsAPIService.GetGlobalRankedKeywordRecommendation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/global/targets/keywords/recommendations/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.amazonAdsAccountId == nil {
		return localVarReturnValue, nil, reportError("amazonAdsAccountId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spkeywordsrecommendation.v5+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spkeywordsrecommendation.v5+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.amazonAdvertisingAPIMarketplaceId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Advertising-API-MarketplaceId", r.amazonAdvertisingAPIMarketplaceId, "simple", "")
	}
	if r.amazonAdvertisingAPIAdvertiserId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Advertising-API-AdvertiserId", r.amazonAdvertisingAPIAdvertiserId, "simple", "")
	}

	parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	// body params
	localVarPostBody = r.getGlobalRankedKeywordRecommendationRequest
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

type ApiGetRankedKeywordRecommendationRequest struct {
	ctx                                   context.Context
	ApiService                            *KeywordTargetsAPIService
	amazonAdvertisingAPIClientId          *string
	amazonAdvertisingAPIScope             *string
	amazonAdvertisingAPIMarketplaceId     *string
	amazonAdvertisingAPIAdvertiserId      *string
	getRankedKeywordRecommendationRequest *GetRankedKeywordRecommendationRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetRankedKeywordRecommendationRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetRankedKeywordRecommendationRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The advertiser&#39;s Marketplace ID associated with the advertiser. account.
func (r ApiGetRankedKeywordRecommendationRequest) AmazonAdvertisingAPIMarketplaceId(amazonAdvertisingAPIMarketplaceId string) ApiGetRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIMarketplaceId = &amazonAdvertisingAPIMarketplaceId
	return r
}

// The advertiser&#39;s ID associated with the advertiser account.
func (r ApiGetRankedKeywordRecommendationRequest) AmazonAdvertisingAPIAdvertiserId(amazonAdvertisingAPIAdvertiserId string) ApiGetRankedKeywordRecommendationRequest {
	r.amazonAdvertisingAPIAdvertiserId = &amazonAdvertisingAPIAdvertiserId
	return r
}

func (r ApiGetRankedKeywordRecommendationRequest) GetRankedKeywordRecommendationRequest(getRankedKeywordRecommendationRequest GetRankedKeywordRecommendationRequest) ApiGetRankedKeywordRecommendationRequest {
	r.getRankedKeywordRecommendationRequest = &getRankedKeywordRecommendationRequest
	return r
}

func (r ApiGetRankedKeywordRecommendationRequest) Execute() (*KeywordTargetResponse, *http.Response, error) {
	return r.ApiService.GetRankedKeywordRecommendationExecute(r)
}

func (a *KeywordTargetsAPIService) GetRankedKeywordRecommendation(ctx context.Context) ApiGetRankedKeywordRecommendationRequest {
	return ApiGetRankedKeywordRecommendationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KeywordTargetResponse
func (a *KeywordTargetsAPIService) GetRankedKeywordRecommendationExecute(r ApiGetRankedKeywordRecommendationRequest) (*KeywordTargetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeywordTargetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordTargetsAPIService.GetRankedKeywordRecommendation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/keywords/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spkeywordsrecommendation.v3+json", "application/vnd.spkeywordsrecommendation.v5+json", "application/vnd.spkeywordsrecommendation.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spkeywordsrecommendation.v3+json", "application/vnd.spkeywordsrecommendation.v5+json", "application/vnd.spkeywordsrecommendation.v4+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.amazonAdvertisingAPIMarketplaceId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Advertising-API-MarketplaceId", r.amazonAdvertisingAPIMarketplaceId, "simple", "")
	}
	if r.amazonAdvertisingAPIAdvertiserId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Advertising-API-AdvertiserId", r.amazonAdvertisingAPIAdvertiserId, "simple", "")
	}

	// body params
	localVarPostBody = r.getRankedKeywordRecommendationRequest
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
