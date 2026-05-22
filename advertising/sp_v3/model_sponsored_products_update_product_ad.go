package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateProductAd{}

// SponsoredProductsUpdateProductAd struct for SponsoredProductsUpdateProductAd
type SponsoredProductsUpdateProductAd struct {
	// The product ad identifier.
	AdId  string                                      `json:"adId"`
	State *SponsoredProductsCreateOrUpdateEntityState `json:"state,omitempty"`
}

type _SponsoredProductsUpdateProductAd SponsoredProductsUpdateProductAd

// NewSponsoredProductsUpdateProductAd instantiates a new SponsoredProductsUpdateProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateProductAd(adId string) *SponsoredProductsUpdateProductAd {
	this := SponsoredProductsUpdateProductAd{}
	this.AdId = adId
	return &this
}

// NewSponsoredProductsUpdateProductAdWithDefaults instantiates a new SponsoredProductsUpdateProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateProductAdWithDefaults() *SponsoredProductsUpdateProductAd {
	this := SponsoredProductsUpdateProductAd{}
	return &this
}

// GetAdId returns the AdId field value
func (o *SponsoredProductsUpdateProductAd) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateProductAd) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *SponsoredProductsUpdateProductAd) SetAdId(v string) {
	o.AdId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateProductAd) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateProductAd) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateProductAd) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateProductAd) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = &v
}

func (o SponsoredProductsUpdateProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateProductAd struct {
	value *SponsoredProductsUpdateProductAd
	isSet bool
}

func (v NullableSponsoredProductsUpdateProductAd) Get() *SponsoredProductsUpdateProductAd {
	return v.value
}

func (v *NullableSponsoredProductsUpdateProductAd) Set(val *SponsoredProductsUpdateProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateProductAd(val *SponsoredProductsUpdateProductAd) *NullableSponsoredProductsUpdateProductAd {
	return &NullableSponsoredProductsUpdateProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
