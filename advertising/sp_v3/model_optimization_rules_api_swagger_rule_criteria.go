package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIRuleCriteria - struct for OptimizationRulesAPIRuleCriteria
type OptimizationRulesAPIRuleCriteria struct {
	OptimizationRulesAPIRangeTypeRuleCriteria *OptimizationRulesAPIRangeTypeRuleCriteria
	OptimizationRulesAPIValueTypeRuleCriteria *OptimizationRulesAPIValueTypeRuleCriteria
}

// OptimizationRulesAPIRangeTypeRuleCriteriaAsOptimizationRulesAPIRuleCriteria is a convenience function that returns OptimizationRulesAPIRangeTypeRuleCriteria wrapped in OptimizationRulesAPIRuleCriteria
func OptimizationRulesAPIRangeTypeRuleCriteriaAsOptimizationRulesAPIRuleCriteria(v *OptimizationRulesAPIRangeTypeRuleCriteria) OptimizationRulesAPIRuleCriteria {
	return OptimizationRulesAPIRuleCriteria{
		OptimizationRulesAPIRangeTypeRuleCriteria: v,
	}
}

// OptimizationRulesAPIValueTypeRuleCriteriaAsOptimizationRulesAPIRuleCriteria is a convenience function that returns OptimizationRulesAPIValueTypeRuleCriteria wrapped in OptimizationRulesAPIRuleCriteria
func OptimizationRulesAPIValueTypeRuleCriteriaAsOptimizationRulesAPIRuleCriteria(v *OptimizationRulesAPIValueTypeRuleCriteria) OptimizationRulesAPIRuleCriteria {
	return OptimizationRulesAPIRuleCriteria{
		OptimizationRulesAPIValueTypeRuleCriteria: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OptimizationRulesAPIRuleCriteria) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OptimizationRulesAPIRangeTypeRuleCriteria
	err = sonic.Unmarshal(data, &dst.OptimizationRulesAPIRangeTypeRuleCriteria)
	if err == nil {
		jsonOptimizationRulesAPIRangeTypeRuleCriteria, _ := sonic.Marshal(dst.OptimizationRulesAPIRangeTypeRuleCriteria)
		if string(jsonOptimizationRulesAPIRangeTypeRuleCriteria) == "{}" { // empty struct
			dst.OptimizationRulesAPIRangeTypeRuleCriteria = nil
		} else {
			match++
		}
	} else {
		dst.OptimizationRulesAPIRangeTypeRuleCriteria = nil
	}

	// try to unmarshal data into OptimizationRulesAPIValueTypeRuleCriteria
	err = sonic.Unmarshal(data, &dst.OptimizationRulesAPIValueTypeRuleCriteria)
	if err == nil {
		jsonOptimizationRulesAPIValueTypeRuleCriteria, _ := sonic.Marshal(dst.OptimizationRulesAPIValueTypeRuleCriteria)
		if string(jsonOptimizationRulesAPIValueTypeRuleCriteria) == "{}" { // empty struct
			dst.OptimizationRulesAPIValueTypeRuleCriteria = nil
		} else {
			match++
		}
	} else {
		dst.OptimizationRulesAPIValueTypeRuleCriteria = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OptimizationRulesAPIRangeTypeRuleCriteria = nil
		dst.OptimizationRulesAPIValueTypeRuleCriteria = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OptimizationRulesAPIRuleCriteria)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OptimizationRulesAPIRuleCriteria)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OptimizationRulesAPIRuleCriteria) MarshalJSON() ([]byte, error) {
	if src.OptimizationRulesAPIRangeTypeRuleCriteria != nil {
		return sonic.Marshal(&src.OptimizationRulesAPIRangeTypeRuleCriteria)
	}

	if src.OptimizationRulesAPIValueTypeRuleCriteria != nil {
		return sonic.Marshal(&src.OptimizationRulesAPIValueTypeRuleCriteria)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OptimizationRulesAPIRuleCriteria) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OptimizationRulesAPIRangeTypeRuleCriteria != nil {
		return obj.OptimizationRulesAPIRangeTypeRuleCriteria
	}

	if obj.OptimizationRulesAPIValueTypeRuleCriteria != nil {
		return obj.OptimizationRulesAPIValueTypeRuleCriteria
	}

	// all schemas are nil
	return nil
}

type NullableOptimizationRulesAPIRuleCriteria struct {
	value *OptimizationRulesAPIRuleCriteria
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleCriteria) Get() *OptimizationRulesAPIRuleCriteria {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleCriteria) Set(val *OptimizationRulesAPIRuleCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleCriteria(val *OptimizationRulesAPIRuleCriteria) *NullableOptimizationRulesAPIRuleCriteria {
	return &NullableOptimizationRulesAPIRuleCriteria{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleCriteria) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
