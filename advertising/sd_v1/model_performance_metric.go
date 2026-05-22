package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// PerformanceMetric The advertising performance metric.
type PerformanceMetric string

// List of PerformanceMetric
const (
	PERFORMANCEMETRIC_ACOS PerformanceMetric = "ACOS"
	PERFORMANCEMETRIC_CTR  PerformanceMetric = "CTR"
	PERFORMANCEMETRIC_CVR  PerformanceMetric = "CVR"
	PERFORMANCEMETRIC_ROAS PerformanceMetric = "ROAS"
)

// All allowed values of PerformanceMetric enum
var AllowedPerformanceMetricEnumValues = []PerformanceMetric{
	"ACOS",
	"CTR",
	"CVR",
	"ROAS",
}

func (v *PerformanceMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PerformanceMetric(value)
	for _, existing := range AllowedPerformanceMetricEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PerformanceMetric", value)
}

// NewPerformanceMetricFromValue returns a pointer to a valid PerformanceMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPerformanceMetricFromValue(v string) (*PerformanceMetric, error) {
	ev := PerformanceMetric(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PerformanceMetric: valid values are %v", v, AllowedPerformanceMetricEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PerformanceMetric) IsValid() bool {
	for _, existing := range AllowedPerformanceMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PerformanceMetric value
func (v PerformanceMetric) Ptr() *PerformanceMetric {
	return &v
}

type NullablePerformanceMetric struct {
	value *PerformanceMetric
	isSet bool
}

func (v NullablePerformanceMetric) Get() *PerformanceMetric {
	return v.value
}

func (v *NullablePerformanceMetric) Set(val *PerformanceMetric) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceMetric) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceMetric(val *PerformanceMetric) *NullablePerformanceMetric {
	return &NullablePerformanceMetric{value: val, isSet: true}
}

func (v NullablePerformanceMetric) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePerformanceMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
