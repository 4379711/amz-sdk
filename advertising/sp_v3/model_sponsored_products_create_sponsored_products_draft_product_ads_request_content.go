package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent{}

// SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent struct for SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent
type SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent struct {
	// An array of ads.
	ProductAds []SponsoredProductsCreateDraftProductAd `json:"productAds,omitempty"`
}

// NewSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent() *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent{}
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent{}
	return &this
}

// GetProductAds returns the ProductAds field value if set, zero value otherwise.
func (o *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) GetProductAds() []SponsoredProductsCreateDraftProductAd {
	if o == nil || IsNil(o.ProductAds) {
		var ret []SponsoredProductsCreateDraftProductAd
		return ret
	}
	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) GetProductAdsOk() ([]SponsoredProductsCreateDraftProductAd, bool) {
	if o == nil || IsNil(o.ProductAds) {
		return nil, false
	}
	return o.ProductAds, true
}

// HasProductAds returns a boolean if a field has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) HasProductAds() bool {
	if o != nil && !IsNil(o.ProductAds) {
		return true
	}

	return false
}

// SetProductAds gets a reference to the given []SponsoredProductsCreateDraftProductAd and assigns it to the ProductAds field.
func (o *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) SetProductAds(v []SponsoredProductsCreateDraftProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductAds) {
		toSerialize["productAds"] = o.ProductAds
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent(val *SponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
