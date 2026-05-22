package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent struct {
	// An array of keywords.
	Keywords []SponsoredProductsCreateGlobalKeyword `json:"keywords"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent

// NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent(keywords []SponsoredProductsCreateGlobalKeyword) *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) GetKeywords() []SponsoredProductsCreateGlobalKeyword {
	if o == nil {
		var ret []SponsoredProductsCreateGlobalKeyword
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) GetKeywordsOk() ([]SponsoredProductsCreateGlobalKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) SetKeywords(v []SponsoredProductsCreateGlobalKeyword) {
	o.Keywords = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
