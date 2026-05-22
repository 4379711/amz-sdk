package sb_v4

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// OptimizationRulesAPIService OptimizationRulesAPI service
type OptimizationRulesAPIService service

type ApiAssociateSponsoredBrandsOptimizationRulesRequest struct {
	ctx                                                     context.Context
	ApiService                                              *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                            *string
	amazonAdvertisingAPIScope                               *string
	associateSponsoredBrandsOptimizationRulesRequestContent *AssociateSponsoredBrandsOptimizationRulesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiAssociateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiAssociateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiAssociateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiAssociateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiAssociateSponsoredBrandsOptimizationRulesRequest) AssociateSponsoredBrandsOptimizationRulesRequestContent(associateSponsoredBrandsOptimizationRulesRequestContent AssociateSponsoredBrandsOptimizationRulesRequestContent) ApiAssociateSponsoredBrandsOptimizationRulesRequest {
	r.associateSponsoredBrandsOptimizationRulesRequestContent = &associateSponsoredBrandsOptimizationRulesRequestContent
	return r
}

func (r ApiAssociateSponsoredBrandsOptimizationRulesRequest) Execute() (*AssociateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	return r.ApiService.AssociateSponsoredBrandsOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) AssociateSponsoredBrandsOptimizationRules(ctx context.Context) ApiAssociateSponsoredBrandsOptimizationRulesRequest {
	return ApiAssociateSponsoredBrandsOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AssociateSponsoredBrandsOptimizationRulesResponseContent
func (a *OptimizationRulesAPIService) AssociateSponsoredBrandsOptimizationRulesExecute(r ApiAssociateSponsoredBrandsOptimizationRulesRequest) (*AssociateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AssociateSponsoredBrandsOptimizationRulesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.AssociateSponsoredBrandsOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/rules/optimization/associate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.associateSponsoredBrandsOptimizationRulesRequestContent == nil {
		return localVarReturnValue, nil, reportError("associateSponsoredBrandsOptimizationRulesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbruleoptimization.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbruleoptimization.v4+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.associateSponsoredBrandsOptimizationRulesRequestContent
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

type ApiCreateSponsoredBrandsOptimizationRulesRequest struct {
	ctx                                                  context.Context
	ApiService                                           *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                         *string
	amazonAdvertisingAPIScope                            *string
	createSponsoredBrandsOptimizationRulesRequestContent *CreateSponsoredBrandsOptimizationRulesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiCreateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiCreateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiCreateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiCreateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiCreateSponsoredBrandsOptimizationRulesRequest) CreateSponsoredBrandsOptimizationRulesRequestContent(createSponsoredBrandsOptimizationRulesRequestContent CreateSponsoredBrandsOptimizationRulesRequestContent) ApiCreateSponsoredBrandsOptimizationRulesRequest {
	r.createSponsoredBrandsOptimizationRulesRequestContent = &createSponsoredBrandsOptimizationRulesRequestContent
	return r
}

func (r ApiCreateSponsoredBrandsOptimizationRulesRequest) Execute() (*CreateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	return r.ApiService.CreateSponsoredBrandsOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) CreateSponsoredBrandsOptimizationRules(ctx context.Context) ApiCreateSponsoredBrandsOptimizationRulesRequest {
	return ApiCreateSponsoredBrandsOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSponsoredBrandsOptimizationRulesResponseContent
func (a *OptimizationRulesAPIService) CreateSponsoredBrandsOptimizationRulesExecute(r ApiCreateSponsoredBrandsOptimizationRulesRequest) (*CreateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateSponsoredBrandsOptimizationRulesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.CreateSponsoredBrandsOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/rules/optimization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createSponsoredBrandsOptimizationRulesRequestContent == nil {
		return localVarReturnValue, nil, reportError("createSponsoredBrandsOptimizationRulesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbruleoptimization.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbruleoptimization.v4+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.createSponsoredBrandsOptimizationRulesRequestContent
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

type ApiDisassociateSponsoredBrandsOptimizationRulesRequest struct {
	ctx                                                        context.Context
	ApiService                                                 *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                               *string
	amazonAdvertisingAPIScope                                  *string
	disassociateSponsoredBrandsOptimizationRulesRequestContent *DisassociateSponsoredBrandsOptimizationRulesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiDisassociateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiDisassociateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiDisassociateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiDisassociateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiDisassociateSponsoredBrandsOptimizationRulesRequest) DisassociateSponsoredBrandsOptimizationRulesRequestContent(disassociateSponsoredBrandsOptimizationRulesRequestContent DisassociateSponsoredBrandsOptimizationRulesRequestContent) ApiDisassociateSponsoredBrandsOptimizationRulesRequest {
	r.disassociateSponsoredBrandsOptimizationRulesRequestContent = &disassociateSponsoredBrandsOptimizationRulesRequestContent
	return r
}

func (r ApiDisassociateSponsoredBrandsOptimizationRulesRequest) Execute() (*DisassociateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	return r.ApiService.DisassociateSponsoredBrandsOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) DisassociateSponsoredBrandsOptimizationRules(ctx context.Context) ApiDisassociateSponsoredBrandsOptimizationRulesRequest {
	return ApiDisassociateSponsoredBrandsOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DisassociateSponsoredBrandsOptimizationRulesResponseContent
func (a *OptimizationRulesAPIService) DisassociateSponsoredBrandsOptimizationRulesExecute(r ApiDisassociateSponsoredBrandsOptimizationRulesRequest) (*DisassociateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DisassociateSponsoredBrandsOptimizationRulesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.DisassociateSponsoredBrandsOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/rules/optimization/disassociate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.disassociateSponsoredBrandsOptimizationRulesRequestContent == nil {
		return localVarReturnValue, nil, reportError("disassociateSponsoredBrandsOptimizationRulesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbruleoptimization.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbruleoptimization.v4+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.disassociateSponsoredBrandsOptimizationRulesRequestContent
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

type ApiListSponsoredBrandsOptimizationRulesRequest struct {
	ctx                                                context.Context
	ApiService                                         *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                       *string
	amazonAdvertisingAPIScope                          *string
	listSponsoredBrandsOptimizationRulesRequestContent *ListSponsoredBrandsOptimizationRulesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiListSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiListSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiListSponsoredBrandsOptimizationRulesRequest) ListSponsoredBrandsOptimizationRulesRequestContent(listSponsoredBrandsOptimizationRulesRequestContent ListSponsoredBrandsOptimizationRulesRequestContent) ApiListSponsoredBrandsOptimizationRulesRequest {
	r.listSponsoredBrandsOptimizationRulesRequestContent = &listSponsoredBrandsOptimizationRulesRequestContent
	return r
}

func (r ApiListSponsoredBrandsOptimizationRulesRequest) Execute() (*ListSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	return r.ApiService.ListSponsoredBrandsOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) ListSponsoredBrandsOptimizationRules(ctx context.Context) ApiListSponsoredBrandsOptimizationRulesRequest {
	return ApiListSponsoredBrandsOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListSponsoredBrandsOptimizationRulesResponseContent
func (a *OptimizationRulesAPIService) ListSponsoredBrandsOptimizationRulesExecute(r ApiListSponsoredBrandsOptimizationRulesRequest) (*ListSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListSponsoredBrandsOptimizationRulesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.ListSponsoredBrandsOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/rules/optimization/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbruleoptimization.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbruleoptimization.v4+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.listSponsoredBrandsOptimizationRulesRequestContent
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

type ApiUpdateSponsoredBrandsOptimizationRulesRequest struct {
	ctx                                                  context.Context
	ApiService                                           *OptimizationRulesAPIService
	amazonAdvertisingAPIClientId                         *string
	amazonAdvertisingAPIScope                            *string
	updateSponsoredBrandsOptimizationRulesRequestContent *UpdateSponsoredBrandsOptimizationRulesRequestContent
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input.
func (r ApiUpdateSponsoredBrandsOptimizationRulesRequest) AmazonAdvertisingAPIScope(amazonAdvertisingAPIScope string) ApiUpdateSponsoredBrandsOptimizationRulesRequest {
	r.amazonAdvertisingAPIScope = &amazonAdvertisingAPIScope
	return r
}

func (r ApiUpdateSponsoredBrandsOptimizationRulesRequest) UpdateSponsoredBrandsOptimizationRulesRequestContent(updateSponsoredBrandsOptimizationRulesRequestContent UpdateSponsoredBrandsOptimizationRulesRequestContent) ApiUpdateSponsoredBrandsOptimizationRulesRequest {
	r.updateSponsoredBrandsOptimizationRulesRequestContent = &updateSponsoredBrandsOptimizationRulesRequestContent
	return r
}

func (r ApiUpdateSponsoredBrandsOptimizationRulesRequest) Execute() (*UpdateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	return r.ApiService.UpdateSponsoredBrandsOptimizationRulesExecute(r)
}

func (a *OptimizationRulesAPIService) UpdateSponsoredBrandsOptimizationRules(ctx context.Context) ApiUpdateSponsoredBrandsOptimizationRulesRequest {
	return ApiUpdateSponsoredBrandsOptimizationRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateSponsoredBrandsOptimizationRulesResponseContent
func (a *OptimizationRulesAPIService) UpdateSponsoredBrandsOptimizationRulesExecute(r ApiUpdateSponsoredBrandsOptimizationRulesRequest) (*UpdateSponsoredBrandsOptimizationRulesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSponsoredBrandsOptimizationRulesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OptimizationRulesAPIService.UpdateSponsoredBrandsOptimizationRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sb/rules/optimization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.updateSponsoredBrandsOptimizationRulesRequestContent == nil {
		return localVarReturnValue, nil, reportError("updateSponsoredBrandsOptimizationRulesRequestContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.sbruleoptimization.v4+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.sbruleoptimization.v4+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.updateSponsoredBrandsOptimizationRulesRequestContent
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
