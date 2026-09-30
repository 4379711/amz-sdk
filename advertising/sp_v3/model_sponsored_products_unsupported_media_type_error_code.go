package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsUnsupportedMediaTypeErrorCode the model 'SponsoredProductsUnsupportedMediaTypeErrorCode'
type SponsoredProductsUnsupportedMediaTypeErrorCode string

// List of SponsoredProductsUnsupportedMediaTypeErrorCode
const (
	SPONSOREDPRODUCTSUNSUPPORTEDMEDIATYPEERRORCODE_UNSUPPORTED_MEDIA_TYPE SponsoredProductsUnsupportedMediaTypeErrorCode = "UNSUPPORTED_MEDIA_TYPE"
)

// All allowed values of SponsoredProductsUnsupportedMediaTypeErrorCode enum
var AllowedSponsoredProductsUnsupportedMediaTypeErrorCodeEnumValues = []SponsoredProductsUnsupportedMediaTypeErrorCode{
	"UNSUPPORTED_MEDIA_TYPE",
}

func (v *SponsoredProductsUnsupportedMediaTypeErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = SponsoredProductsUnsupportedMediaTypeErrorCode(value)
	return nil
}

// NewSponsoredProductsUnsupportedMediaTypeErrorCodeFromValue returns a pointer to a valid SponsoredProductsUnsupportedMediaTypeErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsUnsupportedMediaTypeErrorCodeFromValue(v string) (*SponsoredProductsUnsupportedMediaTypeErrorCode, error) {
	ev := SponsoredProductsUnsupportedMediaTypeErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsUnsupportedMediaTypeErrorCode: valid values are %v", v, AllowedSponsoredProductsUnsupportedMediaTypeErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsUnsupportedMediaTypeErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsUnsupportedMediaTypeErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsUnsupportedMediaTypeErrorCode value
func (v SponsoredProductsUnsupportedMediaTypeErrorCode) Ptr() *SponsoredProductsUnsupportedMediaTypeErrorCode {
	return &v
}

type NullableSponsoredProductsUnsupportedMediaTypeErrorCode struct {
	value *SponsoredProductsUnsupportedMediaTypeErrorCode
	isSet bool
}

func (v NullableSponsoredProductsUnsupportedMediaTypeErrorCode) Get() *SponsoredProductsUnsupportedMediaTypeErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsUnsupportedMediaTypeErrorCode) Set(val *SponsoredProductsUnsupportedMediaTypeErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUnsupportedMediaTypeErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUnsupportedMediaTypeErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUnsupportedMediaTypeErrorCode(val *SponsoredProductsUnsupportedMediaTypeErrorCode) *NullableSponsoredProductsUnsupportedMediaTypeErrorCode {
	return &NullableSponsoredProductsUnsupportedMediaTypeErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsUnsupportedMediaTypeErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUnsupportedMediaTypeErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
