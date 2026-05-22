package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIComparisonOperator The comparison operator.
type OptimizationRulesAPIComparisonOperator string

// List of OptimizationRulesAPIComparisonOperator
const (
	OPTIMIZATIONRULESAPICOMPARISONOPERATOR_GREATER_THAN             OptimizationRulesAPIComparisonOperator = "GREATER_THAN"
	OPTIMIZATIONRULESAPICOMPARISONOPERATOR_LESS_THAN                OptimizationRulesAPIComparisonOperator = "LESS_THAN"
	OPTIMIZATIONRULESAPICOMPARISONOPERATOR_EQUAL_TO                 OptimizationRulesAPIComparisonOperator = "EQUAL_TO"
	OPTIMIZATIONRULESAPICOMPARISONOPERATOR_LESS_THAN_OR_EQUAL_TO    OptimizationRulesAPIComparisonOperator = "LESS_THAN_OR_EQUAL_TO"
	OPTIMIZATIONRULESAPICOMPARISONOPERATOR_GREATER_THAN_OR_EQUAL_TO OptimizationRulesAPIComparisonOperator = "GREATER_THAN_OR_EQUAL_TO"
)

// All allowed values of OptimizationRulesAPIComparisonOperator enum
var AllowedOptimizationRulesAPIComparisonOperatorEnumValues = []OptimizationRulesAPIComparisonOperator{
	"GREATER_THAN",
	"LESS_THAN",
	"EQUAL_TO",
	"LESS_THAN_OR_EQUAL_TO",
	"GREATER_THAN_OR_EQUAL_TO",
}

func (v *OptimizationRulesAPIComparisonOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIComparisonOperator(value)
	for _, existing := range AllowedOptimizationRulesAPIComparisonOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIComparisonOperator", value)
}

// NewOptimizationRulesAPIComparisonOperatorFromValue returns a pointer to a valid OptimizationRulesAPIComparisonOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIComparisonOperatorFromValue(v string) (*OptimizationRulesAPIComparisonOperator, error) {
	ev := OptimizationRulesAPIComparisonOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIComparisonOperator: valid values are %v", v, AllowedOptimizationRulesAPIComparisonOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIComparisonOperator) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIComparisonOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIComparisonOperator value
func (v OptimizationRulesAPIComparisonOperator) Ptr() *OptimizationRulesAPIComparisonOperator {
	return &v
}

type NullableOptimizationRulesAPIComparisonOperator struct {
	value *OptimizationRulesAPIComparisonOperator
	isSet bool
}

func (v NullableOptimizationRulesAPIComparisonOperator) Get() *OptimizationRulesAPIComparisonOperator {
	return v.value
}

func (v *NullableOptimizationRulesAPIComparisonOperator) Set(val *OptimizationRulesAPIComparisonOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIComparisonOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIComparisonOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIComparisonOperator(val *OptimizationRulesAPIComparisonOperator) *NullableOptimizationRulesAPIComparisonOperator {
	return &NullableOptimizationRulesAPIComparisonOperator{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIComparisonOperator) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIComparisonOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
