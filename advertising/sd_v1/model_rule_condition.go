package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the RuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleCondition{}

// RuleCondition A rule condition that defines the advertiser's intent for the outcome of the rule. Certain actions are performed by the product to achieve and maintain the rule condition.
type RuleCondition struct {
	// The name of the metric. Supported rule metrics and corresponding supported comparisonOperators: |      MetricName      |ComparisonOperator  |Description| |------------------|--------------------|-------------------| |COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS     |              LESS_THAN_OR_EQUAL_TO             |Maximize viewable impressions while cost per 1000 views less than or equal to `threshold`| |COST_PER_CLICK    |              LESS_THAN_OR_EQUAL_TO            |Maximize page visits while cost per click less than or equal to `threshold`| |COST_PER_ORDER    |              LESS_THAN_OR_EQUAL_TO            |Maximize viewable impressions/page visits/conversion while cost per order less than or equal to `threshold`|
	MetricName string `json:"metricName"`
	// The comparison operator.
	ComparisonOperator string `json:"comparisonOperator"`
	// The value of the threshold associated with the metric. The threshold values has defined minimums depending on the metric names in the following table: |                  MetricName            | Minimum of `threshold` Value  | |----------------------------------------|-----------------------------------| |COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS  | 1                                 | |COST_PER_CLICK                          | 0.5                               | |COST_PER_ORDER                          | 5                                 |
	Threshold float64 `json:"threshold"`
}

type _RuleCondition RuleCondition

// NewRuleCondition instantiates a new RuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleCondition(metricName string, comparisonOperator string, threshold float64) *RuleCondition {
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
func (o *RuleCondition) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *RuleCondition) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *RuleCondition) SetMetricName(v string) {
	o.MetricName = v
}

// GetComparisonOperator returns the ComparisonOperator field value
func (o *RuleCondition) GetComparisonOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComparisonOperator
}

// GetComparisonOperatorOk returns a tuple with the ComparisonOperator field value
// and a boolean to check if the value has been set.
func (o *RuleCondition) GetComparisonOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonOperator, true
}

// SetComparisonOperator sets field value
func (o *RuleCondition) SetComparisonOperator(v string) {
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
