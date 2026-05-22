package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdMutationError{}

// AdMutationError struct for AdMutationError
type AdMutationError struct {
	// The type of the error.
	ErrorType  string                  `json:"errorType"`
	ErrorValue AdMutationErrorSelector `json:"errorValue"`
}

type _AdMutationError AdMutationError

// NewAdMutationError instantiates a new AdMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdMutationError(errorType string, errorValue AdMutationErrorSelector) *AdMutationError {
	this := AdMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewAdMutationErrorWithDefaults instantiates a new AdMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdMutationErrorWithDefaults() *AdMutationError {
	this := AdMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *AdMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *AdMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *AdMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *AdMutationError) GetErrorValue() AdMutationErrorSelector {
	if o == nil {
		var ret AdMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *AdMutationError) GetErrorValueOk() (*AdMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *AdMutationError) SetErrorValue(v AdMutationErrorSelector) {
	o.ErrorValue = v
}

func (o AdMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableAdMutationError struct {
	value *AdMutationError
	isSet bool
}

func (v NullableAdMutationError) Get() *AdMutationError {
	return v.value
}

func (v *NullableAdMutationError) Set(val *AdMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableAdMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableAdMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdMutationError(val *AdMutationError) *NullableAdMutationError {
	return &NullableAdMutationError{value: val, isSet: true}
}

func (v NullableAdMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
