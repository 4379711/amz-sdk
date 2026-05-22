package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SBKeywordRecommendationThemeType Theme type for targeting. Used to get keyword recommendations for theme.
type SBKeywordRecommendationThemeType string

// List of SBKeywordRecommendationThemeType
const (
	SBKEYWORDRECOMMENDATIONTHEMETYPE_BRAND         SBKeywordRecommendationThemeType = "KEYWORDS_RELATED_TO_YOUR_BRAND"
	SBKEYWORDRECOMMENDATIONTHEMETYPE_LANDING_PAGES SBKeywordRecommendationThemeType = "KEYWORDS_RELATED_TO_YOUR_LANDING_PAGES"
)

// All allowed values of SBKeywordRecommendationThemeType enum
var AllowedSBKeywordRecommendationThemeTypeEnumValues = []SBKeywordRecommendationThemeType{
	"KEYWORDS_RELATED_TO_YOUR_BRAND",
	"KEYWORDS_RELATED_TO_YOUR_LANDING_PAGES",
}

func (v *SBKeywordRecommendationThemeType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SBKeywordRecommendationThemeType(value)
	for _, existing := range AllowedSBKeywordRecommendationThemeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SBKeywordRecommendationThemeType", value)
}

// NewSBKeywordRecommendationThemeTypeFromValue returns a pointer to a valid SBKeywordRecommendationThemeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSBKeywordRecommendationThemeTypeFromValue(v string) (*SBKeywordRecommendationThemeType, error) {
	ev := SBKeywordRecommendationThemeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SBKeywordRecommendationThemeType: valid values are %v", v, AllowedSBKeywordRecommendationThemeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SBKeywordRecommendationThemeType) IsValid() bool {
	for _, existing := range AllowedSBKeywordRecommendationThemeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBKeywordRecommendationThemeType value
func (v SBKeywordRecommendationThemeType) Ptr() *SBKeywordRecommendationThemeType {
	return &v
}

type NullableSBKeywordRecommendationThemeType struct {
	value *SBKeywordRecommendationThemeType
	isSet bool
}

func (v NullableSBKeywordRecommendationThemeType) Get() *SBKeywordRecommendationThemeType {
	return v.value
}

func (v *NullableSBKeywordRecommendationThemeType) Set(val *SBKeywordRecommendationThemeType) {
	v.value = val
	v.isSet = true
}

func (v NullableSBKeywordRecommendationThemeType) IsSet() bool {
	return v.isSet
}

func (v *NullableSBKeywordRecommendationThemeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBKeywordRecommendationThemeType(val *SBKeywordRecommendationThemeType) *NullableSBKeywordRecommendationThemeType {
	return &NullableSBKeywordRecommendationThemeType{value: val, isSet: true}
}

func (v NullableSBKeywordRecommendationThemeType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBKeywordRecommendationThemeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
