package sp_v3

import "github.com/bytedance/sonic"

// checks if the OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem{}

// OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem struct for OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem
type OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem struct {
	// Index of the recommendation id in the request.
	Index int32 `json:"index"`
	// The id that uniquely identifies the recommendation rule that failed to be updated.
	RecommendationId string `json:"recommendationId"`
	// The http status error code for machine use.
	HttpStatusCode string                                              `json:"httpStatusCode"`
	SubErrors      []OptimizationRulesAPIOptimizationRuleBatchSubError `json:"subErrors"`
}

type _OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem(index int32, recommendationId string, httpStatusCode string, subErrors []OptimizationRulesAPIOptimizationRuleBatchSubError) *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem{}
	this.Index = index
	this.RecommendationId = recommendationId
	this.HttpStatusCode = httpStatusCode
	this.SubErrors = subErrors
	return &this
}

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItemWithDefaults instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItemWithDefaults() *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) SetIndex(v int32) {
	o.Index = v
}

// GetRecommendationId returns the RecommendationId field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetRecommendationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetRecommendationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationId, true
}

// SetRecommendationId sets field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) SetRecommendationId(v string) {
	o.RecommendationId = v
}

// GetHttpStatusCode returns the HttpStatusCode field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetHttpStatusCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatusCode, true
}

// SetHttpStatusCode sets field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) SetHttpStatusCode(v string) {
	o.HttpStatusCode = v
}

// GetSubErrors returns the SubErrors field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetSubErrors() []OptimizationRulesAPIOptimizationRuleBatchSubError {
	if o == nil {
		var ret []OptimizationRulesAPIOptimizationRuleBatchSubError
		return ret
	}

	return o.SubErrors
}

// GetSubErrorsOk returns a tuple with the SubErrors field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) GetSubErrorsOk() ([]OptimizationRulesAPIOptimizationRuleBatchSubError, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubErrors, true
}

// SetSubErrors sets field value
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) SetSubErrors(v []OptimizationRulesAPIOptimizationRuleBatchSubError) {
	o.SubErrors = v
}

func (o OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["recommendationId"] = o.RecommendationId
	toSerialize["httpStatusCode"] = o.HttpStatusCode
	toSerialize["subErrors"] = o.SubErrors
	return toSerialize, nil
}

type NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem struct {
	value *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem
	isSet bool
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) Get() *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) Set(val *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem(val *OptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem {
	return &NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusResponseErrorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
