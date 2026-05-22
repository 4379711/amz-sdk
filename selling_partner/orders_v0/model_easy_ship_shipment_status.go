package orders_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// EasyShipShipmentStatus The status of the Amazon Easy Ship order. This property is only included for Amazon Easy Ship orders.
type EasyShipShipmentStatus string

// List of EasyShipShipmentStatus
const (
	EASYSHIPSHIPMENTSTATUS_PENDING_SCHEDULE    EasyShipShipmentStatus = "PendingSchedule"
	EASYSHIPSHIPMENTSTATUS_PENDING_PICK_UP     EasyShipShipmentStatus = "PendingPickUp"
	EASYSHIPSHIPMENTSTATUS_PENDING_DROP_OFF    EasyShipShipmentStatus = "PendingDropOff"
	EASYSHIPSHIPMENTSTATUS_LABEL_CANCELED      EasyShipShipmentStatus = "LabelCanceled"
	EASYSHIPSHIPMENTSTATUS_PICKED_UP           EasyShipShipmentStatus = "PickedUp"
	EASYSHIPSHIPMENTSTATUS_DROPPED_OFF         EasyShipShipmentStatus = "DroppedOff"
	EASYSHIPSHIPMENTSTATUS_AT_ORIGIN_FC        EasyShipShipmentStatus = "AtOriginFC"
	EASYSHIPSHIPMENTSTATUS_AT_DESTINATION_FC   EasyShipShipmentStatus = "AtDestinationFC"
	EASYSHIPSHIPMENTSTATUS_DELIVERED           EasyShipShipmentStatus = "Delivered"
	EASYSHIPSHIPMENTSTATUS_REJECTED_BY_BUYER   EasyShipShipmentStatus = "RejectedByBuyer"
	EASYSHIPSHIPMENTSTATUS_UNDELIVERABLE       EasyShipShipmentStatus = "Undeliverable"
	EASYSHIPSHIPMENTSTATUS_RETURNING_TO_SELLER EasyShipShipmentStatus = "ReturningToSeller"
	EASYSHIPSHIPMENTSTATUS_RETURNED_TO_SELLER  EasyShipShipmentStatus = "ReturnedToSeller"
	EASYSHIPSHIPMENTSTATUS_LOST                EasyShipShipmentStatus = "Lost"
	EASYSHIPSHIPMENTSTATUS_OUT_FOR_DELIVERY    EasyShipShipmentStatus = "OutForDelivery"
	EASYSHIPSHIPMENTSTATUS_DAMAGED             EasyShipShipmentStatus = "Damaged"
)

// All allowed values of EasyShipShipmentStatus enum
var AllowedEasyShipShipmentStatusEnumValues = []EasyShipShipmentStatus{
	EASYSHIPSHIPMENTSTATUS_PENDING_SCHEDULE,
	EASYSHIPSHIPMENTSTATUS_PENDING_PICK_UP,
	EASYSHIPSHIPMENTSTATUS_PENDING_DROP_OFF,
	EASYSHIPSHIPMENTSTATUS_LABEL_CANCELED,
	EASYSHIPSHIPMENTSTATUS_PICKED_UP,
	EASYSHIPSHIPMENTSTATUS_DROPPED_OFF,
	EASYSHIPSHIPMENTSTATUS_AT_ORIGIN_FC,
	EASYSHIPSHIPMENTSTATUS_AT_DESTINATION_FC,
	EASYSHIPSHIPMENTSTATUS_DELIVERED,
	EASYSHIPSHIPMENTSTATUS_REJECTED_BY_BUYER,
	EASYSHIPSHIPMENTSTATUS_UNDELIVERABLE,
	EASYSHIPSHIPMENTSTATUS_RETURNING_TO_SELLER,
	EASYSHIPSHIPMENTSTATUS_RETURNED_TO_SELLER,
	EASYSHIPSHIPMENTSTATUS_LOST,
	EASYSHIPSHIPMENTSTATUS_OUT_FOR_DELIVERY,
	EASYSHIPSHIPMENTSTATUS_DAMAGED,
}

func (v *EasyShipShipmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EasyShipShipmentStatus(value)
	for _, existing := range AllowedEasyShipShipmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EasyShipShipmentStatus", value)
}

// NewEasyShipShipmentStatusFromValue returns a pointer to a valid EasyShipShipmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEasyShipShipmentStatusFromValue(v string) (*EasyShipShipmentStatus, error) {
	ev := EasyShipShipmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EasyShipShipmentStatus: valid values are %v", v, AllowedEasyShipShipmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EasyShipShipmentStatus) IsValid() bool {
	for _, existing := range AllowedEasyShipShipmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EasyShipShipmentStatus value
func (v EasyShipShipmentStatus) Ptr() *EasyShipShipmentStatus {
	return &v
}

type NullableEasyShipShipmentStatus struct {
	value *EasyShipShipmentStatus
	isSet bool
}

func (v NullableEasyShipShipmentStatus) Get() *EasyShipShipmentStatus {
	return v.value
}

func (v *NullableEasyShipShipmentStatus) Set(val *EasyShipShipmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEasyShipShipmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEasyShipShipmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasyShipShipmentStatus(val *EasyShipShipmentStatus) *NullableEasyShipShipmentStatus {
	return &NullableEasyShipShipmentStatus{value: val, isSet: true}
}

func (v NullableEasyShipShipmentStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEasyShipShipmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
