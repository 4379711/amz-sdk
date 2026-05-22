package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRuleAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRuleAction{}

// OptimizationRulesAPIRuleAction Action to be taken by the rule.
type OptimizationRulesAPIRuleAction struct {
	ActionType    OptimizationRulesAPIActionType    `json:"actionType"`
	ActionDetails OptimizationRulesAPIActionDetails `json:"actionDetails"`
}

type _OptimizationRulesAPIRuleAction OptimizationRulesAPIRuleAction

// NewOptimizationRulesAPIRuleAction instantiates a new OptimizationRulesAPIRuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRuleAction(actionType OptimizationRulesAPIActionType, actionDetails OptimizationRulesAPIActionDetails) *OptimizationRulesAPIRuleAction {
	this := OptimizationRulesAPIRuleAction{}
	this.ActionType = actionType
	this.ActionDetails = actionDetails
	return &this
}

// NewOptimizationRulesAPIRuleActionWithDefaults instantiates a new OptimizationRulesAPIRuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRuleActionWithDefaults() *OptimizationRulesAPIRuleAction {
	this := OptimizationRulesAPIRuleAction{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *OptimizationRulesAPIRuleAction) GetActionType() OptimizationRulesAPIActionType {
	if o == nil {
		var ret OptimizationRulesAPIActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleAction) GetActionTypeOk() (*OptimizationRulesAPIActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *OptimizationRulesAPIRuleAction) SetActionType(v OptimizationRulesAPIActionType) {
	o.ActionType = v
}

// GetActionDetails returns the ActionDetails field value
func (o *OptimizationRulesAPIRuleAction) GetActionDetails() OptimizationRulesAPIActionDetails {
	if o == nil {
		var ret OptimizationRulesAPIActionDetails
		return ret
	}

	return o.ActionDetails
}

// GetActionDetailsOk returns a tuple with the ActionDetails field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleAction) GetActionDetailsOk() (*OptimizationRulesAPIActionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionDetails, true
}

// SetActionDetails sets field value
func (o *OptimizationRulesAPIRuleAction) SetActionDetails(v OptimizationRulesAPIActionDetails) {
	o.ActionDetails = v
}

func (o OptimizationRulesAPIRuleAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionType"] = o.ActionType
	toSerialize["actionDetails"] = o.ActionDetails
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRuleAction struct {
	value *OptimizationRulesAPIRuleAction
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleAction) Get() *OptimizationRulesAPIRuleAction {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleAction) Set(val *OptimizationRulesAPIRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleAction(val *OptimizationRulesAPIRuleAction) *NullableOptimizationRulesAPIRuleAction {
	return &NullableOptimizationRulesAPIRuleAction{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
