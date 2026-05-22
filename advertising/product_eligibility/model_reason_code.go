package product_eligibility

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ReasonCode the model 'ReasonCode'
type ReasonCode string

// List of ReasonCode
const (
	REASONCODE_BILLING_ACCOUNT_NOT_FOUND                 ReasonCode = "BILLING_ACCOUNT_NOT_FOUND"
	REASONCODE_PAYMENT_PROFILE_NOT_FOUND                 ReasonCode = "PAYMENT_PROFILE_NOT_FOUND"
	REASONCODE_PAYMENT_METHOD_NOT_FOUND                  ReasonCode = "PAYMENT_METHOD_NOT_FOUND"
	REASONCODE_PAYMENT_METHOD_NOT_VALID                  ReasonCode = "PAYMENT_METHOD_NOT_VALID"
	REASONCODE_EXPIRED_PAYMENT_METHOD                    ReasonCode = "EXPIRED_PAYMENT_METHOD"
	REASONCODE_VETTING_FAILURE                           ReasonCode = "VETTING_FAILURE"
	REASONCODE_ACCOUNT_SUSPENDED                         ReasonCode = "ACCOUNT_SUSPENDED"
	REASONCODE_TAX_INFO_NOT_COMPLETE                     ReasonCode = "TAX_INFO_NOT_COMPLETE"
	REASONCODE_PREPAY_BALANCE_TOO_LOW                    ReasonCode = "PREPAY_BALANCE_TOO_LOW"
	REASONCODE_RO_BALANCE_TOO_LOW                        ReasonCode = "RO_BALANCE_TOO_LOW"
	REASONCODE_NO_BRAND_RELATIONS                        ReasonCode = "NO_BRAND_RELATIONS"
	REASONCODE_MTA_NOT_ELIGIBLE                          ReasonCode = "MTA_NOT_ELIGIBLE"
	REASONCODE_NOT_BRAND_REPRESENTATIVE                  ReasonCode = "NOT_BRAND_REPRESENTATIVE"
	REASONCODE_NO_TACTIC_ENABLED                         ReasonCode = "NO_TACTIC_ENABLED"
	REASONCODE_DIRECT_TO_CONSUMER_OWNER_TAG_ID_NOT_FOUND ReasonCode = "DIRECT_TO_CONSUMER_OWNER_TAG_ID_NOT_FOUND"
	REASONCODE_DIRECT_TO_CONSUMER_SUBSCRIPTION_NOT_FOUND ReasonCode = "DIRECT_TO_CONSUMER_SUBSCRIPTION_NOT_FOUND"
	REASONCODE_SUBSCRIPTION_NOT_FOUND                    ReasonCode = "SUBSCRIPTION_NOT_FOUND"
	REASONCODE_ADVERTISING_ACCOUNT_NOT_FOUND             ReasonCode = "ADVERTISING_ACCOUNT_NOT_FOUND"
	REASONCODE_NOT_LAUNCHED_IN_MARKETPLACE               ReasonCode = "NOT_LAUNCHED_IN_MARKETPLACE"
	REASONCODE_UNKNOWN                                   ReasonCode = "UNKNOWN"
	REASONCODE_BLOCKED                                   ReasonCode = "BLOCKED"
	REASONCODE_ADVERTISER_TYPE_NOT_SUPPORTED             ReasonCode = "ADVERTISER_TYPE_NOT_SUPPORTED"
	REASONCODE_BUSINESS_NOT_VERIFIED                     ReasonCode = "BUSINESS_NOT_VERIFIED"
	REASONCODE_BUSINESS_THRESHOLDS_NOT_MET               ReasonCode = "BUSINESS_THRESHOLDS_NOT_MET"
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
}

func (v *ReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReasonCode(value)
	for _, existing := range AllowedReasonCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReasonCode", value)
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
