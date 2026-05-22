package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent(keywords SponsoredProductsBulkKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) GetKeywords() SponsoredProductsBulkKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
