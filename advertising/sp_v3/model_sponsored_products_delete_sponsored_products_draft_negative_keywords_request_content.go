package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent struct {
	NegativeKeywordIdFilter SponsoredProductsObjectIdFilter `json:"negativeKeywordIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent(negativeKeywordIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent{}
	this.NegativeKeywordIdFilter = negativeKeywordIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywordIdFilter returns the NegativeKeywordIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.NegativeKeywordIdFilter
}

// GetNegativeKeywordIdFilterOk returns a tuple with the NegativeKeywordIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywordIdFilter, true
}

// SetNegativeKeywordIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) SetNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeKeywordIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywordIdFilter"] = o.NegativeKeywordIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
