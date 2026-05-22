package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ComparisonOperator The comparison operator.
type ComparisonOperator string

// List of ComparisonOperator
const (
	COMPARISONOPERATOR_GREATER_THAN             ComparisonOperator = "GREATER_THAN"
	COMPARISONOPERATOR_LESS_THAN                ComparisonOperator = "LESS_THAN"
	COMPARISONOPERATOR_LESS_THAN_OR_EQUAL_TO    ComparisonOperator = "LESS_THAN_OR_EQUAL_TO"
	COMPARISONOPERATOR_GREATER_THAN_OR_EQUAL_TO ComparisonOperator = "GREATER_THAN_OR_EQUAL_TO"
)

// All allowed values of ComparisonOperator enum
var AllowedComparisonOperatorEnumValues = []ComparisonOperator{
	"GREATER_THAN",
	"LESS_THAN",
	"LESS_THAN_OR_EQUAL_TO",
	"GREATER_THAN_OR_EQUAL_TO",
}

func (v *ComparisonOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComparisonOperator(value)
	for _, existing := range AllowedComparisonOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComparisonOperator", value)
}

// NewComparisonOperatorFromValue returns a pointer to a valid ComparisonOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComparisonOperatorFromValue(v string) (*ComparisonOperator, error) {
	ev := ComparisonOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComparisonOperator: valid values are %v", v, AllowedComparisonOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComparisonOperator) IsValid() bool {
	for _, existing := range AllowedComparisonOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComparisonOperator value
func (v ComparisonOperator) Ptr() *ComparisonOperator {
	return &v
}

type NullableComparisonOperator struct {
	value *ComparisonOperator
	isSet bool
}

func (v NullableComparisonOperator) Get() *ComparisonOperator {
	return v.value
}

func (v *NullableComparisonOperator) Set(val *ComparisonOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableComparisonOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableComparisonOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComparisonOperator(val *ComparisonOperator) *NullableComparisonOperator {
	return &NullableComparisonOperator{value: val, isSet: true}
}

func (v NullableComparisonOperator) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableComparisonOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
