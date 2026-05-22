package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeTargetMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeTargetMutationError{}

// SponsoredProductsNegativeTargetMutationError struct for SponsoredProductsNegativeTargetMutationError
type SponsoredProductsNegativeTargetMutationError struct {
	// The type of the error
	ErrorType  string                                               `json:"errorType"`
	ErrorValue SponsoredProductsNegativeTargetMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsNegativeTargetMutationError SponsoredProductsNegativeTargetMutationError

// NewSponsoredProductsNegativeTargetMutationError instantiates a new SponsoredProductsNegativeTargetMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeTargetMutationError(errorType string, errorValue SponsoredProductsNegativeTargetMutationErrorSelector) *SponsoredProductsNegativeTargetMutationError {
	this := SponsoredProductsNegativeTargetMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsNegativeTargetMutationErrorWithDefaults instantiates a new SponsoredProductsNegativeTargetMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeTargetMutationErrorWithDefaults() *SponsoredProductsNegativeTargetMutationError {
	this := SponsoredProductsNegativeTargetMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsNegativeTargetMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsNegativeTargetMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsNegativeTargetMutationError) GetErrorValue() SponsoredProductsNegativeTargetMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsNegativeTargetMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetMutationError) GetErrorValueOk() (*SponsoredProductsNegativeTargetMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsNegativeTargetMutationError) SetErrorValue(v SponsoredProductsNegativeTargetMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsNegativeTargetMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeTargetMutationError struct {
	value *SponsoredProductsNegativeTargetMutationError
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetMutationError) Get() *SponsoredProductsNegativeTargetMutationError {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetMutationError) Set(val *SponsoredProductsNegativeTargetMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetMutationError(val *SponsoredProductsNegativeTargetMutationError) *NullableSponsoredProductsNegativeTargetMutationError {
	return &NullableSponsoredProductsNegativeTargetMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
