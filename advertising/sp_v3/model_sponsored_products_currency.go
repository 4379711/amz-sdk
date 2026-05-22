package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCurrency the model 'SponsoredProductsCurrency'
type SponsoredProductsCurrency string

// List of SponsoredProductsCurrency
const (
	SPONSOREDPRODUCTSCURRENCY_USD SponsoredProductsCurrency = "USD"
	SPONSOREDPRODUCTSCURRENCY_EUR SponsoredProductsCurrency = "EUR"
	SPONSOREDPRODUCTSCURRENCY_GBP SponsoredProductsCurrency = "GBP"
	SPONSOREDPRODUCTSCURRENCY_MXN SponsoredProductsCurrency = "MXN"
	SPONSOREDPRODUCTSCURRENCY_JPY SponsoredProductsCurrency = "JPY"
	SPONSOREDPRODUCTSCURRENCY_INR SponsoredProductsCurrency = "INR"
	SPONSOREDPRODUCTSCURRENCY_PLN SponsoredProductsCurrency = "PLN"
	SPONSOREDPRODUCTSCURRENCY_SGD SponsoredProductsCurrency = "SGD"
	SPONSOREDPRODUCTSCURRENCY_AUD SponsoredProductsCurrency = "AUD"
	SPONSOREDPRODUCTSCURRENCY_CAD SponsoredProductsCurrency = "CAD"
	SPONSOREDPRODUCTSCURRENCY_BRL SponsoredProductsCurrency = "BRL"
	SPONSOREDPRODUCTSCURRENCY_SEK SponsoredProductsCurrency = "SEK"
	SPONSOREDPRODUCTSCURRENCY_TRY SponsoredProductsCurrency = "TRY"
	SPONSOREDPRODUCTSCURRENCY_AED SponsoredProductsCurrency = "AED"
	SPONSOREDPRODUCTSCURRENCY_SAR SponsoredProductsCurrency = "SAR"
	SPONSOREDPRODUCTSCURRENCY_EGP SponsoredProductsCurrency = "EGP"
)

// All allowed values of SponsoredProductsCurrency enum
var AllowedSponsoredProductsCurrencyEnumValues = []SponsoredProductsCurrency{
	"USD",
	"EUR",
	"GBP",
	"MXN",
	"JPY",
	"INR",
	"PLN",
	"SGD",
	"AUD",
	"CAD",
	"BRL",
	"SEK",
	"TRY",
	"AED",
	"SAR",
	"EGP",
}

func (v *SponsoredProductsCurrency) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCurrency(value)
	for _, existing := range AllowedSponsoredProductsCurrencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCurrency", value)
}

// NewSponsoredProductsCurrencyFromValue returns a pointer to a valid SponsoredProductsCurrency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCurrencyFromValue(v string) (*SponsoredProductsCurrency, error) {
	ev := SponsoredProductsCurrency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCurrency: valid values are %v", v, AllowedSponsoredProductsCurrencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCurrency) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCurrencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCurrency value
func (v SponsoredProductsCurrency) Ptr() *SponsoredProductsCurrency {
	return &v
}

type NullableSponsoredProductsCurrency struct {
	value *SponsoredProductsCurrency
	isSet bool
}

func (v NullableSponsoredProductsCurrency) Get() *SponsoredProductsCurrency {
	return v.value
}

func (v *NullableSponsoredProductsCurrency) Set(val *SponsoredProductsCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCurrency(val *SponsoredProductsCurrency) *NullableSponsoredProductsCurrency {
	return &NullableSponsoredProductsCurrency{value: val, isSet: true}
}

func (v NullableSponsoredProductsCurrency) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
