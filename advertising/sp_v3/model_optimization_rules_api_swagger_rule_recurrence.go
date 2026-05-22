package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRuleRecurrence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRuleRecurrence{}

// OptimizationRulesAPIRuleRecurrence The recurrence of the optimization rule application.
type OptimizationRulesAPIRuleRecurrence struct {
	Duration OptimizationRulesAPIDuration `json:"duration"`
	// List of times of the day.
	TimesOfDay []OptimizationRulesAPIRuleRecurrenceTimesOfDayInner `json:"timesOfDay,omitempty"`
	Type       OptimizationRulesAPIRuleRecurrenceType              `json:"type"`
	// A list of days of the week.
	DaysOfWeek []OptimizationRulesAPIDayOfTheWeek `json:"daysOfWeek,omitempty"`
}

type _OptimizationRulesAPIRuleRecurrence OptimizationRulesAPIRuleRecurrence

// NewOptimizationRulesAPIRuleRecurrence instantiates a new OptimizationRulesAPIRuleRecurrence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRuleRecurrence(duration OptimizationRulesAPIDuration, type_ OptimizationRulesAPIRuleRecurrenceType) *OptimizationRulesAPIRuleRecurrence {
	this := OptimizationRulesAPIRuleRecurrence{}
	this.Duration = duration
	this.Type = type_
	return &this
}

// NewOptimizationRulesAPIRuleRecurrenceWithDefaults instantiates a new OptimizationRulesAPIRuleRecurrence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRuleRecurrenceWithDefaults() *OptimizationRulesAPIRuleRecurrence {
	this := OptimizationRulesAPIRuleRecurrence{}
	return &this
}

// GetDuration returns the Duration field value
func (o *OptimizationRulesAPIRuleRecurrence) GetDuration() OptimizationRulesAPIDuration {
	if o == nil {
		var ret OptimizationRulesAPIDuration
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecurrence) GetDurationOk() (*OptimizationRulesAPIDuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *OptimizationRulesAPIRuleRecurrence) SetDuration(v OptimizationRulesAPIDuration) {
	o.Duration = v
}

// GetTimesOfDay returns the TimesOfDay field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecurrence) GetTimesOfDay() []OptimizationRulesAPIRuleRecurrenceTimesOfDayInner {
	if o == nil || IsNil(o.TimesOfDay) {
		var ret []OptimizationRulesAPIRuleRecurrenceTimesOfDayInner
		return ret
	}
	return o.TimesOfDay
}

// GetTimesOfDayOk returns a tuple with the TimesOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecurrence) GetTimesOfDayOk() ([]OptimizationRulesAPIRuleRecurrenceTimesOfDayInner, bool) {
	if o == nil || IsNil(o.TimesOfDay) {
		return nil, false
	}
	return o.TimesOfDay, true
}

// HasTimesOfDay returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecurrence) HasTimesOfDay() bool {
	if o != nil && !IsNil(o.TimesOfDay) {
		return true
	}

	return false
}

// SetTimesOfDay gets a reference to the given []OptimizationRulesAPIRuleRecurrenceTimesOfDayInner and assigns it to the TimesOfDay field.
func (o *OptimizationRulesAPIRuleRecurrence) SetTimesOfDay(v []OptimizationRulesAPIRuleRecurrenceTimesOfDayInner) {
	o.TimesOfDay = v
}

// GetType returns the Type field value
func (o *OptimizationRulesAPIRuleRecurrence) GetType() OptimizationRulesAPIRuleRecurrenceType {
	if o == nil {
		var ret OptimizationRulesAPIRuleRecurrenceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecurrence) GetTypeOk() (*OptimizationRulesAPIRuleRecurrenceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OptimizationRulesAPIRuleRecurrence) SetType(v OptimizationRulesAPIRuleRecurrenceType) {
	o.Type = v
}

// GetDaysOfWeek returns the DaysOfWeek field value if set, zero value otherwise.
func (o *OptimizationRulesAPIRuleRecurrence) GetDaysOfWeek() []OptimizationRulesAPIDayOfTheWeek {
	if o == nil || IsNil(o.DaysOfWeek) {
		var ret []OptimizationRulesAPIDayOfTheWeek
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleRecurrence) GetDaysOfWeekOk() ([]OptimizationRulesAPIDayOfTheWeek, bool) {
	if o == nil || IsNil(o.DaysOfWeek) {
		return nil, false
	}
	return o.DaysOfWeek, true
}

// HasDaysOfWeek returns a boolean if a field has been set.
func (o *OptimizationRulesAPIRuleRecurrence) HasDaysOfWeek() bool {
	if o != nil && !IsNil(o.DaysOfWeek) {
		return true
	}

	return false
}

// SetDaysOfWeek gets a reference to the given []OptimizationRulesAPIDayOfTheWeek and assigns it to the DaysOfWeek field.
func (o *OptimizationRulesAPIRuleRecurrence) SetDaysOfWeek(v []OptimizationRulesAPIDayOfTheWeek) {
	o.DaysOfWeek = v
}

func (o OptimizationRulesAPIRuleRecurrence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	if !IsNil(o.TimesOfDay) {
		toSerialize["timesOfDay"] = o.TimesOfDay
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.DaysOfWeek) {
		toSerialize["daysOfWeek"] = o.DaysOfWeek
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRuleRecurrence struct {
	value *OptimizationRulesAPIRuleRecurrence
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleRecurrence) Get() *OptimizationRulesAPIRuleRecurrence {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleRecurrence) Set(val *OptimizationRulesAPIRuleRecurrence) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleRecurrence) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleRecurrence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleRecurrence(val *OptimizationRulesAPIRuleRecurrence) *NullableOptimizationRulesAPIRuleRecurrence {
	return &NullableOptimizationRulesAPIRuleRecurrence{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleRecurrence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleRecurrence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
