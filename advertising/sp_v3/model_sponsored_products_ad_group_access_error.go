package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupAccessError{}

// SponsoredProductsAdGroupAccessError struct for SponsoredProductsAdGroupAccessError
type SponsoredProductsAdGroupAccessError struct {
	// The type of the error
	ErrorType  string                                      `json:"errorType"`
	ErrorValue SponsoredProductsAdGroupAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsAdGroupAccessError SponsoredProductsAdGroupAccessError

// NewSponsoredProductsAdGroupAccessError instantiates a new SponsoredProductsAdGroupAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupAccessError(errorType string, errorValue SponsoredProductsAdGroupAccessErrorSelector) *SponsoredProductsAdGroupAccessError {
	this := SponsoredProductsAdGroupAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsAdGroupAccessErrorWithDefaults instantiates a new SponsoredProductsAdGroupAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupAccessErrorWithDefaults() *SponsoredProductsAdGroupAccessError {
	this := SponsoredProductsAdGroupAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsAdGroupAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsAdGroupAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsAdGroupAccessError) GetErrorValue() SponsoredProductsAdGroupAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsAdGroupAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupAccessError) GetErrorValueOk() (*SponsoredProductsAdGroupAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsAdGroupAccessError) SetErrorValue(v SponsoredProductsAdGroupAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsAdGroupAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsAdGroupAccessError struct {
	value *SponsoredProductsAdGroupAccessError
	isSet bool
}

func (v NullableSponsoredProductsAdGroupAccessError) Get() *SponsoredProductsAdGroupAccessError {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupAccessError) Set(val *SponsoredProductsAdGroupAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupAccessError(val *SponsoredProductsAdGroupAccessError) *NullableSponsoredProductsAdGroupAccessError {
	return &NullableSponsoredProductsAdGroupAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
