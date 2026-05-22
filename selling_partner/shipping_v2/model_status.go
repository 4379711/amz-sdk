package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// Status The status of the package being shipped.
type Status string

// List of Status
const (
	STATUS_PRE_TRANSIT              Status = "PreTransit"
	STATUS_IN_TRANSIT               Status = "InTransit"
	STATUS_DELIVERED                Status = "Delivered"
	STATUS_LOST                     Status = "Lost"
	STATUS_OUT_FOR_DELIVERY         Status = "OutForDelivery"
	STATUS_REJECTED                 Status = "Rejected"
	STATUS_UNDELIVERABLE            Status = "Undeliverable"
	STATUS_DELIVERY_ATTEMPTED       Status = "DeliveryAttempted"
	STATUS_PICKUP_CANCELLED         Status = "PickupCancelled"
	STATUS_AWAITING_CUSTOMER_PICKUP Status = "AwaitingCustomerPickup"
)

// All allowed values of Status enum
var AllowedStatusEnumValues = []Status{
	STATUS_PRE_TRANSIT,
	STATUS_IN_TRANSIT,
	STATUS_DELIVERED,
	STATUS_LOST,
	STATUS_OUT_FOR_DELIVERY,
	STATUS_REJECTED,
	STATUS_UNDELIVERABLE,
	STATUS_DELIVERY_ATTEMPTED,
	STATUS_PICKUP_CANCELLED,
	STATUS_AWAITING_CUSTOMER_PICKUP,
}

func (v *Status) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Status(value)
	for _, existing := range AllowedStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Status", value)
}

// NewStatusFromValue returns a pointer to a valid Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusFromValue(v string) (*Status, error) {
	ev := Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Status: valid values are %v", v, AllowedStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Status) IsValid() bool {
	for _, existing := range AllowedStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Status value
func (v Status) Ptr() *Status {
	return &v
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
