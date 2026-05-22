package product_eligibility

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ProgramEligibilityAPIService ProgramEligibilityAPI service
type ProgramEligibilityAPIService service

type ApiProgramEligibilityRequest struct {
	ctx                              context.Context
	ApiService                       *ProgramEligibilityAPIService
	amazonAdvertisingAPIClientId     *string
	acceptLanguage                   *AcceptLanguage
	amazonAdsAccountId               *string
	amazonAdvertisingAPIScope        *string
	contentType                      *string
	programEligibilityRequestContent *ProgramEligibilityRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiProgramEligibilityRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiProgramEligibilityRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// Specify the language in which the response is returned.
func (r ApiProgramEligibilityRequest) AcceptLanguage(acceptLanguage AcceptLanguage) ApiProgramEligibilityRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// The header used to pass global account associated with the advertiser account Use &#x60;GET&#x60; method on the Global Ads Account resource to list the global ads account associated with the access token passed in the HTTP Authorization header and choose AdvertisingAccountIdentifier id from the response to pass it as input. Use for v2 global calls
func (r ApiProgramEligibilityRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiProgramEligibilityRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiProgramEligibilityRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiProgramEligibilityRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// The content type of the request.
func (r ApiProgramEligibilityRequest) ContentType(contentType string) ApiProgramEligibilityRequest {
	r.contentType = &contentType
	return r
}

func (r ApiProgramEligibilityRequest) ProgramEligibilityRequestContent(programEligibilityRequestContent ProgramEligibilityRequestContent) ApiProgramEligibilityRequest {
	r.programEligibilityRequestContent = &programEligibilityRequestContent
	return r
}

func (r ApiProgramEligibilityRequest) Execute() (*ProgramEligibilityResponseContent, *http.Response, error) {
	return r.ApiService.ProgramEligibilityExecute(r)
}

func (a *ProgramEligibilityAPIService) ProgramEligibility(ctx context.Context) ApiProgramEligibilityRequest {
	return ApiProgramEligibilityRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProgramEligibilityResponseContent
func (a *ProgramEligibilityAPIService) ProgramEligibilityExecute(r ApiProgramEligibilityRequest) (*ProgramEligibilityResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProgramEligibilityResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgramEligibilityAPIService.ProgramEligibility")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eligibility/programs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if strlen(*r.amazonAdvertisingAPIClientId) > 180 {
		return localVarReturnValue, nil, reportError("amazonAdvertisingAPIClientId must have less than 180 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/vnd.programeligibility.v2+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.programeligibility.v2+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "simple", "")
	}
	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	}

	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	}
	// body params
	localVarPostBody = r.programEligibilityRequestContent
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
