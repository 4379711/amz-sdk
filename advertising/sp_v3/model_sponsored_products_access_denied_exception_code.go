package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsAccessDeniedExceptionCode the model 'SponsoredProductsAccessDeniedExceptionCode'
type SponsoredProductsAccessDeniedExceptionCode string

// List of SponsoredProductsAccessDeniedExceptionCode
const (
	SPONSOREDPRODUCTSACCESSDENIEDEXCEPTIONCODE_ACCESS_DENIED SponsoredProductsAccessDeniedExceptionCode = "accessDenied"
)

// All allowed values of SponsoredProductsAccessDeniedExceptionCode enum
var AllowedSponsoredProductsAccessDeniedExceptionCodeEnumValues = []SponsoredProductsAccessDeniedExceptionCode{
	"accessDenied",
}

func (v *SponsoredProductsAccessDeniedExceptionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAccessDeniedExceptionCode(value)
	for _, existing := range AllowedSponsoredProductsAccessDeniedExceptionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAccessDeniedExceptionCode", value)
}

// NewSponsoredProductsAccessDeniedExceptionCodeFromValue returns a pointer to a valid SponsoredProductsAccessDeniedExceptionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAccessDeniedExceptionCodeFromValue(v string) (*SponsoredProductsAccessDeniedExceptionCode, error) {
	ev := SponsoredProductsAccessDeniedExceptionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAccessDeniedExceptionCode: valid values are %v", v, AllowedSponsoredProductsAccessDeniedExceptionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAccessDeniedExceptionCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAccessDeniedExceptionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAccessDeniedExceptionCode value
func (v SponsoredProductsAccessDeniedExceptionCode) Ptr() *SponsoredProductsAccessDeniedExceptionCode {
	return &v
}

type NullableSponsoredProductsAccessDeniedExceptionCode struct {
	value *SponsoredProductsAccessDeniedExceptionCode
	isSet bool
}

func (v NullableSponsoredProductsAccessDeniedExceptionCode) Get() *SponsoredProductsAccessDeniedExceptionCode {
	return v.value
}

func (v *NullableSponsoredProductsAccessDeniedExceptionCode) Set(val *SponsoredProductsAccessDeniedExceptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAccessDeniedExceptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAccessDeniedExceptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAccessDeniedExceptionCode(val *SponsoredProductsAccessDeniedExceptionCode) *NullableSponsoredProductsAccessDeniedExceptionCode {
	return &NullableSponsoredProductsAccessDeniedExceptionCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsAccessDeniedExceptionCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAccessDeniedExceptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
