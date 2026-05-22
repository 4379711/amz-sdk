package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsAdGroupsResponseContent{}

// ListSponsoredBrandsAdGroupsResponseContent struct for ListSponsoredBrandsAdGroupsResponseContent
type ListSponsoredBrandsAdGroupsResponseContent struct {
	// The total number of entities.
	TotalResults *float32  `json:"totalResults,omitempty"`
	AdGroups     []AdGroup `json:"adGroups,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewListSponsoredBrandsAdGroupsResponseContent instantiates a new ListSponsoredBrandsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsAdGroupsResponseContent() *ListSponsoredBrandsAdGroupsResponseContent {
	this := ListSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// NewListSponsoredBrandsAdGroupsResponseContentWithDefaults instantiates a new ListSponsoredBrandsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsAdGroupsResponseContentWithDefaults() *ListSponsoredBrandsAdGroupsResponseContent {
	this := ListSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsResponseContent) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsResponseContent) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *ListSponsoredBrandsAdGroupsResponseContent) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsResponseContent) GetAdGroups() []AdGroup {
	if o == nil || IsNil(o.AdGroups) {
		var ret []AdGroup
		return ret
	}
	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsResponseContent) GetAdGroupsOk() ([]AdGroup, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given []AdGroup and assigns it to the AdGroups field.
func (o *ListSponsoredBrandsAdGroupsResponseContent) SetAdGroups(v []AdGroup) {
	o.AdGroups = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdGroupsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdGroupsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdGroupsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsAdGroupsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListSponsoredBrandsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableListSponsoredBrandsAdGroupsResponseContent struct {
	value *ListSponsoredBrandsAdGroupsResponseContent
	isSet bool
}

func (v NullableListSponsoredBrandsAdGroupsResponseContent) Get() *ListSponsoredBrandsAdGroupsResponseContent {
	return v.value
}

func (v *NullableListSponsoredBrandsAdGroupsResponseContent) Set(val *ListSponsoredBrandsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsAdGroupsResponseContent(val *ListSponsoredBrandsAdGroupsResponseContent) *NullableListSponsoredBrandsAdGroupsResponseContent {
	return &NullableListSponsoredBrandsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
