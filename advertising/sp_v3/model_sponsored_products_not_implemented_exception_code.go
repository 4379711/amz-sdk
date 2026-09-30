package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsNotImplementedExceptionCode the model 'SponsoredProductsNotImplementedExceptionCode'
type SponsoredProductsNotImplementedExceptionCode string

// List of SponsoredProductsNotImplementedExceptionCode
const (
	SPONSOREDPRODUCTSNOTIMPLEMENTEDEXCEPTIONCODE_NOT_IMPLEMENTED SponsoredProductsNotImplementedExceptionCode = "NOT_IMPLEMENTED"
)

// All allowed values of SponsoredProductsNotImplementedExceptionCode enum
var AllowedSponsoredProductsNotImplementedExceptionCodeEnumValues = []SponsoredProductsNotImplementedExceptionCode{
	"NOT_IMPLEMENTED",
}

func (v *SponsoredProductsNotImplementedExceptionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsNotImplementedExceptionCode(value)
	return nil
}

// NewSponsoredProductsNotImplementedExceptionCodeFromValue returns a pointer to a valid SponsoredProductsNotImplementedExceptionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsNotImplementedExceptionCodeFromValue(v string) (*SponsoredProductsNotImplementedExceptionCode, error) {
	ev := SponsoredProductsNotImplementedExceptionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsNotImplementedExceptionCode: valid values are %v", v, AllowedSponsoredProductsNotImplementedExceptionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsNotImplementedExceptionCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsNotImplementedExceptionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsNotImplementedExceptionCode value
func (v SponsoredProductsNotImplementedExceptionCode) Ptr() *SponsoredProductsNotImplementedExceptionCode {
	return &v
}

type NullableSponsoredProductsNotImplementedExceptionCode struct {
	value *SponsoredProductsNotImplementedExceptionCode
	isSet bool
}

func (v NullableSponsoredProductsNotImplementedExceptionCode) Get() *SponsoredProductsNotImplementedExceptionCode {
	return v.value
}

func (v *NullableSponsoredProductsNotImplementedExceptionCode) Set(val *SponsoredProductsNotImplementedExceptionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNotImplementedExceptionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNotImplementedExceptionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNotImplementedExceptionCode(val *SponsoredProductsNotImplementedExceptionCode) *NullableSponsoredProductsNotImplementedExceptionCode {
	return &NullableSponsoredProductsNotImplementedExceptionCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsNotImplementedExceptionCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNotImplementedExceptionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
