package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ThrottlingErrorCode the model 'ThrottlingErrorCode'
type ThrottlingErrorCode string

// List of ThrottlingErrorCode
const (
	THROTTLINGERRORCODE_THROTTLED ThrottlingErrorCode = "THROTTLED"
)

// All allowed values of ThrottlingErrorCode enum
var AllowedThrottlingErrorCodeEnumValues = []ThrottlingErrorCode{
	"THROTTLED",
}

func (v *ThrottlingErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThrottlingErrorCode(value)
	for _, existing := range AllowedThrottlingErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThrottlingErrorCode", value)
}

// NewThrottlingErrorCodeFromValue returns a pointer to a valid ThrottlingErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThrottlingErrorCodeFromValue(v string) (*ThrottlingErrorCode, error) {
	ev := ThrottlingErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ThrottlingErrorCode: valid values are %v", v, AllowedThrottlingErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThrottlingErrorCode) IsValid() bool {
	for _, existing := range AllowedThrottlingErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThrottlingErrorCode value
func (v ThrottlingErrorCode) Ptr() *ThrottlingErrorCode {
	return &v
}

type NullableThrottlingErrorCode struct {
	value *ThrottlingErrorCode
	isSet bool
}

func (v NullableThrottlingErrorCode) Get() *ThrottlingErrorCode {
	return v.value
}

func (v *NullableThrottlingErrorCode) Set(val *ThrottlingErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingErrorCode(val *ThrottlingErrorCode) *NullableThrottlingErrorCode {
	return &NullableThrottlingErrorCode{value: val, isSet: true}
}

func (v NullableThrottlingErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThrottlingErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
