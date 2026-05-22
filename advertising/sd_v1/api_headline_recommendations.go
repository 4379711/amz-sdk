package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// HeadlineRecommendationsAPIService HeadlineRecommendationsAPI service
type HeadlineRecommendationsAPIService service

type ApiGetHeadlineRecommendationsForSDRequest struct {
	ctx                             context.Context
	ApiService                      *HeadlineRecommendationsAPIService
	amazonAdvertisingAPIClientId    *string
	amazonAdvertisingAPIScope       *string
	sDHeadlineRecommendationRequest *SDHeadlineRecommendationRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetHeadlineRecommendationsForSDRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetHeadlineRecommendationsForSDRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetHeadlineRecommendationsForSDRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetHeadlineRecommendationsForSDRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Request body for SD headline recommendations API.
func (r ApiGetHeadlineRecommendationsForSDRequest) SDHeadlineRecommendationRequest(sDHeadlineRecommendationRequest SDHeadlineRecommendationRequest) ApiGetHeadlineRecommendationsForSDRequest {
	r.sDHeadlineRecommendationRequest = &sDHeadlineRecommendationRequest
	return r
}

func (r ApiGetHeadlineRecommendationsForSDRequest) Execute() (*SDHeadlineRecommendationResponse, *http.Response, error) {
	return r.ApiService.GetHeadlineRecommendationsForSDExecute(r)
}

func (a *HeadlineRecommendationsAPIService) GetHeadlineRecommendationsForSD(ctx context.Context) ApiGetHeadlineRecommendationsForSDRequest {
	return ApiGetHeadlineRecommendationsForSDRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SDHeadlineRecommendationResponse
func (a *HeadlineRecommendationsAPIService) GetHeadlineRecommendationsForSDExecute(r ApiGetHeadlineRecommendationsForSDRequest) (*SDHeadlineRecommendationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SDHeadlineRecommendationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HeadlineRecommendationsAPIService.GetHeadlineRecommendationsForSD")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/recommendations/creative/headline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sDHeadlineRecommendationRequest == nil {
		return localVarReturnValue, nil, reportError("sDHeadlineRecommendationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sdheadlinerecommendationrequest.v4.0+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sdheadlinerecommendationresponse.v4.0+json", "application/vnd.sdheadlinerecommendationschemavalidationexception.v4.0+json", "application/vnd.sdheadlinerecommendationaccessdeniedexception.v4.0+json", "application/vnd.sdheadlinerecommendationidentifiernotfoundexception.v4.0+json", "application/vnd.sdheadlinerecommendationthrottlingexception.v4.0+json", "application/vnd.sdheadlinerecommendationinternalserverexception.v4.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sDHeadlineRecommendationRequest
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
