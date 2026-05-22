package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordTextMarketplaceSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordTextMarketplaceSettings{}

// SponsoredProductsKeywordTextMarketplaceSettings struct for SponsoredProductsKeywordTextMarketplaceSettings
type SponsoredProductsKeywordTextMarketplaceSettings struct {
	Marketplace *SponsoredProductsMarketplace `json:"marketplace,omitempty"`
	Value       string                        `json:"value"`
}

type _SponsoredProductsKeywordTextMarketplaceSettings SponsoredProductsKeywordTextMarketplaceSettings

// NewSponsoredProductsKeywordTextMarketplaceSettings instantiates a new SponsoredProductsKeywordTextMarketplaceSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordTextMarketplaceSettings(value string) *SponsoredProductsKeywordTextMarketplaceSettings {
	this := SponsoredProductsKeywordTextMarketplaceSettings{}
	this.Value = value
	return &this
}

// NewSponsoredProductsKeywordTextMarketplaceSettingsWithDefaults instantiates a new SponsoredProductsKeywordTextMarketplaceSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordTextMarketplaceSettingsWithDefaults() *SponsoredProductsKeywordTextMarketplaceSettings {
	this := SponsoredProductsKeywordTextMarketplaceSettings{}
	return &this
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordTextMarketplaceSettings) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordTextMarketplaceSettings) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordTextMarketplaceSettings) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsKeywordTextMarketplaceSettings) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetValue returns the Value field value
func (o *SponsoredProductsKeywordTextMarketplaceSettings) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordTextMarketplaceSettings) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SponsoredProductsKeywordTextMarketplaceSettings) SetValue(v string) {
	o.Value = v
}

func (o SponsoredProductsKeywordTextMarketplaceSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordTextMarketplaceSettings struct {
	value *SponsoredProductsKeywordTextMarketplaceSettings
	isSet bool
}

func (v NullableSponsoredProductsKeywordTextMarketplaceSettings) Get() *SponsoredProductsKeywordTextMarketplaceSettings {
	return v.value
}

func (v *NullableSponsoredProductsKeywordTextMarketplaceSettings) Set(val *SponsoredProductsKeywordTextMarketplaceSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordTextMarketplaceSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordTextMarketplaceSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordTextMarketplaceSettings(val *SponsoredProductsKeywordTextMarketplaceSettings) *NullableSponsoredProductsKeywordTextMarketplaceSettings {
	return &NullableSponsoredProductsKeywordTextMarketplaceSettings{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordTextMarketplaceSettings) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordTextMarketplaceSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
