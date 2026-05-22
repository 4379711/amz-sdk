package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalTargetingExpressionPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalTargetingExpressionPredicate{}

// SponsoredProductsGlobalTargetingExpressionPredicate struct for SponsoredProductsGlobalTargetingExpressionPredicate
type SponsoredProductsGlobalTargetingExpressionPredicate struct {
	// The marketplace settings for expression to be overridden for marketplace.
	MarketplaceSettings []SponsoredProductsTargetingExpressionPredicateMarketValue `json:"marketplaceSettings,omitempty"`
	Type                *SponsoredProductsTargetingExpressionPredicateType         `json:"type,omitempty"`
}

// NewSponsoredProductsGlobalTargetingExpressionPredicate instantiates a new SponsoredProductsGlobalTargetingExpressionPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalTargetingExpressionPredicate() *SponsoredProductsGlobalTargetingExpressionPredicate {
	this := SponsoredProductsGlobalTargetingExpressionPredicate{}
	return &this
}

// NewSponsoredProductsGlobalTargetingExpressionPredicateWithDefaults instantiates a new SponsoredProductsGlobalTargetingExpressionPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalTargetingExpressionPredicateWithDefaults() *SponsoredProductsGlobalTargetingExpressionPredicate {
	this := SponsoredProductsGlobalTargetingExpressionPredicate{}
	return &this
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) GetMarketplaceSettings() []SponsoredProductsTargetingExpressionPredicateMarketValue {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsTargetingExpressionPredicateMarketValue
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) GetMarketplaceSettingsOk() ([]SponsoredProductsTargetingExpressionPredicateMarketValue, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsTargetingExpressionPredicateMarketValue and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) SetMarketplaceSettings(v []SponsoredProductsTargetingExpressionPredicateMarketValue) {
	o.MarketplaceSettings = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) GetType() SponsoredProductsTargetingExpressionPredicateType {
	if o == nil || IsNil(o.Type) {
		var ret SponsoredProductsTargetingExpressionPredicateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) GetTypeOk() (*SponsoredProductsTargetingExpressionPredicateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SponsoredProductsTargetingExpressionPredicateType and assigns it to the Type field.
func (o *SponsoredProductsGlobalTargetingExpressionPredicate) SetType(v SponsoredProductsTargetingExpressionPredicateType) {
	o.Type = &v
}

func (o SponsoredProductsGlobalTargetingExpressionPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalTargetingExpressionPredicate struct {
	value *SponsoredProductsGlobalTargetingExpressionPredicate
	isSet bool
}

func (v NullableSponsoredProductsGlobalTargetingExpressionPredicate) Get() *SponsoredProductsGlobalTargetingExpressionPredicate {
	return v.value
}

func (v *NullableSponsoredProductsGlobalTargetingExpressionPredicate) Set(val *SponsoredProductsGlobalTargetingExpressionPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalTargetingExpressionPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalTargetingExpressionPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalTargetingExpressionPredicate(val *SponsoredProductsGlobalTargetingExpressionPredicate) *NullableSponsoredProductsGlobalTargetingExpressionPredicate {
	return &NullableSponsoredProductsGlobalTargetingExpressionPredicate{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalTargetingExpressionPredicate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalTargetingExpressionPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
