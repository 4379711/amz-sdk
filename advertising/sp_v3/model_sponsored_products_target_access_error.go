package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsTargetAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetAccessError{}

// SponsoredProductsTargetAccessError struct for SponsoredProductsTargetAccessError
type SponsoredProductsTargetAccessError struct {
	// The type of the error
	ErrorType  string                                     `json:"errorType"`
	ErrorValue SponsoredProductsTargetAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsTargetAccessError SponsoredProductsTargetAccessError

// NewSponsoredProductsTargetAccessError instantiates a new SponsoredProductsTargetAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetAccessError(errorType string, errorValue SponsoredProductsTargetAccessErrorSelector) *SponsoredProductsTargetAccessError {
	this := SponsoredProductsTargetAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsTargetAccessErrorWithDefaults instantiates a new SponsoredProductsTargetAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetAccessErrorWithDefaults() *SponsoredProductsTargetAccessError {
	this := SponsoredProductsTargetAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsTargetAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsTargetAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsTargetAccessError) GetErrorValue() SponsoredProductsTargetAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsTargetAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetAccessError) GetErrorValueOk() (*SponsoredProductsTargetAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsTargetAccessError) SetErrorValue(v SponsoredProductsTargetAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsTargetAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsTargetAccessError struct {
	value *SponsoredProductsTargetAccessError
	isSet bool
}

func (v NullableSponsoredProductsTargetAccessError) Get() *SponsoredProductsTargetAccessError {
	return v.value
}

func (v *NullableSponsoredProductsTargetAccessError) Set(val *SponsoredProductsTargetAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetAccessError(val *SponsoredProductsTargetAccessError) *NullableSponsoredProductsTargetAccessError {
	return &NullableSponsoredProductsTargetAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
