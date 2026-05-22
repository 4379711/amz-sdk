package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIRuleActionOperator The action operation for the rule.
type OptimizationRulesAPIRuleActionOperator string

// List of OptimizationRulesAPIRuleActionOperator
const (
	OPTIMIZATIONRULESAPIRULEACTIONOPERATOR_INCREMENT OptimizationRulesAPIRuleActionOperator = "INCREMENT"
)

// All allowed values of OptimizationRulesAPIRuleActionOperator enum
var AllowedOptimizationRulesAPIRuleActionOperatorEnumValues = []OptimizationRulesAPIRuleActionOperator{
	"INCREMENT",
}

func (v *OptimizationRulesAPIRuleActionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIRuleActionOperator(value)
	for _, existing := range AllowedOptimizationRulesAPIRuleActionOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIRuleActionOperator", value)
}

// NewOptimizationRulesAPIRuleActionOperatorFromValue returns a pointer to a valid OptimizationRulesAPIRuleActionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIRuleActionOperatorFromValue(v string) (*OptimizationRulesAPIRuleActionOperator, error) {
	ev := OptimizationRulesAPIRuleActionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIRuleActionOperator: valid values are %v", v, AllowedOptimizationRulesAPIRuleActionOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIRuleActionOperator) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIRuleActionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIRuleActionOperator value
func (v OptimizationRulesAPIRuleActionOperator) Ptr() *OptimizationRulesAPIRuleActionOperator {
	return &v
}

type NullableOptimizationRulesAPIRuleActionOperator struct {
	value *OptimizationRulesAPIRuleActionOperator
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleActionOperator) Get() *OptimizationRulesAPIRuleActionOperator {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleActionOperator) Set(val *OptimizationRulesAPIRuleActionOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleActionOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleActionOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleActionOperator(val *OptimizationRulesAPIRuleActionOperator) *NullableOptimizationRulesAPIRuleActionOperator {
	return &NullableOptimizationRulesAPIRuleActionOperator{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleActionOperator) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleActionOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
