package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAdAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAdAccessError{}

// SponsoredProductsProductAdAccessError struct for SponsoredProductsProductAdAccessError
type SponsoredProductsProductAdAccessError struct {
	// The type of the error
	ErrorType  string                                        `json:"errorType"`
	ErrorValue SponsoredProductsProductAdAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsProductAdAccessError SponsoredProductsProductAdAccessError

// NewSponsoredProductsProductAdAccessError instantiates a new SponsoredProductsProductAdAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAdAccessError(errorType string, errorValue SponsoredProductsProductAdAccessErrorSelector) *SponsoredProductsProductAdAccessError {
	this := SponsoredProductsProductAdAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsProductAdAccessErrorWithDefaults instantiates a new SponsoredProductsProductAdAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdAccessErrorWithDefaults() *SponsoredProductsProductAdAccessError {
	this := SponsoredProductsProductAdAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsProductAdAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsProductAdAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsProductAdAccessError) GetErrorValue() SponsoredProductsProductAdAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsProductAdAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdAccessError) GetErrorValueOk() (*SponsoredProductsProductAdAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsProductAdAccessError) SetErrorValue(v SponsoredProductsProductAdAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsProductAdAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsProductAdAccessError struct {
	value *SponsoredProductsProductAdAccessError
	isSet bool
}

func (v NullableSponsoredProductsProductAdAccessError) Get() *SponsoredProductsProductAdAccessError {
	return v.value
}

func (v *NullableSponsoredProductsProductAdAccessError) Set(val *SponsoredProductsProductAdAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAdAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAdAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAdAccessError(val *SponsoredProductsProductAdAccessError) *NullableSponsoredProductsProductAdAccessError {
	return &NullableSponsoredProductsProductAdAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAdAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAdAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
