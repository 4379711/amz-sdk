package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetGlobalProductRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGlobalProductRecommendationsRequest{}

// GetGlobalProductRecommendationsRequest Request structure to get ASIN recommendations for a set of input ASINs in a given marketplace
type GetGlobalProductRecommendationsRequest struct {
	// An optional token value that can be used to fetch previous set of records.
	PreviousToken *string `json:"previousToken,omitempty"`
	// The <b>maxResults</b> parameter in the request body determines the number of objects requested in the response. The count will be applied to the objects returned under the array belonging to each country under the `countryAdRecommendations` field in the response body. <ul> <li>When requesting the `application/vnd.spproductrecommendationresponse.asins.v3+json` media type, the count applies to the `GlobalProductRecommendation` objects in the response body. If no count value is provided, a default of `100` is assumed. The API will return a maximum of `1000` recommendations, regardless of how large the count is set.</li> </ul> Please refer to the response schemas for more information on the structure and fields of the response.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// An optional token value that can be used to fetch next set of records.
	NextToken *string `json:"nextToken,omitempty"`
	// Theme names and descriptions will be provided in the language for your supported locale. Available options are en_US (U.S. English), en_GB (UK English), zh_CN (Chinese), es_ES (Spanish), jp_JP (Japanese), de_DE (German), fr_FR (French), it_IT(Italian). If locale is not provided or unsupported, the theme names and descriptions will be returned in U.S. English (en_US).
	Locale *string `json:"locale,omitempty"`
	// A map containing the country code as the key and a list of advertised ASINs for that country as the value. Currently, only one country entry is supported in the request.
	CountryAdAsins map[string][]string `json:"countryAdAsins"`
}

type _GetGlobalProductRecommendationsRequest GetGlobalProductRecommendationsRequest

// NewGetGlobalProductRecommendationsRequest instantiates a new GetGlobalProductRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGlobalProductRecommendationsRequest(countryAdAsins map[string][]string) *GetGlobalProductRecommendationsRequest {
	this := GetGlobalProductRecommendationsRequest{}
	this.CountryAdAsins = countryAdAsins
	return &this
}

// NewGetGlobalProductRecommendationsRequestWithDefaults instantiates a new GetGlobalProductRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGlobalProductRecommendationsRequestWithDefaults() *GetGlobalProductRecommendationsRequest {
	this := GetGlobalProductRecommendationsRequest{}
	return &this
}

// GetPreviousToken returns the PreviousToken field value if set, zero value otherwise.
func (o *GetGlobalProductRecommendationsRequest) GetPreviousToken() string {
	if o == nil || IsNil(o.PreviousToken) {
		var ret string
		return ret
	}
	return *o.PreviousToken
}

// GetPreviousTokenOk returns a tuple with the PreviousToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGlobalProductRecommendationsRequest) GetPreviousTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousToken) {
		return nil, false
	}
	return o.PreviousToken, true
}

// HasPreviousToken returns a boolean if a field has been set.
func (o *GetGlobalProductRecommendationsRequest) HasPreviousToken() bool {
	if o != nil && !IsNil(o.PreviousToken) {
		return true
	}

	return false
}

// SetPreviousToken gets a reference to the given string and assigns it to the PreviousToken field.
func (o *GetGlobalProductRecommendationsRequest) SetPreviousToken(v string) {
	o.PreviousToken = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *GetGlobalProductRecommendationsRequest) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGlobalProductRecommendationsRequest) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *GetGlobalProductRecommendationsRequest) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *GetGlobalProductRecommendationsRequest) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetGlobalProductRecommendationsRequest) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGlobalProductRecommendationsRequest) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetGlobalProductRecommendationsRequest) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetGlobalProductRecommendationsRequest) SetNextToken(v string) {
	o.NextToken = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GetGlobalProductRecommendationsRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGlobalProductRecommendationsRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GetGlobalProductRecommendationsRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GetGlobalProductRecommendationsRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetCountryAdAsins returns the CountryAdAsins field value
func (o *GetGlobalProductRecommendationsRequest) GetCountryAdAsins() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.CountryAdAsins
}

// GetCountryAdAsinsOk returns a tuple with the CountryAdAsins field value
// and a boolean to check if the value has been set.
func (o *GetGlobalProductRecommendationsRequest) GetCountryAdAsinsOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryAdAsins, true
}

// SetCountryAdAsins sets field value
func (o *GetGlobalProductRecommendationsRequest) SetCountryAdAsins(v map[string][]string) {
	o.CountryAdAsins = v
}

func (o GetGlobalProductRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreviousToken) {
		toSerialize["previousToken"] = o.PreviousToken
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	toSerialize["countryAdAsins"] = o.CountryAdAsins
	return toSerialize, nil
}

type NullableGetGlobalProductRecommendationsRequest struct {
	value *GetGlobalProductRecommendationsRequest
	isSet bool
}

func (v NullableGetGlobalProductRecommendationsRequest) Get() *GetGlobalProductRecommendationsRequest {
	return v.value
}

func (v *NullableGetGlobalProductRecommendationsRequest) Set(val *GetGlobalProductRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGlobalProductRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGlobalProductRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGlobalProductRecommendationsRequest(val *GetGlobalProductRecommendationsRequest) *NullableGetGlobalProductRecommendationsRequest {
	return &NullableGetGlobalProductRecommendationsRequest{value: val, isSet: true}
}

func (v NullableGetGlobalProductRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetGlobalProductRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
