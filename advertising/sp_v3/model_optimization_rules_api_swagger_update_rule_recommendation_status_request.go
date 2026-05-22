package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIUpdateRuleRecommendationStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIUpdateRuleRecommendationStatusRequest{}

// OptimizationRulesAPIUpdateRuleRecommendationStatusRequest Request object for getting rule recommendation. All the filters in the request are used in conjunction with one another.
type OptimizationRulesAPIUpdateRuleRecommendationStatusRequest struct {
	Recommendations []OptimizationRulesAPIRecommendationStatusUpdate `json:"recommendations,omitempty"`
}

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusRequest instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusRequest() *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusRequest{}
	return &this
}

// NewOptimizationRulesAPIUpdateRuleRecommendationStatusRequestWithDefaults instantiates a new OptimizationRulesAPIUpdateRuleRecommendationStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIUpdateRuleRecommendationStatusRequestWithDefaults() *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest {
	this := OptimizationRulesAPIUpdateRuleRecommendationStatusRequest{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest) GetRecommendations() []OptimizationRulesAPIRecommendationStatusUpdate {
	if o == nil || IsNil(o.Recommendations) {
		var ret []OptimizationRulesAPIRecommendationStatusUpdate
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest) GetRecommendationsOk() ([]OptimizationRulesAPIRecommendationStatusUpdate, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []OptimizationRulesAPIRecommendationStatusUpdate and assigns it to the Recommendations field.
func (o *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest) SetRecommendations(v []OptimizationRulesAPIRecommendationStatusUpdate) {
	o.Recommendations = v
}

func (o OptimizationRulesAPIUpdateRuleRecommendationStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest struct {
	value *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest) Get() *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest) Set(val *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest(val *OptimizationRulesAPIUpdateRuleRecommendationStatusRequest) *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest {
	return &NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIUpdateRuleRecommendationStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
