package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsSchemaValidationExceptionCode the model 'SponsoredProductsSchemaValidationExceptionCode'
type SponsoredProductsSchemaValidationExceptionCode string

// List of SponsoredProductsSchemaValidationExceptionCode
const (
	SPONSOREDPRODUCTSSCHEMAVALIDATIONEXCEPTIONCODE_INVALID_SCHEMA SponsoredProductsSchemaValidationExceptionCode = "invalidSchema"
)

// All allowed values of SponsoredProductsSchemaValidationExceptionCode enum
var AllowedSponsoredProductsSchemaValidationExceptionCodeEnumValues = []SponsoredProductsSchemaValidationExceptionCode{
	"invalidSchema",
}

func (v *SponsoredProductsSchemaValidationExceptionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsSchemaValidationExceptionCode(value)
	for _, existing := range AllowedSponsoredProductsSchemaValidationExceptionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsSchemaValidationExceptionCode", value)
}

// NewSponsoredProductsSchemaValidationExceptionCodeFromValue returns a pointer to a valid SponsoredProductsSchemaValidationExceptionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsSchemaValidationExceptionCodeFromValue(v string) (*SponsoredProductsSchemaValidationExceptionCode, error) {
	ev := SponsoredProductsSchemaValidationExceptionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsSchemaValidationExceptionCode: valid values are %v", v, AllowedSponsoredProductsSchemaValidationExceptionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsSchemaValidationExceptionCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsSchemaValidationExceptionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsSchemaValidationExceptionCode value
func (v SponsoredProductsSchemaValidationExceptionCode) Ptr() *SponsoredProductsSchemaValidationExceptionCode {
	return &v
}

type NullableSponsoredProductsSchemaValidationExceptionCode struct {
	value *SponsoredProductsSchemaValidationExceptionCode
	isSet bool
}

func (v NullableSponsoredProductsSchemaValidationExceptionCode) Get() *SponsoredProductsSchemaValidationExceptionCode {
	return v.value
}

func (v *NullableSponsoredProductsSchemaValidationExceptionCode) Set(val *SponsoredProductsSchemaValidationExceptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsSchemaValidationExceptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsSchemaValidationExceptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsSchemaValidationExceptionCode(val *SponsoredProductsSchemaValidationExceptionCode) *NullableSponsoredProductsSchemaValidationExceptionCode {
	return &NullableSponsoredProductsSchemaValidationExceptionCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsSchemaValidationExceptionCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsSchemaValidationExceptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
