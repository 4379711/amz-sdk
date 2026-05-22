package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// EventCode The tracking event type.
type EventCode string

// List of EventCode
const (
	EVENTCODE_READY_FOR_RECEIVE                             EventCode = "ReadyForReceive"
	EVENTCODE_PICKUP_DONE                                   EventCode = "PickupDone"
	EVENTCODE_DELIVERED                                     EventCode = "Delivered"
	EVENTCODE_DEPARTED                                      EventCode = "Departed"
	EVENTCODE_DELIVERY_ATTEMPTED                            EventCode = "DeliveryAttempted"
	EVENTCODE_LOST                                          EventCode = "Lost"
	EVENTCODE_OUT_FOR_DELIVERY                              EventCode = "OutForDelivery"
	EVENTCODE_ARRIVED_AT_CARRIER_FACILITY                   EventCode = "ArrivedAtCarrierFacility"
	EVENTCODE_REJECTED                                      EventCode = "Rejected"
	EVENTCODE_UNDELIVERABLE                                 EventCode = "Undeliverable"
	EVENTCODE_PICKUP_CANCELLED                              EventCode = "PickupCancelled"
	EVENTCODE_RETURN_INITIATED                              EventCode = "ReturnInitiated"
	EVENTCODE_AVAILABLE_FOR_PICKUP                          EventCode = "AvailableForPickup"
	EVENTCODE_RECIPIENT_REQUESTED_ALTERNATE_DELIVERY_TIMING EventCode = "RecipientRequestedAlternateDeliveryTiming"
)

// All allowed values of EventCode enum
var AllowedEventCodeEnumValues = []EventCode{
	EVENTCODE_READY_FOR_RECEIVE,
	EVENTCODE_PICKUP_DONE,
	EVENTCODE_DELIVERED,
	EVENTCODE_DEPARTED,
	EVENTCODE_DELIVERY_ATTEMPTED,
	EVENTCODE_LOST,
	EVENTCODE_OUT_FOR_DELIVERY,
	EVENTCODE_ARRIVED_AT_CARRIER_FACILITY,
	EVENTCODE_REJECTED,
	EVENTCODE_UNDELIVERABLE,
	EVENTCODE_PICKUP_CANCELLED,
	EVENTCODE_RETURN_INITIATED,
	EVENTCODE_AVAILABLE_FOR_PICKUP,
	EVENTCODE_RECIPIENT_REQUESTED_ALTERNATE_DELIVERY_TIMING,
}

func (v *EventCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventCode(value)
	for _, existing := range AllowedEventCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventCode", value)
}

// NewEventCodeFromValue returns a pointer to a valid EventCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventCodeFromValue(v string) (*EventCode, error) {
	ev := EventCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventCode: valid values are %v", v, AllowedEventCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventCode) IsValid() bool {
	for _, existing := range AllowedEventCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventCode value
func (v EventCode) Ptr() *EventCode {
	return &v
}

type NullableEventCode struct {
	value *EventCode
	isSet bool
}

func (v NullableEventCode) Get() *EventCode {
	return v.value
}

func (v *NullableEventCode) Set(val *EventCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCode(val *EventCode) *NullableEventCode {
	return &NullableEventCode{value: val, isSet: true}
}

func (v NullableEventCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEventCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
