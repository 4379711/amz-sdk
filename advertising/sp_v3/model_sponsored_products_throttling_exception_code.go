package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsThrottlingExceptionCode the model 'SponsoredProductsThrottlingExceptionCode'
type SponsoredProductsThrottlingExceptionCode string

// List of SponsoredProductsThrottlingExceptionCode
const (
	SPONSOREDPRODUCTSTHROTTLINGEXCEPTIONCODE_THROTTLED SponsoredProductsThrottlingExceptionCode = "THROTTLED"
)

// All allowed values of SponsoredProductsThrottlingExceptionCode enum
var AllowedSponsoredProductsThrottlingExceptionCodeEnumValues = []SponsoredProductsThrottlingExceptionCode{
	"THROTTLED",
}

func (v *SponsoredProductsThrottlingExceptionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsThrottlingExceptionCode(value)
	return nil
}

// NewSponsoredProductsThrottlingExceptionCodeFromValue returns a pointer to a valid SponsoredProductsThrottlingExceptionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsThrottlingExceptionCodeFromValue(v string) (*SponsoredProductsThrottlingExceptionCode, error) {
	ev := SponsoredProductsThrottlingExceptionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsThrottlingExceptionCode: valid values are %v", v, AllowedSponsoredProductsThrottlingExceptionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsThrottlingExceptionCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsThrottlingExceptionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsThrottlingExceptionCode value
func (v SponsoredProductsThrottlingExceptionCode) Ptr() *SponsoredProductsThrottlingExceptionCode {
	return &v
}

type NullableSponsoredProductsThrottlingExceptionCode struct {
	value *SponsoredProductsThrottlingExceptionCode
	isSet bool
}

func (v NullableSponsoredProductsThrottlingExceptionCode) Get() *SponsoredProductsThrottlingExceptionCode {
	return v.value
}

func (v *NullableSponsoredProductsThrottlingExceptionCode) Set(val *SponsoredProductsThrottlingExceptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsThrottlingExceptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsThrottlingExceptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsThrottlingExceptionCode(val *SponsoredProductsThrottlingExceptionCode) *NullableSponsoredProductsThrottlingExceptionCode {
	return &NullableSponsoredProductsThrottlingExceptionCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsThrottlingExceptionCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsThrottlingExceptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
