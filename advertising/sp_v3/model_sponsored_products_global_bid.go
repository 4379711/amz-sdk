package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalBid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalBid{}

// SponsoredProductsGlobalBid struct for SponsoredProductsGlobalBid
type SponsoredProductsGlobalBid struct {
	// marketplace bid settings.
	MarketplaceSettings []SponsoredProductsMarketplaceBid `json:"marketplaceSettings,omitempty"`
	Currency            *SponsoredProductsCurrency        `json:"currency,omitempty"`
}

// NewSponsoredProductsGlobalBid instantiates a new SponsoredProductsGlobalBid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalBid() *SponsoredProductsGlobalBid {
	this := SponsoredProductsGlobalBid{}
	return &this
}

// NewSponsoredProductsGlobalBidWithDefaults instantiates a new SponsoredProductsGlobalBid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalBidWithDefaults() *SponsoredProductsGlobalBid {
	this := SponsoredProductsGlobalBid{}
	return &this
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalBid) GetMarketplaceSettings() []SponsoredProductsMarketplaceBid {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsMarketplaceBid
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalBid) GetMarketplaceSettingsOk() ([]SponsoredProductsMarketplaceBid, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalBid) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsMarketplaceBid and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsGlobalBid) SetMarketplaceSettings(v []SponsoredProductsMarketplaceBid) {
	o.MarketplaceSettings = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalBid) GetCurrency() SponsoredProductsCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret SponsoredProductsCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalBid) GetCurrencyOk() (*SponsoredProductsCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalBid) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given SponsoredProductsCurrency and assigns it to the Currency field.
func (o *SponsoredProductsGlobalBid) SetCurrency(v SponsoredProductsCurrency) {
	o.Currency = &v
}

func (o SponsoredProductsGlobalBid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalBid struct {
	value *SponsoredProductsGlobalBid
	isSet bool
}

func (v NullableSponsoredProductsGlobalBid) Get() *SponsoredProductsGlobalBid {
	return v.value
}

func (v *NullableSponsoredProductsGlobalBid) Set(val *SponsoredProductsGlobalBid) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalBid) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalBid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalBid(val *SponsoredProductsGlobalBid) *NullableSponsoredProductsGlobalBid {
	return &NullableSponsoredProductsGlobalBid{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalBid) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalBid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
