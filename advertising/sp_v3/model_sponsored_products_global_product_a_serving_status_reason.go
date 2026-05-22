package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalProductAServingStatusReason the model 'SponsoredProductsGlobalProductAServingStatusReason'
type SponsoredProductsGlobalProductAServingStatusReason string

// List of SponsoredProductsGlobalProductAServingStatusReason
const (
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_PRODUCT_AD_DELIVERING_IN_ALL_MARKETPLACES                       SponsoredProductsGlobalProductAServingStatusReason = "PRODUCT_AD_DELIVERING_IN_ALL_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE          SponsoredProductsGlobalProductAServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES  SponsoredProductsGlobalProductAServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES       SponsoredProductsGlobalProductAServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES                   SponsoredProductsGlobalProductAServingStatusReason = "ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES                 SponsoredProductsGlobalProductAServingStatusReason = "ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE              SponsoredProductsGlobalProductAServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                     SponsoredProductsGlobalProductAServingStatusReason = "CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_CAMPAIGN_ARCHIVED                                               SponsoredProductsGlobalProductAServingStatusReason = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE                 SponsoredProductsGlobalProductAServingStatusReason = "CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_CAMPAIGN_PENDING_START_DATE                                     SponsoredProductsGlobalProductAServingStatusReason = "CAMPAIGN_PENDING_START_DATE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_CAMPAIGN_ENDED                                                  SponsoredProductsGlobalProductAServingStatusReason = "CAMPAIGN_ENDED"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                     SponsoredProductsGlobalProductAServingStatusReason = "AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_GROUP_ARCHIVED                                               SponsoredProductsGlobalProductAServingStatusReason = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE                 SponsoredProductsGlobalProductAServingStatusReason = "AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_GROUP_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE            SponsoredProductsGlobalProductAServingStatusReason = "AD_GROUP_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE    SponsoredProductsGlobalProductAServingStatusReason = "AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE SponsoredProductsGlobalProductAServingStatusReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE                    SponsoredProductsGlobalProductAServingStatusReason = "AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_ARCHIVED                                                     SponsoredProductsGlobalProductAServingStatusReason = "AD_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                           SponsoredProductsGlobalProductAServingStatusReason = "AD_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_POLICING_SUSPENDED_IN_AT_LEAST_ONE_MARKETPLACE               SponsoredProductsGlobalProductAServingStatusReason = "AD_POLICING_SUSPENDED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_MISSING_DECORATION_IN_AT_LEAST_ONE_MARKETPLACE               SponsoredProductsGlobalProductAServingStatusReason = "AD_MISSING_DECORATION_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_MISSING_IMAGE_IN_AT_LEAST_ONE_MARKETPLACE                    SponsoredProductsGlobalProductAServingStatusReason = "AD_MISSING_IMAGE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE          SponsoredProductsGlobalProductAServingStatusReason = "AD_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_NOT_BUYABLE_IN_AT_LEAST_ONE_MARKETPLACE                      SponsoredProductsGlobalProductAServingStatusReason = "AD_NOT_BUYABLE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_NOT_IN_BUYBOX_IN_AT_LEAST_ONE_MARKETPLACE                    SponsoredProductsGlobalProductAServingStatusReason = "AD_NOT_IN_BUYBOX_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_OUT_OF_STOCK_IN_AT_LEAST_ONE_MARKETPLACE                     SponsoredProductsGlobalProductAServingStatusReason = "AD_OUT_OF_STOCK_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_LANDING_PAGE_NOT_AVAILABLE_IN_AT_LEAST_ONE_MARKETPLACE       SponsoredProductsGlobalProductAServingStatusReason = "AD_LANDING_PAGE_NOT_AVAILABLE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_STATUS_UNAVAILABLE_IN_AT_LEAST_ONE_MARKETPLACE               SponsoredProductsGlobalProductAServingStatusReason = "AD_STATUS_UNAVAILABLE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_NO_PURCHASABLE_OFFER_IN_AT_LEAST_ONE_MARKETPLACE             SponsoredProductsGlobalProductAServingStatusReason = "AD_NO_PURCHASABLE_OFFER_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_INELIGIBLE_IN_AT_LEAST_ONE_MARKETPLACE                       SponsoredProductsGlobalProductAServingStatusReason = "AD_INELIGIBLE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALPRODUCTASERVINGSTATUSREASON_AD_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE                  SponsoredProductsGlobalProductAServingStatusReason = "AD_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE"
)

// All allowed values of SponsoredProductsGlobalProductAServingStatusReason enum
var AllowedSponsoredProductsGlobalProductAServingStatusReasonEnumValues = []SponsoredProductsGlobalProductAServingStatusReason{
	"PRODUCT_AD_DELIVERING_IN_ALL_MARKETPLACES",
	"ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE",
	"ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES",
	"CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_ARCHIVED",
	"CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_PENDING_START_DATE",
	"CAMPAIGN_ENDED",
	"AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_ARCHIVED",
	"AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_ARCHIVED",
	"AD_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_POLICING_SUSPENDED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_MISSING_DECORATION_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_MISSING_IMAGE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_NOT_BUYABLE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_NOT_IN_BUYBOX_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_OUT_OF_STOCK_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_LANDING_PAGE_NOT_AVAILABLE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_STATUS_UNAVAILABLE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_NO_PURCHASABLE_OFFER_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_INELIGIBLE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE",
}

func (v *SponsoredProductsGlobalProductAServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalProductAServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalProductAServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalProductAServingStatusReason", value)
}

// NewSponsoredProductsGlobalProductAServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsGlobalProductAServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalProductAServingStatusReasonFromValue(v string) (*SponsoredProductsGlobalProductAServingStatusReason, error) {
	ev := SponsoredProductsGlobalProductAServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalProductAServingStatusReason: valid values are %v", v, AllowedSponsoredProductsGlobalProductAServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalProductAServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalProductAServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalProductAServingStatusReason value
func (v SponsoredProductsGlobalProductAServingStatusReason) Ptr() *SponsoredProductsGlobalProductAServingStatusReason {
	return &v
}

type NullableSponsoredProductsGlobalProductAServingStatusReason struct {
	value *SponsoredProductsGlobalProductAServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAServingStatusReason) Get() *SponsoredProductsGlobalProductAServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAServingStatusReason) Set(val *SponsoredProductsGlobalProductAServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAServingStatusReason(val *SponsoredProductsGlobalProductAServingStatusReason) *NullableSponsoredProductsGlobalProductAServingStatusReason {
	return &NullableSponsoredProductsGlobalProductAServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
