package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioServingStatusReason the model 'PortfolioServingStatusReason'
type PortfolioServingStatusReason string

// List of PortfolioServingStatusReason
const (
	PORTFOLIOSERVINGSTATUSREASON_PORTFOLIO_STATUS_ENABLED_DETAIL         PortfolioServingStatusReason = "PORTFOLIO_STATUS_ENABLED_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_PORTFOLIO_PAUSED_DETAIL                 PortfolioServingStatusReason = "PORTFOLIO_PAUSED_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_PORTFOLIO_OUT_OF_BUDGET_DETAIL          PortfolioServingStatusReason = "PORTFOLIO_OUT_OF_BUDGET_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_PORTFOLIO_PENDING_START_DATE_DETAIL     PortfolioServingStatusReason = "PORTFOLIO_PENDING_START_DATE_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_PORTFOLIO_ENDED_DETAIL                  PortfolioServingStatusReason = "PORTFOLIO_ENDED_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_ADVERTISER_OUT_OF_BUDGET_DETAIL         PortfolioServingStatusReason = "ADVERTISER_OUT_OF_BUDGET_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_ADVERTISER_PAYMENT_FAILURE_DETAIL       PortfolioServingStatusReason = "ADVERTISER_PAYMENT_FAILURE_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_ADVERTISER_EXCEED_SPENDS_LIMIT_DETAIL   PortfolioServingStatusReason = "ADVERTISER_EXCEED_SPENDS_LIMIT_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_ADVERTISER_PAUSED_DETAIL                PortfolioServingStatusReason = "ADVERTISER_PAUSED_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_ADVERTISER_ARCHIVED_DETAIL              PortfolioServingStatusReason = "ADVERTISER_ARCHIVED_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_ADVERTISER_OUT_OF_PREPAY_BALANCE_DETAIL PortfolioServingStatusReason = "ADVERTISER_OUT_OF_PREPAY_BALANCE_DETAIL"
	PORTFOLIOSERVINGSTATUSREASON_ADVERTISER_ACCOUNT_OUT_OF_BUDGET_DETAIL PortfolioServingStatusReason = "ADVERTISER_ACCOUNT_OUT_OF_BUDGET_DETAIL"
)

// All allowed values of PortfolioServingStatusReason enum
var AllowedPortfolioServingStatusReasonEnumValues = []PortfolioServingStatusReason{
	"PORTFOLIO_STATUS_ENABLED_DETAIL",
	"PORTFOLIO_PAUSED_DETAIL",
	"PORTFOLIO_OUT_OF_BUDGET_DETAIL",
	"PORTFOLIO_PENDING_START_DATE_DETAIL",
	"PORTFOLIO_ENDED_DETAIL",
	"ADVERTISER_OUT_OF_BUDGET_DETAIL",
	"ADVERTISER_PAYMENT_FAILURE_DETAIL",
	"ADVERTISER_EXCEED_SPENDS_LIMIT_DETAIL",
	"ADVERTISER_PAUSED_DETAIL",
	"ADVERTISER_ARCHIVED_DETAIL",
	"ADVERTISER_OUT_OF_PREPAY_BALANCE_DETAIL",
	"ADVERTISER_ACCOUNT_OUT_OF_BUDGET_DETAIL",
}

func (v *PortfolioServingStatusReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioServingStatusReason(value)
	for _, existing := range AllowedPortfolioServingStatusReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioServingStatusReason", value)
}

// NewPortfolioServingStatusReasonFromValue returns a pointer to a valid PortfolioServingStatusReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioServingStatusReasonFromValue(v string) (*PortfolioServingStatusReason, error) {
	ev := PortfolioServingStatusReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioServingStatusReason: valid values are %v", v, AllowedPortfolioServingStatusReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioServingStatusReason) IsValid() bool {
	for _, existing := range AllowedPortfolioServingStatusReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioServingStatusReason value
func (v PortfolioServingStatusReason) Ptr() *PortfolioServingStatusReason {
	return &v
}

type NullablePortfolioServingStatusReason struct {
	value *PortfolioServingStatusReason
	isSet bool
}

func (v NullablePortfolioServingStatusReason) Get() *PortfolioServingStatusReason {
	return v.value
}

func (v *NullablePortfolioServingStatusReason) Set(val *PortfolioServingStatusReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioServingStatusReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioServingStatusReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioServingStatusReason(val *PortfolioServingStatusReason) *NullablePortfolioServingStatusReason {
	return &NullablePortfolioServingStatusReason{value: val, isSet: true}
}

func (v NullablePortfolioServingStatusReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioServingStatusReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
