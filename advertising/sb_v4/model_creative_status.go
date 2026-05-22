package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CreativeStatus The lifecycle status of a creative
type CreativeStatus string

// List of CreativeStatus
const (
	CREATIVESTATUS_SUBMITTED_FOR_MODERATION  CreativeStatus = "SUBMITTED_FOR_MODERATION"
	CREATIVESTATUS_PENDING_TRANSLATION       CreativeStatus = "PENDING_TRANSLATION"
	CREATIVESTATUS_PENDING_MODERATION_REVIEW CreativeStatus = "PENDING_MODERATION_REVIEW"
	CREATIVESTATUS_APPROVED_BY_MODERATION    CreativeStatus = "APPROVED_BY_MODERATION"
	CREATIVESTATUS_REJECTED_BY_MODERATION    CreativeStatus = "REJECTED_BY_MODERATION"
	CREATIVESTATUS_PUBLISHED                 CreativeStatus = "PUBLISHED"
)

// All allowed values of CreativeStatus enum
var AllowedCreativeStatusEnumValues = []CreativeStatus{
	"SUBMITTED_FOR_MODERATION",
	"PENDING_TRANSLATION",
	"PENDING_MODERATION_REVIEW",
	"APPROVED_BY_MODERATION",
	"REJECTED_BY_MODERATION",
	"PUBLISHED",
}

func (v *CreativeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreativeStatus(value)
	for _, existing := range AllowedCreativeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreativeStatus", value)
}

// NewCreativeStatusFromValue returns a pointer to a valid CreativeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeStatusFromValue(v string) (*CreativeStatus, error) {
	ev := CreativeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeStatus: valid values are %v", v, AllowedCreativeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeStatus) IsValid() bool {
	for _, existing := range AllowedCreativeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreativeStatus value
func (v CreativeStatus) Ptr() *CreativeStatus {
	return &v
}

type NullableCreativeStatus struct {
	value *CreativeStatus
	isSet bool
}

func (v NullableCreativeStatus) Get() *CreativeStatus {
	return v.value
}

func (v *NullableCreativeStatus) Set(val *CreativeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeStatus(val *CreativeStatus) *NullableCreativeStatus {
	return &NullableCreativeStatus{value: val, isSet: true}
}

func (v NullableCreativeStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
