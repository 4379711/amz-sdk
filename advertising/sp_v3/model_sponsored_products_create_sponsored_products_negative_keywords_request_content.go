package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent struct {
	// An array of negativeKeywords.
	NegativeKeywords []SponsoredProductsCreateNegativeKeyword `json:"negativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent

// NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent(negativeKeywords []SponsoredProductsCreateNegativeKeyword) *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) GetNegativeKeywords() []SponsoredProductsCreateNegativeKeyword {
	if o == nil {
		var ret []SponsoredProductsCreateNegativeKeyword
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) GetNegativeKeywordsOk() ([]SponsoredProductsCreateNegativeKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) SetNegativeKeywords(v []SponsoredProductsCreateNegativeKeyword) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
