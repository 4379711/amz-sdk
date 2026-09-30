package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsKeywordMatchType the model 'SponsoredProductsKeywordMatchType'
type SponsoredProductsKeywordMatchType string

// List of SponsoredProductsKeywordMatchType
const (
	SPONSOREDPRODUCTSKEYWORDMATCHTYPE_BROAD  SponsoredProductsKeywordMatchType = "BROAD"
	SPONSOREDPRODUCTSKEYWORDMATCHTYPE_PHRASE SponsoredProductsKeywordMatchType = "PHRASE"
	SPONSOREDPRODUCTSKEYWORDMATCHTYPE_EXACT  SponsoredProductsKeywordMatchType = "EXACT"
)

// All allowed values of SponsoredProductsKeywordMatchType enum
var AllowedSponsoredProductsKeywordMatchTypeEnumValues = []SponsoredProductsKeywordMatchType{
	"BROAD",
	"PHRASE",
	"EXACT",
}

func (v *SponsoredProductsKeywordMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsKeywordMatchType(value)
	return nil
}

// NewSponsoredProductsKeywordMatchTypeFromValue returns a pointer to a valid SponsoredProductsKeywordMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsKeywordMatchTypeFromValue(v string) (*SponsoredProductsKeywordMatchType, error) {
	ev := SponsoredProductsKeywordMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsKeywordMatchType: valid values are %v", v, AllowedSponsoredProductsKeywordMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsKeywordMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsKeywordMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsKeywordMatchType value
func (v SponsoredProductsKeywordMatchType) Ptr() *SponsoredProductsKeywordMatchType {
	return &v
}

type NullableSponsoredProductsKeywordMatchType struct {
	value *SponsoredProductsKeywordMatchType
	isSet bool
}

func (v NullableSponsoredProductsKeywordMatchType) Get() *SponsoredProductsKeywordMatchType {
	return v.value
}

func (v *NullableSponsoredProductsKeywordMatchType) Set(val *SponsoredProductsKeywordMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordMatchType(val *SponsoredProductsKeywordMatchType) *NullableSponsoredProductsKeywordMatchType {
	return &NullableSponsoredProductsKeywordMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
