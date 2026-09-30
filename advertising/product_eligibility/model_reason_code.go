package product_eligibility

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ReasonCode the model 'ReasonCode'
type ReasonCode string

// List of ReasonCode
const (
	REASONCODE_BILLING_ACCOUNT_NOT_FOUND                          ReasonCode = "BILLING_ACCOUNT_NOT_FOUND"
	REASONCODE_PAYMENT_PROFILE_NOT_FOUND                          ReasonCode = "PAYMENT_PROFILE_NOT_FOUND"
	REASONCODE_PAYMENT_METHOD_NOT_FOUND                           ReasonCode = "PAYMENT_METHOD_NOT_FOUND"
	REASONCODE_PAYMENT_METHOD_NOT_VALID                           ReasonCode = "PAYMENT_METHOD_NOT_VALID"
	REASONCODE_EXPIRED_PAYMENT_METHOD                             ReasonCode = "EXPIRED_PAYMENT_METHOD"
	REASONCODE_VETTING_FAILURE                                    ReasonCode = "VETTING_FAILURE"
	REASONCODE_ACCOUNT_SUSPENDED                                  ReasonCode = "ACCOUNT_SUSPENDED"
	REASONCODE_TAX_INFO_NOT_COMPLETE                              ReasonCode = "TAX_INFO_NOT_COMPLETE"
	REASONCODE_PREPAY_BALANCE_TOO_LOW                             ReasonCode = "PREPAY_BALANCE_TOO_LOW"
	REASONCODE_RO_BALANCE_TOO_LOW                                 ReasonCode = "RO_BALANCE_TOO_LOW"
	REASONCODE_NO_BRAND_RELATIONS                                 ReasonCode = "NO_BRAND_RELATIONS"
	REASONCODE_MTA_NOT_ELIGIBLE                                   ReasonCode = "MTA_NOT_ELIGIBLE"
	REASONCODE_NOT_BRAND_REPRESENTATIVE                           ReasonCode = "NOT_BRAND_REPRESENTATIVE"
	REASONCODE_NO_TACTIC_ENABLED                                  ReasonCode = "NO_TACTIC_ENABLED"
	REASONCODE_DIRECT_TO_CONSUMER_OWNER_TAG_ID_NOT_FOUND          ReasonCode = "DIRECT_TO_CONSUMER_OWNER_TAG_ID_NOT_FOUND"
	REASONCODE_DIRECT_TO_CONSUMER_SUBSCRIPTION_NOT_FOUND          ReasonCode = "DIRECT_TO_CONSUMER_SUBSCRIPTION_NOT_FOUND"
	REASONCODE_SUBSCRIPTION_NOT_FOUND                             ReasonCode = "SUBSCRIPTION_NOT_FOUND"
	REASONCODE_ADVERTISING_ACCOUNT_NOT_FOUND                      ReasonCode = "ADVERTISING_ACCOUNT_NOT_FOUND"
	REASONCODE_NOT_LAUNCHED_IN_MARKETPLACE                        ReasonCode = "NOT_LAUNCHED_IN_MARKETPLACE"
	REASONCODE_UNKNOWN                                            ReasonCode = "UNKNOWN"
	REASONCODE_BLOCKED                                            ReasonCode = "BLOCKED"
	REASONCODE_ADVERTISER_TYPE_NOT_SUPPORTED                      ReasonCode = "ADVERTISER_TYPE_NOT_SUPPORTED"
	REASONCODE_BUSINESS_NOT_VERIFIED                              ReasonCode = "BUSINESS_NOT_VERIFIED"
	REASONCODE_BUSINESS_THRESHOLDS_NOT_MET                        ReasonCode = "BUSINESS_THRESHOLDS_NOT_MET"
	REASONCODE_ADS_TERMS_NOT_ACCEPTED                             ReasonCode = "ADS_TERMS_NOT_ACCEPTED"
	REASONCODE_AMAZON_BUSINESS_EXCLUSIVE_CAMPAIGN_NOT_ELIGIBLE    ReasonCode = "AMAZON_BUSINESS_EXCLUSIVE_CAMPAIGN_NOT_ELIGIBLE"
	REASONCODE_AMAZON_HAUL_EXCLUSIVE_CAMPAIGN_NOT_ELIGIBLE        ReasonCode = "AMAZON_HAUL_EXCLUSIVE_CAMPAIGN_NOT_ELIGIBLE"
	REASONCODE_AMAZON_MARKETING_CLOUD_ON_DEMAND_NOT_ELIGIBLE      ReasonCode = "AMAZON_MARKETING_CLOUD_ON_DEMAND_NOT_ELIGIBLE"
	REASONCODE_AUTONOMOUS_CAMPAIGNS_FEATURE_NOT_ELIGIBLE          ReasonCode = "AUTONOMOUS_CAMPAIGNS_FEATURE_NOT_ELIGIBLE"
	REASONCODE_BILL_TO_CC                                         ReasonCode = "BILL_TO_CC"
	REASONCODE_DSP_NOT_REQUESTED                                  ReasonCode = "DSP_NOT_REQUESTED"
	REASONCODE_DSP_PENDING_SETUP                                  ReasonCode = "DSP_PENDING_SETUP"
	REASONCODE_DSP_REQUEST_PENDING                                ReasonCode = "DSP_REQUEST_PENDING"
	REASONCODE_DSP_REQUEST_REJECTED                               ReasonCode = "DSP_REQUEST_REJECTED"
	REASONCODE_DVA_BUSINESS_VERIFICATION_NOT_COMPLETE             ReasonCode = "DVA_BUSINESS_VERIFICATION_NOT_COMPLETE"
	REASONCODE_DYNAMIC_PRODUCT_SETS_CAMPAIGN_FEATURE_NOT_ELIGIBLE ReasonCode = "DYNAMIC_PRODUCT_SETS_CAMPAIGN_FEATURE_NOT_ELIGIBLE"
	REASONCODE_EXPERT_CAMPAIGNS_FEATURE_NOT_ELIGIBLE              ReasonCode = "EXPERT_CAMPAIGNS_FEATURE_NOT_ELIGIBLE"
	REASONCODE_GEO_GATED_CAMPAIGN_FEATURE_NOT_ELIGIBLE            ReasonCode = "GEO_GATED_CAMPAIGN_FEATURE_NOT_ELIGIBLE"
	REASONCODE_GLOBAL_ACCOUNT_ALREADY_EXISTS                      ReasonCode = "GLOBAL_ACCOUNT_ALREADY_EXISTS"
	REASONCODE_GLOBAL_AUTO_SCALING_CAMPAIGNS_NOT_ELIGIBLE         ReasonCode = "GLOBAL_AUTO_SCALING_CAMPAIGNS_NOT_ELIGIBLE"
	REASONCODE_GLOBAL_CAMPAIGNS_NOT_ELIGIBLE                      ReasonCode = "GLOBAL_CAMPAIGNS_NOT_ELIGIBLE"
	REASONCODE_KDP_AUTHOR_NOT_ELIGIBLE                            ReasonCode = "KDP_AUTHOR_NOT_ELIGIBLE"
	REASONCODE_MULTIPLE_BILLING_PROFILES_FOUND                    ReasonCode = "MULTIPLE_BILLING_PROFILES_FOUND"
	REASONCODE_NOT_SETUP_FOR_DSP                                  ReasonCode = "NOT_SETUP_FOR_DSP"
	REASONCODE_NO_BILLING_PROFILES_FOUND                          ReasonCode = "NO_BILLING_PROFILES_FOUND"
	REASONCODE_SMART_CAMPAIGNS_FEATURE_NOT_ELIGIBLE               ReasonCode = "SMART_CAMPAIGNS_FEATURE_NOT_ELIGIBLE"
	REASONCODE_STOCK_FILTER_CAMPAIGN_FEATURE_NOT_ELIGIBLE         ReasonCode = "STOCK_FILTER_CAMPAIGN_FEATURE_NOT_ELIGIBLE"
)

// All allowed values of ReasonCode enum
var AllowedReasonCodeEnumValues = []ReasonCode{
	"BILLING_ACCOUNT_NOT_FOUND",
	"PAYMENT_PROFILE_NOT_FOUND",
	"PAYMENT_METHOD_NOT_FOUND",
	"PAYMENT_METHOD_NOT_VALID",
	"EXPIRED_PAYMENT_METHOD",
	"VETTING_FAILURE",
	"ACCOUNT_SUSPENDED",
	"TAX_INFO_NOT_COMPLETE",
	"PREPAY_BALANCE_TOO_LOW",
	"RO_BALANCE_TOO_LOW",
	"NO_BRAND_RELATIONS",
	"MTA_NOT_ELIGIBLE",
	"NOT_BRAND_REPRESENTATIVE",
	"NO_TACTIC_ENABLED",
	"DIRECT_TO_CONSUMER_OWNER_TAG_ID_NOT_FOUND",
	"DIRECT_TO_CONSUMER_SUBSCRIPTION_NOT_FOUND",
	"SUBSCRIPTION_NOT_FOUND",
	"ADVERTISING_ACCOUNT_NOT_FOUND",
	"NOT_LAUNCHED_IN_MARKETPLACE",
	"UNKNOWN",
	"BLOCKED",
	"ADVERTISER_TYPE_NOT_SUPPORTED",
	"BUSINESS_NOT_VERIFIED",
	"BUSINESS_THRESHOLDS_NOT_MET",
	"ADS_TERMS_NOT_ACCEPTED",
	"AMAZON_BUSINESS_EXCLUSIVE_CAMPAIGN_NOT_ELIGIBLE",
	"AMAZON_HAUL_EXCLUSIVE_CAMPAIGN_NOT_ELIGIBLE",
	"AMAZON_MARKETING_CLOUD_ON_DEMAND_NOT_ELIGIBLE",
	"AUTONOMOUS_CAMPAIGNS_FEATURE_NOT_ELIGIBLE",
	"BILL_TO_CC",
	"DSP_NOT_REQUESTED",
	"DSP_PENDING_SETUP",
	"DSP_REQUEST_PENDING",
	"DSP_REQUEST_REJECTED",
	"DVA_BUSINESS_VERIFICATION_NOT_COMPLETE",
	"DYNAMIC_PRODUCT_SETS_CAMPAIGN_FEATURE_NOT_ELIGIBLE",
	"EXPERT_CAMPAIGNS_FEATURE_NOT_ELIGIBLE",
	"GEO_GATED_CAMPAIGN_FEATURE_NOT_ELIGIBLE",
	"GLOBAL_ACCOUNT_ALREADY_EXISTS",
	"GLOBAL_AUTO_SCALING_CAMPAIGNS_NOT_ELIGIBLE",
	"GLOBAL_CAMPAIGNS_NOT_ELIGIBLE",
	"KDP_AUTHOR_NOT_ELIGIBLE",
	"MULTIPLE_BILLING_PROFILES_FOUND",
	"NOT_SETUP_FOR_DSP",
	"NO_BILLING_PROFILES_FOUND",
	"SMART_CAMPAIGNS_FEATURE_NOT_ELIGIBLE",
	"STOCK_FILTER_CAMPAIGN_FEATURE_NOT_ELIGIBLE",
}

func (v *ReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = ReasonCode(value)
	return nil
}

// NewReasonCodeFromValue returns a pointer to a valid ReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReasonCodeFromValue(v string) (*ReasonCode, error) {
	ev := ReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReasonCode: valid values are %v", v, AllowedReasonCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReasonCode) IsValid() bool {
	for _, existing := range AllowedReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReasonCode value
func (v ReasonCode) Ptr() *ReasonCode {
	return &v
}

type NullableReasonCode struct {
	value *ReasonCode
	isSet bool
}

func (v NullableReasonCode) Get() *ReasonCode {
	return v.value
}

func (v *NullableReasonCode) Set(val *ReasonCode) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonCode) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonCode(val *ReasonCode) *NullableReasonCode {
	return &NullableReasonCode{value: val, isSet: true}
}

func (v NullableReasonCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReasonCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
