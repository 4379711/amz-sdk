package transfers_20240601

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AssignmentType The default payment method type.
type AssignmentType string

// List of AssignmentType
const (
	ASSIGNMENTTYPE_DEFAULT_DEPOSIT_METHOD AssignmentType = "DEFAULT_DEPOSIT_METHOD"
)

// All allowed values of AssignmentType enum
var AllowedAssignmentTypeEnumValues = []AssignmentType{
	ASSIGNMENTTYPE_DEFAULT_DEPOSIT_METHOD,
}

func (v *AssignmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssignmentType(value)
	for _, existing := range AllowedAssignmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssignmentType", value)
}

// NewAssignmentTypeFromValue returns a pointer to a valid AssignmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssignmentTypeFromValue(v string) (*AssignmentType, error) {
	ev := AssignmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssignmentType: valid values are %v", v, AllowedAssignmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssignmentType) IsValid() bool {
	for _, existing := range AllowedAssignmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssignmentType value
func (v AssignmentType) Ptr() *AssignmentType {
	return &v
}

type NullableAssignmentType struct {
	value *AssignmentType
	isSet bool
}

func (v NullableAssignmentType) Get() *AssignmentType {
	return v.value
}

func (v *NullableAssignmentType) Set(val *AssignmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentType(val *AssignmentType) *NullableAssignmentType {
	return &NullableAssignmentType{value: val, isSet: true}
}

func (v NullableAssignmentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssignmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
