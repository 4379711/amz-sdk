package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateOptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOptimizationRule{}

// UpdateOptimizationRule struct for UpdateOptimizationRule
type UpdateOptimizationRule struct {
	// The state of the optimization rule.
	State *string `json:"state,omitempty"`
	// The name of the optimization rule.
	RuleName *string `json:"ruleName,omitempty"`
	// A list of rule conditions that define the advertiser's intent for the outcome of the rule. The rule uses 'AND' logic to combine every condition in this list, and will validate the combination when the rule is created or updated.
	RuleConditions []RuleCondition `json:"ruleConditions,omitempty"`
	// The identifier of the optimization rule.
	RuleId string `json:"ruleId"`
}

type _UpdateOptimizationRule UpdateOptimizationRule

// NewUpdateOptimizationRule instantiates a new UpdateOptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOptimizationRule(ruleId string) *UpdateOptimizationRule {
	this := UpdateOptimizationRule{}
	this.RuleId = ruleId
	return &this
}

// NewUpdateOptimizationRuleWithDefaults instantiates a new UpdateOptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOptimizationRuleWithDefaults() *UpdateOptimizationRule {
	this := UpdateOptimizationRule{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateOptimizationRule) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRule) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateOptimizationRule) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateOptimizationRule) SetState(v string) {
	o.State = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *UpdateOptimizationRule) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRule) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *UpdateOptimizationRule) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *UpdateOptimizationRule) SetRuleName(v string) {
	o.RuleName = &v
}

// GetRuleConditions returns the RuleConditions field value if set, zero value otherwise.
func (o *UpdateOptimizationRule) GetRuleConditions() []RuleCondition {
	if o == nil || IsNil(o.RuleConditions) {
		var ret []RuleCondition
		return ret
	}
	return o.RuleConditions
}

// GetRuleConditionsOk returns a tuple with the RuleConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRule) GetRuleConditionsOk() ([]RuleCondition, bool) {
	if o == nil || IsNil(o.RuleConditions) {
		return nil, false
	}
	return o.RuleConditions, true
}

// HasRuleConditions returns a boolean if a field has been set.
func (o *UpdateOptimizationRule) HasRuleConditions() bool {
	if o != nil && !IsNil(o.RuleConditions) {
		return true
	}

	return false
}

// SetRuleConditions gets a reference to the given []RuleCondition and assigns it to the RuleConditions field.
func (o *UpdateOptimizationRule) SetRuleConditions(v []RuleCondition) {
	o.RuleConditions = v
}

// GetRuleId returns the RuleId field value
func (o *UpdateOptimizationRule) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRule) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *UpdateOptimizationRule) SetRuleId(v string) {
	o.RuleId = v
}

func (o UpdateOptimizationRule) ToMap() (map[string]interface{}, error) {
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
	toSerialize["ruleId"] = o.RuleId
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
