package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// InvalidArgumentErrorCode the model 'InvalidArgumentErrorCode'
type InvalidArgumentErrorCode string

// List of InvalidArgumentErrorCode
const (
	INVALIDARGUMENTERRORCODE_INVALID_ARGUMENT InvalidArgumentErrorCode = "INVALID_ARGUMENT"
)

// All allowed values of InvalidArgumentErrorCode enum
var AllowedInvalidArgumentErrorCodeEnumValues = []InvalidArgumentErrorCode{
	"INVALID_ARGUMENT",
}

func (v *InvalidArgumentErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvalidArgumentErrorCode(value)
	for _, existing := range AllowedInvalidArgumentErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvalidArgumentErrorCode", value)
}

// NewInvalidArgumentErrorCodeFromValue returns a pointer to a valid InvalidArgumentErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvalidArgumentErrorCodeFromValue(v string) (*InvalidArgumentErrorCode, error) {
	ev := InvalidArgumentErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvalidArgumentErrorCode: valid values are %v", v, AllowedInvalidArgumentErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvalidArgumentErrorCode) IsValid() bool {
	for _, existing := range AllowedInvalidArgumentErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvalidArgumentErrorCode value
func (v InvalidArgumentErrorCode) Ptr() *InvalidArgumentErrorCode {
	return &v
}

type NullableInvalidArgumentErrorCode struct {
	value *InvalidArgumentErrorCode
	isSet bool
}

func (v NullableInvalidArgumentErrorCode) Get() *InvalidArgumentErrorCode {
	return v.value
}

func (v *NullableInvalidArgumentErrorCode) Set(val *InvalidArgumentErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidArgumentErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidArgumentErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidArgumentErrorCode(val *InvalidArgumentErrorCode) *NullableInvalidArgumentErrorCode {
	return &NullableInvalidArgumentErrorCode{value: val, isSet: true}
}

func (v NullableInvalidArgumentErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidArgumentErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
