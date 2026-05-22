package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalProductAdMarketplaceServingReason the model 'SponsoredProductsGlobalProductAdMarketplaceServingReason'
type SponsoredProductsGlobalProductAdMarketplaceServingReason string

// List of SponsoredProductsGlobalProductAdMarketplaceServingReason
const (
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_GROUP_STATUS_ENABLED_DETAIL             SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_GROUP_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_GROUP_PAUSED_DETAIL                     SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_GROUP_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_GROUP_ARCHIVED_DETAIL                   SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_GROUP_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_GROUP_INCOMPLETE_DETAIL                 SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_GROUP_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_GROUP_POLICING_PENDING_REVIEW_DETAIL    SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_GROUP_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_GROUP_LOW_BID_DETAIL                    SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_GROUP_LOW_BID_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_CAMPAIGN_STATUS_ENABLED_DETAIL             SponsoredProductsGlobalProductAdMarketplaceServingReason = "CAMPAIGN_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_CAMPAIGN_PAUSED_DETAIL                     SponsoredProductsGlobalProductAdMarketplaceServingReason = "CAMPAIGN_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_CAMPAIGN_ARCHIVED_DETAIL                   SponsoredProductsGlobalProductAdMarketplaceServingReason = "CAMPAIGN_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_PENDING_REVIEW_DETAIL                      SponsoredProductsGlobalProductAdMarketplaceServingReason = "PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_REJECTED_DETAIL                            SponsoredProductsGlobalProductAdMarketplaceServingReason = "REJECTED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_PENDING_START_DATE_DETAIL                  SponsoredProductsGlobalProductAdMarketplaceServingReason = "PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_ENDED_DETAIL                               SponsoredProductsGlobalProductAdMarketplaceServingReason = "ENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_CAMPAIGN_OUT_OF_BUDGET_DETAIL              SponsoredProductsGlobalProductAdMarketplaceServingReason = "CAMPAIGN_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_CAMPAIGN_INCOMPLETE_DETAIL                 SponsoredProductsGlobalProductAdMarketplaceServingReason = "CAMPAIGN_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_ADVERTISER_POLICING_SUSPENDED_DETAIL       SponsoredProductsGlobalProductAdMarketplaceServingReason = "ADVERTISER_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_ADVERTISER_POLICING_PENDING_REVIEW_DETAIL  SponsoredProductsGlobalProductAdMarketplaceServingReason = "ADVERTISER_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_ADVERTISER_ARCHIVED_DETAIL                 SponsoredProductsGlobalProductAdMarketplaceServingReason = "ADVERTISER_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_ADVERTISER_PAUSED_DETAIL                   SponsoredProductsGlobalProductAdMarketplaceServingReason = "ADVERTISER_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL            SponsoredProductsGlobalProductAdMarketplaceServingReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL          SponsoredProductsGlobalProductAdMarketplaceServingReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_STATUS_LIVE_DETAIL                      SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_STATUS_LIVE_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_POLICING_PENDING_REVIEW_DETAIL          SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_POLICING_SUSPENDED_DETAIL               SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_PAUSED_DETAIL                           SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGREASON_AD_ARCHIVED_DETAIL                         SponsoredProductsGlobalProductAdMarketplaceServingReason = "AD_ARCHIVED_DETAIL"
)

// All allowed values of SponsoredProductsGlobalProductAdMarketplaceServingReason enum
var AllowedSponsoredProductsGlobalProductAdMarketplaceServingReasonEnumValues = []SponsoredProductsGlobalProductAdMarketplaceServingReason{
	"AD_GROUP_STATUS_ENABLED_DETAIL",
	"AD_GROUP_PAUSED_DETAIL",
	"AD_GROUP_ARCHIVED_DETAIL",
	"AD_GROUP_INCOMPLETE_DETAIL",
	"AD_GROUP_POLICING_PENDING_REVIEW_DETAIL",
	"AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL",
	"AD_GROUP_LOW_BID_DETAIL",
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
	"AD_STATUS_LIVE_DETAIL",
	"AD_POLICING_PENDING_REVIEW_DETAIL",
	"AD_POLICING_SUSPENDED_DETAIL",
	"AD_PAUSED_DETAIL",
	"AD_ARCHIVED_DETAIL",
}

func (v *SponsoredProductsGlobalProductAdMarketplaceServingReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalProductAdMarketplaceServingReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalProductAdMarketplaceServingReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalProductAdMarketplaceServingReason", value)
}

// NewSponsoredProductsGlobalProductAdMarketplaceServingReasonFromValue returns a pointer to a valid SponsoredProductsGlobalProductAdMarketplaceServingReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalProductAdMarketplaceServingReasonFromValue(v string) (*SponsoredProductsGlobalProductAdMarketplaceServingReason, error) {
	ev := SponsoredProductsGlobalProductAdMarketplaceServingReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalProductAdMarketplaceServingReason: valid values are %v", v, AllowedSponsoredProductsGlobalProductAdMarketplaceServingReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalProductAdMarketplaceServingReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalProductAdMarketplaceServingReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalProductAdMarketplaceServingReason value
func (v SponsoredProductsGlobalProductAdMarketplaceServingReason) Ptr() *SponsoredProductsGlobalProductAdMarketplaceServingReason {
	return &v
}

type NullableSponsoredProductsGlobalProductAdMarketplaceServingReason struct {
	value *SponsoredProductsGlobalProductAdMarketplaceServingReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAdMarketplaceServingReason) Get() *SponsoredProductsGlobalProductAdMarketplaceServingReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAdMarketplaceServingReason) Set(val *SponsoredProductsGlobalProductAdMarketplaceServingReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAdMarketplaceServingReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAdMarketplaceServingReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAdMarketplaceServingReason(val *SponsoredProductsGlobalProductAdMarketplaceServingReason) *NullableSponsoredProductsGlobalProductAdMarketplaceServingReason {
	return &NullableSponsoredProductsGlobalProductAdMarketplaceServingReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAdMarketplaceServingReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAdMarketplaceServingReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
