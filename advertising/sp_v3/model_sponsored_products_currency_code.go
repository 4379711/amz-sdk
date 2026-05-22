package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCurrencyCode the model 'SponsoredProductsCurrencyCode'
type SponsoredProductsCurrencyCode string

// List of SponsoredProductsCurrencyCode
const (
	SPONSOREDPRODUCTSCURRENCYCODE_AUD SponsoredProductsCurrencyCode = "AUD"
	SPONSOREDPRODUCTSCURRENCYCODE_BRL SponsoredProductsCurrencyCode = "BRL"
	SPONSOREDPRODUCTSCURRENCYCODE_CAD SponsoredProductsCurrencyCode = "CAD"
	SPONSOREDPRODUCTSCURRENCYCODE_CHF SponsoredProductsCurrencyCode = "CHF"
	SPONSOREDPRODUCTSCURRENCYCODE_CNY SponsoredProductsCurrencyCode = "CNY"
	SPONSOREDPRODUCTSCURRENCYCODE_DKK SponsoredProductsCurrencyCode = "DKK"
	SPONSOREDPRODUCTSCURRENCYCODE_EUR SponsoredProductsCurrencyCode = "EUR"
	SPONSOREDPRODUCTSCURRENCYCODE_GBP SponsoredProductsCurrencyCode = "GBP"
	SPONSOREDPRODUCTSCURRENCYCODE_INR SponsoredProductsCurrencyCode = "INR"
	SPONSOREDPRODUCTSCURRENCYCODE_JPY SponsoredProductsCurrencyCode = "JPY"
	SPONSOREDPRODUCTSCURRENCYCODE_MXN SponsoredProductsCurrencyCode = "MXN"
	SPONSOREDPRODUCTSCURRENCYCODE_NOK SponsoredProductsCurrencyCode = "NOK"
	SPONSOREDPRODUCTSCURRENCYCODE_SAR SponsoredProductsCurrencyCode = "SAR"
	SPONSOREDPRODUCTSCURRENCYCODE_SEK SponsoredProductsCurrencyCode = "SEK"
	SPONSOREDPRODUCTSCURRENCYCODE_SGD SponsoredProductsCurrencyCode = "SGD"
	SPONSOREDPRODUCTSCURRENCYCODE_TRY SponsoredProductsCurrencyCode = "TRY"
	SPONSOREDPRODUCTSCURRENCYCODE_USD SponsoredProductsCurrencyCode = "USD"
)

// All allowed values of SponsoredProductsCurrencyCode enum
var AllowedSponsoredProductsCurrencyCodeEnumValues = []SponsoredProductsCurrencyCode{
	"AUD",
	"BRL",
	"CAD",
	"CHF",
	"CNY",
	"DKK",
	"EUR",
	"GBP",
	"INR",
	"JPY",
	"MXN",
	"NOK",
	"SAR",
	"SEK",
	"SGD",
	"TRY",
	"USD",
}

func (v *SponsoredProductsCurrencyCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCurrencyCode(value)
	for _, existing := range AllowedSponsoredProductsCurrencyCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCurrencyCode", value)
}

// NewSponsoredProductsCurrencyCodeFromValue returns a pointer to a valid SponsoredProductsCurrencyCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCurrencyCodeFromValue(v string) (*SponsoredProductsCurrencyCode, error) {
	ev := SponsoredProductsCurrencyCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCurrencyCode: valid values are %v", v, AllowedSponsoredProductsCurrencyCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCurrencyCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCurrencyCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCurrencyCode value
func (v SponsoredProductsCurrencyCode) Ptr() *SponsoredProductsCurrencyCode {
	return &v
}

type NullableSponsoredProductsCurrencyCode struct {
	value *SponsoredProductsCurrencyCode
	isSet bool
}

func (v NullableSponsoredProductsCurrencyCode) Get() *SponsoredProductsCurrencyCode {
	return v.value
}

func (v *NullableSponsoredProductsCurrencyCode) Set(val *SponsoredProductsCurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCurrencyCode(val *SponsoredProductsCurrencyCode) *NullableSponsoredProductsCurrencyCode {
	return &NullableSponsoredProductsCurrencyCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsCurrencyCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
