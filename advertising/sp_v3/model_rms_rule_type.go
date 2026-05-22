package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// RMSRuleType The type of budget rule. SCHEDULE: A budget rule based on a start and end date. PERFORMANCE: A budget rule based on advertising performance criteria.
type RMSRuleType string

// List of RMSRuleType
const (
	RMSRULETYPE_SCHEDULE    RMSRuleType = "SCHEDULE"
	RMSRULETYPE_PERFORMANCE RMSRuleType = "PERFORMANCE"
)

// All allowed values of RMSRuleType enum
var AllowedRMSRuleTypeEnumValues = []RMSRuleType{
	"SCHEDULE",
	"PERFORMANCE",
}

func (v *RMSRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RMSRuleType(value)
	for _, existing := range AllowedRMSRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RMSRuleType", value)
}

// NewRMSRuleTypeFromValue returns a pointer to a valid RMSRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRMSRuleTypeFromValue(v string) (*RMSRuleType, error) {
	ev := RMSRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RMSRuleType: valid values are %v", v, AllowedRMSRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RMSRuleType) IsValid() bool {
	for _, existing := range AllowedRMSRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RMSRuleType value
func (v RMSRuleType) Ptr() *RMSRuleType {
	return &v
}

type NullableRMSRuleType struct {
	value *RMSRuleType
	isSet bool
}

func (v NullableRMSRuleType) Get() *RMSRuleType {
	return v.value
}

func (v *NullableRMSRuleType) Set(val *RMSRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableRMSRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableRMSRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRMSRuleType(val *RMSRuleType) *NullableRMSRuleType {
	return &NullableRMSRuleType{value: val, isSet: true}
}

func (v NullableRMSRuleType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRMSRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
