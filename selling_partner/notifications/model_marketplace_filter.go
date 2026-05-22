package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the MarketplaceFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceFilter{}

// MarketplaceFilter An event filter to customize your subscription to send notifications for only the specified `marketplaceId`s.
type MarketplaceFilter struct {
	// A list of marketplace identifiers to subscribe to (for example: ATVPDKIKX0DER). To receive notifications in every marketplace, do not provide this list.
	MarketplaceIds []string `json:"marketplaceIds,omitempty"`
}

// NewMarketplaceFilter instantiates a new MarketplaceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceFilter() *MarketplaceFilter {
	this := MarketplaceFilter{}
	return &this
}

// NewMarketplaceFilterWithDefaults instantiates a new MarketplaceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceFilterWithDefaults() *MarketplaceFilter {
	this := MarketplaceFilter{}
	return &this
}

// GetMarketplaceIds returns the MarketplaceIds field value if set, zero value otherwise.
func (o *MarketplaceFilter) GetMarketplaceIds() []string {
	if o == nil || IsNil(o.MarketplaceIds) {
		var ret []string
		return ret
	}
	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceFilter) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MarketplaceIds) {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// HasMarketplaceIds returns a boolean if a field has been set.
func (o *MarketplaceFilter) HasMarketplaceIds() bool {
	if o != nil && !IsNil(o.MarketplaceIds) {
		return true
	}

	return false
}

// SetMarketplaceIds gets a reference to the given []string and assigns it to the MarketplaceIds field.
func (o *MarketplaceFilter) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

func (o MarketplaceFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceIds) {
		toSerialize["marketplaceIds"] = o.MarketplaceIds
	}
	return toSerialize, nil
}

type NullableMarketplaceFilter struct {
	value *MarketplaceFilter
	isSet bool
}

func (v NullableMarketplaceFilter) Get() *MarketplaceFilter {
	return v.value
}

func (v *NullableMarketplaceFilter) Set(val *MarketplaceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceFilter(val *MarketplaceFilter) *NullableMarketplaceFilter {
	return &NullableMarketplaceFilter{value: val, isSet: true}
}

func (v NullableMarketplaceFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMarketplaceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
