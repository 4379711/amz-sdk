package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDRecommendationTypeV33 Signifies a type of recommendation. PRODUCT, CATEGORY, and CONTENT_CATEGORY are supported by tactic T00020. CATEGORY, AUDIENCE, and CONTENT_CATEGORY are supported by tactic T00030.
type SDRecommendationTypeV33 string

// List of SDRecommendationTypeV33
const (
	SDRECOMMENDATIONTYPEV33_PRODUCT          SDRecommendationTypeV33 = "PRODUCT"
	SDRECOMMENDATIONTYPEV33_CATEGORY         SDRecommendationTypeV33 = "CATEGORY"
	SDRECOMMENDATIONTYPEV33_AUDIENCE         SDRecommendationTypeV33 = "AUDIENCE"
	SDRECOMMENDATIONTYPEV33_CONTENT_CATEGORY SDRecommendationTypeV33 = "CONTENT_CATEGORY"
)

// All allowed values of SDRecommendationTypeV33 enum
var AllowedSDRecommendationTypeV33EnumValues = []SDRecommendationTypeV33{
	"PRODUCT",
	"CATEGORY",
	"AUDIENCE",
	"CONTENT_CATEGORY",
}

func (v *SDRecommendationTypeV33) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDRecommendationTypeV33(value)
	for _, existing := range AllowedSDRecommendationTypeV33EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDRecommendationTypeV33", value)
}

// NewSDRecommendationTypeV33FromValue returns a pointer to a valid SDRecommendationTypeV33
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDRecommendationTypeV33FromValue(v string) (*SDRecommendationTypeV33, error) {
	ev := SDRecommendationTypeV33(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDRecommendationTypeV33: valid values are %v", v, AllowedSDRecommendationTypeV33EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDRecommendationTypeV33) IsValid() bool {
	for _, existing := range AllowedSDRecommendationTypeV33EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDRecommendationTypeV33 value
func (v SDRecommendationTypeV33) Ptr() *SDRecommendationTypeV33 {
	return &v
}

type NullableSDRecommendationTypeV33 struct {
	value *SDRecommendationTypeV33
	isSet bool
}

func (v NullableSDRecommendationTypeV33) Get() *SDRecommendationTypeV33 {
	return v.value
}

func (v *NullableSDRecommendationTypeV33) Set(val *SDRecommendationTypeV33) {
	v.value = val
	v.isSet = true
}

func (v NullableSDRecommendationTypeV33) IsSet() bool {
	return v.isSet
}

func (v *NullableSDRecommendationTypeV33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDRecommendationTypeV33(val *SDRecommendationTypeV33) *NullableSDRecommendationTypeV33 {
	return &NullableSDRecommendationTypeV33{value: val, isSet: true}
}

func (v NullableSDRecommendationTypeV33) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDRecommendationTypeV33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
