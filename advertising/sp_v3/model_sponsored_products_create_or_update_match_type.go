package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateMatchType the model 'SponsoredProductsCreateOrUpdateMatchType'
type SponsoredProductsCreateOrUpdateMatchType string

// List of SponsoredProductsCreateOrUpdateMatchType
const (
	SPONSOREDPRODUCTSCREATEORUPDATEMATCHTYPE_EXACT  SponsoredProductsCreateOrUpdateMatchType = "EXACT"
	SPONSOREDPRODUCTSCREATEORUPDATEMATCHTYPE_PHRASE SponsoredProductsCreateOrUpdateMatchType = "PHRASE"
	SPONSOREDPRODUCTSCREATEORUPDATEMATCHTYPE_BROAD  SponsoredProductsCreateOrUpdateMatchType = "BROAD"
)

// All allowed values of SponsoredProductsCreateOrUpdateMatchType enum
var AllowedSponsoredProductsCreateOrUpdateMatchTypeEnumValues = []SponsoredProductsCreateOrUpdateMatchType{
	"EXACT",
	"PHRASE",
	"BROAD",
}

func (v *SponsoredProductsCreateOrUpdateMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateOrUpdateMatchType(value)
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateOrUpdateMatchType", value)
}

// NewSponsoredProductsCreateOrUpdateMatchTypeFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateMatchTypeFromValue(v string) (*SponsoredProductsCreateOrUpdateMatchType, error) {
	ev := SponsoredProductsCreateOrUpdateMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateMatchType: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateMatchType value
func (v SponsoredProductsCreateOrUpdateMatchType) Ptr() *SponsoredProductsCreateOrUpdateMatchType {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateMatchType struct {
	value *SponsoredProductsCreateOrUpdateMatchType
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateMatchType) Get() *SponsoredProductsCreateOrUpdateMatchType {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateMatchType) Set(val *SponsoredProductsCreateOrUpdateMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateMatchType(val *SponsoredProductsCreateOrUpdateMatchType) *NullableSponsoredProductsCreateOrUpdateMatchType {
	return &NullableSponsoredProductsCreateOrUpdateMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
