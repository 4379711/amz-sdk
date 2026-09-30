package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsCurrencyErrorReason the model 'SponsoredProductsCurrencyErrorReason'
type SponsoredProductsCurrencyErrorReason string

// List of SponsoredProductsCurrencyErrorReason
const (
	SPONSOREDPRODUCTSCURRENCYERRORREASON_PREFERRED_CURRENCY_NOT_SET               SponsoredProductsCurrencyErrorReason = "PREFERRED_CURRENCY_NOT_SET"
	SPONSOREDPRODUCTSCURRENCYERRORREASON_CURRENCY_NOT_SUPPORTED                   SponsoredProductsCurrencyErrorReason = "CURRENCY_NOT_SUPPORTED"
	SPONSOREDPRODUCTSCURRENCYERRORREASON_CANNOT_UPDATE_CURRENCY                   SponsoredProductsCurrencyErrorReason = "CANNOT_UPDATE_CURRENCY"
	SPONSOREDPRODUCTSCURRENCYERRORREASON_CURRENCY_NOT_MATCHING_PREFERRED_CURRENCY SponsoredProductsCurrencyErrorReason = "CURRENCY_NOT_MATCHING_PREFERRED_CURRENCY"
)

// All allowed values of SponsoredProductsCurrencyErrorReason enum
var AllowedSponsoredProductsCurrencyErrorReasonEnumValues = []SponsoredProductsCurrencyErrorReason{
	"PREFERRED_CURRENCY_NOT_SET",
	"CURRENCY_NOT_SUPPORTED",
	"CANNOT_UPDATE_CURRENCY",
	"CURRENCY_NOT_MATCHING_PREFERRED_CURRENCY",
}

func (v *SponsoredProductsCurrencyErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsCurrencyErrorReason(value)
	return nil
}

// NewSponsoredProductsCurrencyErrorReasonFromValue returns a pointer to a valid SponsoredProductsCurrencyErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCurrencyErrorReasonFromValue(v string) (*SponsoredProductsCurrencyErrorReason, error) {
	ev := SponsoredProductsCurrencyErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCurrencyErrorReason: valid values are %v", v, AllowedSponsoredProductsCurrencyErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCurrencyErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCurrencyErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCurrencyErrorReason value
func (v SponsoredProductsCurrencyErrorReason) Ptr() *SponsoredProductsCurrencyErrorReason {
	return &v
}

type NullableSponsoredProductsCurrencyErrorReason struct {
	value *SponsoredProductsCurrencyErrorReason
	isSet bool
}

func (v NullableSponsoredProductsCurrencyErrorReason) Get() *SponsoredProductsCurrencyErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsCurrencyErrorReason) Set(val *SponsoredProductsCurrencyErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCurrencyErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCurrencyErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCurrencyErrorReason(val *SponsoredProductsCurrencyErrorReason) *NullableSponsoredProductsCurrencyErrorReason {
	return &NullableSponsoredProductsCurrencyErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsCurrencyErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCurrencyErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
