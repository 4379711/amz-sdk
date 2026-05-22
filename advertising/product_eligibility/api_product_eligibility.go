package product_eligibility

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ProductEligibilityAPIService ProductEligibilityAPI service
type ProductEligibilityAPIService service

type ApiProductEligibilityRequest struct {
	ctx                          context.Context
	ApiService                   *ProductEligibilityAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	productEligibilityRequest    *ProductEligibilityRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiProductEligibilityRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiProductEligibilityRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiProductEligibilityRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiProductEligibilityRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Request for Product Eligibility
func (r ApiProductEligibilityRequest) ProductEligibilityRequest(productEligibilityRequest ProductEligibilityRequest) ApiProductEligibilityRequest {
	r.productEligibilityRequest = &productEligibilityRequest
	return r
}

func (r ApiProductEligibilityRequest) Execute() (*ProductEligibilityResponse, *http.Response, error) {
	return r.ApiService.ProductEligibilityExecute(r)
}

func (a *ProductEligibilityAPIService) ProductEligibility(ctx context.Context) ApiProductEligibilityRequest {
	return ApiProductEligibilityRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductEligibilityResponse
func (a *ProductEligibilityAPIService) ProductEligibilityExecute(r ApiProductEligibilityRequest) (*ProductEligibilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductEligibilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductEligibilityAPIService.ProductEligibility")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eligibility/product/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productEligibilityRequest == nil {
		return localVarReturnValue, nil, reportError("productEligibilityRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.productEligibilityRequest
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
