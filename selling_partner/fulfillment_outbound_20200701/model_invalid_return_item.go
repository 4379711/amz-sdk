package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the InvalidReturnItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidReturnItem{}

// InvalidReturnItem An item that is invalid for return.
type InvalidReturnItem struct {
	// An identifier assigned by the seller to the return item.
	SellerReturnItemId string `json:"sellerReturnItemId"`
	// The identifier assigned to the item by the seller when the fulfillment order was created.
	SellerFulfillmentOrderItemId string            `json:"sellerFulfillmentOrderItemId"`
	InvalidItemReason            InvalidItemReason `json:"invalidItemReason"`
}

type _InvalidReturnItem InvalidReturnItem

// NewInvalidReturnItem instantiates a new InvalidReturnItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidReturnItem(sellerReturnItemId string, sellerFulfillmentOrderItemId string, invalidItemReason InvalidItemReason) *InvalidReturnItem {
	this := InvalidReturnItem{}
	this.SellerReturnItemId = sellerReturnItemId
	this.SellerFulfillmentOrderItemId = sellerFulfillmentOrderItemId
	this.InvalidItemReason = invalidItemReason
	return &this
}

// NewInvalidReturnItemWithDefaults instantiates a new InvalidReturnItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidReturnItemWithDefaults() *InvalidReturnItem {
	this := InvalidReturnItem{}
	return &this
}

// GetSellerReturnItemId returns the SellerReturnItemId field value
func (o *InvalidReturnItem) GetSellerReturnItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerReturnItemId
}

// GetSellerReturnItemIdOk returns a tuple with the SellerReturnItemId field value
// and a boolean to check if the value has been set.
func (o *InvalidReturnItem) GetSellerReturnItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerReturnItemId, true
}

// SetSellerReturnItemId sets field value
func (o *InvalidReturnItem) SetSellerReturnItemId(v string) {
	o.SellerReturnItemId = v
}

// GetSellerFulfillmentOrderItemId returns the SellerFulfillmentOrderItemId field value
func (o *InvalidReturnItem) GetSellerFulfillmentOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerFulfillmentOrderItemId
}

// GetSellerFulfillmentOrderItemIdOk returns a tuple with the SellerFulfillmentOrderItemId field value
// and a boolean to check if the value has been set.
func (o *InvalidReturnItem) GetSellerFulfillmentOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerFulfillmentOrderItemId, true
}

// SetSellerFulfillmentOrderItemId sets field value
func (o *InvalidReturnItem) SetSellerFulfillmentOrderItemId(v string) {
	o.SellerFulfillmentOrderItemId = v
}

// GetInvalidItemReason returns the InvalidItemReason field value
func (o *InvalidReturnItem) GetInvalidItemReason() InvalidItemReason {
	if o == nil {
		var ret InvalidItemReason
		return ret
	}

	return o.InvalidItemReason
}

// GetInvalidItemReasonOk returns a tuple with the InvalidItemReason field value
// and a boolean to check if the value has been set.
func (o *InvalidReturnItem) GetInvalidItemReasonOk() (*InvalidItemReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvalidItemReason, true
}

// SetInvalidItemReason sets field value
func (o *InvalidReturnItem) SetInvalidItemReason(v InvalidItemReason) {
	o.InvalidItemReason = v
}

func (o InvalidReturnItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellerReturnItemId"] = o.SellerReturnItemId
	toSerialize["sellerFulfillmentOrderItemId"] = o.SellerFulfillmentOrderItemId
	toSerialize["invalidItemReason"] = o.InvalidItemReason
	return toSerialize, nil
}

type NullableInvalidReturnItem struct {
	value *InvalidReturnItem
	isSet bool
}

func (v NullableInvalidReturnItem) Get() *InvalidReturnItem {
	return v.value
}

func (v *NullableInvalidReturnItem) Set(val *InvalidReturnItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidReturnItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidReturnItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidReturnItem(val *InvalidReturnItem) *NullableInvalidReturnItem {
	return &NullableInvalidReturnItem{value: val, isSet: true}
}

func (v NullableInvalidReturnItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInvalidReturnItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
