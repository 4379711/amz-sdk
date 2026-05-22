package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeywordAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeywordAccessError{}

// SponsoredProductsNegativeKeywordAccessError struct for SponsoredProductsNegativeKeywordAccessError
type SponsoredProductsNegativeKeywordAccessError struct {
	// The type of the error
	ErrorType  string                                              `json:"errorType"`
	ErrorValue SponsoredProductsNegativeKeywordAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsNegativeKeywordAccessError SponsoredProductsNegativeKeywordAccessError

// NewSponsoredProductsNegativeKeywordAccessError instantiates a new SponsoredProductsNegativeKeywordAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeywordAccessError(errorType string, errorValue SponsoredProductsNegativeKeywordAccessErrorSelector) *SponsoredProductsNegativeKeywordAccessError {
	this := SponsoredProductsNegativeKeywordAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsNegativeKeywordAccessErrorWithDefaults instantiates a new SponsoredProductsNegativeKeywordAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordAccessErrorWithDefaults() *SponsoredProductsNegativeKeywordAccessError {
	this := SponsoredProductsNegativeKeywordAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsNegativeKeywordAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsNegativeKeywordAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsNegativeKeywordAccessError) GetErrorValue() SponsoredProductsNegativeKeywordAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsNegativeKeywordAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordAccessError) GetErrorValueOk() (*SponsoredProductsNegativeKeywordAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsNegativeKeywordAccessError) SetErrorValue(v SponsoredProductsNegativeKeywordAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsNegativeKeywordAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeKeywordAccessError struct {
	value *SponsoredProductsNegativeKeywordAccessError
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeywordAccessError) Get() *SponsoredProductsNegativeKeywordAccessError {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeywordAccessError) Set(val *SponsoredProductsNegativeKeywordAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeywordAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeywordAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeywordAccessError(val *SponsoredProductsNegativeKeywordAccessError) *NullableSponsoredProductsNegativeKeywordAccessError {
	return &NullableSponsoredProductsNegativeKeywordAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeywordAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeywordAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
