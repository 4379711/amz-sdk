package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent{}

// SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent struct for SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent
type SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent struct {
	NegativeKeywords []SponsoredProductsDraftNegativeKeyword `json:"negativeKeywords,omitempty"`
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent() *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) GetNegativeKeywords() []SponsoredProductsDraftNegativeKeyword {
	if o == nil || IsNil(o.NegativeKeywords) {
		var ret []SponsoredProductsDraftNegativeKeyword
		return ret
	}
	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) GetNegativeKeywordsOk() ([]SponsoredProductsDraftNegativeKeyword, bool) {
	if o == nil || IsNil(o.NegativeKeywords) {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// HasNegativeKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) HasNegativeKeywords() bool {
	if o != nil && !IsNil(o.NegativeKeywords) {
		return true
	}

	return false
}

// SetNegativeKeywords gets a reference to the given []SponsoredProductsDraftNegativeKeyword and assigns it to the NegativeKeywords field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) SetNegativeKeywords(v []SponsoredProductsDraftNegativeKeyword) {
	o.NegativeKeywords = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegativeKeywords) {
		toSerialize["negativeKeywords"] = o.NegativeKeywords
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent(val *SponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
