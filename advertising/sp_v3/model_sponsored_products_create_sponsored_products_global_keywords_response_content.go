package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkGlobalKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent(keywords SponsoredProductsBulkGlobalKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) GetKeywords() SponsoredProductsBulkGlobalKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkGlobalKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkGlobalKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
