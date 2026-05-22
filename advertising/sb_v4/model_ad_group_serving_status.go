package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AdGroupServingStatus `Notice: the servingStatus enums have not been finalized yet.` The ad group serving status determined by system. - AD_GROUP_STATUS_ENABLED - Ad group's status is enabled. - AD_GROUP_PAUSED - Ad group's status is paused. - AD_GROUP_ARCHIVED - Ad group's status is archived. - AD_GROUP_INCOMPLETE - Ad group does not contain any ads or targeting clauses. - AD_GROUP_POLICING_PENDING_REVIEW - Ad group is pending review because of policing reason - AD_GROUP_POLICING_CREATIVE_REJECTED - Ad group is rejected due to creative because of policing reason - AD_GROUP_LOW_BID - Ad group is less than the minimum allowed bid in its marketplace  - ADVERTISER_STATUS_ENABLED - Advertiser's status is enabled - ADVERTISER_POLICING_PENDING_REVIEW - Avertiser is pending review because of policing reason - ADVERTISER_POLICING_SUSPENDED - Advertiser's status is suspended because of policing reason - ADVERTISER_PAUSED - Advertiser's status is paused - ADVERTISER_ARCHIVED - Advertiser's status is archived - ADVERTISER_PAYMENT_FAILURE - Advertiser's internal status is suspended - ADVERTISER_ACCOUNT_OUT_OF_BUDGET - Advertiser is out of budget for all Sponsored Ads campaigns - ADVERTISER_OUT_OF_PREPAY_BALANCE - Advertiser is out of prepay balance for all Sponsored Ads campaigns - ADVERTISER_EXCEED_SPENDS_LIMIT - Advertiser spends over the daily limit  - CAMPAIGN_STATUS_ENABLED - Campaign's (parent) status is enabled. - CAMPAIGN_PAUSED - Campaign's (parent) status is paused. - CAMPAIGN_ARCHIVED - Campaign's (parent) status is archived. - CAMPAIGN_INCOMPLETE - Campaign (parent) does not contain any ads or targeting clauses. - CAMPAIGN_OUT_OF_BUDGET - Campaign (parent) is out of budget.  - PORTFOLIO_STATUS_ENABLED - Portfolio's (parent) status is enabled - PORTFOLIO_PAUSED - Portfolio's (parent) status is paused - PORTFOLIO_ARCHIVED - Portfolio's (parent) status is archived - PORTFOLIO_OUT_OF_BUDGET - Portfolio (parent) is out of budget - PORTFOLIO_PENDING_START_DATE - Portfolio's (parent) start date is in the future - PORTFOLIO_ENDED - Portfolio's (parent) end date is in the past.  - INELIGIBLE - Ad group is ineligible. - ELIGIBLE - Ad group is eligible. - ENDED - Campaign's (parent) end date is in the past. - PENDING_REVIEW - Campaign (parent) is pending review. - PENDING_START_DATE - Campaign's (parent) start date is in the future. - REJECTED - Campaign (parent) is rejected by moderation process. - UNKNOWN - Serving status is unknown. Please contact us for support.
type AdGroupServingStatus string

// List of AdGroupServingStatus
const (
	ADGROUPSERVINGSTATUS_AD_GROUP_STATUS_ENABLED             AdGroupServingStatus = "AD_GROUP_STATUS_ENABLED"
	ADGROUPSERVINGSTATUS_AD_GROUP_PAUSED                     AdGroupServingStatus = "AD_GROUP_PAUSED"
	ADGROUPSERVINGSTATUS_AD_GROUP_ARCHIVED                   AdGroupServingStatus = "AD_GROUP_ARCHIVED"
	ADGROUPSERVINGSTATUS_AD_GROUP_INCOMPLETE                 AdGroupServingStatus = "AD_GROUP_INCOMPLETE"
	ADGROUPSERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    AdGroupServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	ADGROUPSERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED AdGroupServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	ADGROUPSERVINGSTATUS_AD_GROUP_LOW_BID                    AdGroupServingStatus = "AD_GROUP_LOW_BID"
	ADGROUPSERVINGSTATUS_ADVERTISER_STATUS_ENABLED           AdGroupServingStatus = "ADVERTISER_STATUS_ENABLED"
	ADGROUPSERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  AdGroupServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	ADGROUPSERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       AdGroupServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	ADGROUPSERVINGSTATUS_ADVERTISER_PAUSED                   AdGroupServingStatus = "ADVERTISER_PAUSED"
	ADGROUPSERVINGSTATUS_ADVERTISER_ARCHIVED                 AdGroupServingStatus = "ADVERTISER_ARCHIVED"
	ADGROUPSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          AdGroupServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	ADGROUPSERVINGSTATUS_ADVERTISER_ACCOUNT_OUT_OF_BUDGET    AdGroupServingStatus = "ADVERTISER_ACCOUNT_OUT_OF_BUDGET"
	ADGROUPSERVINGSTATUS_ADVERTISER_OUT_OF_PREPAY_BALANCE    AdGroupServingStatus = "ADVERTISER_OUT_OF_PREPAY_BALANCE"
	ADGROUPSERVINGSTATUS_ADVERTISER_EXCEED_SPENDS_LIMIT      AdGroupServingStatus = "ADVERTISER_EXCEED_SPENDS_LIMIT"
	ADGROUPSERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             AdGroupServingStatus = "CAMPAIGN_STATUS_ENABLED"
	ADGROUPSERVINGSTATUS_CAMPAIGN_PAUSED                     AdGroupServingStatus = "CAMPAIGN_PAUSED"
	ADGROUPSERVINGSTATUS_CAMPAIGN_ARCHIVED                   AdGroupServingStatus = "CAMPAIGN_ARCHIVED"
	ADGROUPSERVINGSTATUS_CAMPAIGN_INCOMPLETE                 AdGroupServingStatus = "CAMPAIGN_INCOMPLETE"
	ADGROUPSERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              AdGroupServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	ADGROUPSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED            AdGroupServingStatus = "PORTFOLIO_STATUS_ENABLED"
	ADGROUPSERVINGSTATUS_PORTFOLIO_PAUSED                    AdGroupServingStatus = "PORTFOLIO_PAUSED"
	ADGROUPSERVINGSTATUS_PORTFOLIO_ARCHIVED                  AdGroupServingStatus = "PORTFOLIO_ARCHIVED"
	ADGROUPSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET             AdGroupServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	ADGROUPSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE        AdGroupServingStatus = "PORTFOLIO_PENDING_START_DATE"
	ADGROUPSERVINGSTATUS_PORTFOLIO_ENDED                     AdGroupServingStatus = "PORTFOLIO_ENDED"
	ADGROUPSERVINGSTATUS_INELIGIBLE                          AdGroupServingStatus = "INELIGIBLE"
	ADGROUPSERVINGSTATUS_ELIGIBLE                            AdGroupServingStatus = "ELIGIBLE"
	ADGROUPSERVINGSTATUS_ENDED                               AdGroupServingStatus = "ENDED"
	ADGROUPSERVINGSTATUS_PENDING_REVIEW                      AdGroupServingStatus = "PENDING_REVIEW"
	ADGROUPSERVINGSTATUS_PENDING_START_DATE                  AdGroupServingStatus = "PENDING_START_DATE"
	ADGROUPSERVINGSTATUS_REJECTED                            AdGroupServingStatus = "REJECTED"
	ADGROUPSERVINGSTATUS_UNKNOWN                             AdGroupServingStatus = "UNKNOWN"
)

// All allowed values of AdGroupServingStatus enum
var AllowedAdGroupServingStatusEnumValues = []AdGroupServingStatus{
	"AD_GROUP_STATUS_ENABLED",
	"AD_GROUP_PAUSED",
	"AD_GROUP_ARCHIVED",
	"AD_GROUP_INCOMPLETE",
	"AD_GROUP_POLICING_PENDING_REVIEW",
	"AD_GROUP_POLICING_CREATIVE_REJECTED",
	"AD_GROUP_LOW_BID",
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

func (v *AdGroupServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdGroupServingStatus(value)
	for _, existing := range AllowedAdGroupServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdGroupServingStatus", value)
}

// NewAdGroupServingStatusFromValue returns a pointer to a valid AdGroupServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGroupServingStatusFromValue(v string) (*AdGroupServingStatus, error) {
	ev := AdGroupServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGroupServingStatus: valid values are %v", v, AllowedAdGroupServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGroupServingStatus) IsValid() bool {
	for _, existing := range AllowedAdGroupServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdGroupServingStatus value
func (v AdGroupServingStatus) Ptr() *AdGroupServingStatus {
	return &v
}

type NullableAdGroupServingStatus struct {
	value *AdGroupServingStatus
	isSet bool
}

func (v NullableAdGroupServingStatus) Get() *AdGroupServingStatus {
	return v.value
}

func (v *NullableAdGroupServingStatus) Set(val *AdGroupServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupServingStatus(val *AdGroupServingStatus) *NullableAdGroupServingStatus {
	return &NullableAdGroupServingStatus{value: val, isSet: true}
}

func (v NullableAdGroupServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
