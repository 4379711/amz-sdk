package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalCampaignMarketplaceServingReason the model 'SponsoredProductsGlobalCampaignMarketplaceServingReason'
type SponsoredProductsGlobalCampaignMarketplaceServingReason string

// List of SponsoredProductsGlobalCampaignMarketplaceServingReason
const (
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_CAMPAIGN_STATUS_ENABLED_DETAIL            SponsoredProductsGlobalCampaignMarketplaceServingReason = "CAMPAIGN_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_CAMPAIGN_PAUSED_DETAIL                    SponsoredProductsGlobalCampaignMarketplaceServingReason = "CAMPAIGN_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_CAMPAIGN_ARCHIVED_DETAIL                  SponsoredProductsGlobalCampaignMarketplaceServingReason = "CAMPAIGN_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_PENDING_REVIEW_DETAIL                     SponsoredProductsGlobalCampaignMarketplaceServingReason = "PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_REJECTED_DETAIL                           SponsoredProductsGlobalCampaignMarketplaceServingReason = "REJECTED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_PENDING_START_DATE_DETAIL                 SponsoredProductsGlobalCampaignMarketplaceServingReason = "PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_ENDED_DETAIL                              SponsoredProductsGlobalCampaignMarketplaceServingReason = "ENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_CAMPAIGN_OUT_OF_BUDGET_DETAIL             SponsoredProductsGlobalCampaignMarketplaceServingReason = "CAMPAIGN_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_CAMPAIGN_INCOMPLETE_DETAIL                SponsoredProductsGlobalCampaignMarketplaceServingReason = "CAMPAIGN_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_ADVERTISER_POLICING_SUSPENDED_DETAIL      SponsoredProductsGlobalCampaignMarketplaceServingReason = "ADVERTISER_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_ADVERTISER_POLICING_PENDING_REVIEW_DETAIL SponsoredProductsGlobalCampaignMarketplaceServingReason = "ADVERTISER_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_ADVERTISER_ARCHIVED_DETAIL                SponsoredProductsGlobalCampaignMarketplaceServingReason = "ADVERTISER_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_ADVERTISER_PAUSED_DETAIL                  SponsoredProductsGlobalCampaignMarketplaceServingReason = "ADVERTISER_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL           SponsoredProductsGlobalCampaignMarketplaceServingReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL         SponsoredProductsGlobalCampaignMarketplaceServingReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
)

// All allowed values of SponsoredProductsGlobalCampaignMarketplaceServingReason enum
var AllowedSponsoredProductsGlobalCampaignMarketplaceServingReasonEnumValues = []SponsoredProductsGlobalCampaignMarketplaceServingReason{
	"CAMPAIGN_STATUS_ENABLED_DETAIL",
	"CAMPAIGN_PAUSED_DETAIL",
	"CAMPAIGN_ARCHIVED_DETAIL",
	"PENDING_REVIEW_DETAIL",
	"REJECTED_DETAIL",
	"PENDING_START_DATE_DETAIL",
	"ENDED_DETAIL",
	"CAMPAIGN_OUT_OF_BUDGET_DETAIL",
	"CAMPAIGN_INCOMPLETE_DETAIL",
	"ADVERTISER_POLICING_SUSPENDED_DETAIL",
	"ADVERTISER_POLICING_PENDING_REVIEW_DETAIL",
	"ADVERTISER_ARCHIVED_DETAIL",
	"ADVERTISER_PAUSED_DETAIL",
	"ADVERTISER_OUT_OF_BUDGET_DETAIL",
	"ADVERTISER_PAYMENT_FAILURE_DETAIL",
}

func (v *SponsoredProductsGlobalCampaignMarketplaceServingReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalCampaignMarketplaceServingReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalCampaignMarketplaceServingReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalCampaignMarketplaceServingReason", value)
}

// NewSponsoredProductsGlobalCampaignMarketplaceServingReasonFromValue returns a pointer to a valid SponsoredProductsGlobalCampaignMarketplaceServingReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalCampaignMarketplaceServingReasonFromValue(v string) (*SponsoredProductsGlobalCampaignMarketplaceServingReason, error) {
	ev := SponsoredProductsGlobalCampaignMarketplaceServingReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalCampaignMarketplaceServingReason: valid values are %v", v, AllowedSponsoredProductsGlobalCampaignMarketplaceServingReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalCampaignMarketplaceServingReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalCampaignMarketplaceServingReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalCampaignMarketplaceServingReason value
func (v SponsoredProductsGlobalCampaignMarketplaceServingReason) Ptr() *SponsoredProductsGlobalCampaignMarketplaceServingReason {
	return &v
}

type NullableSponsoredProductsGlobalCampaignMarketplaceServingReason struct {
	value *SponsoredProductsGlobalCampaignMarketplaceServingReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignMarketplaceServingReason) Get() *SponsoredProductsGlobalCampaignMarketplaceServingReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignMarketplaceServingReason) Set(val *SponsoredProductsGlobalCampaignMarketplaceServingReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignMarketplaceServingReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignMarketplaceServingReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignMarketplaceServingReason(val *SponsoredProductsGlobalCampaignMarketplaceServingReason) *NullableSponsoredProductsGlobalCampaignMarketplaceServingReason {
	return &NullableSponsoredProductsGlobalCampaignMarketplaceServingReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignMarketplaceServingReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignMarketplaceServingReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
