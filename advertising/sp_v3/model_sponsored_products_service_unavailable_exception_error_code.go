package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsServiceUnavailableExceptionErrorCode the model 'SponsoredProductsServiceUnavailableExceptionErrorCode'
type SponsoredProductsServiceUnavailableExceptionErrorCode string

// List of SponsoredProductsServiceUnavailableExceptionErrorCode
const (
	SPONSOREDPRODUCTSSERVICEUNAVAILABLEEXCEPTIONERRORCODE_SERVICE_UNAVAILABLE_EXCEPTION SponsoredProductsServiceUnavailableExceptionErrorCode = "SERVICE_UNAVAILABLE_EXCEPTION"
)

// All allowed values of SponsoredProductsServiceUnavailableExceptionErrorCode enum
var AllowedSponsoredProductsServiceUnavailableExceptionErrorCodeEnumValues = []SponsoredProductsServiceUnavailableExceptionErrorCode{
	"SERVICE_UNAVAILABLE_EXCEPTION",
}

func (v *SponsoredProductsServiceUnavailableExceptionErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsServiceUnavailableExceptionErrorCode(value)
	for _, existing := range AllowedSponsoredProductsServiceUnavailableExceptionErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsServiceUnavailableExceptionErrorCode", value)
}

// NewSponsoredProductsServiceUnavailableExceptionErrorCodeFromValue returns a pointer to a valid SponsoredProductsServiceUnavailableExceptionErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsServiceUnavailableExceptionErrorCodeFromValue(v string) (*SponsoredProductsServiceUnavailableExceptionErrorCode, error) {
	ev := SponsoredProductsServiceUnavailableExceptionErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsServiceUnavailableExceptionErrorCode: valid values are %v", v, AllowedSponsoredProductsServiceUnavailableExceptionErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsServiceUnavailableExceptionErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsServiceUnavailableExceptionErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsServiceUnavailableExceptionErrorCode value
func (v SponsoredProductsServiceUnavailableExceptionErrorCode) Ptr() *SponsoredProductsServiceUnavailableExceptionErrorCode {
	return &v
}

type NullableSponsoredProductsServiceUnavailableExceptionErrorCode struct {
	value *SponsoredProductsServiceUnavailableExceptionErrorCode
	isSet bool
}

func (v NullableSponsoredProductsServiceUnavailableExceptionErrorCode) Get() *SponsoredProductsServiceUnavailableExceptionErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionErrorCode) Set(val *SponsoredProductsServiceUnavailableExceptionErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsServiceUnavailableExceptionErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsServiceUnavailableExceptionErrorCode(val *SponsoredProductsServiceUnavailableExceptionErrorCode) *NullableSponsoredProductsServiceUnavailableExceptionErrorCode {
	return &NullableSponsoredProductsServiceUnavailableExceptionErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsServiceUnavailableExceptionErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsServiceUnavailableExceptionErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
