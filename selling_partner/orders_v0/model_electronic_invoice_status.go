package orders_v0

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// ElectronicInvoiceStatus The status of the electronic invoice. Only available for Easy Ship orders and orders in the BR marketplace.
type ElectronicInvoiceStatus string

// List of ElectronicInvoiceStatus
const (
	ELECTRONICINVOICESTATUS_NOT_REQUIRED ElectronicInvoiceStatus = "NotRequired"
	ELECTRONICINVOICESTATUS_NOT_FOUND    ElectronicInvoiceStatus = "NotFound"
	ELECTRONICINVOICESTATUS_PROCESSING   ElectronicInvoiceStatus = "Processing"
	ELECTRONICINVOICESTATUS_ERRORED      ElectronicInvoiceStatus = "Errored"
	ELECTRONICINVOICESTATUS_ACCEPTED     ElectronicInvoiceStatus = "Accepted"
)

// All allowed values of ElectronicInvoiceStatus enum
var AllowedElectronicInvoiceStatusEnumValues = []ElectronicInvoiceStatus{
	ELECTRONICINVOICESTATUS_NOT_REQUIRED,
	ELECTRONICINVOICESTATUS_NOT_FOUND,
	ELECTRONICINVOICESTATUS_PROCESSING,
	ELECTRONICINVOICESTATUS_ERRORED,
	ELECTRONICINVOICESTATUS_ACCEPTED,
}

func (v *ElectronicInvoiceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ElectronicInvoiceStatus(value)
	for _, existing := range AllowedElectronicInvoiceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ElectronicInvoiceStatus", value)
}

// NewElectronicInvoiceStatusFromValue returns a pointer to a valid ElectronicInvoiceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewElectronicInvoiceStatusFromValue(v string) (*ElectronicInvoiceStatus, error) {
	ev := ElectronicInvoiceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ElectronicInvoiceStatus: valid values are %v", v, AllowedElectronicInvoiceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ElectronicInvoiceStatus) IsValid() bool {
	for _, existing := range AllowedElectronicInvoiceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElectronicInvoiceStatus value
func (v ElectronicInvoiceStatus) Ptr() *ElectronicInvoiceStatus {
	return &v
}

type NullableElectronicInvoiceStatus struct {
	value *ElectronicInvoiceStatus
	isSet bool
}

func (v NullableElectronicInvoiceStatus) Get() *ElectronicInvoiceStatus {
	return v.value
}

func (v *NullableElectronicInvoiceStatus) Set(val *ElectronicInvoiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableElectronicInvoiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableElectronicInvoiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElectronicInvoiceStatus(val *ElectronicInvoiceStatus) *NullableElectronicInvoiceStatus {
	return &NullableElectronicInvoiceStatus{value: val, isSet: true}
}

func (v NullableElectronicInvoiceStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableElectronicInvoiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
