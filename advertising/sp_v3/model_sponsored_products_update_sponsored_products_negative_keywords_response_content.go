package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent struct for SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent
type SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkNegativeKeywordOperationResponse) *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent(val *SponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
