package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsThrottledErrorCode the model 'SponsoredProductsThrottledErrorCode'
type SponsoredProductsThrottledErrorCode string

// List of SponsoredProductsThrottledErrorCode
const (
	SPONSOREDPRODUCTSTHROTTLEDERRORCODE_THROTTLED SponsoredProductsThrottledErrorCode = "THROTTLED"
)

// All allowed values of SponsoredProductsThrottledErrorCode enum
var AllowedSponsoredProductsThrottledErrorCodeEnumValues = []SponsoredProductsThrottledErrorCode{
	"THROTTLED",
}

func (v *SponsoredProductsThrottledErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsThrottledErrorCode(value)
	return nil
}

// NewSponsoredProductsThrottledErrorCodeFromValue returns a pointer to a valid SponsoredProductsThrottledErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsThrottledErrorCodeFromValue(v string) (*SponsoredProductsThrottledErrorCode, error) {
	ev := SponsoredProductsThrottledErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsThrottledErrorCode: valid values are %v", v, AllowedSponsoredProductsThrottledErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsThrottledErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsThrottledErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsThrottledErrorCode value
func (v SponsoredProductsThrottledErrorCode) Ptr() *SponsoredProductsThrottledErrorCode {
	return &v
}

type NullableSponsoredProductsThrottledErrorCode struct {
	value *SponsoredProductsThrottledErrorCode
	isSet bool
}

func (v NullableSponsoredProductsThrottledErrorCode) Get() *SponsoredProductsThrottledErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsThrottledErrorCode) Set(val *SponsoredProductsThrottledErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsThrottledErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsThrottledErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsThrottledErrorCode(val *SponsoredProductsThrottledErrorCode) *NullableSponsoredProductsThrottledErrorCode {
	return &NullableSponsoredProductsThrottledErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsThrottledErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsThrottledErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
