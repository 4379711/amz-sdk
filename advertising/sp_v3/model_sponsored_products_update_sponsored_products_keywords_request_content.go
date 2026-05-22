package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent struct for SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent
type SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent struct {
	// An array of keywords with updated values.
	Keywords []SponsoredProductsUpdateKeyword `json:"keywords"`
}

type _SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent(keywords []SponsoredProductsUpdateKeyword) *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsKeywordsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) GetKeywords() []SponsoredProductsUpdateKeyword {
	if o == nil {
		var ret []SponsoredProductsUpdateKeyword
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) GetKeywordsOk() ([]SponsoredProductsUpdateKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) SetKeywords(v []SponsoredProductsUpdateKeyword) {
	o.Keywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent(val *SponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
