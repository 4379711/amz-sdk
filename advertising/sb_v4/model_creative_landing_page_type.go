package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CreativeLandingPageType Landing page type
type CreativeLandingPageType string

// List of CreativeLandingPageType
const (
	CREATIVELANDINGPAGETYPE_PRODUCT_LIST             CreativeLandingPageType = "PRODUCT_LIST"
	CREATIVELANDINGPAGETYPE_STORE                    CreativeLandingPageType = "STORE"
	CREATIVELANDINGPAGETYPE_DETAIL_PAGE              CreativeLandingPageType = "DETAIL_PAGE"
	CREATIVELANDINGPAGETYPE_CUSTOM_URL               CreativeLandingPageType = "CUSTOM_URL"
	CREATIVELANDINGPAGETYPE_AD_LANDING_PREVIEW       CreativeLandingPageType = "AD_LANDING_PREVIEW"
	CREATIVELANDINGPAGETYPE_SEARCH                   CreativeLandingPageType = "SEARCH"
	CREATIVELANDINGPAGETYPE_BROWSE                   CreativeLandingPageType = "BROWSE"
	CREATIVELANDINGPAGETYPE_ADVERTISING_LANDING_PAGE CreativeLandingPageType = "ADVERTISING_LANDING_PAGE"
	CREATIVELANDINGPAGETYPE_UNKNOWN                  CreativeLandingPageType = "UNKNOWN"
)

// All allowed values of CreativeLandingPageType enum
var AllowedCreativeLandingPageTypeEnumValues = []CreativeLandingPageType{
	"PRODUCT_LIST",
	"STORE",
	"DETAIL_PAGE",
	"CUSTOM_URL",
	"AD_LANDING_PREVIEW",
	"SEARCH",
	"BROWSE",
	"ADVERTISING_LANDING_PAGE",
	"UNKNOWN",
}

func (v *CreativeLandingPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreativeLandingPageType(value)
	for _, existing := range AllowedCreativeLandingPageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreativeLandingPageType", value)
}

// NewCreativeLandingPageTypeFromValue returns a pointer to a valid CreativeLandingPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeLandingPageTypeFromValue(v string) (*CreativeLandingPageType, error) {
	ev := CreativeLandingPageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeLandingPageType: valid values are %v", v, AllowedCreativeLandingPageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeLandingPageType) IsValid() bool {
	for _, existing := range AllowedCreativeLandingPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreativeLandingPageType value
func (v CreativeLandingPageType) Ptr() *CreativeLandingPageType {
	return &v
}

type NullableCreativeLandingPageType struct {
	value *CreativeLandingPageType
	isSet bool
}

func (v NullableCreativeLandingPageType) Get() *CreativeLandingPageType {
	return v.value
}

func (v *NullableCreativeLandingPageType) Set(val *CreativeLandingPageType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeLandingPageType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeLandingPageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeLandingPageType(val *CreativeLandingPageType) *NullableCreativeLandingPageType {
	return &NullableCreativeLandingPageType{value: val, isSet: true}
}

func (v NullableCreativeLandingPageType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeLandingPageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
