package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the Item type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item{}

// Item Details of the item being shipped.
type Item struct {
	// Item sequence number for the item. The first item will be 001, the second 002, and so on. This number is used as a reference to refer to this item from the carton or pallet level.
	ItemSequenceNumber string `json:"itemSequenceNumber"`
	// Buyer Standard Identification Number (ASIN) of an item.
	AmazonProductIdentifier *string `json:"amazonProductIdentifier,omitempty"`
	// The vendor selected product identification of the item. Should be the same as was sent in the purchase order.
	VendorProductIdentifier *string      `json:"vendorProductIdentifier,omitempty"`
	ShippedQuantity         ItemQuantity `json:"shippedQuantity"`
	ItemDetails             *ItemDetails `json:"itemDetails,omitempty"`
}

type _Item Item

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem(itemSequenceNumber string, shippedQuantity ItemQuantity) *Item {
	this := Item{}
	this.ItemSequenceNumber = itemSequenceNumber
	this.ShippedQuantity = shippedQuantity
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetItemSequenceNumber returns the ItemSequenceNumber field value
func (o *Item) GetItemSequenceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemSequenceNumber
}

// GetItemSequenceNumberOk returns a tuple with the ItemSequenceNumber field value
// and a boolean to check if the value has been set.
func (o *Item) GetItemSequenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemSequenceNumber, true
}

// SetItemSequenceNumber sets field value
func (o *Item) SetItemSequenceNumber(v string) {
	o.ItemSequenceNumber = v
}

// GetAmazonProductIdentifier returns the AmazonProductIdentifier field value if set, zero value otherwise.
func (o *Item) GetAmazonProductIdentifier() string {
	if o == nil || IsNil(o.AmazonProductIdentifier) {
		var ret string
		return ret
	}
	return *o.AmazonProductIdentifier
}

// GetAmazonProductIdentifierOk returns a tuple with the AmazonProductIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetAmazonProductIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonProductIdentifier) {
		return nil, false
	}
	return o.AmazonProductIdentifier, true
}

// HasAmazonProductIdentifier returns a boolean if a field has been set.
func (o *Item) HasAmazonProductIdentifier() bool {
	if o != nil && !IsNil(o.AmazonProductIdentifier) {
		return true
	}

	return false
}

// SetAmazonProductIdentifier gets a reference to the given string and assigns it to the AmazonProductIdentifier field.
func (o *Item) SetAmazonProductIdentifier(v string) {
	o.AmazonProductIdentifier = &v
}

// GetVendorProductIdentifier returns the VendorProductIdentifier field value if set, zero value otherwise.
func (o *Item) GetVendorProductIdentifier() string {
	if o == nil || IsNil(o.VendorProductIdentifier) {
		var ret string
		return ret
	}
	return *o.VendorProductIdentifier
}

// GetVendorProductIdentifierOk returns a tuple with the VendorProductIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetVendorProductIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.VendorProductIdentifier) {
		return nil, false
	}
	return o.VendorProductIdentifier, true
}

// HasVendorProductIdentifier returns a boolean if a field has been set.
func (o *Item) HasVendorProductIdentifier() bool {
	if o != nil && !IsNil(o.VendorProductIdentifier) {
		return true
	}

	return false
}

// SetVendorProductIdentifier gets a reference to the given string and assigns it to the VendorProductIdentifier field.
func (o *Item) SetVendorProductIdentifier(v string) {
	o.VendorProductIdentifier = &v
}

// GetShippedQuantity returns the ShippedQuantity field value
func (o *Item) GetShippedQuantity() ItemQuantity {
	if o == nil {
		var ret ItemQuantity
		return ret
	}

	return o.ShippedQuantity
}

// GetShippedQuantityOk returns a tuple with the ShippedQuantity field value
// and a boolean to check if the value has been set.
func (o *Item) GetShippedQuantityOk() (*ItemQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippedQuantity, true
}

// SetShippedQuantity sets field value
func (o *Item) SetShippedQuantity(v ItemQuantity) {
	o.ShippedQuantity = v
}

// GetItemDetails returns the ItemDetails field value if set, zero value otherwise.
func (o *Item) GetItemDetails() ItemDetails {
	if o == nil || IsNil(o.ItemDetails) {
		var ret ItemDetails
		return ret
	}
	return *o.ItemDetails
}

// GetItemDetailsOk returns a tuple with the ItemDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemDetailsOk() (*ItemDetails, bool) {
	if o == nil || IsNil(o.ItemDetails) {
		return nil, false
	}
	return o.ItemDetails, true
}

// HasItemDetails returns a boolean if a field has been set.
func (o *Item) HasItemDetails() bool {
	if o != nil && !IsNil(o.ItemDetails) {
		return true
	}

	return false
}

// SetItemDetails gets a reference to the given ItemDetails and assigns it to the ItemDetails field.
func (o *Item) SetItemDetails(v ItemDetails) {
	o.ItemDetails = &v
}

func (o Item) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["itemSequenceNumber"] = o.ItemSequenceNumber
	if !IsNil(o.AmazonProductIdentifier) {
		toSerialize["amazonProductIdentifier"] = o.AmazonProductIdentifier
	}
	if !IsNil(o.VendorProductIdentifier) {
		toSerialize["vendorProductIdentifier"] = o.VendorProductIdentifier
	}
	toSerialize["shippedQuantity"] = o.ShippedQuantity
	if !IsNil(o.ItemDetails) {
		toSerialize["itemDetails"] = o.ItemDetails
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
