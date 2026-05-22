package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsProductTargetMatchType the model 'SponsoredProductsProductTargetMatchType'
type SponsoredProductsProductTargetMatchType string

// List of SponsoredProductsProductTargetMatchType
const (
	SPONSOREDPRODUCTSPRODUCTTARGETMATCHTYPE_EXACT   SponsoredProductsProductTargetMatchType = "PRODUCT_EXACT"
	SPONSOREDPRODUCTSPRODUCTTARGETMATCHTYPE_SIMILAR SponsoredProductsProductTargetMatchType = "PRODUCT_SIMILAR"
)

// All allowed values of SponsoredProductsProductTargetMatchType enum
var AllowedSponsoredProductsProductTargetMatchTypeEnumValues = []SponsoredProductsProductTargetMatchType{
	"PRODUCT_EXACT",
	"PRODUCT_SIMILAR",
}

func (v *SponsoredProductsProductTargetMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsProductTargetMatchType(value)
	for _, existing := range AllowedSponsoredProductsProductTargetMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsProductTargetMatchType", value)
}

// NewSponsoredProductsProductTargetMatchTypeFromValue returns a pointer to a valid SponsoredProductsProductTargetMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsProductTargetMatchTypeFromValue(v string) (*SponsoredProductsProductTargetMatchType, error) {
	ev := SponsoredProductsProductTargetMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsProductTargetMatchType: valid values are %v", v, AllowedSponsoredProductsProductTargetMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsProductTargetMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsProductTargetMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsProductTargetMatchType value
func (v SponsoredProductsProductTargetMatchType) Ptr() *SponsoredProductsProductTargetMatchType {
	return &v
}

type NullableSponsoredProductsProductTargetMatchType struct {
	value *SponsoredProductsProductTargetMatchType
	isSet bool
}

func (v NullableSponsoredProductsProductTargetMatchType) Get() *SponsoredProductsProductTargetMatchType {
	return v.value
}

func (v *NullableSponsoredProductsProductTargetMatchType) Set(val *SponsoredProductsProductTargetMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductTargetMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductTargetMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductTargetMatchType(val *SponsoredProductsProductTargetMatchType) *NullableSponsoredProductsProductTargetMatchType {
	return &NullableSponsoredProductsProductTargetMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductTargetMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductTargetMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
