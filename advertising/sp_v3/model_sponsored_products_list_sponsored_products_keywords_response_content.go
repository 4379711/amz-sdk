package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsKeywordsResponseContent{}

// SponsoredProductsListSponsoredProductsKeywordsResponseContent struct for SponsoredProductsListSponsoredProductsKeywordsResponseContent
type SponsoredProductsListSponsoredProductsKeywordsResponseContent struct {
	// The total number of entities
	TotalResults *int64                     `json:"totalResults,omitempty"`
	Keywords     []SponsoredProductsKeyword `json:"keywords,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsKeywordsResponseContent instantiates a new SponsoredProductsListSponsoredProductsKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsKeywordsResponseContent() *SponsoredProductsListSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsKeywordsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsKeywordsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsKeywordsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) GetKeywords() []SponsoredProductsKeyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []SponsoredProductsKeyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) GetKeywordsOk() ([]SponsoredProductsKeyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []SponsoredProductsKeyword and assigns it to the Keywords field.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) SetKeywords(v []SponsoredProductsKeyword) {
	o.Keywords = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsKeywordsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent) Get() *SponsoredProductsListSponsoredProductsKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent) Set(val *SponsoredProductsListSponsoredProductsKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsKeywordsResponseContent(val *SponsoredProductsListSponsoredProductsKeywordsResponseContent) *NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
