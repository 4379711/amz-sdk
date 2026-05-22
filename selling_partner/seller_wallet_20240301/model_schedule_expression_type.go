package seller_wallet_20240301

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ScheduleExpressionType The type of scheduled transfer expression.
type ScheduleExpressionType string

// List of ScheduleExpressionType
const (
	SCHEDULEEXPRESSIONTYPE_RECURRING ScheduleExpressionType = "RECURRING"
	SCHEDULEEXPRESSIONTYPE_ONE_TIME  ScheduleExpressionType = "ONE_TIME"
)

// All allowed values of ScheduleExpressionType enum
var AllowedScheduleExpressionTypeEnumValues = []ScheduleExpressionType{
	SCHEDULEEXPRESSIONTYPE_RECURRING,
	SCHEDULEEXPRESSIONTYPE_ONE_TIME,
}

func (v *ScheduleExpressionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduleExpressionType(value)
	for _, existing := range AllowedScheduleExpressionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleExpressionType", value)
}

// NewScheduleExpressionTypeFromValue returns a pointer to a valid ScheduleExpressionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduleExpressionTypeFromValue(v string) (*ScheduleExpressionType, error) {
	ev := ScheduleExpressionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduleExpressionType: valid values are %v", v, AllowedScheduleExpressionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduleExpressionType) IsValid() bool {
	for _, existing := range AllowedScheduleExpressionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleExpressionType value
func (v ScheduleExpressionType) Ptr() *ScheduleExpressionType {
	return &v
}

type NullableScheduleExpressionType struct {
	value *ScheduleExpressionType
	isSet bool
}

func (v NullableScheduleExpressionType) Get() *ScheduleExpressionType {
	return v.value
}

func (v *NullableScheduleExpressionType) Set(val *ScheduleExpressionType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleExpressionType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleExpressionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleExpressionType(val *ScheduleExpressionType) *NullableScheduleExpressionType {
	return &NullableScheduleExpressionType{value: val, isSet: true}
}

func (v NullableScheduleExpressionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScheduleExpressionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
