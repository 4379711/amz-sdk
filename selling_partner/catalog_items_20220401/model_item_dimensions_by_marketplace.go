package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemDimensionsByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemDimensionsByMarketplace{}

// ItemDimensionsByMarketplace Dimensions that are associated with the item in the Amazon catalog for the indicated `marketplaceId`.
type ItemDimensionsByMarketplace struct {
	// Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string      `json:"marketplaceId"`
	Item          *Dimensions `json:"item,omitempty"`
	Package       *Dimensions `json:"package,omitempty"`
}

type _ItemDimensionsByMarketplace ItemDimensionsByMarketplace

// NewItemDimensionsByMarketplace instantiates a new ItemDimensionsByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemDimensionsByMarketplace(marketplaceId string) *ItemDimensionsByMarketplace {
	this := ItemDimensionsByMarketplace{}
	this.MarketplaceId = marketplaceId
	return &this
}

// NewItemDimensionsByMarketplaceWithDefaults instantiates a new ItemDimensionsByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemDimensionsByMarketplaceWithDefaults() *ItemDimensionsByMarketplace {
	this := ItemDimensionsByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemDimensionsByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemDimensionsByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemDimensionsByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *ItemDimensionsByMarketplace) GetItem() Dimensions {
	if o == nil || IsNil(o.Item) {
		var ret Dimensions
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDimensionsByMarketplace) GetItemOk() (*Dimensions, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *ItemDimensionsByMarketplace) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given Dimensions and assigns it to the Item field.
func (o *ItemDimensionsByMarketplace) SetItem(v Dimensions) {
	o.Item = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *ItemDimensionsByMarketplace) GetPackage() Dimensions {
	if o == nil || IsNil(o.Package) {
		var ret Dimensions
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDimensionsByMarketplace) GetPackageOk() (*Dimensions, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *ItemDimensionsByMarketplace) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given Dimensions and assigns it to the Package field.
func (o *ItemDimensionsByMarketplace) SetPackage(v Dimensions) {
	o.Package = &v
}

func (o ItemDimensionsByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	return toSerialize, nil
}

type NullableItemDimensionsByMarketplace struct {
	value *ItemDimensionsByMarketplace
	isSet bool
}

func (v NullableItemDimensionsByMarketplace) Get() *ItemDimensionsByMarketplace {
	return v.value
}

func (v *NullableItemDimensionsByMarketplace) Set(val *ItemDimensionsByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemDimensionsByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemDimensionsByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemDimensionsByMarketplace(val *ItemDimensionsByMarketplace) *NullableItemDimensionsByMarketplace {
	return &NullableItemDimensionsByMarketplace{value: val, isSet: true}
}

func (v NullableItemDimensionsByMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemDimensionsByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
