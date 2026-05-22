package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent struct {
	// An array of negativeKeywords.
	NegativeKeywords []SponsoredProductsCreateDraftNegativeKeyword `json:"negativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent(negativeKeywords []SponsoredProductsCreateDraftNegativeKeyword) *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywords() []SponsoredProductsCreateDraftNegativeKeyword {
	if o == nil {
		var ret []SponsoredProductsCreateDraftNegativeKeyword
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) GetNegativeKeywordsOk() ([]SponsoredProductsCreateDraftNegativeKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) SetNegativeKeywords(v []SponsoredProductsCreateDraftNegativeKeyword) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
