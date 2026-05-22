package portfolios_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// PortfolioDuplicateValueErrorReason the model 'PortfolioDuplicateValueErrorReason'
type PortfolioDuplicateValueErrorReason string

// List of PortfolioDuplicateValueErrorReason
const (
	PORTFOLIODUPLICATEVALUEERRORREASON_DUPLICATE_VALUE PortfolioDuplicateValueErrorReason = "DUPLICATE_VALUE"
)

// All allowed values of PortfolioDuplicateValueErrorReason enum
var AllowedPortfolioDuplicateValueErrorReasonEnumValues = []PortfolioDuplicateValueErrorReason{
	"DUPLICATE_VALUE",
}

func (v *PortfolioDuplicateValueErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortfolioDuplicateValueErrorReason(value)
	for _, existing := range AllowedPortfolioDuplicateValueErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortfolioDuplicateValueErrorReason", value)
}

// NewPortfolioDuplicateValueErrorReasonFromValue returns a pointer to a valid PortfolioDuplicateValueErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortfolioDuplicateValueErrorReasonFromValue(v string) (*PortfolioDuplicateValueErrorReason, error) {
	ev := PortfolioDuplicateValueErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortfolioDuplicateValueErrorReason: valid values are %v", v, AllowedPortfolioDuplicateValueErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortfolioDuplicateValueErrorReason) IsValid() bool {
	for _, existing := range AllowedPortfolioDuplicateValueErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortfolioDuplicateValueErrorReason value
func (v PortfolioDuplicateValueErrorReason) Ptr() *PortfolioDuplicateValueErrorReason {
	return &v
}

type NullablePortfolioDuplicateValueErrorReason struct {
	value *PortfolioDuplicateValueErrorReason
	isSet bool
}

func (v NullablePortfolioDuplicateValueErrorReason) Get() *PortfolioDuplicateValueErrorReason {
	return v.value
}

func (v *NullablePortfolioDuplicateValueErrorReason) Set(val *PortfolioDuplicateValueErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioDuplicateValueErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioDuplicateValueErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioDuplicateValueErrorReason(val *PortfolioDuplicateValueErrorReason) *NullablePortfolioDuplicateValueErrorReason {
	return &NullablePortfolioDuplicateValueErrorReason{value: val, isSet: true}
}

func (v NullablePortfolioDuplicateValueErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioDuplicateValueErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
