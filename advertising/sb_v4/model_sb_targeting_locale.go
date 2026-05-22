package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SBTargetingLocale The locale to which the caller wishes to translate the targetable categories or refinements to. For example, if the caller wishes to receive the targetable categories in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned tagetable categories will be in the default language of the marketplace.
type SBTargetingLocale string

// List of SBTargetingLocale
const (
	SBTARGETINGLOCALE_AR_AE SBTargetingLocale = "ar_AE"
	SBTARGETINGLOCALE_DE_DE SBTargetingLocale = "de_DE"
	SBTARGETINGLOCALE_EN_AE SBTargetingLocale = "en_AE"
	SBTARGETINGLOCALE_EN_AU SBTargetingLocale = "en_AU"
	SBTARGETINGLOCALE_EN_CA SBTargetingLocale = "en_CA"
	SBTARGETINGLOCALE_EN_GB SBTargetingLocale = "en_GB"
	SBTARGETINGLOCALE_EN_IN SBTargetingLocale = "en_IN"
	SBTARGETINGLOCALE_EN_SG SBTargetingLocale = "en_SG"
	SBTARGETINGLOCALE_EN_US SBTargetingLocale = "en_US"
	SBTARGETINGLOCALE_ES_ES SBTargetingLocale = "es_ES"
	SBTARGETINGLOCALE_ES_MX SBTargetingLocale = "es_MX"
	SBTARGETINGLOCALE_FR_CA SBTargetingLocale = "fr_CA"
	SBTARGETINGLOCALE_FR_FR SBTargetingLocale = "fr_FR"
	SBTARGETINGLOCALE_HI_IN SBTargetingLocale = "hi_IN"
	SBTARGETINGLOCALE_IT_IT SBTargetingLocale = "it_IT"
	SBTARGETINGLOCALE_JA_JP SBTargetingLocale = "ja_JP"
	SBTARGETINGLOCALE_KO_KR SBTargetingLocale = "ko_KR"
	SBTARGETINGLOCALE_NL_NL SBTargetingLocale = "nl_NL"
	SBTARGETINGLOCALE_PL_PL SBTargetingLocale = "pl_PL"
	SBTARGETINGLOCALE_PT_BR SBTargetingLocale = "pt_BR"
	SBTARGETINGLOCALE_SV_SE SBTargetingLocale = "sv_SE"
	SBTARGETINGLOCALE_TA_IN SBTargetingLocale = "ta_IN"
	SBTARGETINGLOCALE_TH_TH SBTargetingLocale = "th_TH"
	SBTARGETINGLOCALE_TR_TR SBTargetingLocale = "tr_TR"
	SBTARGETINGLOCALE_VI_VN SBTargetingLocale = "vi_VN"
	SBTARGETINGLOCALE_ZH_CN SBTargetingLocale = "zh_CN"
)

// All allowed values of SBTargetingLocale enum
var AllowedSBTargetingLocaleEnumValues = []SBTargetingLocale{
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

func (v *SBTargetingLocale) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SBTargetingLocale(value)
	for _, existing := range AllowedSBTargetingLocaleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SBTargetingLocale", value)
}

// NewSBTargetingLocaleFromValue returns a pointer to a valid SBTargetingLocale
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSBTargetingLocaleFromValue(v string) (*SBTargetingLocale, error) {
	ev := SBTargetingLocale(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SBTargetingLocale: valid values are %v", v, AllowedSBTargetingLocaleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SBTargetingLocale) IsValid() bool {
	for _, existing := range AllowedSBTargetingLocaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBTargetingLocale value
func (v SBTargetingLocale) Ptr() *SBTargetingLocale {
	return &v
}

type NullableSBTargetingLocale struct {
	value *SBTargetingLocale
	isSet bool
}

func (v NullableSBTargetingLocale) Get() *SBTargetingLocale {
	return v.value
}

func (v *NullableSBTargetingLocale) Set(val *SBTargetingLocale) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingLocale(val *SBTargetingLocale) *NullableSBTargetingLocale {
	return &NullableSBTargetingLocale{value: val, isSet: true}
}

func (v NullableSBTargetingLocale) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
