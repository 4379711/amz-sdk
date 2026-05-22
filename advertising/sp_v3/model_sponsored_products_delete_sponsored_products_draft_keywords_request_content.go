package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent struct {
	KeywordIdFilter SponsoredProductsObjectIdFilter `json:"keywordIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent(keywordIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent{}
	this.KeywordIdFilter = keywordIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent{}
	return &this
}

// GetKeywordIdFilter returns the KeywordIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) GetKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.KeywordIdFilter
}

// GetKeywordIdFilterOk returns a tuple with the KeywordIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) GetKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordIdFilter, true
}

// SetKeywordIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) SetKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.KeywordIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordIdFilter"] = o.KeywordIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
