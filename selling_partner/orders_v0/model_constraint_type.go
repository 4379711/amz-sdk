package orders_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ConstraintType Details the importance of the constraint present on the item
type ConstraintType string

// List of ConstraintType
const (
	CONSTRAINTTYPE_MANDATORY ConstraintType = "MANDATORY"
)

// All allowed values of ConstraintType enum
var AllowedConstraintTypeEnumValues = []ConstraintType{
	CONSTRAINTTYPE_MANDATORY,
}

func (v *ConstraintType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConstraintType(value)
	for _, existing := range AllowedConstraintTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConstraintType", value)
}

// NewConstraintTypeFromValue returns a pointer to a valid ConstraintType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConstraintTypeFromValue(v string) (*ConstraintType, error) {
	ev := ConstraintType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConstraintType: valid values are %v", v, AllowedConstraintTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConstraintType) IsValid() bool {
	for _, existing := range AllowedConstraintTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConstraintType value
func (v ConstraintType) Ptr() *ConstraintType {
	return &v
}

type NullableConstraintType struct {
	value *ConstraintType
	isSet bool
}

func (v NullableConstraintType) Get() *ConstraintType {
	return v.value
}

func (v *NullableConstraintType) Set(val *ConstraintType) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraintType) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraintType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraintType(val *ConstraintType) *NullableConstraintType {
	return &NullableConstraintType{value: val, isSet: true}
}

func (v NullableConstraintType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConstraintType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
