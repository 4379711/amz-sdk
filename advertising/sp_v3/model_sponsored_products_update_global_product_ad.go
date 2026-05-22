package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateGlobalProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateGlobalProductAd{}

// SponsoredProductsUpdateGlobalProductAd struct for SponsoredProductsUpdateGlobalProductAd
type SponsoredProductsUpdateGlobalProductAd struct {
	// The product ad identifier.
	AdId string `json:"adId"`
	// Name for the product Ad
	Name  *string                                           `json:"name,omitempty"`
	Asin  *SponsoredProductsGlobalProductIdentifiers        `json:"asin,omitempty"`
	State *SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state,omitempty"`
	Sku   *SponsoredProductsGlobalProductIdentifiers        `json:"sku,omitempty"`
}

type _SponsoredProductsUpdateGlobalProductAd SponsoredProductsUpdateGlobalProductAd

// NewSponsoredProductsUpdateGlobalProductAd instantiates a new SponsoredProductsUpdateGlobalProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateGlobalProductAd(adId string) *SponsoredProductsUpdateGlobalProductAd {
	this := SponsoredProductsUpdateGlobalProductAd{}
	this.AdId = adId
	return &this
}

// NewSponsoredProductsUpdateGlobalProductAdWithDefaults instantiates a new SponsoredProductsUpdateGlobalProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateGlobalProductAdWithDefaults() *SponsoredProductsUpdateGlobalProductAd {
	this := SponsoredProductsUpdateGlobalProductAd{}
	return &this
}

// GetAdId returns the AdId field value
func (o *SponsoredProductsUpdateGlobalProductAd) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *SponsoredProductsUpdateGlobalProductAd) SetAdId(v string) {
	o.AdId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalProductAd) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsUpdateGlobalProductAd) SetName(v string) {
	o.Name = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalProductAd) GetAsin() SponsoredProductsGlobalProductIdentifiers {
	if o == nil || IsNil(o.Asin) {
		var ret SponsoredProductsGlobalProductIdentifiers
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) GetAsinOk() (*SponsoredProductsGlobalProductIdentifiers, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given SponsoredProductsGlobalProductIdentifiers and assigns it to the Asin field.
func (o *SponsoredProductsUpdateGlobalProductAd) SetAsin(v SponsoredProductsGlobalProductIdentifiers) {
	o.Asin = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalProductAd) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil || IsNil(o.State) {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given SponsoredProductsCreateOrUpdateGlobalEntityState and assigns it to the State field.
func (o *SponsoredProductsUpdateGlobalProductAd) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SponsoredProductsUpdateGlobalProductAd) GetSku() SponsoredProductsGlobalProductIdentifiers {
	if o == nil || IsNil(o.Sku) {
		var ret SponsoredProductsGlobalProductIdentifiers
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) GetSkuOk() (*SponsoredProductsGlobalProductIdentifiers, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SponsoredProductsUpdateGlobalProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given SponsoredProductsGlobalProductIdentifiers and assigns it to the Sku field.
func (o *SponsoredProductsUpdateGlobalProductAd) SetSku(v SponsoredProductsGlobalProductIdentifiers) {
	o.Sku = &v
}

func (o SponsoredProductsUpdateGlobalProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateGlobalProductAd struct {
	value *SponsoredProductsUpdateGlobalProductAd
	isSet bool
}

func (v NullableSponsoredProductsUpdateGlobalProductAd) Get() *SponsoredProductsUpdateGlobalProductAd {
	return v.value
}

func (v *NullableSponsoredProductsUpdateGlobalProductAd) Set(val *SponsoredProductsUpdateGlobalProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateGlobalProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateGlobalProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateGlobalProductAd(val *SponsoredProductsUpdateGlobalProductAd) *NullableSponsoredProductsUpdateGlobalProductAd {
	return &NullableSponsoredProductsUpdateGlobalProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateGlobalProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateGlobalProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
