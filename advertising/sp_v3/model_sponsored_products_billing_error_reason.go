package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsBillingErrorReason the model 'SponsoredProductsBillingErrorReason'
type SponsoredProductsBillingErrorReason string

// List of SponsoredProductsBillingErrorReason
const (
	SPONSOREDPRODUCTSBILLINGERRORREASON_ADVERTISER_SUSPENDED                SponsoredProductsBillingErrorReason = "ADVERTISER_SUSPENDED"
	SPONSOREDPRODUCTSBILLINGERRORREASON_BILLING_ACCOUNT_NOT_FOUND           SponsoredProductsBillingErrorReason = "BILLING_ACCOUNT_NOT_FOUND"
	SPONSOREDPRODUCTSBILLINGERRORREASON_PAYMENT_PROFILE_NOT_FOUND           SponsoredProductsBillingErrorReason = "PAYMENT_PROFILE_NOT_FOUND"
	SPONSOREDPRODUCTSBILLINGERRORREASON_EXPIRED_PAYMENT_METHOD              SponsoredProductsBillingErrorReason = "EXPIRED_PAYMENT_METHOD"
	SPONSOREDPRODUCTSBILLINGERRORREASON_VETTING_FAILURE                     SponsoredProductsBillingErrorReason = "VETTING_FAILURE"
	SPONSOREDPRODUCTSBILLINGERRORREASON_ADVERTISER_BILLING_SETUP_INCOMPLETE SponsoredProductsBillingErrorReason = "ADVERTISER_BILLING_SETUP_INCOMPLETE"
)

// All allowed values of SponsoredProductsBillingErrorReason enum
var AllowedSponsoredProductsBillingErrorReasonEnumValues = []SponsoredProductsBillingErrorReason{
	"ADVERTISER_SUSPENDED",
	"BILLING_ACCOUNT_NOT_FOUND",
	"PAYMENT_PROFILE_NOT_FOUND",
	"EXPIRED_PAYMENT_METHOD",
	"VETTING_FAILURE",
	"ADVERTISER_BILLING_SETUP_INCOMPLETE",
}

func (v *SponsoredProductsBillingErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsBillingErrorReason(value)
	for _, existing := range AllowedSponsoredProductsBillingErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsBillingErrorReason", value)
}

// NewSponsoredProductsBillingErrorReasonFromValue returns a pointer to a valid SponsoredProductsBillingErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsBillingErrorReasonFromValue(v string) (*SponsoredProductsBillingErrorReason, error) {
	ev := SponsoredProductsBillingErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsBillingErrorReason: valid values are %v", v, AllowedSponsoredProductsBillingErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsBillingErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsBillingErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsBillingErrorReason value
func (v SponsoredProductsBillingErrorReason) Ptr() *SponsoredProductsBillingErrorReason {
	return &v
}

type NullableSponsoredProductsBillingErrorReason struct {
	value *SponsoredProductsBillingErrorReason
	isSet bool
}

func (v NullableSponsoredProductsBillingErrorReason) Get() *SponsoredProductsBillingErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsBillingErrorReason) Set(val *SponsoredProductsBillingErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBillingErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBillingErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBillingErrorReason(val *SponsoredProductsBillingErrorReason) *NullableSponsoredProductsBillingErrorReason {
	return &NullableSponsoredProductsBillingErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsBillingErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBillingErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
