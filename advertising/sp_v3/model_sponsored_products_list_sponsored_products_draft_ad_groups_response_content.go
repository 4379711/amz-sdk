package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent{}

// SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent struct for SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent
type SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent struct {
	// The total number of entities
	TotalResults *int64                          `json:"totalResults,omitempty"`
	AdGroups     []SponsoredProductsDraftAdGroup `json:"adGroups,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent() *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) GetAdGroups() []SponsoredProductsDraftAdGroup {
	if o == nil || IsNil(o.AdGroups) {
		var ret []SponsoredProductsDraftAdGroup
		return ret
	}
	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) GetAdGroupsOk() ([]SponsoredProductsDraftAdGroup, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given []SponsoredProductsDraftAdGroup and assigns it to the AdGroups field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) SetAdGroups(v []SponsoredProductsDraftAdGroup) {
	o.AdGroups = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent(val *SponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
