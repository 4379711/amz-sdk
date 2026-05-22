package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsMatchType the model 'SponsoredProductsMatchType'
type SponsoredProductsMatchType string

// List of SponsoredProductsMatchType
const (
	SPONSOREDPRODUCTSMATCHTYPE_EXACT  SponsoredProductsMatchType = "EXACT"
	SPONSOREDPRODUCTSMATCHTYPE_PHRASE SponsoredProductsMatchType = "PHRASE"
	SPONSOREDPRODUCTSMATCHTYPE_BROAD  SponsoredProductsMatchType = "BROAD"
	SPONSOREDPRODUCTSMATCHTYPE_OTHER  SponsoredProductsMatchType = "OTHER"
)

// All allowed values of SponsoredProductsMatchType enum
var AllowedSponsoredProductsMatchTypeEnumValues = []SponsoredProductsMatchType{
	"EXACT",
	"PHRASE",
	"BROAD",
	"OTHER",
}

func (v *SponsoredProductsMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsMatchType(value)
	for _, existing := range AllowedSponsoredProductsMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsMatchType", value)
}

// NewSponsoredProductsMatchTypeFromValue returns a pointer to a valid SponsoredProductsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsMatchTypeFromValue(v string) (*SponsoredProductsMatchType, error) {
	ev := SponsoredProductsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsMatchType: valid values are %v", v, AllowedSponsoredProductsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsMatchType value
func (v SponsoredProductsMatchType) Ptr() *SponsoredProductsMatchType {
	return &v
}

type NullableSponsoredProductsMatchType struct {
	value *SponsoredProductsMatchType
	isSet bool
}

func (v NullableSponsoredProductsMatchType) Get() *SponsoredProductsMatchType {
	return v.value
}

func (v *NullableSponsoredProductsMatchType) Set(val *SponsoredProductsMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMatchType(val *SponsoredProductsMatchType) *NullableSponsoredProductsMatchType {
	return &NullableSponsoredProductsMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
