package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkGlobalKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent(keywords SponsoredProductsBulkGlobalKeywordOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) GetKeywords() SponsoredProductsBulkGlobalKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkGlobalKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkGlobalKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
