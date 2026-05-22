package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// LandingPageInternalErrorCode the model 'LandingPageInternalErrorCode'
type LandingPageInternalErrorCode string

// List of LandingPageInternalErrorCode
const (
	LANDINGPAGEINTERNALERRORCODE_INTERNAL_ERROR LandingPageInternalErrorCode = "INTERNAL_ERROR"
)

// All allowed values of LandingPageInternalErrorCode enum
var AllowedLandingPageInternalErrorCodeEnumValues = []LandingPageInternalErrorCode{
	"INTERNAL_ERROR",
}

func (v *LandingPageInternalErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LandingPageInternalErrorCode(value)
	for _, existing := range AllowedLandingPageInternalErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LandingPageInternalErrorCode", value)
}

// NewLandingPageInternalErrorCodeFromValue returns a pointer to a valid LandingPageInternalErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLandingPageInternalErrorCodeFromValue(v string) (*LandingPageInternalErrorCode, error) {
	ev := LandingPageInternalErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LandingPageInternalErrorCode: valid values are %v", v, AllowedLandingPageInternalErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LandingPageInternalErrorCode) IsValid() bool {
	for _, existing := range AllowedLandingPageInternalErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LandingPageInternalErrorCode value
func (v LandingPageInternalErrorCode) Ptr() *LandingPageInternalErrorCode {
	return &v
}

type NullableLandingPageInternalErrorCode struct {
	value *LandingPageInternalErrorCode
	isSet bool
}

func (v NullableLandingPageInternalErrorCode) Get() *LandingPageInternalErrorCode {
	return v.value
}

func (v *NullableLandingPageInternalErrorCode) Set(val *LandingPageInternalErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageInternalErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageInternalErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageInternalErrorCode(val *LandingPageInternalErrorCode) *NullableLandingPageInternalErrorCode {
	return &NullableLandingPageInternalErrorCode{value: val, isSet: true}
}

func (v NullableLandingPageInternalErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLandingPageInternalErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
