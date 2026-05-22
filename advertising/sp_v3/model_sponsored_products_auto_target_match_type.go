package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsAutoTargetMatchType Entity State. | Type |  description | |-----------|------------| | `SEARCH_LOOSE_MATCH` | Search terms loosely matching advertised product. | | `SEARCH_CLOSE_MATCH` | Search terms closely matching advertised product. | | `PRODUCT_SUBSTITUTES` | Products that can be substituted for advertised product. | | `PRODUCT_COMPLEMENTS` | Products that complement advertised product. | | `PRODUCT_SIMILAR` | Products similar to advertised product. |
type SponsoredProductsAutoTargetMatchType string

// List of SponsoredProductsAutoTargetMatchType
const (
	SPONSOREDPRODUCTSAUTOTARGETMATCHTYPE_SEARCH_LOOSE_MATCH  SponsoredProductsAutoTargetMatchType = "SEARCH_LOOSE_MATCH"
	SPONSOREDPRODUCTSAUTOTARGETMATCHTYPE_SEARCH_CLOSE_MATCH  SponsoredProductsAutoTargetMatchType = "SEARCH_CLOSE_MATCH"
	SPONSOREDPRODUCTSAUTOTARGETMATCHTYPE_PRODUCT_SUBSTITUTES SponsoredProductsAutoTargetMatchType = "PRODUCT_SUBSTITUTES"
	SPONSOREDPRODUCTSAUTOTARGETMATCHTYPE_PRODUCT_COMPLEMENTS SponsoredProductsAutoTargetMatchType = "PRODUCT_COMPLEMENTS"
	SPONSOREDPRODUCTSAUTOTARGETMATCHTYPE_PRODUCT_SIMILAR     SponsoredProductsAutoTargetMatchType = "PRODUCT_SIMILAR"
)

// All allowed values of SponsoredProductsAutoTargetMatchType enum
var AllowedSponsoredProductsAutoTargetMatchTypeEnumValues = []SponsoredProductsAutoTargetMatchType{
	"SEARCH_LOOSE_MATCH",
	"SEARCH_CLOSE_MATCH",
	"PRODUCT_SUBSTITUTES",
	"PRODUCT_COMPLEMENTS",
	"PRODUCT_SIMILAR",
}

func (v *SponsoredProductsAutoTargetMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAutoTargetMatchType(value)
	for _, existing := range AllowedSponsoredProductsAutoTargetMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAutoTargetMatchType", value)
}

// NewSponsoredProductsAutoTargetMatchTypeFromValue returns a pointer to a valid SponsoredProductsAutoTargetMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAutoTargetMatchTypeFromValue(v string) (*SponsoredProductsAutoTargetMatchType, error) {
	ev := SponsoredProductsAutoTargetMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAutoTargetMatchType: valid values are %v", v, AllowedSponsoredProductsAutoTargetMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAutoTargetMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAutoTargetMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAutoTargetMatchType value
func (v SponsoredProductsAutoTargetMatchType) Ptr() *SponsoredProductsAutoTargetMatchType {
	return &v
}

type NullableSponsoredProductsAutoTargetMatchType struct {
	value *SponsoredProductsAutoTargetMatchType
	isSet bool
}

func (v NullableSponsoredProductsAutoTargetMatchType) Get() *SponsoredProductsAutoTargetMatchType {
	return v.value
}

func (v *NullableSponsoredProductsAutoTargetMatchType) Set(val *SponsoredProductsAutoTargetMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAutoTargetMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAutoTargetMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAutoTargetMatchType(val *SponsoredProductsAutoTargetMatchType) *NullableSponsoredProductsAutoTargetMatchType {
	return &NullableSponsoredProductsAutoTargetMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsAutoTargetMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAutoTargetMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
