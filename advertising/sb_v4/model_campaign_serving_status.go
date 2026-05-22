package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CampaignServingStatus `Notice: the servingStatus enums have not been finalized yet.` The campaign serving status determined by system. - ADVERTISER_STATUS_ENABLED - Advertiser's status is enabled - ADVERTISER_POLICING_PENDING_REVIEW - Avertiser is pending review because of policing reason - ADVERTISER_POLICING_SUSPENDED - Advertiser's status is suspended because of policing reason - ADVERTISER_PAUSED - Advertiser's status is paused - ADVERTISER_ARCHIVED - Advertiser's status is archived - ADVERTISER_PAYMENT_FAILURE - Advertiser's internal status is suspended - ADVERTISER_ACCOUNT_OUT_OF_BUDGET - Advertiser is out of budget for all Sponsored Ads campaigns - ADVERTISER_OUT_OF_PREPAY_BALANCE - Advertiser is out of prepay balance for all Sponsored Ads campaigns - ADVERTISER_EXCEED_SPENDS_LIMIT - Advertiser spends over the daily limit  - CAMPAIGN_STATUS_ENABLED - Campaign's status is enabled. - CAMPAIGN_PAUSED - Campaign's status is paused. - CAMPAIGN_ARCHIVED - Campaign's status is archived. - CAMPAIGN_INCOMPLETE - Campaign does not contain any ads or targeting clauses. - CAMPAIGN_OUT_OF_BUDGET - Campaign is out of budget.  - PORTFOLIO_STATUS_ENABLED - Portfolio's status is enabled - PORTFOLIO_PAUSED - Portfolio's status is paused - PORTFOLIO_ARCHIVED - Portfolio's status is archived - PORTFOLIO_OUT_OF_BUDGET - Portfolio is out of budget - PORTFOLIO_PENDING_START_DATE - Portfolio's start date is in the future - PORTFOLIO_ENDED - Portfolio's end date is in the past.  - INELIGIBLE - Ad Offer is ineligible - ELIGIBLE - Ad Offer is eligible - ENDED - Campaign's end date is in the past. - PENDING_REVIEW - Campaign is pending review. - PENDING_START_DATE - Campaign's start date is in the future. - REJECTED - Campaign is rejected by moderation process. - UNKNOWN - Serving status is unknown. Please contact us for support.
type CampaignServingStatus string

// List of CampaignServingStatus
const (
	CAMPAIGNSERVINGSTATUS_ADVERTISER_STATUS_ENABLED          CampaignServingStatus = "ADVERTISER_STATUS_ENABLED"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW CampaignServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED      CampaignServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_PAUSED                  CampaignServingStatus = "ADVERTISER_PAUSED"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_ARCHIVED                CampaignServingStatus = "ADVERTISER_ARCHIVED"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE         CampaignServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_ACCOUNT_OUT_OF_BUDGET   CampaignServingStatus = "ADVERTISER_ACCOUNT_OUT_OF_BUDGET"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_OUT_OF_PREPAY_BALANCE   CampaignServingStatus = "ADVERTISER_OUT_OF_PREPAY_BALANCE"
	CAMPAIGNSERVINGSTATUS_ADVERTISER_EXCEED_SPENDS_LIMIT     CampaignServingStatus = "ADVERTISER_EXCEED_SPENDS_LIMIT"
	CAMPAIGNSERVINGSTATUS_CAMPAIGN_STATUS_ENABLED            CampaignServingStatus = "CAMPAIGN_STATUS_ENABLED"
	CAMPAIGNSERVINGSTATUS_CAMPAIGN_PAUSED                    CampaignServingStatus = "CAMPAIGN_PAUSED"
	CAMPAIGNSERVINGSTATUS_CAMPAIGN_ARCHIVED                  CampaignServingStatus = "CAMPAIGN_ARCHIVED"
	CAMPAIGNSERVINGSTATUS_CAMPAIGN_INCOMPLETE                CampaignServingStatus = "CAMPAIGN_INCOMPLETE"
	CAMPAIGNSERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET             CampaignServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	CAMPAIGNSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED           CampaignServingStatus = "PORTFOLIO_STATUS_ENABLED"
	CAMPAIGNSERVINGSTATUS_PORTFOLIO_PAUSED                   CampaignServingStatus = "PORTFOLIO_PAUSED"
	CAMPAIGNSERVINGSTATUS_PORTFOLIO_ARCHIVED                 CampaignServingStatus = "PORTFOLIO_ARCHIVED"
	CAMPAIGNSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET            CampaignServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	CAMPAIGNSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE       CampaignServingStatus = "PORTFOLIO_PENDING_START_DATE"
	CAMPAIGNSERVINGSTATUS_PORTFOLIO_ENDED                    CampaignServingStatus = "PORTFOLIO_ENDED"
	CAMPAIGNSERVINGSTATUS_INELIGIBLE                         CampaignServingStatus = "INELIGIBLE"
	CAMPAIGNSERVINGSTATUS_ELIGIBLE                           CampaignServingStatus = "ELIGIBLE"
	CAMPAIGNSERVINGSTATUS_ENDED                              CampaignServingStatus = "ENDED"
	CAMPAIGNSERVINGSTATUS_PENDING_REVIEW                     CampaignServingStatus = "PENDING_REVIEW"
	CAMPAIGNSERVINGSTATUS_PENDING_START_DATE                 CampaignServingStatus = "PENDING_START_DATE"
	CAMPAIGNSERVINGSTATUS_REJECTED                           CampaignServingStatus = "REJECTED"
	CAMPAIGNSERVINGSTATUS_UNKNOWN                            CampaignServingStatus = "UNKNOWN"
)

// All allowed values of CampaignServingStatus enum
var AllowedCampaignServingStatusEnumValues = []CampaignServingStatus{
	"ADVERTISER_STATUS_ENABLED",
	"ADVERTISER_POLICING_PENDING_REVIEW",
	"ADVERTISER_POLICING_SUSPENDED",
	"ADVERTISER_PAUSED",
	"ADVERTISER_ARCHIVED",
	"ADVERTISER_PAYMENT_FAILURE",
	"ADVERTISER_ACCOUNT_OUT_OF_BUDGET",
	"ADVERTISER_OUT_OF_PREPAY_BALANCE",
	"ADVERTISER_EXCEED_SPENDS_LIMIT",
	"CAMPAIGN_STATUS_ENABLED",
	"CAMPAIGN_PAUSED",
	"CAMPAIGN_ARCHIVED",
	"CAMPAIGN_INCOMPLETE",
	"CAMPAIGN_OUT_OF_BUDGET",
	"PORTFOLIO_STATUS_ENABLED",
	"PORTFOLIO_PAUSED",
	"PORTFOLIO_ARCHIVED",
	"PORTFOLIO_OUT_OF_BUDGET",
	"PORTFOLIO_PENDING_START_DATE",
	"PORTFOLIO_ENDED",
	"INELIGIBLE",
	"ELIGIBLE",
	"ENDED",
	"PENDING_REVIEW",
	"PENDING_START_DATE",
	"REJECTED",
	"UNKNOWN",
}

func (v *CampaignServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignServingStatus(value)
	for _, existing := range AllowedCampaignServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignServingStatus", value)
}

// NewCampaignServingStatusFromValue returns a pointer to a valid CampaignServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignServingStatusFromValue(v string) (*CampaignServingStatus, error) {
	ev := CampaignServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignServingStatus: valid values are %v", v, AllowedCampaignServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignServingStatus) IsValid() bool {
	for _, existing := range AllowedCampaignServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignServingStatus value
func (v CampaignServingStatus) Ptr() *CampaignServingStatus {
	return &v
}

type NullableCampaignServingStatus struct {
	value *CampaignServingStatus
	isSet bool
}

func (v NullableCampaignServingStatus) Get() *CampaignServingStatus {
	return v.value
}

func (v *NullableCampaignServingStatus) Set(val *CampaignServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignServingStatus(val *CampaignServingStatus) *NullableCampaignServingStatus {
	return &NullableCampaignServingStatus{value: val, isSet: true}
}

func (v NullableCampaignServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
