package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the PerformanceMeasureCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceMeasureCondition{}

// PerformanceMeasureCondition struct for PerformanceMeasureCondition
type PerformanceMeasureCondition struct {
	MetricName         PerformanceMetric  `json:"metricName"`
	ComparisonOperator ComparisonOperator `json:"comparisonOperator"`
	// The performance threshold value.
	Threshold float64 `json:"threshold"`
}

type _PerformanceMeasureCondition PerformanceMeasureCondition

// NewPerformanceMeasureCondition instantiates a new PerformanceMeasureCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceMeasureCondition(metricName PerformanceMetric, comparisonOperator ComparisonOperator, threshold float64) *PerformanceMeasureCondition {
	this := PerformanceMeasureCondition{}
	this.MetricName = metricName
	this.ComparisonOperator = comparisonOperator
	this.Threshold = threshold
	return &this
}

// NewPerformanceMeasureConditionWithDefaults instantiates a new PerformanceMeasureCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceMeasureConditionWithDefaults() *PerformanceMeasureCondition {
	this := PerformanceMeasureCondition{}
	return &this
}

// GetMetricName returns the MetricName field value
func (o *PerformanceMeasureCondition) GetMetricName() PerformanceMetric {
	if o == nil {
		var ret PerformanceMetric
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *PerformanceMeasureCondition) GetMetricNameOk() (*PerformanceMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *PerformanceMeasureCondition) SetMetricName(v PerformanceMetric) {
	o.MetricName = v
}

// GetComparisonOperator returns the ComparisonOperator field value
func (o *PerformanceMeasureCondition) GetComparisonOperator() ComparisonOperator {
	if o == nil {
		var ret ComparisonOperator
		return ret
	}

	return o.ComparisonOperator
}

// GetComparisonOperatorOk returns a tuple with the ComparisonOperator field value
// and a boolean to check if the value has been set.
func (o *PerformanceMeasureCondition) GetComparisonOperatorOk() (*ComparisonOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonOperator, true
}

// SetComparisonOperator sets field value
func (o *PerformanceMeasureCondition) SetComparisonOperator(v ComparisonOperator) {
	o.ComparisonOperator = v
}

// GetThreshold returns the Threshold field value
func (o *PerformanceMeasureCondition) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *PerformanceMeasureCondition) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *PerformanceMeasureCondition) SetThreshold(v float64) {
	o.Threshold = v
}

func (o PerformanceMeasureCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricName"] = o.MetricName
	toSerialize["comparisonOperator"] = o.ComparisonOperator
	toSerialize["threshold"] = o.Threshold
	return toSerialize, nil
}

type NullablePerformanceMeasureCondition struct {
	value *PerformanceMeasureCondition
	isSet bool
}

func (v NullablePerformanceMeasureCondition) Get() *PerformanceMeasureCondition {
	return v.value
}

func (v *NullablePerformanceMeasureCondition) Set(val *PerformanceMeasureCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceMeasureCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceMeasureCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceMeasureCondition(val *PerformanceMeasureCondition) *NullablePerformanceMeasureCondition {
	return &NullablePerformanceMeasureCondition{value: val, isSet: true}
}

func (v NullablePerformanceMeasureCondition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePerformanceMeasureCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
