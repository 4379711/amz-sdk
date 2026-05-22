package seller_wallet_20240301

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PaymentPreferencePaymentType The type of payment preference.
type PaymentPreferencePaymentType string

// List of PaymentPreferencePaymentType
const (
	PAYMENTPREFERENCEPAYMENTTYPE_PERCENTAGE PaymentPreferencePaymentType = "PERCENTAGE"
	PAYMENTPREFERENCEPAYMENTTYPE_AMOUNT     PaymentPreferencePaymentType = "AMOUNT"
)

// All allowed values of PaymentPreferencePaymentType enum
var AllowedPaymentPreferencePaymentTypeEnumValues = []PaymentPreferencePaymentType{
	PAYMENTPREFERENCEPAYMENTTYPE_PERCENTAGE,
	PAYMENTPREFERENCEPAYMENTTYPE_AMOUNT,
}

func (v *PaymentPreferencePaymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentPreferencePaymentType(value)
	for _, existing := range AllowedPaymentPreferencePaymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentPreferencePaymentType", value)
}

// NewPaymentPreferencePaymentTypeFromValue returns a pointer to a valid PaymentPreferencePaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentPreferencePaymentTypeFromValue(v string) (*PaymentPreferencePaymentType, error) {
	ev := PaymentPreferencePaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentPreferencePaymentType: valid values are %v", v, AllowedPaymentPreferencePaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentPreferencePaymentType) IsValid() bool {
	for _, existing := range AllowedPaymentPreferencePaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentPreferencePaymentType value
func (v PaymentPreferencePaymentType) Ptr() *PaymentPreferencePaymentType {
	return &v
}

type NullablePaymentPreferencePaymentType struct {
	value *PaymentPreferencePaymentType
	isSet bool
}

func (v NullablePaymentPreferencePaymentType) Get() *PaymentPreferencePaymentType {
	return v.value
}

func (v *NullablePaymentPreferencePaymentType) Set(val *PaymentPreferencePaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentPreferencePaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentPreferencePaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentPreferencePaymentType(val *PaymentPreferencePaymentType) *NullablePaymentPreferencePaymentType {
	return &NullablePaymentPreferencePaymentType{value: val, isSet: true}
}

func (v NullablePaymentPreferencePaymentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaymentPreferencePaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
