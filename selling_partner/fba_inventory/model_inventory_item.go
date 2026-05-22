package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the InventoryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryItem{}

// InventoryItem An item in the list of inventory to be added.
type InventoryItem struct {
	// The seller SKU of the item.
	SellerSku string `json:"sellerSku"`
	// The marketplaceId.
	MarketplaceId string `json:"marketplaceId"`
	// The quantity of item to add.
	Quantity int32 `json:"quantity"`
}

type _InventoryItem InventoryItem

// NewInventoryItem instantiates a new InventoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItem(sellerSku string, marketplaceId string, quantity int32) *InventoryItem {
	this := InventoryItem{}
	this.SellerSku = sellerSku
	this.MarketplaceId = marketplaceId
	this.Quantity = quantity
	return &this
}

// NewInventoryItemWithDefaults instantiates a new InventoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemWithDefaults() *InventoryItem {
	this := InventoryItem{}
	return &this
}

// GetSellerSku returns the SellerSku field value
func (o *InventoryItem) GetSellerSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerSku
}

// GetSellerSkuOk returns a tuple with the SellerSku field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetSellerSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerSku, true
}

// SetSellerSku sets field value
func (o *InventoryItem) SetSellerSku(v string) {
	o.SellerSku = v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *InventoryItem) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *InventoryItem) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetQuantity returns the Quantity field value
func (o *InventoryItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InventoryItem) SetQuantity(v int32) {
	o.Quantity = v
}

func (o InventoryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellerSku"] = o.SellerSku
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

type NullableInventoryItem struct {
	value *InventoryItem
	isSet bool
}

func (v NullableInventoryItem) Get() *InventoryItem {
	return v.value
}

func (v *NullableInventoryItem) Set(val *InventoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItem(val *InventoryItem) *NullableInventoryItem {
	return &NullableInventoryItem{value: val, isSet: true}
}

func (v NullableInventoryItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInventoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
