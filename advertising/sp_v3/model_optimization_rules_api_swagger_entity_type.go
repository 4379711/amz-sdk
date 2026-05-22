package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIEntityType Types of advertising entities that may be associated to an optimization rule.
type OptimizationRulesAPIEntityType string

// List of OptimizationRulesAPIEntityType
const (
	OPTIMIZATIONRULESAPIENTITYTYPE_CAMPAIGN               OptimizationRulesAPIEntityType = "CAMPAIGN"
	OPTIMIZATIONRULESAPIENTITYTYPE_AD_GROUP               OptimizationRulesAPIEntityType = "AD_GROUP"
	OPTIMIZATIONRULESAPIENTITYTYPE_TARGET_PROMOTION_GROUP OptimizationRulesAPIEntityType = "TARGET_PROMOTION_GROUP"
)

// All allowed values of OptimizationRulesAPIEntityType enum
var AllowedOptimizationRulesAPIEntityTypeEnumValues = []OptimizationRulesAPIEntityType{
	"CAMPAIGN",
	"AD_GROUP",
	"TARGET_PROMOTION_GROUP",
}

func (v *OptimizationRulesAPIEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIEntityType(value)
	for _, existing := range AllowedOptimizationRulesAPIEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIEntityType", value)
}

// NewOptimizationRulesAPIEntityTypeFromValue returns a pointer to a valid OptimizationRulesAPIEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIEntityTypeFromValue(v string) (*OptimizationRulesAPIEntityType, error) {
	ev := OptimizationRulesAPIEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIEntityType: valid values are %v", v, AllowedOptimizationRulesAPIEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIEntityType) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIEntityType value
func (v OptimizationRulesAPIEntityType) Ptr() *OptimizationRulesAPIEntityType {
	return &v
}

type NullableOptimizationRulesAPIEntityType struct {
	value *OptimizationRulesAPIEntityType
	isSet bool
}

func (v NullableOptimizationRulesAPIEntityType) Get() *OptimizationRulesAPIEntityType {
	return v.value
}

func (v *NullableOptimizationRulesAPIEntityType) Set(val *OptimizationRulesAPIEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIEntityType(val *OptimizationRulesAPIEntityType) *NullableOptimizationRulesAPIEntityType {
	return &NullableOptimizationRulesAPIEntityType{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIEntityType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
