package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIGetRuleRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIGetRuleRecommendationsRequest{}

// OptimizationRulesAPIGetRuleRecommendationsRequest Request object for getting rule recommendation. All the filters in the request are used in conjunction with one another.
type OptimizationRulesAPIGetRuleRecommendationsRequest struct {
	RuleSubCategoryFilter  OptimizationRulesAPIEntityFieldFilter `json:"ruleSubCategoryFilter"`
	RecommendationIdFilter *string                               `json:"recommendationIdFilter,omitempty"`
	CampaignIdListFilter   []OptimizationRulesAPICampaignFilter  `json:"campaignIdListFilter,omitempty"`
	MaxResults             *float32                              `json:"maxResults,omitempty"`
	NextToken              *string                               `json:"nextToken,omitempty"`
	RuleCategoryFilter     OptimizationRulesAPIEntityFieldFilter `json:"ruleCategoryFilter"`
}

type _OptimizationRulesAPIGetRuleRecommendationsRequest OptimizationRulesAPIGetRuleRecommendationsRequest

// NewOptimizationRulesAPIGetRuleRecommendationsRequest instantiates a new OptimizationRulesAPIGetRuleRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIGetRuleRecommendationsRequest(ruleSubCategoryFilter OptimizationRulesAPIEntityFieldFilter, ruleCategoryFilter OptimizationRulesAPIEntityFieldFilter) *OptimizationRulesAPIGetRuleRecommendationsRequest {
	this := OptimizationRulesAPIGetRuleRecommendationsRequest{}
	this.RuleSubCategoryFilter = ruleSubCategoryFilter
	this.RuleCategoryFilter = ruleCategoryFilter
	return &this
}

// NewOptimizationRulesAPIGetRuleRecommendationsRequestWithDefaults instantiates a new OptimizationRulesAPIGetRuleRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIGetRuleRecommendationsRequestWithDefaults() *OptimizationRulesAPIGetRuleRecommendationsRequest {
	this := OptimizationRulesAPIGetRuleRecommendationsRequest{}
	return &this
}

// GetRuleSubCategoryFilter returns the RuleSubCategoryFilter field value
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetRuleSubCategoryFilter() OptimizationRulesAPIEntityFieldFilter {
	if o == nil {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}

	return o.RuleSubCategoryFilter
}

// GetRuleSubCategoryFilterOk returns a tuple with the RuleSubCategoryFilter field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetRuleSubCategoryFilterOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleSubCategoryFilter, true
}

// SetRuleSubCategoryFilter sets field value
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) SetRuleSubCategoryFilter(v OptimizationRulesAPIEntityFieldFilter) {
	o.RuleSubCategoryFilter = v
}

// GetRecommendationIdFilter returns the RecommendationIdFilter field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetRecommendationIdFilter() string {
	if o == nil || IsNil(o.RecommendationIdFilter) {
		var ret string
		return ret
	}
	return *o.RecommendationIdFilter
}

// GetRecommendationIdFilterOk returns a tuple with the RecommendationIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetRecommendationIdFilterOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationIdFilter) {
		return nil, false
	}
	return o.RecommendationIdFilter, true
}

// HasRecommendationIdFilter returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) HasRecommendationIdFilter() bool {
	if o != nil && !IsNil(o.RecommendationIdFilter) {
		return true
	}

	return false
}

// SetRecommendationIdFilter gets a reference to the given string and assigns it to the RecommendationIdFilter field.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) SetRecommendationIdFilter(v string) {
	o.RecommendationIdFilter = &v
}

// GetCampaignIdListFilter returns the CampaignIdListFilter field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetCampaignIdListFilter() []OptimizationRulesAPICampaignFilter {
	if o == nil || IsNil(o.CampaignIdListFilter) {
		var ret []OptimizationRulesAPICampaignFilter
		return ret
	}
	return o.CampaignIdListFilter
}

// GetCampaignIdListFilterOk returns a tuple with the CampaignIdListFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetCampaignIdListFilterOk() ([]OptimizationRulesAPICampaignFilter, bool) {
	if o == nil || IsNil(o.CampaignIdListFilter) {
		return nil, false
	}
	return o.CampaignIdListFilter, true
}

// HasCampaignIdListFilter returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) HasCampaignIdListFilter() bool {
	if o != nil && !IsNil(o.CampaignIdListFilter) {
		return true
	}

	return false
}

// SetCampaignIdListFilter gets a reference to the given []OptimizationRulesAPICampaignFilter and assigns it to the CampaignIdListFilter field.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) SetCampaignIdListFilter(v []OptimizationRulesAPICampaignFilter) {
	o.CampaignIdListFilter = v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetMaxResults() float32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret float32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetMaxResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given float32 and assigns it to the MaxResults field.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) SetMaxResults(v float32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) SetNextToken(v string) {
	o.NextToken = &v
}

// GetRuleCategoryFilter returns the RuleCategoryFilter field value
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetRuleCategoryFilter() OptimizationRulesAPIEntityFieldFilter {
	if o == nil {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}

	return o.RuleCategoryFilter
}

// GetRuleCategoryFilterOk returns a tuple with the RuleCategoryFilter field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) GetRuleCategoryFilterOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleCategoryFilter, true
}

// SetRuleCategoryFilter sets field value
func (o *OptimizationRulesAPIGetRuleRecommendationsRequest) SetRuleCategoryFilter(v OptimizationRulesAPIEntityFieldFilter) {
	o.RuleCategoryFilter = v
}

func (o OptimizationRulesAPIGetRuleRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ruleSubCategoryFilter"] = o.RuleSubCategoryFilter
	if !IsNil(o.RecommendationIdFilter) {
		toSerialize["recommendationIdFilter"] = o.RecommendationIdFilter
	}
	if !IsNil(o.CampaignIdListFilter) {
		toSerialize["campaignIdListFilter"] = o.CampaignIdListFilter
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["ruleCategoryFilter"] = o.RuleCategoryFilter
	return toSerialize, nil
}

type NullableOptimizationRulesAPIGetRuleRecommendationsRequest struct {
	value *OptimizationRulesAPIGetRuleRecommendationsRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIGetRuleRecommendationsRequest) Get() *OptimizationRulesAPIGetRuleRecommendationsRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIGetRuleRecommendationsRequest) Set(val *OptimizationRulesAPIGetRuleRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIGetRuleRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIGetRuleRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIGetRuleRecommendationsRequest(val *OptimizationRulesAPIGetRuleRecommendationsRequest) *NullableOptimizationRulesAPIGetRuleRecommendationsRequest {
	return &NullableOptimizationRulesAPIGetRuleRecommendationsRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIGetRuleRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIGetRuleRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
