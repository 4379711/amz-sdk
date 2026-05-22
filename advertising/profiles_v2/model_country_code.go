package profiles_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CountryCode The countryCode for a given country |Region|`countryCode`|Country Name| |------|-----|-------| |NA|BR|Brazil| |NA|CA|Canada| |NA|MX|Mexico| |NA|US|United States| |EU|AE|United Arab Emirates| |EU|BE|Belgium|  |EU|DE|Germany| |EU|EG|Egypt| |EU|ES|Spain| |EU|FR|France| |EU|IN|India| |EU|IT|Italy| |EU|NL|The Netherlands| |EU|PL|Poland| |EU|SA|Saudi Arabia|  |EU|SE|Sweden|   |EU|TR|Turkey| |EU|UK|United Kingdom| |EU|ZA| South Africa | |FE|AU|Australia|     |FE|JP|Japan| |FE|SG|Singapore|
type CountryCode string

// List of countryCode
const (
	COUNTRYCODE_BR CountryCode = "BR"
	COUNTRYCODE_CA CountryCode = "CA"
	COUNTRYCODE_MX CountryCode = "MX"
	COUNTRYCODE_US CountryCode = "US"
	COUNTRYCODE_AE CountryCode = "AE"
	COUNTRYCODE_BE CountryCode = "BE"
	COUNTRYCODE_DE CountryCode = "DE"
	COUNTRYCODE_EG CountryCode = "EG"
	COUNTRYCODE_ES CountryCode = "ES"
	COUNTRYCODE_FR CountryCode = "FR"
	COUNTRYCODE_IN CountryCode = "IN"
	COUNTRYCODE_IT CountryCode = "IT"
	COUNTRYCODE_NL CountryCode = "NL"
	COUNTRYCODE_PL CountryCode = "PL"
	COUNTRYCODE_SA CountryCode = "SA"
	COUNTRYCODE_SE CountryCode = "SE"
	COUNTRYCODE_TR CountryCode = "TR"
	COUNTRYCODE_UK CountryCode = "UK"
	COUNTRYCODE_AU CountryCode = "AU"
	COUNTRYCODE_JP CountryCode = "JP"
	COUNTRYCODE_SG CountryCode = "SG"
	COUNTRYCODE_ZA CountryCode = "ZA"
)

// All allowed values of CountryCode enum
var AllowedCountryCodeEnumValues = []CountryCode{
	"BR",
	"CA",
	"MX",
	"US",
	"AE",
	"BE",
	"DE",
	"EG",
	"ES",
	"FR",
	"IN",
	"IT",
	"NL",
	"PL",
	"SA",
	"SE",
	"TR",
	"UK",
	"AU",
	"JP",
	"SG",
	"ZA",
}

func (v *CountryCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CountryCode(value)
	for _, existing := range AllowedCountryCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CountryCode", value)
}

// NewCountryCodeFromValue returns a pointer to a valid CountryCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountryCodeFromValue(v string) (*CountryCode, error) {
	ev := CountryCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CountryCode: valid values are %v", v, AllowedCountryCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountryCode) IsValid() bool {
	for _, existing := range AllowedCountryCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to countryCode value
func (v CountryCode) Ptr() *CountryCode {
	return &v
}

type NullableCountryCode struct {
	value *CountryCode
	isSet bool
}

func (v NullableCountryCode) Get() *CountryCode {
	return v.value
}

func (v *NullableCountryCode) Set(val *CountryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryCode(val *CountryCode) *NullableCountryCode {
	return &NullableCountryCode{value: val, isSet: true}
}

func (v NullableCountryCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCountryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
