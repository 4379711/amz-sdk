package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalAdGroupServingStatusReason the model 'SponsoredProductsGlobalAdGroupServingStatusReason'
type SponsoredProductsGlobalAdGroupServingStatusReason string

// List of SponsoredProductsGlobalAdGroupServingStatusReason
const (
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_AD_GROUP_DELIVERING_IN_ALL_MARKETPLACES                         SponsoredProductsGlobalAdGroupServingStatusReason = "AD_GROUP_DELIVERING_IN_ALL_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                     SponsoredProductsGlobalAdGroupServingStatusReason = "AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE                 SponsoredProductsGlobalAdGroupServingStatusReason = "AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_AD_GROUP_ARCHIVED                                               SponsoredProductsGlobalAdGroupServingStatusReason = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE    SponsoredProductsGlobalAdGroupServingStatusReason = "AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE SponsoredProductsGlobalAdGroupServingStatusReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE                    SponsoredProductsGlobalAdGroupServingStatusReason = "AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_CAMPAIGN_ARCHIVED                                               SponsoredProductsGlobalAdGroupServingStatusReason = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                     SponsoredProductsGlobalAdGroupServingStatusReason = "CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE                 SponsoredProductsGlobalAdGroupServingStatusReason = "CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_CAMPAIGN_PENDING_START                                          SponsoredProductsGlobalAdGroupServingStatusReason = "CAMPAIGN_PENDING_START"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_CAMPAIGN_ENDED                                                  SponsoredProductsGlobalAdGroupServingStatusReason = "CAMPAIGN_ENDED"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE              SponsoredProductsGlobalAdGroupServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_PAYMENT_FAILED_IN_ATLEAST_ONE_MARKETPLACE                       SponsoredProductsGlobalAdGroupServingStatusReason = "PAYMENT_FAILED_IN_ATLEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES  SponsoredProductsGlobalAdGroupServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES       SponsoredProductsGlobalAdGroupServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES                   SponsoredProductsGlobalAdGroupServingStatusReason = "ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALADGROUPSERVINGSTATUSREASON_ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES                 SponsoredProductsGlobalAdGroupServingStatusReason = "ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES"
)

// All allowed values of SponsoredProductsGlobalAdGroupServingStatusReason enum
var AllowedSponsoredProductsGlobalAdGroupServingStatusReasonEnumValues = []SponsoredProductsGlobalAdGroupServingStatusReason{
	"AD_GROUP_DELIVERING_IN_ALL_MARKETPLACES",
	"AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_ARCHIVED",
	"AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_ARCHIVED",
	"CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_PENDING_START",
	"CAMPAIGN_ENDED",
	"CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE",
	"PAYMENT_FAILED_IN_ATLEAST_ONE_MARKETPLACE",
	"ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES",
}

func (v *SponsoredProductsGlobalAdGroupServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalAdGroupServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalAdGroupServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalAdGroupServingStatusReason", value)
}

// NewSponsoredProductsGlobalAdGroupServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsGlobalAdGroupServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalAdGroupServingStatusReasonFromValue(v string) (*SponsoredProductsGlobalAdGroupServingStatusReason, error) {
	ev := SponsoredProductsGlobalAdGroupServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalAdGroupServingStatusReason: valid values are %v", v, AllowedSponsoredProductsGlobalAdGroupServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalAdGroupServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalAdGroupServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalAdGroupServingStatusReason value
func (v SponsoredProductsGlobalAdGroupServingStatusReason) Ptr() *SponsoredProductsGlobalAdGroupServingStatusReason {
	return &v
}

type NullableSponsoredProductsGlobalAdGroupServingStatusReason struct {
	value *SponsoredProductsGlobalAdGroupServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroupServingStatusReason) Get() *SponsoredProductsGlobalAdGroupServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroupServingStatusReason) Set(val *SponsoredProductsGlobalAdGroupServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroupServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroupServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroupServingStatusReason(val *SponsoredProductsGlobalAdGroupServingStatusReason) *NullableSponsoredProductsGlobalAdGroupServingStatusReason {
	return &NullableSponsoredProductsGlobalAdGroupServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroupServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroupServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
