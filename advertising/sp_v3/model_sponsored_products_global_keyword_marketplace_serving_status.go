package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalKeywordMarketplaceServingStatus the model 'SponsoredProductsGlobalKeywordMarketplaceServingStatus'
type SponsoredProductsGlobalKeywordMarketplaceServingStatus string

// List of SponsoredProductsGlobalKeywordMarketplaceServingStatus
const (
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_TARGETING_CLAUSE_STATUS_LIVE        SponsoredProductsGlobalKeywordMarketplaceServingStatus = "TARGETING_CLAUSE_STATUS_LIVE"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_TARGETING_CLAUSE_POLICING_SUSPENDED SponsoredProductsGlobalKeywordMarketplaceServingStatus = "TARGETING_CLAUSE_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_TARGETING_CLAUSE_PAUSED             SponsoredProductsGlobalKeywordMarketplaceServingStatus = "TARGETING_CLAUSE_PAUSED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_TARGETING_CLAUSE_ARCHIVED           SponsoredProductsGlobalKeywordMarketplaceServingStatus = "TARGETING_CLAUSE_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_TARGETING_CLAUSE_BLOCKED            SponsoredProductsGlobalKeywordMarketplaceServingStatus = "TARGETING_CLAUSE_BLOCKED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_AD_GROUP_STATUS_ENABLED             SponsoredProductsGlobalKeywordMarketplaceServingStatus = "AD_GROUP_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_AD_GROUP_PAUSED                     SponsoredProductsGlobalKeywordMarketplaceServingStatus = "AD_GROUP_PAUSED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_AD_GROUP_ARCHIVED                   SponsoredProductsGlobalKeywordMarketplaceServingStatus = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_AD_GROUP_INCOMPLETE                 SponsoredProductsGlobalKeywordMarketplaceServingStatus = "AD_GROUP_INCOMPLETE"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    SponsoredProductsGlobalKeywordMarketplaceServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED SponsoredProductsGlobalKeywordMarketplaceServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_AD_GROUP_LOW_BID                    SponsoredProductsGlobalKeywordMarketplaceServingStatus = "AD_GROUP_LOW_BID"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             SponsoredProductsGlobalKeywordMarketplaceServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_CAMPAIGN_PAUSED                     SponsoredProductsGlobalKeywordMarketplaceServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_CAMPAIGN_ARCHIVED                   SponsoredProductsGlobalKeywordMarketplaceServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_PENDING_REVIEW                      SponsoredProductsGlobalKeywordMarketplaceServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_REJECTED                            SponsoredProductsGlobalKeywordMarketplaceServingStatus = "REJECTED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_PENDING_START_DATE                  SponsoredProductsGlobalKeywordMarketplaceServingStatus = "PENDING_START_DATE"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_ENDED                               SponsoredProductsGlobalKeywordMarketplaceServingStatus = "ENDED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              SponsoredProductsGlobalKeywordMarketplaceServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_CAMPAIGN_INCOMPLETE                 SponsoredProductsGlobalKeywordMarketplaceServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_PORTFOLIO_STATUS_ENABLED            SponsoredProductsGlobalKeywordMarketplaceServingStatus = "PORTFOLIO_STATUS_ENABLED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       SponsoredProductsGlobalKeywordMarketplaceServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  SponsoredProductsGlobalKeywordMarketplaceServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_ADVERTISER_ARCHIVED                 SponsoredProductsGlobalKeywordMarketplaceServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_ADVERTISER_PAUSED                   SponsoredProductsGlobalKeywordMarketplaceServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_ADVERTISER_OUT_OF_BUDGET            SponsoredProductsGlobalKeywordMarketplaceServingStatus = "ADVERTISER_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSGLOBALKEYWORDMARKETPLACESERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          SponsoredProductsGlobalKeywordMarketplaceServingStatus = "ADVERTISER_PAYMENT_FAILURE"
)

// All allowed values of SponsoredProductsGlobalKeywordMarketplaceServingStatus enum
var AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusEnumValues = []SponsoredProductsGlobalKeywordMarketplaceServingStatus{
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
	"ADVERTISER_POLICING_SUSPENDED",
	"ADVERTISER_POLICING_PENDING_REVIEW",
	"ADVERTISER_ARCHIVED",
	"ADVERTISER_PAUSED",
	"ADVERTISER_OUT_OF_BUDGET",
	"ADVERTISER_PAYMENT_FAILURE",
}

func (v *SponsoredProductsGlobalKeywordMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalKeywordMarketplaceServingStatus(value)
	for _, existing := range AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalKeywordMarketplaceServingStatus", value)
}

// NewSponsoredProductsGlobalKeywordMarketplaceServingStatusFromValue returns a pointer to a valid SponsoredProductsGlobalKeywordMarketplaceServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalKeywordMarketplaceServingStatusFromValue(v string) (*SponsoredProductsGlobalKeywordMarketplaceServingStatus, error) {
	ev := SponsoredProductsGlobalKeywordMarketplaceServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalKeywordMarketplaceServingStatus: valid values are %v", v, AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalKeywordMarketplaceServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalKeywordMarketplaceServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalKeywordMarketplaceServingStatus value
func (v SponsoredProductsGlobalKeywordMarketplaceServingStatus) Ptr() *SponsoredProductsGlobalKeywordMarketplaceServingStatus {
	return &v
}

type NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus struct {
	value *SponsoredProductsGlobalKeywordMarketplaceServingStatus
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus) Get() *SponsoredProductsGlobalKeywordMarketplaceServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus) Set(val *SponsoredProductsGlobalKeywordMarketplaceServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordMarketplaceServingStatus(val *SponsoredProductsGlobalKeywordMarketplaceServingStatus) *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus {
	return &NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordMarketplaceServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
