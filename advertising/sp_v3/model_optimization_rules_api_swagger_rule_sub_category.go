package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIRuleSubCategory The sub-category of the optimization rule.
type OptimizationRulesAPIRuleSubCategory string

// List of OptimizationRulesAPIRuleSubCategory
const (
	OPTIMIZATIONRULESAPIRULESUBCATEGORY_SCHEDULE OptimizationRulesAPIRuleSubCategory = "SCHEDULE"
)

// All allowed values of OptimizationRulesAPIRuleSubCategory enum
var AllowedOptimizationRulesAPIRuleSubCategoryEnumValues = []OptimizationRulesAPIRuleSubCategory{
	"SCHEDULE",
}

func (v *OptimizationRulesAPIRuleSubCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIRuleSubCategory(value)
	for _, existing := range AllowedOptimizationRulesAPIRuleSubCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIRuleSubCategory", value)
}

// NewOptimizationRulesAPIRuleSubCategoryFromValue returns a pointer to a valid OptimizationRulesAPIRuleSubCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIRuleSubCategoryFromValue(v string) (*OptimizationRulesAPIRuleSubCategory, error) {
	ev := OptimizationRulesAPIRuleSubCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIRuleSubCategory: valid values are %v", v, AllowedOptimizationRulesAPIRuleSubCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIRuleSubCategory) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIRuleSubCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIRuleSubCategory value
func (v OptimizationRulesAPIRuleSubCategory) Ptr() *OptimizationRulesAPIRuleSubCategory {
	return &v
}

type NullableOptimizationRulesAPIRuleSubCategory struct {
	value *OptimizationRulesAPIRuleSubCategory
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleSubCategory) Get() *OptimizationRulesAPIRuleSubCategory {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleSubCategory) Set(val *OptimizationRulesAPIRuleSubCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleSubCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleSubCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleSubCategory(val *OptimizationRulesAPIRuleSubCategory) *NullableOptimizationRulesAPIRuleSubCategory {
	return &NullableOptimizationRulesAPIRuleSubCategory{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleSubCategory) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleSubCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
