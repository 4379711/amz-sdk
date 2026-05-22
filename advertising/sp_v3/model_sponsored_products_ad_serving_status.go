package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsAdServingStatus the model 'SponsoredProductsAdServingStatus'
type SponsoredProductsAdServingStatus string

// List of SponsoredProductsAdServingStatus
const (
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_GROUP_STATUS_ENABLED             SponsoredProductsAdServingStatus = "AD_GROUP_STATUS_ENABLED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_GROUP_PAUSED                     SponsoredProductsAdServingStatus = "AD_GROUP_PAUSED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_GROUP_ARCHIVED                   SponsoredProductsAdServingStatus = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_GROUP_INCOMPLETE                 SponsoredProductsAdServingStatus = "AD_GROUP_INCOMPLETE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_GROUP_POLICING_PENDING_REVIEW    SponsoredProductsAdServingStatus = "AD_GROUP_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_GROUP_POLICING_CREATIVE_REJECTED SponsoredProductsAdServingStatus = "AD_GROUP_POLICING_CREATIVE_REJECTED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_GROUP_LOW_BID                    SponsoredProductsAdServingStatus = "AD_GROUP_LOW_BID"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_STATUS_ENABLED             SponsoredProductsAdServingStatus = "CAMPAIGN_STATUS_ENABLED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_PAUSED                     SponsoredProductsAdServingStatus = "CAMPAIGN_PAUSED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_ARCHIVED                   SponsoredProductsAdServingStatus = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PENDING_REVIEW                      SponsoredProductsAdServingStatus = "PENDING_REVIEW"
	SPONSOREDPRODUCTSADSERVINGSTATUS_REJECTED                            SponsoredProductsAdServingStatus = "REJECTED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PENDING_START_DATE                  SponsoredProductsAdServingStatus = "PENDING_START_DATE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ENDED                               SponsoredProductsAdServingStatus = "ENDED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_OUT_OF_BUDGET              SponsoredProductsAdServingStatus = "CAMPAIGN_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_INCOMPLETE                 SponsoredProductsAdServingStatus = "CAMPAIGN_INCOMPLETE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED            SponsoredProductsAdServingStatus = "PORTFOLIO_STATUS_ENABLED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PORTFOLIO_PAUSED                    SponsoredProductsAdServingStatus = "PORTFOLIO_PAUSED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PORTFOLIO_ARCHIVED                  SponsoredProductsAdServingStatus = "PORTFOLIO_ARCHIVED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET             SponsoredProductsAdServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE        SponsoredProductsAdServingStatus = "PORTFOLIO_PENDING_START_DATE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PORTFOLIO_ENDED                     SponsoredProductsAdServingStatus = "PORTFOLIO_ENDED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_POLICING_SUSPENDED       SponsoredProductsAdServingStatus = "ADVERTISER_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_POLICING_PENDING_REVIEW  SponsoredProductsAdServingStatus = "ADVERTISER_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_ARCHIVED                 SponsoredProductsAdServingStatus = "ADVERTISER_ARCHIVED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_PAUSED                   SponsoredProductsAdServingStatus = "ADVERTISER_PAUSED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_OUT_OF_BUDGET            SponsoredProductsAdServingStatus = "ADVERTISER_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE          SponsoredProductsAdServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_STATUS_LIVE                      SponsoredProductsAdServingStatus = "AD_STATUS_LIVE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_POLICING_PENDING_REVIEW          SponsoredProductsAdServingStatus = "AD_POLICING_PENDING_REVIEW"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_POLICING_SUSPENDED               SponsoredProductsAdServingStatus = "AD_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_PAUSED                           SponsoredProductsAdServingStatus = "AD_PAUSED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_ARCHIVED                         SponsoredProductsAdServingStatus = "AD_ARCHIVED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ACCOUNT_OUT_OF_BUDGET               SponsoredProductsAdServingStatus = "ACCOUNT_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_ACCOUNT_OUT_OF_BUDGET    SponsoredProductsAdServingStatus = "ADVERTISER_ACCOUNT_OUT_OF_BUDGET"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_EXCEED_SPENDS_LIMIT      SponsoredProductsAdServingStatus = "ADVERTISER_EXCEED_SPENDS_LIMIT"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ADVERTISER_STATUS_ENABLED           SponsoredProductsAdServingStatus = "ADVERTISER_STATUS_ENABLED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_CREATION_OFFLINE_FAILED          SponsoredProductsAdServingStatus = "AD_CREATION_OFFLINE_FAILED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_CREATION_OFFLINE_IN_PROGRESS     SponsoredProductsAdServingStatus = "AD_CREATION_OFFLINE_IN_PROGRESS"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_CREATION_OFFLINE_PENDING         SponsoredProductsAdServingStatus = "AD_CREATION_OFFLINE_PENDING"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_ADS_NOT_DELIVERING         SponsoredProductsAdServingStatus = "CAMPAIGN_ADS_NOT_DELIVERING"
	SPONSOREDPRODUCTSADSERVINGSTATUS_LANDING_PAGE_NOT_AVAILABLE          SponsoredProductsAdServingStatus = "LANDING_PAGE_NOT_AVAILABLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_MISSING_DECORATION                  SponsoredProductsAdServingStatus = "MISSING_DECORATION"
	SPONSOREDPRODUCTSADSERVINGSTATUS_MISSING_IMAGE                       SponsoredProductsAdServingStatus = "MISSING_IMAGE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_NOT_BUYABLE                         SponsoredProductsAdServingStatus = "NOT_BUYABLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_NOT_IN_BUYBOX                       SponsoredProductsAdServingStatus = "NOT_IN_BUYBOX"
	SPONSOREDPRODUCTSADSERVINGSTATUS_OUT_OF_STOCK                        SponsoredProductsAdServingStatus = "OUT_OF_STOCK"
	SPONSOREDPRODUCTSADSERVINGSTATUS_SECURITY_SCAN_PENDING_REVIEW        SponsoredProductsAdServingStatus = "SECURITY_SCAN_PENDING_REVIEW"
	SPONSOREDPRODUCTSADSERVINGSTATUS_SECURITY_SCAN_REJECTED              SponsoredProductsAdServingStatus = "SECURITY_SCAN_REJECTED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_STATUS_UNAVAILABLE                  SponsoredProductsAdServingStatus = "STATUS_UNAVAILABLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_TARGETING_CLAUSE_ARCHIVED           SponsoredProductsAdServingStatus = "TARGETING_CLAUSE_ARCHIVED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_TARGETING_CLAUSE_BLOCKED            SponsoredProductsAdServingStatus = "TARGETING_CLAUSE_BLOCKED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_TARGETING_CLAUSE_PAUSED             SponsoredProductsAdServingStatus = "TARGETING_CLAUSE_PAUSED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_TARGETING_CLAUSE_POLICING_SUSPENDED SponsoredProductsAdServingStatus = "TARGETING_CLAUSE_POLICING_SUSPENDED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_TARGETING_CLAUSE_STATUS_LIVE        SponsoredProductsAdServingStatus = "TARGETING_CLAUSE_STATUS_LIVE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_ELIGIBLE                            SponsoredProductsAdServingStatus = "ELIGIBLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_INELIGIBLE                          SponsoredProductsAdServingStatus = "INELIGIBLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_NO_INVENTORY                        SponsoredProductsAdServingStatus = "NO_INVENTORY"
	SPONSOREDPRODUCTSADSERVINGSTATUS_NO_PURCHASABLE_OFFER                SponsoredProductsAdServingStatus = "NO_PURCHASABLE_OFFER"
	SPONSOREDPRODUCTSADSERVINGSTATUS_PIR_RULE_EXCLUDED                   SponsoredProductsAdServingStatus = "PIR_RULE_EXCLUDED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_PENDING_START_DATE         SponsoredProductsAdServingStatus = "CAMPAIGN_PENDING_START_DATE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_CAMPAIGN_ENDED                      SponsoredProductsAdServingStatus = "CAMPAIGN_ENDED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_MISSING_IMAGE                    SponsoredProductsAdServingStatus = "AD_MISSING_IMAGE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_MISSING_DECORATION               SponsoredProductsAdServingStatus = "AD_MISSING_DECORATION"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_NOT_BUYABLE                      SponsoredProductsAdServingStatus = "AD_NOT_BUYABLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_NOT_IN_BUYBOX                    SponsoredProductsAdServingStatus = "AD_NOT_IN_BUYBOX"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_OUT_OF_STOCK                     SponsoredProductsAdServingStatus = "AD_OUT_OF_STOCK"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_LANDING_PAGE_NOT_AVAILABLE       SponsoredProductsAdServingStatus = "AD_LANDING_PAGE_NOT_AVAILABLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_NO_PURCHASABLE_OFFER             SponsoredProductsAdServingStatus = "AD_NO_PURCHASABLE_OFFER"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_INELIGIBLE                       SponsoredProductsAdServingStatus = "AD_INELIGIBLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_ELIGIBLE                         SponsoredProductsAdServingStatus = "AD_ELIGIBLE"
	SPONSOREDPRODUCTSADSERVINGSTATUS_AD_CREATION_FAILED                  SponsoredProductsAdServingStatus = "AD_CREATION_FAILED"
	SPONSOREDPRODUCTSADSERVINGSTATUS_OTHER                               SponsoredProductsAdServingStatus = "OTHER"
)

// All allowed values of SponsoredProductsAdServingStatus enum
var AllowedSponsoredProductsAdServingStatusEnumValues = []SponsoredProductsAdServingStatus{
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
	"AD_STATUS_LIVE",
	"AD_POLICING_PENDING_REVIEW",
	"AD_POLICING_SUSPENDED",
	"AD_PAUSED",
	"AD_ARCHIVED",
	"ACCOUNT_OUT_OF_BUDGET",
	"ADVERTISER_ACCOUNT_OUT_OF_BUDGET",
	"ADVERTISER_EXCEED_SPENDS_LIMIT",
	"ADVERTISER_STATUS_ENABLED",
	"AD_CREATION_OFFLINE_FAILED",
	"AD_CREATION_OFFLINE_IN_PROGRESS",
	"AD_CREATION_OFFLINE_PENDING",
	"CAMPAIGN_ADS_NOT_DELIVERING",
	"LANDING_PAGE_NOT_AVAILABLE",
	"MISSING_DECORATION",
	"MISSING_IMAGE",
	"NOT_BUYABLE",
	"NOT_IN_BUYBOX",
	"OUT_OF_STOCK",
	"SECURITY_SCAN_PENDING_REVIEW",
	"SECURITY_SCAN_REJECTED",
	"STATUS_UNAVAILABLE",
	"TARGETING_CLAUSE_ARCHIVED",
	"TARGETING_CLAUSE_BLOCKED",
	"TARGETING_CLAUSE_PAUSED",
	"TARGETING_CLAUSE_POLICING_SUSPENDED",
	"TARGETING_CLAUSE_STATUS_LIVE",
	"ELIGIBLE",
	"INELIGIBLE",
	"NO_INVENTORY",
	"NO_PURCHASABLE_OFFER",
	"PIR_RULE_EXCLUDED",
	"CAMPAIGN_PENDING_START_DATE",
	"CAMPAIGN_ENDED",
	"AD_MISSING_IMAGE",
	"AD_MISSING_DECORATION",
	"AD_NOT_BUYABLE",
	"AD_NOT_IN_BUYBOX",
	"AD_OUT_OF_STOCK",
	"AD_LANDING_PAGE_NOT_AVAILABLE",
	"AD_NO_PURCHASABLE_OFFER",
	"AD_INELIGIBLE",
	"AD_ELIGIBLE",
	"AD_CREATION_FAILED",
	"OTHER",
}

func (v *SponsoredProductsAdServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsAdServingStatus(value)
	for _, existing := range AllowedSponsoredProductsAdServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsAdServingStatus", value)
}

// NewSponsoredProductsAdServingStatusFromValue returns a pointer to a valid SponsoredProductsAdServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsAdServingStatusFromValue(v string) (*SponsoredProductsAdServingStatus, error) {
	ev := SponsoredProductsAdServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsAdServingStatus: valid values are %v", v, AllowedSponsoredProductsAdServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsAdServingStatus) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsAdServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsAdServingStatus value
func (v SponsoredProductsAdServingStatus) Ptr() *SponsoredProductsAdServingStatus {
	return &v
}

type NullableSponsoredProductsAdServingStatus struct {
	value *SponsoredProductsAdServingStatus
	isSet bool
}

func (v NullableSponsoredProductsAdServingStatus) Get() *SponsoredProductsAdServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsAdServingStatus) Set(val *SponsoredProductsAdServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdServingStatus(val *SponsoredProductsAdServingStatus) *NullableSponsoredProductsAdServingStatus {
	return &NullableSponsoredProductsAdServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
