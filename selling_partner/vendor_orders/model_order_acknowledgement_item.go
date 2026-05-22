package vendor_orders

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderAcknowledgementItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAcknowledgementItem{}

// OrderAcknowledgementItem Details of the item being acknowledged.
type OrderAcknowledgementItem struct {
	// Line item sequence number for the item.
	ItemSequenceNumber *string `json:"itemSequenceNumber,omitempty"`
	// Amazon Standard Identification Number (ASIN) of an item.
	AmazonProductIdentifier *string `json:"amazonProductIdentifier,omitempty"`
	// The vendor selected product identification of the item. Should be the same as was sent in the purchase order.
	VendorProductIdentifier *string      `json:"vendorProductIdentifier,omitempty"`
	OrderedQuantity         ItemQuantity `json:"orderedQuantity"`
	NetCost                 *Money       `json:"netCost,omitempty"`
	ListPrice               *Money       `json:"listPrice,omitempty"`
	// The discount multiplier that should be applied to the price if a vendor sells books with a list price. This is a multiplier factor to arrive at a final discounted price. A multiplier of .90 would be the factor if a 10% discount is given.
	DiscountMultiplier *string `json:"discountMultiplier,omitempty"`
	// This is used to indicate acknowledged quantity.
	ItemAcknowledgements []OrderItemAcknowledgement `json:"itemAcknowledgements"`
}

type _OrderAcknowledgementItem OrderAcknowledgementItem

// NewOrderAcknowledgementItem instantiates a new OrderAcknowledgementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAcknowledgementItem(orderedQuantity ItemQuantity, itemAcknowledgements []OrderItemAcknowledgement) *OrderAcknowledgementItem {
	this := OrderAcknowledgementItem{}
	this.OrderedQuantity = orderedQuantity
	this.ItemAcknowledgements = itemAcknowledgements
	return &this
}

// NewOrderAcknowledgementItemWithDefaults instantiates a new OrderAcknowledgementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAcknowledgementItemWithDefaults() *OrderAcknowledgementItem {
	this := OrderAcknowledgementItem{}
	return &this
}

// GetItemSequenceNumber returns the ItemSequenceNumber field value if set, zero value otherwise.
func (o *OrderAcknowledgementItem) GetItemSequenceNumber() string {
	if o == nil || IsNil(o.ItemSequenceNumber) {
		var ret string
		return ret
	}
	return *o.ItemSequenceNumber
}

// GetItemSequenceNumberOk returns a tuple with the ItemSequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetItemSequenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ItemSequenceNumber) {
		return nil, false
	}
	return o.ItemSequenceNumber, true
}

// HasItemSequenceNumber returns a boolean if a field has been set.
func (o *OrderAcknowledgementItem) HasItemSequenceNumber() bool {
	if o != nil && !IsNil(o.ItemSequenceNumber) {
		return true
	}

	return false
}

// SetItemSequenceNumber gets a reference to the given string and assigns it to the ItemSequenceNumber field.
func (o *OrderAcknowledgementItem) SetItemSequenceNumber(v string) {
	o.ItemSequenceNumber = &v
}

// GetAmazonProductIdentifier returns the AmazonProductIdentifier field value if set, zero value otherwise.
func (o *OrderAcknowledgementItem) GetAmazonProductIdentifier() string {
	if o == nil || IsNil(o.AmazonProductIdentifier) {
		var ret string
		return ret
	}
	return *o.AmazonProductIdentifier
}

// GetAmazonProductIdentifierOk returns a tuple with the AmazonProductIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetAmazonProductIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonProductIdentifier) {
		return nil, false
	}
	return o.AmazonProductIdentifier, true
}

// HasAmazonProductIdentifier returns a boolean if a field has been set.
func (o *OrderAcknowledgementItem) HasAmazonProductIdentifier() bool {
	if o != nil && !IsNil(o.AmazonProductIdentifier) {
		return true
	}

	return false
}

// SetAmazonProductIdentifier gets a reference to the given string and assigns it to the AmazonProductIdentifier field.
func (o *OrderAcknowledgementItem) SetAmazonProductIdentifier(v string) {
	o.AmazonProductIdentifier = &v
}

// GetVendorProductIdentifier returns the VendorProductIdentifier field value if set, zero value otherwise.
func (o *OrderAcknowledgementItem) GetVendorProductIdentifier() string {
	if o == nil || IsNil(o.VendorProductIdentifier) {
		var ret string
		return ret
	}
	return *o.VendorProductIdentifier
}

// GetVendorProductIdentifierOk returns a tuple with the VendorProductIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetVendorProductIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.VendorProductIdentifier) {
		return nil, false
	}
	return o.VendorProductIdentifier, true
}

// HasVendorProductIdentifier returns a boolean if a field has been set.
func (o *OrderAcknowledgementItem) HasVendorProductIdentifier() bool {
	if o != nil && !IsNil(o.VendorProductIdentifier) {
		return true
	}

	return false
}

// SetVendorProductIdentifier gets a reference to the given string and assigns it to the VendorProductIdentifier field.
func (o *OrderAcknowledgementItem) SetVendorProductIdentifier(v string) {
	o.VendorProductIdentifier = &v
}

// GetOrderedQuantity returns the OrderedQuantity field value
func (o *OrderAcknowledgementItem) GetOrderedQuantity() ItemQuantity {
	if o == nil {
		var ret ItemQuantity
		return ret
	}

	return o.OrderedQuantity
}

// GetOrderedQuantityOk returns a tuple with the OrderedQuantity field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetOrderedQuantityOk() (*ItemQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderedQuantity, true
}

// SetOrderedQuantity sets field value
func (o *OrderAcknowledgementItem) SetOrderedQuantity(v ItemQuantity) {
	o.OrderedQuantity = v
}

// GetNetCost returns the NetCost field value if set, zero value otherwise.
func (o *OrderAcknowledgementItem) GetNetCost() Money {
	if o == nil || IsNil(o.NetCost) {
		var ret Money
		return ret
	}
	return *o.NetCost
}

// GetNetCostOk returns a tuple with the NetCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetNetCostOk() (*Money, bool) {
	if o == nil || IsNil(o.NetCost) {
		return nil, false
	}
	return o.NetCost, true
}

// HasNetCost returns a boolean if a field has been set.
func (o *OrderAcknowledgementItem) HasNetCost() bool {
	if o != nil && !IsNil(o.NetCost) {
		return true
	}

	return false
}

// SetNetCost gets a reference to the given Money and assigns it to the NetCost field.
func (o *OrderAcknowledgementItem) SetNetCost(v Money) {
	o.NetCost = &v
}

// GetListPrice returns the ListPrice field value if set, zero value otherwise.
func (o *OrderAcknowledgementItem) GetListPrice() Money {
	if o == nil || IsNil(o.ListPrice) {
		var ret Money
		return ret
	}
	return *o.ListPrice
}

// GetListPriceOk returns a tuple with the ListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetListPriceOk() (*Money, bool) {
	if o == nil || IsNil(o.ListPrice) {
		return nil, false
	}
	return o.ListPrice, true
}

// HasListPrice returns a boolean if a field has been set.
func (o *OrderAcknowledgementItem) HasListPrice() bool {
	if o != nil && !IsNil(o.ListPrice) {
		return true
	}

	return false
}

// SetListPrice gets a reference to the given Money and assigns it to the ListPrice field.
func (o *OrderAcknowledgementItem) SetListPrice(v Money) {
	o.ListPrice = &v
}

// GetDiscountMultiplier returns the DiscountMultiplier field value if set, zero value otherwise.
func (o *OrderAcknowledgementItem) GetDiscountMultiplier() string {
	if o == nil || IsNil(o.DiscountMultiplier) {
		var ret string
		return ret
	}
	return *o.DiscountMultiplier
}

// GetDiscountMultiplierOk returns a tuple with the DiscountMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetDiscountMultiplierOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountMultiplier) {
		return nil, false
	}
	return o.DiscountMultiplier, true
}

// HasDiscountMultiplier returns a boolean if a field has been set.
func (o *OrderAcknowledgementItem) HasDiscountMultiplier() bool {
	if o != nil && !IsNil(o.DiscountMultiplier) {
		return true
	}

	return false
}

// SetDiscountMultiplier gets a reference to the given string and assigns it to the DiscountMultiplier field.
func (o *OrderAcknowledgementItem) SetDiscountMultiplier(v string) {
	o.DiscountMultiplier = &v
}

// GetItemAcknowledgements returns the ItemAcknowledgements field value
func (o *OrderAcknowledgementItem) GetItemAcknowledgements() []OrderItemAcknowledgement {
	if o == nil {
		var ret []OrderItemAcknowledgement
		return ret
	}

	return o.ItemAcknowledgements
}

// GetItemAcknowledgementsOk returns a tuple with the ItemAcknowledgements field value
// and a boolean to check if the value has been set.
func (o *OrderAcknowledgementItem) GetItemAcknowledgementsOk() ([]OrderItemAcknowledgement, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemAcknowledgements, true
}

// SetItemAcknowledgements sets field value
func (o *OrderAcknowledgementItem) SetItemAcknowledgements(v []OrderItemAcknowledgement) {
	o.ItemAcknowledgements = v
}

func (o OrderAcknowledgementItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemSequenceNumber) {
		toSerialize["itemSequenceNumber"] = o.ItemSequenceNumber
	}
	if !IsNil(o.AmazonProductIdentifier) {
		toSerialize["amazonProductIdentifier"] = o.AmazonProductIdentifier
	}
	if !IsNil(o.VendorProductIdentifier) {
		toSerialize["vendorProductIdentifier"] = o.VendorProductIdentifier
	}
	toSerialize["orderedQuantity"] = o.OrderedQuantity
	if !IsNil(o.NetCost) {
		toSerialize["netCost"] = o.NetCost
	}
	if !IsNil(o.ListPrice) {
		toSerialize["listPrice"] = o.ListPrice
	}
	if !IsNil(o.DiscountMultiplier) {
		toSerialize["discountMultiplier"] = o.DiscountMultiplier
	}
	toSerialize["itemAcknowledgements"] = o.ItemAcknowledgements
	return toSerialize, nil
}

type NullableOrderAcknowledgementItem struct {
	value *OrderAcknowledgementItem
	isSet bool
}

func (v NullableOrderAcknowledgementItem) Get() *OrderAcknowledgementItem {
	return v.value
}

func (v *NullableOrderAcknowledgementItem) Set(val *OrderAcknowledgementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAcknowledgementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAcknowledgementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAcknowledgementItem(val *OrderAcknowledgementItem) *NullableOrderAcknowledgementItem {
	return &NullableOrderAcknowledgementItem{value: val, isSet: true}
}

func (v NullableOrderAcknowledgementItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderAcknowledgementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
