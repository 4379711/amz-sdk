package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the AssociatedOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociatedOrder{}

// AssociatedOrder Another order that has a direct business relationship with the current order, such as replacements or exchanges.
type AssociatedOrder struct {
	// The unique identifier of the related order that is associated with the current order.
	OrderId *string `json:"orderId,omitempty"`
	// The relationship between the current order and the associated order.  **Possible values**: `REPLACEMENT_ORIGINAL_ID`, `EXCHANGE_ORIGINAL_ID`
	AssociationType *string `json:"associationType,omitempty"`
}

// NewAssociatedOrder instantiates a new AssociatedOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedOrder() *AssociatedOrder {
	this := AssociatedOrder{}
	return &this
}

// NewAssociatedOrderWithDefaults instantiates a new AssociatedOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedOrderWithDefaults() *AssociatedOrder {
	this := AssociatedOrder{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AssociatedOrder) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedOrder) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AssociatedOrder) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AssociatedOrder) SetOrderId(v string) {
	o.OrderId = &v
}

// GetAssociationType returns the AssociationType field value if set, zero value otherwise.
func (o *AssociatedOrder) GetAssociationType() string {
	if o == nil || IsNil(o.AssociationType) {
		var ret string
		return ret
	}
	return *o.AssociationType
}

// GetAssociationTypeOk returns a tuple with the AssociationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedOrder) GetAssociationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssociationType) {
		return nil, false
	}
	return o.AssociationType, true
}

// HasAssociationType returns a boolean if a field has been set.
func (o *AssociatedOrder) HasAssociationType() bool {
	if o != nil && !IsNil(o.AssociationType) {
		return true
	}

	return false
}

// SetAssociationType gets a reference to the given string and assigns it to the AssociationType field.
func (o *AssociatedOrder) SetAssociationType(v string) {
	o.AssociationType = &v
}

func (o AssociatedOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.AssociationType) {
		toSerialize["associationType"] = o.AssociationType
	}
	return toSerialize, nil
}

type NullableAssociatedOrder struct {
	value *AssociatedOrder
	isSet bool
}

func (v NullableAssociatedOrder) Get() *AssociatedOrder {
	return v.value
}

func (v *NullableAssociatedOrder) Set(val *AssociatedOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedOrder(val *AssociatedOrder) *NullableAssociatedOrder {
	return &NullableAssociatedOrder{value: val, isSet: true}
}

func (v NullableAssociatedOrder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociatedOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
