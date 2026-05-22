package awd_20240509

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// InboundShipmentStatus Possible shipment statuses used by shipments.
type InboundShipmentStatus string

// List of InboundShipmentStatus
const (
	INBOUNDSHIPMENTSTATUS_CREATED    InboundShipmentStatus = "CREATED"
	INBOUNDSHIPMENTSTATUS_SHIPPED    InboundShipmentStatus = "SHIPPED"
	INBOUNDSHIPMENTSTATUS_IN_TRANSIT InboundShipmentStatus = "IN_TRANSIT"
	INBOUNDSHIPMENTSTATUS_RECEIVING  InboundShipmentStatus = "RECEIVING"
	INBOUNDSHIPMENTSTATUS_DELIVERED  InboundShipmentStatus = "DELIVERED"
	INBOUNDSHIPMENTSTATUS_CLOSED     InboundShipmentStatus = "CLOSED"
	INBOUNDSHIPMENTSTATUS_CANCELLED  InboundShipmentStatus = "CANCELLED"
)

// All allowed values of InboundShipmentStatus enum
var AllowedInboundShipmentStatusEnumValues = []InboundShipmentStatus{
	INBOUNDSHIPMENTSTATUS_CREATED,
	INBOUNDSHIPMENTSTATUS_SHIPPED,
	INBOUNDSHIPMENTSTATUS_IN_TRANSIT,
	INBOUNDSHIPMENTSTATUS_RECEIVING,
	INBOUNDSHIPMENTSTATUS_DELIVERED,
	INBOUNDSHIPMENTSTATUS_CLOSED,
	INBOUNDSHIPMENTSTATUS_CANCELLED,
}

func (v *InboundShipmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InboundShipmentStatus(value)
	for _, existing := range AllowedInboundShipmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InboundShipmentStatus", value)
}

// NewInboundShipmentStatusFromValue returns a pointer to a valid InboundShipmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInboundShipmentStatusFromValue(v string) (*InboundShipmentStatus, error) {
	ev := InboundShipmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InboundShipmentStatus: valid values are %v", v, AllowedInboundShipmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InboundShipmentStatus) IsValid() bool {
	for _, existing := range AllowedInboundShipmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InboundShipmentStatus value
func (v InboundShipmentStatus) Ptr() *InboundShipmentStatus {
	return &v
}

type NullableInboundShipmentStatus struct {
	value *InboundShipmentStatus
	isSet bool
}

func (v NullableInboundShipmentStatus) Get() *InboundShipmentStatus {
	return v.value
}

func (v *NullableInboundShipmentStatus) Set(val *InboundShipmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundShipmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundShipmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundShipmentStatus(val *InboundShipmentStatus) *NullableInboundShipmentStatus {
	return &NullableInboundShipmentStatus{value: val, isSet: true}
}

func (v NullableInboundShipmentStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundShipmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
