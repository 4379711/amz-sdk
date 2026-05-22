package replenishment20221107

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AggregationFrequency The time period used to group data in the response. Note that this is only valid for the `PERFORMANCE` time period type.
type AggregationFrequency string

// List of AggregationFrequency
const (
	AGGREGATIONFREQUENCY_WEEK    AggregationFrequency = "WEEK"
	AGGREGATIONFREQUENCY_MONTH   AggregationFrequency = "MONTH"
	AGGREGATIONFREQUENCY_QUARTER AggregationFrequency = "QUARTER"
	AGGREGATIONFREQUENCY_YEAR    AggregationFrequency = "YEAR"
)

// All allowed values of AggregationFrequency enum
var AllowedAggregationFrequencyEnumValues = []AggregationFrequency{
	AGGREGATIONFREQUENCY_WEEK,
	AGGREGATIONFREQUENCY_MONTH,
	AGGREGATIONFREQUENCY_QUARTER,
	AGGREGATIONFREQUENCY_YEAR,
}

func (v *AggregationFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AggregationFrequency(value)
	for _, existing := range AllowedAggregationFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AggregationFrequency", value)
}

// NewAggregationFrequencyFromValue returns a pointer to a valid AggregationFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAggregationFrequencyFromValue(v string) (*AggregationFrequency, error) {
	ev := AggregationFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AggregationFrequency: valid values are %v", v, AllowedAggregationFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AggregationFrequency) IsValid() bool {
	for _, existing := range AllowedAggregationFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregationFrequency value
func (v AggregationFrequency) Ptr() *AggregationFrequency {
	return &v
}

type NullableAggregationFrequency struct {
	value *AggregationFrequency
	isSet bool
}

func (v NullableAggregationFrequency) Get() *AggregationFrequency {
	return v.value
}

func (v *NullableAggregationFrequency) Set(val *AggregationFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationFrequency(val *AggregationFrequency) *NullableAggregationFrequency {
	return &NullableAggregationFrequency{value: val, isSet: true}
}

func (v NullableAggregationFrequency) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAggregationFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
