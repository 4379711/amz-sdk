package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// LandingPageInvalidArgumentErrorCode the model 'LandingPageInvalidArgumentErrorCode'
type LandingPageInvalidArgumentErrorCode string

// List of LandingPageInvalidArgumentErrorCode
const (
	LANDINGPAGEINVALIDARGUMENTERRORCODE_INVALID_ARGUMENT LandingPageInvalidArgumentErrorCode = "INVALID_ARGUMENT"
)

// All allowed values of LandingPageInvalidArgumentErrorCode enum
var AllowedLandingPageInvalidArgumentErrorCodeEnumValues = []LandingPageInvalidArgumentErrorCode{
	"INVALID_ARGUMENT",
}

func (v *LandingPageInvalidArgumentErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LandingPageInvalidArgumentErrorCode(value)
	for _, existing := range AllowedLandingPageInvalidArgumentErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LandingPageInvalidArgumentErrorCode", value)
}

// NewLandingPageInvalidArgumentErrorCodeFromValue returns a pointer to a valid LandingPageInvalidArgumentErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLandingPageInvalidArgumentErrorCodeFromValue(v string) (*LandingPageInvalidArgumentErrorCode, error) {
	ev := LandingPageInvalidArgumentErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LandingPageInvalidArgumentErrorCode: valid values are %v", v, AllowedLandingPageInvalidArgumentErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LandingPageInvalidArgumentErrorCode) IsValid() bool {
	for _, existing := range AllowedLandingPageInvalidArgumentErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LandingPageInvalidArgumentErrorCode value
func (v LandingPageInvalidArgumentErrorCode) Ptr() *LandingPageInvalidArgumentErrorCode {
	return &v
}

type NullableLandingPageInvalidArgumentErrorCode struct {
	value *LandingPageInvalidArgumentErrorCode
	isSet bool
}

func (v NullableLandingPageInvalidArgumentErrorCode) Get() *LandingPageInvalidArgumentErrorCode {
	return v.value
}

func (v *NullableLandingPageInvalidArgumentErrorCode) Set(val *LandingPageInvalidArgumentErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageInvalidArgumentErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageInvalidArgumentErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageInvalidArgumentErrorCode(val *LandingPageInvalidArgumentErrorCode) *NullableLandingPageInvalidArgumentErrorCode {
	return &NullableLandingPageInvalidArgumentErrorCode{value: val, isSet: true}
}

func (v NullableLandingPageInvalidArgumentErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPageInvalidArgumentErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
