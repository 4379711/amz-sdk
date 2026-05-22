package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SBRuleType The type of budget rule. SCHEDULE: A budget rule based on a start and end date. PERFORMANCE: A budget rule based on advertising performance criteria.
type SBRuleType string

// List of SBRuleType
const (
	SBRULETYPE_SCHEDULE    SBRuleType = "SCHEDULE"
	SBRULETYPE_PERFORMANCE SBRuleType = "PERFORMANCE"
)

// All allowed values of SBRuleType enum
var AllowedSBRuleTypeEnumValues = []SBRuleType{
	"SCHEDULE",
	"PERFORMANCE",
}

func (v *SBRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SBRuleType(value)
	for _, existing := range AllowedSBRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SBRuleType", value)
}

// NewSBRuleTypeFromValue returns a pointer to a valid SBRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSBRuleTypeFromValue(v string) (*SBRuleType, error) {
	ev := SBRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SBRuleType: valid values are %v", v, AllowedSBRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SBRuleType) IsValid() bool {
	for _, existing := range AllowedSBRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBRuleType value
func (v SBRuleType) Ptr() *SBRuleType {
	return &v
}

type NullableSBRuleType struct {
	value *SBRuleType
	isSet bool
}

func (v NullableSBRuleType) Get() *SBRuleType {
	return v.value
}

func (v *NullableSBRuleType) Set(val *SBRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableSBRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableSBRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBRuleType(val *SBRuleType) *NullableSBRuleType {
	return &NullableSBRuleType{value: val, isSet: true}
}

func (v NullableSBRuleType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
