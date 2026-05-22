package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PortfolioBillingErrorReason the model 'PortfolioBillingErrorReason'
type PortfolioBillingErrorReason string

// List of PortfolioBillingErrorReason
const (
	PORTFOLIOBILLINGERRORREASON_ADVERTISER_SUSPENDED      PortfolioBillingErrorReason = "ADVERTISER_SUSPENDED"
	PORTFOLIOBILLINGERRORREASON_BILLING_ACCOUNT_NOT_FOUND PortfolioBillingErrorReason = "BILLING_ACCOUNT_NOT_FOUND"
	PORTFOLIOBILLINGERRORREASON_PAYMENT_PROFILE_NOT_FOUND PortfolioBillingErrorReason = "PAYMENT_PROFILE_NOT_FOUND"
	PORTFOLIOBILLINGERRORREASON_EXPIRED_PAYMENT_METHOD    PortfolioBillingErrorReason = "EXPIRED_PAYMENT_METHOD"
	PORTFOLIOBILLINGERRORREASON_VETTING_FAILURE           PortfolioBillingErrorReason = "VETTING_FAILURE"
)

// All allowed values of PortfolioBillingErrorReason enum
var AllowedPortfolioBillingErrorReasonEnumValues = []PortfolioBillingErrorReason{
	"ADVERTISER_SUSPENDED",
	"BILLING_ACCOUNT_NOT_FOUND",
	"PAYMENT_PROFILE_NOT_FOUND",
	"EXPIRED_PAYMENT_METHOD",
	"VETTING_FAILURE",
}

func (v *PortfolioBillingErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioBillingErrorReason(value)
	for _, existing := range AllowedPortfolioBillingErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioBillingErrorReason", value)
}

// NewPortfolioBillingErrorReasonFromValue returns a pointer to a valid PortfolioBillingErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioBillingErrorReasonFromValue(v string) (*PortfolioBillingErrorReason, error) {
	ev := PortfolioBillingErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioBillingErrorReason: valid values are %v", v, AllowedPortfolioBillingErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioBillingErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioBillingErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioBillingErrorReason value
func (v PortfolioBillingErrorReason) Ptr() *PortfolioBillingErrorReason {
	return &v
}

type NullablePortfolioBillingErrorReason struct {
	value *PortfolioBillingErrorReason
	isSet bool
}

func (v NullablePortfolioBillingErrorReason) Get() *PortfolioBillingErrorReason {
	return v.value
}

func (v *NullablePortfolioBillingErrorReason) Set(val *PortfolioBillingErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioBillingErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioBillingErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioBillingErrorReason(val *PortfolioBillingErrorReason) *NullablePortfolioBillingErrorReason {
	return &NullablePortfolioBillingErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioBillingErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioBillingErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
