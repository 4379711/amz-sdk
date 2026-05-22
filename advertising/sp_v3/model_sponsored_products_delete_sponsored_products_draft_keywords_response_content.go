package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkDraftKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent(keywords SponsoredProductsBulkDraftKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) GetKeywords() SponsoredProductsBulkDraftKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkDraftKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkDraftKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
