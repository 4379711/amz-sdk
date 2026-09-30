package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsThrottledErrorReason the model 'SponsoredProductsThrottledErrorReason'
type SponsoredProductsThrottledErrorReason string

// List of SponsoredProductsThrottledErrorReason
const (
	SPONSOREDPRODUCTSTHROTTLEDERRORREASON_THROTTLED SponsoredProductsThrottledErrorReason = "THROTTLED"
)

// All allowed values of SponsoredProductsThrottledErrorReason enum
var AllowedSponsoredProductsThrottledErrorReasonEnumValues = []SponsoredProductsThrottledErrorReason{
	"THROTTLED",
}

func (v *SponsoredProductsThrottledErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsThrottledErrorReason(value)
	return nil
}

// NewSponsoredProductsThrottledErrorReasonFromValue returns a pointer to a valid SponsoredProductsThrottledErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsThrottledErrorReasonFromValue(v string) (*SponsoredProductsThrottledErrorReason, error) {
	ev := SponsoredProductsThrottledErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsThrottledErrorReason: valid values are %v", v, AllowedSponsoredProductsThrottledErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsThrottledErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsThrottledErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsThrottledErrorReason value
func (v SponsoredProductsThrottledErrorReason) Ptr() *SponsoredProductsThrottledErrorReason {
	return &v
}

type NullableSponsoredProductsThrottledErrorReason struct {
	value *SponsoredProductsThrottledErrorReason
	isSet bool
}

func (v NullableSponsoredProductsThrottledErrorReason) Get() *SponsoredProductsThrottledErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsThrottledErrorReason) Set(val *SponsoredProductsThrottledErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsThrottledErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsThrottledErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsThrottledErrorReason(val *SponsoredProductsThrottledErrorReason) *NullableSponsoredProductsThrottledErrorReason {
	return &NullableSponsoredProductsThrottledErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsThrottledErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsThrottledErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
