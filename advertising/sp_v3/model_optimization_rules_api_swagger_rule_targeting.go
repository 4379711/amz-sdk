package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRuleTargeting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRuleTargeting{}

// OptimizationRulesAPIRuleTargeting struct for OptimizationRulesAPIRuleTargeting
type OptimizationRulesAPIRuleTargeting struct {
	ExpressionTypes []OptimizationRulesAPIExpressionType `json:"expressionTypes"`
	TargetingType   OptimizationRulesAPITargetingType    `json:"targetingType"`
	// The number of days of data to look back on for the rule.
	LookbackDays int32 `json:"lookbackDays"`
}

type _OptimizationRulesAPIRuleTargeting OptimizationRulesAPIRuleTargeting

// NewOptimizationRulesAPIRuleTargeting instantiates a new OptimizationRulesAPIRuleTargeting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRuleTargeting(expressionTypes []OptimizationRulesAPIExpressionType, targetingType OptimizationRulesAPITargetingType, lookbackDays int32) *OptimizationRulesAPIRuleTargeting {
	this := OptimizationRulesAPIRuleTargeting{}
	this.ExpressionTypes = expressionTypes
	this.TargetingType = targetingType
	this.LookbackDays = lookbackDays
	return &this
}

// NewOptimizationRulesAPIRuleTargetingWithDefaults instantiates a new OptimizationRulesAPIRuleTargeting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRuleTargetingWithDefaults() *OptimizationRulesAPIRuleTargeting {
	this := OptimizationRulesAPIRuleTargeting{}
	return &this
}

// GetExpressionTypes returns the ExpressionTypes field value
func (o *OptimizationRulesAPIRuleTargeting) GetExpressionTypes() []OptimizationRulesAPIExpressionType {
	if o == nil {
		var ret []OptimizationRulesAPIExpressionType
		return ret
	}

	return o.ExpressionTypes
}

// GetExpressionTypesOk returns a tuple with the ExpressionTypes field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleTargeting) GetExpressionTypesOk() ([]OptimizationRulesAPIExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpressionTypes, true
}

// SetExpressionTypes sets field value
func (o *OptimizationRulesAPIRuleTargeting) SetExpressionTypes(v []OptimizationRulesAPIExpressionType) {
	o.ExpressionTypes = v
}

// GetTargetingType returns the TargetingType field value
func (o *OptimizationRulesAPIRuleTargeting) GetTargetingType() OptimizationRulesAPITargetingType {
	if o == nil {
		var ret OptimizationRulesAPITargetingType
		return ret
	}

	return o.TargetingType
}

// GetTargetingTypeOk returns a tuple with the TargetingType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleTargeting) GetTargetingTypeOk() (*OptimizationRulesAPITargetingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingType, true
}

// SetTargetingType sets field value
func (o *OptimizationRulesAPIRuleTargeting) SetTargetingType(v OptimizationRulesAPITargetingType) {
	o.TargetingType = v
}

// GetLookbackDays returns the LookbackDays field value
func (o *OptimizationRulesAPIRuleTargeting) GetLookbackDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LookbackDays
}

// GetLookbackDaysOk returns a tuple with the LookbackDays field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleTargeting) GetLookbackDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookbackDays, true
}

// SetLookbackDays sets field value
func (o *OptimizationRulesAPIRuleTargeting) SetLookbackDays(v int32) {
	o.LookbackDays = v
}

func (o OptimizationRulesAPIRuleTargeting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expressionTypes"] = o.ExpressionTypes
	toSerialize["targetingType"] = o.TargetingType
	toSerialize["lookbackDays"] = o.LookbackDays
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRuleTargeting struct {
	value *OptimizationRulesAPIRuleTargeting
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleTargeting) Get() *OptimizationRulesAPIRuleTargeting {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleTargeting) Set(val *OptimizationRulesAPIRuleTargeting) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleTargeting) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleTargeting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleTargeting(val *OptimizationRulesAPIRuleTargeting) *NullableOptimizationRulesAPIRuleTargeting {
	return &NullableOptimizationRulesAPIRuleTargeting{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleTargeting) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleTargeting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
