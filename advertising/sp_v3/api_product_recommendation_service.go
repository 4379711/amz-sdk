package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ProductRecommendationServiceAPIService ProductRecommendationServiceAPI service
type ProductRecommendationServiceAPIService service

type ApiGetProductRecommendationsRequest struct {
	ctx                              context.Context
	ApiService                       *ProductRecommendationServiceAPIService
	amazonAdvertisingAPIClientId     *string
	amazonAdvertisingAPIScope        *string
	amazonAdvertisingAPIAdvertiserId *string
	getProductRecommendationsRequest *GetProductRecommendationsRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetProductRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetProductRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API.
func (r ApiGetProductRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetProductRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The Advertiser ID associated with the advertiser account.
func (r ApiGetProductRecommendationsRequest) AmazonAdvertisingAPIAdvertiserId(amazonAdvertisingAPIAdvertiserId string) ApiGetProductRecommendationsRequest {
	r.amazonAdvertisingAPIAdvertiserId = &amazonAdvertisingAPIAdvertiserId
	return r
}

func (r ApiGetProductRecommendationsRequest) GetProductRecommendationsRequest(getProductRecommendationsRequest GetProductRecommendationsRequest) ApiGetProductRecommendationsRequest {
	r.getProductRecommendationsRequest = &getProductRecommendationsRequest
	return r
}

func (r ApiGetProductRecommendationsRequest) Execute() (*ProductRecommendationsByTheme, *http.Response, error) {
	return r.ApiService.GetProductRecommendationsExecute(r)
}

func (a *ProductRecommendationServiceAPIService) GetProductRecommendations(ctx context.Context) ApiGetProductRecommendationsRequest {
	return ApiGetProductRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductRecommendationsByTheme
func (a *ProductRecommendationServiceAPIService) GetProductRecommendationsExecute(r ApiGetProductRecommendationsRequest) (*ProductRecommendationsByTheme, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductRecommendationsByTheme
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductRecommendationServiceAPIService.GetProductRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targets/products/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spproductrecommendation.v3+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spproductrecommendationresponse.themes.v3+json", "application/vnd.spproductrecommendationresponse.asins.v3+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.amazonAdvertisingAPIAdvertiserId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Advertising-API-AdvertiserId", r.amazonAdvertisingAPIAdvertiserId, "simple", "")
	}

	// body params
	localVarPostBody = r.getProductRecommendationsRequest
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
