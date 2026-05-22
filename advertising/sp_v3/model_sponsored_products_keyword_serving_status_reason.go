package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsKeywordServingStatusReason the model 'SponsoredProductsKeywordServingStatusReason'
type SponsoredProductsKeywordServingStatusReason string

// List of SponsoredProductsKeywordServingStatusReason
const (
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_STATUS_LIVE_DETAIL        SponsoredProductsKeywordServingStatusReason = "TARGETING_CLAUSE_STATUS_LIVE_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_POLICING_SUSPENDED_DETAIL SponsoredProductsKeywordServingStatusReason = "TARGETING_CLAUSE_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_PAUSED_DETAIL             SponsoredProductsKeywordServingStatusReason = "TARGETING_CLAUSE_PAUSED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_ARCHIVED_DETAIL           SponsoredProductsKeywordServingStatusReason = "TARGETING_CLAUSE_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_BLOCKED_DETAIL            SponsoredProductsKeywordServingStatusReason = "TARGETING_CLAUSE_BLOCKED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_AD_GROUP_STATUS_ENABLED_DETAIL             SponsoredProductsKeywordServingStatusReason = "AD_GROUP_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_AD_GROUP_PAUSED_DETAIL                     SponsoredProductsKeywordServingStatusReason = "AD_GROUP_PAUSED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_AD_GROUP_ARCHIVED_DETAIL                   SponsoredProductsKeywordServingStatusReason = "AD_GROUP_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_AD_GROUP_INCOMPLETE_DETAIL                 SponsoredProductsKeywordServingStatusReason = "AD_GROUP_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_AD_GROUP_POLICING_PENDING_REVIEW_DETAIL    SponsoredProductsKeywordServingStatusReason = "AD_GROUP_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL SponsoredProductsKeywordServingStatusReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_AD_GROUP_LOW_BID_DETAIL                    SponsoredProductsKeywordServingStatusReason = "AD_GROUP_LOW_BID_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_CAMPAIGN_STATUS_ENABLED_DETAIL             SponsoredProductsKeywordServingStatusReason = "CAMPAIGN_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_CAMPAIGN_PAUSED_DETAIL                     SponsoredProductsKeywordServingStatusReason = "CAMPAIGN_PAUSED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_CAMPAIGN_ARCHIVED_DETAIL                   SponsoredProductsKeywordServingStatusReason = "CAMPAIGN_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PENDING_REVIEW_DETAIL                      SponsoredProductsKeywordServingStatusReason = "PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_REJECTED_DETAIL                            SponsoredProductsKeywordServingStatusReason = "REJECTED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PENDING_START_DATE_DETAIL                  SponsoredProductsKeywordServingStatusReason = "PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ENDED_DETAIL                               SponsoredProductsKeywordServingStatusReason = "ENDED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_DETAIL              SponsoredProductsKeywordServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_DETAIL                 SponsoredProductsKeywordServingStatusReason = "CAMPAIGN_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PORTFOLIO_STATUS_ENABLED_DETAIL            SponsoredProductsKeywordServingStatusReason = "PORTFOLIO_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PORTFOLIO_PAUSED_DETAIL                    SponsoredProductsKeywordServingStatusReason = "PORTFOLIO_PAUSED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PORTFOLIO_ARCHIVED_DETAIL                  SponsoredProductsKeywordServingStatusReason = "PORTFOLIO_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PORTFOLIO_OUT_OF_BUDGET_DETAIL             SponsoredProductsKeywordServingStatusReason = "PORTFOLIO_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PORTFOLIO_PENDING_START_DATE_DETAIL        SponsoredProductsKeywordServingStatusReason = "PORTFOLIO_PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_PORTFOLIO_ENDED_DETAIL                     SponsoredProductsKeywordServingStatusReason = "PORTFOLIO_ENDED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_DETAIL       SponsoredProductsKeywordServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_DETAIL  SponsoredProductsKeywordServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ADVERTISER_ARCHIVED_DETAIL                 SponsoredProductsKeywordServingStatusReason = "ADVERTISER_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ADVERTISER_PAUSED_DETAIL                   SponsoredProductsKeywordServingStatusReason = "ADVERTISER_PAUSED_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL            SponsoredProductsKeywordServingStatusReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL          SponsoredProductsKeywordServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_ACCOUNT_OUT_OF_BUDGET_DETAIL               SponsoredProductsKeywordServingStatusReason = "ACCOUNT_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUSREASON_OTHER                                      SponsoredProductsKeywordServingStatusReason = "OTHER"
)

// All allowed values of SponsoredProductsKeywordServingStatusReason enum
var AllowedSponsoredProductsKeywordServingStatusReasonEnumValues = []SponsoredProductsKeywordServingStatusReason{
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
	"PORTFOLIO_STATUS_ENABLED_DETAIL",
	"PORTFOLIO_PAUSED_DETAIL",
	"PORTFOLIO_ARCHIVED_DETAIL",
	"PORTFOLIO_OUT_OF_BUDGET_DETAIL",
	"PORTFOLIO_PENDING_START_DATE_DETAIL",
	"PORTFOLIO_ENDED_DETAIL",
	"ADVERTISER_POLICING_SUSPENDED_DETAIL",
	"ADVERTISER_POLICING_PENDING_REVIEW_DETAIL",
	"ADVERTISER_ARCHIVED_DETAIL",
	"ADVERTISER_PAUSED_DETAIL",
	"ADVERTISER_OUT_OF_BUDGET_DETAIL",
	"ADVERTISER_PAYMENT_FAILURE_DETAIL",
	"ACCOUNT_OUT_OF_BUDGET_DETAIL",
	"OTHER",
}

func (v *SponsoredProductsKeywordServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsKeywordServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsKeywordServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsKeywordServingStatusReason", value)
}

// NewSponsoredProductsKeywordServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsKeywordServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsKeywordServingStatusReasonFromValue(v string) (*SponsoredProductsKeywordServingStatusReason, error) {
	ev := SponsoredProductsKeywordServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsKeywordServingStatusReason: valid values are %v", v, AllowedSponsoredProductsKeywordServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsKeywordServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsKeywordServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsKeywordServingStatusReason value
func (v SponsoredProductsKeywordServingStatusReason) Ptr() *SponsoredProductsKeywordServingStatusReason {
	return &v
}

type NullableSponsoredProductsKeywordServingStatusReason struct {
	value *SponsoredProductsKeywordServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsKeywordServingStatusReason) Get() *SponsoredProductsKeywordServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsKeywordServingStatusReason) Set(val *SponsoredProductsKeywordServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordServingStatusReason(val *SponsoredProductsKeywordServingStatusReason) *NullableSponsoredProductsKeywordServingStatusReason {
	return &NullableSponsoredProductsKeywordServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
