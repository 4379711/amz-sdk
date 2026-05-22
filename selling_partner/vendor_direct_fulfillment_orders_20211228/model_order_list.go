package vendor_direct_fulfillment_orders_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderList{}

// OrderList A list of purchase orders.
type OrderList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Represents a purchase order within the OrderList.
	Orders []Order `json:"orders,omitempty"`
}

// NewOrderList instantiates a new OrderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderList() *OrderList {
	this := OrderList{}
	return &this
}

// NewOrderListWithDefaults instantiates a new OrderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListWithDefaults() *OrderList {
	this := OrderList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *OrderList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *OrderList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *OrderList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderList) GetOrders() []Order {
	if o == nil || IsNil(o.Orders) {
		var ret []Order
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderList) GetOrdersOk() ([]Order, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderList) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []Order and assigns it to the Orders field.
func (o *OrderList) SetOrders(v []Order) {
	o.Orders = v
}

func (o OrderList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	return toSerialize, nil
}

type NullableOrderList struct {
	value *OrderList
	isSet bool
}

func (v NullableOrderList) Get() *OrderList {
	return v.value
}

func (v *NullableOrderList) Set(val *OrderList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderList(val *OrderList) *NullableOrderList {
	return &NullableOrderList{value: val, isSet: true}
}

func (v NullableOrderList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
