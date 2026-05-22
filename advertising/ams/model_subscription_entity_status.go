package ams

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SubscriptionEntityStatus Current status of the entity. Possible values are PROVISIONING, PENDING_CONFIRMATION, ACTIVE, ARCHIVED, FAILED_CONFIRMATION, SUSPENDED, FAILED_PROVISIONING
type SubscriptionEntityStatus string

// List of SubscriptionEntityStatus
const (
	SUBSCRIPTIONENTITYSTATUS_PROVISIONING         SubscriptionEntityStatus = "PROVISIONING"
	SUBSCRIPTIONENTITYSTATUS_PENDING_CONFIRMATION SubscriptionEntityStatus = "PENDING_CONFIRMATION"
	SUBSCRIPTIONENTITYSTATUS_ACTIVE               SubscriptionEntityStatus = "ACTIVE"
	SUBSCRIPTIONENTITYSTATUS_ARCHIVED             SubscriptionEntityStatus = "ARCHIVED"
	SUBSCRIPTIONENTITYSTATUS_FAILED_CONFIRMATION  SubscriptionEntityStatus = "FAILED_CONFIRMATION"
	SUBSCRIPTIONENTITYSTATUS_SUSPENDED            SubscriptionEntityStatus = "SUSPENDED"
	SUBSCRIPTIONENTITYSTATUS_FAILED_PROVISIONING  SubscriptionEntityStatus = "FAILED_PROVISIONING"
)

// All allowed values of SubscriptionEntityStatus enum
var AllowedSubscriptionEntityStatusEnumValues = []SubscriptionEntityStatus{
	"PROVISIONING",
	"PENDING_CONFIRMATION",
	"ACTIVE",
	"ARCHIVED",
	"FAILED_CONFIRMATION",
	"SUSPENDED",
	"FAILED_PROVISIONING",
}

func (v *SubscriptionEntityStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionEntityStatus(value)
	for _, existing := range AllowedSubscriptionEntityStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionEntityStatus", value)
}

// NewSubscriptionEntityStatusFromValue returns a pointer to a valid SubscriptionEntityStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionEntityStatusFromValue(v string) (*SubscriptionEntityStatus, error) {
	ev := SubscriptionEntityStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionEntityStatus: valid values are %v", v, AllowedSubscriptionEntityStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionEntityStatus) IsValid() bool {
	for _, existing := range AllowedSubscriptionEntityStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionEntityStatus value
func (v SubscriptionEntityStatus) Ptr() *SubscriptionEntityStatus {
	return &v
}

type NullableSubscriptionEntityStatus struct {
	value *SubscriptionEntityStatus
	isSet bool
}

func (v NullableSubscriptionEntityStatus) Get() *SubscriptionEntityStatus {
	return v.value
}

func (v *NullableSubscriptionEntityStatus) Set(val *SubscriptionEntityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionEntityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionEntityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionEntityStatus(val *SubscriptionEntityStatus) *NullableSubscriptionEntityStatus {
	return &NullableSubscriptionEntityStatus{value: val, isSet: true}
}

func (v NullableSubscriptionEntityStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubscriptionEntityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
