package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderChangeTypeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderChangeTypeFilter{}

// OrderChangeTypeFilter An event filter to customize your subscription to send notifications for only the specified `orderChangeType`.
type OrderChangeTypeFilter struct {
	// A list of order change types to subscribe to (for example: `BuyerRequestedChange`). To receive notifications of all change types, do not provide this list.
	OrderChangeTypes []OrderChangeTypeEnum `json:"orderChangeTypes,omitempty"`
}

// NewOrderChangeTypeFilter instantiates a new OrderChangeTypeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderChangeTypeFilter() *OrderChangeTypeFilter {
	this := OrderChangeTypeFilter{}
	return &this
}

// NewOrderChangeTypeFilterWithDefaults instantiates a new OrderChangeTypeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderChangeTypeFilterWithDefaults() *OrderChangeTypeFilter {
	this := OrderChangeTypeFilter{}
	return &this
}

// GetOrderChangeTypes returns the OrderChangeTypes field value if set, zero value otherwise.
func (o *OrderChangeTypeFilter) GetOrderChangeTypes() []OrderChangeTypeEnum {
	if o == nil || IsNil(o.OrderChangeTypes) {
		var ret []OrderChangeTypeEnum
		return ret
	}
	return o.OrderChangeTypes
}

// GetOrderChangeTypesOk returns a tuple with the OrderChangeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChangeTypeFilter) GetOrderChangeTypesOk() ([]OrderChangeTypeEnum, bool) {
	if o == nil || IsNil(o.OrderChangeTypes) {
		return nil, false
	}
	return o.OrderChangeTypes, true
}

// HasOrderChangeTypes returns a boolean if a field has been set.
func (o *OrderChangeTypeFilter) HasOrderChangeTypes() bool {
	if o != nil && !IsNil(o.OrderChangeTypes) {
		return true
	}

	return false
}

// SetOrderChangeTypes gets a reference to the given []OrderChangeTypeEnum and assigns it to the OrderChangeTypes field.
func (o *OrderChangeTypeFilter) SetOrderChangeTypes(v []OrderChangeTypeEnum) {
	o.OrderChangeTypes = v
}

func (o OrderChangeTypeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderChangeTypes) {
		toSerialize["orderChangeTypes"] = o.OrderChangeTypes
	}
	return toSerialize, nil
}

type NullableOrderChangeTypeFilter struct {
	value *OrderChangeTypeFilter
	isSet bool
}

func (v NullableOrderChangeTypeFilter) Get() *OrderChangeTypeFilter {
	return v.value
}

func (v *NullableOrderChangeTypeFilter) Set(val *OrderChangeTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderChangeTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderChangeTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderChangeTypeFilter(val *OrderChangeTypeFilter) *NullableOrderChangeTypeFilter {
	return &NullableOrderChangeTypeFilter{value: val, isSet: true}
}

func (v NullableOrderChangeTypeFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderChangeTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
