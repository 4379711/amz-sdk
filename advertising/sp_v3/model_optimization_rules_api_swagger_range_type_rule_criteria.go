package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRangeTypeRuleCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRangeTypeRuleCriteria{}

// OptimizationRulesAPIRangeTypeRuleCriteria Represents the range of rule attribute value.
type OptimizationRulesAPIRangeTypeRuleCriteria struct {
	MinValue float64 `json:"minValue"`
	MaxValue float64 `json:"maxValue"`
}

type _OptimizationRulesAPIRangeTypeRuleCriteria OptimizationRulesAPIRangeTypeRuleCriteria

// NewOptimizationRulesAPIRangeTypeRuleCriteria instantiates a new OptimizationRulesAPIRangeTypeRuleCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRangeTypeRuleCriteria(minValue float64, maxValue float64) *OptimizationRulesAPIRangeTypeRuleCriteria {
	this := OptimizationRulesAPIRangeTypeRuleCriteria{}
	this.MinValue = minValue
	this.MaxValue = maxValue
	return &this
}

// NewOptimizationRulesAPIRangeTypeRuleCriteriaWithDefaults instantiates a new OptimizationRulesAPIRangeTypeRuleCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRangeTypeRuleCriteriaWithDefaults() *OptimizationRulesAPIRangeTypeRuleCriteria {
	this := OptimizationRulesAPIRangeTypeRuleCriteria{}
	return &this
}

// GetMinValue returns the MinValue field value
func (o *OptimizationRulesAPIRangeTypeRuleCriteria) GetMinValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRangeTypeRuleCriteria) GetMinValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinValue, true
}

// SetMinValue sets field value
func (o *OptimizationRulesAPIRangeTypeRuleCriteria) SetMinValue(v float64) {
	o.MinValue = v
}

// GetMaxValue returns the MaxValue field value
func (o *OptimizationRulesAPIRangeTypeRuleCriteria) GetMaxValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRangeTypeRuleCriteria) GetMaxValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxValue, true
}

// SetMaxValue sets field value
func (o *OptimizationRulesAPIRangeTypeRuleCriteria) SetMaxValue(v float64) {
	o.MaxValue = v
}

func (o OptimizationRulesAPIRangeTypeRuleCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minValue"] = o.MinValue
	toSerialize["maxValue"] = o.MaxValue
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRangeTypeRuleCriteria struct {
	value *OptimizationRulesAPIRangeTypeRuleCriteria
	isSet bool
}

func (v NullableOptimizationRulesAPIRangeTypeRuleCriteria) Get() *OptimizationRulesAPIRangeTypeRuleCriteria {
	return v.value
}

func (v *NullableOptimizationRulesAPIRangeTypeRuleCriteria) Set(val *OptimizationRulesAPIRangeTypeRuleCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRangeTypeRuleCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRangeTypeRuleCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRangeTypeRuleCriteria(val *OptimizationRulesAPIRangeTypeRuleCriteria) *NullableOptimizationRulesAPIRangeTypeRuleCriteria {
	return &NullableOptimizationRulesAPIRangeTypeRuleCriteria{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRangeTypeRuleCriteria) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRangeTypeRuleCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
