package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalProductAdMarketplaceServingStatus the model 'SponsoredProductsGlobalProductAdMarketplaceServingStatus'
type SponsoredProductsGlobalProductAdMarketplaceServingStatus string

// List of SponsoredProductsGlobalProductAdMarketplaceServingStatus
const (
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_GROUP_STATUS_ENABLED             SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_GROUP_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_GROUP_PAUSED                     SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_GROUP_PAUSED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_GROUP_ARCHIVED                   SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_GROUP_INCOMPLETE                 SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_GROUP_INCOMPLETE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_GROUP_LOW_BID                    SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_GROUP_LOW_BID"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             SponsoredProductsGlobalProductAdMarketplaceServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_CAMPAIGN_PAUSED                     SponsoredProductsGlobalProductAdMarketplaceServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_CAMPAIGN_ARCHIVED                   SponsoredProductsGlobalProductAdMarketplaceServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_PENDING_REVIEW                      SponsoredProductsGlobalProductAdMarketplaceServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_REJECTED                            SponsoredProductsGlobalProductAdMarketplaceServingStatus = "REJECTED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_CAMPAIGN_PENDING_START_DATE         SponsoredProductsGlobalProductAdMarketplaceServingStatus = "CAMPAIGN_PENDING_START_DATE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_CAMPAIGN_ENDED                      SponsoredProductsGlobalProductAdMarketplaceServingStatus = "CAMPAIGN_ENDED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              SponsoredProductsGlobalProductAdMarketplaceServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_CAMPAIGN_INCOMPLETE                 SponsoredProductsGlobalProductAdMarketplaceServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_ADVERTISER_STATUS_ENABLED           SponsoredProductsGlobalProductAdMarketplaceServingStatus = "ADVERTISER_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       SponsoredProductsGlobalProductAdMarketplaceServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  SponsoredProductsGlobalProductAdMarketplaceServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_ADVERTISER_ARCHIVED                 SponsoredProductsGlobalProductAdMarketplaceServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_ADVERTISER_PAUSED                   SponsoredProductsGlobalProductAdMarketplaceServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          SponsoredProductsGlobalProductAdMarketplaceServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_STATUS_LIVE                      SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_STATUS_LIVE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_POLICING_PENDING_REVIEW          SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_POLICING_SUSPENDED               SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_PAUSED                           SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_PAUSED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_ARCHIVED                         SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_CREATION_FAILED                  SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_CREATION_FAILED"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_STATUS_UNAVAILABLE                  SponsoredProductsGlobalProductAdMarketplaceServingStatus = "STATUS_UNAVAILABLE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_MISSING_IMAGE                    SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_MISSING_IMAGE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_MISSING_DECORATION               SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_MISSING_DECORATION"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_NOT_BUYABLE                      SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_NOT_BUYABLE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_NOT_IN_BUYBOX                    SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_NOT_IN_BUYBOX"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_OUT_OF_STOCK                     SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_OUT_OF_STOCK"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_LANDING_PAGE_NOT_AVAILABLE       SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_LANDING_PAGE_NOT_AVAILABLE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_NO_PURCHASABLE_OFFER             SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_NO_PURCHASABLE_OFFER"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_INELIGIBLE                       SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_INELIGIBLE"
	SPONSOREDPRODUCTSGLOBALPRODUCTADMARKETPLACESERVINGSTATUS_AD_ELIGIBLE                         SponsoredProductsGlobalProductAdMarketplaceServingStatus = "AD_ELIGIBLE"
)

// All allowed values of SponsoredProductsGlobalProductAdMarketplaceServingStatus enum
var AllowedSponsoredProductsGlobalProductAdMarketplaceServingStatusEnumValues = []SponsoredProductsGlobalProductAdMarketplaceServingStatus{
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
	"CAMPAIGN_PENDING_START_DATE",
	"CAMPAIGN_ENDED",
	"CAMPAIGN_OUT_OF_BUDGET",
	"CAMPAIGN_INCOMPLETE",
	"ADVERTISER_STATUS_ENABLED",
	"ADVERTISER_POLICING_SUSPENDED",
	"ADVERTISER_POLICING_PENDING_REVIEW",
	"ADVERTISER_ARCHIVED",
	"ADVERTISER_PAUSED",
	"ADVERTISER_PAYMENT_FAILURE",
	"AD_STATUS_LIVE",
	"AD_POLICING_PENDING_REVIEW",
	"AD_POLICING_SUSPENDED",
	"AD_PAUSED",
	"AD_ARCHIVED",
	"AD_CREATION_FAILED",
	"STATUS_UNAVAILABLE",
	"AD_MISSING_IMAGE",
	"AD_MISSING_DECORATION",
	"AD_NOT_BUYABLE",
	"AD_NOT_IN_BUYBOX",
	"AD_OUT_OF_STOCK",
	"AD_LANDING_PAGE_NOT_AVAILABLE",
	"AD_NO_PURCHASABLE_OFFER",
	"AD_INELIGIBLE",
	"AD_ELIGIBLE",
}

func (v *SponsoredProductsGlobalProductAdMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalProductAdMarketplaceServingStatus(value)
	for _, existing := range AllowedSponsoredProductsGlobalProductAdMarketplaceServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalProductAdMarketplaceServingStatus", value)
}

// NewSponsoredProductsGlobalProductAdMarketplaceServingStatusFromValue returns a pointer to a valid SponsoredProductsGlobalProductAdMarketplaceServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalProductAdMarketplaceServingStatusFromValue(v string) (*SponsoredProductsGlobalProductAdMarketplaceServingStatus, error) {
	ev := SponsoredProductsGlobalProductAdMarketplaceServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalProductAdMarketplaceServingStatus: valid values are %v", v, AllowedSponsoredProductsGlobalProductAdMarketplaceServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalProductAdMarketplaceServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalProductAdMarketplaceServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalProductAdMarketplaceServingStatus value
func (v SponsoredProductsGlobalProductAdMarketplaceServingStatus) Ptr() *SponsoredProductsGlobalProductAdMarketplaceServingStatus {
	return &v
}

type NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus struct {
	value *SponsoredProductsGlobalProductAdMarketplaceServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus) Get() *SponsoredProductsGlobalProductAdMarketplaceServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus) Set(val *SponsoredProductsGlobalProductAdMarketplaceServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAdMarketplaceServingStatus(val *SponsoredProductsGlobalProductAdMarketplaceServingStatus) *NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus {
	return &NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAdMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
