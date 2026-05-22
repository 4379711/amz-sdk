package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateNegativeMatchType the model 'SponsoredProductsCreateOrUpdateNegativeMatchType'
type SponsoredProductsCreateOrUpdateNegativeMatchType string

// List of SponsoredProductsCreateOrUpdateNegativeMatchType
const (
	SPONSOREDPRODUCTSCREATEORUPDATENEGATIVEMATCHTYPE_EXACT  SponsoredProductsCreateOrUpdateNegativeMatchType = "NEGATIVE_EXACT"
	SPONSOREDPRODUCTSCREATEORUPDATENEGATIVEMATCHTYPE_PHRASE SponsoredProductsCreateOrUpdateNegativeMatchType = "NEGATIVE_PHRASE"
	SPONSOREDPRODUCTSCREATEORUPDATENEGATIVEMATCHTYPE_BROAD  SponsoredProductsCreateOrUpdateNegativeMatchType = "NEGATIVE_BROAD"
)

// All allowed values of SponsoredProductsCreateOrUpdateNegativeMatchType enum
var AllowedSponsoredProductsCreateOrUpdateNegativeMatchTypeEnumValues = []SponsoredProductsCreateOrUpdateNegativeMatchType{
	"NEGATIVE_EXACT",
	"NEGATIVE_PHRASE",
	"NEGATIVE_BROAD",
}

func (v *SponsoredProductsCreateOrUpdateNegativeMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateOrUpdateNegativeMatchType(value)
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateNegativeMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateOrUpdateNegativeMatchType", value)
}

// NewSponsoredProductsCreateOrUpdateNegativeMatchTypeFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateNegativeMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateNegativeMatchTypeFromValue(v string) (*SponsoredProductsCreateOrUpdateNegativeMatchType, error) {
	ev := SponsoredProductsCreateOrUpdateNegativeMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateNegativeMatchType: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateNegativeMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateNegativeMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateNegativeMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateNegativeMatchType value
func (v SponsoredProductsCreateOrUpdateNegativeMatchType) Ptr() *SponsoredProductsCreateOrUpdateNegativeMatchType {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateNegativeMatchType struct {
	value *SponsoredProductsCreateOrUpdateNegativeMatchType
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateNegativeMatchType) Get() *SponsoredProductsCreateOrUpdateNegativeMatchType {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateNegativeMatchType) Set(val *SponsoredProductsCreateOrUpdateNegativeMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateNegativeMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateNegativeMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateNegativeMatchType(val *SponsoredProductsCreateOrUpdateNegativeMatchType) *NullableSponsoredProductsCreateOrUpdateNegativeMatchType {
	return &NullableSponsoredProductsCreateOrUpdateNegativeMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateNegativeMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateNegativeMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
