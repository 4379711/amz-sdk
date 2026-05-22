package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioEntityNotFoundErrorReason the model 'PortfolioEntityNotFoundErrorReason'
type PortfolioEntityNotFoundErrorReason string

// List of PortfolioEntityNotFoundErrorReason
const (
	PORTFOLIOENTITYNOTFOUNDERRORREASON_ENTITY_NOT_FOUND PortfolioEntityNotFoundErrorReason = "ENTITY_NOT_FOUND"
)

// All allowed values of PortfolioEntityNotFoundErrorReason enum
var AllowedPortfolioEntityNotFoundErrorReasonEnumValues = []PortfolioEntityNotFoundErrorReason{
	"ENTITY_NOT_FOUND",
}

func (v *PortfolioEntityNotFoundErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioEntityNotFoundErrorReason(value)
	for _, existing := range AllowedPortfolioEntityNotFoundErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioEntityNotFoundErrorReason", value)
}

// NewPortfolioEntityNotFoundErrorReasonFromValue returns a pointer to a valid PortfolioEntityNotFoundErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioEntityNotFoundErrorReasonFromValue(v string) (*PortfolioEntityNotFoundErrorReason, error) {
	ev := PortfolioEntityNotFoundErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioEntityNotFoundErrorReason: valid values are %v", v, AllowedPortfolioEntityNotFoundErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioEntityNotFoundErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioEntityNotFoundErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioEntityNotFoundErrorReason value
func (v PortfolioEntityNotFoundErrorReason) Ptr() *PortfolioEntityNotFoundErrorReason {
	return &v
}

type NullablePortfolioEntityNotFoundErrorReason struct {
	value *PortfolioEntityNotFoundErrorReason
	isSet bool
}

func (v NullablePortfolioEntityNotFoundErrorReason) Get() *PortfolioEntityNotFoundErrorReason {
	return v.value
}

func (v *NullablePortfolioEntityNotFoundErrorReason) Set(val *PortfolioEntityNotFoundErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioEntityNotFoundErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioEntityNotFoundErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioEntityNotFoundErrorReason(val *PortfolioEntityNotFoundErrorReason) *NullablePortfolioEntityNotFoundErrorReason {
	return &NullablePortfolioEntityNotFoundErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioEntityNotFoundErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioEntityNotFoundErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
