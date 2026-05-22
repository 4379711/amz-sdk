package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateOrUpdateGlobalEntityState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateOrUpdateGlobalEntityState{}

// SponsoredProductsCreateOrUpdateGlobalEntityState struct for SponsoredProductsCreateOrUpdateGlobalEntityState
type SponsoredProductsCreateOrUpdateGlobalEntityState struct {
	MarketplaceSettings []SponsoredProductsCreateOrUpdateMarketplaceState `json:"marketplaceSettings,omitempty"`
	State               SponsoredProductsCreateOrUpdateEntityState        `json:"state"`
}

type _SponsoredProductsCreateOrUpdateGlobalEntityState SponsoredProductsCreateOrUpdateGlobalEntityState

// NewSponsoredProductsCreateOrUpdateGlobalEntityState instantiates a new SponsoredProductsCreateOrUpdateGlobalEntityState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateOrUpdateGlobalEntityState(state SponsoredProductsCreateOrUpdateEntityState) *SponsoredProductsCreateOrUpdateGlobalEntityState {
	this := SponsoredProductsCreateOrUpdateGlobalEntityState{}
	this.State = state
	return &this
}

// NewSponsoredProductsCreateOrUpdateGlobalEntityStateWithDefaults instantiates a new SponsoredProductsCreateOrUpdateGlobalEntityState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateOrUpdateGlobalEntityStateWithDefaults() *SponsoredProductsCreateOrUpdateGlobalEntityState {
	this := SponsoredProductsCreateOrUpdateGlobalEntityState{}
	return &this
}

// GetMarketplaceSettings returns the MarketplaceSettings field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateGlobalEntityState) GetMarketplaceSettings() []SponsoredProductsCreateOrUpdateMarketplaceState {
	if o == nil || IsNil(o.MarketplaceSettings) {
		var ret []SponsoredProductsCreateOrUpdateMarketplaceState
		return ret
	}
	return o.MarketplaceSettings
}

// GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateGlobalEntityState) GetMarketplaceSettingsOk() ([]SponsoredProductsCreateOrUpdateMarketplaceState, bool) {
	if o == nil || IsNil(o.MarketplaceSettings) {
		return nil, false
	}
	return o.MarketplaceSettings, true
}

// HasMarketplaceSettings returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateGlobalEntityState) HasMarketplaceSettings() bool {
	if o != nil && !IsNil(o.MarketplaceSettings) {
		return true
	}

	return false
}

// SetMarketplaceSettings gets a reference to the given []SponsoredProductsCreateOrUpdateMarketplaceState and assigns it to the MarketplaceSettings field.
func (o *SponsoredProductsCreateOrUpdateGlobalEntityState) SetMarketplaceSettings(v []SponsoredProductsCreateOrUpdateMarketplaceState) {
	o.MarketplaceSettings = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateOrUpdateGlobalEntityState) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateGlobalEntityState) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateOrUpdateGlobalEntityState) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

func (o SponsoredProductsCreateOrUpdateGlobalEntityState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceSettings) {
		toSerialize["marketplaceSettings"] = o.MarketplaceSettings
	}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableSponsoredProductsCreateOrUpdateGlobalEntityState struct {
	value *SponsoredProductsCreateOrUpdateGlobalEntityState
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateGlobalEntityState) Get() *SponsoredProductsCreateOrUpdateGlobalEntityState {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateGlobalEntityState) Set(val *SponsoredProductsCreateOrUpdateGlobalEntityState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateGlobalEntityState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateGlobalEntityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateGlobalEntityState(val *SponsoredProductsCreateOrUpdateGlobalEntityState) *NullableSponsoredProductsCreateOrUpdateGlobalEntityState {
	return &NullableSponsoredProductsCreateOrUpdateGlobalEntityState{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateGlobalEntityState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateGlobalEntityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
