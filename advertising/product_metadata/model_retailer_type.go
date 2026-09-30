package product_metadata

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// RetailerType the model 'RetailerType'
type RetailerType string

// List of RetailerType
const (
	RETAILERTYPE_RMS RetailerType = "RMS"
)

// All allowed values of RetailerType enum
var AllowedRetailerTypeEnumValues = []RetailerType{
	"RMS",
}

func (v *RetailerType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = RetailerType(value)
	return nil
}

// NewRetailerTypeFromValue returns a pointer to a valid RetailerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRetailerTypeFromValue(v string) (*RetailerType, error) {
	ev := RetailerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RetailerType: valid values are %v", v, AllowedRetailerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RetailerType) IsValid() bool {
	for _, existing := range AllowedRetailerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetailerType value
func (v RetailerType) Ptr() *RetailerType {
	return &v
}

type NullableRetailerType struct {
	value *RetailerType
	isSet bool
}

func (v NullableRetailerType) Get() *RetailerType {
	return v.value
}

func (v *NullableRetailerType) Set(val *RetailerType) {
	v.value = val
	v.isSet = true
}

func (v NullableRetailerType) IsSet() bool {
	return v.isSet
}

func (v *NullableRetailerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetailerType(val *RetailerType) *NullableRetailerType {
	return &NullableRetailerType{value: val, isSet: true}
}

func (v NullableRetailerType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRetailerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
