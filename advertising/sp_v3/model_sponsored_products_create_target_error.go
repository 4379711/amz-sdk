package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetError{}

// SponsoredProductsCreateTargetError Response object of failed target promotion group target.
type SponsoredProductsCreateTargetError struct {
	// The type of the error.
	ErrorType  *string                                     `json:"errorType,omitempty"`
	ErrorValue *SponsoredProductsCreateTargetErrorSelector `json:"errorValue,omitempty"`
}

// NewSponsoredProductsCreateTargetError instantiates a new SponsoredProductsCreateTargetError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetError() *SponsoredProductsCreateTargetError {
	this := SponsoredProductsCreateTargetError{}
	return &this
}

// NewSponsoredProductsCreateTargetErrorWithDefaults instantiates a new SponsoredProductsCreateTargetError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetErrorWithDefaults() *SponsoredProductsCreateTargetError {
	this := SponsoredProductsCreateTargetError{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetError) GetErrorType() string {
	if o == nil || IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetError) GetErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetError) HasErrorType() bool {
	if o != nil && !IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *SponsoredProductsCreateTargetError) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetError) GetErrorValue() SponsoredProductsCreateTargetErrorSelector {
	if o == nil || IsNil(o.ErrorValue) {
		var ret SponsoredProductsCreateTargetErrorSelector
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetError) GetErrorValueOk() (*SponsoredProductsCreateTargetErrorSelector, bool) {
	if o == nil || IsNil(o.ErrorValue) {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetError) HasErrorValue() bool {
	if o != nil && !IsNil(o.ErrorValue) {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given SponsoredProductsCreateTargetErrorSelector and assigns it to the ErrorValue field.
func (o *SponsoredProductsCreateTargetError) SetErrorValue(v SponsoredProductsCreateTargetErrorSelector) {
	o.ErrorValue = &v
}

func (o SponsoredProductsCreateTargetError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	if !IsNil(o.ErrorValue) {
		toSerialize["errorValue"] = o.ErrorValue
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetError struct {
	value *SponsoredProductsCreateTargetError
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetError) Get() *SponsoredProductsCreateTargetError {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetError) Set(val *SponsoredProductsCreateTargetError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetError(val *SponsoredProductsCreateTargetError) *NullableSponsoredProductsCreateTargetError {
	return &NullableSponsoredProductsCreateTargetError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
