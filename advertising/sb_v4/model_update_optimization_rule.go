package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateOptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOptimizationRule{}

// UpdateOptimizationRule struct for UpdateOptimizationRule
type UpdateOptimizationRule struct {
	// The identifier of the optimization rule.
	OptimizationRuleId *string         `json:"optimizationRuleId,omitempty"`
	Conditions         []RuleCondition `json:"conditions,omitempty"`
}

// NewUpdateOptimizationRule instantiates a new UpdateOptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOptimizationRule() *UpdateOptimizationRule {
	this := UpdateOptimizationRule{}
	return &this
}

// NewUpdateOptimizationRuleWithDefaults instantiates a new UpdateOptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOptimizationRuleWithDefaults() *UpdateOptimizationRule {
	this := UpdateOptimizationRule{}
	return &this
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *UpdateOptimizationRule) GetOptimizationRuleId() string {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret string
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRule) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *UpdateOptimizationRule) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given string and assigns it to the OptimizationRuleId field.
func (o *UpdateOptimizationRule) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *UpdateOptimizationRule) GetConditions() []RuleCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []RuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRule) GetConditionsOk() ([]RuleCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *UpdateOptimizationRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []RuleCondition and assigns it to the Conditions field.
func (o *UpdateOptimizationRule) SetConditions(v []RuleCondition) {
	o.Conditions = v
}

func (o UpdateOptimizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableUpdateOptimizationRule struct {
	value *UpdateOptimizationRule
	isSet bool
}

func (v NullableUpdateOptimizationRule) Get() *UpdateOptimizationRule {
	return v.value
}

func (v *NullableUpdateOptimizationRule) Set(val *UpdateOptimizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOptimizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOptimizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOptimizationRule(val *UpdateOptimizationRule) *NullableUpdateOptimizationRule {
	return &NullableUpdateOptimizationRule{value: val, isSet: true}
}

func (v NullableUpdateOptimizationRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateOptimizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
