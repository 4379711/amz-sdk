package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PaymentType Payment type of the purchase.
type PaymentType string

// List of PaymentType
const (
	PAYMENTTYPE_PAY_THROUGH_AMAZON    PaymentType = "PAY_THROUGH_AMAZON"
	PAYMENTTYPE_PAY_DIRECT_TO_CARRIER PaymentType = "PAY_DIRECT_TO_CARRIER"
)

// All allowed values of PaymentType enum
var AllowedPaymentTypeEnumValues = []PaymentType{
	PAYMENTTYPE_PAY_THROUGH_AMAZON,
	PAYMENTTYPE_PAY_DIRECT_TO_CARRIER,
}

func (v *PaymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = PaymentType(value)
	return nil
}

// NewPaymentTypeFromValue returns a pointer to a valid PaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentTypeFromValue(v string) (*PaymentType, error) {
	ev := PaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentType: valid values are %v", v, AllowedPaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentType) IsValid() bool {
	for _, existing := range AllowedPaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentType value
func (v PaymentType) Ptr() *PaymentType {
	return &v
}

type NullablePaymentType struct {
	value *PaymentType
	isSet bool
}

func (v NullablePaymentType) Get() *PaymentType {
	return v.value
}

func (v *NullablePaymentType) Set(val *PaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentType(val *PaymentType) *NullablePaymentType {
	return &NullablePaymentType{value: val, isSet: true}
}

func (v NullablePaymentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
