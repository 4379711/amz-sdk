package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIOptimizationRuleError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIOptimizationRuleError{}

// OptimizationRulesAPIOptimizationRuleError Object defining one reason for a failure of a optimization rule operation.
type OptimizationRulesAPIOptimizationRuleError struct {
	// The http status error code for machine use.
	Code string `json:"code"`
	// A human-readable description of the error.
	Message string `json:"message"`
}

type _OptimizationRulesAPIOptimizationRuleError OptimizationRulesAPIOptimizationRuleError

// NewOptimizationRulesAPIOptimizationRuleError instantiates a new OptimizationRulesAPIOptimizationRuleError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIOptimizationRuleError(code string, message string) *OptimizationRulesAPIOptimizationRuleError {
	this := OptimizationRulesAPIOptimizationRuleError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewOptimizationRulesAPIOptimizationRuleErrorWithDefaults instantiates a new OptimizationRulesAPIOptimizationRuleError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIOptimizationRuleErrorWithDefaults() *OptimizationRulesAPIOptimizationRuleError {
	this := OptimizationRulesAPIOptimizationRuleError{}
	return &this
}

// GetCode returns the Code field value
func (o *OptimizationRulesAPIOptimizationRuleError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRuleError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OptimizationRulesAPIOptimizationRuleError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *OptimizationRulesAPIOptimizationRuleError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRuleError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OptimizationRulesAPIOptimizationRuleError) SetMessage(v string) {
	o.Message = v
}

func (o OptimizationRulesAPIOptimizationRuleError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableOptimizationRulesAPIOptimizationRuleError struct {
	value *OptimizationRulesAPIOptimizationRuleError
	isSet bool
}

func (v NullableOptimizationRulesAPIOptimizationRuleError) Get() *OptimizationRulesAPIOptimizationRuleError {
	return v.value
}

func (v *NullableOptimizationRulesAPIOptimizationRuleError) Set(val *OptimizationRulesAPIOptimizationRuleError) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIOptimizationRuleError) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIOptimizationRuleError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIOptimizationRuleError(val *OptimizationRulesAPIOptimizationRuleError) *NullableOptimizationRulesAPIOptimizationRuleError {
	return &NullableOptimizationRulesAPIOptimizationRuleError{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIOptimizationRuleError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIOptimizationRuleError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
