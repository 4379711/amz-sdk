package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsInternalServerExceptionCode the model 'SponsoredProductsInternalServerExceptionCode'
type SponsoredProductsInternalServerExceptionCode string

// List of SponsoredProductsInternalServerExceptionCode
const (
	SPONSOREDPRODUCTSINTERNALSERVEREXCEPTIONCODE_INTERNAL_SERVER_EXCEPTION SponsoredProductsInternalServerExceptionCode = "internalServerException"
)

// All allowed values of SponsoredProductsInternalServerExceptionCode enum
var AllowedSponsoredProductsInternalServerExceptionCodeEnumValues = []SponsoredProductsInternalServerExceptionCode{
	"internalServerException",
}

func (v *SponsoredProductsInternalServerExceptionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsInternalServerExceptionCode(value)
	for _, existing := range AllowedSponsoredProductsInternalServerExceptionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsInternalServerExceptionCode", value)
}

// NewSponsoredProductsInternalServerExceptionCodeFromValue returns a pointer to a valid SponsoredProductsInternalServerExceptionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsInternalServerExceptionCodeFromValue(v string) (*SponsoredProductsInternalServerExceptionCode, error) {
	ev := SponsoredProductsInternalServerExceptionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsInternalServerExceptionCode: valid values are %v", v, AllowedSponsoredProductsInternalServerExceptionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsInternalServerExceptionCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsInternalServerExceptionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsInternalServerExceptionCode value
func (v SponsoredProductsInternalServerExceptionCode) Ptr() *SponsoredProductsInternalServerExceptionCode {
	return &v
}

type NullableSponsoredProductsInternalServerExceptionCode struct {
	value *SponsoredProductsInternalServerExceptionCode
	isSet bool
}

func (v NullableSponsoredProductsInternalServerExceptionCode) Get() *SponsoredProductsInternalServerExceptionCode {
	return v.value
}

func (v *NullableSponsoredProductsInternalServerExceptionCode) Set(val *SponsoredProductsInternalServerExceptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInternalServerExceptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInternalServerExceptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInternalServerExceptionCode(val *SponsoredProductsInternalServerExceptionCode) *NullableSponsoredProductsInternalServerExceptionCode {
	return &NullableSponsoredProductsInternalServerExceptionCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsInternalServerExceptionCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInternalServerExceptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
