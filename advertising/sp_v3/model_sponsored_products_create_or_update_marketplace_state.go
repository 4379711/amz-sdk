package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateOrUpdateMarketplaceState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateOrUpdateMarketplaceState{}

// SponsoredProductsCreateOrUpdateMarketplaceState struct for SponsoredProductsCreateOrUpdateMarketplaceState
type SponsoredProductsCreateOrUpdateMarketplaceState struct {
	Marketplace SponsoredProductsMarketplace               `json:"marketplace"`
	State       SponsoredProductsCreateOrUpdateEntityState `json:"state"`
}

type _SponsoredProductsCreateOrUpdateMarketplaceState SponsoredProductsCreateOrUpdateMarketplaceState

// NewSponsoredProductsCreateOrUpdateMarketplaceState instantiates a new SponsoredProductsCreateOrUpdateMarketplaceState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateOrUpdateMarketplaceState(marketplace SponsoredProductsMarketplace, state SponsoredProductsCreateOrUpdateEntityState) *SponsoredProductsCreateOrUpdateMarketplaceState {
	this := SponsoredProductsCreateOrUpdateMarketplaceState{}
	this.Marketplace = marketplace
	this.State = state
	return &this
}

// NewSponsoredProductsCreateOrUpdateMarketplaceStateWithDefaults instantiates a new SponsoredProductsCreateOrUpdateMarketplaceState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateOrUpdateMarketplaceStateWithDefaults() *SponsoredProductsCreateOrUpdateMarketplaceState {
	this := SponsoredProductsCreateOrUpdateMarketplaceState{}
	return &this
}

// GetMarketplace returns the Marketplace field value
func (o *SponsoredProductsCreateOrUpdateMarketplaceState) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil {
		var ret SponsoredProductsMarketplace
		return ret
	}

	return o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateMarketplaceState) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Marketplace, true
}

// SetMarketplace sets field value
func (o *SponsoredProductsCreateOrUpdateMarketplaceState) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateOrUpdateMarketplaceState) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateMarketplaceState) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateOrUpdateMarketplaceState) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

func (o SponsoredProductsCreateOrUpdateMarketplaceState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplace"] = o.Marketplace
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableSponsoredProductsCreateOrUpdateMarketplaceState struct {
	value *SponsoredProductsCreateOrUpdateMarketplaceState
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateMarketplaceState) Get() *SponsoredProductsCreateOrUpdateMarketplaceState {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateMarketplaceState) Set(val *SponsoredProductsCreateOrUpdateMarketplaceState) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateMarketplaceState) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateMarketplaceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateMarketplaceState(val *SponsoredProductsCreateOrUpdateMarketplaceState) *NullableSponsoredProductsCreateOrUpdateMarketplaceState {
	return &NullableSponsoredProductsCreateOrUpdateMarketplaceState{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateMarketplaceState) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateMarketplaceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
