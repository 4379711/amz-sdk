package notifications

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OrderChangeTypeEnum The supported order change type of ORDER_CHANGE notification.
type OrderChangeTypeEnum string

// List of OrderChangeTypeEnum
const (
	ORDERCHANGETYPEENUM_ORDER_STATUS_CHANGE    OrderChangeTypeEnum = "OrderStatusChange"
	ORDERCHANGETYPEENUM_BUYER_REQUESTED_CHANGE OrderChangeTypeEnum = "BuyerRequestedChange"
)

// All allowed values of OrderChangeTypeEnum enum
var AllowedOrderChangeTypeEnumEnumValues = []OrderChangeTypeEnum{
	ORDERCHANGETYPEENUM_ORDER_STATUS_CHANGE,
	ORDERCHANGETYPEENUM_BUYER_REQUESTED_CHANGE,
}

func (v *OrderChangeTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderChangeTypeEnum(value)
	for _, existing := range AllowedOrderChangeTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderChangeTypeEnum", value)
}

// NewOrderChangeTypeEnumFromValue returns a pointer to a valid OrderChangeTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderChangeTypeEnumFromValue(v string) (*OrderChangeTypeEnum, error) {
	ev := OrderChangeTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderChangeTypeEnum: valid values are %v", v, AllowedOrderChangeTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderChangeTypeEnum) IsValid() bool {
	for _, existing := range AllowedOrderChangeTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderChangeTypeEnum value
func (v OrderChangeTypeEnum) Ptr() *OrderChangeTypeEnum {
	return &v
}

type NullableOrderChangeTypeEnum struct {
	value *OrderChangeTypeEnum
	isSet bool
}

func (v NullableOrderChangeTypeEnum) Get() *OrderChangeTypeEnum {
	return v.value
}

func (v *NullableOrderChangeTypeEnum) Set(val *OrderChangeTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderChangeTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderChangeTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderChangeTypeEnum(val *OrderChangeTypeEnum) *NullableOrderChangeTypeEnum {
	return &NullableOrderChangeTypeEnum{value: val, isSet: true}
}

func (v NullableOrderChangeTypeEnum) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderChangeTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
