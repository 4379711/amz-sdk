package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIOptimizationRulesError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIOptimizationRulesError{}

// OptimizationRulesAPIOptimizationRulesError Error response object.
type OptimizationRulesAPIOptimizationRulesError struct {
	// An enumerated error code for machine use.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Message *string `json:"message,omitempty"`
}

// NewOptimizationRulesAPIOptimizationRulesError instantiates a new OptimizationRulesAPIOptimizationRulesError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIOptimizationRulesError() *OptimizationRulesAPIOptimizationRulesError {
	this := OptimizationRulesAPIOptimizationRulesError{}
	return &this
}

// NewOptimizationRulesAPIOptimizationRulesErrorWithDefaults instantiates a new OptimizationRulesAPIOptimizationRulesError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIOptimizationRulesErrorWithDefaults() *OptimizationRulesAPIOptimizationRulesError {
	this := OptimizationRulesAPIOptimizationRulesError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRulesError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRulesError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRulesError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPIOptimizationRulesError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRulesError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRulesError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRulesError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OptimizationRulesAPIOptimizationRulesError) SetMessage(v string) {
	o.Message = &v
}

func (o OptimizationRulesAPIOptimizationRulesError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIOptimizationRulesError struct {
	value *OptimizationRulesAPIOptimizationRulesError
	isSet bool
}

func (v NullableOptimizationRulesAPIOptimizationRulesError) Get() *OptimizationRulesAPIOptimizationRulesError {
	return v.value
}

func (v *NullableOptimizationRulesAPIOptimizationRulesError) Set(val *OptimizationRulesAPIOptimizationRulesError) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIOptimizationRulesError) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIOptimizationRulesError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIOptimizationRulesError(val *OptimizationRulesAPIOptimizationRulesError) *NullableOptimizationRulesAPIOptimizationRulesError {
	return &NullableOptimizationRulesAPIOptimizationRulesError{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIOptimizationRulesError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIOptimizationRulesError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
