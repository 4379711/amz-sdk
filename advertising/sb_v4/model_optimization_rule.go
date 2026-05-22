package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRule{}

// OptimizationRule struct for OptimizationRule
type OptimizationRule struct {
	// The identifier of the optimization rule.
	OptimizationRuleId *string         `json:"optimizationRuleId,omitempty"`
	Conditions         []RuleCondition `json:"conditions,omitempty"`
}

// NewOptimizationRule instantiates a new OptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRule() *OptimizationRule {
	this := OptimizationRule{}
	return &this
}

// NewOptimizationRuleWithDefaults instantiates a new OptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRuleWithDefaults() *OptimizationRule {
	this := OptimizationRule{}
	return &this
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *OptimizationRule) GetOptimizationRuleId() string {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret string
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRule) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *OptimizationRule) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given string and assigns it to the OptimizationRuleId field.
func (o *OptimizationRule) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OptimizationRule) GetConditions() []RuleCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []RuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRule) GetConditionsOk() ([]RuleCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OptimizationRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []RuleCondition and assigns it to the Conditions field.
func (o *OptimizationRule) SetConditions(v []RuleCondition) {
	o.Conditions = v
}

func (o OptimizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableOptimizationRule struct {
	value *OptimizationRule
	isSet bool
}

func (v NullableOptimizationRule) Get() *OptimizationRule {
	return v.value
}

func (v *NullableOptimizationRule) Set(val *OptimizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRule(val *OptimizationRule) *NullableOptimizationRule {
	return &NullableOptimizationRule{value: val, isSet: true}
}

func (v NullableOptimizationRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
