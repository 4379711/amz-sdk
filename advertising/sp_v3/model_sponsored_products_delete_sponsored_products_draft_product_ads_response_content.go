package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent struct {
	ProductAds SponsoredProductsBulkDraftProductAdOperationResponse `json:"productAds"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent(productAds SponsoredProductsBulkDraftProductAdOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) GetProductAds() SponsoredProductsBulkDraftProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftProductAdOperationResponse
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) GetProductAdsOk() (*SponsoredProductsBulkDraftProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) SetProductAds(v SponsoredProductsBulkDraftProductAdOperationResponse) {
	o.ProductAds = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
