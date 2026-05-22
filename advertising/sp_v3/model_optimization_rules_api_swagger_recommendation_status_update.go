package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRecommendationStatusUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRecommendationStatusUpdate{}

// OptimizationRulesAPIRecommendationStatusUpdate struct for OptimizationRulesAPIRecommendationStatusUpdate
type OptimizationRulesAPIRecommendationStatusUpdate struct {
	RecommendationId *OptimizationRulesAPIEntityFieldFilter `json:"recommendationId,omitempty"`
	Status           *OptimizationRulesAPIEntityFieldFilter `json:"status,omitempty"`
}

// NewOptimizationRulesAPIRecommendationStatusUpdate instantiates a new OptimizationRulesAPIRecommendationStatusUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRecommendationStatusUpdate() *OptimizationRulesAPIRecommendationStatusUpdate {
	this := OptimizationRulesAPIRecommendationStatusUpdate{}
	return &this
}

// NewOptimizationRulesAPIRecommendationStatusUpdateWithDefaults instantiates a new OptimizationRulesAPIRecommendationStatusUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRecommendationStatusUpdateWithDefaults() *OptimizationRulesAPIRecommendationStatusUpdate {
	this := OptimizationRulesAPIRecommendationStatusUpdate{}
	return &this
}

// GetRecommendationId returns the RecommendationId field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) GetRecommendationId() OptimizationRulesAPIEntityFieldFilter {
	if o == nil || IsNil(o.RecommendationId) {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}
	return *o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) GetRecommendationIdOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil || IsNil(o.RecommendationId) {
		return nil, false
	}
	return o.RecommendationId, true
}

// HasRecommendationId returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) HasRecommendationId() bool {
	if o != nil && !IsNil(o.RecommendationId) {
		return true
	}

	return false
}

// SetRecommendationId gets a reference to the given OptimizationRulesAPIEntityFieldFilter and assigns it to the RecommendationId field.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) SetRecommendationId(v OptimizationRulesAPIEntityFieldFilter) {
	o.RecommendationId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) GetStatus() OptimizationRulesAPIEntityFieldFilter {
	if o == nil || IsNil(o.Status) {
		var ret OptimizationRulesAPIEntityFieldFilter
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) GetStatusOk() (*OptimizationRulesAPIEntityFieldFilter, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OptimizationRulesAPIEntityFieldFilter and assigns it to the Status field.
func (o *OptimizationRulesAPIRecommendationStatusUpdate) SetStatus(v OptimizationRulesAPIEntityFieldFilter) {
	o.Status = &v
}

func (o OptimizationRulesAPIRecommendationStatusUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecommendationId) {
		toSerialize["recommendationId"] = o.RecommendationId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRecommendationStatusUpdate struct {
	value *OptimizationRulesAPIRecommendationStatusUpdate
	isSet bool
}

func (v NullableOptimizationRulesAPIRecommendationStatusUpdate) Get() *OptimizationRulesAPIRecommendationStatusUpdate {
	return v.value
}

func (v *NullableOptimizationRulesAPIRecommendationStatusUpdate) Set(val *OptimizationRulesAPIRecommendationStatusUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRecommendationStatusUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRecommendationStatusUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRecommendationStatusUpdate(val *OptimizationRulesAPIRecommendationStatusUpdate) *NullableOptimizationRulesAPIRecommendationStatusUpdate {
	return &NullableOptimizationRulesAPIRecommendationStatusUpdate{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRecommendationStatusUpdate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRecommendationStatusUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
