package aplus_content_20201101

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ColorType The relative color scheme of your content.
type ColorType string

// List of ColorType
const (
	COLORTYPE_DARK  ColorType = "DARK"
	COLORTYPE_LIGHT ColorType = "LIGHT"
)

// All allowed values of ColorType enum
var AllowedColorTypeEnumValues = []ColorType{
	COLORTYPE_DARK,
	COLORTYPE_LIGHT,
}

func (v *ColorType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ColorType(value)
	for _, existing := range AllowedColorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ColorType", value)
}

// NewColorTypeFromValue returns a pointer to a valid ColorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewColorTypeFromValue(v string) (*ColorType, error) {
	ev := ColorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ColorType: valid values are %v", v, AllowedColorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ColorType) IsValid() bool {
	for _, existing := range AllowedColorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ColorType value
func (v ColorType) Ptr() *ColorType {
	return &v
}

type NullableColorType struct {
	value *ColorType
	isSet bool
}

func (v NullableColorType) Get() *ColorType {
	return v.value
}

func (v *NullableColorType) Set(val *ColorType) {
	v.value = val
	v.isSet = true
}

func (v NullableColorType) IsSet() bool {
	return v.isSet
}

func (v *NullableColorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColorType(val *ColorType) *NullableColorType {
	return &NullableColorType{value: val, isSet: true}
}

func (v NullableColorType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableColorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
