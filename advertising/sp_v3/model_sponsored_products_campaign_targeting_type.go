package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsCampaignTargetingType Targeting type of a campaign.
type SponsoredProductsCampaignTargetingType string

// List of SponsoredProductsCampaignTargetingType
const (
	SPONSOREDPRODUCTSCAMPAIGNTARGETINGTYPE_MANUAL SponsoredProductsCampaignTargetingType = "MANUAL"
	SPONSOREDPRODUCTSCAMPAIGNTARGETINGTYPE_AUTO   SponsoredProductsCampaignTargetingType = "AUTO"
)

// All allowed values of SponsoredProductsCampaignTargetingType enum
var AllowedSponsoredProductsCampaignTargetingTypeEnumValues = []SponsoredProductsCampaignTargetingType{
	"MANUAL",
	"AUTO",
}

func (v *SponsoredProductsCampaignTargetingType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCampaignTargetingType(value)
	for _, existing := range AllowedSponsoredProductsCampaignTargetingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCampaignTargetingType", value)
}

// NewSponsoredProductsCampaignTargetingTypeFromValue returns a pointer to a valid SponsoredProductsCampaignTargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCampaignTargetingTypeFromValue(v string) (*SponsoredProductsCampaignTargetingType, error) {
	ev := SponsoredProductsCampaignTargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCampaignTargetingType: valid values are %v", v, AllowedSponsoredProductsCampaignTargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCampaignTargetingType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCampaignTargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCampaignTargetingType value
func (v SponsoredProductsCampaignTargetingType) Ptr() *SponsoredProductsCampaignTargetingType {
	return &v
}

type NullableSponsoredProductsCampaignTargetingType struct {
	value *SponsoredProductsCampaignTargetingType
	isSet bool
}

func (v NullableSponsoredProductsCampaignTargetingType) Get() *SponsoredProductsCampaignTargetingType {
	return v.value
}

func (v *NullableSponsoredProductsCampaignTargetingType) Set(val *SponsoredProductsCampaignTargetingType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignTargetingType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignTargetingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignTargetingType(val *SponsoredProductsCampaignTargetingType) *NullableSponsoredProductsCampaignTargetingType {
	return &NullableSponsoredProductsCampaignTargetingType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignTargetingType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignTargetingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
