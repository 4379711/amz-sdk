package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// NotFoundErrorCode the model 'NotFoundErrorCode'
type NotFoundErrorCode string

// List of NotFoundErrorCode
const (
	NOTFOUNDERRORCODE_NOT_FOUND NotFoundErrorCode = "NOT_FOUND"
)

// All allowed values of NotFoundErrorCode enum
var AllowedNotFoundErrorCodeEnumValues = []NotFoundErrorCode{
	"NOT_FOUND",
}

func (v *NotFoundErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotFoundErrorCode(value)
	for _, existing := range AllowedNotFoundErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotFoundErrorCode", value)
}

// NewNotFoundErrorCodeFromValue returns a pointer to a valid NotFoundErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotFoundErrorCodeFromValue(v string) (*NotFoundErrorCode, error) {
	ev := NotFoundErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotFoundErrorCode: valid values are %v", v, AllowedNotFoundErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotFoundErrorCode) IsValid() bool {
	for _, existing := range AllowedNotFoundErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotFoundErrorCode value
func (v NotFoundErrorCode) Ptr() *NotFoundErrorCode {
	return &v
}

type NullableNotFoundErrorCode struct {
	value *NotFoundErrorCode
	isSet bool
}

func (v NullableNotFoundErrorCode) Get() *NotFoundErrorCode {
	return v.value
}

func (v *NullableNotFoundErrorCode) Set(val *NotFoundErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundErrorCode(val *NotFoundErrorCode) *NullableNotFoundErrorCode {
	return &NullableNotFoundErrorCode{value: val, isSet: true}
}

func (v NullableNotFoundErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNotFoundErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
