package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateDraftCampaignBudgetType the model 'SponsoredProductsCreateOrUpdateDraftCampaignBudgetType'
type SponsoredProductsCreateOrUpdateDraftCampaignBudgetType string

// List of SponsoredProductsCreateOrUpdateDraftCampaignBudgetType
const (
	SPONSOREDPRODUCTSCREATEORUPDATEDRAFTCAMPAIGNBUDGETTYPE_DAILY SponsoredProductsCreateOrUpdateDraftCampaignBudgetType = "DAILY"
)

// All allowed values of SponsoredProductsCreateOrUpdateDraftCampaignBudgetType enum
var AllowedSponsoredProductsCreateOrUpdateDraftCampaignBudgetTypeEnumValues = []SponsoredProductsCreateOrUpdateDraftCampaignBudgetType{
	"DAILY",
}

func (v *SponsoredProductsCreateOrUpdateDraftCampaignBudgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateOrUpdateDraftCampaignBudgetType(value)
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateDraftCampaignBudgetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateOrUpdateDraftCampaignBudgetType", value)
}

// NewSponsoredProductsCreateOrUpdateDraftCampaignBudgetTypeFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateDraftCampaignBudgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateDraftCampaignBudgetTypeFromValue(v string) (*SponsoredProductsCreateOrUpdateDraftCampaignBudgetType, error) {
	ev := SponsoredProductsCreateOrUpdateDraftCampaignBudgetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateDraftCampaignBudgetType: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateDraftCampaignBudgetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateDraftCampaignBudgetType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateDraftCampaignBudgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateDraftCampaignBudgetType value
func (v SponsoredProductsCreateOrUpdateDraftCampaignBudgetType) Ptr() *SponsoredProductsCreateOrUpdateDraftCampaignBudgetType {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType struct {
	value *SponsoredProductsCreateOrUpdateDraftCampaignBudgetType
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType) Get() *SponsoredProductsCreateOrUpdateDraftCampaignBudgetType {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType) Set(val *SponsoredProductsCreateOrUpdateDraftCampaignBudgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType(val *SponsoredProductsCreateOrUpdateDraftCampaignBudgetType) *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType {
	return &NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
