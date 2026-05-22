package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListMatchedAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListMatchedAdGroupsResponseContent{}

// SponsoredProductsListMatchedAdGroupsResponseContent Response object for querying adGroups with ASIN/SKU matches.
type SponsoredProductsListMatchedAdGroupsResponseContent struct {
	// The total number of results available.
	TotalResults *int64 `json:"totalResults,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the     request. If the nextToken field is empty, there are no further results.
	NextToken       *string                           `json:"nextToken,omitempty"`
	MatchedAdGroups []SponsoredProductsMatchedAdGroup `json:"matchedAdGroups,omitempty"`
}

// NewSponsoredProductsListMatchedAdGroupsResponseContent instantiates a new SponsoredProductsListMatchedAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListMatchedAdGroupsResponseContent() *SponsoredProductsListMatchedAdGroupsResponseContent {
	this := SponsoredProductsListMatchedAdGroupsResponseContent{}
	return &this
}

// NewSponsoredProductsListMatchedAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsListMatchedAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListMatchedAdGroupsResponseContentWithDefaults() *SponsoredProductsListMatchedAdGroupsResponseContent {
	this := SponsoredProductsListMatchedAdGroupsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetMatchedAdGroups returns the MatchedAdGroups field value if set, zero value otherwise.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) GetMatchedAdGroups() []SponsoredProductsMatchedAdGroup {
	if o == nil || IsNil(o.MatchedAdGroups) {
		var ret []SponsoredProductsMatchedAdGroup
		return ret
	}
	return o.MatchedAdGroups
}

// GetMatchedAdGroupsOk returns a tuple with the MatchedAdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) GetMatchedAdGroupsOk() ([]SponsoredProductsMatchedAdGroup, bool) {
	if o == nil || IsNil(o.MatchedAdGroups) {
		return nil, false
	}
	return o.MatchedAdGroups, true
}

// HasMatchedAdGroups returns a boolean if a field has been set.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) HasMatchedAdGroups() bool {
	if o != nil && !IsNil(o.MatchedAdGroups) {
		return true
	}

	return false
}

// SetMatchedAdGroups gets a reference to the given []SponsoredProductsMatchedAdGroup and assigns it to the MatchedAdGroups field.
func (o *SponsoredProductsListMatchedAdGroupsResponseContent) SetMatchedAdGroups(v []SponsoredProductsMatchedAdGroup) {
	o.MatchedAdGroups = v
}

func (o SponsoredProductsListMatchedAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.MatchedAdGroups) {
		toSerialize["matchedAdGroups"] = o.MatchedAdGroups
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListMatchedAdGroupsResponseContent struct {
	value *SponsoredProductsListMatchedAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListMatchedAdGroupsResponseContent) Get() *SponsoredProductsListMatchedAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListMatchedAdGroupsResponseContent) Set(val *SponsoredProductsListMatchedAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListMatchedAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListMatchedAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListMatchedAdGroupsResponseContent(val *SponsoredProductsListMatchedAdGroupsResponseContent) *NullableSponsoredProductsListMatchedAdGroupsResponseContent {
	return &NullableSponsoredProductsListMatchedAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListMatchedAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListMatchedAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
