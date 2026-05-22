package fulfillment_outbound_20200701

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// FulfillmentReturnItemStatus Indicates if the return item has been processed by a fulfillment center.
type FulfillmentReturnItemStatus string

// List of FulfillmentReturnItemStatus
const (
	FULFILLMENTRETURNITEMSTATUS_NEW       FulfillmentReturnItemStatus = "New"
	FULFILLMENTRETURNITEMSTATUS_PROCESSED FulfillmentReturnItemStatus = "Processed"
)

// All allowed values of FulfillmentReturnItemStatus enum
var AllowedFulfillmentReturnItemStatusEnumValues = []FulfillmentReturnItemStatus{
	FULFILLMENTRETURNITEMSTATUS_NEW,
	FULFILLMENTRETURNITEMSTATUS_PROCESSED,
}

func (v *FulfillmentReturnItemStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FulfillmentReturnItemStatus(value)
	for _, existing := range AllowedFulfillmentReturnItemStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FulfillmentReturnItemStatus", value)
}

// NewFulfillmentReturnItemStatusFromValue returns a pointer to a valid FulfillmentReturnItemStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFulfillmentReturnItemStatusFromValue(v string) (*FulfillmentReturnItemStatus, error) {
	ev := FulfillmentReturnItemStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FulfillmentReturnItemStatus: valid values are %v", v, AllowedFulfillmentReturnItemStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FulfillmentReturnItemStatus) IsValid() bool {
	for _, existing := range AllowedFulfillmentReturnItemStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FulfillmentReturnItemStatus value
func (v FulfillmentReturnItemStatus) Ptr() *FulfillmentReturnItemStatus {
	return &v
}

type NullableFulfillmentReturnItemStatus struct {
	value *FulfillmentReturnItemStatus
	isSet bool
}

func (v NullableFulfillmentReturnItemStatus) Get() *FulfillmentReturnItemStatus {
	return v.value
}

func (v *NullableFulfillmentReturnItemStatus) Set(val *FulfillmentReturnItemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentReturnItemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentReturnItemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentReturnItemStatus(val *FulfillmentReturnItemStatus) *NullableFulfillmentReturnItemStatus {
	return &NullableFulfillmentReturnItemStatus{value: val, isSet: true}
}

func (v NullableFulfillmentReturnItemStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentReturnItemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
