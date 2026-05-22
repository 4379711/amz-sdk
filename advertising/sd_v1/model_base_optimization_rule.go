package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BaseOptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseOptimizationRule{}

// BaseOptimizationRule struct for BaseOptimizationRule
type BaseOptimizationRule struct {
	// The state of the optimization rule.
	State *string `json:"state,omitempty"`
	// The name of the optimization rule.
	RuleName *string `json:"ruleName,omitempty"`
	// A list of rule conditions that define the advertiser's intent for the outcome of the rule. The rule uses 'AND' logic to combine every condition in this list, and will validate the combination when the rule is created or updated.
	RuleConditions []RuleCondition `json:"ruleConditions,omitempty"`
}

// NewBaseOptimizationRule instantiates a new BaseOptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseOptimizationRule() *BaseOptimizationRule {
	this := BaseOptimizationRule{}
	return &this
}

// NewBaseOptimizationRuleWithDefaults instantiates a new BaseOptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseOptimizationRuleWithDefaults() *BaseOptimizationRule {
	this := BaseOptimizationRule{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BaseOptimizationRule) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOptimizationRule) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BaseOptimizationRule) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BaseOptimizationRule) SetState(v string) {
	o.State = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *BaseOptimizationRule) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOptimizationRule) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *BaseOptimizationRule) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *BaseOptimizationRule) SetRuleName(v string) {
	o.RuleName = &v
}

// GetRuleConditions returns the RuleConditions field value if set, zero value otherwise.
func (o *BaseOptimizationRule) GetRuleConditions() []RuleCondition {
	if o == nil || IsNil(o.RuleConditions) {
		var ret []RuleCondition
		return ret
	}
	return o.RuleConditions
}

// GetRuleConditionsOk returns a tuple with the RuleConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOptimizationRule) GetRuleConditionsOk() ([]RuleCondition, bool) {
	if o == nil || IsNil(o.RuleConditions) {
		return nil, false
	}
	return o.RuleConditions, true
}

// HasRuleConditions returns a boolean if a field has been set.
func (o *BaseOptimizationRule) HasRuleConditions() bool {
	if o != nil && !IsNil(o.RuleConditions) {
		return true
	}

	return false
}

// SetRuleConditions gets a reference to the given []RuleCondition and assigns it to the RuleConditions field.
func (o *BaseOptimizationRule) SetRuleConditions(v []RuleCondition) {
	o.RuleConditions = v
}

func (o BaseOptimizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.RuleName) {
		toSerialize["ruleName"] = o.RuleName
	}
	if !IsNil(o.RuleConditions) {
		toSerialize["ruleConditions"] = o.RuleConditions
	}
	return toSerialize, nil
}

type NullableBaseOptimizationRule struct {
	value *BaseOptimizationRule
	isSet bool
}

func (v NullableBaseOptimizationRule) Get() *BaseOptimizationRule {
	return v.value
}

func (v *NullableBaseOptimizationRule) Set(val *BaseOptimizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseOptimizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseOptimizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseOptimizationRule(val *BaseOptimizationRule) *NullableBaseOptimizationRule {
	return &NullableBaseOptimizationRule{value: val, isSet: true}
}

func (v NullableBaseOptimizationRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBaseOptimizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
