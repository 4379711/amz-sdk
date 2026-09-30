package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIExpressionType The expression types of targets for the rule.
type OptimizationRulesAPIExpressionType string

// List of OptimizationRulesAPIExpressionType
const (
	OPTIMIZATIONRULESAPIEXPRESSIONTYPE_EXACT    OptimizationRulesAPIExpressionType = "EXACT"
	OPTIMIZATIONRULESAPIEXPRESSIONTYPE_EXTENDED OptimizationRulesAPIExpressionType = "EXTENDED"
	OPTIMIZATIONRULESAPIEXPRESSIONTYPE_BROAD    OptimizationRulesAPIExpressionType = "BROAD"
	OPTIMIZATIONRULESAPIEXPRESSIONTYPE_PHRASE   OptimizationRulesAPIExpressionType = "PHRASE"
	OPTIMIZATIONRULESAPIEXPRESSIONTYPE_EXPANDED OptimizationRulesAPIExpressionType = "EXPANDED"
)

// All allowed values of OptimizationRulesAPIExpressionType enum
var AllowedOptimizationRulesAPIExpressionTypeEnumValues = []OptimizationRulesAPIExpressionType{
	"EXACT",
	"EXTENDED",
	"BROAD",
	"PHRASE",
	"EXPANDED",
}

func (v *OptimizationRulesAPIExpressionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = OptimizationRulesAPIExpressionType(value)
	return nil
}

// NewOptimizationRulesAPIExpressionTypeFromValue returns a pointer to a valid OptimizationRulesAPIExpressionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIExpressionTypeFromValue(v string) (*OptimizationRulesAPIExpressionType, error) {
	ev := OptimizationRulesAPIExpressionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIExpressionType: valid values are %v", v, AllowedOptimizationRulesAPIExpressionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIExpressionType) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIExpressionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIExpressionType value
func (v OptimizationRulesAPIExpressionType) Ptr() *OptimizationRulesAPIExpressionType {
	return &v
}

type NullableOptimizationRulesAPIExpressionType struct {
	value *OptimizationRulesAPIExpressionType
	isSet bool
}

func (v NullableOptimizationRulesAPIExpressionType) Get() *OptimizationRulesAPIExpressionType {
	return v.value
}

func (v *NullableOptimizationRulesAPIExpressionType) Set(val *OptimizationRulesAPIExpressionType) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIExpressionType) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIExpressionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIExpressionType(val *OptimizationRulesAPIExpressionType) *NullableOptimizationRulesAPIExpressionType {
	return &NullableOptimizationRulesAPIExpressionType{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIExpressionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIExpressionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
