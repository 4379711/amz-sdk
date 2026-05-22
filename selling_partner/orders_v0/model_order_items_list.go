package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderItemsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemsList{}

// OrderItemsList The order items list along with the order ID.
type OrderItemsList struct {
	// A list of order items.
	OrderItems []OrderItem `json:"OrderItems"`
	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"NextToken,omitempty"`
	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`
}

type _OrderItemsList OrderItemsList

// NewOrderItemsList instantiates a new OrderItemsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemsList(orderItems []OrderItem, amazonOrderId string) *OrderItemsList {
	this := OrderItemsList{}
	this.OrderItems = orderItems
	this.AmazonOrderId = amazonOrderId
	return &this
}

// NewOrderItemsListWithDefaults instantiates a new OrderItemsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemsListWithDefaults() *OrderItemsList {
	this := OrderItemsList{}
	return &this
}

// GetOrderItems returns the OrderItems field value
func (o *OrderItemsList) GetOrderItems() []OrderItem {
	if o == nil {
		var ret []OrderItem
		return ret
	}

	return o.OrderItems
}

// GetOrderItemsOk returns a tuple with the OrderItems field value
// and a boolean to check if the value has been set.
func (o *OrderItemsList) GetOrderItemsOk() ([]OrderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderItems, true
}

// SetOrderItems sets field value
func (o *OrderItemsList) SetOrderItems(v []OrderItem) {
	o.OrderItems = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OrderItemsList) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemsList) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OrderItemsList) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OrderItemsList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *OrderItemsList) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderItemsList) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *OrderItemsList) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

func (o OrderItemsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["OrderItems"] = o.OrderItems
	if !IsNil(o.NextToken) {
		toSerialize["NextToken"] = o.NextToken
	}
	toSerialize["AmazonOrderId"] = o.AmazonOrderId
	return toSerialize, nil
}

type NullableOrderItemsList struct {
	value *OrderItemsList
	isSet bool
}

func (v NullableOrderItemsList) Get() *OrderItemsList {
	return v.value
}

func (v *NullableOrderItemsList) Set(val *OrderItemsList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemsList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemsList(val *OrderItemsList) *NullableOrderItemsList {
	return &NullableOrderItemsList{value: val, isSet: true}
}

func (v NullableOrderItemsList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderItemsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
