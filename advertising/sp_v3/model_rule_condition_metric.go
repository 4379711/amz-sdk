package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// RuleConditionMetric The advertising performance metric. ROAS is the only supported metric.
type RuleConditionMetric string

// List of RuleConditionMetric
const (
	RULECONDITIONMETRIC_ROAS        RuleConditionMetric = "ROAS"
	RULECONDITIONMETRIC_AVERAGE_BID RuleConditionMetric = "AVERAGE_BID"
)

// All allowed values of RuleConditionMetric enum
var AllowedRuleConditionMetricEnumValues = []RuleConditionMetric{
	"ROAS",
	"AVERAGE_BID",
}

func (v *RuleConditionMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuleConditionMetric(value)
	for _, existing := range AllowedRuleConditionMetricEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuleConditionMetric", value)
}

// NewRuleConditionMetricFromValue returns a pointer to a valid RuleConditionMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuleConditionMetricFromValue(v string) (*RuleConditionMetric, error) {
	ev := RuleConditionMetric(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuleConditionMetric: valid values are %v", v, AllowedRuleConditionMetricEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuleConditionMetric) IsValid() bool {
	for _, existing := range AllowedRuleConditionMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleConditionMetric value
func (v RuleConditionMetric) Ptr() *RuleConditionMetric {
	return &v
}

type NullableRuleConditionMetric struct {
	value *RuleConditionMetric
	isSet bool
}

func (v NullableRuleConditionMetric) Get() *RuleConditionMetric {
	return v.value
}

func (v *NullableRuleConditionMetric) Set(val *RuleConditionMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleConditionMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleConditionMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleConditionMetric(val *RuleConditionMetric) *NullableRuleConditionMetric {
	return &NullableRuleConditionMetric{value: val, isSet: true}
}

func (v NullableRuleConditionMetric) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleConditionMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
