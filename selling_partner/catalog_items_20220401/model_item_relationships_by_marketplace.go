package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemRelationshipsByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemRelationshipsByMarketplace{}

// ItemRelationshipsByMarketplace Relationship details for the Amazon catalog item for the specified Amazon `marketplaceId`.
type ItemRelationshipsByMarketplace struct {
	// Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string `json:"marketplaceId"`
	// Relationships for the item.
	Relationships []ItemRelationship `json:"relationships"`
}

type _ItemRelationshipsByMarketplace ItemRelationshipsByMarketplace

// NewItemRelationshipsByMarketplace instantiates a new ItemRelationshipsByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemRelationshipsByMarketplace(marketplaceId string, relationships []ItemRelationship) *ItemRelationshipsByMarketplace {
	this := ItemRelationshipsByMarketplace{}
	this.MarketplaceId = marketplaceId
	this.Relationships = relationships
	return &this
}

// NewItemRelationshipsByMarketplaceWithDefaults instantiates a new ItemRelationshipsByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemRelationshipsByMarketplaceWithDefaults() *ItemRelationshipsByMarketplace {
	this := ItemRelationshipsByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemRelationshipsByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemRelationshipsByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemRelationshipsByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetRelationships returns the Relationships field value
func (o *ItemRelationshipsByMarketplace) GetRelationships() []ItemRelationship {
	if o == nil {
		var ret []ItemRelationship
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *ItemRelationshipsByMarketplace) GetRelationshipsOk() ([]ItemRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relationships, true
}

// SetRelationships sets field value
func (o *ItemRelationshipsByMarketplace) SetRelationships(v []ItemRelationship) {
	o.Relationships = v
}

func (o ItemRelationshipsByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableItemRelationshipsByMarketplace struct {
	value *ItemRelationshipsByMarketplace
	isSet bool
}

func (v NullableItemRelationshipsByMarketplace) Get() *ItemRelationshipsByMarketplace {
	return v.value
}

func (v *NullableItemRelationshipsByMarketplace) Set(val *ItemRelationshipsByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemRelationshipsByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemRelationshipsByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemRelationshipsByMarketplace(val *ItemRelationshipsByMarketplace) *NullableItemRelationshipsByMarketplace {
	return &NullableItemRelationshipsByMarketplace{value: val, isSet: true}
}

func (v NullableItemRelationshipsByMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemRelationshipsByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
