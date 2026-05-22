package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent struct for SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent
type SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent struct {
	ProductAds SponsoredProductsBulkProductAdOperationResponse `json:"productAds"`
}

type _SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent(productAds SponsoredProductsBulkProductAdOperationResponse) *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsProductAdsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) GetProductAds() SponsoredProductsBulkProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkProductAdOperationResponse
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) GetProductAdsOk() (*SponsoredProductsBulkProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) SetProductAds(v SponsoredProductsBulkProductAdOperationResponse) {
	o.ProductAds = v
}

func (o SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent(val *SponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
