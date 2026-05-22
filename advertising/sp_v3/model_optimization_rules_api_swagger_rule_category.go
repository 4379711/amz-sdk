package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIRuleCategory The type of the optimization rule.
type OptimizationRulesAPIRuleCategory string

// List of OptimizationRulesAPIRuleCategory
const (
	OPTIMIZATIONRULESAPIRULECATEGORY_BID OptimizationRulesAPIRuleCategory = "BID"
)

// All allowed values of OptimizationRulesAPIRuleCategory enum
var AllowedOptimizationRulesAPIRuleCategoryEnumValues = []OptimizationRulesAPIRuleCategory{
	"BID",
}

func (v *OptimizationRulesAPIRuleCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIRuleCategory(value)
	for _, existing := range AllowedOptimizationRulesAPIRuleCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIRuleCategory", value)
}

// NewOptimizationRulesAPIRuleCategoryFromValue returns a pointer to a valid OptimizationRulesAPIRuleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIRuleCategoryFromValue(v string) (*OptimizationRulesAPIRuleCategory, error) {
	ev := OptimizationRulesAPIRuleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIRuleCategory: valid values are %v", v, AllowedOptimizationRulesAPIRuleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIRuleCategory) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIRuleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIRuleCategory value
func (v OptimizationRulesAPIRuleCategory) Ptr() *OptimizationRulesAPIRuleCategory {
	return &v
}

type NullableOptimizationRulesAPIRuleCategory struct {
	value *OptimizationRulesAPIRuleCategory
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleCategory) Get() *OptimizationRulesAPIRuleCategory {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleCategory) Set(val *OptimizationRulesAPIRuleCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleCategory(val *OptimizationRulesAPIRuleCategory) *NullableOptimizationRulesAPIRuleCategory {
	return &NullableOptimizationRulesAPIRuleCategory{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleCategory) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
