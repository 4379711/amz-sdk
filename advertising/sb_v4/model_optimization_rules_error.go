package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesError{}

// OptimizationRulesError struct for OptimizationRulesError
type OptimizationRulesError struct {
	// The type of the error.
	Code string `json:"code"`
	// Human readable error message.
	Message string `json:"message"`
}

type _OptimizationRulesError OptimizationRulesError

// NewOptimizationRulesError instantiates a new OptimizationRulesError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesError(code string, message string) *OptimizationRulesError {
	this := OptimizationRulesError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewOptimizationRulesErrorWithDefaults instantiates a new OptimizationRulesError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesErrorWithDefaults() *OptimizationRulesError {
	this := OptimizationRulesError{}
	return &this
}

// GetCode returns the Code field value
func (o *OptimizationRulesError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OptimizationRulesError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *OptimizationRulesError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OptimizationRulesError) SetMessage(v string) {
	o.Message = v
}

func (o OptimizationRulesError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableOptimizationRulesError struct {
	value *OptimizationRulesError
	isSet bool
}

func (v NullableOptimizationRulesError) Get() *OptimizationRulesError {
	return v.value
}

func (v *NullableOptimizationRulesError) Set(val *OptimizationRulesError) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesError) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesError(val *OptimizationRulesError) *NullableOptimizationRulesError {
	return &NullableOptimizationRulesError{value: val, isSet: true}
}

func (v NullableOptimizationRulesError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
