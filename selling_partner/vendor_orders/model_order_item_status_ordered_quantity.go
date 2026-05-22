package vendor_orders

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderItemStatusOrderedQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemStatusOrderedQuantity{}

// OrderItemStatusOrderedQuantity Ordered quantity information.
type OrderItemStatusOrderedQuantity struct {
	OrderedQuantity *ItemQuantity `json:"orderedQuantity,omitempty"`
	// Details of item quantity ordered.
	OrderedQuantityDetails []OrderedQuantityDetails `json:"orderedQuantityDetails,omitempty"`
}

// NewOrderItemStatusOrderedQuantity instantiates a new OrderItemStatusOrderedQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemStatusOrderedQuantity() *OrderItemStatusOrderedQuantity {
	this := OrderItemStatusOrderedQuantity{}
	return &this
}

// NewOrderItemStatusOrderedQuantityWithDefaults instantiates a new OrderItemStatusOrderedQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemStatusOrderedQuantityWithDefaults() *OrderItemStatusOrderedQuantity {
	this := OrderItemStatusOrderedQuantity{}
	return &this
}

// GetOrderedQuantity returns the OrderedQuantity field value if set, zero value otherwise.
func (o *OrderItemStatusOrderedQuantity) GetOrderedQuantity() ItemQuantity {
	if o == nil || IsNil(o.OrderedQuantity) {
		var ret ItemQuantity
		return ret
	}
	return *o.OrderedQuantity
}

// GetOrderedQuantityOk returns a tuple with the OrderedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemStatusOrderedQuantity) GetOrderedQuantityOk() (*ItemQuantity, bool) {
	if o == nil || IsNil(o.OrderedQuantity) {
		return nil, false
	}
	return o.OrderedQuantity, true
}

// HasOrderedQuantity returns a boolean if a field has been set.
func (o *OrderItemStatusOrderedQuantity) HasOrderedQuantity() bool {
	if o != nil && !IsNil(o.OrderedQuantity) {
		return true
	}

	return false
}

// SetOrderedQuantity gets a reference to the given ItemQuantity and assigns it to the OrderedQuantity field.
func (o *OrderItemStatusOrderedQuantity) SetOrderedQuantity(v ItemQuantity) {
	o.OrderedQuantity = &v
}

// GetOrderedQuantityDetails returns the OrderedQuantityDetails field value if set, zero value otherwise.
func (o *OrderItemStatusOrderedQuantity) GetOrderedQuantityDetails() []OrderedQuantityDetails {
	if o == nil || IsNil(o.OrderedQuantityDetails) {
		var ret []OrderedQuantityDetails
		return ret
	}
	return o.OrderedQuantityDetails
}

// GetOrderedQuantityDetailsOk returns a tuple with the OrderedQuantityDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemStatusOrderedQuantity) GetOrderedQuantityDetailsOk() ([]OrderedQuantityDetails, bool) {
	if o == nil || IsNil(o.OrderedQuantityDetails) {
		return nil, false
	}
	return o.OrderedQuantityDetails, true
}

// HasOrderedQuantityDetails returns a boolean if a field has been set.
func (o *OrderItemStatusOrderedQuantity) HasOrderedQuantityDetails() bool {
	if o != nil && !IsNil(o.OrderedQuantityDetails) {
		return true
	}

	return false
}

// SetOrderedQuantityDetails gets a reference to the given []OrderedQuantityDetails and assigns it to the OrderedQuantityDetails field.
func (o *OrderItemStatusOrderedQuantity) SetOrderedQuantityDetails(v []OrderedQuantityDetails) {
	o.OrderedQuantityDetails = v
}

func (o OrderItemStatusOrderedQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderedQuantity) {
		toSerialize["orderedQuantity"] = o.OrderedQuantity
	}
	if !IsNil(o.OrderedQuantityDetails) {
		toSerialize["orderedQuantityDetails"] = o.OrderedQuantityDetails
	}
	return toSerialize, nil
}

type NullableOrderItemStatusOrderedQuantity struct {
	value *OrderItemStatusOrderedQuantity
	isSet bool
}

func (v NullableOrderItemStatusOrderedQuantity) Get() *OrderItemStatusOrderedQuantity {
	return v.value
}

func (v *NullableOrderItemStatusOrderedQuantity) Set(val *OrderItemStatusOrderedQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemStatusOrderedQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemStatusOrderedQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemStatusOrderedQuantity(val *OrderItemStatusOrderedQuantity) *NullableOrderItemStatusOrderedQuantity {
	return &NullableOrderItemStatusOrderedQuantity{value: val, isSet: true}
}

func (v NullableOrderItemStatusOrderedQuantity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderItemStatusOrderedQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
