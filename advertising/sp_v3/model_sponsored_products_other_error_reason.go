package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsOtherErrorReason the model 'SponsoredProductsOtherErrorReason'
type SponsoredProductsOtherErrorReason string

// List of SponsoredProductsOtherErrorReason
const (
	SPONSOREDPRODUCTSOTHERERRORREASON_OTHER_ERROR SponsoredProductsOtherErrorReason = "OTHER_ERROR"
)

// All allowed values of SponsoredProductsOtherErrorReason enum
var AllowedSponsoredProductsOtherErrorReasonEnumValues = []SponsoredProductsOtherErrorReason{
	"OTHER_ERROR",
}

func (v *SponsoredProductsOtherErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsOtherErrorReason(value)
	return nil
}

// NewSponsoredProductsOtherErrorReasonFromValue returns a pointer to a valid SponsoredProductsOtherErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsOtherErrorReasonFromValue(v string) (*SponsoredProductsOtherErrorReason, error) {
	ev := SponsoredProductsOtherErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsOtherErrorReason: valid values are %v", v, AllowedSponsoredProductsOtherErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsOtherErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsOtherErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsOtherErrorReason value
func (v SponsoredProductsOtherErrorReason) Ptr() *SponsoredProductsOtherErrorReason {
	return &v
}

type NullableSponsoredProductsOtherErrorReason struct {
	value *SponsoredProductsOtherErrorReason
	isSet bool
}

func (v NullableSponsoredProductsOtherErrorReason) Get() *SponsoredProductsOtherErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsOtherErrorReason) Set(val *SponsoredProductsOtherErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsOtherErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsOtherErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsOtherErrorReason(val *SponsoredProductsOtherErrorReason) *NullableSponsoredProductsOtherErrorReason {
	return &NullableSponsoredProductsOtherErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsOtherErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsOtherErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
