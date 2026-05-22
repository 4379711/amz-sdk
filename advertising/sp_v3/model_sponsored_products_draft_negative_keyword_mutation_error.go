package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeKeywordMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeKeywordMutationError{}

// SponsoredProductsDraftNegativeKeywordMutationError struct for SponsoredProductsDraftNegativeKeywordMutationError
type SponsoredProductsDraftNegativeKeywordMutationError struct {
	// The type of the error
	ErrorType  string                                                     `json:"errorType"`
	ErrorValue SponsoredProductsDraftNegativeKeywordMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftNegativeKeywordMutationError SponsoredProductsDraftNegativeKeywordMutationError

// NewSponsoredProductsDraftNegativeKeywordMutationError instantiates a new SponsoredProductsDraftNegativeKeywordMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeKeywordMutationError(errorType string, errorValue SponsoredProductsDraftNegativeKeywordMutationErrorSelector) *SponsoredProductsDraftNegativeKeywordMutationError {
	this := SponsoredProductsDraftNegativeKeywordMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftNegativeKeywordMutationErrorWithDefaults instantiates a new SponsoredProductsDraftNegativeKeywordMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeKeywordMutationErrorWithDefaults() *SponsoredProductsDraftNegativeKeywordMutationError {
	this := SponsoredProductsDraftNegativeKeywordMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftNegativeKeywordMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftNegativeKeywordMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftNegativeKeywordMutationError) GetErrorValue() SponsoredProductsDraftNegativeKeywordMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftNegativeKeywordMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordMutationError) GetErrorValueOk() (*SponsoredProductsDraftNegativeKeywordMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftNegativeKeywordMutationError) SetErrorValue(v SponsoredProductsDraftNegativeKeywordMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftNegativeKeywordMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeKeywordMutationError struct {
	value *SponsoredProductsDraftNegativeKeywordMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeKeywordMutationError) Get() *SponsoredProductsDraftNegativeKeywordMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeKeywordMutationError) Set(val *SponsoredProductsDraftNegativeKeywordMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeKeywordMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeKeywordMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeKeywordMutationError(val *SponsoredProductsDraftNegativeKeywordMutationError) *NullableSponsoredProductsDraftNegativeKeywordMutationError {
	return &NullableSponsoredProductsDraftNegativeKeywordMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeKeywordMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeKeywordMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
