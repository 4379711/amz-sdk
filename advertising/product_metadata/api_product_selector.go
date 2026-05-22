package product_metadata

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ProductSelectorAPIService ProductSelectorAPI service
type ProductSelectorAPIService service

type ApiProductMetadataRequest struct {
	ctx                          context.Context
	ApiService                   *ProductSelectorAPIService
	amazonAdvertisingAPIClientId *string
	productMetadataRequest       *ProductMetadataRequest
	amazonAdvertisingAPIScope    *string
	amazonAdsAccountId           *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiProductMetadataRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiProductMetadataRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiProductMetadataRequest) ProductMetadataRequest(productMetadataRequest ProductMetadataRequest) ApiProductMetadataRequest {
	r.productMetadataRequest = &productMetadataRequest
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiProductMetadataRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiProductMetadataRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

// Account identifier you use to access the DSP. This is your Amazon DSP Advertiser ID.
func (r ApiProductMetadataRequest) AmazonAdsAccountId(amazonAdsAccountId string) ApiProductMetadataRequest {
	r.amazonAdsAccountId = &amazonAdsAccountId
	return r
}

func (r ApiProductMetadataRequest) Execute() (*ProductMetadataResponse, *http.Response, error) {
	return r.ApiService.ProductMetadataExecute(r)
}

func (a *ProductSelectorAPIService) ProductMetadata(ctx context.Context) ApiProductMetadataRequest {
	return ApiProductMetadataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductMetadataResponse
func (a *ProductSelectorAPIService) ProductMetadataExecute(r ApiProductMetadataRequest) (*ProductMetadataResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductMetadataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductSelectorAPIService.ProductMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/product/metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productMetadataRequest == nil {
		return localVarReturnValue, nil, reportError("productMetadataRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.productmetadatarequest.v1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.productmetadataresponse.v1+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if r.amazonAdsAccountId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Amazon-Ads-AccountId", r.amazonAdsAccountId, "simple", "")
	}
	// body params
	localVarPostBody = r.productMetadataRequest
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
