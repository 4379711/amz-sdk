package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PortfolioDateErrorReason the model 'PortfolioDateErrorReason'
type PortfolioDateErrorReason string

// List of PortfolioDateErrorReason
const (
	PORTFOLIODATEERRORREASON_START_DATE_NOT_NULL           PortfolioDateErrorReason = "START_DATE_NOT_NULL"
	PORTFOLIODATEERRORREASON_INVALID_DATE                  PortfolioDateErrorReason = "INVALID_DATE"
	PORTFOLIODATEERRORREASON_START_DATE_EARLIER_THAN_TODAY PortfolioDateErrorReason = "START_DATE_EARLIER_THAN_TODAY"
	PORTFOLIODATEERRORREASON_END_DATE_EARLIER_THAN_TODAY   PortfolioDateErrorReason = "END_DATE_EARLIER_THAN_TODAY"
	PORTFOLIODATEERRORREASON_START_DATE_AFTER_END_DATE     PortfolioDateErrorReason = "START_DATE_AFTER_END_DATE"
	PORTFOLIODATEERRORREASON_START_DATE_EQUAL_END_DATE     PortfolioDateErrorReason = "START_DATE_EQUAL_END_DATE"
)

// All allowed values of PortfolioDateErrorReason enum
var AllowedPortfolioDateErrorReasonEnumValues = []PortfolioDateErrorReason{
	"START_DATE_NOT_NULL",
	"INVALID_DATE",
	"START_DATE_EARLIER_THAN_TODAY",
	"END_DATE_EARLIER_THAN_TODAY",
	"START_DATE_AFTER_END_DATE",
	"START_DATE_EQUAL_END_DATE",
}

func (v *PortfolioDateErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioDateErrorReason(value)
	for _, existing := range AllowedPortfolioDateErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioDateErrorReason", value)
}

// NewPortfolioDateErrorReasonFromValue returns a pointer to a valid PortfolioDateErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioDateErrorReasonFromValue(v string) (*PortfolioDateErrorReason, error) {
	ev := PortfolioDateErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioDateErrorReason: valid values are %v", v, AllowedPortfolioDateErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioDateErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioDateErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioDateErrorReason value
func (v PortfolioDateErrorReason) Ptr() *PortfolioDateErrorReason {
	return &v
}

type NullablePortfolioDateErrorReason struct {
	value *PortfolioDateErrorReason
	isSet bool
}

func (v NullablePortfolioDateErrorReason) Get() *PortfolioDateErrorReason {
	return v.value
}

func (v *NullablePortfolioDateErrorReason) Set(val *PortfolioDateErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioDateErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioDateErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioDateErrorReason(val *PortfolioDateErrorReason) *NullablePortfolioDateErrorReason {
	return &NullablePortfolioDateErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioDateErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioDateErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
