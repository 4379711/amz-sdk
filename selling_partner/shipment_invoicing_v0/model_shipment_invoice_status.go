package shipment_invoicing_v0

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ShipmentInvoiceStatus The shipment invoice status.
type ShipmentInvoiceStatus string

// List of ShipmentInvoiceStatus
const (
	SHIPMENTINVOICESTATUS_PROCESSING ShipmentInvoiceStatus = "Processing"
	SHIPMENTINVOICESTATUS_ACCEPTED   ShipmentInvoiceStatus = "Accepted"
	SHIPMENTINVOICESTATUS_ERRORED    ShipmentInvoiceStatus = "Errored"
	SHIPMENTINVOICESTATUS_NOT_FOUND  ShipmentInvoiceStatus = "NotFound"
)

// All allowed values of ShipmentInvoiceStatus enum
var AllowedShipmentInvoiceStatusEnumValues = []ShipmentInvoiceStatus{
	SHIPMENTINVOICESTATUS_PROCESSING,
	SHIPMENTINVOICESTATUS_ACCEPTED,
	SHIPMENTINVOICESTATUS_ERRORED,
	SHIPMENTINVOICESTATUS_NOT_FOUND,
}

func (v *ShipmentInvoiceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = ShipmentInvoiceStatus(value)
	return nil
}

// NewShipmentInvoiceStatusFromValue returns a pointer to a valid ShipmentInvoiceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentInvoiceStatusFromValue(v string) (*ShipmentInvoiceStatus, error) {
	ev := ShipmentInvoiceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentInvoiceStatus: valid values are %v", v, AllowedShipmentInvoiceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentInvoiceStatus) IsValid() bool {
	for _, existing := range AllowedShipmentInvoiceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipmentInvoiceStatus value
func (v ShipmentInvoiceStatus) Ptr() *ShipmentInvoiceStatus {
	return &v
}

type NullableShipmentInvoiceStatus struct {
	value *ShipmentInvoiceStatus
	isSet bool
}

func (v NullableShipmentInvoiceStatus) Get() *ShipmentInvoiceStatus {
	return v.value
}

func (v *NullableShipmentInvoiceStatus) Set(val *ShipmentInvoiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInvoiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInvoiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInvoiceStatus(val *ShipmentInvoiceStatus) *NullableShipmentInvoiceStatus {
	return &NullableShipmentInvoiceStatus{value: val, isSet: true}
}

func (v NullableShipmentInvoiceStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentInvoiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
