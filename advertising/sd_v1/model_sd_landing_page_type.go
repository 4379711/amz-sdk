package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDLandingPageType The type of the landingPage used. This field is not supported when using asin field.
type SDLandingPageType string

// List of SDLandingPageType
const (
	SDLANDINGPAGETYPE_OFF_AMAZON_LINK SDLandingPageType = "OFF_AMAZON_LINK"
)

// All allowed values of SDLandingPageType enum
var AllowedSDLandingPageTypeEnumValues = []SDLandingPageType{
	"OFF_AMAZON_LINK",
}

func (v *SDLandingPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDLandingPageType(value)
	for _, existing := range AllowedSDLandingPageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDLandingPageType", value)
}

// NewSDLandingPageTypeFromValue returns a pointer to a valid SDLandingPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDLandingPageTypeFromValue(v string) (*SDLandingPageType, error) {
	ev := SDLandingPageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDLandingPageType: valid values are %v", v, AllowedSDLandingPageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDLandingPageType) IsValid() bool {
	for _, existing := range AllowedSDLandingPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDLandingPageType value
func (v SDLandingPageType) Ptr() *SDLandingPageType {
	return &v
}

type NullableSDLandingPageType struct {
	value *SDLandingPageType
	isSet bool
}

func (v NullableSDLandingPageType) Get() *SDLandingPageType {
	return v.value
}

func (v *NullableSDLandingPageType) Set(val *SDLandingPageType) {
	v.value = val
	v.isSet = true
}

func (v NullableSDLandingPageType) IsSet() bool {
	return v.isSet
}

func (v *NullableSDLandingPageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDLandingPageType(val *SDLandingPageType) *NullableSDLandingPageType {
	return &NullableSDLandingPageType{value: val, isSet: true}
}

func (v NullableSDLandingPageType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDLandingPageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
