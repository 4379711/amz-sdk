package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemIdentifiersByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemIdentifiersByMarketplace{}

// ItemIdentifiersByMarketplace Identifiers that are associated with the item in the Amazon catalog, grouped by `marketplaceId`.
type ItemIdentifiersByMarketplace struct {
	// Amazon marketplace identifier. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).identifier.
	MarketplaceId string `json:"marketplaceId"`
	// Identifiers associated with the item in the Amazon catalog for the indicated `marketplaceId`.
	Identifiers []ItemIdentifier `json:"identifiers"`
}

type _ItemIdentifiersByMarketplace ItemIdentifiersByMarketplace

// NewItemIdentifiersByMarketplace instantiates a new ItemIdentifiersByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemIdentifiersByMarketplace(marketplaceId string, identifiers []ItemIdentifier) *ItemIdentifiersByMarketplace {
	this := ItemIdentifiersByMarketplace{}
	this.MarketplaceId = marketplaceId
	this.Identifiers = identifiers
	return &this
}

// NewItemIdentifiersByMarketplaceWithDefaults instantiates a new ItemIdentifiersByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemIdentifiersByMarketplaceWithDefaults() *ItemIdentifiersByMarketplace {
	this := ItemIdentifiersByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ItemIdentifiersByMarketplace) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ItemIdentifiersByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ItemIdentifiersByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetIdentifiers returns the Identifiers field value
func (o *ItemIdentifiersByMarketplace) GetIdentifiers() []ItemIdentifier {
	if o == nil {
		var ret []ItemIdentifier
		return ret
	}

	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value
// and a boolean to check if the value has been set.
func (o *ItemIdentifiersByMarketplace) GetIdentifiersOk() ([]ItemIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// SetIdentifiers sets field value
func (o *ItemIdentifiersByMarketplace) SetIdentifiers(v []ItemIdentifier) {
	o.Identifiers = v
}

func (o ItemIdentifiersByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["identifiers"] = o.Identifiers
	return toSerialize, nil
}

type NullableItemIdentifiersByMarketplace struct {
	value *ItemIdentifiersByMarketplace
	isSet bool
}

func (v NullableItemIdentifiersByMarketplace) Get() *ItemIdentifiersByMarketplace {
	return v.value
}

func (v *NullableItemIdentifiersByMarketplace) Set(val *ItemIdentifiersByMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableItemIdentifiersByMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableItemIdentifiersByMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemIdentifiersByMarketplace(val *ItemIdentifiersByMarketplace) *NullableItemIdentifiersByMarketplace {
	return &NullableItemIdentifiersByMarketplace{value: val, isSet: true}
}

func (v NullableItemIdentifiersByMarketplace) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemIdentifiersByMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
