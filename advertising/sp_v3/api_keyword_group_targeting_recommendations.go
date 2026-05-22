package sp_v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// KeywordGroupTargetingRecommendationsAPIService KeywordGroupTargetingRecommendationsAPI service
type KeywordGroupTargetingRecommendationsAPIService service

type ApiGetKeywordGroupRecommendationsRequest struct {
	ctx                                 context.Context
	ApiService                          *KeywordGroupTargetingRecommendationsAPIService
	amazonAdvertisingAPIClientId        *string
	amazonAdvertisingAPIScope           *string
	locale                              *string
	keywordGroupsRecommendationsRequest *KeywordGroupsRecommendationsRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetKeywordGroupRecommendationsRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetKeywordGroupRecommendationsRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiGetKeywordGroupRecommendationsRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiGetKeywordGroupRecommendationsRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// User specified locale. If nothing is passed the default for the marketplace will be used. The value should confirm to the IETF BCP 47 standard, using language tags composed of language- and optionally region specific sub-tags (e.g., &#39;en-us&#39; for American English and &#39;fr-CA&#39; for Canadian French).
func (r ApiGetKeywordGroupRecommendationsRequest) Locale(locale string) ApiGetKeywordGroupRecommendationsRequest {
	r.locale = &locale
	return r
}

func (r ApiGetKeywordGroupRecommendationsRequest) KeywordGroupsRecommendationsRequest(keywordGroupsRecommendationsRequest KeywordGroupsRecommendationsRequest) ApiGetKeywordGroupRecommendationsRequest {
	r.keywordGroupsRecommendationsRequest = &keywordGroupsRecommendationsRequest
	return r
}

func (r ApiGetKeywordGroupRecommendationsRequest) Execute() (*KeywordGroupsRecommendationsResponse, *http.Response, error) {
	return r.ApiService.GetKeywordGroupRecommendationsExecute(r)
}

func (a *KeywordGroupTargetingRecommendationsAPIService) GetKeywordGroupRecommendations(ctx context.Context) ApiGetKeywordGroupRecommendationsRequest {
	return ApiGetKeywordGroupRecommendationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KeywordGroupsRecommendationsResponse
func (a *KeywordGroupTargetingRecommendationsAPIService) GetKeywordGroupRecommendationsExecute(r ApiGetKeywordGroupRecommendationsRequest) (*KeywordGroupsRecommendationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *KeywordGroupsRecommendationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeywordGroupTargetingRecommendationsAPIService.GetKeywordGroupRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sp/targeting/recommendations/keywordGroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.spkeywordgroupsrecommendations.v1.0+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.spkeywordgroupsrecommendations.v1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.locale != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "locale", r.locale, "simple", "")
	}
	// body params
	localVarPostBody = r.keywordGroupsRecommendationsRequest
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
