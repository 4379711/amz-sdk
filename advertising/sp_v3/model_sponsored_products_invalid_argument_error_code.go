package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsInvalidArgumentErrorCode the model 'SponsoredProductsInvalidArgumentErrorCode'
type SponsoredProductsInvalidArgumentErrorCode string

// List of SponsoredProductsInvalidArgumentErrorCode
const (
	SPONSOREDPRODUCTSINVALIDARGUMENTERRORCODE_INVALID_ARGUMENT SponsoredProductsInvalidArgumentErrorCode = "INVALID_ARGUMENT"
)

// All allowed values of SponsoredProductsInvalidArgumentErrorCode enum
var AllowedSponsoredProductsInvalidArgumentErrorCodeEnumValues = []SponsoredProductsInvalidArgumentErrorCode{
	"INVALID_ARGUMENT",
}

func (v *SponsoredProductsInvalidArgumentErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsInvalidArgumentErrorCode(value)
	for _, existing := range AllowedSponsoredProductsInvalidArgumentErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsInvalidArgumentErrorCode", value)
}

// NewSponsoredProductsInvalidArgumentErrorCodeFromValue returns a pointer to a valid SponsoredProductsInvalidArgumentErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsInvalidArgumentErrorCodeFromValue(v string) (*SponsoredProductsInvalidArgumentErrorCode, error) {
	ev := SponsoredProductsInvalidArgumentErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsInvalidArgumentErrorCode: valid values are %v", v, AllowedSponsoredProductsInvalidArgumentErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsInvalidArgumentErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsInvalidArgumentErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsInvalidArgumentErrorCode value
func (v SponsoredProductsInvalidArgumentErrorCode) Ptr() *SponsoredProductsInvalidArgumentErrorCode {
	return &v
}

type NullableSponsoredProductsInvalidArgumentErrorCode struct {
	value *SponsoredProductsInvalidArgumentErrorCode
	isSet bool
}

func (v NullableSponsoredProductsInvalidArgumentErrorCode) Get() *SponsoredProductsInvalidArgumentErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsInvalidArgumentErrorCode) Set(val *SponsoredProductsInvalidArgumentErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInvalidArgumentErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInvalidArgumentErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInvalidArgumentErrorCode(val *SponsoredProductsInvalidArgumentErrorCode) *NullableSponsoredProductsInvalidArgumentErrorCode {
	return &NullableSponsoredProductsInvalidArgumentErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsInvalidArgumentErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInvalidArgumentErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
