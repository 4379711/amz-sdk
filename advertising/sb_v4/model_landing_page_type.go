package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// LandingPageType The type of landing page, such as store page, product list (simple landing page), custom url.
type LandingPageType string

// List of LandingPageType
const (
	LANDINGPAGETYPE_PRODUCT_LIST LandingPageType = "PRODUCT_LIST"
	LANDINGPAGETYPE_STORE        LandingPageType = "STORE"
	LANDINGPAGETYPE_CUSTOM_URL   LandingPageType = "CUSTOM_URL"
	LANDINGPAGETYPE_DETAIL_PAGE  LandingPageType = "DETAIL_PAGE"
)

// All allowed values of LandingPageType enum
var AllowedLandingPageTypeEnumValues = []LandingPageType{
	"PRODUCT_LIST",
	"STORE",
	"CUSTOM_URL",
	"DETAIL_PAGE",
}

func (v *LandingPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LandingPageType(value)
	for _, existing := range AllowedLandingPageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LandingPageType", value)
}

// NewLandingPageTypeFromValue returns a pointer to a valid LandingPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLandingPageTypeFromValue(v string) (*LandingPageType, error) {
	ev := LandingPageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LandingPageType: valid values are %v", v, AllowedLandingPageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LandingPageType) IsValid() bool {
	for _, existing := range AllowedLandingPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LandingPageType value
func (v LandingPageType) Ptr() *LandingPageType {
	return &v
}

type NullableLandingPageType struct {
	value *LandingPageType
	isSet bool
}

func (v NullableLandingPageType) Get() *LandingPageType {
	return v.value
}

func (v *NullableLandingPageType) Set(val *LandingPageType) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageType) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageType(val *LandingPageType) *NullableLandingPageType {
	return &NullableLandingPageType{value: val, isSet: true}
}

func (v NullableLandingPageType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
