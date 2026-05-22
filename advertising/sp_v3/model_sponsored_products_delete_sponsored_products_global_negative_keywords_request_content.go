package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	NegativeKeywordIdFilter *SponsoredProductsObjectIdFilter `json:"negativeKeywordIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywordIdFilter returns the NegativeKeywordIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywordIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.NegativeKeywordIdFilter
}

// GetNegativeKeywordIdFilterOk returns a tuple with the NegativeKeywordIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywordIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.NegativeKeywordIdFilter) {
		return nil, false
	}
	return o.NegativeKeywordIdFilter, true
}

// HasNegativeKeywordIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) HasNegativeKeywordIdFilter() bool {
	if o != nil && !IsNil(o.NegativeKeywordIdFilter) {
		return true
	}

	return false
}

// SetNegativeKeywordIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the NegativeKeywordIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) SetNegativeKeywordIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeKeywordIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegativeKeywordIdFilter) {
		toSerialize["negativeKeywordIdFilter"] = o.NegativeKeywordIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
