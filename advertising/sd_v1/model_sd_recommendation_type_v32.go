package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SDRecommendationTypeV32 Signifies a type of recommendation. PRODUCT and CATEGORY are supported by tactic T00020. CATEGORY and AUDIENCE are supported by tactic T00030.
type SDRecommendationTypeV32 string

// List of SDRecommendationTypeV32
const (
	SDRECOMMENDATIONTYPEV32_PRODUCT  SDRecommendationTypeV32 = "PRODUCT"
	SDRECOMMENDATIONTYPEV32_CATEGORY SDRecommendationTypeV32 = "CATEGORY"
	SDRECOMMENDATIONTYPEV32_AUDIENCE SDRecommendationTypeV32 = "AUDIENCE"
)

// All allowed values of SDRecommendationTypeV32 enum
var AllowedSDRecommendationTypeV32EnumValues = []SDRecommendationTypeV32{
	"PRODUCT",
	"CATEGORY",
	"AUDIENCE",
}

func (v *SDRecommendationTypeV32) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDRecommendationTypeV32(value)
	for _, existing := range AllowedSDRecommendationTypeV32EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDRecommendationTypeV32", value)
}

// NewSDRecommendationTypeV32FromValue returns a pointer to a valid SDRecommendationTypeV32
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDRecommendationTypeV32FromValue(v string) (*SDRecommendationTypeV32, error) {
	ev := SDRecommendationTypeV32(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDRecommendationTypeV32: valid values are %v", v, AllowedSDRecommendationTypeV32EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDRecommendationTypeV32) IsValid() bool {
	for _, existing := range AllowedSDRecommendationTypeV32EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDRecommendationTypeV32 value
func (v SDRecommendationTypeV32) Ptr() *SDRecommendationTypeV32 {
	return &v
}

type NullableSDRecommendationTypeV32 struct {
	value *SDRecommendationTypeV32
	isSet bool
}

func (v NullableSDRecommendationTypeV32) Get() *SDRecommendationTypeV32 {
	return v.value
}

func (v *NullableSDRecommendationTypeV32) Set(val *SDRecommendationTypeV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDRecommendationTypeV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDRecommendationTypeV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDRecommendationTypeV32(val *SDRecommendationTypeV32) *NullableSDRecommendationTypeV32 {
	return &NullableSDRecommendationTypeV32{value: val, isSet: true}
}

func (v NullableSDRecommendationTypeV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDRecommendationTypeV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
