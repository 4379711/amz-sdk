package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent struct {
	NegativeKeywordIdFilter SponsoredProductsObjectIdFilter `json:"negativeKeywordIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent(negativeKeywordIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent{}
	this.NegativeKeywordIdFilter = negativeKeywordIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywordIdFilter returns the NegativeKeywordIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) GetNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.NegativeKeywordIdFilter
}

// GetNegativeKeywordIdFilterOk returns a tuple with the NegativeKeywordIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) GetNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywordIdFilter, true
}

// SetNegativeKeywordIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) SetNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeKeywordIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywordIdFilter"] = o.NegativeKeywordIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
