package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalProductIdentifiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalProductIdentifiers{}

// SponsoredProductsGlobalProductIdentifiers struct for SponsoredProductsGlobalProductIdentifiers
type SponsoredProductsGlobalProductIdentifiers struct {
	MarketplaceSettings []SponsoredProductsMarketplaceLevelProductIdentifier `json:"marketplaceSettings,omitempty"`
}

// NewSponsoredProductsGlobalProductIdentifiers instantiates a new SponsoredProductsGlobalProductIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalProductIdentifiers() *SponsoredProductsGlobalProductIdentifiers {
	this := SponsoredProductsGlobalProductIdentifiers{}
	return &this
}

// NewSponsoredProductsGlobalProductIdentifiersWithDefaults instantiates a new SponsoredProductsGlobalProductIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalProductIdentifiersWithDefaults() *SponsoredProductsGlobalProductIdentifiers {
	this := SponsoredProductsGlobalProductIdentifiers{}
	return &this
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductIdentifiers) GetMarketplaceSettings() []SponsoredProductsMarketplaceLevelProductIdentifier {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsMarketplaceLevelProductIdentifier
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductIdentifiers) GetMarketplaceSettingsOk() ([]SponsoredProductsMarketplaceLevelProductIdentifier, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductIdentifiers) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsMarketplaceLevelProductIdentifier and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsGlobalProductIdentifiers) SetMarketplaceSettings(v []SponsoredProductsMarketplaceLevelProductIdentifier) {
	o.MarketplaceSettings = v
}

func (o SponsoredProductsGlobalProductIdentifiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalProductIdentifiers struct {
	value *SponsoredProductsGlobalProductIdentifiers
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductIdentifiers) Get() *SponsoredProductsGlobalProductIdentifiers {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductIdentifiers) Set(val *SponsoredProductsGlobalProductIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductIdentifiers(val *SponsoredProductsGlobalProductIdentifiers) *NullableSponsoredProductsGlobalProductIdentifiers {
	return &NullableSponsoredProductsGlobalProductIdentifiers{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductIdentifiers) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
