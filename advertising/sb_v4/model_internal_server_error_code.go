package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// InternalServerErrorCode the model 'InternalServerErrorCode'
type InternalServerErrorCode string

// List of InternalServerErrorCode
const (
	INTERNALSERVERERRORCODE_INTERNAL_ERROR InternalServerErrorCode = "INTERNAL_ERROR"
)

// All allowed values of InternalServerErrorCode enum
var AllowedInternalServerErrorCodeEnumValues = []InternalServerErrorCode{
	"INTERNAL_ERROR",
}

func (v *InternalServerErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalServerErrorCode(value)
	for _, existing := range AllowedInternalServerErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalServerErrorCode", value)
}

// NewInternalServerErrorCodeFromValue returns a pointer to a valid InternalServerErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalServerErrorCodeFromValue(v string) (*InternalServerErrorCode, error) {
	ev := InternalServerErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalServerErrorCode: valid values are %v", v, AllowedInternalServerErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalServerErrorCode) IsValid() bool {
	for _, existing := range AllowedInternalServerErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InternalServerErrorCode value
func (v InternalServerErrorCode) Ptr() *InternalServerErrorCode {
	return &v
}

type NullableInternalServerErrorCode struct {
	value *InternalServerErrorCode
	isSet bool
}

func (v NullableInternalServerErrorCode) Get() *InternalServerErrorCode {
	return v.value
}

func (v *NullableInternalServerErrorCode) Set(val *InternalServerErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerErrorCode(val *InternalServerErrorCode) *NullableInternalServerErrorCode {
	return &NullableInternalServerErrorCode{value: val, isSet: true}
}

func (v NullableInternalServerErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInternalServerErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
