package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent struct for SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent
type SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent struct {
	// An array of negativeKeywords with updated values.
	NegativeKeywords []SponsoredProductsUpdateNegativeKeyword `json:"negativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent(negativeKeywords []SponsoredProductsUpdateNegativeKeyword) *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) GetNegativeKeywords() []SponsoredProductsUpdateNegativeKeyword {
	if o == nil {
		var ret []SponsoredProductsUpdateNegativeKeyword
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) GetNegativeKeywordsOk() ([]SponsoredProductsUpdateNegativeKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) SetNegativeKeywords(v []SponsoredProductsUpdateNegativeKeyword) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent(val *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
