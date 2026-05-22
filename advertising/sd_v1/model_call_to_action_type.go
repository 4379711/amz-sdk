package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CallToActionType The call to action type. |Call to action type|Description| |--------|-----------| |SIGN_UP|Call to action type for customers to sign up.| |INQUIRE|Call to action type for customers to inquire.| |REQUEST_QUOTE|Call to action type for customers to request a quote.|
type CallToActionType string

// List of callToActionType
const (
	CALLTOACTIONTYPE_SIGN_UP       CallToActionType = "SIGN_UP"
	CALLTOACTIONTYPE_INQUIRE       CallToActionType = "INQUIRE"
	CALLTOACTIONTYPE_REQUEST_QUOTE CallToActionType = "REQUEST_QUOTE"
)

// All allowed values of CallToActionType enum
var AllowedCallToActionTypeEnumValues = []CallToActionType{
	"SIGN_UP",
	"INQUIRE",
	"REQUEST_QUOTE",
}

func (v *CallToActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CallToActionType(value)
	for _, existing := range AllowedCallToActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CallToActionType", value)
}

// NewCallToActionTypeFromValue returns a pointer to a valid CallToActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCallToActionTypeFromValue(v string) (*CallToActionType, error) {
	ev := CallToActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CallToActionType: valid values are %v", v, AllowedCallToActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CallToActionType) IsValid() bool {
	for _, existing := range AllowedCallToActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to callToActionType value
func (v CallToActionType) Ptr() *CallToActionType {
	return &v
}

type NullableCallToActionType struct {
	value *CallToActionType
	isSet bool
}

func (v NullableCallToActionType) Get() *CallToActionType {
	return v.value
}

func (v *NullableCallToActionType) Set(val *CallToActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableCallToActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableCallToActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallToActionType(val *CallToActionType) *NullableCallToActionType {
	return &NullableCallToActionType{value: val, isSet: true}
}

func (v NullableCallToActionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCallToActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
