package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BudgetType For the lifetime budget type, `startDate` and `endDate` must be specified.
type BudgetType string

// List of BudgetType
const (
	BUDGETTYPE_DAILY    BudgetType = "DAILY"
	BUDGETTYPE_LIFETIME BudgetType = "LIFETIME"
)

// All allowed values of BudgetType enum
var AllowedBudgetTypeEnumValues = []BudgetType{
	"DAILY",
	"LIFETIME",
}

func (v *BudgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BudgetType(value)
	for _, existing := range AllowedBudgetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BudgetType", value)
}

// NewBudgetTypeFromValue returns a pointer to a valid BudgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetTypeFromValue(v string) (*BudgetType, error) {
	ev := BudgetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetType: valid values are %v", v, AllowedBudgetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetType) IsValid() bool {
	for _, existing := range AllowedBudgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BudgetType value
func (v BudgetType) Ptr() *BudgetType {
	return &v
}

type NullableBudgetType struct {
	value *BudgetType
	isSet bool
}

func (v NullableBudgetType) Get() *BudgetType {
	return v.value
}

func (v *NullableBudgetType) Set(val *BudgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetType(val *BudgetType) *NullableBudgetType {
	return &NullableBudgetType{value: val, isSet: true}
}

func (v NullableBudgetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
