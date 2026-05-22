package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordMutationError{}

// SponsoredProductsKeywordMutationError struct for SponsoredProductsKeywordMutationError
type SponsoredProductsKeywordMutationError struct {
	// The type of the error
	ErrorType  string                                        `json:"errorType"`
	ErrorValue SponsoredProductsKeywordMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsKeywordMutationError SponsoredProductsKeywordMutationError

// NewSponsoredProductsKeywordMutationError instantiates a new SponsoredProductsKeywordMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordMutationError(errorType string, errorValue SponsoredProductsKeywordMutationErrorSelector) *SponsoredProductsKeywordMutationError {
	this := SponsoredProductsKeywordMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsKeywordMutationErrorWithDefaults instantiates a new SponsoredProductsKeywordMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordMutationErrorWithDefaults() *SponsoredProductsKeywordMutationError {
	this := SponsoredProductsKeywordMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsKeywordMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsKeywordMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsKeywordMutationError) GetErrorValue() SponsoredProductsKeywordMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsKeywordMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordMutationError) GetErrorValueOk() (*SponsoredProductsKeywordMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsKeywordMutationError) SetErrorValue(v SponsoredProductsKeywordMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsKeywordMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordMutationError struct {
	value *SponsoredProductsKeywordMutationError
	isSet bool
}

func (v NullableSponsoredProductsKeywordMutationError) Get() *SponsoredProductsKeywordMutationError {
	return v.value
}

func (v *NullableSponsoredProductsKeywordMutationError) Set(val *SponsoredProductsKeywordMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordMutationError(val *SponsoredProductsKeywordMutationError) *NullableSponsoredProductsKeywordMutationError {
	return &NullableSponsoredProductsKeywordMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
