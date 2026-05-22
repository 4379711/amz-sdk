package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalProductEligibilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductEligibilityResponse{}

// GlobalProductEligibilityResponse A multi-country product advertising eligibility response object.
type GlobalProductEligibilityResponse struct {
	// A list of country and the county's corresponding product advertising eligibility responses.
	CountryProductResponseList []CountryProductResponse `json:"countryProductResponseList,omitempty"`
}

// NewGlobalProductEligibilityResponse instantiates a new GlobalProductEligibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductEligibilityResponse() *GlobalProductEligibilityResponse {
	this := GlobalProductEligibilityResponse{}
	return &this
}

// NewGlobalProductEligibilityResponseWithDefaults instantiates a new GlobalProductEligibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductEligibilityResponseWithDefaults() *GlobalProductEligibilityResponse {
	this := GlobalProductEligibilityResponse{}
	return &this
}

// GetCountryProductResponseList returns the CountryProductResponseList field value if set, zero value otherwise.
func (o *GlobalProductEligibilityResponse) GetCountryProductResponseList() []CountryProductResponse {
	if o == nil || IsNil(o.CountryProductResponseList) {
		var ret []CountryProductResponse
		return ret
	}
	return o.CountryProductResponseList
}

// GetCountryProductResponseListOk returns a tuple with the CountryProductResponseList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductEligibilityResponse) GetCountryProductResponseListOk() ([]CountryProductResponse, bool) {
	if o == nil || IsNil(o.CountryProductResponseList) {
		return nil, false
	}
	return o.CountryProductResponseList, true
}

// HasCountryProductResponseList returns a boolean if a field has been set.
func (o *GlobalProductEligibilityResponse) HasCountryProductResponseList() bool {
	if o != nil && !IsNil(o.CountryProductResponseList) {
		return true
	}

	return false
}

// SetCountryProductResponseList gets a reference to the given []CountryProductResponse and assigns it to the CountryProductResponseList field.
func (o *GlobalProductEligibilityResponse) SetCountryProductResponseList(v []CountryProductResponse) {
	o.CountryProductResponseList = v
}

func (o GlobalProductEligibilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryProductResponseList) {
		toSerialize["countryProductResponseList"] = o.CountryProductResponseList
	}
	return toSerialize, nil
}

type NullableGlobalProductEligibilityResponse struct {
	value *GlobalProductEligibilityResponse
	isSet bool
}

func (v NullableGlobalProductEligibilityResponse) Get() *GlobalProductEligibilityResponse {
	return v.value
}

func (v *NullableGlobalProductEligibilityResponse) Set(val *GlobalProductEligibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductEligibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductEligibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductEligibilityResponse(val *GlobalProductEligibilityResponse) *NullableGlobalProductEligibilityResponse {
	return &NullableGlobalProductEligibilityResponse{value: val, isSet: true}
}

func (v NullableGlobalProductEligibilityResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductEligibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
