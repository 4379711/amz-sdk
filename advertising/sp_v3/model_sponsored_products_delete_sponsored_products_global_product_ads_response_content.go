package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent struct {
	ProductAds SponsoredProductsBulkGlobalProductAdOperationResponse `json:"productAds"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent(productAds SponsoredProductsBulkGlobalProductAdOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) GetProductAds() SponsoredProductsBulkGlobalProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalProductAdOperationResponse
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) GetProductAdsOk() (*SponsoredProductsBulkGlobalProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) SetProductAds(v SponsoredProductsBulkGlobalProductAdOperationResponse) {
	o.ProductAds = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
