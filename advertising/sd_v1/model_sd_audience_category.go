package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDAudienceCategory An audience category determines the goal of the audience such as In-market, Interest, Lifestyle and Life Event
type SDAudienceCategory string

// List of SDAudienceCategory
const (
	SDAUDIENCECATEGORY_IN_MARKET  SDAudienceCategory = "In-market"
	SDAUDIENCECATEGORY_LIFESTYLE  SDAudienceCategory = "Lifestyle"
	SDAUDIENCECATEGORY_INTEREST   SDAudienceCategory = "Interest"
	SDAUDIENCECATEGORY_LIFE_EVENT SDAudienceCategory = "Life event"
)

// All allowed values of SDAudienceCategory enum
var AllowedSDAudienceCategoryEnumValues = []SDAudienceCategory{
	"In-market",
	"Lifestyle",
	"Interest",
	"Life event",
}

func (v *SDAudienceCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDAudienceCategory(value)
	for _, existing := range AllowedSDAudienceCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDAudienceCategory", value)
}

// NewSDAudienceCategoryFromValue returns a pointer to a valid SDAudienceCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDAudienceCategoryFromValue(v string) (*SDAudienceCategory, error) {
	ev := SDAudienceCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDAudienceCategory: valid values are %v", v, AllowedSDAudienceCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDAudienceCategory) IsValid() bool {
	for _, existing := range AllowedSDAudienceCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDAudienceCategory value
func (v SDAudienceCategory) Ptr() *SDAudienceCategory {
	return &v
}

type NullableSDAudienceCategory struct {
	value *SDAudienceCategory
	isSet bool
}

func (v NullableSDAudienceCategory) Get() *SDAudienceCategory {
	return v.value
}

func (v *NullableSDAudienceCategory) Set(val *SDAudienceCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableSDAudienceCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableSDAudienceCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDAudienceCategory(val *SDAudienceCategory) *NullableSDAudienceCategory {
	return &NullableSDAudienceCategory{value: val, isSet: true}
}

func (v NullableSDAudienceCategory) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDAudienceCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
