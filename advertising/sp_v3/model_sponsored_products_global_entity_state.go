package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalEntityState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalEntityState{}

// SponsoredProductsGlobalEntityState struct for SponsoredProductsGlobalEntityState
type SponsoredProductsGlobalEntityState struct {
	MarketplaceSettings []SponsoredProductsMarketplaceState `json:"marketplaceSettings,omitempty"`
	State               SponsoredProductsEntityState        `json:"state"`
}

type _SponsoredProductsGlobalEntityState SponsoredProductsGlobalEntityState

// NewSponsoredProductsGlobalEntityState instantiates a new SponsoredProductsGlobalEntityState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalEntityState(state SponsoredProductsEntityState) *SponsoredProductsGlobalEntityState {
	this := SponsoredProductsGlobalEntityState{}
	this.State = state
	return &this
}

// NewSponsoredProductsGlobalEntityStateWithDefaults instantiates a new SponsoredProductsGlobalEntityState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalEntityStateWithDefaults() *SponsoredProductsGlobalEntityState {
	this := SponsoredProductsGlobalEntityState{}
	return &this
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalEntityState) GetMarketplaceSettings() []SponsoredProductsMarketplaceState {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsMarketplaceState
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalEntityState) GetMarketplaceSettingsOk() ([]SponsoredProductsMarketplaceState, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalEntityState) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsMarketplaceState and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsGlobalEntityState) SetMarketplaceSettings(v []SponsoredProductsMarketplaceState) {
	o.MarketplaceSettings = v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalEntityState) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalEntityState) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalEntityState) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

func (o SponsoredProductsGlobalEntityState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalEntityState struct {
	value *SponsoredProductsGlobalEntityState
	isSet bool
}

func (v NullableSponsoredProductsGlobalEntityState) Get() *SponsoredProductsGlobalEntityState {
	return v.value
}

func (v *NullableSponsoredProductsGlobalEntityState) Set(val *SponsoredProductsGlobalEntityState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalEntityState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalEntityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalEntityState(val *SponsoredProductsGlobalEntityState) *NullableSponsoredProductsGlobalEntityState {
	return &NullableSponsoredProductsGlobalEntityState{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalEntityState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalEntityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
