package aplus_content_20201101

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PositionType The content's relative positioning.
type PositionType string

// List of PositionType
const (
	POSITIONTYPE_LEFT  PositionType = "LEFT"
	POSITIONTYPE_RIGHT PositionType = "RIGHT"
)

// All allowed values of PositionType enum
var AllowedPositionTypeEnumValues = []PositionType{
	POSITIONTYPE_LEFT,
	POSITIONTYPE_RIGHT,
}

func (v *PositionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PositionType(value)
	for _, existing := range AllowedPositionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PositionType", value)
}

// NewPositionTypeFromValue returns a pointer to a valid PositionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPositionTypeFromValue(v string) (*PositionType, error) {
	ev := PositionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PositionType: valid values are %v", v, AllowedPositionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PositionType) IsValid() bool {
	for _, existing := range AllowedPositionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PositionType value
func (v PositionType) Ptr() *PositionType {
	return &v
}

type NullablePositionType struct {
	value *PositionType
	isSet bool
}

func (v NullablePositionType) Get() *PositionType {
	return v.value
}

func (v *NullablePositionType) Set(val *PositionType) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionType) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionType(val *PositionType) *NullablePositionType {
	return &NullablePositionType{value: val, isSet: true}
}

func (v NullablePositionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePositionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
