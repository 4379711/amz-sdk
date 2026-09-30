package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsInternalServerErrorReason the model 'SponsoredProductsInternalServerErrorReason'
type SponsoredProductsInternalServerErrorReason string

// List of SponsoredProductsInternalServerErrorReason
const (
	SPONSOREDPRODUCTSINTERNALSERVERERRORREASON_INTERNAL_ERROR SponsoredProductsInternalServerErrorReason = "INTERNAL_ERROR"
)

// All allowed values of SponsoredProductsInternalServerErrorReason enum
var AllowedSponsoredProductsInternalServerErrorReasonEnumValues = []SponsoredProductsInternalServerErrorReason{
	"INTERNAL_ERROR",
}

func (v *SponsoredProductsInternalServerErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsInternalServerErrorReason(value)
	return nil
}

// NewSponsoredProductsInternalServerErrorReasonFromValue returns a pointer to a valid SponsoredProductsInternalServerErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsInternalServerErrorReasonFromValue(v string) (*SponsoredProductsInternalServerErrorReason, error) {
	ev := SponsoredProductsInternalServerErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsInternalServerErrorReason: valid values are %v", v, AllowedSponsoredProductsInternalServerErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsInternalServerErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsInternalServerErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsInternalServerErrorReason value
func (v SponsoredProductsInternalServerErrorReason) Ptr() *SponsoredProductsInternalServerErrorReason {
	return &v
}

type NullableSponsoredProductsInternalServerErrorReason struct {
	value *SponsoredProductsInternalServerErrorReason
	isSet bool
}

func (v NullableSponsoredProductsInternalServerErrorReason) Get() *SponsoredProductsInternalServerErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsInternalServerErrorReason) Set(val *SponsoredProductsInternalServerErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInternalServerErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInternalServerErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInternalServerErrorReason(val *SponsoredProductsInternalServerErrorReason) *NullableSponsoredProductsInternalServerErrorReason {
	return &NullableSponsoredProductsInternalServerErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsInternalServerErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInternalServerErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
