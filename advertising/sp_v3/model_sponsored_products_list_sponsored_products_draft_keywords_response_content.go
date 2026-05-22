package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent{}

// SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent struct for SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent
type SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent struct {
	// The total number of entities
	TotalResults *int64                          `json:"totalResults,omitempty"`
	Keywords     []SponsoredProductsDraftKeyword `json:"keywords,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent() *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftKeywordsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) GetKeywords() []SponsoredProductsDraftKeyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []SponsoredProductsDraftKeyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) GetKeywordsOk() ([]SponsoredProductsDraftKeyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []SponsoredProductsDraftKeyword and assigns it to the Keywords field.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) SetKeywords(v []SponsoredProductsDraftKeyword) {
	o.Keywords = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent(val *SponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
