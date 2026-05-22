package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPISearchOptimizationRuleAssociationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISearchOptimizationRuleAssociationsRequest{}

// OptimizationRulesAPISearchOptimizationRuleAssociationsRequest Request body for getting optimization rule associations.
type OptimizationRulesAPISearchOptimizationRuleAssociationsRequest struct {
	// Sets a limit on the number of results returned. Defaults to 50. Maximum limit of maxResults is 1000.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the nextToken field is empty, there are no further results.
	NextToken              *string                                `json:"nextToken,omitempty"`
	OptimizationRuleFilter *OptimizationRulesAPIEntityFieldFilter `json:"optimizationRuleFilter,omitempty"`
}

// NewOptimizationRulesAPISearchOptimizationRuleAssociationsRequest instantiates a new OptimizationRulesAPISearchOptimizationRuleAssociationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISearchOptimizationRuleAssociationsRequest() *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest {
	this := OptimizationRulesAPISearchOptimizationRuleAssociationsRequest{}
	return &this
}

// NewOptimizationRulesAPISearchOptimizationRuleAssociationsRequestWithDefaults instantiates a new OptimizationRulesAPISearchOptimizationRuleAssociationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISearchOptimizationRuleAssociationsRequestWithDefaults() *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest {
	this := OptimizationRulesAPISearchOptimizationRuleAssociationsRequest{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) SetNextToken(v string) {
	o.NextToken = &v
}

// GetOptimizationRuleFilter returns the OptimizationRuleFilter field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) GetOptimizationRuleFilter() OptimizationRulesAPIEntityFieldFilter {
	if o == nil || IsNil(o.OptimizationRuleFilter) {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}
	return *o.OptimizationRuleFilter
}

// GetOptimizationRuleFilterOk returns a tuple with the OptimizationRuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) GetOptimizationRuleFilterOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil || IsNil(o.OptimizationRuleFilter) {
		return nil, false
	}
	return o.OptimizationRuleFilter, true
}

// HasOptimizationRuleFilter returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) HasOptimizationRuleFilter() bool {
	if o != nil && !IsNil(o.OptimizationRuleFilter) {
		return true
	}

	return false
}

// SetOptimizationRuleFilter gets a reference to the given OptimizationRulesAPIEntityFieldFilter and assigns it to the OptimizationRuleFilter field.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) SetOptimizationRuleFilter(v OptimizationRulesAPIEntityFieldFilter) {
	o.OptimizationRuleFilter = &v
}

func (o OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.OptimizationRuleFilter) {
		toSerialize["optimizationRuleFilter"] = o.OptimizationRuleFilter
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest struct {
	value *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest
	isSet bool
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest) Get() *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest) Set(val *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest(val *OptimizationRulesAPISearchOptimizationRuleAssociationsRequest) *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest {
	return &NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
