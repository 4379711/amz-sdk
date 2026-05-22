package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CurrencyCode The currency used for all monetary values for entities under this profile. Cannot be `null`. |Region|`countryCode`|Country Name|`currencyCode`| |-----|------|------|------| |NA|US|United States|USD| |NA|CA|Canada|CAD| |NA|MX|Mexico|MXN| |NA|BR|Brazil|BRL| |EU|UK|United Kingdom|GBP| |EU|DE|Germany|EUR| |EU|FR|France|EUR| |EU|ES|Spain|EUR| |EU|IT|Italy|EUR| |EU|NL|The Netherlands|EUR| |EU|SE|Sweden|SEK| |EU|PL|Poland|PLN| |EU|AE|United Arab Emirates|AED| |EU|TR|Turkey|TRY| |FE|JP|Japan|JPY| |FE|AU|Australia|AUD| |FE|SG|Singapore|SGD|
type CurrencyCode string

// List of CurrencyCode
const (
	CURRENCYCODE_USD CurrencyCode = "USD"
	CURRENCYCODE_CAD CurrencyCode = "CAD"
	CURRENCYCODE_GBP CurrencyCode = "GBP"
	CURRENCYCODE_EUR CurrencyCode = "EUR"
	CURRENCYCODE_CNY CurrencyCode = "CNY"
	CURRENCYCODE_JPY CurrencyCode = "JPY"
	CURRENCYCODE_INR CurrencyCode = "INR"
	CURRENCYCODE_MXN CurrencyCode = "MXN"
	CURRENCYCODE_AUD CurrencyCode = "AUD"
	CURRENCYCODE_AED CurrencyCode = "AED"
	CURRENCYCODE_BRL CurrencyCode = "BRL"
	CURRENCYCODE_SGD CurrencyCode = "SGD"
	CURRENCYCODE_SEK CurrencyCode = "SEK"
	CURRENCYCODE_TRY CurrencyCode = "TRY"
	CURRENCYCODE_SAR CurrencyCode = "SAR"
	CURRENCYCODE_EGP CurrencyCode = "EGP"
	CURRENCYCODE_PLN CurrencyCode = "PLN"
	CURRENCYCODE_CLP CurrencyCode = "CLP"
	CURRENCYCODE_COP CurrencyCode = "COP"
	CURRENCYCODE_NGN CurrencyCode = "NGN"
	CURRENCYCODE_ZAR CurrencyCode = "ZAR"
)

// All allowed values of CurrencyCode enum
var AllowedCurrencyCodeEnumValues = []CurrencyCode{
	"USD",
	"CAD",
	"GBP",
	"EUR",
	"CNY",
	"JPY",
	"INR",
	"MXN",
	"AUD",
	"AED",
	"BRL",
	"SGD",
	"SEK",
	"TRY",
	"SAR",
	"EGP",
	"PLN",
	"CLP",
	"COP",
	"NGN",
	"ZAR",
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
