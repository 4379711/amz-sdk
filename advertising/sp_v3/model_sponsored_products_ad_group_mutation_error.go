package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupMutationError{}

// SponsoredProductsAdGroupMutationError struct for SponsoredProductsAdGroupMutationError
type SponsoredProductsAdGroupMutationError struct {
	// The type of the error
	ErrorType  string                                        `json:"errorType"`
	ErrorValue SponsoredProductsAdGroupMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsAdGroupMutationError SponsoredProductsAdGroupMutationError

// NewSponsoredProductsAdGroupMutationError instantiates a new SponsoredProductsAdGroupMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupMutationError(errorType string, errorValue SponsoredProductsAdGroupMutationErrorSelector) *SponsoredProductsAdGroupMutationError {
	this := SponsoredProductsAdGroupMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsAdGroupMutationErrorWithDefaults instantiates a new SponsoredProductsAdGroupMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupMutationErrorWithDefaults() *SponsoredProductsAdGroupMutationError {
	this := SponsoredProductsAdGroupMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsAdGroupMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsAdGroupMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsAdGroupMutationError) GetErrorValue() SponsoredProductsAdGroupMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsAdGroupMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupMutationError) GetErrorValueOk() (*SponsoredProductsAdGroupMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsAdGroupMutationError) SetErrorValue(v SponsoredProductsAdGroupMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsAdGroupMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsAdGroupMutationError struct {
	value *SponsoredProductsAdGroupMutationError
	isSet bool
}

func (v NullableSponsoredProductsAdGroupMutationError) Get() *SponsoredProductsAdGroupMutationError {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupMutationError) Set(val *SponsoredProductsAdGroupMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupMutationError(val *SponsoredProductsAdGroupMutationError) *NullableSponsoredProductsAdGroupMutationError {
	return &NullableSponsoredProductsAdGroupMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
