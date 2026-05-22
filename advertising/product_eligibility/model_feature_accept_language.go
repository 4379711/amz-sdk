package product_eligibility

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// FeatureAcceptLanguage the model 'FeatureAcceptLanguage'
type FeatureAcceptLanguage string

// List of FeatureAcceptLanguage
const (
	FEATUREACCEPTLANGUAGE_EN_US FeatureAcceptLanguage = "en_US"
	FEATUREACCEPTLANGUAGE_AR_AE FeatureAcceptLanguage = "ar_AE"
	FEATUREACCEPTLANGUAGE_ZH_CN FeatureAcceptLanguage = "zh_CN"
	FEATUREACCEPTLANGUAGE_NL_NL FeatureAcceptLanguage = "nl_NL"
	FEATUREACCEPTLANGUAGE_EN_AU FeatureAcceptLanguage = "en_AU"
	FEATUREACCEPTLANGUAGE_EN_CA FeatureAcceptLanguage = "en_CA"
	FEATUREACCEPTLANGUAGE_EN_IN FeatureAcceptLanguage = "en_IN"
	FEATUREACCEPTLANGUAGE_EN_GB FeatureAcceptLanguage = "en_GB"
	FEATUREACCEPTLANGUAGE_FR_CA FeatureAcceptLanguage = "fr_CA"
	FEATUREACCEPTLANGUAGE_FR_FR FeatureAcceptLanguage = "fr_FR"
	FEATUREACCEPTLANGUAGE_DE_DE FeatureAcceptLanguage = "de_DE"
	FEATUREACCEPTLANGUAGE_IT_IT FeatureAcceptLanguage = "it_IT"
	FEATUREACCEPTLANGUAGE_JA_JP FeatureAcceptLanguage = "ja_JP"
	FEATUREACCEPTLANGUAGE_KO_KR FeatureAcceptLanguage = "ko_KR"
	FEATUREACCEPTLANGUAGE_PL_PL FeatureAcceptLanguage = "pl_PL"
	FEATUREACCEPTLANGUAGE_PT_BR FeatureAcceptLanguage = "pt_BR"
	FEATUREACCEPTLANGUAGE_ES_ES FeatureAcceptLanguage = "es_ES"
	FEATUREACCEPTLANGUAGE_ES_US FeatureAcceptLanguage = "es_US"
	FEATUREACCEPTLANGUAGE_ES_MX FeatureAcceptLanguage = "es_MX"
	FEATUREACCEPTLANGUAGE_TR_TR FeatureAcceptLanguage = "tr_TR"
)

// All allowed values of FeatureAcceptLanguage enum
var AllowedFeatureAcceptLanguageEnumValues = []FeatureAcceptLanguage{
	"en_US",
	"ar_AE",
	"zh_CN",
	"nl_NL",
	"en_AU",
	"en_CA",
	"en_IN",
	"en_GB",
	"fr_CA",
	"fr_FR",
	"de_DE",
	"it_IT",
	"ja_JP",
	"ko_KR",
	"pl_PL",
	"pt_BR",
	"es_ES",
	"es_US",
	"es_MX",
	"tr_TR",
}

func (v *FeatureAcceptLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeatureAcceptLanguage(value)
	for _, existing := range AllowedFeatureAcceptLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeatureAcceptLanguage", value)
}

// NewFeatureAcceptLanguageFromValue returns a pointer to a valid FeatureAcceptLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeatureAcceptLanguageFromValue(v string) (*FeatureAcceptLanguage, error) {
	ev := FeatureAcceptLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeatureAcceptLanguage: valid values are %v", v, AllowedFeatureAcceptLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeatureAcceptLanguage) IsValid() bool {
	for _, existing := range AllowedFeatureAcceptLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureAcceptLanguage value
func (v FeatureAcceptLanguage) Ptr() *FeatureAcceptLanguage {
	return &v
}

type NullableFeatureAcceptLanguage struct {
	value *FeatureAcceptLanguage
	isSet bool
}

func (v NullableFeatureAcceptLanguage) Get() *FeatureAcceptLanguage {
	return v.value
}

func (v *NullableFeatureAcceptLanguage) Set(val *FeatureAcceptLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureAcceptLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureAcceptLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureAcceptLanguage(val *FeatureAcceptLanguage) *NullableFeatureAcceptLanguage {
	return &NullableFeatureAcceptLanguage{value: val, isSet: true}
}

func (v NullableFeatureAcceptLanguage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeatureAcceptLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
