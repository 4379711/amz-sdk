package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsAccessDeniedErrorCode the model 'SponsoredProductsAccessDeniedErrorCode'
type SponsoredProductsAccessDeniedErrorCode string

// List of SponsoredProductsAccessDeniedErrorCode
const (
	SPONSOREDPRODUCTSACCESSDENIEDERRORCODE_ACCESS_DENIED SponsoredProductsAccessDeniedErrorCode = "ACCESS_DENIED"
)

// All allowed values of SponsoredProductsAccessDeniedErrorCode enum
var AllowedSponsoredProductsAccessDeniedErrorCodeEnumValues = []SponsoredProductsAccessDeniedErrorCode{
	"ACCESS_DENIED",
}

func (v *SponsoredProductsAccessDeniedErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsAccessDeniedErrorCode(value)
	return nil
}

// NewSponsoredProductsAccessDeniedErrorCodeFromValue returns a pointer to a valid SponsoredProductsAccessDeniedErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAccessDeniedErrorCodeFromValue(v string) (*SponsoredProductsAccessDeniedErrorCode, error) {
	ev := SponsoredProductsAccessDeniedErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAccessDeniedErrorCode: valid values are %v", v, AllowedSponsoredProductsAccessDeniedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAccessDeniedErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAccessDeniedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAccessDeniedErrorCode value
func (v SponsoredProductsAccessDeniedErrorCode) Ptr() *SponsoredProductsAccessDeniedErrorCode {
	return &v
}

type NullableSponsoredProductsAccessDeniedErrorCode struct {
	value *SponsoredProductsAccessDeniedErrorCode
	isSet bool
}

func (v NullableSponsoredProductsAccessDeniedErrorCode) Get() *SponsoredProductsAccessDeniedErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsAccessDeniedErrorCode) Set(val *SponsoredProductsAccessDeniedErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAccessDeniedErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAccessDeniedErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAccessDeniedErrorCode(val *SponsoredProductsAccessDeniedErrorCode) *NullableSponsoredProductsAccessDeniedErrorCode {
	return &NullableSponsoredProductsAccessDeniedErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsAccessDeniedErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAccessDeniedErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
