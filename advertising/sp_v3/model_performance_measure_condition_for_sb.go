package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PerformanceMeasureConditionForSB type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceMeasureConditionForSB{}

// PerformanceMeasureConditionForSB struct for PerformanceMeasureConditionForSB
type PerformanceMeasureConditionForSB struct {
	MetricName         PerformanceMetricForSB `json:"metricName"`
	ComparisonOperator ComparisonOperator     `json:"comparisonOperator"`
	// The performance threshold value.
	Threshold float64 `json:"threshold"`
}

type _PerformanceMeasureConditionForSB PerformanceMeasureConditionForSB

// NewPerformanceMeasureConditionForSB instantiates a new PerformanceMeasureConditionForSB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceMeasureConditionForSB(metricName PerformanceMetricForSB, comparisonOperator ComparisonOperator, threshold float64) *PerformanceMeasureConditionForSB {
	this := PerformanceMeasureConditionForSB{}
	this.MetricName = metricName
	this.ComparisonOperator = comparisonOperator
	this.Threshold = threshold
	return &this
}

// NewPerformanceMeasureConditionForSBWithDefaults instantiates a new PerformanceMeasureConditionForSB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceMeasureConditionForSBWithDefaults() *PerformanceMeasureConditionForSB {
	this := PerformanceMeasureConditionForSB{}
	return &this
}

// GetMetricName returns the MetricName field value
func (o *PerformanceMeasureConditionForSB) GetMetricName() PerformanceMetricForSB {
	if o == nil {
		var ret PerformanceMetricForSB
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *PerformanceMeasureConditionForSB) GetMetricNameOk() (*PerformanceMetricForSB, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *PerformanceMeasureConditionForSB) SetMetricName(v PerformanceMetricForSB) {
	o.MetricName = v
}

// GetComparisonOperator returns the ComparisonOperator field value
func (o *PerformanceMeasureConditionForSB) GetComparisonOperator() ComparisonOperator {
	if o == nil {
		var ret ComparisonOperator
		return ret
	}

	return o.ComparisonOperator
}

// GetComparisonOperatorOk returns a tuple with the ComparisonOperator field value
// and a boolean to check if the value has been set.
func (o *PerformanceMeasureConditionForSB) GetComparisonOperatorOk() (*ComparisonOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonOperator, true
}

// SetComparisonOperator sets field value
func (o *PerformanceMeasureConditionForSB) SetComparisonOperator(v ComparisonOperator) {
	o.ComparisonOperator = v
}

// GetThreshold returns the Threshold field value
func (o *PerformanceMeasureConditionForSB) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *PerformanceMeasureConditionForSB) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *PerformanceMeasureConditionForSB) SetThreshold(v float64) {
	o.Threshold = v
}

func (o PerformanceMeasureConditionForSB) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricName"] = o.MetricName
	toSerialize["comparisonOperator"] = o.ComparisonOperator
	toSerialize["threshold"] = o.Threshold
	return toSerialize, nil
}

type NullablePerformanceMeasureConditionForSB struct {
	value *PerformanceMeasureConditionForSB
	isSet bool
}

func (v NullablePerformanceMeasureConditionForSB) Get() *PerformanceMeasureConditionForSB {
	return v.value
}

func (v *NullablePerformanceMeasureConditionForSB) Set(val *PerformanceMeasureConditionForSB) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceMeasureConditionForSB) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceMeasureConditionForSB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceMeasureConditionForSB(val *PerformanceMeasureConditionForSB) *NullablePerformanceMeasureConditionForSB {
	return &NullablePerformanceMeasureConditionForSB{value: val, isSet: true}
}

func (v NullablePerformanceMeasureConditionForSB) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePerformanceMeasureConditionForSB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
