package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRule{}

// OptimizationRule struct for OptimizationRule
type OptimizationRule struct {
	// The state of the optimization rule.
	State *string `json:"state,omitempty"`
	// The name of the optimization rule.
	RuleName *string `json:"ruleName,omitempty"`
	// A list of rule conditions that define the advertiser's intent for the outcome of the rule. The rule uses 'AND' logic to combine every condition in this list, and will validate the combination when the rule is created or updated.
	RuleConditions []RuleCondition `json:"ruleConditions,omitempty"`
	// The identifier of the optimization rule.
	RuleId *string `json:"ruleId,omitempty"`
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

// GetState returns the State field value if set, zero value otherwise.
func (o *OptimizationRule) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRule) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OptimizationRule) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OptimizationRule) SetState(v string) {
	o.State = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *OptimizationRule) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRule) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *OptimizationRule) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *OptimizationRule) SetRuleName(v string) {
	o.RuleName = &v
}

// GetRuleConditions returns the RuleConditions field value if set, zero value otherwise.
func (o *OptimizationRule) GetRuleConditions() []RuleCondition {
	if o == nil || IsNil(o.RuleConditions) {
		var ret []RuleCondition
		return ret
	}
	return o.RuleConditions
}

// GetRuleConditionsOk returns a tuple with the RuleConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRule) GetRuleConditionsOk() ([]RuleCondition, bool) {
	if o == nil || IsNil(o.RuleConditions) {
		return nil, false
	}
	return o.RuleConditions, true
}

// HasRuleConditions returns a boolean if a field has been set.
func (o *OptimizationRule) HasRuleConditions() bool {
	if o != nil && !IsNil(o.RuleConditions) {
		return true
	}

	return false
}

// SetRuleConditions gets a reference to the given []RuleCondition and assigns it to the RuleConditions field.
func (o *OptimizationRule) SetRuleConditions(v []RuleCondition) {
	o.RuleConditions = v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *OptimizationRule) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRule) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *OptimizationRule) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *OptimizationRule) SetRuleId(v string) {
	o.RuleId = &v
}

func (o OptimizationRule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
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
