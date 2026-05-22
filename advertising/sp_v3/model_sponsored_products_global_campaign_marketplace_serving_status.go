package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalCampaignMarketplaceServingStatus the model 'SponsoredProductsGlobalCampaignMarketplaceServingStatus'
type SponsoredProductsGlobalCampaignMarketplaceServingStatus string

// List of SponsoredProductsGlobalCampaignMarketplaceServingStatus
const (
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_CAMPAIGN_STATUS_ENABLED            SponsoredProductsGlobalCampaignMarketplaceServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_CAMPAIGN_PAUSED                    SponsoredProductsGlobalCampaignMarketplaceServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_CAMPAIGN_ARCHIVED                  SponsoredProductsGlobalCampaignMarketplaceServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_PENDING_REVIEW                     SponsoredProductsGlobalCampaignMarketplaceServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_REJECTED                           SponsoredProductsGlobalCampaignMarketplaceServingStatus = "REJECTED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_PENDING_START_DATE                 SponsoredProductsGlobalCampaignMarketplaceServingStatus = "PENDING_START_DATE"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_ENDED                              SponsoredProductsGlobalCampaignMarketplaceServingStatus = "ENDED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET             SponsoredProductsGlobalCampaignMarketplaceServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_CAMPAIGN_INCOMPLETE                SponsoredProductsGlobalCampaignMarketplaceServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED      SponsoredProductsGlobalCampaignMarketplaceServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW SponsoredProductsGlobalCampaignMarketplaceServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_ADVERTISER_ARCHIVED                SponsoredProductsGlobalCampaignMarketplaceServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_ADVERTISER_PAUSED                  SponsoredProductsGlobalCampaignMarketplaceServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSGLOBALCAMPAIGNMARKETPLACESERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE         SponsoredProductsGlobalCampaignMarketplaceServingStatus = "ADVERTISER_PAYMENT_FAILURE"
)

// All allowed values of SponsoredProductsGlobalCampaignMarketplaceServingStatus enum
var AllowedSponsoredProductsGlobalCampaignMarketplaceServingStatusEnumValues = []SponsoredProductsGlobalCampaignMarketplaceServingStatus{
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
}

func (v *SponsoredProductsGlobalCampaignMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalCampaignMarketplaceServingStatus(value)
	for _, existing := range AllowedSponsoredProductsGlobalCampaignMarketplaceServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalCampaignMarketplaceServingStatus", value)
}

// NewSponsoredProductsGlobalCampaignMarketplaceServingStatusFromValue returns a pointer to a valid SponsoredProductsGlobalCampaignMarketplaceServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalCampaignMarketplaceServingStatusFromValue(v string) (*SponsoredProductsGlobalCampaignMarketplaceServingStatus, error) {
	ev := SponsoredProductsGlobalCampaignMarketplaceServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalCampaignMarketplaceServingStatus: valid values are %v", v, AllowedSponsoredProductsGlobalCampaignMarketplaceServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalCampaignMarketplaceServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalCampaignMarketplaceServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalCampaignMarketplaceServingStatus value
func (v SponsoredProductsGlobalCampaignMarketplaceServingStatus) Ptr() *SponsoredProductsGlobalCampaignMarketplaceServingStatus {
	return &v
}

type NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus struct {
	value *SponsoredProductsGlobalCampaignMarketplaceServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus) Get() *SponsoredProductsGlobalCampaignMarketplaceServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus) Set(val *SponsoredProductsGlobalCampaignMarketplaceServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignMarketplaceServingStatus(val *SponsoredProductsGlobalCampaignMarketplaceServingStatus) *NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus {
	return &NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
