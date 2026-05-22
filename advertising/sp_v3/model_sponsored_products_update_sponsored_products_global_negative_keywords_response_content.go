package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkGlobalNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkGlobalNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkGlobalNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
