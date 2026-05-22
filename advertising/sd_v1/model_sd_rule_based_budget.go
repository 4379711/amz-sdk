package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDRuleBasedBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDRuleBasedBudget{}

// SDRuleBasedBudget struct for SDRuleBasedBudget
type SDRuleBasedBudget struct {
	// Epoch time of budget rule execution.
	ExecutionTime *float32      `json:"executionTime,omitempty"`
	AppliedRule   *SDBudgetRule `json:"appliedRule,omitempty"`
	// The budget value.
	RuleBasedBudgetValue *float32 `json:"ruleBasedBudgetValue,omitempty"`
	// The daily budget value.
	DailyBudgetValue  *float32                `json:"dailyBudgetValue,omitempty"`
	PerformanceMetric *PerformanceMetricValue `json:"performanceMetric,omitempty"`
}

// NewSDRuleBasedBudget instantiates a new SDRuleBasedBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDRuleBasedBudget() *SDRuleBasedBudget {
	this := SDRuleBasedBudget{}
	return &this
}

// NewSDRuleBasedBudgetWithDefaults instantiates a new SDRuleBasedBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDRuleBasedBudgetWithDefaults() *SDRuleBasedBudget {
	this := SDRuleBasedBudget{}
	return &this
}

// GetExecutionTime returns the ExecutionTime field value if set, zero value otherwise.
func (o *SDRuleBasedBudget) GetExecutionTime() float32 {
	if o == nil || IsNil(o.ExecutionTime) {
		var ret float32
		return ret
	}
	return *o.ExecutionTime
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDRuleBasedBudget) GetExecutionTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ExecutionTime) {
		return nil, false
	}
	return o.ExecutionTime, true
}

// HasExecutionTime returns a boolean if a field has been set.
func (o *SDRuleBasedBudget) HasExecutionTime() bool {
	if o != nil && !IsNil(o.ExecutionTime) {
		return true
	}

	return false
}

// SetExecutionTime gets a reference to the given float32 and assigns it to the ExecutionTime field.
func (o *SDRuleBasedBudget) SetExecutionTime(v float32) {
	o.ExecutionTime = &v
}

// GetAppliedRule returns the AppliedRule field value if set, zero value otherwise.
func (o *SDRuleBasedBudget) GetAppliedRule() SDBudgetRule {
	if o == nil || IsNil(o.AppliedRule) {
		var ret SDBudgetRule
		return ret
	}
	return *o.AppliedRule
}

// GetAppliedRuleOk returns a tuple with the AppliedRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDRuleBasedBudget) GetAppliedRuleOk() (*SDBudgetRule, bool) {
	if o == nil || IsNil(o.AppliedRule) {
		return nil, false
	}
	return o.AppliedRule, true
}

// HasAppliedRule returns a boolean if a field has been set.
func (o *SDRuleBasedBudget) HasAppliedRule() bool {
	if o != nil && !IsNil(o.AppliedRule) {
		return true
	}

	return false
}

// SetAppliedRule gets a reference to the given SDBudgetRule and assigns it to the AppliedRule field.
func (o *SDRuleBasedBudget) SetAppliedRule(v SDBudgetRule) {
	o.AppliedRule = &v
}

// GetRuleBasedBudgetValue returns the RuleBasedBudgetValue field value if set, zero value otherwise.
func (o *SDRuleBasedBudget) GetRuleBasedBudgetValue() float32 {
	if o == nil || IsNil(o.RuleBasedBudgetValue) {
		var ret float32
		return ret
	}
	return *o.RuleBasedBudgetValue
}

// GetRuleBasedBudgetValueOk returns a tuple with the RuleBasedBudgetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDRuleBasedBudget) GetRuleBasedBudgetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.RuleBasedBudgetValue) {
		return nil, false
	}
	return o.RuleBasedBudgetValue, true
}

// HasRuleBasedBudgetValue returns a boolean if a field has been set.
func (o *SDRuleBasedBudget) HasRuleBasedBudgetValue() bool {
	if o != nil && !IsNil(o.RuleBasedBudgetValue) {
		return true
	}

	return false
}

// SetRuleBasedBudgetValue gets a reference to the given float32 and assigns it to the RuleBasedBudgetValue field.
func (o *SDRuleBasedBudget) SetRuleBasedBudgetValue(v float32) {
	o.RuleBasedBudgetValue = &v
}

// GetDailyBudgetValue returns the DailyBudgetValue field value if set, zero value otherwise.
func (o *SDRuleBasedBudget) GetDailyBudgetValue() float32 {
	if o == nil || IsNil(o.DailyBudgetValue) {
		var ret float32
		return ret
	}
	return *o.DailyBudgetValue
}

// GetDailyBudgetValueOk returns a tuple with the DailyBudgetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDRuleBasedBudget) GetDailyBudgetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.DailyBudgetValue) {
		return nil, false
	}
	return o.DailyBudgetValue, true
}

// HasDailyBudgetValue returns a boolean if a field has been set.
func (o *SDRuleBasedBudget) HasDailyBudgetValue() bool {
	if o != nil && !IsNil(o.DailyBudgetValue) {
		return true
	}

	return false
}

// SetDailyBudgetValue gets a reference to the given float32 and assigns it to the DailyBudgetValue field.
func (o *SDRuleBasedBudget) SetDailyBudgetValue(v float32) {
	o.DailyBudgetValue = &v
}

// GetPerformanceMetric returns the PerformanceMetric field value if set, zero value otherwise.
func (o *SDRuleBasedBudget) GetPerformanceMetric() PerformanceMetricValue {
	if o == nil || IsNil(o.PerformanceMetric) {
		var ret PerformanceMetricValue
		return ret
	}
	return *o.PerformanceMetric
}

// GetPerformanceMetricOk returns a tuple with the PerformanceMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDRuleBasedBudget) GetPerformanceMetricOk() (*PerformanceMetricValue, bool) {
	if o == nil || IsNil(o.PerformanceMetric) {
		return nil, false
	}
	return o.PerformanceMetric, true
}

// HasPerformanceMetric returns a boolean if a field has been set.
func (o *SDRuleBasedBudget) HasPerformanceMetric() bool {
	if o != nil && !IsNil(o.PerformanceMetric) {
		return true
	}

	return false
}

// SetPerformanceMetric gets a reference to the given PerformanceMetricValue and assigns it to the PerformanceMetric field.
func (o *SDRuleBasedBudget) SetPerformanceMetric(v PerformanceMetricValue) {
	o.PerformanceMetric = &v
}

func (o SDRuleBasedBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutionTime) {
		toSerialize["executionTime"] = o.ExecutionTime
	}
	if !IsNil(o.AppliedRule) {
		toSerialize["appliedRule"] = o.AppliedRule
	}
	if !IsNil(o.RuleBasedBudgetValue) {
		toSerialize["ruleBasedBudgetValue"] = o.RuleBasedBudgetValue
	}
	if !IsNil(o.DailyBudgetValue) {
		toSerialize["dailyBudgetValue"] = o.DailyBudgetValue
	}
	if !IsNil(o.PerformanceMetric) {
		toSerialize["performanceMetric"] = o.PerformanceMetric
	}
	return toSerialize, nil
}

type NullableSDRuleBasedBudget struct {
	value *SDRuleBasedBudget
	isSet bool
}

func (v NullableSDRuleBasedBudget) Get() *SDRuleBasedBudget {
	return v.value
}

func (v *NullableSDRuleBasedBudget) Set(val *SDRuleBasedBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSDRuleBasedBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSDRuleBasedBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDRuleBasedBudget(val *SDRuleBasedBudget) *NullableSDRuleBasedBudget {
	return &NullableSDRuleBasedBudget{value: val, isSet: true}
}

func (v NullableSDRuleBasedBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDRuleBasedBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
