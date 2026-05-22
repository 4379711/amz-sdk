package sd_v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ForecastsAPIService ForecastsAPI service
type ForecastsAPIService service

type ApiCreateSDForecastRequest struct {
	ctx                          context.Context
	ApiService                   *ForecastsAPIService
	amazonAdvertisingAPIClientId *string
	amazonAdvertisingAPIScope    *string
	sDForecastRequest            *SDForecastRequest
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSDForecastRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSDForecastRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
func (r ApiCreateSDForecastRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSDForecastRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSDForecastRequest) SDForecastRequest(sDForecastRequest SDForecastRequest) ApiCreateSDForecastRequest {
	r.sDForecastRequest = &sDForecastRequest
	return r
}

func (r ApiCreateSDForecastRequest) Execute() (*SDForecastResponse, *http.Response, error) {
	return r.ApiService.CreateSDForecastExecute(r)
}

func (a *ForecastsAPIService) CreateSDForecast(ctx context.Context) ApiCreateSDForecastRequest {
	return ApiCreateSDForecastRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SDForecastResponse
func (a *ForecastsAPIService) CreateSDForecastExecute(r ApiCreateSDForecastRequest) (*SDForecastResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SDForecastResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ForecastsAPIService.CreateSDForecast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sd/forecasts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sdforecasts.v3.1+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sdforecasts.v3.1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.sDForecastRequest
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
