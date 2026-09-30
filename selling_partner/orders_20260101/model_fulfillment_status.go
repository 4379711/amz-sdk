package orders_20260101

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// FulfillmentStatus The current fulfillment status of an order, indicating where the order is in the fulfillment process from placement to handover to carrier.
type FulfillmentStatus string

// List of FulfillmentStatus
const (
	FULFILLMENTSTATUS_PENDING_AVAILABILITY FulfillmentStatus = "PENDING_AVAILABILITY"
	FULFILLMENTSTATUS_PENDING              FulfillmentStatus = "PENDING"
	FULFILLMENTSTATUS_UNSHIPPED            FulfillmentStatus = "UNSHIPPED"
	FULFILLMENTSTATUS_PARTIALLY_SHIPPED    FulfillmentStatus = "PARTIALLY_SHIPPED"
	FULFILLMENTSTATUS_SHIPPED              FulfillmentStatus = "SHIPPED"
	FULFILLMENTSTATUS_CANCELLED            FulfillmentStatus = "CANCELLED"
	FULFILLMENTSTATUS_UNFULFILLABLE        FulfillmentStatus = "UNFULFILLABLE"
	// 官方模型 JSON 未收录此值,取自 Orders API 迁移指南(对应 v0 的 InvoiceUnconfirmed);按模型重新生成时需保留。
	FULFILLMENTSTATUS_INVOICE_UNCONFIRMED FulfillmentStatus = "INVOICE_UNCONFIRMED"
)

// All allowed values of FulfillmentStatus enum
var AllowedFulfillmentStatusEnumValues = []FulfillmentStatus{
	"PENDING_AVAILABILITY",
	"PENDING",
	"UNSHIPPED",
	"PARTIALLY_SHIPPED",
	"SHIPPED",
	"CANCELLED",
	"UNFULFILLABLE",
	"INVOICE_UNCONFIRMED",
}

func (v *FulfillmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = FulfillmentStatus(value)
	return nil
}

// NewFulfillmentStatusFromValue returns a pointer to a valid FulfillmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFulfillmentStatusFromValue(v string) (*FulfillmentStatus, error) {
	ev := FulfillmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FulfillmentStatus: valid values are %v", v, AllowedFulfillmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FulfillmentStatus) IsValid() bool {
	for _, existing := range AllowedFulfillmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FulfillmentStatus value
func (v FulfillmentStatus) Ptr() *FulfillmentStatus {
	return &v
}

type NullableFulfillmentStatus struct {
	value *FulfillmentStatus
	isSet bool
}

func (v NullableFulfillmentStatus) Get() *FulfillmentStatus {
	return v.value
}

func (v *NullableFulfillmentStatus) Set(val *FulfillmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentStatus(val *FulfillmentStatus) *NullableFulfillmentStatus {
	return &NullableFulfillmentStatus{value: val, isSet: true}
}

func (v NullableFulfillmentStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
