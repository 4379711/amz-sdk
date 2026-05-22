package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftAdGroupMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroupMutationError{}

// SponsoredProductsDraftAdGroupMutationError struct for SponsoredProductsDraftAdGroupMutationError
type SponsoredProductsDraftAdGroupMutationError struct {
	// The type of the error
	ErrorType  string                                             `json:"errorType"`
	ErrorValue SponsoredProductsDraftAdGroupMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftAdGroupMutationError SponsoredProductsDraftAdGroupMutationError

// NewSponsoredProductsDraftAdGroupMutationError instantiates a new SponsoredProductsDraftAdGroupMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroupMutationError(errorType string, errorValue SponsoredProductsDraftAdGroupMutationErrorSelector) *SponsoredProductsDraftAdGroupMutationError {
	this := SponsoredProductsDraftAdGroupMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftAdGroupMutationErrorWithDefaults instantiates a new SponsoredProductsDraftAdGroupMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupMutationErrorWithDefaults() *SponsoredProductsDraftAdGroupMutationError {
	this := SponsoredProductsDraftAdGroupMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftAdGroupMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftAdGroupMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftAdGroupMutationError) GetErrorValue() SponsoredProductsDraftAdGroupMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftAdGroupMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupMutationError) GetErrorValueOk() (*SponsoredProductsDraftAdGroupMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftAdGroupMutationError) SetErrorValue(v SponsoredProductsDraftAdGroupMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftAdGroupMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftAdGroupMutationError struct {
	value *SponsoredProductsDraftAdGroupMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroupMutationError) Get() *SponsoredProductsDraftAdGroupMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroupMutationError) Set(val *SponsoredProductsDraftAdGroupMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroupMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroupMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroupMutationError(val *SponsoredProductsDraftAdGroupMutationError) *NullableSponsoredProductsDraftAdGroupMutationError {
	return &NullableSponsoredProductsDraftAdGroupMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroupMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroupMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
