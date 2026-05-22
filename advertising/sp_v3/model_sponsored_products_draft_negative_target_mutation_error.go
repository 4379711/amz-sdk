package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeTargetMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeTargetMutationError{}

// SponsoredProductsDraftNegativeTargetMutationError struct for SponsoredProductsDraftNegativeTargetMutationError
type SponsoredProductsDraftNegativeTargetMutationError struct {
	// The type of the error
	ErrorType  string                                                    `json:"errorType"`
	ErrorValue SponsoredProductsDraftNegativeTargetMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftNegativeTargetMutationError SponsoredProductsDraftNegativeTargetMutationError

// NewSponsoredProductsDraftNegativeTargetMutationError instantiates a new SponsoredProductsDraftNegativeTargetMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeTargetMutationError(errorType string, errorValue SponsoredProductsDraftNegativeTargetMutationErrorSelector) *SponsoredProductsDraftNegativeTargetMutationError {
	this := SponsoredProductsDraftNegativeTargetMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftNegativeTargetMutationErrorWithDefaults instantiates a new SponsoredProductsDraftNegativeTargetMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeTargetMutationErrorWithDefaults() *SponsoredProductsDraftNegativeTargetMutationError {
	this := SponsoredProductsDraftNegativeTargetMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftNegativeTargetMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftNegativeTargetMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftNegativeTargetMutationError) GetErrorValue() SponsoredProductsDraftNegativeTargetMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftNegativeTargetMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetMutationError) GetErrorValueOk() (*SponsoredProductsDraftNegativeTargetMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftNegativeTargetMutationError) SetErrorValue(v SponsoredProductsDraftNegativeTargetMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftNegativeTargetMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeTargetMutationError struct {
	value *SponsoredProductsDraftNegativeTargetMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeTargetMutationError) Get() *SponsoredProductsDraftNegativeTargetMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeTargetMutationError) Set(val *SponsoredProductsDraftNegativeTargetMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeTargetMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeTargetMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeTargetMutationError(val *SponsoredProductsDraftNegativeTargetMutationError) *NullableSponsoredProductsDraftNegativeTargetMutationError {
	return &NullableSponsoredProductsDraftNegativeTargetMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeTargetMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeTargetMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
