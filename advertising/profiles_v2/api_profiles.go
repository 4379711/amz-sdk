package profiles_v2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ProfilesAPIService ProfilesAPI service
type ProfilesAPIService service

type ApiGetProfileByIdRequest struct {
	ctx                          context.Context
	ApiService                   *ProfilesAPIService
	amazonAdvertisingAPIClientId *string
	profileId                    int64
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiGetProfileByIdRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiGetProfileByIdRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiGetProfileByIdRequest) Execute() (*Profile, *http.Response, error) {
	return r.ApiService.GetProfileByIdExecute(r)
}

func (a *ProfilesAPIService) GetProfileById(ctx context.Context, profileId int64) ApiGetProfileByIdRequest {
	return ApiGetProfileByIdRequest{
		ApiService: a,
		ctx:        ctx,
		profileId:  profileId,
	}
}

// Execute executes the request
//
//	@return Profile
func (a *ProfilesAPIService) GetProfileByIdExecute(r ApiGetProfileByIdRequest) (*Profile, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Profile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfilesAPIService.GetProfileById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/profiles/{profileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"profileId"+"}", url.PathEscape(parameterValueToString(r.profileId, "profileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiListProfilesRequest struct {
	ctx                          context.Context
	ApiService                   *ProfilesAPIService
	amazonAdvertisingAPIClientId *string
	apiProgram                   *string
	accessLevel                  *string
	profileTypeFilter            *string
	validPaymentMethodFilter     *string
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiListProfilesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiListProfilesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

// Filters response to include profiles that have permissions for the specified Advertising API program only. Setting &#x60;apiProgram&#x3D;billing&#x60; filters the response to include only profiles to which the user and application associated with the access token have permission to view or edit billing information.
func (r ApiListProfilesRequest) ApiProgram(apiProgram string) ApiListProfilesRequest {
	r.apiProgram = &apiProgram
	return r
}

// Filters response to include profiles that have specified permissions for the specified Advertising API program only. Currently, the only supported access level is &#x60;view&#x60; and &#x60;edit&#x60;. Setting &#x60;accessLevel&#x3D;view&#x60; filters the response to include only profiles to which the user and application associated with the access token have view permission to the provided api program.
func (r ApiListProfilesRequest) AccessLevel(accessLevel string) ApiListProfilesRequest {
	r.accessLevel = &accessLevel
	return r
}

// Filters response to include profiles that are of the specified types in the comma-delimited list. Default is all types. Note that this filter performs an inclusive AND operation on the types.
func (r ApiListProfilesRequest) ProfileTypeFilter(profileTypeFilter string) ApiListProfilesRequest {
	r.profileTypeFilter = &profileTypeFilter
	return r
}

// Filter response to include profiles that have valid payment methods. Default is to include all profiles. Setting this filter to &#x60;true&#x60; returns only profiles with either no &#x60;validPaymentMethod&#x60; field, or the &#x60;validPaymentMethod&#x60; field set to &#x60;true&#x60;.  Setting this to &#x60;false&#x60; returns profiles with the &#x60;validPaymentMethod&#x60; field set to &#x60;false&#x60; only.
func (r ApiListProfilesRequest) ValidPaymentMethodFilter(validPaymentMethodFilter string) ApiListProfilesRequest {
	r.validPaymentMethodFilter = &validPaymentMethodFilter
	return r
}

func (r ApiListProfilesRequest) Execute() ([]Profile, *http.Response, error) {
	return r.ApiService.ListProfilesExecute(r)
}

func (a *ProfilesAPIService) ListProfiles(ctx context.Context) ApiListProfilesRequest {
	return ApiListProfilesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Profile
func (a *ProfilesAPIService) ListProfilesExecute(r ApiListProfilesRequest) ([]Profile, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Profile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfilesAPIService.ListProfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.apiProgram != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apiProgram", r.apiProgram, "form", "")
	} else {
		var defaultValue string = "campaign"
		r.apiProgram = &defaultValue
	}
	if r.accessLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accessLevel", r.accessLevel, "form", "")
	} else {
		var defaultValue string = "edit"
		r.accessLevel = &defaultValue
	}
	if r.profileTypeFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "profileTypeFilter", r.profileTypeFilter, "form", "")
	}
	if r.validPaymentMethodFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validPaymentMethodFilter", r.validPaymentMethodFilter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiUpdateProfilesRequest struct {
	ctx                          context.Context
	ApiService                   *ProfilesAPIService
	amazonAdvertisingAPIClientId *string
	profile                      *[]Profile
}

// The identifier of a client associated with a \&quot;Login with Amazon\&quot; account.
func (r ApiUpdateProfilesRequest) AmazonAdvertisingAPIClientId(amazonAdvertisingAPIClientId string) ApiUpdateProfilesRequest {
	r.amazonAdvertisingAPIClientId = &amazonAdvertisingAPIClientId
	return r
}

func (r ApiUpdateProfilesRequest) Profile(profile []Profile) ApiUpdateProfilesRequest {
	r.profile = &profile
	return r
}

func (r ApiUpdateProfilesRequest) Execute() ([]ProfileResponse, *http.Response, error) {
	return r.ApiService.UpdateProfilesExecute(r)
}

func (a *ProfilesAPIService) UpdateProfiles(ctx context.Context) ApiUpdateProfilesRequest {
	return ApiUpdateProfilesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ProfileResponse
func (a *ProfilesAPIService) UpdateProfilesExecute(r ApiUpdateProfilesRequest) ([]ProfileResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ProfileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfilesAPIService.UpdateProfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.profile
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
