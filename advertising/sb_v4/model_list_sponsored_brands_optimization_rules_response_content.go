package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsOptimizationRulesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsOptimizationRulesResponseContent{}

// ListSponsoredBrandsOptimizationRulesResponseContent struct for ListSponsoredBrandsOptimizationRulesResponseContent
type ListSponsoredBrandsOptimizationRulesResponseContent struct {
	// Token value allowing to navigate to the next response page.
	NextToken *string `json:"nextToken,omitempty"`
	// The total number of entities.
	TotalCount        *float32           `json:"totalCount,omitempty"`
	OptimizationRules []OptimizationRule `json:"optimizationRules"`
}

type _ListSponsoredBrandsOptimizationRulesResponseContent ListSponsoredBrandsOptimizationRulesResponseContent

// NewListSponsoredBrandsOptimizationRulesResponseContent instantiates a new ListSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsOptimizationRulesResponseContent(optimizationRules []OptimizationRule) *ListSponsoredBrandsOptimizationRulesResponseContent {
	this := ListSponsoredBrandsOptimizationRulesResponseContent{}
	this.OptimizationRules = optimizationRules
	return &this
}

// NewListSponsoredBrandsOptimizationRulesResponseContentWithDefaults instantiates a new ListSponsoredBrandsOptimizationRulesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsOptimizationRulesResponseContentWithDefaults() *ListSponsoredBrandsOptimizationRulesResponseContent {
	this := ListSponsoredBrandsOptimizationRulesResponseContent{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) GetTotalCount() float32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) GetTotalCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) SetTotalCount(v float32) {
	o.TotalCount = &v
}

// GetOptimizationRules returns the OptimizationRules field value
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRules() []OptimizationRule {
	if o == nil {
		var ret []OptimizationRule
		return ret
	}

	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) GetOptimizationRulesOk() ([]OptimizationRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptimizationRules, true
}

// SetOptimizationRules sets field value
func (o *ListSponsoredBrandsOptimizationRulesResponseContent) SetOptimizationRules(v []OptimizationRule) {
	o.OptimizationRules = v
}

func (o ListSponsoredBrandsOptimizationRulesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	toSerialize["optimizationRules"] = o.OptimizationRules
	return toSerialize, nil
}

type NullableListSponsoredBrandsOptimizationRulesResponseContent struct {
	value *ListSponsoredBrandsOptimizationRulesResponseContent
	isSet bool
}

func (v NullableListSponsoredBrandsOptimizationRulesResponseContent) Get() *ListSponsoredBrandsOptimizationRulesResponseContent {
	return v.value
}

func (v *NullableListSponsoredBrandsOptimizationRulesResponseContent) Set(val *ListSponsoredBrandsOptimizationRulesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsOptimizationRulesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsOptimizationRulesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsOptimizationRulesResponseContent(val *ListSponsoredBrandsOptimizationRulesResponseContent) *NullableListSponsoredBrandsOptimizationRulesResponseContent {
	return &NullableListSponsoredBrandsOptimizationRulesResponseContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsOptimizationRulesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsOptimizationRulesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
