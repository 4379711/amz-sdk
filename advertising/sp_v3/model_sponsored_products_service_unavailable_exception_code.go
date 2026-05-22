package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsServiceUnavailableExceptionCode the model 'SponsoredProductsServiceUnavailableExceptionCode'
type SponsoredProductsServiceUnavailableExceptionCode string

// List of SponsoredProductsServiceUnavailableExceptionCode
const (
	SPONSOREDPRODUCTSSERVICEUNAVAILABLEEXCEPTIONCODE_SERVICE_UNAVAILABLE SponsoredProductsServiceUnavailableExceptionCode = "serviceUnavailable"
)

// All allowed values of SponsoredProductsServiceUnavailableExceptionCode enum
var AllowedSponsoredProductsServiceUnavailableExceptionCodeEnumValues = []SponsoredProductsServiceUnavailableExceptionCode{
	"serviceUnavailable",
}

func (v *SponsoredProductsServiceUnavailableExceptionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsServiceUnavailableExceptionCode(value)
	for _, existing := range AllowedSponsoredProductsServiceUnavailableExceptionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsServiceUnavailableExceptionCode", value)
}

// NewSponsoredProductsServiceUnavailableExceptionCodeFromValue returns a pointer to a valid SponsoredProductsServiceUnavailableExceptionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsServiceUnavailableExceptionCodeFromValue(v string) (*SponsoredProductsServiceUnavailableExceptionCode, error) {
	ev := SponsoredProductsServiceUnavailableExceptionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsServiceUnavailableExceptionCode: valid values are %v", v, AllowedSponsoredProductsServiceUnavailableExceptionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsServiceUnavailableExceptionCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsServiceUnavailableExceptionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsServiceUnavailableExceptionCode value
func (v SponsoredProductsServiceUnavailableExceptionCode) Ptr() *SponsoredProductsServiceUnavailableExceptionCode {
	return &v
}

type NullableSponsoredProductsServiceUnavailableExceptionCode struct {
	value *SponsoredProductsServiceUnavailableExceptionCode
	isSet bool
}

func (v NullableSponsoredProductsServiceUnavailableExceptionCode) Get() *SponsoredProductsServiceUnavailableExceptionCode {
	return v.value
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionCode) Set(val *SponsoredProductsServiceUnavailableExceptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsServiceUnavailableExceptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsServiceUnavailableExceptionCode(val *SponsoredProductsServiceUnavailableExceptionCode) *NullableSponsoredProductsServiceUnavailableExceptionCode {
	return &NullableSponsoredProductsServiceUnavailableExceptionCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsServiceUnavailableExceptionCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
