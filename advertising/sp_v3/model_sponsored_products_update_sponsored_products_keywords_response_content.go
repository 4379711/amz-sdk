package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent struct for SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent
type SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent struct {
	Keywords SponsoredProductsBulkKeywordOperationResponse `json:"keywords"`
}

type _SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent(keywords SponsoredProductsBulkKeywordOperationResponse) *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent{}
	this.Keywords = keywords
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsKeywordsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent{}
	return &this
}

// GetKeywords returns the Keywords field value
func (o *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) GetKeywords() SponsoredProductsBulkKeywordOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkKeywordOperationResponse
		return ret
	}

	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) GetKeywordsOk() (*SponsoredProductsBulkKeywordOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value
func (o *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) SetKeywords(v SponsoredProductsBulkKeywordOperationResponse) {
	o.Keywords = v
}

func (o SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywords"] = o.Keywords
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent(val *SponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
