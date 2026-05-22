package portfolios_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PortfolioMalformedValueErrorReason the model 'PortfolioMalformedValueErrorReason'
type PortfolioMalformedValueErrorReason string

// List of PortfolioMalformedValueErrorReason
const (
	PORTFOLIOMALFORMEDVALUEERRORREASON_FORBIDDEN_CHARS     PortfolioMalformedValueErrorReason = "FORBIDDEN_CHARS"
	PORTFOLIOMALFORMEDVALUEERRORREASON_PATTERN_NOT_MATCHED PortfolioMalformedValueErrorReason = "PATTERN_NOT_MATCHED"
	PORTFOLIOMALFORMEDVALUEERRORREASON_TOO_LONG            PortfolioMalformedValueErrorReason = "TOO_LONG"
	PORTFOLIOMALFORMEDVALUEERRORREASON_TOO_SHORT           PortfolioMalformedValueErrorReason = "TOO_SHORT"
)

// All allowed values of PortfolioMalformedValueErrorReason enum
var AllowedPortfolioMalformedValueErrorReasonEnumValues = []PortfolioMalformedValueErrorReason{
	"FORBIDDEN_CHARS",
	"PATTERN_NOT_MATCHED",
	"TOO_LONG",
	"TOO_SHORT",
}

func (v *PortfolioMalformedValueErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioMalformedValueErrorReason(value)
	for _, existing := range AllowedPortfolioMalformedValueErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioMalformedValueErrorReason", value)
}

// NewPortfolioMalformedValueErrorReasonFromValue returns a pointer to a valid PortfolioMalformedValueErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioMalformedValueErrorReasonFromValue(v string) (*PortfolioMalformedValueErrorReason, error) {
	ev := PortfolioMalformedValueErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioMalformedValueErrorReason: valid values are %v", v, AllowedPortfolioMalformedValueErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioMalformedValueErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioMalformedValueErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioMalformedValueErrorReason value
func (v PortfolioMalformedValueErrorReason) Ptr() *PortfolioMalformedValueErrorReason {
	return &v
}

type NullablePortfolioMalformedValueErrorReason struct {
	value *PortfolioMalformedValueErrorReason
	isSet bool
}

func (v NullablePortfolioMalformedValueErrorReason) Get() *PortfolioMalformedValueErrorReason {
	return v.value
}

func (v *NullablePortfolioMalformedValueErrorReason) Set(val *PortfolioMalformedValueErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMalformedValueErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMalformedValueErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMalformedValueErrorReason(val *PortfolioMalformedValueErrorReason) *NullablePortfolioMalformedValueErrorReason {
	return &NullablePortfolioMalformedValueErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioMalformedValueErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioMalformedValueErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
