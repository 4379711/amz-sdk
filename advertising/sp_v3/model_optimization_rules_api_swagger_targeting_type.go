package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OptimizationRulesAPITargetingType The type of targets for the rule.
type OptimizationRulesAPITargetingType string

// List of OptimizationRulesAPITargetingType
const (
	OPTIMIZATIONRULESAPITARGETINGTYPE_PRODUCT OptimizationRulesAPITargetingType = "PRODUCT"
	OPTIMIZATIONRULESAPITARGETINGTYPE_KEYWORD OptimizationRulesAPITargetingType = "KEYWORD"
)

// All allowed values of OptimizationRulesAPITargetingType enum
var AllowedOptimizationRulesAPITargetingTypeEnumValues = []OptimizationRulesAPITargetingType{
	"PRODUCT",
	"KEYWORD",
}

func (v *OptimizationRulesAPITargetingType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = OptimizationRulesAPITargetingType(value)
	return nil
}

// NewOptimizationRulesAPITargetingTypeFromValue returns a pointer to a valid OptimizationRulesAPITargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPITargetingTypeFromValue(v string) (*OptimizationRulesAPITargetingType, error) {
	ev := OptimizationRulesAPITargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPITargetingType: valid values are %v", v, AllowedOptimizationRulesAPITargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPITargetingType) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPITargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPITargetingType value
func (v OptimizationRulesAPITargetingType) Ptr() *OptimizationRulesAPITargetingType {
	return &v
}

type NullableOptimizationRulesAPITargetingType struct {
	value *OptimizationRulesAPITargetingType
	isSet bool
}

func (v NullableOptimizationRulesAPITargetingType) Get() *OptimizationRulesAPITargetingType {
	return v.value
}

func (v *NullableOptimizationRulesAPITargetingType) Set(val *OptimizationRulesAPITargetingType) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPITargetingType) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPITargetingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPITargetingType(val *OptimizationRulesAPITargetingType) *NullableOptimizationRulesAPITargetingType {
	return &NullableOptimizationRulesAPITargetingType{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPITargetingType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPITargetingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
