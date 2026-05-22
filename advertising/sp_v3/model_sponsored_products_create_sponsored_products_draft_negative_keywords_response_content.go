package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkDraftNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkDraftNegativeKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkDraftNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkDraftNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkDraftNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
