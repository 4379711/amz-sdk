package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent{}

// SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent struct for SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent
type SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent struct {
	// The total number of entities
	TotalResults *int64                     `json:"totalResults,omitempty"`
	Keywords     []SponsoredProductsKeyword `json:"keywords,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent instantiates a new SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent() *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent {
	this := SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent {
	this := SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) GetKeywords() []SponsoredProductsKeyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []SponsoredProductsKeyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) GetKeywordsOk() ([]SponsoredProductsKeyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []SponsoredProductsKeyword and assigns it to the Keywords field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) SetKeywords(v []SponsoredProductsKeyword) {
	o.Keywords = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent struct {
	value *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) Get() *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) Set(val *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent(val *SponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsKeywordsPreviewResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
