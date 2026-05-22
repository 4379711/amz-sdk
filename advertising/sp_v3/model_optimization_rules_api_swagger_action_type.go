package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIActionType The action taken when the optimization rule is enabled. Defaults to ADOPT.
type OptimizationRulesAPIActionType string

// List of OptimizationRulesAPIActionType
const (
	OPTIMIZATIONRULESAPIACTIONTYPE_ADOPT OptimizationRulesAPIActionType = "ADOPT"
)

// All allowed values of OptimizationRulesAPIActionType enum
var AllowedOptimizationRulesAPIActionTypeEnumValues = []OptimizationRulesAPIActionType{
	"ADOPT",
}

func (v *OptimizationRulesAPIActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIActionType(value)
	for _, existing := range AllowedOptimizationRulesAPIActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIActionType", value)
}

// NewOptimizationRulesAPIActionTypeFromValue returns a pointer to a valid OptimizationRulesAPIActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIActionTypeFromValue(v string) (*OptimizationRulesAPIActionType, error) {
	ev := OptimizationRulesAPIActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIActionType: valid values are %v", v, AllowedOptimizationRulesAPIActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIActionType) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIActionType value
func (v OptimizationRulesAPIActionType) Ptr() *OptimizationRulesAPIActionType {
	return &v
}

type NullableOptimizationRulesAPIActionType struct {
	value *OptimizationRulesAPIActionType
	isSet bool
}

func (v NullableOptimizationRulesAPIActionType) Get() *OptimizationRulesAPIActionType {
	return v.value
}

func (v *NullableOptimizationRulesAPIActionType) Set(val *OptimizationRulesAPIActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIActionType(val *OptimizationRulesAPIActionType) *NullableOptimizationRulesAPIActionType {
	return &NullableOptimizationRulesAPIActionType{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIActionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
