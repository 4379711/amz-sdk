package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleCondition{}

// RuleCondition struct for RuleCondition
type RuleCondition struct {
	Criteria ValueTypeRuleCriteria `json:"criteria"`
	// Enum: \"COST_PER_CLICK\"  The name of the attribute.   Supported rule metrics and corresponding supported comparisonOperators: | AttributeName                      |  ComparisonOperator       |  Description                                                                            | |------------------------------------|---------------------------|-----------------------------------------------------------------------------------------| | COST_PER_CLICK                     | LESS_THAN_OR_EQUAL_TO     | Maximize page visits while cost per click less than or equal to threshold.              |
	AttributeName string `json:"attributeName"`
}

type _RuleCondition RuleCondition

// NewRuleCondition instantiates a new RuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleCondition(criteria ValueTypeRuleCriteria, attributeName string) *RuleCondition {
	this := RuleCondition{}
	this.Criteria = criteria
	this.AttributeName = attributeName
	return &this
}

// NewRuleConditionWithDefaults instantiates a new RuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConditionWithDefaults() *RuleCondition {
	this := RuleCondition{}
	return &this
}

// GetCriteria returns the Criteria field value
func (o *RuleCondition) GetCriteria() ValueTypeRuleCriteria {
	if o == nil {
		var ret ValueTypeRuleCriteria
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *RuleCondition) GetCriteriaOk() (*ValueTypeRuleCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// SetCriteria sets field value
func (o *RuleCondition) SetCriteria(v ValueTypeRuleCriteria) {
	o.Criteria = v
}

// GetAttributeName returns the AttributeName field value
func (o *RuleCondition) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *RuleCondition) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *RuleCondition) SetAttributeName(v string) {
	o.AttributeName = v
}

func (o RuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["criteria"] = o.Criteria
	toSerialize["attributeName"] = o.AttributeName
	return toSerialize, nil
}

type NullableRuleCondition struct {
	value *RuleCondition
	isSet bool
}

func (v NullableRuleCondition) Get() *RuleCondition {
	return v.value
}

func (v *NullableRuleCondition) Set(val *RuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleCondition(val *RuleCondition) *NullableRuleCondition {
	return &NullableRuleCondition{value: val, isSet: true}
}

func (v NullableRuleCondition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
