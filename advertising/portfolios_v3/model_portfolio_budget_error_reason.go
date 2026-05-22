package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioBudgetErrorReason the model 'PortfolioBudgetErrorReason'
type PortfolioBudgetErrorReason string

// List of PortfolioBudgetErrorReason
const (
	PORTFOLIOBUDGETERRORREASON_BUDGET_TOO_LOW                                      PortfolioBudgetErrorReason = "BUDGET_TOO_LOW"
	PORTFOLIOBUDGETERRORREASON_BUDGET_TOO_HIGH                                     PortfolioBudgetErrorReason = "BUDGET_TOO_HIGH"
	PORTFOLIOBUDGETERRORREASON_BUDGET_LT_DEFAULT_BIDS                              PortfolioBudgetErrorReason = "BUDGET_LT_DEFAULT_BIDS"
	PORTFOLIOBUDGETERRORREASON_BUDGET_LT_KEYWORD_BIDS                              PortfolioBudgetErrorReason = "BUDGET_LT_KEYWORD_BIDS"
	PORTFOLIOBUDGETERRORREASON_BUDGET_LT_PREDEFINED_TARGET_BIDS                    PortfolioBudgetErrorReason = "BUDGET_LT_PREDEFINED_TARGET_BIDS"
	PORTFOLIOBUDGETERRORREASON_BUDGET_OUT_OF_MARKET_PLACE_RANGE                    PortfolioBudgetErrorReason = "BUDGET_OUT_OF_MARKET_PLACE_RANGE"
	PORTFOLIOBUDGETERRORREASON_BUDGET_CURRENCY_DOES_NOT_MATCH_MARKETPLACE_SETTINGS PortfolioBudgetErrorReason = "BUDGET_CURRENCY_DOES_NOT_MATCH_MARKETPLACE_SETTINGS"
	PORTFOLIOBUDGETERRORREASON_BUDGETING_POLICY_INVALID                            PortfolioBudgetErrorReason = "BUDGETING_POLICY_INVALID"
	PORTFOLIOBUDGETERRORREASON_MISSING_BUDGETING_POLICY                            PortfolioBudgetErrorReason = "MISSING_BUDGETING_POLICY"
	PORTFOLIOBUDGETERRORREASON_MISSING_IN_BUDGET_FLAG                              PortfolioBudgetErrorReason = "MISSING_IN_BUDGET_FLAG"
	PORTFOLIOBUDGETERRORREASON_BUDGET_AMOUNT_INVALID                               PortfolioBudgetErrorReason = "BUDGET_AMOUNT_INVALID"
)

// All allowed values of PortfolioBudgetErrorReason enum
var AllowedPortfolioBudgetErrorReasonEnumValues = []PortfolioBudgetErrorReason{
	"BUDGET_TOO_LOW",
	"BUDGET_TOO_HIGH",
	"BUDGET_LT_DEFAULT_BIDS",
	"BUDGET_LT_KEYWORD_BIDS",
	"BUDGET_LT_PREDEFINED_TARGET_BIDS",
	"BUDGET_OUT_OF_MARKET_PLACE_RANGE",
	"BUDGET_CURRENCY_DOES_NOT_MATCH_MARKETPLACE_SETTINGS",
	"BUDGETING_POLICY_INVALID",
	"MISSING_BUDGETING_POLICY",
	"MISSING_IN_BUDGET_FLAG",
	"BUDGET_AMOUNT_INVALID",
}

func (v *PortfolioBudgetErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioBudgetErrorReason(value)
	for _, existing := range AllowedPortfolioBudgetErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioBudgetErrorReason", value)
}

// NewPortfolioBudgetErrorReasonFromValue returns a pointer to a valid PortfolioBudgetErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioBudgetErrorReasonFromValue(v string) (*PortfolioBudgetErrorReason, error) {
	ev := PortfolioBudgetErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioBudgetErrorReason: valid values are %v", v, AllowedPortfolioBudgetErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioBudgetErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioBudgetErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioBudgetErrorReason value
func (v PortfolioBudgetErrorReason) Ptr() *PortfolioBudgetErrorReason {
	return &v
}

type NullablePortfolioBudgetErrorReason struct {
	value *PortfolioBudgetErrorReason
	isSet bool
}

func (v NullablePortfolioBudgetErrorReason) Get() *PortfolioBudgetErrorReason {
	return v.value
}

func (v *NullablePortfolioBudgetErrorReason) Set(val *PortfolioBudgetErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioBudgetErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioBudgetErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioBudgetErrorReason(val *PortfolioBudgetErrorReason) *NullablePortfolioBudgetErrorReason {
	return &NullablePortfolioBudgetErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioBudgetErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioBudgetErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
