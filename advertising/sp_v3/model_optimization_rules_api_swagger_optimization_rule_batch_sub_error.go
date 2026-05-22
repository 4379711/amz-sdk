package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIOptimizationRuleBatchSubError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIOptimizationRuleBatchSubError{}

// OptimizationRulesAPIOptimizationRuleBatchSubError Response object for operations involving a single optimization rule.
type OptimizationRulesAPIOptimizationRuleBatchSubError struct {
	// A human-readable description of the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewOptimizationRulesAPIOptimizationRuleBatchSubError instantiates a new OptimizationRulesAPIOptimizationRuleBatchSubError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIOptimizationRuleBatchSubError() *OptimizationRulesAPIOptimizationRuleBatchSubError {
	this := OptimizationRulesAPIOptimizationRuleBatchSubError{}
	return &this
}

// NewOptimizationRulesAPIOptimizationRuleBatchSubErrorWithDefaults instantiates a new OptimizationRulesAPIOptimizationRuleBatchSubError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIOptimizationRuleBatchSubErrorWithDefaults() *OptimizationRulesAPIOptimizationRuleBatchSubError {
	this := OptimizationRulesAPIOptimizationRuleBatchSubError{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRuleBatchSubError) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRuleBatchSubError) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRuleBatchSubError) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *OptimizationRulesAPIOptimizationRuleBatchSubError) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o OptimizationRulesAPIOptimizationRuleBatchSubError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIOptimizationRuleBatchSubError struct {
	value *OptimizationRulesAPIOptimizationRuleBatchSubError
	isSet bool
}

func (v NullableOptimizationRulesAPIOptimizationRuleBatchSubError) Get() *OptimizationRulesAPIOptimizationRuleBatchSubError {
	return v.value
}

func (v *NullableOptimizationRulesAPIOptimizationRuleBatchSubError) Set(val *OptimizationRulesAPIOptimizationRuleBatchSubError) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIOptimizationRuleBatchSubError) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIOptimizationRuleBatchSubError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIOptimizationRuleBatchSubError(val *OptimizationRulesAPIOptimizationRuleBatchSubError) *NullableOptimizationRulesAPIOptimizationRuleBatchSubError {
	return &NullableOptimizationRulesAPIOptimizationRuleBatchSubError{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIOptimizationRuleBatchSubError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIOptimizationRuleBatchSubError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
