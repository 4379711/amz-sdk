package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateInventoryItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInventoryItemRequest{}

// CreateInventoryItemRequest An item to be created in the inventory.
type CreateInventoryItemRequest struct {
	// The seller SKU of the item.
	SellerSku string `json:"sellerSku"`
	// The marketplaceId.
	MarketplaceId string `json:"marketplaceId"`
	// The name of the item.
	ProductName string `json:"productName"`
}

type _CreateInventoryItemRequest CreateInventoryItemRequest

// NewCreateInventoryItemRequest instantiates a new CreateInventoryItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInventoryItemRequest(sellerSku string, marketplaceId string, productName string) *CreateInventoryItemRequest {
	this := CreateInventoryItemRequest{}
	this.SellerSku = sellerSku
	this.MarketplaceId = marketplaceId
	this.ProductName = productName
	return &this
}

// NewCreateInventoryItemRequestWithDefaults instantiates a new CreateInventoryItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInventoryItemRequestWithDefaults() *CreateInventoryItemRequest {
	this := CreateInventoryItemRequest{}
	return &this
}

// GetSellerSku returns the SellerSku field value
func (o *CreateInventoryItemRequest) GetSellerSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerSku
}

// GetSellerSkuOk returns a tuple with the SellerSku field value
// and a boolean to check if the value has been set.
func (o *CreateInventoryItemRequest) GetSellerSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerSku, true
}

// SetSellerSku sets field value
func (o *CreateInventoryItemRequest) SetSellerSku(v string) {
	o.SellerSku = v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *CreateInventoryItemRequest) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *CreateInventoryItemRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *CreateInventoryItemRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetProductName returns the ProductName field value
func (o *CreateInventoryItemRequest) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *CreateInventoryItemRequest) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *CreateInventoryItemRequest) SetProductName(v string) {
	o.ProductName = v
}

func (o CreateInventoryItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellerSku"] = o.SellerSku
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["productName"] = o.ProductName
	return toSerialize, nil
}

type NullableCreateInventoryItemRequest struct {
	value *CreateInventoryItemRequest
	isSet bool
}

func (v NullableCreateInventoryItemRequest) Get() *CreateInventoryItemRequest {
	return v.value
}

func (v *NullableCreateInventoryItemRequest) Set(val *CreateInventoryItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInventoryItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInventoryItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInventoryItemRequest(val *CreateInventoryItemRequest) *NullableCreateInventoryItemRequest {
	return &NullableCreateInventoryItemRequest{value: val, isSet: true}
}

func (v NullableCreateInventoryItemRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateInventoryItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
