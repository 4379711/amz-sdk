package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeTargetAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeTargetAccessError{}

// SponsoredProductsNegativeTargetAccessError struct for SponsoredProductsNegativeTargetAccessError
type SponsoredProductsNegativeTargetAccessError struct {
	// The type of the error
	ErrorType  string                                             `json:"errorType"`
	ErrorValue SponsoredProductsNegativeTargetAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsNegativeTargetAccessError SponsoredProductsNegativeTargetAccessError

// NewSponsoredProductsNegativeTargetAccessError instantiates a new SponsoredProductsNegativeTargetAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeTargetAccessError(errorType string, errorValue SponsoredProductsNegativeTargetAccessErrorSelector) *SponsoredProductsNegativeTargetAccessError {
	this := SponsoredProductsNegativeTargetAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsNegativeTargetAccessErrorWithDefaults instantiates a new SponsoredProductsNegativeTargetAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeTargetAccessErrorWithDefaults() *SponsoredProductsNegativeTargetAccessError {
	this := SponsoredProductsNegativeTargetAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsNegativeTargetAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsNegativeTargetAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsNegativeTargetAccessError) GetErrorValue() SponsoredProductsNegativeTargetAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsNegativeTargetAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetAccessError) GetErrorValueOk() (*SponsoredProductsNegativeTargetAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsNegativeTargetAccessError) SetErrorValue(v SponsoredProductsNegativeTargetAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsNegativeTargetAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeTargetAccessError struct {
	value *SponsoredProductsNegativeTargetAccessError
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetAccessError) Get() *SponsoredProductsNegativeTargetAccessError {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetAccessError) Set(val *SponsoredProductsNegativeTargetAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetAccessError(val *SponsoredProductsNegativeTargetAccessError) *NullableSponsoredProductsNegativeTargetAccessError {
	return &NullableSponsoredProductsNegativeTargetAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
