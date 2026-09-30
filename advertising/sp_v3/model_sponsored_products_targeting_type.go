package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsTargetingType the model 'SponsoredProductsTargetingType'
type SponsoredProductsTargetingType string

// List of SponsoredProductsTargetingType
const (
	SPONSOREDPRODUCTSTARGETINGTYPE_AUTO   SponsoredProductsTargetingType = "AUTO"
	SPONSOREDPRODUCTSTARGETINGTYPE_MANUAL SponsoredProductsTargetingType = "MANUAL"
)

// All allowed values of SponsoredProductsTargetingType enum
var AllowedSponsoredProductsTargetingTypeEnumValues = []SponsoredProductsTargetingType{
	"AUTO",
	"MANUAL",
}

func (v *SponsoredProductsTargetingType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsTargetingType(value)
	return nil
}

// NewSponsoredProductsTargetingTypeFromValue returns a pointer to a valid SponsoredProductsTargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsTargetingTypeFromValue(v string) (*SponsoredProductsTargetingType, error) {
	ev := SponsoredProductsTargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsTargetingType: valid values are %v", v, AllowedSponsoredProductsTargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsTargetingType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsTargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsTargetingType value
func (v SponsoredProductsTargetingType) Ptr() *SponsoredProductsTargetingType {
	return &v
}

type NullableSponsoredProductsTargetingType struct {
	value *SponsoredProductsTargetingType
	isSet bool
}

func (v NullableSponsoredProductsTargetingType) Get() *SponsoredProductsTargetingType {
	return v.value
}

func (v *NullableSponsoredProductsTargetingType) Set(val *SponsoredProductsTargetingType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingType(val *SponsoredProductsTargetingType) *NullableSponsoredProductsTargetingType {
	return &NullableSponsoredProductsTargetingType{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
