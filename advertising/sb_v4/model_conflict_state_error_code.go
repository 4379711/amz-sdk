package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ConflictStateErrorCode the model 'ConflictStateErrorCode'
type ConflictStateErrorCode string

// List of ConflictStateErrorCode
const (
	CONFLICTSTATEERRORCODE_CONFLICT_STATE ConflictStateErrorCode = "CONFLICT_STATE"
)

// All allowed values of ConflictStateErrorCode enum
var AllowedConflictStateErrorCodeEnumValues = []ConflictStateErrorCode{
	"CONFLICT_STATE",
}

func (v *ConflictStateErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConflictStateErrorCode(value)
	for _, existing := range AllowedConflictStateErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConflictStateErrorCode", value)
}

// NewConflictStateErrorCodeFromValue returns a pointer to a valid ConflictStateErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConflictStateErrorCodeFromValue(v string) (*ConflictStateErrorCode, error) {
	ev := ConflictStateErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConflictStateErrorCode: valid values are %v", v, AllowedConflictStateErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConflictStateErrorCode) IsValid() bool {
	for _, existing := range AllowedConflictStateErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConflictStateErrorCode value
func (v ConflictStateErrorCode) Ptr() *ConflictStateErrorCode {
	return &v
}

type NullableConflictStateErrorCode struct {
	value *ConflictStateErrorCode
	isSet bool
}

func (v NullableConflictStateErrorCode) Get() *ConflictStateErrorCode {
	return v.value
}

func (v *NullableConflictStateErrorCode) Set(val *ConflictStateErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictStateErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictStateErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictStateErrorCode(val *ConflictStateErrorCode) *NullableConflictStateErrorCode {
	return &NullableConflictStateErrorCode{value: val, isSet: true}
}

func (v NullableConflictStateErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConflictStateErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
