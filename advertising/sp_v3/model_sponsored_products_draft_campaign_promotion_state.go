package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsDraftCampaignPromotionState the model 'SponsoredProductsDraftCampaignPromotionState'
type SponsoredProductsDraftCampaignPromotionState string

// List of SponsoredProductsDraftCampaignPromotionState
const (
	SPONSOREDPRODUCTSDRAFTCAMPAIGNPROMOTIONSTATE_IN_PROGRESS SponsoredProductsDraftCampaignPromotionState = "IN_PROGRESS"
	SPONSOREDPRODUCTSDRAFTCAMPAIGNPROMOTIONSTATE_FAILED      SponsoredProductsDraftCampaignPromotionState = "FAILED"
	SPONSOREDPRODUCTSDRAFTCAMPAIGNPROMOTIONSTATE_SUCCEEDED   SponsoredProductsDraftCampaignPromotionState = "SUCCEEDED"
)

// All allowed values of SponsoredProductsDraftCampaignPromotionState enum
var AllowedSponsoredProductsDraftCampaignPromotionStateEnumValues = []SponsoredProductsDraftCampaignPromotionState{
	"IN_PROGRESS",
	"FAILED",
	"SUCCEEDED",
}

func (v *SponsoredProductsDraftCampaignPromotionState) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsDraftCampaignPromotionState(value)
	for _, existing := range AllowedSponsoredProductsDraftCampaignPromotionStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsDraftCampaignPromotionState", value)
}

// NewSponsoredProductsDraftCampaignPromotionStateFromValue returns a pointer to a valid SponsoredProductsDraftCampaignPromotionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsDraftCampaignPromotionStateFromValue(v string) (*SponsoredProductsDraftCampaignPromotionState, error) {
	ev := SponsoredProductsDraftCampaignPromotionState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsDraftCampaignPromotionState: valid values are %v", v, AllowedSponsoredProductsDraftCampaignPromotionStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsDraftCampaignPromotionState) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsDraftCampaignPromotionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsDraftCampaignPromotionState value
func (v SponsoredProductsDraftCampaignPromotionState) Ptr() *SponsoredProductsDraftCampaignPromotionState {
	return &v
}

type NullableSponsoredProductsDraftCampaignPromotionState struct {
	value *SponsoredProductsDraftCampaignPromotionState
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignPromotionState) Get() *SponsoredProductsDraftCampaignPromotionState {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignPromotionState) Set(val *SponsoredProductsDraftCampaignPromotionState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignPromotionState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignPromotionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignPromotionState(val *SponsoredProductsDraftCampaignPromotionState) *NullableSponsoredProductsDraftCampaignPromotionState {
	return &NullableSponsoredProductsDraftCampaignPromotionState{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignPromotionState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignPromotionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
