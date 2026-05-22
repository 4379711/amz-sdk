package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIValueTypeRuleCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIValueTypeRuleCriteria{}

// OptimizationRulesAPIValueTypeRuleCriteria Represents a criteria by comparing with the rule attribute value.
type OptimizationRulesAPIValueTypeRuleCriteria struct {
	ComparisonOperator OptimizationRulesAPIComparisonOperator `json:"comparisonOperator"`
	Value              float64                                `json:"value"`
}

type _OptimizationRulesAPIValueTypeRuleCriteria OptimizationRulesAPIValueTypeRuleCriteria

// NewOptimizationRulesAPIValueTypeRuleCriteria instantiates a new OptimizationRulesAPIValueTypeRuleCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIValueTypeRuleCriteria(comparisonOperator OptimizationRulesAPIComparisonOperator, value float64) *OptimizationRulesAPIValueTypeRuleCriteria {
	this := OptimizationRulesAPIValueTypeRuleCriteria{}
	this.ComparisonOperator = comparisonOperator
	this.Value = value
	return &this
}

// NewOptimizationRulesAPIValueTypeRuleCriteriaWithDefaults instantiates a new OptimizationRulesAPIValueTypeRuleCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIValueTypeRuleCriteriaWithDefaults() *OptimizationRulesAPIValueTypeRuleCriteria {
	this := OptimizationRulesAPIValueTypeRuleCriteria{}
	return &this
}

// GetComparisonOperator returns the ComparisonOperator field value
func (o *OptimizationRulesAPIValueTypeRuleCriteria) GetComparisonOperator() OptimizationRulesAPIComparisonOperator {
	if o == nil {
		var ret OptimizationRulesAPIComparisonOperator
		return ret
	}

	return o.ComparisonOperator
}

// GetComparisonOperatorOk returns a tuple with the ComparisonOperator field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIValueTypeRuleCriteria) GetComparisonOperatorOk() (*OptimizationRulesAPIComparisonOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonOperator, true
}

// SetComparisonOperator sets field value
func (o *OptimizationRulesAPIValueTypeRuleCriteria) SetComparisonOperator(v OptimizationRulesAPIComparisonOperator) {
	o.ComparisonOperator = v
}

// GetValue returns the Value field value
func (o *OptimizationRulesAPIValueTypeRuleCriteria) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIValueTypeRuleCriteria) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OptimizationRulesAPIValueTypeRuleCriteria) SetValue(v float64) {
	o.Value = v
}

func (o OptimizationRulesAPIValueTypeRuleCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comparisonOperator"] = o.ComparisonOperator
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableOptimizationRulesAPIValueTypeRuleCriteria struct {
	value *OptimizationRulesAPIValueTypeRuleCriteria
	isSet bool
}

func (v NullableOptimizationRulesAPIValueTypeRuleCriteria) Get() *OptimizationRulesAPIValueTypeRuleCriteria {
	return v.value
}

func (v *NullableOptimizationRulesAPIValueTypeRuleCriteria) Set(val *OptimizationRulesAPIValueTypeRuleCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIValueTypeRuleCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIValueTypeRuleCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIValueTypeRuleCriteria(val *OptimizationRulesAPIValueTypeRuleCriteria) *NullableOptimizationRulesAPIValueTypeRuleCriteria {
	return &NullableOptimizationRulesAPIValueTypeRuleCriteria{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIValueTypeRuleCriteria) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIValueTypeRuleCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
