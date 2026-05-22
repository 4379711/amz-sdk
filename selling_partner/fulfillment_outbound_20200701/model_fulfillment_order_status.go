package fulfillment_outbound_20200701

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// FulfillmentOrderStatus The current status of the fulfillment order.
type FulfillmentOrderStatus string

// List of FulfillmentOrderStatus
const (
	FULFILLMENTORDERSTATUS_NEW                 FulfillmentOrderStatus = "New"
	FULFILLMENTORDERSTATUS_RECEIVED            FulfillmentOrderStatus = "Received"
	FULFILLMENTORDERSTATUS_PLANNING            FulfillmentOrderStatus = "Planning"
	FULFILLMENTORDERSTATUS_PROCESSING          FulfillmentOrderStatus = "Processing"
	FULFILLMENTORDERSTATUS_CANCELLED           FulfillmentOrderStatus = "Cancelled"
	FULFILLMENTORDERSTATUS_COMPLETE            FulfillmentOrderStatus = "Complete"
	FULFILLMENTORDERSTATUS_COMPLETE_PARTIALLED FulfillmentOrderStatus = "CompletePartialled"
	FULFILLMENTORDERSTATUS_UNFULFILLABLE       FulfillmentOrderStatus = "Unfulfillable"
	FULFILLMENTORDERSTATUS_INVALID             FulfillmentOrderStatus = "Invalid"
)

// All allowed values of FulfillmentOrderStatus enum
var AllowedFulfillmentOrderStatusEnumValues = []FulfillmentOrderStatus{
	FULFILLMENTORDERSTATUS_NEW,
	FULFILLMENTORDERSTATUS_RECEIVED,
	FULFILLMENTORDERSTATUS_PLANNING,
	FULFILLMENTORDERSTATUS_PROCESSING,
	FULFILLMENTORDERSTATUS_CANCELLED,
	FULFILLMENTORDERSTATUS_COMPLETE,
	FULFILLMENTORDERSTATUS_COMPLETE_PARTIALLED,
	FULFILLMENTORDERSTATUS_UNFULFILLABLE,
	FULFILLMENTORDERSTATUS_INVALID,
}

func (v *FulfillmentOrderStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FulfillmentOrderStatus(value)
	for _, existing := range AllowedFulfillmentOrderStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FulfillmentOrderStatus", value)
}

// NewFulfillmentOrderStatusFromValue returns a pointer to a valid FulfillmentOrderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFulfillmentOrderStatusFromValue(v string) (*FulfillmentOrderStatus, error) {
	ev := FulfillmentOrderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FulfillmentOrderStatus: valid values are %v", v, AllowedFulfillmentOrderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FulfillmentOrderStatus) IsValid() bool {
	for _, existing := range AllowedFulfillmentOrderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FulfillmentOrderStatus value
func (v FulfillmentOrderStatus) Ptr() *FulfillmentOrderStatus {
	return &v
}

type NullableFulfillmentOrderStatus struct {
	value *FulfillmentOrderStatus
	isSet bool
}

func (v NullableFulfillmentOrderStatus) Get() *FulfillmentOrderStatus {
	return v.value
}

func (v *NullableFulfillmentOrderStatus) Set(val *FulfillmentOrderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentOrderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentOrderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentOrderStatus(val *FulfillmentOrderStatus) *NullableFulfillmentOrderStatus {
	return &NullableFulfillmentOrderStatus{value: val, isSet: true}
}

func (v NullableFulfillmentOrderStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentOrderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
