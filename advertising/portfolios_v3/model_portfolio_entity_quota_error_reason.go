package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioEntityQuotaErrorReason the model 'PortfolioEntityQuotaErrorReason'
type PortfolioEntityQuotaErrorReason string

// List of PortfolioEntityQuotaErrorReason
const (
	PORTFOLIOENTITYQUOTAERRORREASON_QUOTA_EXCEEDED PortfolioEntityQuotaErrorReason = "QUOTA_EXCEEDED"
)

// All allowed values of PortfolioEntityQuotaErrorReason enum
var AllowedPortfolioEntityQuotaErrorReasonEnumValues = []PortfolioEntityQuotaErrorReason{
	"QUOTA_EXCEEDED",
}

func (v *PortfolioEntityQuotaErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioEntityQuotaErrorReason(value)
	for _, existing := range AllowedPortfolioEntityQuotaErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioEntityQuotaErrorReason", value)
}

// NewPortfolioEntityQuotaErrorReasonFromValue returns a pointer to a valid PortfolioEntityQuotaErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioEntityQuotaErrorReasonFromValue(v string) (*PortfolioEntityQuotaErrorReason, error) {
	ev := PortfolioEntityQuotaErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioEntityQuotaErrorReason: valid values are %v", v, AllowedPortfolioEntityQuotaErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioEntityQuotaErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioEntityQuotaErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioEntityQuotaErrorReason value
func (v PortfolioEntityQuotaErrorReason) Ptr() *PortfolioEntityQuotaErrorReason {
	return &v
}

type NullablePortfolioEntityQuotaErrorReason struct {
	value *PortfolioEntityQuotaErrorReason
	isSet bool
}

func (v NullablePortfolioEntityQuotaErrorReason) Get() *PortfolioEntityQuotaErrorReason {
	return v.value
}

func (v *NullablePortfolioEntityQuotaErrorReason) Set(val *PortfolioEntityQuotaErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioEntityQuotaErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioEntityQuotaErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioEntityQuotaErrorReason(val *PortfolioEntityQuotaErrorReason) *NullablePortfolioEntityQuotaErrorReason {
	return &NullablePortfolioEntityQuotaErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioEntityQuotaErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioEntityQuotaErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
