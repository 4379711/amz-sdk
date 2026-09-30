package shipment_invoicing_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AddressTypeEnum The shipping address type.
type AddressTypeEnum string

// List of AddressTypeEnum
const (
	ADDRESSTYPEENUM_RESIDENTIAL AddressTypeEnum = "Residential"
	ADDRESSTYPEENUM_COMMERCIAL  AddressTypeEnum = "Commercial"
)

// All allowed values of AddressTypeEnum enum
var AllowedAddressTypeEnumEnumValues = []AddressTypeEnum{
	ADDRESSTYPEENUM_RESIDENTIAL,
	ADDRESSTYPEENUM_COMMERCIAL,
}

func (v *AddressTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = AddressTypeEnum(value)
	return nil
}

// NewAddressTypeEnumFromValue returns a pointer to a valid AddressTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAddressTypeEnumFromValue(v string) (*AddressTypeEnum, error) {
	ev := AddressTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AddressTypeEnum: valid values are %v", v, AllowedAddressTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AddressTypeEnum) IsValid() bool {
	for _, existing := range AllowedAddressTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AddressTypeEnum value
func (v AddressTypeEnum) Ptr() *AddressTypeEnum {
	return &v
}

type NullableAddressTypeEnum struct {
	value *AddressTypeEnum
	isSet bool
}

func (v NullableAddressTypeEnum) Get() *AddressTypeEnum {
	return v.value
}

func (v *NullableAddressTypeEnum) Set(val *AddressTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTypeEnum(val *AddressTypeEnum) *NullableAddressTypeEnum {
	return &NullableAddressTypeEnum{value: val, isSet: true}
}

func (v NullableAddressTypeEnum) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAddressTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
