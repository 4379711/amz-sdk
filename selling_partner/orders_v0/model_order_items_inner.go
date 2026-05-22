package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemsInner{}

// OrderItemsInner struct for OrderItemsInner
type OrderItemsInner struct {
	// The order item's unique identifier.
	OrderItemId *string `json:"orderItemId,omitempty"`
	// The quantity for which to update the shipment status.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewOrderItemsInner instantiates a new OrderItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemsInner() *OrderItemsInner {
	this := OrderItemsInner{}
	return &this
}

// NewOrderItemsInnerWithDefaults instantiates a new OrderItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemsInnerWithDefaults() *OrderItemsInner {
	this := OrderItemsInner{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value if set, zero value otherwise.
func (o *OrderItemsInner) GetOrderItemId() string {
	if o == nil || IsNil(o.OrderItemId) {
		var ret string
		return ret
	}
	return *o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemsInner) GetOrderItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemId) {
		return nil, false
	}
	return o.OrderItemId, true
}

// HasOrderItemId returns a boolean if a field has been set.
func (o *OrderItemsInner) HasOrderItemId() bool {
	if o != nil && !IsNil(o.OrderItemId) {
		return true
	}

	return false
}

// SetOrderItemId gets a reference to the given string and assigns it to the OrderItemId field.
func (o *OrderItemsInner) SetOrderItemId(v string) {
	o.OrderItemId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderItemsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderItemsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderItemsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o OrderItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderItemId) {
		toSerialize["orderItemId"] = o.OrderItemId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableOrderItemsInner struct {
	value *OrderItemsInner
	isSet bool
}

func (v NullableOrderItemsInner) Get() *OrderItemsInner {
	return v.value
}

func (v *NullableOrderItemsInner) Set(val *OrderItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemsInner(val *OrderItemsInner) *NullableOrderItemsInner {
	return &NullableOrderItemsInner{value: val, isSet: true}
}

func (v NullableOrderItemsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
