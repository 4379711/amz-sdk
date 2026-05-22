package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioOtherErrorReason the model 'PortfolioOtherErrorReason'
type PortfolioOtherErrorReason string

// List of PortfolioOtherErrorReason
const (
	PORTFOLIOOTHERERRORREASON_OTHER_ERROR PortfolioOtherErrorReason = "OTHER_ERROR"
)

// All allowed values of PortfolioOtherErrorReason enum
var AllowedPortfolioOtherErrorReasonEnumValues = []PortfolioOtherErrorReason{
	"OTHER_ERROR",
}

func (v *PortfolioOtherErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioOtherErrorReason(value)
	for _, existing := range AllowedPortfolioOtherErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioOtherErrorReason", value)
}

// NewPortfolioOtherErrorReasonFromValue returns a pointer to a valid PortfolioOtherErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioOtherErrorReasonFromValue(v string) (*PortfolioOtherErrorReason, error) {
	ev := PortfolioOtherErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioOtherErrorReason: valid values are %v", v, AllowedPortfolioOtherErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioOtherErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioOtherErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioOtherErrorReason value
func (v PortfolioOtherErrorReason) Ptr() *PortfolioOtherErrorReason {
	return &v
}

type NullablePortfolioOtherErrorReason struct {
	value *PortfolioOtherErrorReason
	isSet bool
}

func (v NullablePortfolioOtherErrorReason) Get() *PortfolioOtherErrorReason {
	return v.value
}

func (v *NullablePortfolioOtherErrorReason) Set(val *PortfolioOtherErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioOtherErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioOtherErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioOtherErrorReason(val *PortfolioOtherErrorReason) *NullablePortfolioOtherErrorReason {
	return &NullablePortfolioOtherErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioOtherErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioOtherErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
