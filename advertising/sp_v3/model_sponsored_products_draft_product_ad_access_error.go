package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftProductAdAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdAccessError{}

// SponsoredProductsDraftProductAdAccessError struct for SponsoredProductsDraftProductAdAccessError
type SponsoredProductsDraftProductAdAccessError struct {
	// The type of the error
	ErrorType  string                                             `json:"errorType"`
	ErrorValue SponsoredProductsDraftProductAdAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftProductAdAccessError SponsoredProductsDraftProductAdAccessError

// NewSponsoredProductsDraftProductAdAccessError instantiates a new SponsoredProductsDraftProductAdAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdAccessError(errorType string, errorValue SponsoredProductsDraftProductAdAccessErrorSelector) *SponsoredProductsDraftProductAdAccessError {
	this := SponsoredProductsDraftProductAdAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftProductAdAccessErrorWithDefaults instantiates a new SponsoredProductsDraftProductAdAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdAccessErrorWithDefaults() *SponsoredProductsDraftProductAdAccessError {
	this := SponsoredProductsDraftProductAdAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftProductAdAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftProductAdAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftProductAdAccessError) GetErrorValue() SponsoredProductsDraftProductAdAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftProductAdAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdAccessError) GetErrorValueOk() (*SponsoredProductsDraftProductAdAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftProductAdAccessError) SetErrorValue(v SponsoredProductsDraftProductAdAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftProductAdAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftProductAdAccessError struct {
	value *SponsoredProductsDraftProductAdAccessError
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdAccessError) Get() *SponsoredProductsDraftProductAdAccessError {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdAccessError) Set(val *SponsoredProductsDraftProductAdAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdAccessError(val *SponsoredProductsDraftProductAdAccessError) *NullableSponsoredProductsDraftProductAdAccessError {
	return &NullableSponsoredProductsDraftProductAdAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
