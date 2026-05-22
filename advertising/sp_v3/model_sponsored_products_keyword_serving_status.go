package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsKeywordServingStatus the model 'SponsoredProductsKeywordServingStatus'
type SponsoredProductsKeywordServingStatus string

// List of SponsoredProductsKeywordServingStatus
const (
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_TARGETING_CLAUSE_STATUS_LIVE        SponsoredProductsKeywordServingStatus = "TARGETING_CLAUSE_STATUS_LIVE"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_TARGETING_CLAUSE_POLICING_SUSPENDED SponsoredProductsKeywordServingStatus = "TARGETING_CLAUSE_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_TARGETING_CLAUSE_PAUSED             SponsoredProductsKeywordServingStatus = "TARGETING_CLAUSE_PAUSED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_TARGETING_CLAUSE_ARCHIVED           SponsoredProductsKeywordServingStatus = "TARGETING_CLAUSE_ARCHIVED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_TARGETING_CLAUSE_BLOCKED            SponsoredProductsKeywordServingStatus = "TARGETING_CLAUSE_BLOCKED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_AD_GROUP_STATUS_ENABLED             SponsoredProductsKeywordServingStatus = "AD_GROUP_STATUS_ENABLED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_AD_GROUP_PAUSED                     SponsoredProductsKeywordServingStatus = "AD_GROUP_PAUSED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_AD_GROUP_ARCHIVED                   SponsoredProductsKeywordServingStatus = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_AD_GROUP_INCOMPLETE                 SponsoredProductsKeywordServingStatus = "AD_GROUP_INCOMPLETE"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    SponsoredProductsKeywordServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED SponsoredProductsKeywordServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_AD_GROUP_LOW_BID                    SponsoredProductsKeywordServingStatus = "AD_GROUP_LOW_BID"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             SponsoredProductsKeywordServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_CAMPAIGN_PAUSED                     SponsoredProductsKeywordServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_CAMPAIGN_ARCHIVED                   SponsoredProductsKeywordServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PENDING_REVIEW                      SponsoredProductsKeywordServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_REJECTED                            SponsoredProductsKeywordServingStatus = "REJECTED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PENDING_START_DATE                  SponsoredProductsKeywordServingStatus = "PENDING_START_DATE"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ENDED                               SponsoredProductsKeywordServingStatus = "ENDED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              SponsoredProductsKeywordServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_CAMPAIGN_INCOMPLETE                 SponsoredProductsKeywordServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED            SponsoredProductsKeywordServingStatus = "PORTFOLIO_STATUS_ENABLED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PORTFOLIO_PAUSED                    SponsoredProductsKeywordServingStatus = "PORTFOLIO_PAUSED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PORTFOLIO_ARCHIVED                  SponsoredProductsKeywordServingStatus = "PORTFOLIO_ARCHIVED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET             SponsoredProductsKeywordServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE        SponsoredProductsKeywordServingStatus = "PORTFOLIO_PENDING_START_DATE"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_PORTFOLIO_ENDED                     SponsoredProductsKeywordServingStatus = "PORTFOLIO_ENDED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       SponsoredProductsKeywordServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  SponsoredProductsKeywordServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ADVERTISER_ARCHIVED                 SponsoredProductsKeywordServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ADVERTISER_PAUSED                   SponsoredProductsKeywordServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ADVERTISER_OUT_OF_BUDGET            SponsoredProductsKeywordServingStatus = "ADVERTISER_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          SponsoredProductsKeywordServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_ACCOUNT_OUT_OF_BUDGET               SponsoredProductsKeywordServingStatus = "ACCOUNT_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSKEYWORDSERVINGSTATUS_OTHER                               SponsoredProductsKeywordServingStatus = "OTHER"
)

// All allowed values of SponsoredProductsKeywordServingStatus enum
var AllowedSponsoredProductsKeywordServingStatusEnumValues = []SponsoredProductsKeywordServingStatus{
	"TARGETING_CLAUSE_STATUS_LIVE",
	"TARGETING_CLAUSE_POLICING_SUSPENDED",
	"TARGETING_CLAUSE_PAUSED",
	"TARGETING_CLAUSE_ARCHIVED",
	"TARGETING_CLAUSE_BLOCKED",
	"AD_GROUP_STATUS_ENABLED",
	"AD_GROUP_PAUSED",
	"AD_GROUP_ARCHIVED",
	"AD_GROUP_INCOMPLETE",
	"AD_GROUP_POLICING_PENDING_REVIEW",
	"AD_GROUP_POLICING_CREATIVE_REJECTED",
	"AD_GROUP_LOW_BID",
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

func (v *SponsoredProductsKeywordServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsKeywordServingStatus(value)
	for _, existing := range AllowedSponsoredProductsKeywordServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsKeywordServingStatus", value)
}

// NewSponsoredProductsKeywordServingStatusFromValue returns a pointer to a valid SponsoredProductsKeywordServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsKeywordServingStatusFromValue(v string) (*SponsoredProductsKeywordServingStatus, error) {
	ev := SponsoredProductsKeywordServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsKeywordServingStatus: valid values are %v", v, AllowedSponsoredProductsKeywordServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsKeywordServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsKeywordServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsKeywordServingStatus value
func (v SponsoredProductsKeywordServingStatus) Ptr() *SponsoredProductsKeywordServingStatus {
	return &v
}

type NullableSponsoredProductsKeywordServingStatus struct {
	value *SponsoredProductsKeywordServingStatus
	isSet bool
}

func (v NullableSponsoredProductsKeywordServingStatus) Get() *SponsoredProductsKeywordServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsKeywordServingStatus) Set(val *SponsoredProductsKeywordServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordServingStatus(val *SponsoredProductsKeywordServingStatus) *NullableSponsoredProductsKeywordServingStatus {
	return &NullableSponsoredProductsKeywordServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
