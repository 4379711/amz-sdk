package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemIdentifiersByMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemIdentifiersByMarketplace{}

// ItemIdentifiersByMarketplace Identity attributes associated with the item in the Amazon catalog for the indicated Amazon marketplace.
type ItemIdentifiersByMarketplace struct {
	// A marketplace identifier. Identifies the Amazon marketplace for the listings item.
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// Amazon Standard Identification Number (ASIN) of the listings item.
	Asin *string `json:"asin,omitempty"`
}

// NewItemIdentifiersByMarketplace instantiates a new ItemIdentifiersByMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemIdentifiersByMarketplace() *ItemIdentifiersByMarketplace {
	this := ItemIdentifiersByMarketplace{}
	return &this
}

// NewItemIdentifiersByMarketplaceWithDefaults instantiates a new ItemIdentifiersByMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemIdentifiersByMarketplaceWithDefaults() *ItemIdentifiersByMarketplace {
	this := ItemIdentifiersByMarketplace{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *ItemIdentifiersByMarketplace) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemIdentifiersByMarketplace) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *ItemIdentifiersByMarketplace) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *ItemIdentifiersByMarketplace) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ItemIdentifiersByMarketplace) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemIdentifiersByMarketplace) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ItemIdentifiersByMarketplace) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ItemIdentifiersByMarketplace) SetAsin(v string) {
	o.Asin = &v
}

func (o ItemIdentifiersByMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
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
