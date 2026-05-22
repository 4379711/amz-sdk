package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIRuleRecurrenceType The frequency of the optimization rule application.
type OptimizationRulesAPIRuleRecurrenceType string

// List of OptimizationRulesAPIRuleRecurrenceType
const (
	OPTIMIZATIONRULESAPIRULERECURRENCETYPE_DAILY  OptimizationRulesAPIRuleRecurrenceType = "DAILY"
	OPTIMIZATIONRULESAPIRULERECURRENCETYPE_WEEKLY OptimizationRulesAPIRuleRecurrenceType = "WEEKLY"
)

// All allowed values of OptimizationRulesAPIRuleRecurrenceType enum
var AllowedOptimizationRulesAPIRuleRecurrenceTypeEnumValues = []OptimizationRulesAPIRuleRecurrenceType{
	"DAILY",
	"WEEKLY",
}

func (v *OptimizationRulesAPIRuleRecurrenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIRuleRecurrenceType(value)
	for _, existing := range AllowedOptimizationRulesAPIRuleRecurrenceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIRuleRecurrenceType", value)
}

// NewOptimizationRulesAPIRuleRecurrenceTypeFromValue returns a pointer to a valid OptimizationRulesAPIRuleRecurrenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIRuleRecurrenceTypeFromValue(v string) (*OptimizationRulesAPIRuleRecurrenceType, error) {
	ev := OptimizationRulesAPIRuleRecurrenceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIRuleRecurrenceType: valid values are %v", v, AllowedOptimizationRulesAPIRuleRecurrenceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIRuleRecurrenceType) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIRuleRecurrenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIRuleRecurrenceType value
func (v OptimizationRulesAPIRuleRecurrenceType) Ptr() *OptimizationRulesAPIRuleRecurrenceType {
	return &v
}

type NullableOptimizationRulesAPIRuleRecurrenceType struct {
	value *OptimizationRulesAPIRuleRecurrenceType
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleRecurrenceType) Get() *OptimizationRulesAPIRuleRecurrenceType {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleRecurrenceType) Set(val *OptimizationRulesAPIRuleRecurrenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleRecurrenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleRecurrenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleRecurrenceType(val *OptimizationRulesAPIRuleRecurrenceType) *NullableOptimizationRulesAPIRuleRecurrenceType {
	return &NullableOptimizationRulesAPIRuleRecurrenceType{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleRecurrenceType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleRecurrenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
