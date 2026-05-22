package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftKeywordAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeywordAccessError{}

// SponsoredProductsDraftKeywordAccessError struct for SponsoredProductsDraftKeywordAccessError
type SponsoredProductsDraftKeywordAccessError struct {
	// The type of the error
	ErrorType  string                                           `json:"errorType"`
	ErrorValue SponsoredProductsDraftKeywordAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftKeywordAccessError SponsoredProductsDraftKeywordAccessError

// NewSponsoredProductsDraftKeywordAccessError instantiates a new SponsoredProductsDraftKeywordAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeywordAccessError(errorType string, errorValue SponsoredProductsDraftKeywordAccessErrorSelector) *SponsoredProductsDraftKeywordAccessError {
	this := SponsoredProductsDraftKeywordAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftKeywordAccessErrorWithDefaults instantiates a new SponsoredProductsDraftKeywordAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordAccessErrorWithDefaults() *SponsoredProductsDraftKeywordAccessError {
	this := SponsoredProductsDraftKeywordAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftKeywordAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftKeywordAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftKeywordAccessError) GetErrorValue() SponsoredProductsDraftKeywordAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftKeywordAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordAccessError) GetErrorValueOk() (*SponsoredProductsDraftKeywordAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftKeywordAccessError) SetErrorValue(v SponsoredProductsDraftKeywordAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftKeywordAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftKeywordAccessError struct {
	value *SponsoredProductsDraftKeywordAccessError
	isSet bool
}

func (v NullableSponsoredProductsDraftKeywordAccessError) Get() *SponsoredProductsDraftKeywordAccessError {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeywordAccessError) Set(val *SponsoredProductsDraftKeywordAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeywordAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeywordAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeywordAccessError(val *SponsoredProductsDraftKeywordAccessError) *NullableSponsoredProductsDraftKeywordAccessError {
	return &NullableSponsoredProductsDraftKeywordAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeywordAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeywordAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
