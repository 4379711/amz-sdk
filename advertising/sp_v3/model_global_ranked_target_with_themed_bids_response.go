package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalRankedTargetWithThemedBidsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalRankedTargetWithThemedBidsResponse{}

// GlobalRankedTargetWithThemedBidsResponse struct for GlobalRankedTargetWithThemedBidsResponse
type GlobalRankedTargetWithThemedBidsResponse struct {
	CountryCodes *map[string]RankedTargetWithThemedBidsResponseWithError `json:"countryCodes,omitempty"`
}

// NewGlobalRankedTargetWithThemedBidsResponse instantiates a new GlobalRankedTargetWithThemedBidsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRankedTargetWithThemedBidsResponse() *GlobalRankedTargetWithThemedBidsResponse {
	this := GlobalRankedTargetWithThemedBidsResponse{}
	return &this
}

// NewGlobalRankedTargetWithThemedBidsResponseWithDefaults instantiates a new GlobalRankedTargetWithThemedBidsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRankedTargetWithThemedBidsResponseWithDefaults() *GlobalRankedTargetWithThemedBidsResponse {
	this := GlobalRankedTargetWithThemedBidsResponse{}
	return &this
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *GlobalRankedTargetWithThemedBidsResponse) GetCountryCodes() map[string]RankedTargetWithThemedBidsResponseWithError {
	if o == nil || IsNil(o.CountryCodes) {
		var ret map[string]RankedTargetWithThemedBidsResponseWithError
		return ret
	}
	return *o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedTargetWithThemedBidsResponse) GetCountryCodesOk() (*map[string]RankedTargetWithThemedBidsResponseWithError, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *GlobalRankedTargetWithThemedBidsResponse) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given map[string]RankedTargetWithThemedBidsResponseWithError and assigns it to the CountryCodes field.
func (o *GlobalRankedTargetWithThemedBidsResponse) SetCountryCodes(v map[string]RankedTargetWithThemedBidsResponseWithError) {
	o.CountryCodes = &v
}

func (o GlobalRankedTargetWithThemedBidsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCodes) {
		toSerialize["countryCodes"] = o.CountryCodes
	}
	return toSerialize, nil
}

type NullableGlobalRankedTargetWithThemedBidsResponse struct {
	value *GlobalRankedTargetWithThemedBidsResponse
	isSet bool
}

func (v NullableGlobalRankedTargetWithThemedBidsResponse) Get() *GlobalRankedTargetWithThemedBidsResponse {
	return v.value
}

func (v *NullableGlobalRankedTargetWithThemedBidsResponse) Set(val *GlobalRankedTargetWithThemedBidsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRankedTargetWithThemedBidsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRankedTargetWithThemedBidsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRankedTargetWithThemedBidsResponse(val *GlobalRankedTargetWithThemedBidsResponse) *NullableGlobalRankedTargetWithThemedBidsResponse {
	return &NullableGlobalRankedTargetWithThemedBidsResponse{value: val, isSet: true}
}

func (v NullableGlobalRankedTargetWithThemedBidsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalRankedTargetWithThemedBidsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
