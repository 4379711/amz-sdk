package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidArgumentError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidArgumentError{}

// InvalidArgumentError struct for InvalidArgumentError
type InvalidArgumentError struct {
	// The type of the error
	ErrorType  string                       `json:"errorType"`
	ErrorValue InvalidArgumentErrorSelector `json:"errorValue"`
}

type _InvalidArgumentError InvalidArgumentError

// NewInvalidArgumentError instantiates a new InvalidArgumentError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidArgumentError(errorType string, errorValue InvalidArgumentErrorSelector) *InvalidArgumentError {
	this := InvalidArgumentError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewInvalidArgumentErrorWithDefaults instantiates a new InvalidArgumentError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidArgumentErrorWithDefaults() *InvalidArgumentError {
	this := InvalidArgumentError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *InvalidArgumentError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *InvalidArgumentError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *InvalidArgumentError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *InvalidArgumentError) GetErrorValue() InvalidArgumentErrorSelector {
	if o == nil {
		var ret InvalidArgumentErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *InvalidArgumentError) GetErrorValueOk() (*InvalidArgumentErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *InvalidArgumentError) SetErrorValue(v InvalidArgumentErrorSelector) {
	o.ErrorValue = v
}

func (o InvalidArgumentError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableInvalidArgumentError struct {
	value *InvalidArgumentError
	isSet bool
}

func (v NullableInvalidArgumentError) Get() *InvalidArgumentError {
	return v.value
}

func (v *NullableInvalidArgumentError) Set(val *InvalidArgumentError) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidArgumentError) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidArgumentError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidArgumentError(val *InvalidArgumentError) *NullableInvalidArgumentError {
	return &NullableInvalidArgumentError{value: val, isSet: true}
}

func (v NullableInvalidArgumentError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidArgumentError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
