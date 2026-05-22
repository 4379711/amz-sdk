package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsMarketplace the model 'SponsoredProductsMarketplace'
type SponsoredProductsMarketplace string

// List of SponsoredProductsMarketplace
const (
	SPONSOREDPRODUCTSMARKETPLACE_US SponsoredProductsMarketplace = "US"
	SPONSOREDPRODUCTSMARKETPLACE_CA SponsoredProductsMarketplace = "CA"
	SPONSOREDPRODUCTSMARKETPLACE_MX SponsoredProductsMarketplace = "MX"
	SPONSOREDPRODUCTSMARKETPLACE_BR SponsoredProductsMarketplace = "BR"
	SPONSOREDPRODUCTSMARKETPLACE_UK SponsoredProductsMarketplace = "UK"
	SPONSOREDPRODUCTSMARKETPLACE_DE SponsoredProductsMarketplace = "DE"
	SPONSOREDPRODUCTSMARKETPLACE_FR SponsoredProductsMarketplace = "FR"
	SPONSOREDPRODUCTSMARKETPLACE_ES SponsoredProductsMarketplace = "ES"
	SPONSOREDPRODUCTSMARKETPLACE_IT SponsoredProductsMarketplace = "IT"
	SPONSOREDPRODUCTSMARKETPLACE_IN SponsoredProductsMarketplace = "IN"
	SPONSOREDPRODUCTSMARKETPLACE_AE SponsoredProductsMarketplace = "AE"
	SPONSOREDPRODUCTSMARKETPLACE_SA SponsoredProductsMarketplace = "SA"
	SPONSOREDPRODUCTSMARKETPLACE_NL SponsoredProductsMarketplace = "NL"
	SPONSOREDPRODUCTSMARKETPLACE_PL SponsoredProductsMarketplace = "PL"
	SPONSOREDPRODUCTSMARKETPLACE_SE SponsoredProductsMarketplace = "SE"
	SPONSOREDPRODUCTSMARKETPLACE_TR SponsoredProductsMarketplace = "TR"
	SPONSOREDPRODUCTSMARKETPLACE_EG SponsoredProductsMarketplace = "EG"
	SPONSOREDPRODUCTSMARKETPLACE_JP SponsoredProductsMarketplace = "JP"
	SPONSOREDPRODUCTSMARKETPLACE_AU SponsoredProductsMarketplace = "AU"
	SPONSOREDPRODUCTSMARKETPLACE_SG SponsoredProductsMarketplace = "SG"
)

// All allowed values of SponsoredProductsMarketplace enum
var AllowedSponsoredProductsMarketplaceEnumValues = []SponsoredProductsMarketplace{
	"US",
	"CA",
	"MX",
	"BR",
	"UK",
	"DE",
	"FR",
	"ES",
	"IT",
	"IN",
	"AE",
	"SA",
	"NL",
	"PL",
	"SE",
	"TR",
	"EG",
	"JP",
	"AU",
	"SG",
}

func (v *SponsoredProductsMarketplace) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsMarketplace(value)
	for _, existing := range AllowedSponsoredProductsMarketplaceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsMarketplace", value)
}

// NewSponsoredProductsMarketplaceFromValue returns a pointer to a valid SponsoredProductsMarketplace
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsMarketplaceFromValue(v string) (*SponsoredProductsMarketplace, error) {
	ev := SponsoredProductsMarketplace(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsMarketplace: valid values are %v", v, AllowedSponsoredProductsMarketplaceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsMarketplace) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsMarketplaceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsMarketplace value
func (v SponsoredProductsMarketplace) Ptr() *SponsoredProductsMarketplace {
	return &v
}

type NullableSponsoredProductsMarketplace struct {
	value *SponsoredProductsMarketplace
	isSet bool
}

func (v NullableSponsoredProductsMarketplace) Get() *SponsoredProductsMarketplace {
	return v.value
}

func (v *NullableSponsoredProductsMarketplace) Set(val *SponsoredProductsMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplace(val *SponsoredProductsMarketplace) *NullableSponsoredProductsMarketplace {
	return &NullableSponsoredProductsMarketplace{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
