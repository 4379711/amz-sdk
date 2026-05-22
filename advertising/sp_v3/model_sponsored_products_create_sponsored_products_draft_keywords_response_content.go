package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkDraftKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent(keywords SponsoredProductsBulkDraftKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) GetKeywords() SponsoredProductsBulkDraftKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkDraftKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkDraftKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
