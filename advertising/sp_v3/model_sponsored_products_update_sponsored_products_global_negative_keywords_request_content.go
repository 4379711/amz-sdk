package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	// An array of negativeKeywords with updated values.
	NegativeKeywords []SponsoredProductsUpdateGlobalNegativeKeyword `json:"negativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent(negativeKeywords []SponsoredProductsUpdateGlobalNegativeKeyword) *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywords() []SponsoredProductsUpdateGlobalNegativeKeyword {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalNegativeKeyword
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywordsOk() ([]SponsoredProductsUpdateGlobalNegativeKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) SetNegativeKeywords(v []SponsoredProductsUpdateGlobalNegativeKeyword) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
