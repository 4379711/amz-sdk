package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SDRecommendationType Signifies a type of recommendation
type SDRecommendationType string

// List of SDRecommendationType
const (
	SDRECOMMENDATIONTYPE_PRODUCT SDRecommendationType = "PRODUCT"
)

// All allowed values of SDRecommendationType enum
var AllowedSDRecommendationTypeEnumValues = []SDRecommendationType{
	"PRODUCT",
}

func (v *SDRecommendationType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDRecommendationType(value)
	for _, existing := range AllowedSDRecommendationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDRecommendationType", value)
}

// NewSDRecommendationTypeFromValue returns a pointer to a valid SDRecommendationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDRecommendationTypeFromValue(v string) (*SDRecommendationType, error) {
	ev := SDRecommendationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDRecommendationType: valid values are %v", v, AllowedSDRecommendationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDRecommendationType) IsValid() bool {
	for _, existing := range AllowedSDRecommendationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDRecommendationType value
func (v SDRecommendationType) Ptr() *SDRecommendationType {
	return &v
}

type NullableSDRecommendationType struct {
	value *SDRecommendationType
	isSet bool
}

func (v NullableSDRecommendationType) Get() *SDRecommendationType {
	return v.value
}

func (v *NullableSDRecommendationType) Set(val *SDRecommendationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSDRecommendationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSDRecommendationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDRecommendationType(val *SDRecommendationType) *NullableSDRecommendationType {
	return &NullableSDRecommendationType{value: val, isSet: true}
}

func (v NullableSDRecommendationType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDRecommendationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
