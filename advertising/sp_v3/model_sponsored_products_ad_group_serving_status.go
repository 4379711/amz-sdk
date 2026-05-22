package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsAdGroupServingStatus the model 'SponsoredProductsAdGroupServingStatus'
type SponsoredProductsAdGroupServingStatus string

// List of SponsoredProductsAdGroupServingStatus
const (
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_AD_GROUP_STATUS_ENABLED             SponsoredProductsAdGroupServingStatus = "AD_GROUP_STATUS_ENABLED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_AD_GROUP_PAUSED                     SponsoredProductsAdGroupServingStatus = "AD_GROUP_PAUSED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_AD_GROUP_ARCHIVED                   SponsoredProductsAdGroupServingStatus = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_AD_GROUP_INCOMPLETE                 SponsoredProductsAdGroupServingStatus = "AD_GROUP_INCOMPLETE"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    SponsoredProductsAdGroupServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED SponsoredProductsAdGroupServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_AD_GROUP_LOW_BID                    SponsoredProductsAdGroupServingStatus = "AD_GROUP_LOW_BID"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             SponsoredProductsAdGroupServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_CAMPAIGN_PAUSED                     SponsoredProductsAdGroupServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_CAMPAIGN_ARCHIVED                   SponsoredProductsAdGroupServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PENDING_REVIEW                      SponsoredProductsAdGroupServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_REJECTED                            SponsoredProductsAdGroupServingStatus = "REJECTED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PENDING_START_DATE                  SponsoredProductsAdGroupServingStatus = "PENDING_START_DATE"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_ENDED                               SponsoredProductsAdGroupServingStatus = "ENDED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              SponsoredProductsAdGroupServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_CAMPAIGN_INCOMPLETE                 SponsoredProductsAdGroupServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED            SponsoredProductsAdGroupServingStatus = "PORTFOLIO_STATUS_ENABLED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PORTFOLIO_PAUSED                    SponsoredProductsAdGroupServingStatus = "PORTFOLIO_PAUSED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PORTFOLIO_ARCHIVED                  SponsoredProductsAdGroupServingStatus = "PORTFOLIO_ARCHIVED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET             SponsoredProductsAdGroupServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE        SponsoredProductsAdGroupServingStatus = "PORTFOLIO_PENDING_START_DATE"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_PORTFOLIO_ENDED                     SponsoredProductsAdGroupServingStatus = "PORTFOLIO_ENDED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       SponsoredProductsAdGroupServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  SponsoredProductsAdGroupServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_ADVERTISER_ARCHIVED                 SponsoredProductsAdGroupServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_ADVERTISER_PAUSED                   SponsoredProductsAdGroupServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_ADVERTISER_OUT_OF_BUDGET            SponsoredProductsAdGroupServingStatus = "ADVERTISER_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          SponsoredProductsAdGroupServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	SPONSOREDPRODUCTSADGROUPSERVINGSTATUS_OTHER                               SponsoredProductsAdGroupServingStatus = "OTHER"
)

// All allowed values of SponsoredProductsAdGroupServingStatus enum
var AllowedSponsoredProductsAdGroupServingStatusEnumValues = []SponsoredProductsAdGroupServingStatus{
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
	"OTHER",
}

func (v *SponsoredProductsAdGroupServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAdGroupServingStatus(value)
	for _, existing := range AllowedSponsoredProductsAdGroupServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAdGroupServingStatus", value)
}

// NewSponsoredProductsAdGroupServingStatusFromValue returns a pointer to a valid SponsoredProductsAdGroupServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAdGroupServingStatusFromValue(v string) (*SponsoredProductsAdGroupServingStatus, error) {
	ev := SponsoredProductsAdGroupServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAdGroupServingStatus: valid values are %v", v, AllowedSponsoredProductsAdGroupServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAdGroupServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAdGroupServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAdGroupServingStatus value
func (v SponsoredProductsAdGroupServingStatus) Ptr() *SponsoredProductsAdGroupServingStatus {
	return &v
}

type NullableSponsoredProductsAdGroupServingStatus struct {
	value *SponsoredProductsAdGroupServingStatus
	isSet bool
}

func (v NullableSponsoredProductsAdGroupServingStatus) Get() *SponsoredProductsAdGroupServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupServingStatus) Set(val *SponsoredProductsAdGroupServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupServingStatus(val *SponsoredProductsAdGroupServingStatus) *NullableSponsoredProductsAdGroupServingStatus {
	return &NullableSponsoredProductsAdGroupServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
