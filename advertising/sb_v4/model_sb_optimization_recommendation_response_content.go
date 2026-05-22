package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBOptimizationRecommendationResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBOptimizationRecommendationResponseContent{}

// SBOptimizationRecommendationResponseContent struct for SBOptimizationRecommendationResponseContent
type SBOptimizationRecommendationResponseContent struct {
	// Minimum accepted value for cost control metric.
	MinimumValue      float64           `json:"minimumValue"`
	CostControlMetric CostControlMetric `json:"costControlMetric"`
	// Recommended value for cost control metric.
	RecommendedValue float64 `json:"recommendedValue"`
}

type _SBOptimizationRecommendationResponseContent SBOptimizationRecommendationResponseContent

// NewSBOptimizationRecommendationResponseContent instantiates a new SBOptimizationRecommendationResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBOptimizationRecommendationResponseContent(minimumValue float64, costControlMetric CostControlMetric, recommendedValue float64) *SBOptimizationRecommendationResponseContent {
	this := SBOptimizationRecommendationResponseContent{}
	this.MinimumValue = minimumValue
	this.CostControlMetric = costControlMetric
	this.RecommendedValue = recommendedValue
	return &this
}

// NewSBOptimizationRecommendationResponseContentWithDefaults instantiates a new SBOptimizationRecommendationResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBOptimizationRecommendationResponseContentWithDefaults() *SBOptimizationRecommendationResponseContent {
	this := SBOptimizationRecommendationResponseContent{}
	return &this
}

// GetMinimumValue returns the MinimumValue field value
func (o *SBOptimizationRecommendationResponseContent) GetMinimumValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MinimumValue
}

// GetMinimumValueOk returns a tuple with the MinimumValue field value
// and a boolean to check if the value has been set.
func (o *SBOptimizationRecommendationResponseContent) GetMinimumValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumValue, true
}

// SetMinimumValue sets field value
func (o *SBOptimizationRecommendationResponseContent) SetMinimumValue(v float64) {
	o.MinimumValue = v
}

// GetCostControlMetric returns the CostControlMetric field value
func (o *SBOptimizationRecommendationResponseContent) GetCostControlMetric() CostControlMetric {
	if o == nil {
		var ret CostControlMetric
		return ret
	}

	return o.CostControlMetric
}

// GetCostControlMetricOk returns a tuple with the CostControlMetric field value
// and a boolean to check if the value has been set.
func (o *SBOptimizationRecommendationResponseContent) GetCostControlMetricOk() (*CostControlMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostControlMetric, true
}

// SetCostControlMetric sets field value
func (o *SBOptimizationRecommendationResponseContent) SetCostControlMetric(v CostControlMetric) {
	o.CostControlMetric = v
}

// GetRecommendedValue returns the RecommendedValue field value
func (o *SBOptimizationRecommendationResponseContent) GetRecommendedValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.RecommendedValue
}

// GetRecommendedValueOk returns a tuple with the RecommendedValue field value
// and a boolean to check if the value has been set.
func (o *SBOptimizationRecommendationResponseContent) GetRecommendedValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendedValue, true
}

// SetRecommendedValue sets field value
func (o *SBOptimizationRecommendationResponseContent) SetRecommendedValue(v float64) {
	o.RecommendedValue = v
}

func (o SBOptimizationRecommendationResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minimumValue"] = o.MinimumValue
	toSerialize["costControlMetric"] = o.CostControlMetric
	toSerialize["recommendedValue"] = o.RecommendedValue
	return toSerialize, nil
}

type NullableSBOptimizationRecommendationResponseContent struct {
	value *SBOptimizationRecommendationResponseContent
	isSet bool
}

func (v NullableSBOptimizationRecommendationResponseContent) Get() *SBOptimizationRecommendationResponseContent {
	return v.value
}

func (v *NullableSBOptimizationRecommendationResponseContent) Set(val *SBOptimizationRecommendationResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBOptimizationRecommendationResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBOptimizationRecommendationResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBOptimizationRecommendationResponseContent(val *SBOptimizationRecommendationResponseContent) *NullableSBOptimizationRecommendationResponseContent {
	return &NullableSBOptimizationRecommendationResponseContent{value: val, isSet: true}
}

func (v NullableSBOptimizationRecommendationResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBOptimizationRecommendationResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
