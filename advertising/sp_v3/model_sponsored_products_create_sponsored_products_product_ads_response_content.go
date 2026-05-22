package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateSponsoredProductsProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsProductAdsResponseContent{}

// SponsoredProductsCreateSponsoredProductsProductAdsResponseContent struct for SponsoredProductsCreateSponsoredProductsProductAdsResponseContent
type SponsoredProductsCreateSponsoredProductsProductAdsResponseContent struct {
	ProductAds SponsoredProductsBulkProductAdOperationResponse `json:"productAds"`
}

type _SponsoredProductsCreateSponsoredProductsProductAdsResponseContent SponsoredProductsCreateSponsoredProductsProductAdsResponseContent

// NewSponsoredProductsCreateSponsoredProductsProductAdsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsProductAdsResponseContent(productAds SponsoredProductsBulkProductAdOperationResponse) *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsProductAdsResponseContent{}
	this.ProductAds = productAds
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsProductAdsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsProductAdsResponseContent{}
	return &this
}

// GetProductAds returns the ProductAds field value
func (o *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent) GetProductAds() SponsoredProductsBulkProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkProductAdOperationResponse
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent) GetProductAdsOk() (*SponsoredProductsBulkProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductAds, true
}

// SetProductAds sets field value
func (o *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent) SetProductAds(v SponsoredProductsBulkProductAdOperationResponse) {
	o.ProductAds = v
}

func (o SponsoredProductsCreateSponsoredProductsProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productAds"] = o.ProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent(val *SponsoredProductsCreateSponsoredProductsProductAdsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
