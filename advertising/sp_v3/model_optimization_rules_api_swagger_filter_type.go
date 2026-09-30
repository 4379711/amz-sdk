package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIFilterType Types of filter used for search.
type OptimizationRulesAPIFilterType string

// List of OptimizationRulesAPIFilterType
const (
	OPTIMIZATIONRULESAPIFILTERTYPE_EXACT_MATCH OptimizationRulesAPIFilterType = "EXACT_MATCH"
	OPTIMIZATIONRULESAPIFILTERTYPE_BROAD_MATCH OptimizationRulesAPIFilterType = "BROAD_MATCH"
)

// All allowed values of OptimizationRulesAPIFilterType enum
var AllowedOptimizationRulesAPIFilterTypeEnumValues = []OptimizationRulesAPIFilterType{
	"EXACT_MATCH",
	"BROAD_MATCH",
}

func (v *OptimizationRulesAPIFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = OptimizationRulesAPIFilterType(value)
	return nil
}

// NewOptimizationRulesAPIFilterTypeFromValue returns a pointer to a valid OptimizationRulesAPIFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIFilterTypeFromValue(v string) (*OptimizationRulesAPIFilterType, error) {
	ev := OptimizationRulesAPIFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIFilterType: valid values are %v", v, AllowedOptimizationRulesAPIFilterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIFilterType) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIFilterType value
func (v OptimizationRulesAPIFilterType) Ptr() *OptimizationRulesAPIFilterType {
	return &v
}

type NullableOptimizationRulesAPIFilterType struct {
	value *OptimizationRulesAPIFilterType
	isSet bool
}

func (v NullableOptimizationRulesAPIFilterType) Get() *OptimizationRulesAPIFilterType {
	return v.value
}

func (v *NullableOptimizationRulesAPIFilterType) Set(val *OptimizationRulesAPIFilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIFilterType(val *OptimizationRulesAPIFilterType) *NullableOptimizationRulesAPIFilterType {
	return &NullableOptimizationRulesAPIFilterType{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIFilterType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
