package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIDayOfTheWeek Day of the week.
type OptimizationRulesAPIDayOfTheWeek string

// List of OptimizationRulesAPIDayOfTheWeek
const (
	OPTIMIZATIONRULESAPIDAYOFTHEWEEK_MONDAY    OptimizationRulesAPIDayOfTheWeek = "MONDAY"
	OPTIMIZATIONRULESAPIDAYOFTHEWEEK_TUESDAY   OptimizationRulesAPIDayOfTheWeek = "TUESDAY"
	OPTIMIZATIONRULESAPIDAYOFTHEWEEK_WEDNESDAY OptimizationRulesAPIDayOfTheWeek = "WEDNESDAY"
	OPTIMIZATIONRULESAPIDAYOFTHEWEEK_THURSDAY  OptimizationRulesAPIDayOfTheWeek = "THURSDAY"
	OPTIMIZATIONRULESAPIDAYOFTHEWEEK_FRIDAY    OptimizationRulesAPIDayOfTheWeek = "FRIDAY"
	OPTIMIZATIONRULESAPIDAYOFTHEWEEK_SATURDAY  OptimizationRulesAPIDayOfTheWeek = "SATURDAY"
	OPTIMIZATIONRULESAPIDAYOFTHEWEEK_SUNDAY    OptimizationRulesAPIDayOfTheWeek = "SUNDAY"
)

// All allowed values of OptimizationRulesAPIDayOfTheWeek enum
var AllowedOptimizationRulesAPIDayOfTheWeekEnumValues = []OptimizationRulesAPIDayOfTheWeek{
	"MONDAY",
	"TUESDAY",
	"WEDNESDAY",
	"THURSDAY",
	"FRIDAY",
	"SATURDAY",
	"SUNDAY",
}

func (v *OptimizationRulesAPIDayOfTheWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIDayOfTheWeek(value)
	for _, existing := range AllowedOptimizationRulesAPIDayOfTheWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIDayOfTheWeek", value)
}

// NewOptimizationRulesAPIDayOfTheWeekFromValue returns a pointer to a valid OptimizationRulesAPIDayOfTheWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIDayOfTheWeekFromValue(v string) (*OptimizationRulesAPIDayOfTheWeek, error) {
	ev := OptimizationRulesAPIDayOfTheWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIDayOfTheWeek: valid values are %v", v, AllowedOptimizationRulesAPIDayOfTheWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIDayOfTheWeek) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIDayOfTheWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIDayOfTheWeek value
func (v OptimizationRulesAPIDayOfTheWeek) Ptr() *OptimizationRulesAPIDayOfTheWeek {
	return &v
}

type NullableOptimizationRulesAPIDayOfTheWeek struct {
	value *OptimizationRulesAPIDayOfTheWeek
	isSet bool
}

func (v NullableOptimizationRulesAPIDayOfTheWeek) Get() *OptimizationRulesAPIDayOfTheWeek {
	return v.value
}

func (v *NullableOptimizationRulesAPIDayOfTheWeek) Set(val *OptimizationRulesAPIDayOfTheWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDayOfTheWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDayOfTheWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDayOfTheWeek(val *OptimizationRulesAPIDayOfTheWeek) *NullableOptimizationRulesAPIDayOfTheWeek {
	return &NullableOptimizationRulesAPIDayOfTheWeek{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDayOfTheWeek) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDayOfTheWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
