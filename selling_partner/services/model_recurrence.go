package services

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the Recurrence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recurrence{}

// Recurrence Repeated occurrence of an event in a time range.
type Recurrence struct {
	// End time of the recurrence.
	EndTime time.Time `json:"endTime"`
	// Days of the week when recurrence is valid. If the schedule is valid every Monday, input will only contain `MONDAY` in the list.
	DaysOfWeek []DayOfWeek `json:"daysOfWeek,omitempty"`
	// Days of the month when recurrence is valid.
	DaysOfMonth []int32 `json:"daysOfMonth,omitempty"`
}

type _Recurrence Recurrence

// NewRecurrence instantiates a new Recurrence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurrence(endTime time.Time) *Recurrence {
	this := Recurrence{}
	this.EndTime = endTime
	return &this
}

// NewRecurrenceWithDefaults instantiates a new Recurrence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurrenceWithDefaults() *Recurrence {
	this := Recurrence{}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *Recurrence) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *Recurrence) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *Recurrence) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetDaysOfWeek returns the DaysOfWeek field value if set, zero value otherwise.
func (o *Recurrence) GetDaysOfWeek() []DayOfWeek {
	if o == nil || IsNil(o.DaysOfWeek) {
		var ret []DayOfWeek
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurrence) GetDaysOfWeekOk() ([]DayOfWeek, bool) {
	if o == nil || IsNil(o.DaysOfWeek) {
		return nil, false
	}
	return o.DaysOfWeek, true
}

// HasDaysOfWeek returns a boolean if a field has been set.
func (o *Recurrence) HasDaysOfWeek() bool {
	if o != nil && !IsNil(o.DaysOfWeek) {
		return true
	}

	return false
}

// SetDaysOfWeek gets a reference to the given []DayOfWeek and assigns it to the DaysOfWeek field.
func (o *Recurrence) SetDaysOfWeek(v []DayOfWeek) {
	o.DaysOfWeek = v
}

// GetDaysOfMonth returns the DaysOfMonth field value if set, zero value otherwise.
func (o *Recurrence) GetDaysOfMonth() []int32 {
	if o == nil || IsNil(o.DaysOfMonth) {
		var ret []int32
		return ret
	}
	return o.DaysOfMonth
}

// GetDaysOfMonthOk returns a tuple with the DaysOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurrence) GetDaysOfMonthOk() ([]int32, bool) {
	if o == nil || IsNil(o.DaysOfMonth) {
		return nil, false
	}
	return o.DaysOfMonth, true
}

// HasDaysOfMonth returns a boolean if a field has been set.
func (o *Recurrence) HasDaysOfMonth() bool {
	if o != nil && !IsNil(o.DaysOfMonth) {
		return true
	}

	return false
}

// SetDaysOfMonth gets a reference to the given []int32 and assigns it to the DaysOfMonth field.
func (o *Recurrence) SetDaysOfMonth(v []int32) {
	o.DaysOfMonth = v
}

func (o Recurrence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endTime"] = o.EndTime
	if !IsNil(o.DaysOfWeek) {
		toSerialize["daysOfWeek"] = o.DaysOfWeek
	}
	if !IsNil(o.DaysOfMonth) {
		toSerialize["daysOfMonth"] = o.DaysOfMonth
	}
	return toSerialize, nil
}

type NullableRecurrence struct {
	value *Recurrence
	isSet bool
}

func (v NullableRecurrence) Get() *Recurrence {
	return v.value
}

func (v *NullableRecurrence) Set(val *Recurrence) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrence) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrence(val *Recurrence) *NullableRecurrence {
	return &NullableRecurrence{value: val, isSet: true}
}

func (v NullableRecurrence) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecurrence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
