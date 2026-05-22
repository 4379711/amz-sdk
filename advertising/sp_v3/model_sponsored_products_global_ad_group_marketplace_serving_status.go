package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalAdGroupMarketplaceServingStatus the model 'SponsoredProductsGlobalAdGroupMarketplaceServingStatus'
type SponsoredProductsGlobalAdGroupMarketplaceServingStatus string

// List of SponsoredProductsGlobalAdGroupMarketplaceServingStatus
const (
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_AD_GROUP_STATUS_ENABLED             SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "AD_GROUP_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_AD_GROUP_PAUSED                     SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "AD_GROUP_PAUSED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_AD_GROUP_ARCHIVED                   SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_AD_GROUP_INCOMPLETE                 SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "AD_GROUP_INCOMPLETE"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_AD_GROUP_LOW_BID                    SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "AD_GROUP_LOW_BID"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_CAMPAIGN_PAUSED                     SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_CAMPAIGN_ARCHIVED                   SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_PENDING_REVIEW                      SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_REJECTED                            SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "REJECTED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_PENDING_START_DATE                  SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "PENDING_START_DATE"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_ENDED                               SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "ENDED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_CAMPAIGN_INCOMPLETE                 SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_ADVERTISER_ARCHIVED                 SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_ADVERTISER_PAUSED                   SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	SPONSOREDPRODUCTSGLOBALADGROUPMARKETPLACESERVINGSTATUS_ADVERTISER_STATUS_ENABLED           SponsoredProductsGlobalAdGroupMarketplaceServingStatus = "ADVERTISER_STATUS_ENABLED"
)

// All allowed values of SponsoredProductsGlobalAdGroupMarketplaceServingStatus enum
var AllowedSponsoredProductsGlobalAdGroupMarketplaceServingStatusEnumValues = []SponsoredProductsGlobalAdGroupMarketplaceServingStatus{
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
	"ADVERTISER_POLICING_SUSPENDED",
	"ADVERTISER_POLICING_PENDING_REVIEW",
	"ADVERTISER_ARCHIVED",
	"ADVERTISER_PAUSED",
	"ADVERTISER_PAYMENT_FAILURE",
	"ADVERTISER_STATUS_ENABLED",
}

func (v *SponsoredProductsGlobalAdGroupMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalAdGroupMarketplaceServingStatus(value)
	for _, existing := range AllowedSponsoredProductsGlobalAdGroupMarketplaceServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalAdGroupMarketplaceServingStatus", value)
}

// NewSponsoredProductsGlobalAdGroupMarketplaceServingStatusFromValue returns a pointer to a valid SponsoredProductsGlobalAdGroupMarketplaceServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalAdGroupMarketplaceServingStatusFromValue(v string) (*SponsoredProductsGlobalAdGroupMarketplaceServingStatus, error) {
	ev := SponsoredProductsGlobalAdGroupMarketplaceServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalAdGroupMarketplaceServingStatus: valid values are %v", v, AllowedSponsoredProductsGlobalAdGroupMarketplaceServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalAdGroupMarketplaceServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalAdGroupMarketplaceServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalAdGroupMarketplaceServingStatus value
func (v SponsoredProductsGlobalAdGroupMarketplaceServingStatus) Ptr() *SponsoredProductsGlobalAdGroupMarketplaceServingStatus {
	return &v
}

type NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus struct {
	value *SponsoredProductsGlobalAdGroupMarketplaceServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus) Get() *SponsoredProductsGlobalAdGroupMarketplaceServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus) Set(val *SponsoredProductsGlobalAdGroupMarketplaceServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus(val *SponsoredProductsGlobalAdGroupMarketplaceServingStatus) *NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus {
	return &NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroupMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
