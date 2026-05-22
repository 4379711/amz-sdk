package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkGlobalNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkGlobalNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkGlobalNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
