package vendor_direct_fulfillment_orders_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Order type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Order{}

// Order Represents a purchase order.
type Order struct {
	// The purchase order number for this order. Formatting Notes: alpha-numeric code.
	PurchaseOrderNumber string        `json:"purchaseOrderNumber"`
	OrderDetails        *OrderDetails `json:"orderDetails,omitempty"`
}

type _Order Order

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(purchaseOrderNumber string) *Order {
	this := Order{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *Order) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *Order) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *Order) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetOrderDetails returns the OrderDetails field value if set, zero value otherwise.
func (o *Order) GetOrderDetails() OrderDetails {
	if o == nil || IsNil(o.OrderDetails) {
		var ret OrderDetails
		return ret
	}
	return *o.OrderDetails
}

// GetOrderDetailsOk returns a tuple with the OrderDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetOrderDetailsOk() (*OrderDetails, bool) {
	if o == nil || IsNil(o.OrderDetails) {
		return nil, false
	}
	return o.OrderDetails, true
}

// HasOrderDetails returns a boolean if a field has been set.
func (o *Order) HasOrderDetails() bool {
	if o != nil && !IsNil(o.OrderDetails) {
		return true
	}

	return false
}

// SetOrderDetails gets a reference to the given OrderDetails and assigns it to the OrderDetails field.
func (o *Order) SetOrderDetails(v OrderDetails) {
	o.OrderDetails = &v
}

func (o Order) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	if !IsNil(o.OrderDetails) {
		toSerialize["orderDetails"] = o.OrderDetails
	}
	return toSerialize, nil
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
