package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRuleCondition{}

// OptimizationRulesAPIRuleCondition struct for OptimizationRulesAPIRuleCondition
type OptimizationRulesAPIRuleCondition struct {
	Criteria      *OptimizationRulesAPIRuleCriteria  `json:"criteria,omitempty"`
	AttributeName *OptimizationRulesAPIRuleAttribute `json:"attributeName,omitempty"`
}

// NewOptimizationRulesAPIRuleCondition instantiates a new OptimizationRulesAPIRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRuleCondition() *OptimizationRulesAPIRuleCondition {
	this := OptimizationRulesAPIRuleCondition{}
	return &this
}

// NewOptimizationRulesAPIRuleConditionWithDefaults instantiates a new OptimizationRulesAPIRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRuleConditionWithDefaults() *OptimizationRulesAPIRuleCondition {
	this := OptimizationRulesAPIRuleCondition{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleCondition) GetCriteria() OptimizationRulesAPIRuleCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret OptimizationRulesAPIRuleCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleCondition) GetCriteriaOk() (*OptimizationRulesAPIRuleCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleCondition) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given OptimizationRulesAPIRuleCriteria and assigns it to the Criteria field.
func (o *OptimizationRulesAPIRuleCondition) SetCriteria(v OptimizationRulesAPIRuleCriteria) {
	o.Criteria = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleCondition) GetAttributeName() OptimizationRulesAPIRuleAttribute {
	if o == nil || IsNil(o.AttributeName) {
		var ret OptimizationRulesAPIRuleAttribute
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleCondition) GetAttributeNameOk() (*OptimizationRulesAPIRuleAttribute, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleCondition) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given OptimizationRulesAPIRuleAttribute and assigns it to the AttributeName field.
func (o *OptimizationRulesAPIRuleCondition) SetAttributeName(v OptimizationRulesAPIRuleAttribute) {
	o.AttributeName = &v
}

func (o OptimizationRulesAPIRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRuleCondition struct {
	value *OptimizationRulesAPIRuleCondition
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleCondition) Get() *OptimizationRulesAPIRuleCondition {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleCondition) Set(val *OptimizationRulesAPIRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleCondition(val *OptimizationRulesAPIRuleCondition) *NullableOptimizationRulesAPIRuleCondition {
	return &NullableOptimizationRulesAPIRuleCondition{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleCondition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
