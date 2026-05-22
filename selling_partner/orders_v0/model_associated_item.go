package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AssociatedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociatedItem{}

// AssociatedItem An item that is associated with an order item. For example, a tire installation service that is purchased with tires.
type AssociatedItem struct {
	// The order item's order identifier, in 3-7-7 format.
	OrderId *string `json:"OrderId,omitempty"`
	// An Amazon-defined item identifier for the associated item.
	OrderItemId     *string          `json:"OrderItemId,omitempty"`
	AssociationType *AssociationType `json:"AssociationType,omitempty"`
}

// NewAssociatedItem instantiates a new AssociatedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedItem() *AssociatedItem {
	this := AssociatedItem{}
	return &this
}

// NewAssociatedItemWithDefaults instantiates a new AssociatedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedItemWithDefaults() *AssociatedItem {
	this := AssociatedItem{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AssociatedItem) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedItem) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AssociatedItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AssociatedItem) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderItemId returns the OrderItemId field value if set, zero value otherwise.
func (o *AssociatedItem) GetOrderItemId() string {
	if o == nil || IsNil(o.OrderItemId) {
		var ret string
		return ret
	}
	return *o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemId) {
		return nil, false
	}
	return o.OrderItemId, true
}

// HasOrderItemId returns a boolean if a field has been set.
func (o *AssociatedItem) HasOrderItemId() bool {
	if o != nil && !IsNil(o.OrderItemId) {
		return true
	}

	return false
}

// SetOrderItemId gets a reference to the given string and assigns it to the OrderItemId field.
func (o *AssociatedItem) SetOrderItemId(v string) {
	o.OrderItemId = &v
}

// GetAssociationType returns the AssociationType field value if set, zero value otherwise.
func (o *AssociatedItem) GetAssociationType() AssociationType {
	if o == nil || IsNil(o.AssociationType) {
		var ret AssociationType
		return ret
	}
	return *o.AssociationType
}

// GetAssociationTypeOk returns a tuple with the AssociationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedItem) GetAssociationTypeOk() (*AssociationType, bool) {
	if o == nil || IsNil(o.AssociationType) {
		return nil, false
	}
	return o.AssociationType, true
}

// HasAssociationType returns a boolean if a field has been set.
func (o *AssociatedItem) HasAssociationType() bool {
	if o != nil && !IsNil(o.AssociationType) {
		return true
	}

	return false
}

// SetAssociationType gets a reference to the given AssociationType and assigns it to the AssociationType field.
func (o *AssociatedItem) SetAssociationType(v AssociationType) {
	o.AssociationType = &v
}

func (o AssociatedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["OrderId"] = o.OrderId
	}
	if !IsNil(o.OrderItemId) {
		toSerialize["OrderItemId"] = o.OrderItemId
	}
	if !IsNil(o.AssociationType) {
		toSerialize["AssociationType"] = o.AssociationType
	}
	return toSerialize, nil
}

type NullableAssociatedItem struct {
	value *AssociatedItem
	isSet bool
}

func (v NullableAssociatedItem) Get() *AssociatedItem {
	return v.value
}

func (v *NullableAssociatedItem) Set(val *AssociatedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedItem(val *AssociatedItem) *NullableAssociatedItem {
	return &NullableAssociatedItem{value: val, isSet: true}
}

func (v NullableAssociatedItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociatedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
