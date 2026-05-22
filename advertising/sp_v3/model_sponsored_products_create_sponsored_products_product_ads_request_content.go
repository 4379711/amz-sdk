package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsProductAdsRequestContent{}

// SponsoredProductsCreateSponsoredProductsProductAdsRequestContent struct for SponsoredProductsCreateSponsoredProductsProductAdsRequestContent
type SponsoredProductsCreateSponsoredProductsProductAdsRequestContent struct {
	// An array of ads.
	ProductAds []SponsoredProductsCreateProductAd `json:"productAds"`
}

type _SponsoredProductsCreateSponsoredProductsProductAdsRequestContent SponsoredProductsCreateSponsoredProductsProductAdsRequestContent

// NewSponsoredProductsCreateSponsoredProductsProductAdsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsProductAdsRequestContent(productAds []SponsoredProductsCreateProductAd) *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsProductAdsRequestContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsProductAdsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsProductAdsRequestContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent) GetProductAds() []SponsoredProductsCreateProductAd {
	if o == nil {
		var ret []SponsoredProductsCreateProductAd
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent) GetProductAdsOk() ([]SponsoredProductsCreateProductAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent) SetProductAds(v []SponsoredProductsCreateProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsCreateSponsoredProductsProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent(val *SponsoredProductsCreateSponsoredProductsProductAdsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
