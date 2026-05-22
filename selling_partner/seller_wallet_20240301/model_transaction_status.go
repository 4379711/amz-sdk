package seller_wallet_20240301

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// TransactionStatus The current status of the transaction.
type TransactionStatus string

// List of TransactionStatus
const (
	TRANSACTIONSTATUS_FAILED                 TransactionStatus = "FAILED"
	TRANSACTIONSTATUS_FAILED_CREDITS_APPLIED TransactionStatus = "FAILED_CREDITS_APPLIED"
	TRANSACTIONSTATUS_INITIATED              TransactionStatus = "INITIATED"
	TRANSACTIONSTATUS_IN_PROGRESS            TransactionStatus = "IN_PROGRESS"
	TRANSACTIONSTATUS_PAYEE_UNDER_REVIEW     TransactionStatus = "PAYEE_UNDER_REVIEW"
	TRANSACTIONSTATUS_SUCCESSFUL             TransactionStatus = "SUCCESSFUL"
)

// All allowed values of TransactionStatus enum
var AllowedTransactionStatusEnumValues = []TransactionStatus{
	TRANSACTIONSTATUS_FAILED,
	TRANSACTIONSTATUS_FAILED_CREDITS_APPLIED,
	TRANSACTIONSTATUS_INITIATED,
	TRANSACTIONSTATUS_IN_PROGRESS,
	TRANSACTIONSTATUS_PAYEE_UNDER_REVIEW,
	TRANSACTIONSTATUS_SUCCESSFUL,
}

func (v *TransactionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionStatus(value)
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionStatus", value)
}

// NewTransactionStatusFromValue returns a pointer to a valid TransactionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionStatusFromValue(v string) (*TransactionStatus, error) {
	ev := TransactionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionStatus: valid values are %v", v, AllowedTransactionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionStatus) IsValid() bool {
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionStatus value
func (v TransactionStatus) Ptr() *TransactionStatus {
	return &v
}

type NullableTransactionStatus struct {
	value *TransactionStatus
	isSet bool
}

func (v NullableTransactionStatus) Get() *TransactionStatus {
	return v.value
}

func (v *NullableTransactionStatus) Set(val *TransactionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStatus(val *TransactionStatus) *NullableTransactionStatus {
	return &NullableTransactionStatus{value: val, isSet: true}
}

func (v NullableTransactionStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
