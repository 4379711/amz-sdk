package awd_20240509

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// InboundStatus The supported statuses for an inbound order.
type InboundStatus string

// List of InboundStatus
const (
	INBOUNDSTATUS_DRAFT      InboundStatus = "DRAFT"
	INBOUNDSTATUS_VALIDATING InboundStatus = "VALIDATING"
	INBOUNDSTATUS_CONFIRMED  InboundStatus = "CONFIRMED"
	INBOUNDSTATUS_CLOSED     InboundStatus = "CLOSED"
	INBOUNDSTATUS_EXPIRED    InboundStatus = "EXPIRED"
	INBOUNDSTATUS_CANCELLED  InboundStatus = "CANCELLED"
)

// All allowed values of InboundStatus enum
var AllowedInboundStatusEnumValues = []InboundStatus{
	INBOUNDSTATUS_DRAFT,
	INBOUNDSTATUS_VALIDATING,
	INBOUNDSTATUS_CONFIRMED,
	INBOUNDSTATUS_CLOSED,
	INBOUNDSTATUS_EXPIRED,
	INBOUNDSTATUS_CANCELLED,
}

func (v *InboundStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InboundStatus(value)
	for _, existing := range AllowedInboundStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InboundStatus", value)
}

// NewInboundStatusFromValue returns a pointer to a valid InboundStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInboundStatusFromValue(v string) (*InboundStatus, error) {
	ev := InboundStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InboundStatus: valid values are %v", v, AllowedInboundStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InboundStatus) IsValid() bool {
	for _, existing := range AllowedInboundStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InboundStatus value
func (v InboundStatus) Ptr() *InboundStatus {
	return &v
}

type NullableInboundStatus struct {
	value *InboundStatus
	isSet bool
}

func (v NullableInboundStatus) Get() *InboundStatus {
	return v.value
}

func (v *NullableInboundStatus) Set(val *InboundStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundStatus(val *InboundStatus) *NullableInboundStatus {
	return &NullableInboundStatus{value: val, isSet: true}
}

func (v NullableInboundStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
