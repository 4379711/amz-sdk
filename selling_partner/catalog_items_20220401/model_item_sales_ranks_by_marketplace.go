package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemSalesRanksByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemSalesRanksByMarketplace{}

// ItemSalesRanksByMarketplace Sales ranks of an Amazon catalog item, grouped by `marketplaceId`.
type ItemSalesRanksByMarketplace struct {
	// Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string `json:"marketplaceId"`
	// Sales ranks of an Amazon catalog item for a `marketplaceId`, grouped by classification.
	ClassificationRanks []ItemClassificationSalesRank `json:"classificationRanks,omitempty"`
	// Sales ranks of an Amazon catalog item for a `marketplaceId`, grouped by website display group.
	DisplayGroupRanks []ItemDisplayGroupSalesRank `json:"displayGroupRanks,omitempty"`
}

type _ItemSalesRanksByMarketplace ItemSalesRanksByMarketplace

// NewItemSalesRanksByMarketplace instantiates a new ItemSalesRanksByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemSalesRanksByMarketplace(marketplaceId string) *ItemSalesRanksByMarketplace {
	this := ItemSalesRanksByMarketplace{}
	this.MarketplaceId = marketplaceId
	return &this
}

// NewItemSalesRanksByMarketplaceWithDefaults instantiates a new ItemSalesRanksByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemSalesRanksByMarketplaceWithDefaults() *ItemSalesRanksByMarketplace {
	this := ItemSalesRanksByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemSalesRanksByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemSalesRanksByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemSalesRanksByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetClassificationRanks returns the ClassificationRanks field value if set, zero value otherwise.
func (o *ItemSalesRanksByMarketplace) GetClassificationRanks() []ItemClassificationSalesRank {
	if o == nil || IsNil(o.ClassificationRanks) {
		var ret []ItemClassificationSalesRank
		return ret
	}
	return o.ClassificationRanks
}

// GetClassificationRanksOk returns a tuple with the ClassificationRanks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSalesRanksByMarketplace) GetClassificationRanksOk() ([]ItemClassificationSalesRank, bool) {
	if o == nil || IsNil(o.ClassificationRanks) {
		return nil, false
	}
	return o.ClassificationRanks, true
}

// HasClassificationRanks returns a boolean if a field has been set.
func (o *ItemSalesRanksByMarketplace) HasClassificationRanks() bool {
	if o != nil && !IsNil(o.ClassificationRanks) {
		return true
	}

	return false
}

// SetClassificationRanks gets a reference to the given []ItemClassificationSalesRank and assigns it to the ClassificationRanks field.
func (o *ItemSalesRanksByMarketplace) SetClassificationRanks(v []ItemClassificationSalesRank) {
	o.ClassificationRanks = v
}

// GetDisplayGroupRanks returns the DisplayGroupRanks field value if set, zero value otherwise.
func (o *ItemSalesRanksByMarketplace) GetDisplayGroupRanks() []ItemDisplayGroupSalesRank {
	if o == nil || IsNil(o.DisplayGroupRanks) {
		var ret []ItemDisplayGroupSalesRank
		return ret
	}
	return o.DisplayGroupRanks
}

// GetDisplayGroupRanksOk returns a tuple with the DisplayGroupRanks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSalesRanksByMarketplace) GetDisplayGroupRanksOk() ([]ItemDisplayGroupSalesRank, bool) {
	if o == nil || IsNil(o.DisplayGroupRanks) {
		return nil, false
	}
	return o.DisplayGroupRanks, true
}

// HasDisplayGroupRanks returns a boolean if a field has been set.
func (o *ItemSalesRanksByMarketplace) HasDisplayGroupRanks() bool {
	if o != nil && !IsNil(o.DisplayGroupRanks) {
		return true
	}

	return false
}

// SetDisplayGroupRanks gets a reference to the given []ItemDisplayGroupSalesRank and assigns it to the DisplayGroupRanks field.
func (o *ItemSalesRanksByMarketplace) SetDisplayGroupRanks(v []ItemDisplayGroupSalesRank) {
	o.DisplayGroupRanks = v
}

func (o ItemSalesRanksByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.ClassificationRanks) {
		toSerialize["classificationRanks"] = o.ClassificationRanks
	}
	if !IsNil(o.DisplayGroupRanks) {
		toSerialize["displayGroupRanks"] = o.DisplayGroupRanks
	}
	return toSerialize, nil
}

type NullableItemSalesRanksByMarketplace struct {
	value *ItemSalesRanksByMarketplace
	isSet bool
}

func (v NullableItemSalesRanksByMarketplace) Get() *ItemSalesRanksByMarketplace {
	return v.value
}

func (v *NullableItemSalesRanksByMarketplace) Set(val *ItemSalesRanksByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemSalesRanksByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemSalesRanksByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemSalesRanksByMarketplace(val *ItemSalesRanksByMarketplace) *NullableItemSalesRanksByMarketplace {
	return &NullableItemSalesRanksByMarketplace{value: val, isSet: true}
}

func (v NullableItemSalesRanksByMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemSalesRanksByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
