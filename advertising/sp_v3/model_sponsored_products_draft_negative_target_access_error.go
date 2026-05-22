package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeTargetAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeTargetAccessError{}

// SponsoredProductsDraftNegativeTargetAccessError struct for SponsoredProductsDraftNegativeTargetAccessError
type SponsoredProductsDraftNegativeTargetAccessError struct {
	// The type of the error
	ErrorType  string                                                  `json:"errorType"`
	ErrorValue SponsoredProductsDraftNegativeTargetAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftNegativeTargetAccessError SponsoredProductsDraftNegativeTargetAccessError

// NewSponsoredProductsDraftNegativeTargetAccessError instantiates a new SponsoredProductsDraftNegativeTargetAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeTargetAccessError(errorType string, errorValue SponsoredProductsDraftNegativeTargetAccessErrorSelector) *SponsoredProductsDraftNegativeTargetAccessError {
	this := SponsoredProductsDraftNegativeTargetAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftNegativeTargetAccessErrorWithDefaults instantiates a new SponsoredProductsDraftNegativeTargetAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeTargetAccessErrorWithDefaults() *SponsoredProductsDraftNegativeTargetAccessError {
	this := SponsoredProductsDraftNegativeTargetAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftNegativeTargetAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftNegativeTargetAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftNegativeTargetAccessError) GetErrorValue() SponsoredProductsDraftNegativeTargetAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftNegativeTargetAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetAccessError) GetErrorValueOk() (*SponsoredProductsDraftNegativeTargetAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftNegativeTargetAccessError) SetErrorValue(v SponsoredProductsDraftNegativeTargetAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftNegativeTargetAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeTargetAccessError struct {
	value *SponsoredProductsDraftNegativeTargetAccessError
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeTargetAccessError) Get() *SponsoredProductsDraftNegativeTargetAccessError {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeTargetAccessError) Set(val *SponsoredProductsDraftNegativeTargetAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeTargetAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeTargetAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeTargetAccessError(val *SponsoredProductsDraftNegativeTargetAccessError) *NullableSponsoredProductsDraftNegativeTargetAccessError {
	return &NullableSponsoredProductsDraftNegativeTargetAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeTargetAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeTargetAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
