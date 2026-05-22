package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// RuleType The type of the campaign optimization rule. Only Support BID as of now
type RuleType string

// List of RuleType
const (
	RULETYPE_BID     RuleType = "BID"
	RULETYPE_KEYWORD RuleType = "KEYWORD"
	RULETYPE_PRODUCT RuleType = "PRODUCT"
)

// All allowed values of RuleType enum
var AllowedRuleTypeEnumValues = []RuleType{
	"BID",
	"KEYWORD",
	"PRODUCT",
}

func (v *RuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuleType(value)
	for _, existing := range AllowedRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuleType", value)
}

// NewRuleTypeFromValue returns a pointer to a valid RuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuleTypeFromValue(v string) (*RuleType, error) {
	ev := RuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuleType: valid values are %v", v, AllowedRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuleType) IsValid() bool {
	for _, existing := range AllowedRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleType value
func (v RuleType) Ptr() *RuleType {
	return &v
}

type NullableRuleType struct {
	value *RuleType
	isSet bool
}

func (v NullableRuleType) Get() *RuleType {
	return v.value
}

func (v *NullableRuleType) Set(val *RuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleType(val *RuleType) *NullableRuleType {
	return &NullableRuleType{value: val, isSet: true}
}

func (v NullableRuleType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
