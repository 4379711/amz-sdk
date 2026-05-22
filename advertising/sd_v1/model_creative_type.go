package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CreativeType The type of the associated creative. If the field is empty or null, a default value of IMAGE will be used. One ad group only supports one type (VIDEO or IMAGE) of creativeType at a time. |Name|Description| |----|-----------| |IMAGE |The creative will display static assets (e.g. headline, brandLogo or custom image).| |VIDEO |The creative will display video assets. This type of creative must have a video asset provided. Only supported when using productAds with ASIN or SKU.|
type CreativeType string

// List of CreativeType
const (
	CREATIVETYPE_IMAGE CreativeType = "IMAGE"
	CREATIVETYPE_VIDEO CreativeType = "VIDEO"
)

// All allowed values of CreativeType enum
var AllowedCreativeTypeEnumValues = []CreativeType{
	"IMAGE",
	"VIDEO",
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
