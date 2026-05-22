package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent struct for SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent
type SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent struct {
	// An array of ads with updated values.
	ProductAds []SponsoredProductsUpdateProductAd `json:"productAds"`
}

type _SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent(productAds []SponsoredProductsUpdateProductAd) *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsProductAdsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) GetProductAds() []SponsoredProductsUpdateProductAd {
	if o == nil {
		var ret []SponsoredProductsUpdateProductAd
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) GetProductAdsOk() ([]SponsoredProductsUpdateProductAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) SetProductAds(v []SponsoredProductsUpdateProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent(val *SponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
