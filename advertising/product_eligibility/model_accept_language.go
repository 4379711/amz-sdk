package product_eligibility

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AcceptLanguage the model 'AcceptLanguage'
type AcceptLanguage string

// List of AcceptLanguage
const (
	ACCEPTLANGUAGE_EN_US AcceptLanguage = "en-US"
	ACCEPTLANGUAGE_AR_AE AcceptLanguage = "ar-AE"
	ACCEPTLANGUAGE_ZH_CN AcceptLanguage = "zh-CN"
	ACCEPTLANGUAGE_NL_NL AcceptLanguage = "nl-NL"
	ACCEPTLANGUAGE_EN_AU AcceptLanguage = "en-AU"
	ACCEPTLANGUAGE_EN_CA AcceptLanguage = "en-CA"
	ACCEPTLANGUAGE_EN_IN AcceptLanguage = "en-IN"
	ACCEPTLANGUAGE_EN_GB AcceptLanguage = "en-GB"
	ACCEPTLANGUAGE_FR_CA AcceptLanguage = "fr-CA"
	ACCEPTLANGUAGE_FR_FR AcceptLanguage = "fr-FR"
	ACCEPTLANGUAGE_DE_DE AcceptLanguage = "de-DE"
	ACCEPTLANGUAGE_IT_IT AcceptLanguage = "it-IT"
	ACCEPTLANGUAGE_JA_JP AcceptLanguage = "ja-JP"
	ACCEPTLANGUAGE_KO_KR AcceptLanguage = "ko-KR"
	ACCEPTLANGUAGE_PL_PL AcceptLanguage = "pl-PL"
	ACCEPTLANGUAGE_PT_BR AcceptLanguage = "pt-BR"
	ACCEPTLANGUAGE_ES_ES AcceptLanguage = "es-ES"
	ACCEPTLANGUAGE_ES_US AcceptLanguage = "es-US"
	ACCEPTLANGUAGE_ES_MX AcceptLanguage = "es-MX"
	ACCEPTLANGUAGE_TR_TR AcceptLanguage = "tr-TR"
	ACCEPTLANGUAGE_EN_AE AcceptLanguage = "en-AE"
	ACCEPTLANGUAGE_EN_SG AcceptLanguage = "en-SG"
	ACCEPTLANGUAGE_ES_CO AcceptLanguage = "es-CO"
	ACCEPTLANGUAGE_HI_IN AcceptLanguage = "hi-IN"
	ACCEPTLANGUAGE_SV_SE AcceptLanguage = "sv-SE"
	ACCEPTLANGUAGE_TA_IN AcceptLanguage = "ta-IN"
	ACCEPTLANGUAGE_TH_TH AcceptLanguage = "th-TH"
	ACCEPTLANGUAGE_VI_VN AcceptLanguage = "vi-VN"
	ACCEPTLANGUAGE_ZH_TW AcceptLanguage = "zh-TW"
)

// All allowed values of AcceptLanguage enum
var AllowedAcceptLanguageEnumValues = []AcceptLanguage{
	"en-US",
	"ar-AE",
	"zh-CN",
	"nl-NL",
	"en-AU",
	"en-CA",
	"en-IN",
	"en-GB",
	"fr-CA",
	"fr-FR",
	"de-DE",
	"it-IT",
	"ja-JP",
	"ko-KR",
	"pl-PL",
	"pt-BR",
	"es-ES",
	"es-US",
	"es-MX",
	"tr-TR",
	"en-AE",
	"en-SG",
	"es-CO",
	"hi-IN",
	"sv-SE",
	"ta-IN",
	"th-TH",
	"vi-VN",
	"zh-TW",
}

func (v *AcceptLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = AcceptLanguage(value)
	return nil
}

// NewAcceptLanguageFromValue returns a pointer to a valid AcceptLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcceptLanguageFromValue(v string) (*AcceptLanguage, error) {
	ev := AcceptLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcceptLanguage: valid values are %v", v, AllowedAcceptLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcceptLanguage) IsValid() bool {
	for _, existing := range AllowedAcceptLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcceptLanguage value
func (v AcceptLanguage) Ptr() *AcceptLanguage {
	return &v
}

type NullableAcceptLanguage struct {
	value *AcceptLanguage
	isSet bool
}

func (v NullableAcceptLanguage) Get() *AcceptLanguage {
	return v.value
}

func (v *NullableAcceptLanguage) Set(val *AcceptLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptLanguage(val *AcceptLanguage) *NullableAcceptLanguage {
	return &NullableAcceptLanguage{value: val, isSet: true}
}

func (v NullableAcceptLanguage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAcceptLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
