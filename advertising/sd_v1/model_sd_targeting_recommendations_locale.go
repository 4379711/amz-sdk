package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SDTargetingRecommendationsLocale List of supported locales.
type SDTargetingRecommendationsLocale string

// List of SDTargetingRecommendationsLocale
const (
	SDTARGETINGRECOMMENDATIONSLOCALE_AR_AE SDTargetingRecommendationsLocale = "ar_AE"
	SDTARGETINGRECOMMENDATIONSLOCALE_DE_DE SDTargetingRecommendationsLocale = "de_DE"
	SDTARGETINGRECOMMENDATIONSLOCALE_EN_AE SDTargetingRecommendationsLocale = "en_AE"
	SDTARGETINGRECOMMENDATIONSLOCALE_EN_AU SDTargetingRecommendationsLocale = "en_AU"
	SDTARGETINGRECOMMENDATIONSLOCALE_EN_CA SDTargetingRecommendationsLocale = "en_CA"
	SDTARGETINGRECOMMENDATIONSLOCALE_EN_GB SDTargetingRecommendationsLocale = "en_GB"
	SDTARGETINGRECOMMENDATIONSLOCALE_EN_IN SDTargetingRecommendationsLocale = "en_IN"
	SDTARGETINGRECOMMENDATIONSLOCALE_EN_SG SDTargetingRecommendationsLocale = "en_SG"
	SDTARGETINGRECOMMENDATIONSLOCALE_EN_US SDTargetingRecommendationsLocale = "en_US"
	SDTARGETINGRECOMMENDATIONSLOCALE_ES_ES SDTargetingRecommendationsLocale = "es_ES"
	SDTARGETINGRECOMMENDATIONSLOCALE_ES_MX SDTargetingRecommendationsLocale = "es_MX"
	SDTARGETINGRECOMMENDATIONSLOCALE_FR_CA SDTargetingRecommendationsLocale = "fr_CA"
	SDTARGETINGRECOMMENDATIONSLOCALE_FR_FR SDTargetingRecommendationsLocale = "fr_FR"
	SDTARGETINGRECOMMENDATIONSLOCALE_HI_IN SDTargetingRecommendationsLocale = "hi_IN"
	SDTARGETINGRECOMMENDATIONSLOCALE_IT_IT SDTargetingRecommendationsLocale = "it_IT"
	SDTARGETINGRECOMMENDATIONSLOCALE_JA_JP SDTargetingRecommendationsLocale = "ja_JP"
	SDTARGETINGRECOMMENDATIONSLOCALE_KO_KR SDTargetingRecommendationsLocale = "ko_KR"
	SDTARGETINGRECOMMENDATIONSLOCALE_NL_NL SDTargetingRecommendationsLocale = "nl_NL"
	SDTARGETINGRECOMMENDATIONSLOCALE_PL_PL SDTargetingRecommendationsLocale = "pl_PL"
	SDTARGETINGRECOMMENDATIONSLOCALE_PT_BR SDTargetingRecommendationsLocale = "pt_BR"
	SDTARGETINGRECOMMENDATIONSLOCALE_SV_SE SDTargetingRecommendationsLocale = "sv_SE"
	SDTARGETINGRECOMMENDATIONSLOCALE_TA_IN SDTargetingRecommendationsLocale = "ta_IN"
	SDTARGETINGRECOMMENDATIONSLOCALE_TH_TH SDTargetingRecommendationsLocale = "th_TH"
	SDTARGETINGRECOMMENDATIONSLOCALE_TR_TR SDTargetingRecommendationsLocale = "tr_TR"
	SDTARGETINGRECOMMENDATIONSLOCALE_VI_VN SDTargetingRecommendationsLocale = "vi_VN"
	SDTARGETINGRECOMMENDATIONSLOCALE_ZH_CN SDTargetingRecommendationsLocale = "zh_CN"
)

// All allowed values of SDTargetingRecommendationsLocale enum
var AllowedSDTargetingRecommendationsLocaleEnumValues = []SDTargetingRecommendationsLocale{
	"ar_AE",
	"de_DE",
	"en_AE",
	"en_AU",
	"en_CA",
	"en_GB",
	"en_IN",
	"en_SG",
	"en_US",
	"es_ES",
	"es_MX",
	"fr_CA",
	"fr_FR",
	"hi_IN",
	"it_IT",
	"ja_JP",
	"ko_KR",
	"nl_NL",
	"pl_PL",
	"pt_BR",
	"sv_SE",
	"ta_IN",
	"th_TH",
	"tr_TR",
	"vi_VN",
	"zh_CN",
}

func (v *SDTargetingRecommendationsLocale) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDTargetingRecommendationsLocale(value)
	for _, existing := range AllowedSDTargetingRecommendationsLocaleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDTargetingRecommendationsLocale", value)
}

// NewSDTargetingRecommendationsLocaleFromValue returns a pointer to a valid SDTargetingRecommendationsLocale
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDTargetingRecommendationsLocaleFromValue(v string) (*SDTargetingRecommendationsLocale, error) {
	ev := SDTargetingRecommendationsLocale(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDTargetingRecommendationsLocale: valid values are %v", v, AllowedSDTargetingRecommendationsLocaleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDTargetingRecommendationsLocale) IsValid() bool {
	for _, existing := range AllowedSDTargetingRecommendationsLocaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDTargetingRecommendationsLocale value
func (v SDTargetingRecommendationsLocale) Ptr() *SDTargetingRecommendationsLocale {
	return &v
}

type NullableSDTargetingRecommendationsLocale struct {
	value *SDTargetingRecommendationsLocale
	isSet bool
}

func (v NullableSDTargetingRecommendationsLocale) Get() *SDTargetingRecommendationsLocale {
	return v.value
}

func (v *NullableSDTargetingRecommendationsLocale) Set(val *SDTargetingRecommendationsLocale) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsLocale(val *SDTargetingRecommendationsLocale) *NullableSDTargetingRecommendationsLocale {
	return &NullableSDTargetingRecommendationsLocale{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsLocale) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
