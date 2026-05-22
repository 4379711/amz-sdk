package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIRuleAttribute The attribute of the rule.
type OptimizationRulesAPIRuleAttribute string

// List of OptimizationRulesAPIRuleAttribute
const (
	OPTIMIZATIONRULESAPIRULEATTRIBUTE_ROAS OptimizationRulesAPIRuleAttribute = "ROAS"
)

// All allowed values of OptimizationRulesAPIRuleAttribute enum
var AllowedOptimizationRulesAPIRuleAttributeEnumValues = []OptimizationRulesAPIRuleAttribute{
	"ROAS",
}

func (v *OptimizationRulesAPIRuleAttribute) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIRuleAttribute(value)
	for _, existing := range AllowedOptimizationRulesAPIRuleAttributeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIRuleAttribute", value)
}

// NewOptimizationRulesAPIRuleAttributeFromValue returns a pointer to a valid OptimizationRulesAPIRuleAttribute
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIRuleAttributeFromValue(v string) (*OptimizationRulesAPIRuleAttribute, error) {
	ev := OptimizationRulesAPIRuleAttribute(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIRuleAttribute: valid values are %v", v, AllowedOptimizationRulesAPIRuleAttributeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIRuleAttribute) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIRuleAttributeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIRuleAttribute value
func (v OptimizationRulesAPIRuleAttribute) Ptr() *OptimizationRulesAPIRuleAttribute {
	return &v
}

type NullableOptimizationRulesAPIRuleAttribute struct {
	value *OptimizationRulesAPIRuleAttribute
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleAttribute) Get() *OptimizationRulesAPIRuleAttribute {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleAttribute) Set(val *OptimizationRulesAPIRuleAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleAttribute(val *OptimizationRulesAPIRuleAttribute) *NullableOptimizationRulesAPIRuleAttribute {
	return &NullableOptimizationRulesAPIRuleAttribute{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleAttribute) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
