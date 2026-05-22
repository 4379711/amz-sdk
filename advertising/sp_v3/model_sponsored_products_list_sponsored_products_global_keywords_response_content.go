package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent{}

// SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent struct for SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent
type SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent struct {
	// The total number of entities
	TotalResults *int64                           `json:"totalResults,omitempty"`
	Keywords     []SponsoredProductsGlobalKeyword `json:"keywords,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent instantiates a new SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent() *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) GetKeywords() []SponsoredProductsGlobalKeyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []SponsoredProductsGlobalKeyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) GetKeywordsOk() ([]SponsoredProductsGlobalKeyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []SponsoredProductsGlobalKeyword and assigns it to the Keywords field.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) SetKeywords(v []SponsoredProductsGlobalKeyword) {
	o.Keywords = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) Get() *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) Set(val *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent(val *SponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalKeywordsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
