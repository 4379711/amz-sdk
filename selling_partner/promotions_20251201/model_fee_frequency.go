package promotions_20251201

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// FeeFrequency For participation fees, the frequency at which fees are charged.
type FeeFrequency string

// List of FeeFrequency
const (
	FEEFREQUENCY_ONE_TIME FeeFrequency = "ONE_TIME"
	FEEFREQUENCY_DAILY    FeeFrequency = "DAILY"
)

// All allowed values of FeeFrequency enum
var AllowedFeeFrequencyEnumValues = []FeeFrequency{
	"ONE_TIME",
	"DAILY",
}

func (v *FeeFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeeFrequency(value)
	for _, existing := range AllowedFeeFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeeFrequency", value)
}

// NewFeeFrequencyFromValue returns a pointer to a valid FeeFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeeFrequencyFromValue(v string) (*FeeFrequency, error) {
	ev := FeeFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeeFrequency: valid values are %v", v, AllowedFeeFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeeFrequency) IsValid() bool {
	for _, existing := range AllowedFeeFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeeFrequency value
func (v FeeFrequency) Ptr() *FeeFrequency {
	return &v
}

type NullableFeeFrequency struct {
	value *FeeFrequency
	isSet bool
}

func (v NullableFeeFrequency) Get() *FeeFrequency {
	return v.value
}

func (v *NullableFeeFrequency) Set(val *FeeFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeFrequency(val *FeeFrequency) *NullableFeeFrequency {
	return &NullableFeeFrequency{value: val, isSet: true}
}

func (v NullableFeeFrequency) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeeFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
