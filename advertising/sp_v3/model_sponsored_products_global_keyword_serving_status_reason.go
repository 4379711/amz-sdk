package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsGlobalKeywordServingStatusReason the model 'SponsoredProductsGlobalKeywordServingStatusReason'
type SponsoredProductsGlobalKeywordServingStatusReason string

// List of SponsoredProductsGlobalKeywordServingStatusReason
const (
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_DELIVERING_IN_ALL_MARKETPLACES                 SponsoredProductsGlobalKeywordServingStatusReason = "TARGETING_CLAUSE_DELIVERING_IN_ALL_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE          SponsoredProductsGlobalKeywordServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_ACCOUNT_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE               SponsoredProductsGlobalKeywordServingStatusReason = "ACCOUNT_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES  SponsoredProductsGlobalKeywordServingStatusReason = "ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES       SponsoredProductsGlobalKeywordServingStatusReason = "ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES                   SponsoredProductsGlobalKeywordServingStatusReason = "ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES                 SponsoredProductsGlobalKeywordServingStatusReason = "ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE              SponsoredProductsGlobalKeywordServingStatusReason = "CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                     SponsoredProductsGlobalKeywordServingStatusReason = "CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_CAMPAIGN_ARCHIVED                                               SponsoredProductsGlobalKeywordServingStatusReason = "CAMPAIGN_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE                 SponsoredProductsGlobalKeywordServingStatusReason = "CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_CAMPAIGN_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE            SponsoredProductsGlobalKeywordServingStatusReason = "CAMPAIGN_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_CAMPAIGN_PENDING_START_DATE                                     SponsoredProductsGlobalKeywordServingStatusReason = "CAMPAIGN_PENDING_START_DATE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_CAMPAIGN_ENDED                                                  SponsoredProductsGlobalKeywordServingStatusReason = "CAMPAIGN_ENDED"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE                     SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_ARCHIVED                                               SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE                 SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_CREATION_IN_PROGRESS_IN_AT_LEAST_ONE_MARKETPLACE       SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_CREATION_IN_PROGRESS_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE            SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE    SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE                    SponsoredProductsGlobalKeywordServingStatusReason = "AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE             SponsoredProductsGlobalKeywordServingStatusReason = "TARGETING_CLAUSE_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_ARCHIVED                                       SponsoredProductsGlobalKeywordServingStatusReason = "TARGETING_CLAUSE_ARCHIVED"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_POLICING_SUSPENDED_IN_AT_LEAST_ONE_MARKETPLACE SponsoredProductsGlobalKeywordServingStatusReason = "TARGETING_CLAUSE_POLICING_SUSPENDED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE    SponsoredProductsGlobalKeywordServingStatusReason = "TARGETING_CLAUSE_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE"
	SPONSOREDPRODUCTSGLOBALKEYWORDSERVINGSTATUSREASON_TARGETING_CLAUSE_BLOCKED_IN_AT_LEAST_ONE_MARKETPLACE            SponsoredProductsGlobalKeywordServingStatusReason = "TARGETING_CLAUSE_BLOCKED_IN_AT_LEAST_ONE_MARKETPLACE"
)

// All allowed values of SponsoredProductsGlobalKeywordServingStatusReason enum
var AllowedSponsoredProductsGlobalKeywordServingStatusReasonEnumValues = []SponsoredProductsGlobalKeywordServingStatusReason{
	"TARGETING_CLAUSE_DELIVERING_IN_ALL_MARKETPLACES",
	"ADVERTISER_PAYMENT_FAILURE_IN_AT_LEAST_ONE_MARKETPLACE",
	"ACCOUNT_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE",
	"ADVERTISER_POLICING_PENDING_REVIEW_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_POLICING_SUSPENDED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_PAUSED_IN_ATLEAST_ONE_MARKETPLACES",
	"ADVERTISER_ARCHIVED_IN_ATLEAST_ONE_MARKETPLACES",
	"CAMPAIGN_OUT_OF_BUDGET_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_ARCHIVED",
	"CAMPAIGN_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE",
	"CAMPAIGN_PENDING_START_DATE",
	"CAMPAIGN_ENDED",
	"AD_GROUP_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_ARCHIVED",
	"AD_GROUP_INCOMPLETE_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_CREATION_IN_PROGRESS_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_POLICING_PENDING_REVIEW_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_POLICING_CREATIVE_REJECTED_IN_AT_LEAST_ONE_MARKETPLACE",
	"AD_GROUP_LOW_BID_IN_AT_LEAST_ONE_MARKETPLACE",
	"TARGETING_CLAUSE_PAUSED_IN_AT_LEAST_ONE_MARKETPLACE",
	"TARGETING_CLAUSE_ARCHIVED",
	"TARGETING_CLAUSE_POLICING_SUSPENDED_IN_AT_LEAST_ONE_MARKETPLACE",
	"TARGETING_CLAUSE_CREATION_FAILED_IN_AT_LEAST_ONE_MARKETPLACE",
	"TARGETING_CLAUSE_BLOCKED_IN_AT_LEAST_ONE_MARKETPLACE",
}

func (v *SponsoredProductsGlobalKeywordServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsGlobalKeywordServingStatusReason(value)
	for _, existing := range AllowedSponsoredProductsGlobalKeywordServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsGlobalKeywordServingStatusReason", value)
}

// NewSponsoredProductsGlobalKeywordServingStatusReasonFromValue returns a pointer to a valid SponsoredProductsGlobalKeywordServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsGlobalKeywordServingStatusReasonFromValue(v string) (*SponsoredProductsGlobalKeywordServingStatusReason, error) {
	ev := SponsoredProductsGlobalKeywordServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsGlobalKeywordServingStatusReason: valid values are %v", v, AllowedSponsoredProductsGlobalKeywordServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsGlobalKeywordServingStatusReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsGlobalKeywordServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsGlobalKeywordServingStatusReason value
func (v SponsoredProductsGlobalKeywordServingStatusReason) Ptr() *SponsoredProductsGlobalKeywordServingStatusReason {
	return &v
}

type NullableSponsoredProductsGlobalKeywordServingStatusReason struct {
	value *SponsoredProductsGlobalKeywordServingStatusReason
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordServingStatusReason) Get() *SponsoredProductsGlobalKeywordServingStatusReason {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordServingStatusReason) Set(val *SponsoredProductsGlobalKeywordServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordServingStatusReason(val *SponsoredProductsGlobalKeywordServingStatusReason) *NullableSponsoredProductsGlobalKeywordServingStatusReason {
	return &NullableSponsoredProductsGlobalKeywordServingStatusReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
