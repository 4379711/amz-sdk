package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	// An array of negativeKeywords.
	NegativeKeywords []SponsoredProductsCreateGlobalNegativeKeyword `json:"negativeKeywords,omitempty"`
}

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent() *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value if set, zero value otherwise.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywords() []SponsoredProductsCreateGlobalNegativeKeyword {
	if o == nil || IsNil(o.NegativeKeywords) {
		var ret []SponsoredProductsCreateGlobalNegativeKeyword
		return ret
	}
	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) GetNegativeKeywordsOk() ([]SponsoredProductsCreateGlobalNegativeKeyword, bool) {
	if o == nil || IsNil(o.NegativeKeywords) {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// HasNegativeKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) HasNegativeKeywords() bool {
	if o != nil && !IsNil(o.NegativeKeywords) {
		return true
	}

	return false
}

// SetNegativeKeywords gets a reference to the given []SponsoredProductsCreateGlobalNegativeKeyword and assigns it to the NegativeKeywords field.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) SetNegativeKeywords(v []SponsoredProductsCreateGlobalNegativeKeyword) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegativeKeywords) {
		toSerialize["negativeKeywords"] = o.NegativeKeywords
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
