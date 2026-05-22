package fulfillment_inbound_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CurrencyCode The currency code.
type CurrencyCode string

// List of CurrencyCode
const (
	CURRENCYCODE_USD CurrencyCode = "USD"
	CURRENCYCODE_GBP CurrencyCode = "GBP"
)

// All allowed values of CurrencyCode enum
var AllowedCurrencyCodeEnumValues = []CurrencyCode{
	CURRENCYCODE_USD,
	CURRENCYCODE_GBP,
}

func (v *CurrencyCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CurrencyCode(value)
	for _, existing := range AllowedCurrencyCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CurrencyCode", value)
}

// NewCurrencyCodeFromValue returns a pointer to a valid CurrencyCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCurrencyCodeFromValue(v string) (*CurrencyCode, error) {
	ev := CurrencyCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CurrencyCode: valid values are %v", v, AllowedCurrencyCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CurrencyCode) IsValid() bool {
	for _, existing := range AllowedCurrencyCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CurrencyCode value
func (v CurrencyCode) Ptr() *CurrencyCode {
	return &v
}

type NullableCurrencyCode struct {
	value *CurrencyCode
	isSet bool
}

func (v NullableCurrencyCode) Get() *CurrencyCode {
	return v.value
}

func (v *NullableCurrencyCode) Set(val *CurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyCode(val *CurrencyCode) *NullableCurrencyCode {
	return &NullableCurrencyCode{value: val, isSet: true}
}

func (v NullableCurrencyCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
