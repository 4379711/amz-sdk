package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemBrowseClassificationsByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemBrowseClassificationsByMarketplace{}

// ItemBrowseClassificationsByMarketplace Classifications (browse nodes) that are associated with the item in the Amazon catalog for the indicated `marketplaceId`.
type ItemBrowseClassificationsByMarketplace struct {
	// Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string `json:"marketplaceId"`
	// Classifications (browse nodes) that are associated with the item in the Amazon catalog.
	Classifications []ItemBrowseClassification `json:"classifications,omitempty"`
}

type _ItemBrowseClassificationsByMarketplace ItemBrowseClassificationsByMarketplace

// NewItemBrowseClassificationsByMarketplace instantiates a new ItemBrowseClassificationsByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemBrowseClassificationsByMarketplace(marketplaceId string) *ItemBrowseClassificationsByMarketplace {
	this := ItemBrowseClassificationsByMarketplace{}
	this.MarketplaceId = marketplaceId
	return &this
}

// NewItemBrowseClassificationsByMarketplaceWithDefaults instantiates a new ItemBrowseClassificationsByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemBrowseClassificationsByMarketplaceWithDefaults() *ItemBrowseClassificationsByMarketplace {
	this := ItemBrowseClassificationsByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemBrowseClassificationsByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemBrowseClassificationsByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemBrowseClassificationsByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetClassifications returns the Classifications field value if set, zero value otherwise.
func (o *ItemBrowseClassificationsByMarketplace) GetClassifications() []ItemBrowseClassification {
	if o == nil || IsNil(o.Classifications) {
		var ret []ItemBrowseClassification
		return ret
	}
	return o.Classifications
}

// GetClassificationsOk returns a tuple with the Classifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemBrowseClassificationsByMarketplace) GetClassificationsOk() ([]ItemBrowseClassification, bool) {
	if o == nil || IsNil(o.Classifications) {
		return nil, false
	}
	return o.Classifications, true
}

// HasClassifications returns a boolean if a field has been set.
func (o *ItemBrowseClassificationsByMarketplace) HasClassifications() bool {
	if o != nil && !IsNil(o.Classifications) {
		return true
	}

	return false
}

// SetClassifications gets a reference to the given []ItemBrowseClassification and assigns it to the Classifications field.
func (o *ItemBrowseClassificationsByMarketplace) SetClassifications(v []ItemBrowseClassification) {
	o.Classifications = v
}

func (o ItemBrowseClassificationsByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.Classifications) {
		toSerialize["classifications"] = o.Classifications
	}
	return toSerialize, nil
}

type NullableItemBrowseClassificationsByMarketplace struct {
	value *ItemBrowseClassificationsByMarketplace
	isSet bool
}

func (v NullableItemBrowseClassificationsByMarketplace) Get() *ItemBrowseClassificationsByMarketplace {
	return v.value
}

func (v *NullableItemBrowseClassificationsByMarketplace) Set(val *ItemBrowseClassificationsByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemBrowseClassificationsByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemBrowseClassificationsByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemBrowseClassificationsByMarketplace(val *ItemBrowseClassificationsByMarketplace) *NullableItemBrowseClassificationsByMarketplace {
	return &NullableItemBrowseClassificationsByMarketplace{value: val, isSet: true}
}

func (v NullableItemBrowseClassificationsByMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemBrowseClassificationsByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
