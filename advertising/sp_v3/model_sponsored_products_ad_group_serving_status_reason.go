package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsAdGroupServingStatusReason the model 'SponsoredProductsAdGroupServingStatusReason'
type SponsoredProductsAdGroupServingStatusReason string

// List of SponsoredProductsAdGroupServingStatusReason
const (
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_AD_GROUP_STATUS_ENABLED_DETAIL             SponsoredProductsAdGroupServingStatusReason = "AD_GROUP_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_AD_GROUP_PAUSED_DETAIL                     SponsoredProductsAdGroupServingStatusReason = "AD_GROUP_PAUSED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_AD_GROUP_ARCHIVED_DETAIL                   SponsoredProductsAdGroupServingStatusReason = "AD_GROUP_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_AD_GROUP_INCOMPLETE_DETAIL                 SponsoredProductsAdGroupServingStatusReason = "AD_GROUP_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_AD_GROUP_POLICING_PENDING_REVIEW_DETAIL    SponsoredProductsAdGroupServingStatusReason = "AD_GROUP_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL SponsoredProductsAdGroupServingStatusReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_AD_GROUP_LOW_BID_DETAIL                    SponsoredProductsAdGroupServingStatusReason = "AD_GROUP_LOW_BID_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_CAMPAIGN_STATUS_ENABLED_DETAIL             SponsoredProductsAdGroupServingStatusReason = "CAMPAIGN_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_CAMPAIGN_PAUSED_DETAIL                     SponsoredProductsAdGroupServingStatusReason = "CAMPAIGN_PAUSED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_CAMPAIGN_ARCHIVED_DETAIL                   SponsoredProductsAdGroupServingStatusReason = "CAMPAIGN_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PENDING_REVIEW_DETAIL                      SponsoredProductsAdGroupServingStatusReason = "PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_REJECTED_DETAIL                            SponsoredProductsAdGroupServingStatusReason = "REJECTED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PENDING_START_DATE_DETAIL                  SponsoredProductsAdGroupServingStatusReason = "PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_ENDED_DETAIL                               SponsoredProductsAdGroupServingStatusReason = "ENDED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_DETAIL              SponsoredProductsAdGroupServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_DETAIL                 SponsoredProductsAdGroupServingStatusReason = "CAMPAIGN_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PORTFOLIO_STATUS_ENABLED_DETAIL            SponsoredProductsAdGroupServingStatusReason = "PORTFOLIO_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PORTFOLIO_PAUSED_DETAIL                    SponsoredProductsAdGroupServingStatusReason = "PORTFOLIO_PAUSED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PORTFOLIO_ARCHIVED_DETAIL                  SponsoredProductsAdGroupServingStatusReason = "PORTFOLIO_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PORTFOLIO_OUT_OF_BUDGET_DETAIL             SponsoredProductsAdGroupServingStatusReason = "PORTFOLIO_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PORTFOLIO_PENDING_START_DATE_DETAIL        SponsoredProductsAdGroupServingStatusReason = "PORTFOLIO_PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_PORTFOLIO_ENDED_DETAIL                     SponsoredProductsAdGroupServingStatusReason = "PORTFOLIO_ENDED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_DETAIL       SponsoredProductsAdGroupServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_DETAIL  SponsoredProductsAdGroupServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_ADVERTISER_ARCHIVED_DETAIL                 SponsoredProductsAdGroupServingStatusReason = "ADVERTISER_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_ADVERTISER_PAUSED_DETAIL                   SponsoredProductsAdGroupServingStatusReason = "ADVERTISER_PAUSED_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL            SponsoredProductsAdGroupServingStatusReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL          SponsoredProductsAdGroupServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUSREASON_OTHER                                      SponsoredProductsAdGroupServingStatusReason = "OTHER"
)

// All allowed values of SponsoredProductsAdGroupServingStatusReason enum
var AllowedSponsoredProductsAdGroupServingStatusReasonEnumValues = []SponsoredProductsAdGroupServingStatusReason{
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
	"OTHER",
}

func (v *SponsoredProductsAdGroupServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAdGroupServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsAdGroupServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAdGroupServingStatusReason", value)
}

// NewSponsoredProductsAdGroupServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsAdGroupServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAdGroupServingStatusReasonFromValue(v string) (*SponsoredProductsAdGroupServingStatusReason, error) {
	ev := SponsoredProductsAdGroupServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAdGroupServingStatusReason: valid values are %v", v, AllowedSponsoredProductsAdGroupServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAdGroupServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAdGroupServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAdGroupServingStatusReason value
func (v SponsoredProductsAdGroupServingStatusReason) Ptr() *SponsoredProductsAdGroupServingStatusReason {
	return &v
}

type NullableSponsoredProductsAdGroupServingStatusReason struct {
	value *SponsoredProductsAdGroupServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsAdGroupServingStatusReason) Get() *SponsoredProductsAdGroupServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupServingStatusReason) Set(val *SponsoredProductsAdGroupServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupServingStatusReason(val *SponsoredProductsAdGroupServingStatusReason) *NullableSponsoredProductsAdGroupServingStatusReason {
	return &NullableSponsoredProductsAdGroupServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
