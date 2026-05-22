package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsKeywordsRequestContent struct {
	// An array of keywords.
	Keywords []SponsoredProductsCreateKeyword `json:"keywords"`
}

type _SponsoredProductsCreateSponsoredProductsKeywordsRequestContent SponsoredProductsCreateSponsoredProductsKeywordsRequestContent

// NewSponsoredProductsCreateSponsoredProductsKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsKeywordsRequestContent(keywords []SponsoredProductsCreateKeyword) *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsKeywordsRequestContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsKeywordsRequestContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent) GetKeywords() []SponsoredProductsCreateKeyword {
	if o == nil {
		var ret []SponsoredProductsCreateKeyword
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent) GetKeywordsOk() ([]SponsoredProductsCreateKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent) SetKeywords(v []SponsoredProductsCreateKeyword) {
	o.Keywords = v
}

func (o SponsoredProductsCreateSponsoredProductsKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
