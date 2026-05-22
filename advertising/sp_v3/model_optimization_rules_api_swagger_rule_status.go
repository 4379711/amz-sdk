package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// OptimizationRulesAPIRuleStatus The status of a rule. Only ENABLED and PAUSED are accepted in requests.
type OptimizationRulesAPIRuleStatus string

// List of OptimizationRulesAPIRuleStatus
const (
	OPTIMIZATIONRULESAPIRULESTATUS_SCHEDULED OptimizationRulesAPIRuleStatus = "SCHEDULED"
	OPTIMIZATIONRULESAPIRULESTATUS_ENABLED   OptimizationRulesAPIRuleStatus = "ENABLED"
	OPTIMIZATIONRULESAPIRULESTATUS_PAUSED    OptimizationRulesAPIRuleStatus = "PAUSED"
	OPTIMIZATIONRULESAPIRULESTATUS_ENDED     OptimizationRulesAPIRuleStatus = "ENDED"
)

// All allowed values of OptimizationRulesAPIRuleStatus enum
var AllowedOptimizationRulesAPIRuleStatusEnumValues = []OptimizationRulesAPIRuleStatus{
	"SCHEDULED",
	"ENABLED",
	"PAUSED",
	"ENDED",
}

func (v *OptimizationRulesAPIRuleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationRulesAPIRuleStatus(value)
	for _, existing := range AllowedOptimizationRulesAPIRuleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationRulesAPIRuleStatus", value)
}

// NewOptimizationRulesAPIRuleStatusFromValue returns a pointer to a valid OptimizationRulesAPIRuleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationRulesAPIRuleStatusFromValue(v string) (*OptimizationRulesAPIRuleStatus, error) {
	ev := OptimizationRulesAPIRuleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationRulesAPIRuleStatus: valid values are %v", v, AllowedOptimizationRulesAPIRuleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationRulesAPIRuleStatus) IsValid() bool {
	for _, existing := range AllowedOptimizationRulesAPIRuleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationRulesAPIRuleStatus value
func (v OptimizationRulesAPIRuleStatus) Ptr() *OptimizationRulesAPIRuleStatus {
	return &v
}

type NullableOptimizationRulesAPIRuleStatus struct {
	value *OptimizationRulesAPIRuleStatus
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleStatus) Get() *OptimizationRulesAPIRuleStatus {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleStatus) Set(val *OptimizationRulesAPIRuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleStatus(val *OptimizationRulesAPIRuleStatus) *NullableOptimizationRulesAPIRuleStatus {
	return &NullableOptimizationRulesAPIRuleStatus{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
