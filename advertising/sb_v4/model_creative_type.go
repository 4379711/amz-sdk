package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CreativeType The creative type of SB ad.
type CreativeType string

// List of CreativeType
const (
	CREATIVETYPE_PRODUCT_COLLECTION CreativeType = "PRODUCT_COLLECTION"
	CREATIVETYPE_STORE_SPOTLIGHT    CreativeType = "STORE_SPOTLIGHT"
	CREATIVETYPE_VIDEO              CreativeType = "VIDEO"
	CREATIVETYPE_BRAND_VIDEO        CreativeType = "BRAND_VIDEO"
)

// All allowed values of CreativeType enum
var AllowedCreativeTypeEnumValues = []CreativeType{
	"PRODUCT_COLLECTION",
	"STORE_SPOTLIGHT",
	"VIDEO",
	"BRAND_VIDEO",
}

func (v *CreativeType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreativeType(value)
	for _, existing := range AllowedCreativeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreativeType", value)
}

// NewCreativeTypeFromValue returns a pointer to a valid CreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTypeFromValue(v string) (*CreativeType, error) {
	ev := CreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeType: valid values are %v", v, AllowedCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeType) IsValid() bool {
	for _, existing := range AllowedCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreativeType value
func (v CreativeType) Ptr() *CreativeType {
	return &v
}

type NullableCreativeType struct {
	value *CreativeType
	isSet bool
}

func (v NullableCreativeType) Get() *CreativeType {
	return v.value
}

func (v *NullableCreativeType) Set(val *CreativeType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeType(val *CreativeType) *NullableCreativeType {
	return &NullableCreativeType{value: val, isSet: true}
}

func (v NullableCreativeType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
