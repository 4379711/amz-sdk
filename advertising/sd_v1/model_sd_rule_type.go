package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SDRuleType The type of budget rule. SCHEDULE: A budget rule based on a start and end date. PERFORMANCE: A budget rule based on advertising performance criteria.
type SDRuleType string

// List of SDRuleType
const (
	SDRULETYPE_SCHEDULE    SDRuleType = "SCHEDULE"
	SDRULETYPE_PERFORMANCE SDRuleType = "PERFORMANCE"
)

// All allowed values of SDRuleType enum
var AllowedSDRuleTypeEnumValues = []SDRuleType{
	"SCHEDULE",
	"PERFORMANCE",
}

func (v *SDRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDRuleType(value)
	for _, existing := range AllowedSDRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDRuleType", value)
}

// NewSDRuleTypeFromValue returns a pointer to a valid SDRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDRuleTypeFromValue(v string) (*SDRuleType, error) {
	ev := SDRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDRuleType: valid values are %v", v, AllowedSDRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDRuleType) IsValid() bool {
	for _, existing := range AllowedSDRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDRuleType value
func (v SDRuleType) Ptr() *SDRuleType {
	return &v
}

type NullableSDRuleType struct {
	value *SDRuleType
	isSet bool
}

func (v NullableSDRuleType) Get() *SDRuleType {
	return v.value
}

func (v *NullableSDRuleType) Set(val *SDRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableSDRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableSDRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDRuleType(val *SDRuleType) *NullableSDRuleType {
	return &NullableSDRuleType{value: val, isSet: true}
}

func (v NullableSDRuleType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
