package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftAdGroupAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroupAccessError{}

// SponsoredProductsDraftAdGroupAccessError struct for SponsoredProductsDraftAdGroupAccessError
type SponsoredProductsDraftAdGroupAccessError struct {
	// The type of the error
	ErrorType  string                                           `json:"errorType"`
	ErrorValue SponsoredProductsDraftAdGroupAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftAdGroupAccessError SponsoredProductsDraftAdGroupAccessError

// NewSponsoredProductsDraftAdGroupAccessError instantiates a new SponsoredProductsDraftAdGroupAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroupAccessError(errorType string, errorValue SponsoredProductsDraftAdGroupAccessErrorSelector) *SponsoredProductsDraftAdGroupAccessError {
	this := SponsoredProductsDraftAdGroupAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftAdGroupAccessErrorWithDefaults instantiates a new SponsoredProductsDraftAdGroupAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupAccessErrorWithDefaults() *SponsoredProductsDraftAdGroupAccessError {
	this := SponsoredProductsDraftAdGroupAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftAdGroupAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftAdGroupAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftAdGroupAccessError) GetErrorValue() SponsoredProductsDraftAdGroupAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftAdGroupAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupAccessError) GetErrorValueOk() (*SponsoredProductsDraftAdGroupAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftAdGroupAccessError) SetErrorValue(v SponsoredProductsDraftAdGroupAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftAdGroupAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftAdGroupAccessError struct {
	value *SponsoredProductsDraftAdGroupAccessError
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroupAccessError) Get() *SponsoredProductsDraftAdGroupAccessError {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroupAccessError) Set(val *SponsoredProductsDraftAdGroupAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroupAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroupAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroupAccessError(val *SponsoredProductsDraftAdGroupAccessError) *NullableSponsoredProductsDraftAdGroupAccessError {
	return &NullableSponsoredProductsDraftAdGroupAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroupAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroupAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
