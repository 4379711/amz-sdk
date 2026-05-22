package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsLocale the model 'SponsoredProductsLocale'
type SponsoredProductsLocale string

// List of SponsoredProductsLocale
const (
	SPONSOREDPRODUCTSLOCALE_EN_AE SponsoredProductsLocale = "en_AE"
	SPONSOREDPRODUCTSLOCALE_DE_DE SponsoredProductsLocale = "de_DE"
	SPONSOREDPRODUCTSLOCALE_AR_EG SponsoredProductsLocale = "ar_EG"
	SPONSOREDPRODUCTSLOCALE_ES_ES SponsoredProductsLocale = "es_ES"
	SPONSOREDPRODUCTSLOCALE_FR_FR SponsoredProductsLocale = "fr_FR"
	SPONSOREDPRODUCTSLOCALE_EN_GB SponsoredProductsLocale = "en_GB"
	SPONSOREDPRODUCTSLOCALE_IT_IT SponsoredProductsLocale = "it_IT"
	SPONSOREDPRODUCTSLOCALE_NL_NL SponsoredProductsLocale = "nl_NL"
	SPONSOREDPRODUCTSLOCALE_PL_PL SponsoredProductsLocale = "pl_PL"
	SPONSOREDPRODUCTSLOCALE_EN_SA SponsoredProductsLocale = "en_SA"
	SPONSOREDPRODUCTSLOCALE_SV_SE SponsoredProductsLocale = "sv_SE"
	SPONSOREDPRODUCTSLOCALE_TR_TR SponsoredProductsLocale = "tr_TR"
	SPONSOREDPRODUCTSLOCALE_EN_AU SponsoredProductsLocale = "en_AU"
	SPONSOREDPRODUCTSLOCALE_EN_IN SponsoredProductsLocale = "en_IN"
	SPONSOREDPRODUCTSLOCALE_JA_JP SponsoredProductsLocale = "ja_JP"
	SPONSOREDPRODUCTSLOCALE_EN_SG SponsoredProductsLocale = "en_SG"
	SPONSOREDPRODUCTSLOCALE_PT_BR SponsoredProductsLocale = "pt_BR"
	SPONSOREDPRODUCTSLOCALE_EN_CA SponsoredProductsLocale = "en_CA"
	SPONSOREDPRODUCTSLOCALE_ES_MX SponsoredProductsLocale = "es_MX"
	SPONSOREDPRODUCTSLOCALE_EN_US SponsoredProductsLocale = "en_US"
)

// All allowed values of SponsoredProductsLocale enum
var AllowedSponsoredProductsLocaleEnumValues = []SponsoredProductsLocale{
	"en_AE",
	"de_DE",
	"ar_EG",
	"es_ES",
	"fr_FR",
	"en_GB",
	"it_IT",
	"nl_NL",
	"pl_PL",
	"en_SA",
	"sv_SE",
	"tr_TR",
	"en_AU",
	"en_IN",
	"ja_JP",
	"en_SG",
	"pt_BR",
	"en_CA",
	"es_MX",
	"en_US",
}

func (v *SponsoredProductsLocale) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsLocale(value)
	for _, existing := range AllowedSponsoredProductsLocaleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsLocale", value)
}

// NewSponsoredProductsLocaleFromValue returns a pointer to a valid SponsoredProductsLocale
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsLocaleFromValue(v string) (*SponsoredProductsLocale, error) {
	ev := SponsoredProductsLocale(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsLocale: valid values are %v", v, AllowedSponsoredProductsLocaleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsLocale) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsLocaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsLocale value
func (v SponsoredProductsLocale) Ptr() *SponsoredProductsLocale {
	return &v
}

type NullableSponsoredProductsLocale struct {
	value *SponsoredProductsLocale
	isSet bool
}

func (v NullableSponsoredProductsLocale) Get() *SponsoredProductsLocale {
	return v.value
}

func (v *NullableSponsoredProductsLocale) Set(val *SponsoredProductsLocale) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsLocale(val *SponsoredProductsLocale) *NullableSponsoredProductsLocale {
	return &NullableSponsoredProductsLocale{value: val, isSet: true}
}

func (v NullableSponsoredProductsLocale) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
