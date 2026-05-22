package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkDraftNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkDraftNegativeKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkDraftNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkDraftNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkDraftNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
