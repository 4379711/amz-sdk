package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent struct {
	KeywordIdFilter SponsoredProductsObjectIdFilter `json:"keywordIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent(keywordIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent{}
	this.KeywordIdFilter = keywordIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent{}
	return &this
}

// GetKeywordIdFilter returns the KeywordIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) GetKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.KeywordIdFilter
}

// GetKeywordIdFilterOk returns a tuple with the KeywordIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) GetKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordIdFilter, true
}

// SetKeywordIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) SetKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.KeywordIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordIdFilter"] = o.KeywordIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
