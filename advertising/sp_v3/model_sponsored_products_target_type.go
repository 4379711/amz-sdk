package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsTargetType the model 'SponsoredProductsTargetType'
type SponsoredProductsTargetType string

// List of SponsoredProductsTargetType
const (
	SPONSOREDPRODUCTSTARGETTYPE_AUTO             SponsoredProductsTargetType = "AUTO"
	SPONSOREDPRODUCTSTARGETTYPE_KEYWORD          SponsoredProductsTargetType = "KEYWORD"
	SPONSOREDPRODUCTSTARGETTYPE_PRODUCT          SponsoredProductsTargetType = "PRODUCT"
	SPONSOREDPRODUCTSTARGETTYPE_PRODUCT_CATEGORY SponsoredProductsTargetType = "PRODUCT_CATEGORY"
)

// All allowed values of SponsoredProductsTargetType enum
var AllowedSponsoredProductsTargetTypeEnumValues = []SponsoredProductsTargetType{
	"AUTO",
	"KEYWORD",
	"PRODUCT",
	"PRODUCT_CATEGORY",
}

func (v *SponsoredProductsTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsTargetType(value)
	for _, existing := range AllowedSponsoredProductsTargetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsTargetType", value)
}

// NewSponsoredProductsTargetTypeFromValue returns a pointer to a valid SponsoredProductsTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsTargetTypeFromValue(v string) (*SponsoredProductsTargetType, error) {
	ev := SponsoredProductsTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsTargetType: valid values are %v", v, AllowedSponsoredProductsTargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsTargetType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsTargetType value
func (v SponsoredProductsTargetType) Ptr() *SponsoredProductsTargetType {
	return &v
}

type NullableSponsoredProductsTargetType struct {
	value *SponsoredProductsTargetType
	isSet bool
}

func (v NullableSponsoredProductsTargetType) Get() *SponsoredProductsTargetType {
	return v.value
}

func (v *NullableSponsoredProductsTargetType) Set(val *SponsoredProductsTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetType(val *SponsoredProductsTargetType) *NullableSponsoredProductsTargetType {
	return &NullableSponsoredProductsTargetType{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
