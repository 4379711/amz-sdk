package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemImagesByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemImagesByMarketplace{}

// ItemImagesByMarketplace Images for an item in the Amazon catalog, grouped by `marketplaceId`.
type ItemImagesByMarketplace struct {
	// Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string `json:"marketplaceId"`
	// Images for an item in the Amazon catalog, grouped by `marketplaceId`.
	Images []ItemImage `json:"images"`
}

type _ItemImagesByMarketplace ItemImagesByMarketplace

// NewItemImagesByMarketplace instantiates a new ItemImagesByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemImagesByMarketplace(marketplaceId string, images []ItemImage) *ItemImagesByMarketplace {
	this := ItemImagesByMarketplace{}
	this.MarketplaceId = marketplaceId
	this.Images = images
	return &this
}

// NewItemImagesByMarketplaceWithDefaults instantiates a new ItemImagesByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemImagesByMarketplaceWithDefaults() *ItemImagesByMarketplace {
	this := ItemImagesByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemImagesByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemImagesByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemImagesByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetImages returns the Images field value
func (o *ItemImagesByMarketplace) GetImages() []ItemImage {
	if o == nil {
		var ret []ItemImage
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *ItemImagesByMarketplace) GetImagesOk() ([]ItemImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *ItemImagesByMarketplace) SetImages(v []ItemImage) {
	o.Images = v
}

func (o ItemImagesByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["images"] = o.Images
	return toSerialize, nil
}

type NullableItemImagesByMarketplace struct {
	value *ItemImagesByMarketplace
	isSet bool
}

func (v NullableItemImagesByMarketplace) Get() *ItemImagesByMarketplace {
	return v.value
}

func (v *NullableItemImagesByMarketplace) Set(val *ItemImagesByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemImagesByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemImagesByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemImagesByMarketplace(val *ItemImagesByMarketplace) *NullableItemImagesByMarketplace {
	return &NullableItemImagesByMarketplace{value: val, isSet: true}
}

func (v NullableItemImagesByMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemImagesByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
