package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkGlobalKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent(keywords SponsoredProductsBulkGlobalKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) GetKeywords() SponsoredProductsBulkGlobalKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkGlobalKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkGlobalKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
