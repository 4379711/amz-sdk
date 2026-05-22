package tokens_20210301

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// TokensAPIService TokensAPI service
type TokensAPIService service

type ApiCreateRestrictedDataTokenRequest struct {
	ctx        context.Context
	ApiService *TokensAPIService
	body       *CreateRestrictedDataTokenRequest
}

// The restricted data token request details.
func (r ApiCreateRestrictedDataTokenRequest) Body(body CreateRestrictedDataTokenRequest) ApiCreateRestrictedDataTokenRequest {
	r.body = &body
	return r
}

func (r ApiCreateRestrictedDataTokenRequest) Execute() (*CreateRestrictedDataTokenResponse, *http.Response, error) {
	return r.ApiService.CreateRestrictedDataTokenExecute(r)
}

/*
CreateRestrictedDataToken Method for CreateRestrictedDataToken

Returns a Restricted Data Token (RDT) for one or more restricted resources that you specify. A restricted resource is the HTTP method and path from a restricted operation that returns Personally Identifiable Information (PII), plus a dataElements value that indicates the type of PII requested. See the Tokens API Use Case Guide for a list of restricted operations. Use the RDT returned here as the access token in subsequent calls to the corresponding restricted operations.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRestrictedDataTokenRequest
*/
func (a *TokensAPIService) CreateRestrictedDataToken(ctx context.Context) ApiCreateRestrictedDataTokenRequest {
	return ApiCreateRestrictedDataTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateRestrictedDataTokenResponse
func (a *TokensAPIService) CreateRestrictedDataTokenExecute(r ApiCreateRestrictedDataTokenRequest) (*CreateRestrictedDataTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateRestrictedDataTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensAPIService.CreateRestrictedDataToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokens/2021-03-01/restrictedDataToken"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
