package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SearchGlobalBrandsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchGlobalBrandsRequest{}

// SearchGlobalBrandsRequest Request object for SearchGlobalBrands API.
type SearchGlobalBrandsRequest struct {
	// Map containing keywords for countries passed in the request.
	CountryKeyword map[string]string `json:"countryKeyword"`
}

type _SearchGlobalBrandsRequest SearchGlobalBrandsRequest

// NewSearchGlobalBrandsRequest instantiates a new SearchGlobalBrandsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchGlobalBrandsRequest(countryKeyword map[string]string) *SearchGlobalBrandsRequest {
	this := SearchGlobalBrandsRequest{}
	this.CountryKeyword = countryKeyword
	return &this
}

// NewSearchGlobalBrandsRequestWithDefaults instantiates a new SearchGlobalBrandsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchGlobalBrandsRequestWithDefaults() *SearchGlobalBrandsRequest {
	this := SearchGlobalBrandsRequest{}
	return &this
}

// GetCountryKeyword returns the CountryKeyword field value
func (o *SearchGlobalBrandsRequest) GetCountryKeyword() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.CountryKeyword
}

// GetCountryKeywordOk returns a tuple with the CountryKeyword field value
// and a boolean to check if the value has been set.
func (o *SearchGlobalBrandsRequest) GetCountryKeywordOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryKeyword, true
}

// SetCountryKeyword sets field value
func (o *SearchGlobalBrandsRequest) SetCountryKeyword(v map[string]string) {
	o.CountryKeyword = v
}

func (o SearchGlobalBrandsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryKeyword"] = o.CountryKeyword
	return toSerialize, nil
}

type NullableSearchGlobalBrandsRequest struct {
	value *SearchGlobalBrandsRequest
	isSet bool
}

func (v NullableSearchGlobalBrandsRequest) Get() *SearchGlobalBrandsRequest {
	return v.value
}

func (v *NullableSearchGlobalBrandsRequest) Set(val *SearchGlobalBrandsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchGlobalBrandsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchGlobalBrandsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchGlobalBrandsRequest(val *SearchGlobalBrandsRequest) *NullableSearchGlobalBrandsRequest {
	return &NullableSearchGlobalBrandsRequest{value: val, isSet: true}
}

func (v NullableSearchGlobalBrandsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSearchGlobalBrandsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
