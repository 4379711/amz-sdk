package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsOptimizationRulesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsOptimizationRulesRequestContent{}

// ListSponsoredBrandsOptimizationRulesRequestContent struct for ListSponsoredBrandsOptimizationRulesRequestContent
type ListSponsoredBrandsOptimizationRulesRequestContent struct {
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API.
	MaxResults *float32 `json:"maxResults,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken                *string                   `json:"nextToken,omitempty"`
	OptimizationRuleIdFilter *OptimizationRuleIdFilter `json:"optimizationRuleIdFilter,omitempty"`
}

// NewListSponsoredBrandsOptimizationRulesRequestContent instantiates a new ListSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsOptimizationRulesRequestContent() *ListSponsoredBrandsOptimizationRulesRequestContent {
	this := ListSponsoredBrandsOptimizationRulesRequestContent{}
	return &this
}

// NewListSponsoredBrandsOptimizationRulesRequestContentWithDefaults instantiates a new ListSponsoredBrandsOptimizationRulesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsOptimizationRulesRequestContentWithDefaults() *ListSponsoredBrandsOptimizationRulesRequestContent {
	this := ListSponsoredBrandsOptimizationRulesRequestContent{}
	return &this
}

// GetEntityFilter returns the EntityFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetEntityFilter() EntityFilter {
	if o == nil || IsNil(o.EntityFilter) {
		var ret EntityFilter
		return ret
	}
	return *o.EntityFilter
}

// GetEntityFilterOk returns a tuple with the EntityFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetEntityFilterOk() (*EntityFilter, bool) {
	if o == nil || IsNil(o.EntityFilter) {
		return nil, false
	}
	return o.EntityFilter, true
}

// HasEntityFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) HasEntityFilter() bool {
	if o != nil && !IsNil(o.EntityFilter) {
		return true
	}

	return false
}

// SetEntityFilter gets a reference to the given EntityFilter and assigns it to the EntityFilter field.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) SetEntityFilter(v EntityFilter) {
	o.EntityFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetOptimizationRuleIdFilter returns the OptimizationRuleIdFilter field value if set, zero value otherwise.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRuleIdFilter() OptimizationRuleIdFilter {
	if o == nil || IsNil(o.OptimizationRuleIdFilter) {
		var ret OptimizationRuleIdFilter
		return ret
	}
	return *o.OptimizationRuleIdFilter
}

// GetOptimizationRuleIdFilterOk returns a tuple with the OptimizationRuleIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) GetOptimizationRuleIdFilterOk() (*OptimizationRuleIdFilter, bool) {
	if o == nil || IsNil(o.OptimizationRuleIdFilter) {
		return nil, false
	}
	return o.OptimizationRuleIdFilter, true
}

// HasOptimizationRuleIdFilter returns a boolean if a field has been set.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) HasOptimizationRuleIdFilter() bool {
	if o != nil && !IsNil(o.OptimizationRuleIdFilter) {
		return true
	}

	return false
}

// SetOptimizationRuleIdFilter gets a reference to the given OptimizationRuleIdFilter and assigns it to the OptimizationRuleIdFilter field.
func (o *ListSponsoredBrandsOptimizationRulesRequestContent) SetOptimizationRuleIdFilter(v OptimizationRuleIdFilter) {
	o.OptimizationRuleIdFilter = &v
}

func (o ListSponsoredBrandsOptimizationRulesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityFilter) {
		toSerialize["entityFilter"] = o.EntityFilter
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.OptimizationRuleIdFilter) {
		toSerialize["optimizationRuleIdFilter"] = o.OptimizationRuleIdFilter
	}
	return toSerialize, nil
}

type NullableListSponsoredBrandsOptimizationRulesRequestContent struct {
	value *ListSponsoredBrandsOptimizationRulesRequestContent
	isSet bool
}

func (v NullableListSponsoredBrandsOptimizationRulesRequestContent) Get() *ListSponsoredBrandsOptimizationRulesRequestContent {
	return v.value
}

func (v *NullableListSponsoredBrandsOptimizationRulesRequestContent) Set(val *ListSponsoredBrandsOptimizationRulesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsOptimizationRulesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsOptimizationRulesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsOptimizationRulesRequestContent(val *ListSponsoredBrandsOptimizationRulesRequestContent) *NullableListSponsoredBrandsOptimizationRulesRequestContent {
	return &NullableListSponsoredBrandsOptimizationRulesRequestContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsOptimizationRulesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsOptimizationRulesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
