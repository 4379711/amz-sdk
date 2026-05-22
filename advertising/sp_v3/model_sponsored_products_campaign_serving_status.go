package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCampaignServingStatus the model 'SponsoredProductsCampaignServingStatus'
type SponsoredProductsCampaignServingStatus string

// List of SponsoredProductsCampaignServingStatus
const (
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_CAMPAIGN_STATUS_ENABLED            SponsoredProductsCampaignServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_CAMPAIGN_PAUSED                    SponsoredProductsCampaignServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_CAMPAIGN_ARCHIVED                  SponsoredProductsCampaignServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PENDING_REVIEW                     SponsoredProductsCampaignServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_REJECTED                           SponsoredProductsCampaignServingStatus = "REJECTED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PENDING_START_DATE                 SponsoredProductsCampaignServingStatus = "PENDING_START_DATE"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ENDED                              SponsoredProductsCampaignServingStatus = "ENDED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET             SponsoredProductsCampaignServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_CAMPAIGN_INCOMPLETE                SponsoredProductsCampaignServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED           SponsoredProductsCampaignServingStatus = "PORTFOLIO_STATUS_ENABLED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PORTFOLIO_PAUSED                   SponsoredProductsCampaignServingStatus = "PORTFOLIO_PAUSED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PORTFOLIO_ARCHIVED                 SponsoredProductsCampaignServingStatus = "PORTFOLIO_ARCHIVED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET            SponsoredProductsCampaignServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE       SponsoredProductsCampaignServingStatus = "PORTFOLIO_PENDING_START_DATE"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_PORTFOLIO_ENDED                    SponsoredProductsCampaignServingStatus = "PORTFOLIO_ENDED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED      SponsoredProductsCampaignServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW SponsoredProductsCampaignServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ADVERTISER_ARCHIVED                SponsoredProductsCampaignServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ADVERTISER_PAUSED                  SponsoredProductsCampaignServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ADVERTISER_OUT_OF_BUDGET           SponsoredProductsCampaignServingStatus = "ADVERTISER_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE         SponsoredProductsCampaignServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_ACCOUNT_OUT_OF_BUDGET              SponsoredProductsCampaignServingStatus = "ACCOUNT_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSCAMPAIGNSERVINGSTATUS_OTHER                              SponsoredProductsCampaignServingStatus = "OTHER"
)

// All allowed values of SponsoredProductsCampaignServingStatus enum
var AllowedSponsoredProductsCampaignServingStatusEnumValues = []SponsoredProductsCampaignServingStatus{
	"CAMPAIGN_STATUS_ENABLED",
	"CAMPAIGN_PAUSED",
	"CAMPAIGN_ARCHIVED",
	"PENDING_REVIEW",
	"REJECTED",
	"PENDING_START_DATE",
	"ENDED",
	"CAMPAIGN_OUT_OF_BUDGET",
	"CAMPAIGN_INCOMPLETE",
	"PORTFOLIO_STATUS_ENABLED",
	"PORTFOLIO_PAUSED",
	"PORTFOLIO_ARCHIVED",
	"PORTFOLIO_OUT_OF_BUDGET",
	"PORTFOLIO_PENDING_START_DATE",
	"PORTFOLIO_ENDED",
	"ADVERTISER_POLICING_SUSPENDED",
	"ADVERTISER_POLICING_PENDING_REVIEW",
	"ADVERTISER_ARCHIVED",
	"ADVERTISER_PAUSED",
	"ADVERTISER_OUT_OF_BUDGET",
	"ADVERTISER_PAYMENT_FAILURE",
	"ACCOUNT_OUT_OF_BUDGET",
	"OTHER",
}

func (v *SponsoredProductsCampaignServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCampaignServingStatus(value)
	for _, existing := range AllowedSponsoredProductsCampaignServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCampaignServingStatus", value)
}

// NewSponsoredProductsCampaignServingStatusFromValue returns a pointer to a valid SponsoredProductsCampaignServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCampaignServingStatusFromValue(v string) (*SponsoredProductsCampaignServingStatus, error) {
	ev := SponsoredProductsCampaignServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCampaignServingStatus: valid values are %v", v, AllowedSponsoredProductsCampaignServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCampaignServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCampaignServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCampaignServingStatus value
func (v SponsoredProductsCampaignServingStatus) Ptr() *SponsoredProductsCampaignServingStatus {
	return &v
}

type NullableSponsoredProductsCampaignServingStatus struct {
	value *SponsoredProductsCampaignServingStatus
	isSet bool
}

func (v NullableSponsoredProductsCampaignServingStatus) Get() *SponsoredProductsCampaignServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsCampaignServingStatus) Set(val *SponsoredProductsCampaignServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignServingStatus(val *SponsoredProductsCampaignServingStatus) *NullableSponsoredProductsCampaignServingStatus {
	return &NullableSponsoredProductsCampaignServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
