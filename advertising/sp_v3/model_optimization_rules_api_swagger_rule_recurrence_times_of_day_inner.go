package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRuleRecurrenceTimesOfDayInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRuleRecurrenceTimesOfDayInner{}

// OptimizationRulesAPIRuleRecurrenceTimesOfDayInner List of times of the day.
type OptimizationRulesAPIRuleRecurrenceTimesOfDayInner struct {
	// Time of the day in HH:MM.
	StartTime string `json:"startTime" validate:"regexp=^\\\\d{2}:\\\\d{2}$"`
	// Time of the day in HH:MM.
	EndTime string `json:"endTime" validate:"regexp=^\\\\d{2}:\\\\d{2}$"`
}

type _OptimizationRulesAPIRuleRecurrenceTimesOfDayInner OptimizationRulesAPIRuleRecurrenceTimesOfDayInner

// NewOptimizationRulesAPIRuleRecurrenceTimesOfDayInner instantiates a new OptimizationRulesAPIRuleRecurrenceTimesOfDayInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRuleRecurrenceTimesOfDayInner(startTime string, endTime string) *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner {
	this := OptimizationRulesAPIRuleRecurrenceTimesOfDayInner{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewOptimizationRulesAPIRuleRecurrenceTimesOfDayInnerWithDefaults instantiates a new OptimizationRulesAPIRuleRecurrenceTimesOfDayInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRuleRecurrenceTimesOfDayInnerWithDefaults() *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner {
	this := OptimizationRulesAPIRuleRecurrenceTimesOfDayInner{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) SetEndTime(v string) {
	o.EndTime = v
}

func (o OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner struct {
	value *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner) Get() *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner) Set(val *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner(val *OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) *NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner {
	return &NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleRecurrenceTimesOfDayInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
