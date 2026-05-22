package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent struct {
	// An array of draft keywords.
	Keywords []SponsoredProductsCreateDraftKeyword `json:"keywords"`
}

type _SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent

// NewSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent(keywords []SponsoredProductsCreateDraftKeyword) *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) GetKeywords() []SponsoredProductsCreateDraftKeyword {
	if o == nil {
		var ret []SponsoredProductsCreateDraftKeyword
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) GetKeywordsOk() ([]SponsoredProductsCreateDraftKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) SetKeywords(v []SponsoredProductsCreateDraftKeyword) {
	o.Keywords = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
