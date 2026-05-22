package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAdMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAdMutationError{}

// SponsoredProductsProductAdMutationError struct for SponsoredProductsProductAdMutationError
type SponsoredProductsProductAdMutationError struct {
	// The type of the error
	ErrorType  string                                          `json:"errorType"`
	ErrorValue SponsoredProductsProductAdMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsProductAdMutationError SponsoredProductsProductAdMutationError

// NewSponsoredProductsProductAdMutationError instantiates a new SponsoredProductsProductAdMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAdMutationError(errorType string, errorValue SponsoredProductsProductAdMutationErrorSelector) *SponsoredProductsProductAdMutationError {
	this := SponsoredProductsProductAdMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsProductAdMutationErrorWithDefaults instantiates a new SponsoredProductsProductAdMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdMutationErrorWithDefaults() *SponsoredProductsProductAdMutationError {
	this := SponsoredProductsProductAdMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsProductAdMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsProductAdMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsProductAdMutationError) GetErrorValue() SponsoredProductsProductAdMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsProductAdMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdMutationError) GetErrorValueOk() (*SponsoredProductsProductAdMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsProductAdMutationError) SetErrorValue(v SponsoredProductsProductAdMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsProductAdMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsProductAdMutationError struct {
	value *SponsoredProductsProductAdMutationError
	isSet bool
}

func (v NullableSponsoredProductsProductAdMutationError) Get() *SponsoredProductsProductAdMutationError {
	return v.value
}

func (v *NullableSponsoredProductsProductAdMutationError) Set(val *SponsoredProductsProductAdMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAdMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAdMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAdMutationError(val *SponsoredProductsProductAdMutationError) *NullableSponsoredProductsProductAdMutationError {
	return &NullableSponsoredProductsProductAdMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAdMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAdMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
