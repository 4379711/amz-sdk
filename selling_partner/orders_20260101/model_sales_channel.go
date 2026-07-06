package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the SalesChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesChannel{}

// SalesChannel Information about where the customer placed this order.
type SalesChannel struct {
	// The name of the sales platform or channel where the customer placed this order.  **Possible values**: `AMAZON`, `NON_AMAZON`
	ChannelName string `json:"channelName"`
	// The unique identifier for the specific marketplace within the sales channel where this order was placed.
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// The human-readable name of the marketplace where this order was placed.
	MarketplaceName *string `json:"marketplaceName,omitempty"`
}

// NewSalesChannel instantiates a new SalesChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesChannel(channelName string) *SalesChannel {
	this := SalesChannel{}
	this.ChannelName = channelName
	return &this
}

// NewSalesChannelWithDefaults instantiates a new SalesChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesChannelWithDefaults() *SalesChannel {
	this := SalesChannel{}
	return &this
}

// GetChannelName returns the ChannelName field value
func (o *SalesChannel) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *SalesChannel) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *SalesChannel) SetChannelName(v string) {
	o.ChannelName = v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *SalesChannel) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesChannel) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *SalesChannel) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *SalesChannel) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetMarketplaceName returns the MarketplaceName field value if set, zero value otherwise.
func (o *SalesChannel) GetMarketplaceName() string {
	if o == nil || IsNil(o.MarketplaceName) {
		var ret string
		return ret
	}
	return *o.MarketplaceName
}

// GetMarketplaceNameOk returns a tuple with the MarketplaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesChannel) GetMarketplaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceName) {
		return nil, false
	}
	return o.MarketplaceName, true
}

// HasMarketplaceName returns a boolean if a field has been set.
func (o *SalesChannel) HasMarketplaceName() bool {
	if o != nil && !IsNil(o.MarketplaceName) {
		return true
	}

	return false
}

// SetMarketplaceName gets a reference to the given string and assigns it to the MarketplaceName field.
func (o *SalesChannel) SetMarketplaceName(v string) {
	o.MarketplaceName = &v
}

func (o SalesChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channelName"] = o.ChannelName
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if !IsNil(o.MarketplaceName) {
		toSerialize["marketplaceName"] = o.MarketplaceName
	}
	return toSerialize, nil
}

type NullableSalesChannel struct {
	value *SalesChannel
	isSet bool
}

func (v NullableSalesChannel) Get() *SalesChannel {
	return v.value
}

func (v *NullableSalesChannel) Set(val *SalesChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesChannel(val *SalesChannel) *NullableSalesChannel {
	return &NullableSalesChannel{value: val, isSet: true}
}

func (v NullableSalesChannel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSalesChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
