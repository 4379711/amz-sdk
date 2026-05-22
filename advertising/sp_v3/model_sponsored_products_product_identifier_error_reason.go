package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsProductIdentifierErrorReason the model 'SponsoredProductsProductIdentifierErrorReason'
type SponsoredProductsProductIdentifierErrorReason string

// List of SponsoredProductsProductIdentifierErrorReason
const (
	SPONSOREDPRODUCTSPRODUCTIDENTIFIERERRORREASON_SKU  SponsoredProductsProductIdentifierErrorReason = "INVALID_SKU"
	SPONSOREDPRODUCTSPRODUCTIDENTIFIERERRORREASON_ASIN SponsoredProductsProductIdentifierErrorReason = "INVALID_ASIN"
)

// All allowed values of SponsoredProductsProductIdentifierErrorReason enum
var AllowedSponsoredProductsProductIdentifierErrorReasonEnumValues = []SponsoredProductsProductIdentifierErrorReason{
	"INVALID_SKU",
	"INVALID_ASIN",
}

func (v *SponsoredProductsProductIdentifierErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsProductIdentifierErrorReason(value)
	for _, existing := range AllowedSponsoredProductsProductIdentifierErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsProductIdentifierErrorReason", value)
}

// NewSponsoredProductsProductIdentifierErrorReasonFromValue returns a pointer to a valid SponsoredProductsProductIdentifierErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsProductIdentifierErrorReasonFromValue(v string) (*SponsoredProductsProductIdentifierErrorReason, error) {
	ev := SponsoredProductsProductIdentifierErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsProductIdentifierErrorReason: valid values are %v", v, AllowedSponsoredProductsProductIdentifierErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsProductIdentifierErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsProductIdentifierErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsProductIdentifierErrorReason value
func (v SponsoredProductsProductIdentifierErrorReason) Ptr() *SponsoredProductsProductIdentifierErrorReason {
	return &v
}

type NullableSponsoredProductsProductIdentifierErrorReason struct {
	value *SponsoredProductsProductIdentifierErrorReason
	isSet bool
}

func (v NullableSponsoredProductsProductIdentifierErrorReason) Get() *SponsoredProductsProductIdentifierErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsProductIdentifierErrorReason) Set(val *SponsoredProductsProductIdentifierErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductIdentifierErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductIdentifierErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductIdentifierErrorReason(val *SponsoredProductsProductIdentifierErrorReason) *NullableSponsoredProductsProductIdentifierErrorReason {
	return &NullableSponsoredProductsProductIdentifierErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductIdentifierErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductIdentifierErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
