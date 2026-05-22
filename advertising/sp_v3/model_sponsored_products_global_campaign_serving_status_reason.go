package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalCampaignServingStatusReason the model 'SponsoredProductsGlobalCampaignServingStatusReason'
type SponsoredProductsGlobalCampaignServingStatusReason string

// List of SponsoredProductsGlobalCampaignServingStatusReason
const (
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_DELIVERING_IN_ALL_MARKETPLACES                        SponsoredProductsGlobalCampaignServingStatusReason = "CAMPAIGN_DELIVERING_IN_ALL_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE         SponsoredProductsGlobalCampaignServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES SponsoredProductsGlobalCampaignServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES      SponsoredProductsGlobalCampaignServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES                  SponsoredProductsGlobalCampaignServingStatusReason = "ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES                SponsoredProductsGlobalCampaignServingStatusReason = "ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE             SponsoredProductsGlobalCampaignServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                    SponsoredProductsGlobalCampaignServingStatusReason = "CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_ARCHIVED                                              SponsoredProductsGlobalCampaignServingStatusReason = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE                SponsoredProductsGlobalCampaignServingStatusReason = "CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_PENDING_START_DATE                                    SponsoredProductsGlobalCampaignServingStatusReason = "CAMPAIGN_PENDING_START_DATE"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_ENDED                                                 SponsoredProductsGlobalCampaignServingStatusReason = "CAMPAIGN_ENDED"
)

// All allowed values of SponsoredProductsGlobalCampaignServingStatusReason enum
var AllowedSponsoredProductsGlobalCampaignServingStatusReasonEnumValues = []SponsoredProductsGlobalCampaignServingStatusReason{
	"CAMPAIGN_DELIVERING_IN_ALL_MARKETPLACES",
	"ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE",
	"ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES",
	"CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_ARCHIVED",
	"CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_PENDING_START_DATE",
	"CAMPAIGN_ENDED",
}

func (v *SponsoredProductsGlobalCampaignServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalCampaignServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalCampaignServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalCampaignServingStatusReason", value)
}

// NewSponsoredProductsGlobalCampaignServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsGlobalCampaignServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalCampaignServingStatusReasonFromValue(v string) (*SponsoredProductsGlobalCampaignServingStatusReason, error) {
	ev := SponsoredProductsGlobalCampaignServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalCampaignServingStatusReason: valid values are %v", v, AllowedSponsoredProductsGlobalCampaignServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalCampaignServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalCampaignServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalCampaignServingStatusReason value
func (v SponsoredProductsGlobalCampaignServingStatusReason) Ptr() *SponsoredProductsGlobalCampaignServingStatusReason {
	return &v
}

type NullableSponsoredProductsGlobalCampaignServingStatusReason struct {
	value *SponsoredProductsGlobalCampaignServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignServingStatusReason) Get() *SponsoredProductsGlobalCampaignServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignServingStatusReason) Set(val *SponsoredProductsGlobalCampaignServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignServingStatusReason(val *SponsoredProductsGlobalCampaignServingStatusReason) *NullableSponsoredProductsGlobalCampaignServingStatusReason {
	return &NullableSponsoredProductsGlobalCampaignServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
