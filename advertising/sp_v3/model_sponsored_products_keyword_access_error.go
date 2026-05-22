package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordAccessError{}

// SponsoredProductsKeywordAccessError struct for SponsoredProductsKeywordAccessError
type SponsoredProductsKeywordAccessError struct {
	// The type of the error
	ErrorType  string                                      `json:"errorType"`
	ErrorValue SponsoredProductsKeywordAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsKeywordAccessError SponsoredProductsKeywordAccessError

// NewSponsoredProductsKeywordAccessError instantiates a new SponsoredProductsKeywordAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordAccessError(errorType string, errorValue SponsoredProductsKeywordAccessErrorSelector) *SponsoredProductsKeywordAccessError {
	this := SponsoredProductsKeywordAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsKeywordAccessErrorWithDefaults instantiates a new SponsoredProductsKeywordAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordAccessErrorWithDefaults() *SponsoredProductsKeywordAccessError {
	this := SponsoredProductsKeywordAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsKeywordAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsKeywordAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsKeywordAccessError) GetErrorValue() SponsoredProductsKeywordAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsKeywordAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordAccessError) GetErrorValueOk() (*SponsoredProductsKeywordAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsKeywordAccessError) SetErrorValue(v SponsoredProductsKeywordAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsKeywordAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordAccessError struct {
	value *SponsoredProductsKeywordAccessError
	isSet bool
}

func (v NullableSponsoredProductsKeywordAccessError) Get() *SponsoredProductsKeywordAccessError {
	return v.value
}

func (v *NullableSponsoredProductsKeywordAccessError) Set(val *SponsoredProductsKeywordAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordAccessError(val *SponsoredProductsKeywordAccessError) *NullableSponsoredProductsKeywordAccessError {
	return &NullableSponsoredProductsKeywordAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
