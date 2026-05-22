package seller_wallet_20240301

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// BankNumberFormat The format of the bank number. Also known as the routing number type.
type BankNumberFormat string

// List of BankNumberFormat
const (
	BANKNUMBERFORMAT_BIC   BankNumberFormat = "BIC"
	BANKNUMBERFORMAT_BASIC BankNumberFormat = "BASIC"
)

// All allowed values of BankNumberFormat enum
var AllowedBankNumberFormatEnumValues = []BankNumberFormat{
	BANKNUMBERFORMAT_BIC,
	BANKNUMBERFORMAT_BASIC,
}

func (v *BankNumberFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BankNumberFormat(value)
	for _, existing := range AllowedBankNumberFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BankNumberFormat", value)
}

// NewBankNumberFormatFromValue returns a pointer to a valid BankNumberFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankNumberFormatFromValue(v string) (*BankNumberFormat, error) {
	ev := BankNumberFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BankNumberFormat: valid values are %v", v, AllowedBankNumberFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankNumberFormat) IsValid() bool {
	for _, existing := range AllowedBankNumberFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankNumberFormat value
func (v BankNumberFormat) Ptr() *BankNumberFormat {
	return &v
}

type NullableBankNumberFormat struct {
	value *BankNumberFormat
	isSet bool
}

func (v NullableBankNumberFormat) Get() *BankNumberFormat {
	return v.value
}

func (v *NullableBankNumberFormat) Set(val *BankNumberFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableBankNumberFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableBankNumberFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankNumberFormat(val *BankNumberFormat) *NullableBankNumberFormat {
	return &NullableBankNumberFormat{value: val, isSet: true}
}

func (v NullableBankNumberFormat) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBankNumberFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
