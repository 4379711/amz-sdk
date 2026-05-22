package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleCondition{}

// RuleCondition struct for RuleCondition
type RuleCondition struct {
	MetricName         RuleConditionMetric `json:"metricName"`
	ComparisonOperator ComparisonOperator  `json:"comparisonOperator"`
	// The performance threshold value.
	Threshold float64 `json:"threshold"`
}

type _RuleCondition RuleCondition

// NewRuleCondition instantiates a new RuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleCondition(metricName RuleConditionMetric, comparisonOperator ComparisonOperator, threshold float64) *RuleCondition {
	this := RuleCondition{}
	this.MetricName = metricName
	this.ComparisonOperator = comparisonOperator
	this.Threshold = threshold
	return &this
}

// NewRuleConditionWithDefaults instantiates a new RuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConditionWithDefaults() *RuleCondition {
	this := RuleCondition{}
	return &this
}

// GetMetricName returns the MetricName field value
func (o *RuleCondition) GetMetricName() RuleConditionMetric {
	if o == nil {
		var ret RuleConditionMetric
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *RuleCondition) GetMetricNameOk() (*RuleConditionMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *RuleCondition) SetMetricName(v RuleConditionMetric) {
	o.MetricName = v
}

// GetComparisonOperator returns the ComparisonOperator field value
func (o *RuleCondition) GetComparisonOperator() ComparisonOperator {
	if o == nil {
		var ret ComparisonOperator
		return ret
	}

	return o.ComparisonOperator
}

// GetComparisonOperatorOk returns a tuple with the ComparisonOperator field value
// and a boolean to check if the value has been set.
func (o *RuleCondition) GetComparisonOperatorOk() (*ComparisonOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonOperator, true
}

// SetComparisonOperator sets field value
func (o *RuleCondition) SetComparisonOperator(v ComparisonOperator) {
	o.ComparisonOperator = v
}

// GetThreshold returns the Threshold field value
func (o *RuleCondition) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *RuleCondition) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *RuleCondition) SetThreshold(v float64) {
	o.Threshold = v
}

func (o RuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricName"] = o.MetricName
	toSerialize["comparisonOperator"] = o.ComparisonOperator
	toSerialize["threshold"] = o.Threshold
	return toSerialize, nil
}

type NullableRuleCondition struct {
	value *RuleCondition
	isSet bool
}

func (v NullableRuleCondition) Get() *RuleCondition {
	return v.value
}

func (v *NullableRuleCondition) Set(val *RuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleCondition(val *RuleCondition) *NullableRuleCondition {
	return &NullableRuleCondition{value: val, isSet: true}
}

func (v NullableRuleCondition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
