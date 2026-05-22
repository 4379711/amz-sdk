package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent struct {
	// An array of ads.
	ProductAds []SponsoredProductsCreateGlobalProductAd `json:"productAds,omitempty"`
}

// NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent() *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent{}
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent{}
	return &this
}

// GetProductAds returns the ProductAds field value if set, zero value otherwise.
func (o *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) GetProductAds() []SponsoredProductsCreateGlobalProductAd {
	if o == nil || IsNil(o.ProductAds) {
		var ret []SponsoredProductsCreateGlobalProductAd
		return ret
	}
	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) GetProductAdsOk() ([]SponsoredProductsCreateGlobalProductAd, bool) {
	if o == nil || IsNil(o.ProductAds) {
		return nil, false
	}
	return o.ProductAds, true
}

// HasProductAds returns a boolean if a field has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) HasProductAds() bool {
	if o != nil && !IsNil(o.ProductAds) {
		return true
	}

	return false
}

// SetProductAds gets a reference to the given []SponsoredProductsCreateGlobalProductAd and assigns it to the ProductAds field.
func (o *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) SetProductAds(v []SponsoredProductsCreateGlobalProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductAds) {
		toSerialize["productAds"] = o.ProductAds
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
