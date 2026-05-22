package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemProductTypeByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemProductTypeByMarketplace{}

// ItemProductTypeByMarketplace Product types that are associated with the listing item for the specified marketplace.
type ItemProductTypeByMarketplace struct {
	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
	// The name of the product type that is submitted by the Selling Partner.
	ProductType string `json:"productType"`
}

type _ItemProductTypeByMarketplace ItemProductTypeByMarketplace

// NewItemProductTypeByMarketplace instantiates a new ItemProductTypeByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemProductTypeByMarketplace(marketplaceId string, productType string) *ItemProductTypeByMarketplace {
	this := ItemProductTypeByMarketplace{}
	this.MarketplaceId = marketplaceId
	this.ProductType = productType
	return &this
}

// NewItemProductTypeByMarketplaceWithDefaults instantiates a new ItemProductTypeByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemProductTypeByMarketplaceWithDefaults() *ItemProductTypeByMarketplace {
	this := ItemProductTypeByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemProductTypeByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemProductTypeByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemProductTypeByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetProductType returns the ProductType field value
func (o *ItemProductTypeByMarketplace) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ItemProductTypeByMarketplace) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ItemProductTypeByMarketplace) SetProductType(v string) {
	o.ProductType = v
}

func (o ItemProductTypeByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["productType"] = o.ProductType
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
