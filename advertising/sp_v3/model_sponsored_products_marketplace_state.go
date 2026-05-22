package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceState{}

// SponsoredProductsMarketplaceState struct for SponsoredProductsMarketplaceState
type SponsoredProductsMarketplaceState struct {
	Marketplace SponsoredProductsMarketplace `json:"marketplace"`
	State       SponsoredProductsEntityState `json:"state"`
}

type _SponsoredProductsMarketplaceState SponsoredProductsMarketplaceState

// NewSponsoredProductsMarketplaceState instantiates a new SponsoredProductsMarketplaceState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceState(marketplace SponsoredProductsMarketplace, state SponsoredProductsEntityState) *SponsoredProductsMarketplaceState {
	this := SponsoredProductsMarketplaceState{}
	this.Marketplace = marketplace
	this.State = state
	return &this
}

// NewSponsoredProductsMarketplaceStateWithDefaults instantiates a new SponsoredProductsMarketplaceState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceStateWithDefaults() *SponsoredProductsMarketplaceState {
	this := SponsoredProductsMarketplaceState{}
	return &this
}

// GetMarketplace returns the Marketplace field value
func (o *SponsoredProductsMarketplaceState) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil {
		var ret SponsoredProductsMarketplace
		return ret
	}

	return o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceState) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Marketplace, true
}

// SetMarketplace sets field value
func (o *SponsoredProductsMarketplaceState) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = v
}

// GetState returns the State field value
func (o *SponsoredProductsMarketplaceState) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceState) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsMarketplaceState) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

func (o SponsoredProductsMarketplaceState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplace"] = o.Marketplace
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableSponsoredProductsMarketplaceState struct {
	value *SponsoredProductsMarketplaceState
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceState) Get() *SponsoredProductsMarketplaceState {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceState) Set(val *SponsoredProductsMarketplaceState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceState(val *SponsoredProductsMarketplaceState) *NullableSponsoredProductsMarketplaceState {
	return &NullableSponsoredProductsMarketplaceState{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
