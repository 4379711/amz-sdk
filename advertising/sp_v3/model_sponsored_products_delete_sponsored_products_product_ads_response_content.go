package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent struct for SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent
type SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent struct {
	ProductAds SponsoredProductsBulkProductAdOperationResponse `json:"productAds"`
}

type _SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent(productAds SponsoredProductsBulkProductAdOperationResponse) *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsProductAdsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) GetProductAds() SponsoredProductsBulkProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkProductAdOperationResponse
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) GetProductAdsOk() (*SponsoredProductsBulkProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) SetProductAds(v SponsoredProductsBulkProductAdOperationResponse) {
	o.ProductAds = v
}

func (o SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent(val *SponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
