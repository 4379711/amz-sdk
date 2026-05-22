package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsDraftCampaignBudgetType the model 'SponsoredProductsDraftCampaignBudgetType'
type SponsoredProductsDraftCampaignBudgetType string

// List of SponsoredProductsDraftCampaignBudgetType
const (
	SPONSOREDPRODUCTSDRAFTCAMPAIGNBUDGETTYPE_DAILY SponsoredProductsDraftCampaignBudgetType = "DAILY"
	SPONSOREDPRODUCTSDRAFTCAMPAIGNBUDGETTYPE_OTHER SponsoredProductsDraftCampaignBudgetType = "OTHER"
)

// All allowed values of SponsoredProductsDraftCampaignBudgetType enum
var AllowedSponsoredProductsDraftCampaignBudgetTypeEnumValues = []SponsoredProductsDraftCampaignBudgetType{
	"DAILY",
	"OTHER",
}

func (v *SponsoredProductsDraftCampaignBudgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsDraftCampaignBudgetType(value)
	for _, existing := range AllowedSponsoredProductsDraftCampaignBudgetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsDraftCampaignBudgetType", value)
}

// NewSponsoredProductsDraftCampaignBudgetTypeFromValue returns a pointer to a valid SponsoredProductsDraftCampaignBudgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsDraftCampaignBudgetTypeFromValue(v string) (*SponsoredProductsDraftCampaignBudgetType, error) {
	ev := SponsoredProductsDraftCampaignBudgetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsDraftCampaignBudgetType: valid values are %v", v, AllowedSponsoredProductsDraftCampaignBudgetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsDraftCampaignBudgetType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsDraftCampaignBudgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsDraftCampaignBudgetType value
func (v SponsoredProductsDraftCampaignBudgetType) Ptr() *SponsoredProductsDraftCampaignBudgetType {
	return &v
}

type NullableSponsoredProductsDraftCampaignBudgetType struct {
	value *SponsoredProductsDraftCampaignBudgetType
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignBudgetType) Get() *SponsoredProductsDraftCampaignBudgetType {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignBudgetType) Set(val *SponsoredProductsDraftCampaignBudgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignBudgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignBudgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignBudgetType(val *SponsoredProductsDraftCampaignBudgetType) *NullableSponsoredProductsDraftCampaignBudgetType {
	return &NullableSponsoredProductsDraftCampaignBudgetType{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignBudgetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignBudgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
