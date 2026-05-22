package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalKeywordMarketplaceServingStatusReason the model 'SponsoredProductsGlobalKeywordMarketplaceServingStatusReason'
type SponsoredProductsGlobalKeywordMarketplaceServingStatusReason string

// List of SponsoredProductsGlobalKeywordMarketplaceServingStatusReason
const (
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_TARGETING_CLAUSE_STATUS_LIVE_DETAIL        SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "TARGETING_CLAUSE_STATUS_LIVE_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_TARGETING_CLAUSE_POLICING_SUSPENDED_DETAIL SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "TARGETING_CLAUSE_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_TARGETING_CLAUSE_PAUSED_DETAIL             SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "TARGETING_CLAUSE_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_TARGETING_CLAUSE_ARCHIVED_DETAIL           SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "TARGETING_CLAUSE_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_TARGETING_CLAUSE_BLOCKED_DETAIL            SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "TARGETING_CLAUSE_BLOCKED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_AD_GROUP_STATUS_ENABLED_DETAIL             SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "AD_GROUP_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_AD_GROUP_PAUSED_DETAIL                     SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "AD_GROUP_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_AD_GROUP_ARCHIVED_DETAIL                   SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "AD_GROUP_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_AD_GROUP_INCOMPLETE_DETAIL                 SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "AD_GROUP_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_AD_GROUP_POLICING_PENDING_REVIEW_DETAIL    SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "AD_GROUP_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_AD_GROUP_LOW_BID_DETAIL                    SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "AD_GROUP_LOW_BID_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_CAMPAIGN_STATUS_ENABLED_DETAIL             SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "CAMPAIGN_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_CAMPAIGN_PAUSED_DETAIL                     SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "CAMPAIGN_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_CAMPAIGN_ARCHIVED_DETAIL                   SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "CAMPAIGN_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_PENDING_REVIEW_DETAIL                      SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_REJECTED_DETAIL                            SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "REJECTED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_PENDING_START_DATE_DETAIL                  SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_ENDED_DETAIL                               SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "ENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_DETAIL              SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_DETAIL                 SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "CAMPAIGN_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_DETAIL       SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_DETAIL  SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_ADVERTISER_ARCHIVED_DETAIL                 SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "ADVERTISER_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_ADVERTISER_PAUSED_DETAIL                   SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "ADVERTISER_PAUSED_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL            SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL          SponsoredProductsGlobalKeywordMarketplaceServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
)

// All allowed values of SponsoredProductsGlobalKeywordMarketplaceServingStatusReason enum
var AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusReasonEnumValues = []SponsoredProductsGlobalKeywordMarketplaceServingStatusReason{
	"TARGETING_CLAUSE_STATUS_LIVE_DETAIL",
	"TARGETING_CLAUSE_POLICING_SUSPENDED_DETAIL",
	"TARGETING_CLAUSE_PAUSED_DETAIL",
	"TARGETING_CLAUSE_ARCHIVED_DETAIL",
	"TARGETING_CLAUSE_BLOCKED_DETAIL",
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
}

func (v *SponsoredProductsGlobalKeywordMarketplaceServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalKeywordMarketplaceServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalKeywordMarketplaceServingStatusReason", value)
}

// NewSponsoredProductsGlobalKeywordMarketplaceServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsGlobalKeywordMarketplaceServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalKeywordMarketplaceServingStatusReasonFromValue(v string) (*SponsoredProductsGlobalKeywordMarketplaceServingStatusReason, error) {
	ev := SponsoredProductsGlobalKeywordMarketplaceServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalKeywordMarketplaceServingStatusReason: valid values are %v", v, AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalKeywordMarketplaceServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalKeywordMarketplaceServingStatusReason value
func (v SponsoredProductsGlobalKeywordMarketplaceServingStatusReason) Ptr() *SponsoredProductsGlobalKeywordMarketplaceServingStatusReason {
	return &v
}

type NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason struct {
	value *SponsoredProductsGlobalKeywordMarketplaceServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason) Get() *SponsoredProductsGlobalKeywordMarketplaceServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason) Set(val *SponsoredProductsGlobalKeywordMarketplaceServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason(val *SponsoredProductsGlobalKeywordMarketplaceServingStatusReason) *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason {
	return &NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
