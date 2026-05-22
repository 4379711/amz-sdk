package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// RuleStatus The campaign optimization rule status. Read-Only
type RuleStatus string

// List of RuleStatus
const (
	RULESTATUS_ACTIVE   RuleStatus = "ACTIVE"
	RULESTATUS_ARCHIVED RuleStatus = "ARCHIVED"
)

// All allowed values of RuleStatus enum
var AllowedRuleStatusEnumValues = []RuleStatus{
	"ACTIVE",
	"ARCHIVED",
}

func (v *RuleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuleStatus(value)
	for _, existing := range AllowedRuleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuleStatus", value)
}

// NewRuleStatusFromValue returns a pointer to a valid RuleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuleStatusFromValue(v string) (*RuleStatus, error) {
	ev := RuleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuleStatus: valid values are %v", v, AllowedRuleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuleStatus) IsValid() bool {
	for _, existing := range AllowedRuleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleStatus value
func (v RuleStatus) Ptr() *RuleStatus {
	return &v
}

type NullableRuleStatus struct {
	value *RuleStatus
	isSet bool
}

func (v NullableRuleStatus) Get() *RuleStatus {
	return v.value
}

func (v *NullableRuleStatus) Set(val *RuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleStatus(val *RuleStatus) *NullableRuleStatus {
	return &NullableRuleStatus{value: val, isSet: true}
}

func (v NullableRuleStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
