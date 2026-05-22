package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsNegativeMatchType the model 'SponsoredProductsNegativeMatchType'
type SponsoredProductsNegativeMatchType string

// List of SponsoredProductsNegativeMatchType
const (
	SPONSOREDPRODUCTSNEGATIVEMATCHTYPE_NEGATIVE_EXACT  SponsoredProductsNegativeMatchType = "NEGATIVE_EXACT"
	SPONSOREDPRODUCTSNEGATIVEMATCHTYPE_NEGATIVE_PHRASE SponsoredProductsNegativeMatchType = "NEGATIVE_PHRASE"
	SPONSOREDPRODUCTSNEGATIVEMATCHTYPE_NEGATIVE_BROAD  SponsoredProductsNegativeMatchType = "NEGATIVE_BROAD"
	SPONSOREDPRODUCTSNEGATIVEMATCHTYPE_OTHER           SponsoredProductsNegativeMatchType = "OTHER"
)

// All allowed values of SponsoredProductsNegativeMatchType enum
var AllowedSponsoredProductsNegativeMatchTypeEnumValues = []SponsoredProductsNegativeMatchType{
	"NEGATIVE_EXACT",
	"NEGATIVE_PHRASE",
	"NEGATIVE_BROAD",
	"OTHER",
}

func (v *SponsoredProductsNegativeMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsNegativeMatchType(value)
	for _, existing := range AllowedSponsoredProductsNegativeMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsNegativeMatchType", value)
}

// NewSponsoredProductsNegativeMatchTypeFromValue returns a pointer to a valid SponsoredProductsNegativeMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsNegativeMatchTypeFromValue(v string) (*SponsoredProductsNegativeMatchType, error) {
	ev := SponsoredProductsNegativeMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsNegativeMatchType: valid values are %v", v, AllowedSponsoredProductsNegativeMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsNegativeMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsNegativeMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsNegativeMatchType value
func (v SponsoredProductsNegativeMatchType) Ptr() *SponsoredProductsNegativeMatchType {
	return &v
}

type NullableSponsoredProductsNegativeMatchType struct {
	value *SponsoredProductsNegativeMatchType
	isSet bool
}

func (v NullableSponsoredProductsNegativeMatchType) Get() *SponsoredProductsNegativeMatchType {
	return v.value
}

func (v *NullableSponsoredProductsNegativeMatchType) Set(val *SponsoredProductsNegativeMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeMatchType(val *SponsoredProductsNegativeMatchType) *NullableSponsoredProductsNegativeMatchType {
	return &NullableSponsoredProductsNegativeMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
