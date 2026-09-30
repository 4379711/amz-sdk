package sp_v3

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OptimizationRulesAPIRuleCriteria) UnmarshalJSON(data []byte) error {
	dst.OptimizationRulesAPIRangeTypeRuleCriteria = nil
	dst.OptimizationRulesAPIValueTypeRuleCriteria = nil

	var fields map[string]json.RawMessage
	if err := sonic.Unmarshal(data, &fields); err != nil {
		return fmt.Errorf("rule criteria must be an object: %w", err)
	}
	if fields == nil {
		return fmt.Errorf("rule criteria must be a non-null object")
	}

	_, hasMin := fields["minValue"]
	_, hasMax := fields["maxValue"]
	_, hasOperator := fields["comparisonOperator"]
	_, hasValue := fields["value"]
	if (hasMin || hasMax) && (hasOperator || hasValue) {
		return fmt.Errorf("rule criteria cannot mix range and value fields")
	}

	// Field presence distinguishes the variants; zero is a valid numeric criterion.
	required := []string{"comparisonOperator", "value"}
	isRange := hasMin || hasMax
	if isRange {
		required = []string{"minValue", "maxValue"}
	}
	for _, name := range required {
		value, ok := fields[name]
		if !ok || bytes.Equal(bytes.TrimSpace(value), []byte("null")) {
			return fmt.Errorf("rule criteria requires non-null %s", name)
		}
	}

	if isRange {
		var criteria OptimizationRulesAPIRangeTypeRuleCriteria
		if err := sonic.Unmarshal(data, &criteria); err != nil {
			return err
		}
		dst.OptimizationRulesAPIRangeTypeRuleCriteria = &criteria
	} else {
		var criteria OptimizationRulesAPIValueTypeRuleCriteria
		if err := sonic.Unmarshal(data, &criteria); err != nil {
			return err
		}
		dst.OptimizationRulesAPIValueTypeRuleCriteria = &criteria
	}
	return nil
}
