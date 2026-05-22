package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent struct {
	NegativeKeywords SponsoredProductsBulkGlobalNegativeKeywordOperationResponse `json:"negativeKeywords"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent(negativeKeywords SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent{}
	this.NegativeKeywords = negativeKeywords
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) GetNegativeKeywords() SponsoredProductsBulkGlobalNegativeKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalNegativeKeywordOperationResponse
		return ret
	}

	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) GetNegativeKeywordsOk() (*SponsoredProductsBulkGlobalNegativeKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeKeywords, true
}

// SetNegativeKeywords sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) SetNegativeKeywords(v SponsoredProductsBulkGlobalNegativeKeywordOperationResponse) {
	o.NegativeKeywords = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeKeywords"] = o.NegativeKeywords
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
