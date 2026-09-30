package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ShipmentType Shipment type.
type ShipmentType string

// List of ShipmentType
const (
	SHIPMENTTYPE_FORWARD ShipmentType = "FORWARD"
	SHIPMENTTYPE_RETURNS ShipmentType = "RETURNS"
)

// All allowed values of ShipmentType enum
var AllowedShipmentTypeEnumValues = []ShipmentType{
	SHIPMENTTYPE_FORWARD,
	SHIPMENTTYPE_RETURNS,
}

func (v *ShipmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = ShipmentType(value)
	return nil
}

// NewShipmentTypeFromValue returns a pointer to a valid ShipmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentTypeFromValue(v string) (*ShipmentType, error) {
	ev := ShipmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentType: valid values are %v", v, AllowedShipmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentType) IsValid() bool {
	for _, existing := range AllowedShipmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipmentType value
func (v ShipmentType) Ptr() *ShipmentType {
	return &v
}

type NullableShipmentType struct {
	value *ShipmentType
	isSet bool
}

func (v NullableShipmentType) Get() *ShipmentType {
	return v.value
}

func (v *NullableShipmentType) Set(val *ShipmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentType(val *ShipmentType) *NullableShipmentType {
	return &NullableShipmentType{value: val, isSet: true}
}

func (v NullableShipmentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
