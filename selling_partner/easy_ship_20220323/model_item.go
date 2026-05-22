package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the Item type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item{}

// Item Item identifier and serial number information.
type Item struct {
	// The Amazon-defined order item identifier.
	OrderItemId *string `json:"orderItemId,omitempty"`
	// A list of serial numbers for the items associated with the `OrderItemId` value.
	OrderItemSerialNumbers []string `json:"orderItemSerialNumbers,omitempty"`
}

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem() *Item {
	this := Item{}
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value if set, zero value otherwise.
func (o *Item) GetOrderItemId() string {
	if o == nil || IsNil(o.OrderItemId) {
		var ret string
		return ret
	}
	return *o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetOrderItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemId) {
		return nil, false
	}
	return o.OrderItemId, true
}

// HasOrderItemId returns a boolean if a field has been set.
func (o *Item) HasOrderItemId() bool {
	if o != nil && !IsNil(o.OrderItemId) {
		return true
	}

	return false
}

// SetOrderItemId gets a reference to the given string and assigns it to the OrderItemId field.
func (o *Item) SetOrderItemId(v string) {
	o.OrderItemId = &v
}

// GetOrderItemSerialNumbers returns the OrderItemSerialNumbers field value if set, zero value otherwise.
func (o *Item) GetOrderItemSerialNumbers() []string {
	if o == nil || IsNil(o.OrderItemSerialNumbers) {
		var ret []string
		return ret
	}
	return o.OrderItemSerialNumbers
}

// GetOrderItemSerialNumbersOk returns a tuple with the OrderItemSerialNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetOrderItemSerialNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderItemSerialNumbers) {
		return nil, false
	}
	return o.OrderItemSerialNumbers, true
}

// HasOrderItemSerialNumbers returns a boolean if a field has been set.
func (o *Item) HasOrderItemSerialNumbers() bool {
	if o != nil && !IsNil(o.OrderItemSerialNumbers) {
		return true
	}

	return false
}

// SetOrderItemSerialNumbers gets a reference to the given []string and assigns it to the OrderItemSerialNumbers field.
func (o *Item) SetOrderItemSerialNumbers(v []string) {
	o.OrderItemSerialNumbers = v
}

func (o Item) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderItemId) {
		toSerialize["orderItemId"] = o.OrderItemId
	}
	if !IsNil(o.OrderItemSerialNumbers) {
		toSerialize["orderItemSerialNumbers"] = o.OrderItemSerialNumbers
	}
	return toSerialize, nil
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
