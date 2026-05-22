package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeKeywordAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeKeywordAccessError{}

// SponsoredProductsDraftNegativeKeywordAccessError struct for SponsoredProductsDraftNegativeKeywordAccessError
type SponsoredProductsDraftNegativeKeywordAccessError struct {
	// The type of the error
	ErrorType  string                                                   `json:"errorType"`
	ErrorValue SponsoredProductsDraftNegativeKeywordAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftNegativeKeywordAccessError SponsoredProductsDraftNegativeKeywordAccessError

// NewSponsoredProductsDraftNegativeKeywordAccessError instantiates a new SponsoredProductsDraftNegativeKeywordAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeKeywordAccessError(errorType string, errorValue SponsoredProductsDraftNegativeKeywordAccessErrorSelector) *SponsoredProductsDraftNegativeKeywordAccessError {
	this := SponsoredProductsDraftNegativeKeywordAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftNegativeKeywordAccessErrorWithDefaults instantiates a new SponsoredProductsDraftNegativeKeywordAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeKeywordAccessErrorWithDefaults() *SponsoredProductsDraftNegativeKeywordAccessError {
	this := SponsoredProductsDraftNegativeKeywordAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftNegativeKeywordAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftNegativeKeywordAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftNegativeKeywordAccessError) GetErrorValue() SponsoredProductsDraftNegativeKeywordAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftNegativeKeywordAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordAccessError) GetErrorValueOk() (*SponsoredProductsDraftNegativeKeywordAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftNegativeKeywordAccessError) SetErrorValue(v SponsoredProductsDraftNegativeKeywordAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftNegativeKeywordAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeKeywordAccessError struct {
	value *SponsoredProductsDraftNegativeKeywordAccessError
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeKeywordAccessError) Get() *SponsoredProductsDraftNegativeKeywordAccessError {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeKeywordAccessError) Set(val *SponsoredProductsDraftNegativeKeywordAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeKeywordAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeKeywordAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeKeywordAccessError(val *SponsoredProductsDraftNegativeKeywordAccessError) *NullableSponsoredProductsDraftNegativeKeywordAccessError {
	return &NullableSponsoredProductsDraftNegativeKeywordAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeKeywordAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeKeywordAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
