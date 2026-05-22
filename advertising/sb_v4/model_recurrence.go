package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Recurrence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recurrence{}

// Recurrence struct for Recurrence
type Recurrence struct {
	Type *RecurrenceType `json:"type,omitempty"`
	// Object representing days of the week for weekly type rule. It is not required for daily recurrence type
	DaysOfWeek []DayOfWeek `json:"daysOfWeek,omitempty"`
	// List of objects representing start and end time of desired intra-day budget rule window
	IntraDaySchedule []TimeOfDay `json:"intraDaySchedule,omitempty"`
}

// NewRecurrence instantiates a new Recurrence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurrence() *Recurrence {
	this := Recurrence{}
	return &this
}

// NewRecurrenceWithDefaults instantiates a new Recurrence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurrenceWithDefaults() *Recurrence {
	this := Recurrence{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Recurrence) GetType() RecurrenceType {
	if o == nil || IsNil(o.Type) {
		var ret RecurrenceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurrence) GetTypeOk() (*RecurrenceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Recurrence) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RecurrenceType and assigns it to the Type field.
func (o *Recurrence) SetType(v RecurrenceType) {
	o.Type = &v
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

// GetIntraDaySchedule returns the IntraDaySchedule field value if set, zero value otherwise.
func (o *Recurrence) GetIntraDaySchedule() []TimeOfDay {
	if o == nil || IsNil(o.IntraDaySchedule) {
		var ret []TimeOfDay
		return ret
	}
	return o.IntraDaySchedule
}

// GetIntraDayScheduleOk returns a tuple with the IntraDaySchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurrence) GetIntraDayScheduleOk() ([]TimeOfDay, bool) {
	if o == nil || IsNil(o.IntraDaySchedule) {
		return nil, false
	}
	return o.IntraDaySchedule, true
}

// HasIntraDaySchedule returns a boolean if a field has been set.
func (o *Recurrence) HasIntraDaySchedule() bool {
	if o != nil && !IsNil(o.IntraDaySchedule) {
		return true
	}

	return false
}

// SetIntraDaySchedule gets a reference to the given []TimeOfDay and assigns it to the IntraDaySchedule field.
func (o *Recurrence) SetIntraDaySchedule(v []TimeOfDay) {
	o.IntraDaySchedule = v
}

func (o Recurrence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DaysOfWeek) {
		toSerialize["daysOfWeek"] = o.DaysOfWeek
	}
	if !IsNil(o.IntraDaySchedule) {
		toSerialize["intraDaySchedule"] = o.IntraDaySchedule
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
