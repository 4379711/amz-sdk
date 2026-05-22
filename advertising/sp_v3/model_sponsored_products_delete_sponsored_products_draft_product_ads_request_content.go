package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent struct {
	AdIdFilter SponsoredProductsObjectIdFilter `json:"adIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent(adIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent{}
	this.AdIdFilter = adIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent{}
	return &this
}

// GetAdIdFilter returns the AdIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) GetAdIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) GetAdIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdIdFilter, true
}

// SetAdIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) SetAdIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adIdFilter"] = o.AdIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
