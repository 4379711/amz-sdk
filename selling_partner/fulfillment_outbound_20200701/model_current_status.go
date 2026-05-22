package fulfillment_outbound_20200701

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CurrentStatus The current delivery status of the package.
type CurrentStatus string

// List of CurrentStatus
const (
	CURRENTSTATUS_IN_TRANSIT              CurrentStatus = "IN_TRANSIT"
	CURRENTSTATUS_DELIVERED               CurrentStatus = "DELIVERED"
	CURRENTSTATUS_RETURNING               CurrentStatus = "RETURNING"
	CURRENTSTATUS_RETURNED                CurrentStatus = "RETURNED"
	CURRENTSTATUS_UNDELIVERABLE           CurrentStatus = "UNDELIVERABLE"
	CURRENTSTATUS_DELAYED                 CurrentStatus = "DELAYED"
	CURRENTSTATUS_AVAILABLE_FOR_PICKUP    CurrentStatus = "AVAILABLE_FOR_PICKUP"
	CURRENTSTATUS_CUSTOMER_ACTION         CurrentStatus = "CUSTOMER_ACTION"
	CURRENTSTATUS_UNKNOWN                 CurrentStatus = "UNKNOWN"
	CURRENTSTATUS_OUT_FOR_DELIVERY        CurrentStatus = "OUT_FOR_DELIVERY"
	CURRENTSTATUS_DELIVERY_ATTEMPTED      CurrentStatus = "DELIVERY_ATTEMPTED"
	CURRENTSTATUS_PICKUP_SUCCESSFUL       CurrentStatus = "PICKUP_SUCCESSFUL"
	CURRENTSTATUS_PICKUP_CANCELLED        CurrentStatus = "PICKUP_CANCELLED"
	CURRENTSTATUS_PICKUP_ATTEMPTED        CurrentStatus = "PICKUP_ATTEMPTED"
	CURRENTSTATUS_PICKUP_SCHEDULED        CurrentStatus = "PICKUP_SCHEDULED"
	CURRENTSTATUS_RETURN_REQUEST_ACCEPTED CurrentStatus = "RETURN_REQUEST_ACCEPTED"
	CURRENTSTATUS_REFUND_ISSUED           CurrentStatus = "REFUND_ISSUED"
	CURRENTSTATUS_RETURN_RECEIVED_IN_FC   CurrentStatus = "RETURN_RECEIVED_IN_FC"
)

// All allowed values of CurrentStatus enum
var AllowedCurrentStatusEnumValues = []CurrentStatus{
	CURRENTSTATUS_IN_TRANSIT,
	CURRENTSTATUS_DELIVERED,
	CURRENTSTATUS_RETURNING,
	CURRENTSTATUS_RETURNED,
	CURRENTSTATUS_UNDELIVERABLE,
	CURRENTSTATUS_DELAYED,
	CURRENTSTATUS_AVAILABLE_FOR_PICKUP,
	CURRENTSTATUS_CUSTOMER_ACTION,
	CURRENTSTATUS_UNKNOWN,
	CURRENTSTATUS_OUT_FOR_DELIVERY,
	CURRENTSTATUS_DELIVERY_ATTEMPTED,
	CURRENTSTATUS_PICKUP_SUCCESSFUL,
	CURRENTSTATUS_PICKUP_CANCELLED,
	CURRENTSTATUS_PICKUP_ATTEMPTED,
	CURRENTSTATUS_PICKUP_SCHEDULED,
	CURRENTSTATUS_RETURN_REQUEST_ACCEPTED,
	CURRENTSTATUS_REFUND_ISSUED,
	CURRENTSTATUS_RETURN_RECEIVED_IN_FC,
}

func (v *CurrentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CurrentStatus(value)
	for _, existing := range AllowedCurrentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CurrentStatus", value)
}

// NewCurrentStatusFromValue returns a pointer to a valid CurrentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCurrentStatusFromValue(v string) (*CurrentStatus, error) {
	ev := CurrentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CurrentStatus: valid values are %v", v, AllowedCurrentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CurrentStatus) IsValid() bool {
	for _, existing := range AllowedCurrentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CurrentStatus value
func (v CurrentStatus) Ptr() *CurrentStatus {
	return &v
}

type NullableCurrentStatus struct {
	value *CurrentStatus
	isSet bool
}

func (v NullableCurrentStatus) Get() *CurrentStatus {
	return v.value
}

func (v *NullableCurrentStatus) Set(val *CurrentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentStatus(val *CurrentStatus) *NullableCurrentStatus {
	return &NullableCurrentStatus{value: val, isSet: true}
}

func (v NullableCurrentStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCurrentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
