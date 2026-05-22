package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PerformanceMetricForSB The advertising performance metric.
type PerformanceMetricForSB string

// List of PerformanceMetricForSB
const (
	PERFORMANCEMETRICFORSB_IS   PerformanceMetricForSB = "IS"
	PERFORMANCEMETRICFORSB_NTB  PerformanceMetricForSB = "NTB"
	PERFORMANCEMETRICFORSB_ROAS PerformanceMetricForSB = "ROAS"
)

// All allowed values of PerformanceMetricForSB enum
var AllowedPerformanceMetricForSBEnumValues = []PerformanceMetricForSB{
	"IS",
	"NTB",
	"ROAS",
}

func (v *PerformanceMetricForSB) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PerformanceMetricForSB(value)
	for _, existing := range AllowedPerformanceMetricForSBEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PerformanceMetricForSB", value)
}

// NewPerformanceMetricForSBFromValue returns a pointer to a valid PerformanceMetricForSB
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPerformanceMetricForSBFromValue(v string) (*PerformanceMetricForSB, error) {
	ev := PerformanceMetricForSB(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PerformanceMetricForSB: valid values are %v", v, AllowedPerformanceMetricForSBEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PerformanceMetricForSB) IsValid() bool {
	for _, existing := range AllowedPerformanceMetricForSBEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PerformanceMetricForSB value
func (v PerformanceMetricForSB) Ptr() *PerformanceMetricForSB {
	return &v
}

type NullablePerformanceMetricForSB struct {
	value *PerformanceMetricForSB
	isSet bool
}

func (v NullablePerformanceMetricForSB) Get() *PerformanceMetricForSB {
	return v.value
}

func (v *NullablePerformanceMetricForSB) Set(val *PerformanceMetricForSB) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceMetricForSB) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceMetricForSB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceMetricForSB(val *PerformanceMetricForSB) *NullablePerformanceMetricForSB {
	return &NullablePerformanceMetricForSB{value: val, isSet: true}
}

func (v NullablePerformanceMetricForSB) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePerformanceMetricForSB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
