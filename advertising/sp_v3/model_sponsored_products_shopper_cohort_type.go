package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsShopperCohortType This field specifies the type of shopper cohort used to apply bid adjustments. `AUDIENCE_SEGMENT` is the only supported type currently. | Value |  Description | |-----------|------------| | `AUDIENCE_SEGMENT` | A predefined list of the shopper ids. |
type SponsoredProductsShopperCohortType string

// List of SponsoredProductsShopperCohortType
const (
	SPONSOREDPRODUCTSSHOPPERCOHORTTYPE_AUDIENCE_SEGMENT SponsoredProductsShopperCohortType = "AUDIENCE_SEGMENT"
)

// All allowed values of SponsoredProductsShopperCohortType enum
var AllowedSponsoredProductsShopperCohortTypeEnumValues = []SponsoredProductsShopperCohortType{
	"AUDIENCE_SEGMENT",
}

func (v *SponsoredProductsShopperCohortType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsShopperCohortType(value)
	for _, existing := range AllowedSponsoredProductsShopperCohortTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsShopperCohortType", value)
}

// NewSponsoredProductsShopperCohortTypeFromValue returns a pointer to a valid SponsoredProductsShopperCohortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsShopperCohortTypeFromValue(v string) (*SponsoredProductsShopperCohortType, error) {
	ev := SponsoredProductsShopperCohortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsShopperCohortType: valid values are %v", v, AllowedSponsoredProductsShopperCohortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsShopperCohortType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsShopperCohortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsShopperCohortType value
func (v SponsoredProductsShopperCohortType) Ptr() *SponsoredProductsShopperCohortType {
	return &v
}

type NullableSponsoredProductsShopperCohortType struct {
	value *SponsoredProductsShopperCohortType
	isSet bool
}

func (v NullableSponsoredProductsShopperCohortType) Get() *SponsoredProductsShopperCohortType {
	return v.value
}

func (v *NullableSponsoredProductsShopperCohortType) Set(val *SponsoredProductsShopperCohortType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsShopperCohortType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsShopperCohortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsShopperCohortType(val *SponsoredProductsShopperCohortType) *NullableSponsoredProductsShopperCohortType {
	return &NullableSponsoredProductsShopperCohortType{value: val, isSet: true}
}

func (v NullableSponsoredProductsShopperCohortType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsShopperCohortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
