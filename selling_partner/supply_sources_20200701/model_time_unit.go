package supply_sources_20200701

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// TimeUnit The time unit
type TimeUnit string

// List of TimeUnit
const (
	TIMEUNIT_HOURS   TimeUnit = "Hours"
	TIMEUNIT_MINUTES TimeUnit = "Minutes"
	TIMEUNIT_DAYS    TimeUnit = "Days"
)

// All allowed values of TimeUnit enum
var AllowedTimeUnitEnumValues = []TimeUnit{
	TIMEUNIT_HOURS,
	TIMEUNIT_MINUTES,
	TIMEUNIT_DAYS,
}

func (v *TimeUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeUnit(value)
	for _, existing := range AllowedTimeUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeUnit", value)
}

// NewTimeUnitFromValue returns a pointer to a valid TimeUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeUnitFromValue(v string) (*TimeUnit, error) {
	ev := TimeUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeUnit: valid values are %v", v, AllowedTimeUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeUnit) IsValid() bool {
	for _, existing := range AllowedTimeUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeUnit value
func (v TimeUnit) Ptr() *TimeUnit {
	return &v
}

type NullableTimeUnit struct {
	value *TimeUnit
	isSet bool
}

func (v NullableTimeUnit) Get() *TimeUnit {
	return v.value
}

func (v *NullableTimeUnit) Set(val *TimeUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeUnit(val *TimeUnit) *NullableTimeUnit {
	return &NullableTimeUnit{value: val, isSet: true}
}

func (v NullableTimeUnit) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTimeUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
