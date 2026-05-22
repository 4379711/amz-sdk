package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// Locale Locale string as described in [BCP 47](https://tools.ietf.org/html/bcp47). For example, `en-US`
type Locale string

// List of Locale
const (
	LOCALE_EN_US Locale = "en-US"
	LOCALE_ES_MX Locale = "es-MX"
	LOCALE_ZH_CN Locale = "zh-CN"
	LOCALE_ES_ES Locale = "es-ES"
	LOCALE_IT_IT Locale = "it-IT"
	LOCALE_FR_FR Locale = "fr-FR"
	LOCALE_FR_CA Locale = "fr-CA"
	LOCALE_DE_DE Locale = "de-DE"
	LOCALE_JA_JP Locale = "ja-JP"
	LOCALE_KO_KR Locale = "ko-KR"
	LOCALE_EN_GB Locale = "en-GB"
	LOCALE_EN_CA Locale = "en-CA"
	LOCALE_HI_IN Locale = "hi-IN"
	LOCALE_EN_IN Locale = "en-IN"
	LOCALE_EN_DE Locale = "en-DE"
	LOCALE_EN_ES Locale = "en-ES"
	LOCALE_EN_FR Locale = "en-FR"
	LOCALE_EN_IT Locale = "en-IT"
	LOCALE_EN_JP Locale = "en-JP"
	LOCALE_EN_AE Locale = "en-AE"
	LOCALE_AR_AE Locale = "ar-AE"
)

// All allowed values of Locale enum
var AllowedLocaleEnumValues = []Locale{
	"en-US",
	"es-MX",
	"zh-CN",
	"es-ES",
	"it-IT",
	"fr-FR",
	"fr-CA",
	"de-DE",
	"ja-JP",
	"ko-KR",
	"en-GB",
	"en-CA",
	"hi-IN",
	"en-IN",
	"en-DE",
	"en-ES",
	"en-FR",
	"en-IT",
	"en-JP",
	"en-AE",
	"ar-AE",
}

func (v *Locale) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Locale(value)
	for _, existing := range AllowedLocaleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Locale", value)
}

// NewLocaleFromValue returns a pointer to a valid Locale
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocaleFromValue(v string) (*Locale, error) {
	ev := Locale(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Locale: valid values are %v", v, AllowedLocaleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Locale) IsValid() bool {
	for _, existing := range AllowedLocaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Locale value
func (v Locale) Ptr() *Locale {
	return &v
}

type NullableLocale struct {
	value *Locale
	isSet bool
}

func (v NullableLocale) Get() *Locale {
	return v.value
}

func (v *NullableLocale) Set(val *Locale) {
	v.value = val
	v.isSet = true
}

func (v NullableLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocale(val *Locale) *NullableLocale {
	return &NullableLocale{value: val, isSet: true}
}

func (v NullableLocale) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
