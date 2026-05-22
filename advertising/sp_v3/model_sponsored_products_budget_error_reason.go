package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsBudgetErrorReason the model 'SponsoredProductsBudgetErrorReason'
type SponsoredProductsBudgetErrorReason string

// List of SponsoredProductsBudgetErrorReason
const (
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGET_TOO_LOW                                      SponsoredProductsBudgetErrorReason = "BUDGET_TOO_LOW"
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGET_TOO_HIGH                                     SponsoredProductsBudgetErrorReason = "BUDGET_TOO_HIGH"
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGET_LT_DEFAULT_BIDS                              SponsoredProductsBudgetErrorReason = "BUDGET_LT_DEFAULT_BIDS"
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGET_LT_KEYWORD_BIDS                              SponsoredProductsBudgetErrorReason = "BUDGET_LT_KEYWORD_BIDS"
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGET_LT_PREDEFINED_TARGET_BIDS                    SponsoredProductsBudgetErrorReason = "BUDGET_LT_PREDEFINED_TARGET_BIDS"
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGET_OUT_OF_MARKET_PLACE_RANGE                    SponsoredProductsBudgetErrorReason = "BUDGET_OUT_OF_MARKET_PLACE_RANGE"
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGET_CURRENCY_DOES_NOT_MATCH_MARKETPLACE_SETTINGS SponsoredProductsBudgetErrorReason = "BUDGET_CURRENCY_DOES_NOT_MATCH_MARKETPLACE_SETTINGS"
	SPONSOREDPRODUCTSBUDGETERRORREASON_BUDGETING_POLICY_INVALID                            SponsoredProductsBudgetErrorReason = "BUDGETING_POLICY_INVALID"
	SPONSOREDPRODUCTSBUDGETERRORREASON_MISSING_BUDGETING_POLICY                            SponsoredProductsBudgetErrorReason = "MISSING_BUDGETING_POLICY"
	SPONSOREDPRODUCTSBUDGETERRORREASON_MISSING_IN_BUDGET_FLAG                              SponsoredProductsBudgetErrorReason = "MISSING_IN_BUDGET_FLAG"
)

// All allowed values of SponsoredProductsBudgetErrorReason enum
var AllowedSponsoredProductsBudgetErrorReasonEnumValues = []SponsoredProductsBudgetErrorReason{
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
}

func (v *SponsoredProductsBudgetErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsBudgetErrorReason(value)
	for _, existing := range AllowedSponsoredProductsBudgetErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsBudgetErrorReason", value)
}

// NewSponsoredProductsBudgetErrorReasonFromValue returns a pointer to a valid SponsoredProductsBudgetErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsBudgetErrorReasonFromValue(v string) (*SponsoredProductsBudgetErrorReason, error) {
	ev := SponsoredProductsBudgetErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsBudgetErrorReason: valid values are %v", v, AllowedSponsoredProductsBudgetErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsBudgetErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsBudgetErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsBudgetErrorReason value
func (v SponsoredProductsBudgetErrorReason) Ptr() *SponsoredProductsBudgetErrorReason {
	return &v
}

type NullableSponsoredProductsBudgetErrorReason struct {
	value *SponsoredProductsBudgetErrorReason
	isSet bool
}

func (v NullableSponsoredProductsBudgetErrorReason) Get() *SponsoredProductsBudgetErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsBudgetErrorReason) Set(val *SponsoredProductsBudgetErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBudgetErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBudgetErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBudgetErrorReason(val *SponsoredProductsBudgetErrorReason) *NullableSponsoredProductsBudgetErrorReason {
	return &NullableSponsoredProductsBudgetErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsBudgetErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBudgetErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
