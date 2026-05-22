package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent struct {
	AdIdFilter *SponsoredProductsObjectIdFilter `json:"adIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent{}
	return &this
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) GetAdIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) GetAdIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) SetAdIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdIdFilter) {
		toSerialize["adIdFilter"] = o.AdIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
