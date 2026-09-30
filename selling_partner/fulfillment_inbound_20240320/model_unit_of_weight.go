package fulfillment_inbound_20240320

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// UnitOfWeight Unit of the weight being measured.
type UnitOfWeight string

// List of UnitOfWeight
const (
	UNITOFWEIGHT_LB UnitOfWeight = "LB"
	UNITOFWEIGHT_KG UnitOfWeight = "KG"
)

// All allowed values of UnitOfWeight enum
var AllowedUnitOfWeightEnumValues = []UnitOfWeight{
	UNITOFWEIGHT_LB,
	UNITOFWEIGHT_KG,
}

func (v *UnitOfWeight) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = UnitOfWeight(value)
	return nil
}

// NewUnitOfWeightFromValue returns a pointer to a valid UnitOfWeight
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnitOfWeightFromValue(v string) (*UnitOfWeight, error) {
	ev := UnitOfWeight(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnitOfWeight: valid values are %v", v, AllowedUnitOfWeightEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnitOfWeight) IsValid() bool {
	for _, existing := range AllowedUnitOfWeightEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnitOfWeight value
func (v UnitOfWeight) Ptr() *UnitOfWeight {
	return &v
}

type NullableUnitOfWeight struct {
	value *UnitOfWeight
	isSet bool
}

func (v NullableUnitOfWeight) Get() *UnitOfWeight {
	return v.value
}

func (v *NullableUnitOfWeight) Set(val *UnitOfWeight) {
	v.value = val
	v.isSet = true
}

func (v NullableUnitOfWeight) IsSet() bool {
	return v.isSet
}

func (v *NullableUnitOfWeight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnitOfWeight(val *UnitOfWeight) *NullableUnitOfWeight {
	return &NullableUnitOfWeight{value: val, isSet: true}
}

func (v NullableUnitOfWeight) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUnitOfWeight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
