package product_eligibility

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// IneligibleLevel the model 'IneligibleLevel'
type IneligibleLevel string

// List of IneligibleLevel
const (
	INELIGIBLELEVEL_INELIGIBLE_WITH_RESOLUTION IneligibleLevel = "INELIGIBLE_WITH_RESOLUTION"
	INELIGIBLELEVEL_INELIGIBLE                 IneligibleLevel = "INELIGIBLE"
)

// All allowed values of IneligibleLevel enum
var AllowedIneligibleLevelEnumValues = []IneligibleLevel{
	"INELIGIBLE_WITH_RESOLUTION",
	"INELIGIBLE",
}

func (v *IneligibleLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IneligibleLevel(value)
	for _, existing := range AllowedIneligibleLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IneligibleLevel", value)
}

// NewIneligibleLevelFromValue returns a pointer to a valid IneligibleLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIneligibleLevelFromValue(v string) (*IneligibleLevel, error) {
	ev := IneligibleLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IneligibleLevel: valid values are %v", v, AllowedIneligibleLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IneligibleLevel) IsValid() bool {
	for _, existing := range AllowedIneligibleLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IneligibleLevel value
func (v IneligibleLevel) Ptr() *IneligibleLevel {
	return &v
}

type NullableIneligibleLevel struct {
	value *IneligibleLevel
	isSet bool
}

func (v NullableIneligibleLevel) Get() *IneligibleLevel {
	return v.value
}

func (v *NullableIneligibleLevel) Set(val *IneligibleLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableIneligibleLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableIneligibleLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIneligibleLevel(val *IneligibleLevel) *NullableIneligibleLevel {
	return &NullableIneligibleLevel{value: val, isSet: true}
}

func (v NullableIneligibleLevel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIneligibleLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
