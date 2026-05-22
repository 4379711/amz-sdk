package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsUnauthenticatedExceptionCode the model 'SponsoredProductsUnauthenticatedExceptionCode'
type SponsoredProductsUnauthenticatedExceptionCode string

// List of SponsoredProductsUnauthenticatedExceptionCode
const (
	SPONSOREDPRODUCTSUNAUTHENTICATEDEXCEPTIONCODE_UNAUTHENTICATED SponsoredProductsUnauthenticatedExceptionCode = "unauthenticated"
)

// All allowed values of SponsoredProductsUnauthenticatedExceptionCode enum
var AllowedSponsoredProductsUnauthenticatedExceptionCodeEnumValues = []SponsoredProductsUnauthenticatedExceptionCode{
	"unauthenticated",
}

func (v *SponsoredProductsUnauthenticatedExceptionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsUnauthenticatedExceptionCode(value)
	for _, existing := range AllowedSponsoredProductsUnauthenticatedExceptionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsUnauthenticatedExceptionCode", value)
}

// NewSponsoredProductsUnauthenticatedExceptionCodeFromValue returns a pointer to a valid SponsoredProductsUnauthenticatedExceptionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsUnauthenticatedExceptionCodeFromValue(v string) (*SponsoredProductsUnauthenticatedExceptionCode, error) {
	ev := SponsoredProductsUnauthenticatedExceptionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsUnauthenticatedExceptionCode: valid values are %v", v, AllowedSponsoredProductsUnauthenticatedExceptionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsUnauthenticatedExceptionCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsUnauthenticatedExceptionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsUnauthenticatedExceptionCode value
func (v SponsoredProductsUnauthenticatedExceptionCode) Ptr() *SponsoredProductsUnauthenticatedExceptionCode {
	return &v
}

type NullableSponsoredProductsUnauthenticatedExceptionCode struct {
	value *SponsoredProductsUnauthenticatedExceptionCode
	isSet bool
}

func (v NullableSponsoredProductsUnauthenticatedExceptionCode) Get() *SponsoredProductsUnauthenticatedExceptionCode {
	return v.value
}

func (v *NullableSponsoredProductsUnauthenticatedExceptionCode) Set(val *SponsoredProductsUnauthenticatedExceptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnauthenticatedExceptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnauthenticatedExceptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnauthenticatedExceptionCode(val *SponsoredProductsUnauthenticatedExceptionCode) *NullableSponsoredProductsUnauthenticatedExceptionCode {
	return &NullableSponsoredProductsUnauthenticatedExceptionCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnauthenticatedExceptionCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnauthenticatedExceptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
