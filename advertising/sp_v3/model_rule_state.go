package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// RuleState The campaign optimization rule state.
type RuleState string

// List of RuleState
const (
	RULESTATE_ENABLED  RuleState = "ENABLED"
	RULESTATE_DISABLED RuleState = "DISABLED"
)

// All allowed values of RuleState enum
var AllowedRuleStateEnumValues = []RuleState{
	"ENABLED",
	"DISABLED",
}

func (v *RuleState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuleState(value)
	for _, existing := range AllowedRuleStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuleState", value)
}

// NewRuleStateFromValue returns a pointer to a valid RuleState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuleStateFromValue(v string) (*RuleState, error) {
	ev := RuleState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuleState: valid values are %v", v, AllowedRuleStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuleState) IsValid() bool {
	for _, existing := range AllowedRuleStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleState value
func (v RuleState) Ptr() *RuleState {
	return &v
}

type NullableRuleState struct {
	value *RuleState
	isSet bool
}

func (v NullableRuleState) Get() *RuleState {
	return v.value
}

func (v *NullableRuleState) Set(val *RuleState) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleState) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleState(val *RuleState) *NullableRuleState {
	return &NullableRuleState{value: val, isSet: true}
}

func (v NullableRuleState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
