package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCampaignServingStatusReason the model 'SponsoredProductsCampaignServingStatusReason'
type SponsoredProductsCampaignServingStatusReason string

// List of SponsoredProductsCampaignServingStatusReason
const (
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_STATUS_ENABLED_DETAIL            SponsoredProductsCampaignServingStatusReason = "CAMPAIGN_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_PAUSED_DETAIL                    SponsoredProductsCampaignServingStatusReason = "CAMPAIGN_PAUSED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_ARCHIVED_DETAIL                  SponsoredProductsCampaignServingStatusReason = "CAMPAIGN_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PENDING_REVIEW_DETAIL                     SponsoredProductsCampaignServingStatusReason = "PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_REJECTED_DETAIL                           SponsoredProductsCampaignServingStatusReason = "REJECTED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PENDING_START_DATE_DETAIL                 SponsoredProductsCampaignServingStatusReason = "PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ENDED_DETAIL                              SponsoredProductsCampaignServingStatusReason = "ENDED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_DETAIL             SponsoredProductsCampaignServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_DETAIL                SponsoredProductsCampaignServingStatusReason = "CAMPAIGN_INCOMPLETE_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PORTFOLIO_STATUS_ENABLED_DETAIL           SponsoredProductsCampaignServingStatusReason = "PORTFOLIO_STATUS_ENABLED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PORTFOLIO_PAUSED_DETAIL                   SponsoredProductsCampaignServingStatusReason = "PORTFOLIO_PAUSED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PORTFOLIO_ARCHIVED_DETAIL                 SponsoredProductsCampaignServingStatusReason = "PORTFOLIO_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PORTFOLIO_OUT_OF_BUDGET_DETAIL            SponsoredProductsCampaignServingStatusReason = "PORTFOLIO_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PORTFOLIO_PENDING_START_DATE_DETAIL       SponsoredProductsCampaignServingStatusReason = "PORTFOLIO_PENDING_START_DATE_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_PORTFOLIO_ENDED_DETAIL                    SponsoredProductsCampaignServingStatusReason = "PORTFOLIO_ENDED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_DETAIL      SponsoredProductsCampaignServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_DETAIL SponsoredProductsCampaignServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_ARCHIVED_DETAIL                SponsoredProductsCampaignServingStatusReason = "ADVERTISER_ARCHIVED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_PAUSED_DETAIL                  SponsoredProductsCampaignServingStatusReason = "ADVERTISER_PAUSED_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL           SponsoredProductsCampaignServingStatusReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL         SponsoredProductsCampaignServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_ACCOUNT_OUT_OF_BUDGET_DETAIL              SponsoredProductsCampaignServingStatusReason = "ACCOUNT_OUT_OF_BUDGET_DETAIL"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUSREASON_OTHER                                     SponsoredProductsCampaignServingStatusReason = "OTHER"
)

// All allowed values of SponsoredProductsCampaignServingStatusReason enum
var AllowedSponsoredProductsCampaignServingStatusReasonEnumValues = []SponsoredProductsCampaignServingStatusReason{
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

func (v *SponsoredProductsCampaignServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCampaignServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsCampaignServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCampaignServingStatusReason", value)
}

// NewSponsoredProductsCampaignServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsCampaignServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCampaignServingStatusReasonFromValue(v string) (*SponsoredProductsCampaignServingStatusReason, error) {
	ev := SponsoredProductsCampaignServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCampaignServingStatusReason: valid values are %v", v, AllowedSponsoredProductsCampaignServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCampaignServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCampaignServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCampaignServingStatusReason value
func (v SponsoredProductsCampaignServingStatusReason) Ptr() *SponsoredProductsCampaignServingStatusReason {
	return &v
}

type NullableSponsoredProductsCampaignServingStatusReason struct {
	value *SponsoredProductsCampaignServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsCampaignServingStatusReason) Get() *SponsoredProductsCampaignServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsCampaignServingStatusReason) Set(val *SponsoredProductsCampaignServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignServingStatusReason(val *SponsoredProductsCampaignServingStatusReason) *NullableSponsoredProductsCampaignServingStatusReason {
	return &NullableSponsoredProductsCampaignServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
