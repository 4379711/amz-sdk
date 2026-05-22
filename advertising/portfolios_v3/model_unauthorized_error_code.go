package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// UnauthorizedErrorCode the model 'UnauthorizedErrorCode'
type UnauthorizedErrorCode string

// List of UnauthorizedErrorCode
const (
	UNAUTHORIZEDERRORCODE_UNAUTHORIZED UnauthorizedErrorCode = "UNAUTHORIZED"
)

// All allowed values of UnauthorizedErrorCode enum
var AllowedUnauthorizedErrorCodeEnumValues = []UnauthorizedErrorCode{
	"UNAUTHORIZED",
}

func (v *UnauthorizedErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnauthorizedErrorCode(value)
	for _, existing := range AllowedUnauthorizedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnauthorizedErrorCode", value)
}

// NewUnauthorizedErrorCodeFromValue returns a pointer to a valid UnauthorizedErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnauthorizedErrorCodeFromValue(v string) (*UnauthorizedErrorCode, error) {
	ev := UnauthorizedErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnauthorizedErrorCode: valid values are %v", v, AllowedUnauthorizedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnauthorizedErrorCode) IsValid() bool {
	for _, existing := range AllowedUnauthorizedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnauthorizedErrorCode value
func (v UnauthorizedErrorCode) Ptr() *UnauthorizedErrorCode {
	return &v
}

type NullableUnauthorizedErrorCode struct {
	value *UnauthorizedErrorCode
	isSet bool
}

func (v NullableUnauthorizedErrorCode) Get() *UnauthorizedErrorCode {
	return v.value
}

func (v *NullableUnauthorizedErrorCode) Set(val *UnauthorizedErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedErrorCode(val *UnauthorizedErrorCode) *NullableUnauthorizedErrorCode {
	return &NullableUnauthorizedErrorCode{value: val, isSet: true}
}

func (v NullableUnauthorizedErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnauthorizedErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
