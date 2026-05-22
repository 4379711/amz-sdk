package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CostControlMetric Cost control metric to retrieve recommended value for.  Currently only COST_PER_CLICK is supported.
type CostControlMetric string

// List of CostControlMetric
const (
	COSTCONTROLMETRIC_COST_PER_CLICK CostControlMetric = "COST_PER_CLICK"
)

// All allowed values of CostControlMetric enum
var AllowedCostControlMetricEnumValues = []CostControlMetric{
	"COST_PER_CLICK",
}

func (v *CostControlMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CostControlMetric(value)
	for _, existing := range AllowedCostControlMetricEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CostControlMetric", value)
}

// NewCostControlMetricFromValue returns a pointer to a valid CostControlMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCostControlMetricFromValue(v string) (*CostControlMetric, error) {
	ev := CostControlMetric(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CostControlMetric: valid values are %v", v, AllowedCostControlMetricEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CostControlMetric) IsValid() bool {
	for _, existing := range AllowedCostControlMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostControlMetric value
func (v CostControlMetric) Ptr() *CostControlMetric {
	return &v
}

type NullableCostControlMetric struct {
	value *CostControlMetric
	isSet bool
}

func (v NullableCostControlMetric) Get() *CostControlMetric {
	return v.value
}

func (v *NullableCostControlMetric) Set(val *CostControlMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableCostControlMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableCostControlMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostControlMetric(val *CostControlMetric) *NullableCostControlMetric {
	return &NullableCostControlMetric{value: val, isSet: true}
}

func (v NullableCostControlMetric) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCostControlMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
