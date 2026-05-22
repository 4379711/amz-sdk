package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem{}

// OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem struct for OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem
type OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem struct {
	// Index of the recommendation id in the request.
	Index int32 `json:"index"`
	// The id that uniquely identifies the recommendation that was updated successfully
	RecommendationId string `json:"recommendationId"`
}

type _OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem(index int32, recommendationId string) *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem{}
	this.Index = index
	this.RecommendationId = recommendationId
	return &this
}

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItemWithDefaults instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItemWithDefaults() *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) SetIndex(v int32) {
	o.Index = v
}

// GetRecommendationId returns the RecommendationId field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) GetRecommendationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) GetRecommendationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationId, true
}

// SetRecommendationId sets field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) SetRecommendationId(v string) {
	o.RecommendationId = v
}

func (o OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["recommendationId"] = o.RecommendationId
	return toSerialize, nil
}

type NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem struct {
	value *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem
	isSet bool
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) Get() *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) Set(val *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem(val *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem {
	return &NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseSuccessItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
