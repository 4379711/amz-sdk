package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsAdGroupsResponseContent{}

// SponsoredProductsListSponsoredProductsAdGroupsResponseContent struct for SponsoredProductsListSponsoredProductsAdGroupsResponseContent
type SponsoredProductsListSponsoredProductsAdGroupsResponseContent struct {
	// The total number of entities
	TotalResults *int64                     `json:"totalResults,omitempty"`
	AdGroups     []SponsoredProductsAdGroup `json:"adGroups,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsAdGroupsResponseContent instantiates a new SponsoredProductsListSponsoredProductsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsAdGroupsResponseContent() *SponsoredProductsListSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsListSponsoredProductsAdGroupsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsAdGroupsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsListSponsoredProductsAdGroupsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) GetAdGroups() []SponsoredProductsAdGroup {
	if o == nil || IsNil(o.AdGroups) {
		var ret []SponsoredProductsAdGroup
		return ret
	}
	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) GetAdGroupsOk() ([]SponsoredProductsAdGroup, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given []SponsoredProductsAdGroup and assigns it to the AdGroups field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) SetAdGroups(v []SponsoredProductsAdGroup) {
	o.AdGroups = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent) Get() *SponsoredProductsListSponsoredProductsAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent) Set(val *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent(val *SponsoredProductsListSponsoredProductsAdGroupsResponseContent) *NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
