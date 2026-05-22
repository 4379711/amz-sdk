package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// BidRecommendationsAPIService BidRecommendationsAPI service
type BidRecommendationsAPIService service

type ApiGetTargetBidRecommendationsRequest struct {
	ctx                                     context.Context
	ApiService                              *BidRecommendationsAPIService
	amazonAdvertisingAPIClientId            *string
	amazonAdvertisingAPIScope               *string
	sDTargetingBidRecommendationsRequestV34 *SDTargetingBidRecommendationsRequestV34
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetTargetBidRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetTargetBidRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetTargetBidRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetTargetBidRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiGetTargetBidRecommendationsRequest) SDTargetingBidRecommendationsRequestV34(sDTargetingBidRecommendationsRequestV34 SDTargetingBidRecommendationsRequestV34) ApiGetTargetBidRecommendationsRequest {
	r.sDTargetingBidRecommendationsRequestV34 = &sDTargetingBidRecommendationsRequestV34
	return r
}

func (r ApiGetTargetBidRecommendationsRequest) Execute() (*SDTargetingBidRecommendationsResponseV32, *http.Response, error) {
	return r.ApiService.GetTargetBidRecommendationsExecute(r)
}

func (a *BidRecommendationsAPIService) GetTargetBidRecommendations(ctx context.Context) ApiGetTargetBidRecommendationsRequest {
	return ApiGetTargetBidRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SDTargetingBidRecommendationsResponseV32
func (a *BidRecommendationsAPIService) GetTargetBidRecommendationsExecute(r ApiGetTargetBidRecommendationsRequest) (*SDTargetingBidRecommendationsResponseV32, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SDTargetingBidRecommendationsResponseV32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BidRecommendationsAPIService.GetTargetBidRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets/bid/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sdtargetingrecommendations.v3.4+json", "application/vnd.sdtargetingrecommendations.v3.3+json", "application/vnd.sdtargetingrecommendations.v3.2+json", "application/vnd.sdtargetingrecommendations.v3.1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sdtargetingrecommendations.v3.3+json", "application/vnd.sdtargetingrecommendations.v3.2+json", "application/vnd.sdtargetingrecommendations.v3.1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sDTargetingBidRecommendationsRequestV34
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
