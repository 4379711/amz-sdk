package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryBidRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryBidRecommendationError{}

// MultiCountryBidRecommendationError struct for MultiCountryBidRecommendationError
type MultiCountryBidRecommendationError struct {
	// Countries where error have occurred
	CountryCodes []string `json:"countryCodes,omitempty"`
	// Machine readable error code.
	Code *string `json:"code,omitempty"`
	// Human readable 1 liner error message
	Message *string `json:"message,omitempty"`
}

// NewMultiCountryBidRecommendationError instantiates a new MultiCountryBidRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryBidRecommendationError() *MultiCountryBidRecommendationError {
	this := MultiCountryBidRecommendationError{}
	return &this
}

// NewMultiCountryBidRecommendationErrorWithDefaults instantiates a new MultiCountryBidRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryBidRecommendationErrorWithDefaults() *MultiCountryBidRecommendationError {
	this := MultiCountryBidRecommendationError{}
	return &this
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *MultiCountryBidRecommendationError) GetCountryCodes() []string {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []string
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCountryBidRecommendationError) GetCountryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *MultiCountryBidRecommendationError) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []string and assigns it to the CountryCodes field.
func (o *MultiCountryBidRecommendationError) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MultiCountryBidRecommendationError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCountryBidRecommendationError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MultiCountryBidRecommendationError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *MultiCountryBidRecommendationError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MultiCountryBidRecommendationError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiCountryBidRecommendationError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MultiCountryBidRecommendationError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MultiCountryBidRecommendationError) SetMessage(v string) {
	o.Message = &v
}

func (o MultiCountryBidRecommendationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCodes) {
		toSerialize["countryCodes"] = o.CountryCodes
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableMultiCountryBidRecommendationError struct {
	value *MultiCountryBidRecommendationError
	isSet bool
}

func (v NullableMultiCountryBidRecommendationError) Get() *MultiCountryBidRecommendationError {
	return v.value
}

func (v *NullableMultiCountryBidRecommendationError) Set(val *MultiCountryBidRecommendationError) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryBidRecommendationError) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryBidRecommendationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryBidRecommendationError(val *MultiCountryBidRecommendationError) *NullableMultiCountryBidRecommendationError {
	return &NullableMultiCountryBidRecommendationError{value: val, isSet: true}
}

func (v NullableMultiCountryBidRecommendationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryBidRecommendationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
