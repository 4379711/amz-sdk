package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// TargetingRecommendationsAPIService TargetingRecommendationsAPI service
type TargetingRecommendationsAPIService service

type ApiGetTargetRecommendationsRequest struct {
	ctx                                  context.Context
	ApiService                           *TargetingRecommendationsAPIService
	amazonAdvertisingAPIClientId         *string
	amazonAdvertisingAPIScope            *string
	locale                               *SDTargetingRecommendationsLocale
	sDTargetingRecommendationsRequestV35 *SDTargetingRecommendationsRequestV35
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetTargetRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetTargetRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiGetTargetRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetTargetRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The requested locale from query parameter to return translated category recommendations.
func (r ApiGetTargetRecommendationsRequest) Locale(locale SDTargetingRecommendationsLocale) ApiGetTargetRecommendationsRequest {
	r.locale = &locale
	return r
}

func (r ApiGetTargetRecommendationsRequest) SDTargetingRecommendationsRequestV35(sDTargetingRecommendationsRequestV35 SDTargetingRecommendationsRequestV35) ApiGetTargetRecommendationsRequest {
	r.sDTargetingRecommendationsRequestV35 = &sDTargetingRecommendationsRequestV35
	return r
}

func (r ApiGetTargetRecommendationsRequest) Execute() (*SDTargetingRecommendationsResponseV35, *http.Response, error) {
	return r.ApiService.GetTargetRecommendationsExecute(r)
}

func (a *TargetingRecommendationsAPIService) GetTargetRecommendations(ctx context.Context) ApiGetTargetRecommendationsRequest {
	return ApiGetTargetRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SDTargetingRecommendationsResponseV35
func (a *TargetingRecommendationsAPIService) GetTargetRecommendationsExecute(r ApiGetTargetRecommendationsRequest) (*SDTargetingRecommendationsResponseV35, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SDTargetingRecommendationsResponseV35
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetingRecommendationsAPIService.GetTargetRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/targets/recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locale", r.locale, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sdtargetingrecommendations.v3.5+json", "application/vnd.sdtargetingrecommendations.v3.4+json", "application/vnd.sdtargetingrecommendations.v3.3+json", "application/vnd.sdtargetingrecommendations.v3.2+json", "application/vnd.sdtargetingrecommendations.v3.1+json", "application/vnd.sdtargetingrecommendations.v3.0+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sdtargetingrecommendations.v3.5+json", "application/vnd.sdtargetingrecommendations.v3.4+json", "application/vnd.sdtargetingrecommendations.v3.3+json", "application/vnd.sdtargetingrecommendations.v3.2+json", "application/vnd.sdtargetingrecommendations.v3.1+json", "application/vnd.sdtargetingrecommendations.v3.0+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sDTargetingRecommendationsRequestV35
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
