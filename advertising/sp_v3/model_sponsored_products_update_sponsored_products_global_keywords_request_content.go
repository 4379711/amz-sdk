package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent struct {
	// An array of keywords with updated values.
	Keywords []SponsoredProductsUpdateGlobalKeyword `json:"keywords"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent(keywords []SponsoredProductsUpdateGlobalKeyword) *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) GetKeywords() []SponsoredProductsUpdateGlobalKeyword {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalKeyword
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) GetKeywordsOk() ([]SponsoredProductsUpdateGlobalKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) SetKeywords(v []SponsoredProductsUpdateGlobalKeyword) {
	o.Keywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
