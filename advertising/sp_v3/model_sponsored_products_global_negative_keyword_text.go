package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeKeywordText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeKeywordText{}

// SponsoredProductsGlobalNegativeKeywordText struct for SponsoredProductsGlobalNegativeKeywordText
type SponsoredProductsGlobalNegativeKeywordText struct {
	// The marketplace settings for keyword text to be overridden for marketplace.
	MarketplaceSettings []SponsoredProductsKeywordTextMarketplaceSettings `json:"marketplaceSettings,omitempty"`
	Value               string                                            `json:"value"`
}

type _SponsoredProductsGlobalNegativeKeywordText SponsoredProductsGlobalNegativeKeywordText

// NewSponsoredProductsGlobalNegativeKeywordText instantiates a new SponsoredProductsGlobalNegativeKeywordText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeKeywordText(value string) *SponsoredProductsGlobalNegativeKeywordText {
	this := SponsoredProductsGlobalNegativeKeywordText{}
	this.Value = value
	return &this
}

// NewSponsoredProductsGlobalNegativeKeywordTextWithDefaults instantiates a new SponsoredProductsGlobalNegativeKeywordText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeKeywordTextWithDefaults() *SponsoredProductsGlobalNegativeKeywordText {
	this := SponsoredProductsGlobalNegativeKeywordText{}
	return &this
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeKeywordText) GetMarketplaceSettings() []SponsoredProductsKeywordTextMarketplaceSettings {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsKeywordTextMarketplaceSettings
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordText) GetMarketplaceSettingsOk() ([]SponsoredProductsKeywordTextMarketplaceSettings, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeKeywordText) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsKeywordTextMarketplaceSettings and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsGlobalNegativeKeywordText) SetMarketplaceSettings(v []SponsoredProductsKeywordTextMarketplaceSettings) {
	o.MarketplaceSettings = v
}

// GetValue returns the Value field value
func (o *SponsoredProductsGlobalNegativeKeywordText) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordText) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SponsoredProductsGlobalNegativeKeywordText) SetValue(v string) {
	o.Value = v
}

func (o SponsoredProductsGlobalNegativeKeywordText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalNegativeKeywordText struct {
	value *SponsoredProductsGlobalNegativeKeywordText
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeKeywordText) Get() *SponsoredProductsGlobalNegativeKeywordText {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordText) Set(val *SponsoredProductsGlobalNegativeKeywordText) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeKeywordText) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeKeywordText(val *SponsoredProductsGlobalNegativeKeywordText) *NullableSponsoredProductsGlobalNegativeKeywordText {
	return &NullableSponsoredProductsGlobalNegativeKeywordText{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeKeywordText) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
