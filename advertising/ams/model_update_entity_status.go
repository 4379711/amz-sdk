package ams

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// UpdateEntityStatus Update the status of the entity
type UpdateEntityStatus string

// List of UpdateEntityStatus
const (
	UPDATEENTITYSTATUS_ARCHIVED UpdateEntityStatus = "ARCHIVED"
)

// All allowed values of UpdateEntityStatus enum
var AllowedUpdateEntityStatusEnumValues = []UpdateEntityStatus{
	"ARCHIVED",
}

func (v *UpdateEntityStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpdateEntityStatus(value)
	for _, existing := range AllowedUpdateEntityStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpdateEntityStatus", value)
}

// NewUpdateEntityStatusFromValue returns a pointer to a valid UpdateEntityStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateEntityStatusFromValue(v string) (*UpdateEntityStatus, error) {
	ev := UpdateEntityStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateEntityStatus: valid values are %v", v, AllowedUpdateEntityStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateEntityStatus) IsValid() bool {
	for _, existing := range AllowedUpdateEntityStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateEntityStatus value
func (v UpdateEntityStatus) Ptr() *UpdateEntityStatus {
	return &v
}

type NullableUpdateEntityStatus struct {
	value *UpdateEntityStatus
	isSet bool
}

func (v NullableUpdateEntityStatus) Get() *UpdateEntityStatus {
	return v.value
}

func (v *NullableUpdateEntityStatus) Set(val *UpdateEntityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityStatus(val *UpdateEntityStatus) *NullableUpdateEntityStatus {
	return &NullableUpdateEntityStatus{value: val, isSet: true}
}

func (v NullableUpdateEntityStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateEntityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
