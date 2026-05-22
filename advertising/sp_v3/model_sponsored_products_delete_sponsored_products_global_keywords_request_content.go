package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent struct {
	KeywordIdFilter *SponsoredProductsObjectIdFilter `json:"keywordIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent{}
	return &this
}

// GetKeywordIdFilter returns the KeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) GetKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.KeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.KeywordIdFilter
}

// GetKeywordIdFilterOk returns a tuple with the KeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) GetKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.KeywordIdFilter) {
		return nil, false
	}
	return o.KeywordIdFilter, true
}

// HasKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) HasKeywordIdFilter() bool {
	if o != nil && !IsNil(o.KeywordIdFilter) {
		return true
	}

	return false
}

// SetKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the KeywordIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) SetKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.KeywordIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordIdFilter) {
		toSerialize["keywordIdFilter"] = o.KeywordIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
