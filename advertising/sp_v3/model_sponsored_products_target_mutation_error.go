package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetMutationError{}

// SponsoredProductsTargetMutationError struct for SponsoredProductsTargetMutationError
type SponsoredProductsTargetMutationError struct {
	// The type of the error
	ErrorType  string                                       `json:"errorType"`
	ErrorValue SponsoredProductsTargetMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsTargetMutationError SponsoredProductsTargetMutationError

// NewSponsoredProductsTargetMutationError instantiates a new SponsoredProductsTargetMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetMutationError(errorType string, errorValue SponsoredProductsTargetMutationErrorSelector) *SponsoredProductsTargetMutationError {
	this := SponsoredProductsTargetMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsTargetMutationErrorWithDefaults instantiates a new SponsoredProductsTargetMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetMutationErrorWithDefaults() *SponsoredProductsTargetMutationError {
	this := SponsoredProductsTargetMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsTargetMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsTargetMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsTargetMutationError) GetErrorValue() SponsoredProductsTargetMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsTargetMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetMutationError) GetErrorValueOk() (*SponsoredProductsTargetMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsTargetMutationError) SetErrorValue(v SponsoredProductsTargetMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsTargetMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsTargetMutationError struct {
	value *SponsoredProductsTargetMutationError
	isSet bool
}

func (v NullableSponsoredProductsTargetMutationError) Get() *SponsoredProductsTargetMutationError {
	return v.value
}

func (v *NullableSponsoredProductsTargetMutationError) Set(val *SponsoredProductsTargetMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetMutationError(val *SponsoredProductsTargetMutationError) *NullableSponsoredProductsTargetMutationError {
	return &NullableSponsoredProductsTargetMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
