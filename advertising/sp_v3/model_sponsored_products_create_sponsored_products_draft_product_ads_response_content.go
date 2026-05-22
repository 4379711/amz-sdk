package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent{}

// SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent struct for SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent
type SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent struct {
	DraftProductAds SponsoredProductsBulkDraftProductAdOperationResponse `json:"draftProductAds"`
}

type _SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent

// NewSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent(draftProductAds SponsoredProductsBulkDraftProductAdOperationResponse) *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent{}
	this.DraftProductAds = draftProductAds
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent{}
	return &this
}

// GetDraftProductAds returns the DraftProductAds field value
func (o *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) GetDraftProductAds() SponsoredProductsBulkDraftProductAdOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftProductAdOperationResponse
		return ret
	}

	return o.DraftProductAds
}

// GetDraftProductAdsOk returns a tuple with the DraftProductAds field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) GetDraftProductAdsOk() (*SponsoredProductsBulkDraftProductAdOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftProductAds, true
}

// SetDraftProductAds sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) SetDraftProductAds(v SponsoredProductsBulkDraftProductAdOperationResponse) {
	o.DraftProductAds = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["draftProductAds"] = o.DraftProductAds
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent(val *SponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
