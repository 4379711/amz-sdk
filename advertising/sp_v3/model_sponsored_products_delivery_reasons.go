package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsDeliveryReasons DeliveryReasons provides a description for the DeliveryStatus.
type SponsoredProductsDeliveryReasons string

// List of SponsoredProductsDeliveryReasons
const (
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_PAUSED                    SponsoredProductsDeliveryReasons = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_ARCHIVED                  SponsoredProductsDeliveryReasons = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_PENDING_START_DATE        SponsoredProductsDeliveryReasons = "CAMPAIGN_PENDING_START_DATE"
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_END_DATE_REACHED          SponsoredProductsDeliveryReasons = "CAMPAIGN_END_DATE_REACHED"
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_REJECTED                  SponsoredProductsDeliveryReasons = "CAMPAIGN_REJECTED"
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_PENDING_REVIEW            SponsoredProductsDeliveryReasons = "CAMPAIGN_PENDING_REVIEW"
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_INCOMPLETE                SponsoredProductsDeliveryReasons = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSDELIVERYREASONS_CAMPAIGN_OUT_OF_BUDGET             SponsoredProductsDeliveryReasons = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSDELIVERYREASONS_PORTFOLIO_OUT_OF_BUDGET            SponsoredProductsDeliveryReasons = "PORTFOLIO_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSDELIVERYREASONS_PORTFOLIO_END_DATE_REACHED         SponsoredProductsDeliveryReasons = "PORTFOLIO_END_DATE_REACHED"
	SPONSOREDPRODUCTSDELIVERYREASONS_PORTFOLIO_ARCHIVED                 SponsoredProductsDeliveryReasons = "PORTFOLIO_ARCHIVED"
	SPONSOREDPRODUCTSDELIVERYREASONS_PORTFOLIO_PAUSED                   SponsoredProductsDeliveryReasons = "PORTFOLIO_PAUSED"
	SPONSOREDPRODUCTSDELIVERYREASONS_PORTFOLIO_PENDING_START_DATE       SponsoredProductsDeliveryReasons = "PORTFOLIO_PENDING_START_DATE"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_POLICING_SUSPENDED      SponsoredProductsDeliveryReasons = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_POLICING_PENDING_REVIEW SponsoredProductsDeliveryReasons = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_ARCHIVED                SponsoredProductsDeliveryReasons = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_OUT_OF_BUDGET           SponsoredProductsDeliveryReasons = "ADVERTISER_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_PAUSED                  SponsoredProductsDeliveryReasons = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_PAYMENT_FAILURE         SponsoredProductsDeliveryReasons = "ADVERTISER_PAYMENT_FAILURE"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_OUT_OF_PREPAY_BALANCE   SponsoredProductsDeliveryReasons = "ADVERTISER_OUT_OF_PREPAY_BALANCE"
	SPONSOREDPRODUCTSDELIVERYREASONS_ADVERTISER_INELIGIBLE              SponsoredProductsDeliveryReasons = "ADVERTISER_INELIGIBLE"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_GROUP_PAUSED                    SponsoredProductsDeliveryReasons = "AD_GROUP_PAUSED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_GROUP_ARCHIVED                  SponsoredProductsDeliveryReasons = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_GROUP_INCOMPLETE                SponsoredProductsDeliveryReasons = "AD_GROUP_INCOMPLETE"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_GROUP_POLICING_PENDING_REVIEW   SponsoredProductsDeliveryReasons = "AD_GROUP_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_GROUP_LOW_BID                   SponsoredProductsDeliveryReasons = "AD_GROUP_LOW_BID"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_GROUP_PENDING_REVIEW            SponsoredProductsDeliveryReasons = "AD_GROUP_PENDING_REVIEW"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_GROUP_REJECTED                  SponsoredProductsDeliveryReasons = "AD_GROUP_REJECTED"
	SPONSOREDPRODUCTSDELIVERYREASONS_TARGET_POLICING_SUSPENDED          SponsoredProductsDeliveryReasons = "TARGET_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSDELIVERYREASONS_TARGET_PAUSED                      SponsoredProductsDeliveryReasons = "TARGET_PAUSED"
	SPONSOREDPRODUCTSDELIVERYREASONS_TARGET_ARCHIVED                    SponsoredProductsDeliveryReasons = "TARGET_ARCHIVED"
	SPONSOREDPRODUCTSDELIVERYREASONS_TARGET_BLOCKED                     SponsoredProductsDeliveryReasons = "TARGET_BLOCKED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_PAUSED                          SponsoredProductsDeliveryReasons = "AD_PAUSED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_ARCHIVED                        SponsoredProductsDeliveryReasons = "AD_ARCHIVED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_POLICING_PENDING_REVIEW         SponsoredProductsDeliveryReasons = "AD_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_POLICING_SUSPENDED              SponsoredProductsDeliveryReasons = "AD_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_CREATION_FAILED                 SponsoredProductsDeliveryReasons = "AD_CREATION_FAILED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_CREATION_IN_PROGRESS            SponsoredProductsDeliveryReasons = "AD_CREATION_IN_PROGRESS"
	SPONSOREDPRODUCTSDELIVERYREASONS_LANDING_PAGE_NOT_AVAILABLE         SponsoredProductsDeliveryReasons = "LANDING_PAGE_NOT_AVAILABLE"
	SPONSOREDPRODUCTSDELIVERYREASONS_NOT_BUYABLE                        SponsoredProductsDeliveryReasons = "NOT_BUYABLE"
	SPONSOREDPRODUCTSDELIVERYREASONS_NOT_IN_BUYBOX                      SponsoredProductsDeliveryReasons = "NOT_IN_BUYBOX"
	SPONSOREDPRODUCTSDELIVERYREASONS_NOT_IN_POLICY                      SponsoredProductsDeliveryReasons = "NOT_IN_POLICY"
	SPONSOREDPRODUCTSDELIVERYREASONS_OUT_OF_STOCK                       SponsoredProductsDeliveryReasons = "OUT_OF_STOCK"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_MISSING_IMAGE                   SponsoredProductsDeliveryReasons = "AD_MISSING_IMAGE"
	SPONSOREDPRODUCTSDELIVERYREASONS_SECURITY_SCAN_PENDING_REVIEW       SponsoredProductsDeliveryReasons = "SECURITY_SCAN_PENDING_REVIEW"
	SPONSOREDPRODUCTSDELIVERYREASONS_SECURITY_SCAN_REJECTED             SponsoredProductsDeliveryReasons = "SECURITY_SCAN_REJECTED"
	SPONSOREDPRODUCTSDELIVERYREASONS_STATUS_UNAVAILABLE                 SponsoredProductsDeliveryReasons = "STATUS_UNAVAILABLE"
	SPONSOREDPRODUCTSDELIVERYREASONS_NO_INVENTORY                       SponsoredProductsDeliveryReasons = "NO_INVENTORY"
	SPONSOREDPRODUCTSDELIVERYREASONS_NO_PURCHASABLE_OFFER               SponsoredProductsDeliveryReasons = "NO_PURCHASABLE_OFFER"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_INELIGIBLE                      SponsoredProductsDeliveryReasons = "AD_INELIGIBLE"
	SPONSOREDPRODUCTSDELIVERYREASONS_CREATIVE_MISSING_ASSET             SponsoredProductsDeliveryReasons = "CREATIVE_MISSING_ASSET"
	SPONSOREDPRODUCTSDELIVERYREASONS_CREATIVE_PENDING_REVIEW            SponsoredProductsDeliveryReasons = "CREATIVE_PENDING_REVIEW"
	SPONSOREDPRODUCTSDELIVERYREASONS_CREATIVE_REJECTED                  SponsoredProductsDeliveryReasons = "CREATIVE_REJECTED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_MISSING_DECORATION              SponsoredProductsDeliveryReasons = "AD_MISSING_DECORATION"
	SPONSOREDPRODUCTSDELIVERYREASONS_PIR_RULE_EXCLUDED                  SponsoredProductsDeliveryReasons = "PIR_RULE_EXCLUDED"
	SPONSOREDPRODUCTSDELIVERYREASONS_AD_NOT_DELIVERING                  SponsoredProductsDeliveryReasons = "AD_NOT_DELIVERING"
	SPONSOREDPRODUCTSDELIVERYREASONS_OTHER                              SponsoredProductsDeliveryReasons = "OTHER"
)

// All allowed values of SponsoredProductsDeliveryReasons enum
var AllowedSponsoredProductsDeliveryReasonsEnumValues = []SponsoredProductsDeliveryReasons{
	"CAMPAIGN_PAUSED",
	"CAMPAIGN_ARCHIVED",
	"CAMPAIGN_PENDING_START_DATE",
	"CAMPAIGN_END_DATE_REACHED",
	"CAMPAIGN_REJECTED",
	"CAMPAIGN_PENDING_REVIEW",
	"CAMPAIGN_INCOMPLETE",
	"CAMPAIGN_OUT_OF_BUDGET",
	"PORTFOLIO_OUT_OF_BUDGET",
	"PORTFOLIO_END_DATE_REACHED",
	"PORTFOLIO_ARCHIVED",
	"PORTFOLIO_PAUSED",
	"PORTFOLIO_PENDING_START_DATE",
	"ADVERTISER_POLICING_SUSPENDED",
	"ADVERTISER_POLICING_PENDING_REVIEW",
	"ADVERTISER_ARCHIVED",
	"ADVERTISER_OUT_OF_BUDGET",
	"ADVERTISER_PAUSED",
	"ADVERTISER_PAYMENT_FAILURE",
	"ADVERTISER_OUT_OF_PREPAY_BALANCE",
	"ADVERTISER_INELIGIBLE",
	"AD_GROUP_PAUSED",
	"AD_GROUP_ARCHIVED",
	"AD_GROUP_INCOMPLETE",
	"AD_GROUP_POLICING_PENDING_REVIEW",
	"AD_GROUP_LOW_BID",
	"AD_GROUP_PENDING_REVIEW",
	"AD_GROUP_REJECTED",
	"TARGET_POLICING_SUSPENDED",
	"TARGET_PAUSED",
	"TARGET_ARCHIVED",
	"TARGET_BLOCKED",
	"AD_PAUSED",
	"AD_ARCHIVED",
	"AD_POLICING_PENDING_REVIEW",
	"AD_POLICING_SUSPENDED",
	"AD_CREATION_FAILED",
	"AD_CREATION_IN_PROGRESS",
	"LANDING_PAGE_NOT_AVAILABLE",
	"NOT_BUYABLE",
	"NOT_IN_BUYBOX",
	"NOT_IN_POLICY",
	"OUT_OF_STOCK",
	"AD_MISSING_IMAGE",
	"SECURITY_SCAN_PENDING_REVIEW",
	"SECURITY_SCAN_REJECTED",
	"STATUS_UNAVAILABLE",
	"NO_INVENTORY",
	"NO_PURCHASABLE_OFFER",
	"AD_INELIGIBLE",
	"CREATIVE_MISSING_ASSET",
	"CREATIVE_PENDING_REVIEW",
	"CREATIVE_REJECTED",
	"AD_MISSING_DECORATION",
	"PIR_RULE_EXCLUDED",
	"AD_NOT_DELIVERING",
	"OTHER",
}

func (v *SponsoredProductsDeliveryReasons) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsDeliveryReasons(value)
	for _, existing := range AllowedSponsoredProductsDeliveryReasonsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsDeliveryReasons", value)
}

// NewSponsoredProductsDeliveryReasonsFromValue returns a pointer to a valid SponsoredProductsDeliveryReasons
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsDeliveryReasonsFromValue(v string) (*SponsoredProductsDeliveryReasons, error) {
	ev := SponsoredProductsDeliveryReasons(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsDeliveryReasons: valid values are %v", v, AllowedSponsoredProductsDeliveryReasonsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsDeliveryReasons) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsDeliveryReasonsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsDeliveryReasons value
func (v SponsoredProductsDeliveryReasons) Ptr() *SponsoredProductsDeliveryReasons {
	return &v
}

type NullableSponsoredProductsDeliveryReasons struct {
	value *SponsoredProductsDeliveryReasons
	isSet bool
}

func (v NullableSponsoredProductsDeliveryReasons) Get() *SponsoredProductsDeliveryReasons {
	return v.value
}

func (v *NullableSponsoredProductsDeliveryReasons) Set(val *SponsoredProductsDeliveryReasons) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeliveryReasons) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeliveryReasons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeliveryReasons(val *SponsoredProductsDeliveryReasons) *NullableSponsoredProductsDeliveryReasons {
	return &NullableSponsoredProductsDeliveryReasons{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeliveryReasons) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeliveryReasons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
