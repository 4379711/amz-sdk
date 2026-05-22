package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent{}

// SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent struct for SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent
type SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent struct {
	NegativeKeywords []SponsoredProductsNegativeKeyword `json:"negativeKeywords,omitempty"`
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent instantiates a new SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent() *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent{}
	return &this
}

// GetNegativeKeywords returns the NegativeKeywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywords() []SponsoredProductsNegativeKeyword {
	if o == nil || IsNil(o.NegativeKeywords) {
		var ret []SponsoredProductsNegativeKeyword
		return ret
	}
	return o.NegativeKeywords
}

// GetNegativeKeywordsOk returns a tuple with the NegativeKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) GetNegativeKeywordsOk() ([]SponsoredProductsNegativeKeyword, bool) {
	if o == nil || IsNil(o.NegativeKeywords) {
		return nil, false
	}
	return o.NegativeKeywords, true
}

// HasNegativeKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) HasNegativeKeywords() bool {
	if o != nil && !IsNil(o.NegativeKeywords) {
		return true
	}

	return false
}

// SetNegativeKeywords gets a reference to the given []SponsoredProductsNegativeKeyword and assigns it to the NegativeKeywords field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) SetNegativeKeywords(v []SponsoredProductsNegativeKeyword) {
	o.NegativeKeywords = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) Get() *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) Set(val *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent(val *SponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
