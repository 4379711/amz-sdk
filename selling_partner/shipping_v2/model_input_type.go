package shipping_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// InputType Type of Input.
type InputType string

// List of InputType
const (
	INPUTTYPE_TEXTBOX  InputType = "TEXTBOX"
	INPUTTYPE_PASSWORD InputType = "PASSWORD"
)

// All allowed values of InputType enum
var AllowedInputTypeEnumValues = []InputType{
	INPUTTYPE_TEXTBOX,
	INPUTTYPE_PASSWORD,
}

func (v *InputType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InputType(value)
	for _, existing := range AllowedInputTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InputType", value)
}

// NewInputTypeFromValue returns a pointer to a valid InputType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInputTypeFromValue(v string) (*InputType, error) {
	ev := InputType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InputType: valid values are %v", v, AllowedInputTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InputType) IsValid() bool {
	for _, existing := range AllowedInputTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InputType value
func (v InputType) Ptr() *InputType {
	return &v
}

type NullableInputType struct {
	value *InputType
	isSet bool
}

func (v NullableInputType) Get() *InputType {
	return v.value
}

func (v *NullableInputType) Set(val *InputType) {
	v.value = val
	v.isSet = true
}

func (v NullableInputType) IsSet() bool {
	return v.isSet
}

func (v *NullableInputType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputType(val *InputType) *NullableInputType {
	return &NullableInputType{value: val, isSet: true}
}

func (v NullableInputType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInputType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
