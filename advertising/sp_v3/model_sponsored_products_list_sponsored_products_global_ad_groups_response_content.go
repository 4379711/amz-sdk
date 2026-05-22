package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent{}

// SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent struct for SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent
type SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent struct {
	// The total number of entities
	TotalResults *int64                           `json:"totalResults,omitempty"`
	AdGroups     []SponsoredProductsGlobalAdGroup `json:"adGroups,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent instantiates a new SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent() *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroups() []SponsoredProductsGlobalAdGroup {
	if o == nil || IsNil(o.AdGroups) {
		var ret []SponsoredProductsGlobalAdGroup
		return ret
	}
	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroupsOk() ([]SponsoredProductsGlobalAdGroup, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given []SponsoredProductsGlobalAdGroup and assigns it to the AdGroups field.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) SetAdGroups(v []SponsoredProductsGlobalAdGroup) {
	o.AdGroups = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) Get() *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) Set(val *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent(val *SponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
