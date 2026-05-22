package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// InternalErrorErrorCode the model 'InternalErrorErrorCode'
type InternalErrorErrorCode string

// List of InternalErrorErrorCode
const (
	INTERNALERRORERRORCODE_INTERNAL_ERROR InternalErrorErrorCode = "INTERNAL_ERROR"
)

// All allowed values of InternalErrorErrorCode enum
var AllowedInternalErrorErrorCodeEnumValues = []InternalErrorErrorCode{
	"INTERNAL_ERROR",
}

func (v *InternalErrorErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalErrorErrorCode(value)
	for _, existing := range AllowedInternalErrorErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalErrorErrorCode", value)
}

// NewInternalErrorErrorCodeFromValue returns a pointer to a valid InternalErrorErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalErrorErrorCodeFromValue(v string) (*InternalErrorErrorCode, error) {
	ev := InternalErrorErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalErrorErrorCode: valid values are %v", v, AllowedInternalErrorErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalErrorErrorCode) IsValid() bool {
	for _, existing := range AllowedInternalErrorErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InternalErrorErrorCode value
func (v InternalErrorErrorCode) Ptr() *InternalErrorErrorCode {
	return &v
}

type NullableInternalErrorErrorCode struct {
	value *InternalErrorErrorCode
	isSet bool
}

func (v NullableInternalErrorErrorCode) Get() *InternalErrorErrorCode {
	return v.value
}

func (v *NullableInternalErrorErrorCode) Set(val *InternalErrorErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalErrorErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalErrorErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalErrorErrorCode(val *InternalErrorErrorCode) *NullableInternalErrorErrorCode {
	return &NullableInternalErrorErrorCode{value: val, isSet: true}
}

func (v NullableInternalErrorErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInternalErrorErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
