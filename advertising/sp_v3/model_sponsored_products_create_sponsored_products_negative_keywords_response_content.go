package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkNegativeKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
