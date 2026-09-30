package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AccessDeniedErrorCode the model 'AccessDeniedErrorCode'
type AccessDeniedErrorCode string

// List of AccessDeniedErrorCode
const (
	ACCESSDENIEDERRORCODE_ACCESS_DENIED AccessDeniedErrorCode = "ACCESS_DENIED"
)

// All allowed values of AccessDeniedErrorCode enum
var AllowedAccessDeniedErrorCodeEnumValues = []AccessDeniedErrorCode{
	"ACCESS_DENIED",
}

func (v *AccessDeniedErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = AccessDeniedErrorCode(value)
	return nil
}

// NewAccessDeniedErrorCodeFromValue returns a pointer to a valid AccessDeniedErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessDeniedErrorCodeFromValue(v string) (*AccessDeniedErrorCode, error) {
	ev := AccessDeniedErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessDeniedErrorCode: valid values are %v", v, AllowedAccessDeniedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessDeniedErrorCode) IsValid() bool {
	for _, existing := range AllowedAccessDeniedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessDeniedErrorCode value
func (v AccessDeniedErrorCode) Ptr() *AccessDeniedErrorCode {
	return &v
}

type NullableAccessDeniedErrorCode struct {
	value *AccessDeniedErrorCode
	isSet bool
}

func (v NullableAccessDeniedErrorCode) Get() *AccessDeniedErrorCode {
	return v.value
}

func (v *NullableAccessDeniedErrorCode) Set(val *AccessDeniedErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDeniedErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDeniedErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDeniedErrorCode(val *AccessDeniedErrorCode) *NullableAccessDeniedErrorCode {
	return &NullableAccessDeniedErrorCode{value: val, isSet: true}
}

func (v NullableAccessDeniedErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessDeniedErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
