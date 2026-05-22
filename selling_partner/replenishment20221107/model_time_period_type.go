package replenishment20221107

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// TimePeriodType The time period type that determines whether the metrics requested are backward-looking (performance) or forward-looking (forecast).
type TimePeriodType string

// List of TimePeriodType
const (
	TIMEPERIODTYPE_PERFORMANCE TimePeriodType = "PERFORMANCE"
	TIMEPERIODTYPE_FORECAST    TimePeriodType = "FORECAST"
)

// All allowed values of TimePeriodType enum
var AllowedTimePeriodTypeEnumValues = []TimePeriodType{
	TIMEPERIODTYPE_PERFORMANCE,
	TIMEPERIODTYPE_FORECAST,
}

func (v *TimePeriodType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimePeriodType(value)
	for _, existing := range AllowedTimePeriodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimePeriodType", value)
}

// NewTimePeriodTypeFromValue returns a pointer to a valid TimePeriodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimePeriodTypeFromValue(v string) (*TimePeriodType, error) {
	ev := TimePeriodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimePeriodType: valid values are %v", v, AllowedTimePeriodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimePeriodType) IsValid() bool {
	for _, existing := range AllowedTimePeriodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimePeriodType value
func (v TimePeriodType) Ptr() *TimePeriodType {
	return &v
}

type NullableTimePeriodType struct {
	value *TimePeriodType
	isSet bool
}

func (v NullableTimePeriodType) Get() *TimePeriodType {
	return v.value
}

func (v *NullableTimePeriodType) Set(val *TimePeriodType) {
	v.value = val
	v.isSet = true
}

func (v NullableTimePeriodType) IsSet() bool {
	return v.isSet
}

func (v *NullableTimePeriodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimePeriodType(val *TimePeriodType) *NullableTimePeriodType {
	return &NullableTimePeriodType{value: val, isSet: true}
}

func (v NullableTimePeriodType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTimePeriodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
