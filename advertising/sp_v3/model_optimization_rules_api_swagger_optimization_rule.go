package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIOptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIOptimizationRule{}

// OptimizationRulesAPIOptimizationRule struct for OptimizationRulesAPIOptimizationRule
type OptimizationRulesAPIOptimizationRule struct {
	Recurrence      OptimizationRulesAPIRuleRecurrence  `json:"recurrence"`
	RuleCategory    OptimizationRulesAPIRuleCategory    `json:"ruleCategory"`
	RuleSubCategory OptimizationRulesAPIRuleSubCategory `json:"ruleSubCategory"`
	// The rule name.
	RuleName   *string                             `json:"ruleName,omitempty"`
	Action     OptimizationRulesAPIRuleAction      `json:"action"`
	Conditions []OptimizationRulesAPIRuleCondition `json:"conditions,omitempty"`
	Status     *OptimizationRulesAPIRuleStatus     `json:"status,omitempty"`
	// The rule identifier.
	OptimizationRuleId *string `json:"optimizationRuleId,omitempty"`
}

type _OptimizationRulesAPIOptimizationRule OptimizationRulesAPIOptimizationRule

// NewOptimizationRulesAPIOptimizationRule instantiates a new OptimizationRulesAPIOptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIOptimizationRule(recurrence OptimizationRulesAPIRuleRecurrence, ruleCategory OptimizationRulesAPIRuleCategory, ruleSubCategory OptimizationRulesAPIRuleSubCategory, action OptimizationRulesAPIRuleAction) *OptimizationRulesAPIOptimizationRule {
	this := OptimizationRulesAPIOptimizationRule{}
	this.Recurrence = recurrence
	this.RuleCategory = ruleCategory
	this.RuleSubCategory = ruleSubCategory
	this.Action = action
	return &this
}

// NewOptimizationRulesAPIOptimizationRuleWithDefaults instantiates a new OptimizationRulesAPIOptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIOptimizationRuleWithDefaults() *OptimizationRulesAPIOptimizationRule {
	this := OptimizationRulesAPIOptimizationRule{}
	return &this
}

// GetRecurrence returns the Recurrence field value
func (o *OptimizationRulesAPIOptimizationRule) GetRecurrence() OptimizationRulesAPIRuleRecurrence {
	if o == nil {
		var ret OptimizationRulesAPIRuleRecurrence
		return ret
	}

	return o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetRecurrenceOk() (*OptimizationRulesAPIRuleRecurrence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrence, true
}

// SetRecurrence sets field value
func (o *OptimizationRulesAPIOptimizationRule) SetRecurrence(v OptimizationRulesAPIRuleRecurrence) {
	o.Recurrence = v
}

// GetRuleCategory returns the RuleCategory field value
func (o *OptimizationRulesAPIOptimizationRule) GetRuleCategory() OptimizationRulesAPIRuleCategory {
	if o == nil {
		var ret OptimizationRulesAPIRuleCategory
		return ret
	}

	return o.RuleCategory
}

// GetRuleCategoryOk returns a tuple with the RuleCategory field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetRuleCategoryOk() (*OptimizationRulesAPIRuleCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleCategory, true
}

// SetRuleCategory sets field value
func (o *OptimizationRulesAPIOptimizationRule) SetRuleCategory(v OptimizationRulesAPIRuleCategory) {
	o.RuleCategory = v
}

// GetRuleSubCategory returns the RuleSubCategory field value
func (o *OptimizationRulesAPIOptimizationRule) GetRuleSubCategory() OptimizationRulesAPIRuleSubCategory {
	if o == nil {
		var ret OptimizationRulesAPIRuleSubCategory
		return ret
	}

	return o.RuleSubCategory
}

// GetRuleSubCategoryOk returns a tuple with the RuleSubCategory field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetRuleSubCategoryOk() (*OptimizationRulesAPIRuleSubCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleSubCategory, true
}

// SetRuleSubCategory sets field value
func (o *OptimizationRulesAPIOptimizationRule) SetRuleSubCategory(v OptimizationRulesAPIRuleSubCategory) {
	o.RuleSubCategory = v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRule) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRule) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *OptimizationRulesAPIOptimizationRule) SetRuleName(v string) {
	o.RuleName = &v
}

// GetAction returns the Action field value
func (o *OptimizationRulesAPIOptimizationRule) GetAction() OptimizationRulesAPIRuleAction {
	if o == nil {
		var ret OptimizationRulesAPIRuleAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetActionOk() (*OptimizationRulesAPIRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *OptimizationRulesAPIOptimizationRule) SetAction(v OptimizationRulesAPIRuleAction) {
	o.Action = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRule) GetConditions() []OptimizationRulesAPIRuleCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []OptimizationRulesAPIRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetConditionsOk() ([]OptimizationRulesAPIRuleCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []OptimizationRulesAPIRuleCondition and assigns it to the Conditions field.
func (o *OptimizationRulesAPIOptimizationRule) SetConditions(v []OptimizationRulesAPIRuleCondition) {
	o.Conditions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRule) GetStatus() OptimizationRulesAPIRuleStatus {
	if o == nil || IsNil(o.Status) {
		var ret OptimizationRulesAPIRuleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetStatusOk() (*OptimizationRulesAPIRuleStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRule) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OptimizationRulesAPIRuleStatus and assigns it to the Status field.
func (o *OptimizationRulesAPIOptimizationRule) SetStatus(v OptimizationRulesAPIRuleStatus) {
	o.Status = &v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *OptimizationRulesAPIOptimizationRule) GetOptimizationRuleId() string {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret string
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIOptimizationRule) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *OptimizationRulesAPIOptimizationRule) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given string and assigns it to the OptimizationRuleId field.
func (o *OptimizationRulesAPIOptimizationRule) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = &v
}

func (o OptimizationRulesAPIOptimizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recurrence"] = o.Recurrence
	toSerialize["ruleCategory"] = o.RuleCategory
	toSerialize["ruleSubCategory"] = o.RuleSubCategory
	if !IsNil(o.RuleName) {
		toSerialize["ruleName"] = o.RuleName
	}
	toSerialize["action"] = o.Action
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIOptimizationRule struct {
	value *OptimizationRulesAPIOptimizationRule
	isSet bool
}

func (v NullableOptimizationRulesAPIOptimizationRule) Get() *OptimizationRulesAPIOptimizationRule {
	return v.value
}

func (v *NullableOptimizationRulesAPIOptimizationRule) Set(val *OptimizationRulesAPIOptimizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIOptimizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIOptimizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIOptimizationRule(val *OptimizationRulesAPIOptimizationRule) *NullableOptimizationRulesAPIOptimizationRule {
	return &NullableOptimizationRulesAPIOptimizationRule{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIOptimizationRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIOptimizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
