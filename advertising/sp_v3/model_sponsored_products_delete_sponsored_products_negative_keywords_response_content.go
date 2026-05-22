package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkNegativeKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
