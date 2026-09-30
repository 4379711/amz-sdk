package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsAdGroupTargetingType Targeting type of an adGroup.
type SponsoredProductsAdGroupTargetingType string

// List of SponsoredProductsAdGroupTargetingType
const (
	SPONSOREDPRODUCTSADGROUPTARGETINGTYPE_KEYWORD SponsoredProductsAdGroupTargetingType = "KEYWORD"
	SPONSOREDPRODUCTSADGROUPTARGETINGTYPE_PRODUCT SponsoredProductsAdGroupTargetingType = "PRODUCT"
	SPONSOREDPRODUCTSADGROUPTARGETINGTYPE_AUTO    SponsoredProductsAdGroupTargetingType = "AUTO"
	SPONSOREDPRODUCTSADGROUPTARGETINGTYPE_OTHER   SponsoredProductsAdGroupTargetingType = "OTHER"
)

// All allowed values of SponsoredProductsAdGroupTargetingType enum
var AllowedSponsoredProductsAdGroupTargetingTypeEnumValues = []SponsoredProductsAdGroupTargetingType{
	"KEYWORD",
	"PRODUCT",
	"AUTO",
	"OTHER",
}

func (v *SponsoredProductsAdGroupTargetingType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsAdGroupTargetingType(value)
	return nil
}

// NewSponsoredProductsAdGroupTargetingTypeFromValue returns a pointer to a valid SponsoredProductsAdGroupTargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAdGroupTargetingTypeFromValue(v string) (*SponsoredProductsAdGroupTargetingType, error) {
	ev := SponsoredProductsAdGroupTargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAdGroupTargetingType: valid values are %v", v, AllowedSponsoredProductsAdGroupTargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAdGroupTargetingType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAdGroupTargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAdGroupTargetingType value
func (v SponsoredProductsAdGroupTargetingType) Ptr() *SponsoredProductsAdGroupTargetingType {
	return &v
}

type NullableSponsoredProductsAdGroupTargetingType struct {
	value *SponsoredProductsAdGroupTargetingType
	isSet bool
}

func (v NullableSponsoredProductsAdGroupTargetingType) Get() *SponsoredProductsAdGroupTargetingType {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupTargetingType) Set(val *SponsoredProductsAdGroupTargetingType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupTargetingType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupTargetingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupTargetingType(val *SponsoredProductsAdGroupTargetingType) *NullableSponsoredProductsAdGroupTargetingType {
	return &NullableSponsoredProductsAdGroupTargetingType{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupTargetingType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupTargetingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
