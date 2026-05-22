package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// UnsupportedMediaTypeErrorCode the model 'UnsupportedMediaTypeErrorCode'
type UnsupportedMediaTypeErrorCode string

// List of UnsupportedMediaTypeErrorCode
const (
	UNSUPPORTEDMEDIATYPEERRORCODE_UNSUPPORTED_MEDIA_TYPE UnsupportedMediaTypeErrorCode = "UNSUPPORTED_MEDIA_TYPE"
)

// All allowed values of UnsupportedMediaTypeErrorCode enum
var AllowedUnsupportedMediaTypeErrorCodeEnumValues = []UnsupportedMediaTypeErrorCode{
	"UNSUPPORTED_MEDIA_TYPE",
}

func (v *UnsupportedMediaTypeErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnsupportedMediaTypeErrorCode(value)
	for _, existing := range AllowedUnsupportedMediaTypeErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnsupportedMediaTypeErrorCode", value)
}

// NewUnsupportedMediaTypeErrorCodeFromValue returns a pointer to a valid UnsupportedMediaTypeErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnsupportedMediaTypeErrorCodeFromValue(v string) (*UnsupportedMediaTypeErrorCode, error) {
	ev := UnsupportedMediaTypeErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnsupportedMediaTypeErrorCode: valid values are %v", v, AllowedUnsupportedMediaTypeErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnsupportedMediaTypeErrorCode) IsValid() bool {
	for _, existing := range AllowedUnsupportedMediaTypeErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnsupportedMediaTypeErrorCode value
func (v UnsupportedMediaTypeErrorCode) Ptr() *UnsupportedMediaTypeErrorCode {
	return &v
}

type NullableUnsupportedMediaTypeErrorCode struct {
	value *UnsupportedMediaTypeErrorCode
	isSet bool
}

func (v NullableUnsupportedMediaTypeErrorCode) Get() *UnsupportedMediaTypeErrorCode {
	return v.value
}

func (v *NullableUnsupportedMediaTypeErrorCode) Set(val *UnsupportedMediaTypeErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsupportedMediaTypeErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsupportedMediaTypeErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsupportedMediaTypeErrorCode(val *UnsupportedMediaTypeErrorCode) *NullableUnsupportedMediaTypeErrorCode {
	return &NullableUnsupportedMediaTypeErrorCode{value: val, isSet: true}
}

func (v NullableUnsupportedMediaTypeErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnsupportedMediaTypeErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
