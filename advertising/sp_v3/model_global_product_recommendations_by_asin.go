package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductRecommendationsByASIN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductRecommendationsByASIN{}

// GlobalProductRecommendationsByASIN Global Product recommendations supplemented with relevant information.
type GlobalProductRecommendationsByASIN struct {
	// A map containing recommended ASINs and their associated themes, grouped by country.
	CountryAdRecommendations *map[string][]GlobalProductRecommendation `json:"countryAdRecommendations,omitempty"`
	// Optional parameter that links to the previous result set served. This parameter will be null on the first request.
	PreviousToken *string `json:"previousToken,omitempty"`
	// An identifier to fetch next set of `GlobalProductRecommendation` records in the result set if available. This will be null when at the end of result set.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewGlobalProductRecommendationsByASIN instantiates a new GlobalProductRecommendationsByASIN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductRecommendationsByASIN() *GlobalProductRecommendationsByASIN {
	this := GlobalProductRecommendationsByASIN{}
	return &this
}

// NewGlobalProductRecommendationsByASINWithDefaults instantiates a new GlobalProductRecommendationsByASIN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductRecommendationsByASINWithDefaults() *GlobalProductRecommendationsByASIN {
	this := GlobalProductRecommendationsByASIN{}
	return &this
}

// GetCountryAdRecommendations returns the CountryAdRecommendations field value if set, zero value otherwise.
func (o *GlobalProductRecommendationsByASIN) GetCountryAdRecommendations() map[string][]GlobalProductRecommendation {
	if o == nil || IsNil(o.CountryAdRecommendations) {
		var ret map[string][]GlobalProductRecommendation
		return ret
	}
	return *o.CountryAdRecommendations
}

// GetCountryAdRecommendationsOk returns a tuple with the CountryAdRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductRecommendationsByASIN) GetCountryAdRecommendationsOk() (*map[string][]GlobalProductRecommendation, bool) {
	if o == nil || IsNil(o.CountryAdRecommendations) {
		return nil, false
	}
	return o.CountryAdRecommendations, true
}

// HasCountryAdRecommendations returns a boolean if a field has been set.
func (o *GlobalProductRecommendationsByASIN) HasCountryAdRecommendations() bool {
	if o != nil && !IsNil(o.CountryAdRecommendations) {
		return true
	}

	return false
}

// SetCountryAdRecommendations gets a reference to the given map[string][]GlobalProductRecommendation and assigns it to the CountryAdRecommendations field.
func (o *GlobalProductRecommendationsByASIN) SetCountryAdRecommendations(v map[string][]GlobalProductRecommendation) {
	o.CountryAdRecommendations = &v
}

// GetPreviousToken returns the PreviousToken field value if set, zero value otherwise.
func (o *GlobalProductRecommendationsByASIN) GetPreviousToken() string {
	if o == nil || IsNil(o.PreviousToken) {
		var ret string
		return ret
	}
	return *o.PreviousToken
}

// GetPreviousTokenOk returns a tuple with the PreviousToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductRecommendationsByASIN) GetPreviousTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousToken) {
		return nil, false
	}
	return o.PreviousToken, true
}

// HasPreviousToken returns a boolean if a field has been set.
func (o *GlobalProductRecommendationsByASIN) HasPreviousToken() bool {
	if o != nil && !IsNil(o.PreviousToken) {
		return true
	}

	return false
}

// SetPreviousToken gets a reference to the given string and assigns it to the PreviousToken field.
func (o *GlobalProductRecommendationsByASIN) SetPreviousToken(v string) {
	o.PreviousToken = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GlobalProductRecommendationsByASIN) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductRecommendationsByASIN) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GlobalProductRecommendationsByASIN) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GlobalProductRecommendationsByASIN) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GlobalProductRecommendationsByASIN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryAdRecommendations) {
		toSerialize["countryAdRecommendations"] = o.CountryAdRecommendations
	}
	if !IsNil(o.PreviousToken) {
		toSerialize["previousToken"] = o.PreviousToken
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGlobalProductRecommendationsByASIN struct {
	value *GlobalProductRecommendationsByASIN
	isSet bool
}

func (v NullableGlobalProductRecommendationsByASIN) Get() *GlobalProductRecommendationsByASIN {
	return v.value
}

func (v *NullableGlobalProductRecommendationsByASIN) Set(val *GlobalProductRecommendationsByASIN) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductRecommendationsByASIN) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductRecommendationsByASIN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductRecommendationsByASIN(val *GlobalProductRecommendationsByASIN) *NullableGlobalProductRecommendationsByASIN {
	return &NullableGlobalProductRecommendationsByASIN{value: val, isSet: true}
}

func (v NullableGlobalProductRecommendationsByASIN) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductRecommendationsByASIN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
