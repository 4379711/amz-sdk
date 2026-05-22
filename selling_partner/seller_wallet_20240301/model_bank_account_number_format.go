package seller_wallet_20240301

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BankAccountNumberFormat The bank account's format type.
type BankAccountNumberFormat string

// List of BankAccountNumberFormat
const (
	BANKACCOUNTNUMBERFORMAT_IBAN BankAccountNumberFormat = "IBAN"
	BANKACCOUNTNUMBERFORMAT_BBAN BankAccountNumberFormat = "BBAN"
)

// All allowed values of BankAccountNumberFormat enum
var AllowedBankAccountNumberFormatEnumValues = []BankAccountNumberFormat{
	BANKACCOUNTNUMBERFORMAT_IBAN,
	BANKACCOUNTNUMBERFORMAT_BBAN,
}

func (v *BankAccountNumberFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BankAccountNumberFormat(value)
	for _, existing := range AllowedBankAccountNumberFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BankAccountNumberFormat", value)
}

// NewBankAccountNumberFormatFromValue returns a pointer to a valid BankAccountNumberFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankAccountNumberFormatFromValue(v string) (*BankAccountNumberFormat, error) {
	ev := BankAccountNumberFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BankAccountNumberFormat: valid values are %v", v, AllowedBankAccountNumberFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankAccountNumberFormat) IsValid() bool {
	for _, existing := range AllowedBankAccountNumberFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankAccountNumberFormat value
func (v BankAccountNumberFormat) Ptr() *BankAccountNumberFormat {
	return &v
}

type NullableBankAccountNumberFormat struct {
	value *BankAccountNumberFormat
	isSet bool
}

func (v NullableBankAccountNumberFormat) Get() *BankAccountNumberFormat {
	return v.value
}

func (v *NullableBankAccountNumberFormat) Set(val *BankAccountNumberFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountNumberFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountNumberFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountNumberFormat(val *BankAccountNumberFormat) *NullableBankAccountNumberFormat {
	return &NullableBankAccountNumberFormat{value: val, isSet: true}
}

func (v NullableBankAccountNumberFormat) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBankAccountNumberFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
