package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalAdGroupMarketplaceServingReason the model 'SponsoredProductsGlobalAdGroupMarketplaceServingReason'
type SponsoredProductsGlobalAdGroupMarketplaceServingReason string

// List of SponsoredProductsGlobalAdGroupMarketplaceServingReason
const (
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_AD_GROUP_STATUS_ENABLED_DETAIL             SponsoredProductsGlobalAdGroupMarketplaceServingReason = "AD_GROUP_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_AD_GROUP_PAUSED_DETAIL                     SponsoredProductsGlobalAdGroupMarketplaceServingReason = "AD_GROUP_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_AD_GROUP_ARCHIVED_DETAIL                   SponsoredProductsGlobalAdGroupMarketplaceServingReason = "AD_GROUP_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_AD_GROUP_INCOMPLETE_DETAIL                 SponsoredProductsGlobalAdGroupMarketplaceServingReason = "AD_GROUP_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_AD_GROUP_POLICING_PENDING_REVIEW_DETAIL    SponsoredProductsGlobalAdGroupMarketplaceServingReason = "AD_GROUP_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL SponsoredProductsGlobalAdGroupMarketplaceServingReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_AD_GROUP_LOW_BID_DETAIL                    SponsoredProductsGlobalAdGroupMarketplaceServingReason = "AD_GROUP_LOW_BID_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_CAMPAIGN_STATUS_ENABLED_DETAIL             SponsoredProductsGlobalAdGroupMarketplaceServingReason = "CAMPAIGN_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_CAMPAIGN_PAUSED_DETAIL                     SponsoredProductsGlobalAdGroupMarketplaceServingReason = "CAMPAIGN_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_CAMPAIGN_ARCHIVED_DETAIL                   SponsoredProductsGlobalAdGroupMarketplaceServingReason = "CAMPAIGN_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_PENDING_REVIEW_DETAIL                      SponsoredProductsGlobalAdGroupMarketplaceServingReason = "PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_REJECTED_DETAIL                            SponsoredProductsGlobalAdGroupMarketplaceServingReason = "REJECTED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_PENDING_START_DATE_DETAIL                  SponsoredProductsGlobalAdGroupMarketplaceServingReason = "PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ENDED_DETAIL                               SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_CAMPAIGN_OUT_OF_BUDGET_DETAIL              SponsoredProductsGlobalAdGroupMarketplaceServingReason = "CAMPAIGN_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_CAMPAIGN_INCOMPLETE_DETAIL                 SponsoredProductsGlobalAdGroupMarketplaceServingReason = "CAMPAIGN_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ADVERTISER_POLICING_SUSPENDED_DETAIL       SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ADVERTISER_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ADVERTISER_POLICING_PENDING_REVIEW_DETAIL  SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ADVERTISER_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ADVERTISER_ARCHIVED_DETAIL                 SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ADVERTISER_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ADVERTISER_PAUSED_DETAIL                   SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ADVERTISER_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL            SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL          SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGREASON_ADVERTISER_STATUS_ENABLED_DETAIL           SponsoredProductsGlobalAdGroupMarketplaceServingReason = "ADVERTISER_STATUS_ENABLED_DETAIL"
)

// All allowed values of SponsoredProductsGlobalAdGroupMarketplaceServingReason enum
var AllowedSponsoredProductsGlobalAdGroupMarketplaceServingReasonEnumValues = []SponsoredProductsGlobalAdGroupMarketplaceServingReason{
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
	"ADVERTISER_STATUS_ENABLED_DETAIL",
}

func (v *SponsoredProductsGlobalAdGroupMarketplaceServingReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalAdGroupMarketplaceServingReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalAdGroupMarketplaceServingReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalAdGroupMarketplaceServingReason", value)
}

// NewSponsoredProductsGlobalAdGroupMarketplaceServingReasonFromValue returns a pointer to a valid SponsoredProductsGlobalAdGroupMarketplaceServingReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalAdGroupMarketplaceServingReasonFromValue(v string) (*SponsoredProductsGlobalAdGroupMarketplaceServingReason, error) {
	ev := SponsoredProductsGlobalAdGroupMarketplaceServingReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalAdGroupMarketplaceServingReason: valid values are %v", v, AllowedSponsoredProductsGlobalAdGroupMarketplaceServingReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalAdGroupMarketplaceServingReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalAdGroupMarketplaceServingReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalAdGroupMarketplaceServingReason value
func (v SponsoredProductsGlobalAdGroupMarketplaceServingReason) Ptr() *SponsoredProductsGlobalAdGroupMarketplaceServingReason {
	return &v
}

type NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason struct {
	value *SponsoredProductsGlobalAdGroupMarketplaceServingReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason) Get() *SponsoredProductsGlobalAdGroupMarketplaceServingReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason) Set(val *SponsoredProductsGlobalAdGroupMarketplaceServingReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroupMarketplaceServingReason(val *SponsoredProductsGlobalAdGroupMarketplaceServingReason) *NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason {
	return &NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroupMarketplaceServingReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
