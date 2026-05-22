package seller_wallet_20240301

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ScheduleTransferType The type of schedule the transfer is on. Schedules based on time patterns use EventBridge.
type ScheduleTransferType string

// List of ScheduleTransferType
const (
	SCHEDULETRANSFERTYPE_TIME_BASED ScheduleTransferType = "TIME_BASED"
)

// All allowed values of ScheduleTransferType enum
var AllowedScheduleTransferTypeEnumValues = []ScheduleTransferType{
	SCHEDULETRANSFERTYPE_TIME_BASED,
}

func (v *ScheduleTransferType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduleTransferType(value)
	for _, existing := range AllowedScheduleTransferTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleTransferType", value)
}

// NewScheduleTransferTypeFromValue returns a pointer to a valid ScheduleTransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduleTransferTypeFromValue(v string) (*ScheduleTransferType, error) {
	ev := ScheduleTransferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduleTransferType: valid values are %v", v, AllowedScheduleTransferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduleTransferType) IsValid() bool {
	for _, existing := range AllowedScheduleTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleTransferType value
func (v ScheduleTransferType) Ptr() *ScheduleTransferType {
	return &v
}

type NullableScheduleTransferType struct {
	value *ScheduleTransferType
	isSet bool
}

func (v NullableScheduleTransferType) Get() *ScheduleTransferType {
	return v.value
}

func (v *NullableScheduleTransferType) Set(val *ScheduleTransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleTransferType(val *ScheduleTransferType) *NullableScheduleTransferType {
	return &NullableScheduleTransferType{value: val, isSet: true}
}

func (v NullableScheduleTransferType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScheduleTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
