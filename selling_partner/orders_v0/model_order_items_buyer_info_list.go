package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderItemsBuyerInfoList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemsBuyerInfoList{}

// OrderItemsBuyerInfoList A single order item's buyer information list with the order ID.
type OrderItemsBuyerInfoList struct {
	// A single order item's buyer information list.
	OrderItems []OrderItemBuyerInfo `json:"OrderItems"`
	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"NextToken,omitempty"`
	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`
}

type _OrderItemsBuyerInfoList OrderItemsBuyerInfoList

// NewOrderItemsBuyerInfoList instantiates a new OrderItemsBuyerInfoList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemsBuyerInfoList(orderItems []OrderItemBuyerInfo, amazonOrderId string) *OrderItemsBuyerInfoList {
	this := OrderItemsBuyerInfoList{}
	this.OrderItems = orderItems
	this.AmazonOrderId = amazonOrderId
	return &this
}

// NewOrderItemsBuyerInfoListWithDefaults instantiates a new OrderItemsBuyerInfoList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemsBuyerInfoListWithDefaults() *OrderItemsBuyerInfoList {
	this := OrderItemsBuyerInfoList{}
	return &this
}

// GetOrderItems returns the OrderItems field value
func (o *OrderItemsBuyerInfoList) GetOrderItems() []OrderItemBuyerInfo {
	if o == nil {
		var ret []OrderItemBuyerInfo
		return ret
	}

	return o.OrderItems
}

// GetOrderItemsOk returns a tuple with the OrderItems field value
// and a boolean to check if the value has been set.
func (o *OrderItemsBuyerInfoList) GetOrderItemsOk() ([]OrderItemBuyerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderItems, true
}

// SetOrderItems sets field value
func (o *OrderItemsBuyerInfoList) SetOrderItems(v []OrderItemBuyerInfo) {
	o.OrderItems = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OrderItemsBuyerInfoList) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemsBuyerInfoList) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OrderItemsBuyerInfoList) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OrderItemsBuyerInfoList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *OrderItemsBuyerInfoList) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderItemsBuyerInfoList) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *OrderItemsBuyerInfoList) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

func (o OrderItemsBuyerInfoList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["OrderItems"] = o.OrderItems
	if !IsNil(o.NextToken) {
		toSerialize["NextToken"] = o.NextToken
	}
	toSerialize["AmazonOrderId"] = o.AmazonOrderId
	return toSerialize, nil
}

type NullableOrderItemsBuyerInfoList struct {
	value *OrderItemsBuyerInfoList
	isSet bool
}

func (v NullableOrderItemsBuyerInfoList) Get() *OrderItemsBuyerInfoList {
	return v.value
}

func (v *NullableOrderItemsBuyerInfoList) Set(val *OrderItemsBuyerInfoList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemsBuyerInfoList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemsBuyerInfoList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemsBuyerInfoList(val *OrderItemsBuyerInfoList) *NullableOrderItemsBuyerInfoList {
	return &NullableOrderItemsBuyerInfoList{value: val, isSet: true}
}

func (v NullableOrderItemsBuyerInfoList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderItemsBuyerInfoList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
