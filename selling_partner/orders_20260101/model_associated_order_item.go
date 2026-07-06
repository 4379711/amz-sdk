package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the AssociatedOrderItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociatedOrderItem{}

// AssociatedOrderItem An associated order item that a customer has purchased with the product. For example, a tire installation service purchased with tires.
type AssociatedOrderItem struct {
	// The order identifier of the associated order item.
	OrderId *string `json:"orderId,omitempty"`
	// The order item identifier of the associated order item.
	OrderItemId *string `json:"orderItemId,omitempty"`
	// The type of association between the order items.  **Possible values**: - `VALUE_ADD_SERVICE` (The associated item is a service order)
	AssociationType *string `json:"associationType,omitempty"`
}

// NewAssociatedOrderItem instantiates a new AssociatedOrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedOrderItem() *AssociatedOrderItem {
	this := AssociatedOrderItem{}
	return &this
}

// NewAssociatedOrderItemWithDefaults instantiates a new AssociatedOrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedOrderItemWithDefaults() *AssociatedOrderItem {
	this := AssociatedOrderItem{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AssociatedOrderItem) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedOrderItem) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AssociatedOrderItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AssociatedOrderItem) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderItemId returns the OrderItemId field value if set, zero value otherwise.
func (o *AssociatedOrderItem) GetOrderItemId() string {
	if o == nil || IsNil(o.OrderItemId) {
		var ret string
		return ret
	}
	return *o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedOrderItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemId) {
		return nil, false
	}
	return o.OrderItemId, true
}

// HasOrderItemId returns a boolean if a field has been set.
func (o *AssociatedOrderItem) HasOrderItemId() bool {
	if o != nil && !IsNil(o.OrderItemId) {
		return true
	}

	return false
}

// SetOrderItemId gets a reference to the given string and assigns it to the OrderItemId field.
func (o *AssociatedOrderItem) SetOrderItemId(v string) {
	o.OrderItemId = &v
}

// GetAssociationType returns the AssociationType field value if set, zero value otherwise.
func (o *AssociatedOrderItem) GetAssociationType() string {
	if o == nil || IsNil(o.AssociationType) {
		var ret string
		return ret
	}
	return *o.AssociationType
}

// GetAssociationTypeOk returns a tuple with the AssociationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedOrderItem) GetAssociationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssociationType) {
		return nil, false
	}
	return o.AssociationType, true
}

// HasAssociationType returns a boolean if a field has been set.
func (o *AssociatedOrderItem) HasAssociationType() bool {
	if o != nil && !IsNil(o.AssociationType) {
		return true
	}

	return false
}

// SetAssociationType gets a reference to the given string and assigns it to the AssociationType field.
func (o *AssociatedOrderItem) SetAssociationType(v string) {
	o.AssociationType = &v
}

func (o AssociatedOrderItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrderItemId) {
		toSerialize["orderItemId"] = o.OrderItemId
	}
	if !IsNil(o.AssociationType) {
		toSerialize["associationType"] = o.AssociationType
	}
	return toSerialize, nil
}

type NullableAssociatedOrderItem struct {
	value *AssociatedOrderItem
	isSet bool
}

func (v NullableAssociatedOrderItem) Get() *AssociatedOrderItem {
	return v.value
}

func (v *NullableAssociatedOrderItem) Set(val *AssociatedOrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedOrderItem(val *AssociatedOrderItem) *NullableAssociatedOrderItem {
	return &NullableAssociatedOrderItem{value: val, isSet: true}
}

func (v NullableAssociatedOrderItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociatedOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
