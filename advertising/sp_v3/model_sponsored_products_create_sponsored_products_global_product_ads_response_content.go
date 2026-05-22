package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent struct {
	ProductAds SponsoredProductsBulkGlobalProductAdOperationResponse `json:"productAds"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent(productAds SponsoredProductsBulkGlobalProductAdOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) GetProductAds() SponsoredProductsBulkGlobalProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalProductAdOperationResponse
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) GetProductAdsOk() (*SponsoredProductsBulkGlobalProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) SetProductAds(v SponsoredProductsBulkGlobalProductAdOperationResponse) {
	o.ProductAds = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
