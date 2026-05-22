package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDCreativeType The type of the associated creative. If the field is empty or null, a default value of IMAGE will be used. Only supports one type (VIDEO or IMAGE) at a time.
type SDCreativeType string

// List of SDCreativeType
const (
	SDCREATIVETYPE_IMAGE SDCreativeType = "IMAGE"
	SDCREATIVETYPE_VIDEO SDCreativeType = "VIDEO"
)

// All allowed values of SDCreativeType enum
var AllowedSDCreativeTypeEnumValues = []SDCreativeType{
	"IMAGE",
	"VIDEO",
}

func (v *SDCreativeType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDCreativeType(value)
	for _, existing := range AllowedSDCreativeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDCreativeType", value)
}

// NewSDCreativeTypeFromValue returns a pointer to a valid SDCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDCreativeTypeFromValue(v string) (*SDCreativeType, error) {
	ev := SDCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDCreativeType: valid values are %v", v, AllowedSDCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDCreativeType) IsValid() bool {
	for _, existing := range AllowedSDCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDCreativeType value
func (v SDCreativeType) Ptr() *SDCreativeType {
	return &v
}

type NullableSDCreativeType struct {
	value *SDCreativeType
	isSet bool
}

func (v NullableSDCreativeType) Get() *SDCreativeType {
	return v.value
}

func (v *NullableSDCreativeType) Set(val *SDCreativeType) {
	v.value = val
	v.isSet = true
}

func (v NullableSDCreativeType) IsSet() bool {
	return v.isSet
}

func (v *NullableSDCreativeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDCreativeType(val *SDCreativeType) *NullableSDCreativeType {
	return &NullableSDCreativeType{value: val, isSet: true}
}

func (v NullableSDCreativeType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDCreativeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
