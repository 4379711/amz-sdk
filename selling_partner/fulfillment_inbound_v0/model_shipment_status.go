package fulfillment_inbound_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ShipmentStatus Indicates the status of the inbound shipment. When used with the createInboundShipment operation, WORKING is the only valid value. When used with the updateInboundShipment operation, possible values are WORKING, SHIPPED or CANCELLED.
type ShipmentStatus string

// List of ShipmentStatus
const (
	SHIPMENTSTATUS_UNCONFIRMED   ShipmentStatus = "UNCONFIRMED"
	SHIPMENTSTATUS_WORKING       ShipmentStatus = "WORKING"
	SHIPMENTSTATUS_READY_TO_SHIP ShipmentStatus = "READY_TO_SHIP"
	SHIPMENTSTATUS_ABANDONED     ShipmentStatus = "ABANDONED"
	SHIPMENTSTATUS_MIXED         ShipmentStatus = "MIXED"
	SHIPMENTSTATUS_SHIPPED       ShipmentStatus = "SHIPPED"
	SHIPMENTSTATUS_RECEIVING     ShipmentStatus = "RECEIVING"
	SHIPMENTSTATUS_CANCELLED     ShipmentStatus = "CANCELLED"
	SHIPMENTSTATUS_DELETED       ShipmentStatus = "DELETED"
	SHIPMENTSTATUS_CLOSED        ShipmentStatus = "CLOSED"
	SHIPMENTSTATUS_ERROR         ShipmentStatus = "ERROR"
	SHIPMENTSTATUS_IN_TRANSIT    ShipmentStatus = "IN_TRANSIT"
	SHIPMENTSTATUS_DELIVERED     ShipmentStatus = "DELIVERED"
	SHIPMENTSTATUS_CHECKED_IN    ShipmentStatus = "CHECKED_IN"
)

// All allowed values of ShipmentStatus enum
var AllowedShipmentStatusEnumValues = []ShipmentStatus{
	SHIPMENTSTATUS_WORKING,
	SHIPMENTSTATUS_UNCONFIRMED,
	SHIPMENTSTATUS_READY_TO_SHIP,
	SHIPMENTSTATUS_MIXED,
	SHIPMENTSTATUS_ABANDONED,
	SHIPMENTSTATUS_SHIPPED,
	SHIPMENTSTATUS_RECEIVING,
	SHIPMENTSTATUS_CANCELLED,
	SHIPMENTSTATUS_DELETED,
	SHIPMENTSTATUS_CLOSED,
	SHIPMENTSTATUS_ERROR,
	SHIPMENTSTATUS_IN_TRANSIT,
	SHIPMENTSTATUS_DELIVERED,
	SHIPMENTSTATUS_CHECKED_IN,
}

func (v *ShipmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipmentStatus(value)
	for _, existing := range AllowedShipmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipmentStatus", value)
}

// NewShipmentStatusFromValue returns a pointer to a valid ShipmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentStatusFromValue(v string) (*ShipmentStatus, error) {
	ev := ShipmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentStatus: valid values are %v", v, AllowedShipmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentStatus) IsValid() bool {
	for _, existing := range AllowedShipmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipmentStatus value
func (v ShipmentStatus) Ptr() *ShipmentStatus {
	return &v
}

type NullableShipmentStatus struct {
	value *ShipmentStatus
	isSet bool
}

func (v NullableShipmentStatus) Get() *ShipmentStatus {
	return v.value
}

func (v *NullableShipmentStatus) Set(val *ShipmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentStatus(val *ShipmentStatus) *NullableShipmentStatus {
	return &NullableShipmentStatus{value: val, isSet: true}
}

func (v NullableShipmentStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
