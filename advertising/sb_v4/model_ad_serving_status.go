package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AdServingStatus The ad serving status determined by system. - AD_STATUS_LIVE - Ad's status is enabled. - AD_POLICING_PENDING_REVIEW - Ad is pending review because of policing reason. - AD_POLICING_SUSPENDED - Ad is suspended review because of policing reason. - AD_PAUSED - Ad's status is paused. - AD_ARCHIVED - Ad's status is archived.  - AD_GROUP_STATUS_ENABLED - Ad group's (parent) status is enabled. - AD_GROUP_PAUSED - Ad group's (parent) status is paused. - AD_GROUP_ARCHIVED - Ad group's (parent) status is archived. - AD_GROUP_INCOMPLETE - Ad group (parent) does not contain any ads or targeting clauses. - AD_GROUP_POLICING_PENDING_REVIEW - Ad group is pending review because of policing reason - AD_GROUP_POLICING_CREATIVE_REJECTED - Ad group is rejected due to creative because of policing reason - AD_GROUP_LOW_BID - Ad group is less than the minimum allowed bid in its marketplace  - ADVERTISER_STATUS_ENABLED - Advertiser's status is enabled - ADVERTISER_POLICING_PENDING_REVIEW - Avertiser is pending review because of policing reason - ADVERTISER_POLICING_SUSPENDED - Advertiser's status is suspended because of policing reason - ADVERTISER_PAUSED - Advertiser's status is paused - ADVERTISER_ARCHIVED - Advertiser's status is archived - ADVERTISER_PAYMENT_FAILURE - Advertiser's internal status is suspended - ADVERTISER_ACCOUNT_OUT_OF_BUDGET - Advertiser is out of budget for all Sponsored Ads campaigns - ADVERTISER_OUT_OF_PREPAY_BALANCE - Advertiser is out of prepay balance for all Sponsored Ads campaigns - ADVERTISER_EXCEED_SPENDS_LIMIT - Advertiser spends over the daily limit  - CAMPAIGN_STATUS_ENABLED - Campaign's (parent) status is enabled. - CAMPAIGN_PAUSED - Campaign's (parent) status is paused. - CAMPAIGN_ARCHIVED - Campaign's (parent) status is archived. - CAMPAIGN_INCOMPLETE - Campaign (parent) does not contain any ads or targeting clauses. - CAMPAIGN_OUT_OF_BUDGET - Campaign (parent) is out of budget.  - PORTFOLIO_STATUS_ENABLED - Portfolio's (parent) status is enabled - PORTFOLIO_PAUSED - Portfolio's (parent) status is paused - PORTFOLIO_ARCHIVED - Portfolio's (parent) status is archived - PORTFOLIO_OUT_OF_BUDGET - Portfolio (parent) is out of budget - PORTFOLIO_PENDING_START_DATE - Portfolio's (parent) start date is in the future - PORTFOLIO_ENDED - Portfolio's (parent) end date is in the past.  - INELIGIBLE - Ad is ineligible. - ELIGIBLE  - Ad is eligible. - ENDED - Campaign's (parent) end date is in the past. - PENDING_REVIEW - Campaign (parent) is pending review. - PENDING_START_DATE - Campaign's (parent) start date is in the future. - REJECTED - Campaign (parent) is rejected by moderation process. - UNKNOWN - Serving status is unknown. Please contact us for support.
type AdServingStatus string

// List of AdServingStatus
const (
	ADSERVINGSTATUS_AD_STATUS_LIVE                      AdServingStatus = "AD_STATUS_LIVE"
	ADSERVINGSTATUS_AD_POLICING_PENDING_REVIEW          AdServingStatus = "AD_POLICING_PENDING_REVIEW"
	ADSERVINGSTATUS_AD_POLICING_SUSPENDED               AdServingStatus = "AD_POLICING_SUSPENDED"
	ADSERVINGSTATUS_AD_PAUSED                           AdServingStatus = "AD_PAUSED"
	ADSERVINGSTATUS_AD_ARCHIVED                         AdServingStatus = "AD_ARCHIVED"
	ADSERVINGSTATUS_AD_GROUP_STATUS_ENABLED             AdServingStatus = "AD_GROUP_STATUS_ENABLED"
	ADSERVINGSTATUS_AD_GROUP_PAUSED                     AdServingStatus = "AD_GROUP_PAUSED"
	ADSERVINGSTATUS_AD_GROUP_ARCHIVED                   AdServingStatus = "AD_GROUP_ARCHIVED"
	ADSERVINGSTATUS_AD_GROUP_INCOMPLETE                 AdServingStatus = "AD_GROUP_INCOMPLETE"
	ADSERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    AdServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	ADSERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED AdServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	ADSERVINGSTATUS_AD_GROUP_LOW_BID                    AdServingStatus = "AD_GROUP_LOW_BID"
	ADSERVINGSTATUS_ADVERTISER_STATUS_ENABLED           AdServingStatus = "ADVERTISER_STATUS_ENABLED"
	ADSERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  AdServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	ADSERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       AdServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	ADSERVINGSTATUS_ADVERTISER_PAUSED                   AdServingStatus = "ADVERTISER_PAUSED"
	ADSERVINGSTATUS_ADVERTISER_ARCHIVED                 AdServingStatus = "ADVERTISER_ARCHIVED"
	ADSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          AdServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	ADSERVINGSTATUS_ADVERTISER_ACCOUNT_OUT_OF_BUDGET    AdServingStatus = "ADVERTISER_ACCOUNT_OUT_OF_BUDGET"
	ADSERVINGSTATUS_ADVERTISER_OUT_OF_PREPAY_BALANCE    AdServingStatus = "ADVERTISER_OUT_OF_PREPAY_BALANCE"
	ADSERVINGSTATUS_ADVERTISER_EXCEED_SPENDS_LIMIT      AdServingStatus = "ADVERTISER_EXCEED_SPENDS_LIMIT"
	ADSERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             AdServingStatus = "CAMPAIGN_STATUS_ENABLED"
	ADSERVINGSTATUS_CAMPAIGN_PAUSED                     AdServingStatus = "CAMPAIGN_PAUSED"
	ADSERVINGSTATUS_CAMPAIGN_ARCHIVED                   AdServingStatus = "CAMPAIGN_ARCHIVED"
	ADSERVINGSTATUS_CAMPAIGN_INCOMPLETE                 AdServingStatus = "CAMPAIGN_INCOMPLETE"
	ADSERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              AdServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	ADSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED            AdServingStatus = "PORTFOLIO_STATUS_ENABLED"
	ADSERVINGSTATUS_PORTFOLIO_PAUSED                    AdServingStatus = "PORTFOLIO_PAUSED"
	ADSERVINGSTATUS_PORTFOLIO_ARCHIVED                  AdServingStatus = "PORTFOLIO_ARCHIVED"
	ADSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET             AdServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	ADSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE        AdServingStatus = "PORTFOLIO_PENDING_START_DATE"
	ADSERVINGSTATUS_PORTFOLIO_ENDED                     AdServingStatus = "PORTFOLIO_ENDED"
	ADSERVINGSTATUS_INELIGIBLE                          AdServingStatus = "INELIGIBLE"
	ADSERVINGSTATUS_ELIGIBLE                            AdServingStatus = "ELIGIBLE"
	ADSERVINGSTATUS_ENDED                               AdServingStatus = "ENDED"
	ADSERVINGSTATUS_PENDING_REVIEW                      AdServingStatus = "PENDING_REVIEW"
	ADSERVINGSTATUS_PENDING_START_DATE                  AdServingStatus = "PENDING_START_DATE"
	ADSERVINGSTATUS_REJECTED                            AdServingStatus = "REJECTED"
	ADSERVINGSTATUS_UNKNOWN                             AdServingStatus = "UNKNOWN"
)

// All allowed values of AdServingStatus enum
var AllowedAdServingStatusEnumValues = []AdServingStatus{
	"AD_STATUS_LIVE",
	"AD_POLICING_PENDING_REVIEW",
	"AD_POLICING_SUSPENDED",
	"AD_PAUSED",
	"AD_ARCHIVED",
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

func (v *AdServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdServingStatus(value)
	for _, existing := range AllowedAdServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdServingStatus", value)
}

// NewAdServingStatusFromValue returns a pointer to a valid AdServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdServingStatusFromValue(v string) (*AdServingStatus, error) {
	ev := AdServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdServingStatus: valid values are %v", v, AllowedAdServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdServingStatus) IsValid() bool {
	for _, existing := range AllowedAdServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdServingStatus value
func (v AdServingStatus) Ptr() *AdServingStatus {
	return &v
}

type NullableAdServingStatus struct {
	value *AdServingStatus
	isSet bool
}

func (v NullableAdServingStatus) Get() *AdServingStatus {
	return v.value
}

func (v *NullableAdServingStatus) Set(val *AdServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAdServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAdServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdServingStatus(val *AdServingStatus) *NullableAdServingStatus {
	return &NullableAdServingStatus{value: val, isSet: true}
}

func (v NullableAdServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
