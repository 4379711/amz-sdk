package seller_wallet_20240301

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// TransactionType The type of transaction.
type TransactionType string

// List of TransactionType
const (
	TRANSACTIONTYPE_CREDIT TransactionType = "CREDIT"
	TRANSACTIONTYPE_DEBIT  TransactionType = "DEBIT"
)

// All allowed values of TransactionType enum
var AllowedTransactionTypeEnumValues = []TransactionType{
	TRANSACTIONTYPE_CREDIT,
	TRANSACTIONTYPE_DEBIT,
}

func (v *TransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionType(value)
	for _, existing := range AllowedTransactionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionType", value)
}

// NewTransactionTypeFromValue returns a pointer to a valid TransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionTypeFromValue(v string) (*TransactionType, error) {
	ev := TransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionType: valid values are %v", v, AllowedTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionType) IsValid() bool {
	for _, existing := range AllowedTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionType value
func (v TransactionType) Ptr() *TransactionType {
	return &v
}

type NullableTransactionType struct {
	value *TransactionType
	isSet bool
}

func (v NullableTransactionType) Get() *TransactionType {
	return v.value
}

func (v *NullableTransactionType) Set(val *TransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionType(val *TransactionType) *NullableTransactionType {
	return &NullableTransactionType{value: val, isSet: true}
}

func (v NullableTransactionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
