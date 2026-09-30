package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// RecurrenceType depicts the type of recurrence
type RecurrenceType string

// List of RecurrenceType
const (
	RECURRENCETYPE_DAILY  RecurrenceType = "DAILY"
	RECURRENCETYPE_WEEKLY RecurrenceType = "WEEKLY"
)

// All allowed values of RecurrenceType enum
var AllowedRecurrenceTypeEnumValues = []RecurrenceType{
	"DAILY",
	"WEEKLY",
}

func (v *RecurrenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = RecurrenceType(value)
	return nil
}

// NewRecurrenceTypeFromValue returns a pointer to a valid RecurrenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecurrenceTypeFromValue(v string) (*RecurrenceType, error) {
	ev := RecurrenceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecurrenceType: valid values are %v", v, AllowedRecurrenceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecurrenceType) IsValid() bool {
	for _, existing := range AllowedRecurrenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecurrenceType value
func (v RecurrenceType) Ptr() *RecurrenceType {
	return &v
}

type NullableRecurrenceType struct {
	value *RecurrenceType
	isSet bool
}

func (v NullableRecurrenceType) Get() *RecurrenceType {
	return v.value
}

func (v *NullableRecurrenceType) Set(val *RecurrenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrenceType(val *RecurrenceType) *NullableRecurrenceType {
	return &NullableRecurrenceType{value: val, isSet: true}
}

func (v NullableRecurrenceType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecurrenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
