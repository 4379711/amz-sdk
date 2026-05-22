package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeywordMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeywordMutationError{}

// SponsoredProductsNegativeKeywordMutationError struct for SponsoredProductsNegativeKeywordMutationError
type SponsoredProductsNegativeKeywordMutationError struct {
	// The type of the error
	ErrorType  string                                                `json:"errorType"`
	ErrorValue SponsoredProductsNegativeKeywordMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsNegativeKeywordMutationError SponsoredProductsNegativeKeywordMutationError

// NewSponsoredProductsNegativeKeywordMutationError instantiates a new SponsoredProductsNegativeKeywordMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeywordMutationError(errorType string, errorValue SponsoredProductsNegativeKeywordMutationErrorSelector) *SponsoredProductsNegativeKeywordMutationError {
	this := SponsoredProductsNegativeKeywordMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsNegativeKeywordMutationErrorWithDefaults instantiates a new SponsoredProductsNegativeKeywordMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordMutationErrorWithDefaults() *SponsoredProductsNegativeKeywordMutationError {
	this := SponsoredProductsNegativeKeywordMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsNegativeKeywordMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsNegativeKeywordMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsNegativeKeywordMutationError) GetErrorValue() SponsoredProductsNegativeKeywordMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsNegativeKeywordMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordMutationError) GetErrorValueOk() (*SponsoredProductsNegativeKeywordMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsNegativeKeywordMutationError) SetErrorValue(v SponsoredProductsNegativeKeywordMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsNegativeKeywordMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeKeywordMutationError struct {
	value *SponsoredProductsNegativeKeywordMutationError
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeywordMutationError) Get() *SponsoredProductsNegativeKeywordMutationError {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeywordMutationError) Set(val *SponsoredProductsNegativeKeywordMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeywordMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeywordMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeywordMutationError(val *SponsoredProductsNegativeKeywordMutationError) *NullableSponsoredProductsNegativeKeywordMutationError {
	return &NullableSponsoredProductsNegativeKeywordMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeywordMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeywordMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
