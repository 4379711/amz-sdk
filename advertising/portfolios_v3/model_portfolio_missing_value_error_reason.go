package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioMissingValueErrorReason the model 'PortfolioMissingValueErrorReason'
type PortfolioMissingValueErrorReason string

// List of PortfolioMissingValueErrorReason
const (
	PORTFOLIOMISSINGVALUEERRORREASON_MISSING_VALUE PortfolioMissingValueErrorReason = "MISSING_VALUE"
)

// All allowed values of PortfolioMissingValueErrorReason enum
var AllowedPortfolioMissingValueErrorReasonEnumValues = []PortfolioMissingValueErrorReason{
	"MISSING_VALUE",
}

func (v *PortfolioMissingValueErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioMissingValueErrorReason(value)
	for _, existing := range AllowedPortfolioMissingValueErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioMissingValueErrorReason", value)
}

// NewPortfolioMissingValueErrorReasonFromValue returns a pointer to a valid PortfolioMissingValueErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioMissingValueErrorReasonFromValue(v string) (*PortfolioMissingValueErrorReason, error) {
	ev := PortfolioMissingValueErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioMissingValueErrorReason: valid values are %v", v, AllowedPortfolioMissingValueErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioMissingValueErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioMissingValueErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioMissingValueErrorReason value
func (v PortfolioMissingValueErrorReason) Ptr() *PortfolioMissingValueErrorReason {
	return &v
}

type NullablePortfolioMissingValueErrorReason struct {
	value *PortfolioMissingValueErrorReason
	isSet bool
}

func (v NullablePortfolioMissingValueErrorReason) Get() *PortfolioMissingValueErrorReason {
	return v.value
}

func (v *NullablePortfolioMissingValueErrorReason) Set(val *PortfolioMissingValueErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMissingValueErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMissingValueErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMissingValueErrorReason(val *PortfolioMissingValueErrorReason) *NullablePortfolioMissingValueErrorReason {
	return &NullablePortfolioMissingValueErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioMissingValueErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioMissingValueErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
