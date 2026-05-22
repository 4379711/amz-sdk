package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDRecommendationTypeV31 Signifies a type of recommendation
type SDRecommendationTypeV31 string

// List of SDRecommendationTypeV31
const (
	SDRECOMMENDATIONTYPEV31_PRODUCT  SDRecommendationTypeV31 = "PRODUCT"
	SDRECOMMENDATIONTYPEV31_CATEGORY SDRecommendationTypeV31 = "CATEGORY"
)

// All allowed values of SDRecommendationTypeV31 enum
var AllowedSDRecommendationTypeV31EnumValues = []SDRecommendationTypeV31{
	"PRODUCT",
	"CATEGORY",
}

func (v *SDRecommendationTypeV31) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDRecommendationTypeV31(value)
	for _, existing := range AllowedSDRecommendationTypeV31EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDRecommendationTypeV31", value)
}

// NewSDRecommendationTypeV31FromValue returns a pointer to a valid SDRecommendationTypeV31
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDRecommendationTypeV31FromValue(v string) (*SDRecommendationTypeV31, error) {
	ev := SDRecommendationTypeV31(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDRecommendationTypeV31: valid values are %v", v, AllowedSDRecommendationTypeV31EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDRecommendationTypeV31) IsValid() bool {
	for _, existing := range AllowedSDRecommendationTypeV31EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDRecommendationTypeV31 value
func (v SDRecommendationTypeV31) Ptr() *SDRecommendationTypeV31 {
	return &v
}

type NullableSDRecommendationTypeV31 struct {
	value *SDRecommendationTypeV31
	isSet bool
}

func (v NullableSDRecommendationTypeV31) Get() *SDRecommendationTypeV31 {
	return v.value
}

func (v *NullableSDRecommendationTypeV31) Set(val *SDRecommendationTypeV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDRecommendationTypeV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDRecommendationTypeV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDRecommendationTypeV31(val *SDRecommendationTypeV31) *NullableSDRecommendationTypeV31 {
	return &NullableSDRecommendationTypeV31{value: val, isSet: true}
}

func (v NullableSDRecommendationTypeV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDRecommendationTypeV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
