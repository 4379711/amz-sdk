package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsCreateSponsoredProductsKeywordsResponseContent SponsoredProductsCreateSponsoredProductsKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsKeywordsResponseContent(keywords SponsoredProductsBulkKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent) GetKeywords() SponsoredProductsBulkKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsCreateSponsoredProductsKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
