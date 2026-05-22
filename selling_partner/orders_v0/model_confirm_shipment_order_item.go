package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmShipmentOrderItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmShipmentOrderItem{}

// ConfirmShipmentOrderItem A single order item.
type ConfirmShipmentOrderItem struct {
	// The order item's unique identifier.
	OrderItemId string `json:"orderItemId"`
	// The item's quantity.
	Quantity int32 `json:"quantity"`
	// A list of order items.
	TransparencyCodes []string `json:"transparencyCodes,omitempty"`
}

type _ConfirmShipmentOrderItem ConfirmShipmentOrderItem

// NewConfirmShipmentOrderItem instantiates a new ConfirmShipmentOrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmShipmentOrderItem(orderItemId string, quantity int32) *ConfirmShipmentOrderItem {
	this := ConfirmShipmentOrderItem{}
	this.OrderItemId = orderItemId
	this.Quantity = quantity
	return &this
}

// NewConfirmShipmentOrderItemWithDefaults instantiates a new ConfirmShipmentOrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmShipmentOrderItemWithDefaults() *ConfirmShipmentOrderItem {
	this := ConfirmShipmentOrderItem{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value
func (o *ConfirmShipmentOrderItem) GetOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentOrderItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemId, true
}

// SetOrderItemId sets field value
func (o *ConfirmShipmentOrderItem) SetOrderItemId(v string) {
	o.OrderItemId = v
}

// GetQuantity returns the Quantity field value
func (o *ConfirmShipmentOrderItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentOrderItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ConfirmShipmentOrderItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetTransparencyCodes returns the TransparencyCodes field value if set, zero value otherwise.
func (o *ConfirmShipmentOrderItem) GetTransparencyCodes() []string {
	if o == nil || IsNil(o.TransparencyCodes) {
		var ret []string
		return ret
	}
	return o.TransparencyCodes
}

// GetTransparencyCodesOk returns a tuple with the TransparencyCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentOrderItem) GetTransparencyCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.TransparencyCodes) {
		return nil, false
	}
	return o.TransparencyCodes, true
}

// HasTransparencyCodes returns a boolean if a field has been set.
func (o *ConfirmShipmentOrderItem) HasTransparencyCodes() bool {
	if o != nil && !IsNil(o.TransparencyCodes) {
		return true
	}

	return false
}

// SetTransparencyCodes gets a reference to the given []string and assigns it to the TransparencyCodes field.
func (o *ConfirmShipmentOrderItem) SetTransparencyCodes(v []string) {
	o.TransparencyCodes = v
}

func (o ConfirmShipmentOrderItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderItemId"] = o.OrderItemId
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.TransparencyCodes) {
		toSerialize["transparencyCodes"] = o.TransparencyCodes
	}
	return toSerialize, nil
}

type NullableConfirmShipmentOrderItem struct {
	value *ConfirmShipmentOrderItem
	isSet bool
}

func (v NullableConfirmShipmentOrderItem) Get() *ConfirmShipmentOrderItem {
	return v.value
}

func (v *NullableConfirmShipmentOrderItem) Set(val *ConfirmShipmentOrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmShipmentOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmShipmentOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmShipmentOrderItem(val *ConfirmShipmentOrderItem) *NullableConfirmShipmentOrderItem {
	return &NullableConfirmShipmentOrderItem{value: val, isSet: true}
}

func (v NullableConfirmShipmentOrderItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmShipmentOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
