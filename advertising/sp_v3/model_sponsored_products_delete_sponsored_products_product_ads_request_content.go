package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent struct for SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent
type SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent struct {
	AdIdFilter SponsoredProductsObjectIdFilter `json:"adIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent(adIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent{}
	this.AdIdFilter = adIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsProductAdsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent{}
	return &this
}

// GetAdIdFilter returns the AdIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) GetAdIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) GetAdIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdIdFilter, true
}

// SetAdIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) SetAdIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adIdFilter"] = o.AdIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent(val *SponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
