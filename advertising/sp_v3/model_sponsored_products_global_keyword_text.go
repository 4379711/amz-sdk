package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalKeywordText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalKeywordText{}

// SponsoredProductsGlobalKeywordText struct for SponsoredProductsGlobalKeywordText
type SponsoredProductsGlobalKeywordText struct {
	// The marketplace settings for keyword text to be overridden for marketplace.
	MarketplaceSettings []SponsoredProductsKeywordTextMarketplaceSettings `json:"marketplaceSettings,omitempty"`
	Locale              *SponsoredProductsLocale                          `json:"locale,omitempty"`
	Value               string                                            `json:"value"`
}

type _SponsoredProductsGlobalKeywordText SponsoredProductsGlobalKeywordText

// NewSponsoredProductsGlobalKeywordText instantiates a new SponsoredProductsGlobalKeywordText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalKeywordText(value string) *SponsoredProductsGlobalKeywordText {
	this := SponsoredProductsGlobalKeywordText{}
	this.Value = value
	return &this
}

// NewSponsoredProductsGlobalKeywordTextWithDefaults instantiates a new SponsoredProductsGlobalKeywordText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalKeywordTextWithDefaults() *SponsoredProductsGlobalKeywordText {
	this := SponsoredProductsGlobalKeywordText{}
	return &this
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordText) GetMarketplaceSettings() []SponsoredProductsKeywordTextMarketplaceSettings {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsKeywordTextMarketplaceSettings
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordText) GetMarketplaceSettingsOk() ([]SponsoredProductsKeywordTextMarketplaceSettings, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordText) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsKeywordTextMarketplaceSettings and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsGlobalKeywordText) SetMarketplaceSettings(v []SponsoredProductsKeywordTextMarketplaceSettings) {
	o.MarketplaceSettings = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordText) GetLocale() SponsoredProductsLocale {
	if o == nil || IsNil(o.Locale) {
		var ret SponsoredProductsLocale
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordText) GetLocaleOk() (*SponsoredProductsLocale, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordText) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given SponsoredProductsLocale and assigns it to the Locale field.
func (o *SponsoredProductsGlobalKeywordText) SetLocale(v SponsoredProductsLocale) {
	o.Locale = &v
}

// GetValue returns the Value field value
func (o *SponsoredProductsGlobalKeywordText) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordText) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SponsoredProductsGlobalKeywordText) SetValue(v string) {
	o.Value = v
}

func (o SponsoredProductsGlobalKeywordText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalKeywordText struct {
	value *SponsoredProductsGlobalKeywordText
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordText) Get() *SponsoredProductsGlobalKeywordText {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordText) Set(val *SponsoredProductsGlobalKeywordText) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordText) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordText(val *SponsoredProductsGlobalKeywordText) *NullableSponsoredProductsGlobalKeywordText {
	return &NullableSponsoredProductsGlobalKeywordText{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordText) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
