package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftTargetMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftTargetMutationError{}

// SponsoredProductsDraftTargetMutationError struct for SponsoredProductsDraftTargetMutationError
type SponsoredProductsDraftTargetMutationError struct {
	// The type of the error
	ErrorType  string                                            `json:"errorType"`
	ErrorValue SponsoredProductsDraftTargetMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftTargetMutationError SponsoredProductsDraftTargetMutationError

// NewSponsoredProductsDraftTargetMutationError instantiates a new SponsoredProductsDraftTargetMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftTargetMutationError(errorType string, errorValue SponsoredProductsDraftTargetMutationErrorSelector) *SponsoredProductsDraftTargetMutationError {
	this := SponsoredProductsDraftTargetMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftTargetMutationErrorWithDefaults instantiates a new SponsoredProductsDraftTargetMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftTargetMutationErrorWithDefaults() *SponsoredProductsDraftTargetMutationError {
	this := SponsoredProductsDraftTargetMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftTargetMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftTargetMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftTargetMutationError) GetErrorValue() SponsoredProductsDraftTargetMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftTargetMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetMutationError) GetErrorValueOk() (*SponsoredProductsDraftTargetMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftTargetMutationError) SetErrorValue(v SponsoredProductsDraftTargetMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftTargetMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftTargetMutationError struct {
	value *SponsoredProductsDraftTargetMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftTargetMutationError) Get() *SponsoredProductsDraftTargetMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftTargetMutationError) Set(val *SponsoredProductsDraftTargetMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftTargetMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftTargetMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftTargetMutationError(val *SponsoredProductsDraftTargetMutationError) *NullableSponsoredProductsDraftTargetMutationError {
	return &NullableSponsoredProductsDraftTargetMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftTargetMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftTargetMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
