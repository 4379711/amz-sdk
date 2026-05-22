package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ModerationStatus The moderation status of the ad.
type ModerationStatus string

// List of ModerationStatus
const (
	MODERATIONSTATUS_APPROVED    ModerationStatus = "APPROVED"
	MODERATIONSTATUS_IN_PROGRESS ModerationStatus = "IN_PROGRESS"
	MODERATIONSTATUS_REJECTED    ModerationStatus = "REJECTED"
	MODERATIONSTATUS_FAILED      ModerationStatus = "FAILED"
)

// All allowed values of ModerationStatus enum
var AllowedModerationStatusEnumValues = []ModerationStatus{
	"APPROVED",
	"IN_PROGRESS",
	"REJECTED",
	"FAILED",
}

func (v *ModerationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModerationStatus(value)
	for _, existing := range AllowedModerationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModerationStatus", value)
}

// NewModerationStatusFromValue returns a pointer to a valid ModerationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModerationStatusFromValue(v string) (*ModerationStatus, error) {
	ev := ModerationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModerationStatus: valid values are %v", v, AllowedModerationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModerationStatus) IsValid() bool {
	for _, existing := range AllowedModerationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModerationStatus value
func (v ModerationStatus) Ptr() *ModerationStatus {
	return &v
}

type NullableModerationStatus struct {
	value *ModerationStatus
	isSet bool
}

func (v NullableModerationStatus) Get() *ModerationStatus {
	return v.value
}

func (v *NullableModerationStatus) Set(val *ModerationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationStatus(val *ModerationStatus) *NullableModerationStatus {
	return &NullableModerationStatus{value: val, isSet: true}
}

func (v NullableModerationStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableModerationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
