package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateOptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOptimizationRule{}

// CreateOptimizationRule struct for CreateOptimizationRule
type CreateOptimizationRule struct {
	// The state of the optimization rule.
	State string `json:"state"`
	// The name of the optimization rule.
	RuleName *string `json:"ruleName,omitempty"`
	// A list of rule conditions that define the advertiser's intent for the outcome of the rule. The rule uses 'AND' logic to combine every condition in this list, and will validate the combination when the rule is created or updated.
	RuleConditions []RuleCondition `json:"ruleConditions"`
}

type _CreateOptimizationRule CreateOptimizationRule

// NewCreateOptimizationRule instantiates a new CreateOptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOptimizationRule(state string, ruleConditions []RuleCondition) *CreateOptimizationRule {
	this := CreateOptimizationRule{}
	this.State = state
	this.RuleConditions = ruleConditions
	return &this
}

// NewCreateOptimizationRuleWithDefaults instantiates a new CreateOptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOptimizationRuleWithDefaults() *CreateOptimizationRule {
	this := CreateOptimizationRule{}
	return &this
}

// GetState returns the State field value
func (o *CreateOptimizationRule) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRule) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateOptimizationRule) SetState(v string) {
	o.State = v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *CreateOptimizationRule) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRule) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *CreateOptimizationRule) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *CreateOptimizationRule) SetRuleName(v string) {
	o.RuleName = &v
}

// GetRuleConditions returns the RuleConditions field value
func (o *CreateOptimizationRule) GetRuleConditions() []RuleCondition {
	if o == nil {
		var ret []RuleCondition
		return ret
	}

	return o.RuleConditions
}

// GetRuleConditionsOk returns a tuple with the RuleConditions field value
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRule) GetRuleConditionsOk() ([]RuleCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleConditions, true
}

// SetRuleConditions sets field value
func (o *CreateOptimizationRule) SetRuleConditions(v []RuleCondition) {
	o.RuleConditions = v
}

func (o CreateOptimizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	if !IsNil(o.RuleName) {
		toSerialize["ruleName"] = o.RuleName
	}
	toSerialize["ruleConditions"] = o.RuleConditions
	return toSerialize, nil
}

type NullableCreateOptimizationRule struct {
	value *CreateOptimizationRule
	isSet bool
}

func (v NullableCreateOptimizationRule) Get() *CreateOptimizationRule {
	return v.value
}

func (v *NullableCreateOptimizationRule) Set(val *CreateOptimizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOptimizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOptimizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOptimizationRule(val *CreateOptimizationRule) *NullableCreateOptimizationRule {
	return &NullableCreateOptimizationRule{value: val, isSet: true}
}

func (v NullableCreateOptimizationRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateOptimizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
