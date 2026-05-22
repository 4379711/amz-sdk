package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListTargetPromotionGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListTargetPromotionGroupsResponseContent{}

// SponsoredProductsListTargetPromotionGroupsResponseContent Response object for querying target promotion groups.
type SponsoredProductsListTargetPromotionGroupsResponseContent struct {
	// The total number of results available.
	TotalResults *int64 `json:"totalResults,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the     request. If the nextToken field is empty, there are no further results.
	NextToken             *string                                 `json:"nextToken,omitempty"`
	TargetPromotionGroups []SponsoredProductsTargetPromotionGroup `json:"targetPromotionGroups,omitempty"`
}

// NewSponsoredProductsListTargetPromotionGroupsResponseContent instantiates a new SponsoredProductsListTargetPromotionGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListTargetPromotionGroupsResponseContent() *SponsoredProductsListTargetPromotionGroupsResponseContent {
	this := SponsoredProductsListTargetPromotionGroupsResponseContent{}
	return &this
}

// NewSponsoredProductsListTargetPromotionGroupsResponseContentWithDefaults instantiates a new SponsoredProductsListTargetPromotionGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListTargetPromotionGroupsResponseContentWithDefaults() *SponsoredProductsListTargetPromotionGroupsResponseContent {
	this := SponsoredProductsListTargetPromotionGroupsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetPromotionGroups returns the TargetPromotionGroups field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) GetTargetPromotionGroups() []SponsoredProductsTargetPromotionGroup {
	if o == nil || IsNil(o.TargetPromotionGroups) {
		var ret []SponsoredProductsTargetPromotionGroup
		return ret
	}
	return o.TargetPromotionGroups
}

// GetTargetPromotionGroupsOk returns a tuple with the TargetPromotionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) GetTargetPromotionGroupsOk() ([]SponsoredProductsTargetPromotionGroup, bool) {
	if o == nil || IsNil(o.TargetPromotionGroups) {
		return nil, false
	}
	return o.TargetPromotionGroups, true
}

// HasTargetPromotionGroups returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) HasTargetPromotionGroups() bool {
	if o != nil && !IsNil(o.TargetPromotionGroups) {
		return true
	}

	return false
}

// SetTargetPromotionGroups gets a reference to the given []SponsoredProductsTargetPromotionGroup and assigns it to the TargetPromotionGroups field.
func (o *SponsoredProductsListTargetPromotionGroupsResponseContent) SetTargetPromotionGroups(v []SponsoredProductsTargetPromotionGroup) {
	o.TargetPromotionGroups = v
}

func (o SponsoredProductsListTargetPromotionGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TargetPromotionGroups) {
		toSerialize["targetPromotionGroups"] = o.TargetPromotionGroups
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListTargetPromotionGroupsResponseContent struct {
	value *SponsoredProductsListTargetPromotionGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListTargetPromotionGroupsResponseContent) Get() *SponsoredProductsListTargetPromotionGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListTargetPromotionGroupsResponseContent) Set(val *SponsoredProductsListTargetPromotionGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListTargetPromotionGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListTargetPromotionGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListTargetPromotionGroupsResponseContent(val *SponsoredProductsListTargetPromotionGroupsResponseContent) *NullableSponsoredProductsListTargetPromotionGroupsResponseContent {
	return &NullableSponsoredProductsListTargetPromotionGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListTargetPromotionGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListTargetPromotionGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
