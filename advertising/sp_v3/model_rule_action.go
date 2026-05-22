package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// RuleAction The action taken when the campaign optimization rule is enabled. Defaults to adopt
type RuleAction string

// List of RuleAction
const (
	RULEACTION_ADOPT RuleAction = "ADOPT"
)

// All allowed values of RuleAction enum
var AllowedRuleActionEnumValues = []RuleAction{
	"ADOPT",
}

func (v *RuleAction) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuleAction(value)
	for _, existing := range AllowedRuleActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuleAction", value)
}

// NewRuleActionFromValue returns a pointer to a valid RuleAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuleActionFromValue(v string) (*RuleAction, error) {
	ev := RuleAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuleAction: valid values are %v", v, AllowedRuleActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuleAction) IsValid() bool {
	for _, existing := range AllowedRuleActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleAction value
func (v RuleAction) Ptr() *RuleAction {
	return &v
}

type NullableRuleAction struct {
	value *RuleAction
	isSet bool
}

func (v NullableRuleAction) Get() *RuleAction {
	return v.value
}

func (v *NullableRuleAction) Set(val *RuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleAction(val *RuleAction) *NullableRuleAction {
	return &NullableRuleAction{value: val, isSet: true}
}

func (v NullableRuleAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
