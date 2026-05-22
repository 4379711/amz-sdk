package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemProductTypeByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemProductTypeByMarketplace{}

// ItemProductTypeByMarketplace Product type that is associated with the Amazon catalog item, grouped by `marketplaceId`.
type ItemProductTypeByMarketplace struct {
	// Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// Name of the product type that is associated with the Amazon catalog item.
	ProductType *string `json:"productType,omitempty"`
}

// NewItemProductTypeByMarketplace instantiates a new ItemProductTypeByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemProductTypeByMarketplace() *ItemProductTypeByMarketplace {
	this := ItemProductTypeByMarketplace{}
	return &this
}

// NewItemProductTypeByMarketplaceWithDefaults instantiates a new ItemProductTypeByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemProductTypeByMarketplaceWithDefaults() *ItemProductTypeByMarketplace {
	this := ItemProductTypeByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *ItemProductTypeByMarketplace) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProductTypeByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *ItemProductTypeByMarketplace) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *ItemProductTypeByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *ItemProductTypeByMarketplace) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProductTypeByMarketplace) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *ItemProductTypeByMarketplace) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *ItemProductTypeByMarketplace) SetProductType(v string) {
	o.ProductType = &v
}

func (o ItemProductTypeByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return toSerialize, nil
}

type NullableItemProductTypeByMarketplace struct {
	value *ItemProductTypeByMarketplace
	isSet bool
}

func (v NullableItemProductTypeByMarketplace) Get() *ItemProductTypeByMarketplace {
	return v.value
}

func (v *NullableItemProductTypeByMarketplace) Set(val *ItemProductTypeByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemProductTypeByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemProductTypeByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemProductTypeByMarketplace(val *ItemProductTypeByMarketplace) *NullableItemProductTypeByMarketplace {
	return &NullableItemProductTypeByMarketplace{value: val, isSet: true}
}

func (v NullableItemProductTypeByMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemProductTypeByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
