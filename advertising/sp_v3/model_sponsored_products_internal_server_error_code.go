package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsInternalServerErrorCode the model 'SponsoredProductsInternalServerErrorCode'
type SponsoredProductsInternalServerErrorCode string

// List of SponsoredProductsInternalServerErrorCode
const (
	SPONSOREDPRODUCTSINTERNALSERVERERRORCODE_INTERNAL_SERVER_EXCEPTION SponsoredProductsInternalServerErrorCode = "INTERNAL_SERVER_EXCEPTION"
)

// All allowed values of SponsoredProductsInternalServerErrorCode enum
var AllowedSponsoredProductsInternalServerErrorCodeEnumValues = []SponsoredProductsInternalServerErrorCode{
	"INTERNAL_SERVER_EXCEPTION",
}

func (v *SponsoredProductsInternalServerErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsInternalServerErrorCode(value)
	for _, existing := range AllowedSponsoredProductsInternalServerErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsInternalServerErrorCode", value)
}

// NewSponsoredProductsInternalServerErrorCodeFromValue returns a pointer to a valid SponsoredProductsInternalServerErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsInternalServerErrorCodeFromValue(v string) (*SponsoredProductsInternalServerErrorCode, error) {
	ev := SponsoredProductsInternalServerErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsInternalServerErrorCode: valid values are %v", v, AllowedSponsoredProductsInternalServerErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsInternalServerErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsInternalServerErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsInternalServerErrorCode value
func (v SponsoredProductsInternalServerErrorCode) Ptr() *SponsoredProductsInternalServerErrorCode {
	return &v
}

type NullableSponsoredProductsInternalServerErrorCode struct {
	value *SponsoredProductsInternalServerErrorCode
	isSet bool
}

func (v NullableSponsoredProductsInternalServerErrorCode) Get() *SponsoredProductsInternalServerErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsInternalServerErrorCode) Set(val *SponsoredProductsInternalServerErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInternalServerErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInternalServerErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInternalServerErrorCode(val *SponsoredProductsInternalServerErrorCode) *NullableSponsoredProductsInternalServerErrorCode {
	return &NullableSponsoredProductsInternalServerErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsInternalServerErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInternalServerErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
