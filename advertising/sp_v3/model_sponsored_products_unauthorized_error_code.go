package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsUnauthorizedErrorCode the model 'SponsoredProductsUnauthorizedErrorCode'
type SponsoredProductsUnauthorizedErrorCode string

// List of SponsoredProductsUnauthorizedErrorCode
const (
	SPONSOREDPRODUCTSUNAUTHORIZEDERRORCODE_UNAUTHORIZED SponsoredProductsUnauthorizedErrorCode = "UNAUTHORIZED"
)

// All allowed values of SponsoredProductsUnauthorizedErrorCode enum
var AllowedSponsoredProductsUnauthorizedErrorCodeEnumValues = []SponsoredProductsUnauthorizedErrorCode{
	"UNAUTHORIZED",
}

func (v *SponsoredProductsUnauthorizedErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsUnauthorizedErrorCode(value)
	for _, existing := range AllowedSponsoredProductsUnauthorizedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsUnauthorizedErrorCode", value)
}

// NewSponsoredProductsUnauthorizedErrorCodeFromValue returns a pointer to a valid SponsoredProductsUnauthorizedErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsUnauthorizedErrorCodeFromValue(v string) (*SponsoredProductsUnauthorizedErrorCode, error) {
	ev := SponsoredProductsUnauthorizedErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsUnauthorizedErrorCode: valid values are %v", v, AllowedSponsoredProductsUnauthorizedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsUnauthorizedErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsUnauthorizedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsUnauthorizedErrorCode value
func (v SponsoredProductsUnauthorizedErrorCode) Ptr() *SponsoredProductsUnauthorizedErrorCode {
	return &v
}

type NullableSponsoredProductsUnauthorizedErrorCode struct {
	value *SponsoredProductsUnauthorizedErrorCode
	isSet bool
}

func (v NullableSponsoredProductsUnauthorizedErrorCode) Get() *SponsoredProductsUnauthorizedErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsUnauthorizedErrorCode) Set(val *SponsoredProductsUnauthorizedErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnauthorizedErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnauthorizedErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnauthorizedErrorCode(val *SponsoredProductsUnauthorizedErrorCode) *NullableSponsoredProductsUnauthorizedErrorCode {
	return &NullableSponsoredProductsUnauthorizedErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnauthorizedErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnauthorizedErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
