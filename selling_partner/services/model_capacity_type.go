package services

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CapacityType Type of capacity
type CapacityType string

// List of CapacityType
const (
	CAPACITYTYPE_SCHEDULED_CAPACITY  CapacityType = "SCHEDULED_CAPACITY"
	CAPACITYTYPE_AVAILABLE_CAPACITY  CapacityType = "AVAILABLE_CAPACITY"
	CAPACITYTYPE_ENCUMBERED_CAPACITY CapacityType = "ENCUMBERED_CAPACITY"
	CAPACITYTYPE_RESERVED_CAPACITY   CapacityType = "RESERVED_CAPACITY"
)

// All allowed values of CapacityType enum
var AllowedCapacityTypeEnumValues = []CapacityType{
	CAPACITYTYPE_SCHEDULED_CAPACITY,
	CAPACITYTYPE_AVAILABLE_CAPACITY,
	CAPACITYTYPE_ENCUMBERED_CAPACITY,
	CAPACITYTYPE_RESERVED_CAPACITY,
}

func (v *CapacityType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CapacityType(value)
	for _, existing := range AllowedCapacityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CapacityType", value)
}

// NewCapacityTypeFromValue returns a pointer to a valid CapacityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCapacityTypeFromValue(v string) (*CapacityType, error) {
	ev := CapacityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CapacityType: valid values are %v", v, AllowedCapacityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CapacityType) IsValid() bool {
	for _, existing := range AllowedCapacityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CapacityType value
func (v CapacityType) Ptr() *CapacityType {
	return &v
}

type NullableCapacityType struct {
	value *CapacityType
	isSet bool
}

func (v NullableCapacityType) Get() *CapacityType {
	return v.value
}

func (v *NullableCapacityType) Set(val *CapacityType) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityType) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityType(val *CapacityType) *NullableCapacityType {
	return &NullableCapacityType{value: val, isSet: true}
}

func (v NullableCapacityType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCapacityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
