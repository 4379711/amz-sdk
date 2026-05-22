package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIActionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIActionDetails{}

// OptimizationRulesAPIActionDetails Details of a rule action.
type OptimizationRulesAPIActionDetails struct {
	ActionUnit     string                                 `json:"actionUnit"`
	Value          float64                                `json:"value"`
	ActionOperator OptimizationRulesAPIRuleActionOperator `json:"actionOperator"`
}

type _OptimizationRulesAPIActionDetails OptimizationRulesAPIActionDetails

// NewOptimizationRulesAPIActionDetails instantiates a new OptimizationRulesAPIActionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIActionDetails(actionUnit string, value float64, actionOperator OptimizationRulesAPIRuleActionOperator) *OptimizationRulesAPIActionDetails {
	this := OptimizationRulesAPIActionDetails{}
	this.ActionUnit = actionUnit
	this.Value = value
	this.ActionOperator = actionOperator
	return &this
}

// NewOptimizationRulesAPIActionDetailsWithDefaults instantiates a new OptimizationRulesAPIActionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIActionDetailsWithDefaults() *OptimizationRulesAPIActionDetails {
	this := OptimizationRulesAPIActionDetails{}
	return &this
}

// GetActionUnit returns the ActionUnit field value
func (o *OptimizationRulesAPIActionDetails) GetActionUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionUnit
}

// GetActionUnitOk returns a tuple with the ActionUnit field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIActionDetails) GetActionUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionUnit, true
}

// SetActionUnit sets field value
func (o *OptimizationRulesAPIActionDetails) SetActionUnit(v string) {
	o.ActionUnit = v
}

// GetValue returns the Value field value
func (o *OptimizationRulesAPIActionDetails) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIActionDetails) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OptimizationRulesAPIActionDetails) SetValue(v float64) {
	o.Value = v
}

// GetActionOperator returns the ActionOperator field value
func (o *OptimizationRulesAPIActionDetails) GetActionOperator() OptimizationRulesAPIRuleActionOperator {
	if o == nil {
		var ret OptimizationRulesAPIRuleActionOperator
		return ret
	}

	return o.ActionOperator
}

// GetActionOperatorOk returns a tuple with the ActionOperator field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIActionDetails) GetActionOperatorOk() (*OptimizationRulesAPIRuleActionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionOperator, true
}

// SetActionOperator sets field value
func (o *OptimizationRulesAPIActionDetails) SetActionOperator(v OptimizationRulesAPIRuleActionOperator) {
	o.ActionOperator = v
}

func (o OptimizationRulesAPIActionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionUnit"] = o.ActionUnit
	toSerialize["value"] = o.Value
	toSerialize["actionOperator"] = o.ActionOperator
	return toSerialize, nil
}

type NullableOptimizationRulesAPIActionDetails struct {
	value *OptimizationRulesAPIActionDetails
	isSet bool
}

func (v NullableOptimizationRulesAPIActionDetails) Get() *OptimizationRulesAPIActionDetails {
	return v.value
}

func (v *NullableOptimizationRulesAPIActionDetails) Set(val *OptimizationRulesAPIActionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIActionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIActionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIActionDetails(val *OptimizationRulesAPIActionDetails) *NullableOptimizationRulesAPIActionDetails {
	return &NullableOptimizationRulesAPIActionDetails{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIActionDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIActionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
