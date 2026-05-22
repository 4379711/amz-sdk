package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent struct {
	ProductAds SponsoredProductsBulkGlobalProductAdOperationResponse `json:"productAds"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent(productAds SponsoredProductsBulkGlobalProductAdOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) GetProductAds() SponsoredProductsBulkGlobalProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalProductAdOperationResponse
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) GetProductAdsOk() (*SponsoredProductsBulkGlobalProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) SetProductAds(v SponsoredProductsBulkGlobalProductAdOperationResponse) {
	o.ProductAds = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
