package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// BudgetChangeType The value by which to update the budget of the budget rule.
type BudgetChangeType string

// List of BudgetChangeType
const (
	BUDGETCHANGETYPE_PERCENT BudgetChangeType = "PERCENT"
)

// All allowed values of BudgetChangeType enum
var AllowedBudgetChangeTypeEnumValues = []BudgetChangeType{
	"PERCENT",
}

func (v *BudgetChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BudgetChangeType(value)
	for _, existing := range AllowedBudgetChangeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BudgetChangeType", value)
}

// NewBudgetChangeTypeFromValue returns a pointer to a valid BudgetChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetChangeTypeFromValue(v string) (*BudgetChangeType, error) {
	ev := BudgetChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetChangeType: valid values are %v", v, AllowedBudgetChangeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetChangeType) IsValid() bool {
	for _, existing := range AllowedBudgetChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BudgetChangeType value
func (v BudgetChangeType) Ptr() *BudgetChangeType {
	return &v
}

type NullableBudgetChangeType struct {
	value *BudgetChangeType
	isSet bool
}

func (v NullableBudgetChangeType) Get() *BudgetChangeType {
	return v.value
}

func (v *NullableBudgetChangeType) Set(val *BudgetChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetChangeType(val *BudgetChangeType) *NullableBudgetChangeType {
	return &NullableBudgetChangeType{value: val, isSet: true}
}

func (v NullableBudgetChangeType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
