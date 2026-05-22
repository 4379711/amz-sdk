package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioServingStatus the model 'PortfolioServingStatus'
type PortfolioServingStatus string

// List of PortfolioServingStatus
const (
	PORTFOLIOSERVINGSTATUS_PORTFOLIO_STATUS_ENABLED         PortfolioServingStatus = "PORTFOLIO_STATUS_ENABLED"
	PORTFOLIOSERVINGSTATUS_PORTFOLIO_PAUSED                 PortfolioServingStatus = "PORTFOLIO_PAUSED"
	PORTFOLIOSERVINGSTATUS_PORTFOLIO_OUT_OF_BUDGET          PortfolioServingStatus = "PORTFOLIO_OUT_OF_BUDGET"
	PORTFOLIOSERVINGSTATUS_PORTFOLIO_PENDING_START_DATE     PortfolioServingStatus = "PORTFOLIO_PENDING_START_DATE"
	PORTFOLIOSERVINGSTATUS_PORTFOLIO_ENDED                  PortfolioServingStatus = "PORTFOLIO_ENDED"
	PORTFOLIOSERVINGSTATUS_ADVERTISER_OUT_OF_BUDGET         PortfolioServingStatus = "ADVERTISER_OUT_OF_BUDGET"
	PORTFOLIOSERVINGSTATUS_ADVERTISER_PAYMENT_FAILURE       PortfolioServingStatus = "ADVERTISER_PAYMENT_FAILURE"
	PORTFOLIOSERVINGSTATUS_ADVERTISER_EXCEED_SPENDS_LIMIT   PortfolioServingStatus = "ADVERTISER_EXCEED_SPENDS_LIMIT"
	PORTFOLIOSERVINGSTATUS_ADVERTISER_PAUSED                PortfolioServingStatus = "ADVERTISER_PAUSED"
	PORTFOLIOSERVINGSTATUS_ADVERTISER_ARCHIVED              PortfolioServingStatus = "ADVERTISER_ARCHIVED"
	PORTFOLIOSERVINGSTATUS_ADVERTISER_OUT_OF_PREPAY_BALANCE PortfolioServingStatus = "ADVERTISER_OUT_OF_PREPAY_BALANCE"
	PORTFOLIOSERVINGSTATUS_ADVERTISER_ACCOUNT_OUT_OF_BUDGET PortfolioServingStatus = "ADVERTISER_ACCOUNT_OUT_OF_BUDGET"
)

// All allowed values of PortfolioServingStatus enum
var AllowedPortfolioServingStatusEnumValues = []PortfolioServingStatus{
	"PORTFOLIO_STATUS_ENABLED",
	"PORTFOLIO_PAUSED",
	"PORTFOLIO_OUT_OF_BUDGET",
	"PORTFOLIO_PENDING_START_DATE",
	"PORTFOLIO_ENDED",
	"ADVERTISER_OUT_OF_BUDGET",
	"ADVERTISER_PAYMENT_FAILURE",
	"ADVERTISER_EXCEED_SPENDS_LIMIT",
	"ADVERTISER_PAUSED",
	"ADVERTISER_ARCHIVED",
	"ADVERTISER_OUT_OF_PREPAY_BALANCE",
	"ADVERTISER_ACCOUNT_OUT_OF_BUDGET",
}

func (v *PortfolioServingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioServingStatus(value)
	for _, existing := range AllowedPortfolioServingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioServingStatus", value)
}

// NewPortfolioServingStatusFromValue returns a pointer to a valid PortfolioServingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioServingStatusFromValue(v string) (*PortfolioServingStatus, error) {
	ev := PortfolioServingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioServingStatus: valid values are %v", v, AllowedPortfolioServingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioServingStatus) IsValid() bool {
	for _, existing := range AllowedPortfolioServingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioServingStatus value
func (v PortfolioServingStatus) Ptr() *PortfolioServingStatus {
	return &v
}

type NullablePortfolioServingStatus struct {
	value *PortfolioServingStatus
	isSet bool
}

func (v NullablePortfolioServingStatus) Get() *PortfolioServingStatus {
	return v.value
}

func (v *NullablePortfolioServingStatus) Set(val *PortfolioServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioServingStatus(val *PortfolioServingStatus) *NullablePortfolioServingStatus {
	return &NullablePortfolioServingStatus{value: val, isSet: true}
}

func (v NullablePortfolioServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
