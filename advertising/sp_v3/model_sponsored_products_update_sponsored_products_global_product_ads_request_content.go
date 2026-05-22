package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent struct {
	// An array of ads with updated values.
	ProductAds []SponsoredProductsUpdateGlobalProductAd `json:"productAds"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent(productAds []SponsoredProductsUpdateGlobalProductAd) *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) GetProductAds() []SponsoredProductsUpdateGlobalProductAd {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalProductAd
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) GetProductAdsOk() ([]SponsoredProductsUpdateGlobalProductAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) SetProductAds(v []SponsoredProductsUpdateGlobalProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
