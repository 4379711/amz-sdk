package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsTargetLevel the model 'SponsoredProductsTargetLevel'
type SponsoredProductsTargetLevel string

// List of SponsoredProductsTargetLevel
const (
	SPONSOREDPRODUCTSTARGETLEVEL_AD_GROUP SponsoredProductsTargetLevel = "AD_GROUP"
	SPONSOREDPRODUCTSTARGETLEVEL_CAMPAIGN SponsoredProductsTargetLevel = "CAMPAIGN"
)

// All allowed values of SponsoredProductsTargetLevel enum
var AllowedSponsoredProductsTargetLevelEnumValues = []SponsoredProductsTargetLevel{
	"AD_GROUP",
	"CAMPAIGN",
}

func (v *SponsoredProductsTargetLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsTargetLevel(value)
	for _, existing := range AllowedSponsoredProductsTargetLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsTargetLevel", value)
}

// NewSponsoredProductsTargetLevelFromValue returns a pointer to a valid SponsoredProductsTargetLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsTargetLevelFromValue(v string) (*SponsoredProductsTargetLevel, error) {
	ev := SponsoredProductsTargetLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsTargetLevel: valid values are %v", v, AllowedSponsoredProductsTargetLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsTargetLevel) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsTargetLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsTargetLevel value
func (v SponsoredProductsTargetLevel) Ptr() *SponsoredProductsTargetLevel {
	return &v
}

type NullableSponsoredProductsTargetLevel struct {
	value *SponsoredProductsTargetLevel
	isSet bool
}

func (v NullableSponsoredProductsTargetLevel) Get() *SponsoredProductsTargetLevel {
	return v.value
}

func (v *NullableSponsoredProductsTargetLevel) Set(val *SponsoredProductsTargetLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetLevel(val *SponsoredProductsTargetLevel) *NullableSponsoredProductsTargetLevel {
	return &NullableSponsoredProductsTargetLevel{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetLevel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
