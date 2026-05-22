package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ThrottledErrorCode the model 'ThrottledErrorCode'
type ThrottledErrorCode string

// List of ThrottledErrorCode
const (
	THROTTLEDERRORCODE_THROTTLED ThrottledErrorCode = "THROTTLED"
)

// All allowed values of ThrottledErrorCode enum
var AllowedThrottledErrorCodeEnumValues = []ThrottledErrorCode{
	"THROTTLED",
}

func (v *ThrottledErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThrottledErrorCode(value)
	for _, existing := range AllowedThrottledErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThrottledErrorCode", value)
}

// NewThrottledErrorCodeFromValue returns a pointer to a valid ThrottledErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThrottledErrorCodeFromValue(v string) (*ThrottledErrorCode, error) {
	ev := ThrottledErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ThrottledErrorCode: valid values are %v", v, AllowedThrottledErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThrottledErrorCode) IsValid() bool {
	for _, existing := range AllowedThrottledErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThrottledErrorCode value
func (v ThrottledErrorCode) Ptr() *ThrottledErrorCode {
	return &v
}

type NullableThrottledErrorCode struct {
	value *ThrottledErrorCode
	isSet bool
}

func (v NullableThrottledErrorCode) Get() *ThrottledErrorCode {
	return v.value
}

func (v *NullableThrottledErrorCode) Set(val *ThrottledErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottledErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottledErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottledErrorCode(val *ThrottledErrorCode) *NullableThrottledErrorCode {
	return &NullableThrottledErrorCode{value: val, isSet: true}
}

func (v NullableThrottledErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThrottledErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
