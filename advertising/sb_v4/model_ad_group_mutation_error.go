package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupMutationError{}

// AdGroupMutationError struct for AdGroupMutationError
type AdGroupMutationError struct {
	// The type of the error.
	ErrorType  string                       `json:"errorType"`
	ErrorValue AdGroupMutationErrorSelector `json:"errorValue"`
}

type _AdGroupMutationError AdGroupMutationError

// NewAdGroupMutationError instantiates a new AdGroupMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupMutationError(errorType string, errorValue AdGroupMutationErrorSelector) *AdGroupMutationError {
	this := AdGroupMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewAdGroupMutationErrorWithDefaults instantiates a new AdGroupMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupMutationErrorWithDefaults() *AdGroupMutationError {
	this := AdGroupMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *AdGroupMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *AdGroupMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *AdGroupMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *AdGroupMutationError) GetErrorValue() AdGroupMutationErrorSelector {
	if o == nil {
		var ret AdGroupMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *AdGroupMutationError) GetErrorValueOk() (*AdGroupMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *AdGroupMutationError) SetErrorValue(v AdGroupMutationErrorSelector) {
	o.ErrorValue = v
}

func (o AdGroupMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableAdGroupMutationError struct {
	value *AdGroupMutationError
	isSet bool
}

func (v NullableAdGroupMutationError) Get() *AdGroupMutationError {
	return v.value
}

func (v *NullableAdGroupMutationError) Set(val *AdGroupMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupMutationError(val *AdGroupMutationError) *NullableAdGroupMutationError {
	return &NullableAdGroupMutationError{value: val, isSet: true}
}

func (v NullableAdGroupMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
