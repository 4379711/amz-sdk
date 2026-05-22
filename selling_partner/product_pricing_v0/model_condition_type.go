package product_pricing_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ConditionType Indicates the condition of the item. Possible values: New, Used, Collectible, Refurbished, Club.
type ConditionType string

// List of ConditionType
const (
	CONDITIONTYPE_NEW         ConditionType = "New"
	CONDITIONTYPE_USED        ConditionType = "Used"
	CONDITIONTYPE_COLLECTIBLE ConditionType = "Collectible"
	CONDITIONTYPE_REFURBISHED ConditionType = "Refurbished"
	CONDITIONTYPE_CLUB        ConditionType = "Club"
)

// All allowed values of ConditionType enum
var AllowedConditionTypeEnumValues = []ConditionType{
	CONDITIONTYPE_NEW,
	CONDITIONTYPE_USED,
	CONDITIONTYPE_COLLECTIBLE,
	CONDITIONTYPE_REFURBISHED,
	CONDITIONTYPE_CLUB,
}

func (v *ConditionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConditionType(value)
	for _, existing := range AllowedConditionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConditionType", value)
}

// NewConditionTypeFromValue returns a pointer to a valid ConditionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionTypeFromValue(v string) (*ConditionType, error) {
	ev := ConditionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConditionType: valid values are %v", v, AllowedConditionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConditionType) IsValid() bool {
	for _, existing := range AllowedConditionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConditionType value
func (v ConditionType) Ptr() *ConditionType {
	return &v
}

type NullableConditionType struct {
	value *ConditionType
	isSet bool
}

func (v NullableConditionType) Get() *ConditionType {
	return v.value
}

func (v *NullableConditionType) Set(val *ConditionType) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionType) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionType(val *ConditionType) *NullableConditionType {
	return &NullableConditionType{value: val, isSet: true}
}

func (v NullableConditionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConditionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
