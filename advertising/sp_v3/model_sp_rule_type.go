package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SPRuleType The type of budget rule. SCHEDULE: A budget rule based on a start and end date. PERFORMANCE: A budget rule based on advertising performance criteria.
type SPRuleType string

// List of SPRuleType
const (
	SPRULETYPE_SCHEDULE    SPRuleType = "SCHEDULE"
	SPRULETYPE_PERFORMANCE SPRuleType = "PERFORMANCE"
)

// All allowed values of SPRuleType enum
var AllowedSPRuleTypeEnumValues = []SPRuleType{
	"SCHEDULE",
	"PERFORMANCE",
}

func (v *SPRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SPRuleType(value)
	for _, existing := range AllowedSPRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SPRuleType", value)
}

// NewSPRuleTypeFromValue returns a pointer to a valid SPRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSPRuleTypeFromValue(v string) (*SPRuleType, error) {
	ev := SPRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SPRuleType: valid values are %v", v, AllowedSPRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SPRuleType) IsValid() bool {
	for _, existing := range AllowedSPRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SPRuleType value
func (v SPRuleType) Ptr() *SPRuleType {
	return &v
}

type NullableSPRuleType struct {
	value *SPRuleType
	isSet bool
}

func (v NullableSPRuleType) Get() *SPRuleType {
	return v.value
}

func (v *NullableSPRuleType) Set(val *SPRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableSPRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableSPRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPRuleType(val *SPRuleType) *NullableSPRuleType {
	return &NullableSPRuleType{value: val, isSet: true}
}

func (v NullableSPRuleType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
