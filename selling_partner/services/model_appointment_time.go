package services

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the AppointmentTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentTime{}

// AppointmentTime The time of the appointment window.
type AppointmentTime struct {
	// The date and time of the start of the appointment window in ISO 8601 format.
	StartTime time.Time `json:"startTime"`
	// The duration of the appointment window, in minutes.
	DurationInMinutes int32 `json:"durationInMinutes"`
}

type _AppointmentTime AppointmentTime

// NewAppointmentTime instantiates a new AppointmentTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentTime(startTime time.Time, durationInMinutes int32) *AppointmentTime {
	this := AppointmentTime{}
	this.StartTime = startTime
	this.DurationInMinutes = durationInMinutes
	return &this
}

// NewAppointmentTimeWithDefaults instantiates a new AppointmentTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentTimeWithDefaults() *AppointmentTime {
	this := AppointmentTime{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *AppointmentTime) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AppointmentTime) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *AppointmentTime) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetDurationInMinutes returns the DurationInMinutes field value
func (o *AppointmentTime) GetDurationInMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationInMinutes
}

// GetDurationInMinutesOk returns a tuple with the DurationInMinutes field value
// and a boolean to check if the value has been set.
func (o *AppointmentTime) GetDurationInMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationInMinutes, true
}

// SetDurationInMinutes sets field value
func (o *AppointmentTime) SetDurationInMinutes(v int32) {
	o.DurationInMinutes = v
}

func (o AppointmentTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["durationInMinutes"] = o.DurationInMinutes
	return toSerialize, nil
}

type NullableAppointmentTime struct {
	value *AppointmentTime
	isSet bool
}

func (v NullableAppointmentTime) Get() *AppointmentTime {
	return v.value
}

func (v *NullableAppointmentTime) Set(val *AppointmentTime) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentTime) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentTime(val *AppointmentTime) *NullableAppointmentTime {
	return &NullableAppointmentTime{value: val, isSet: true}
}

func (v NullableAppointmentTime) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAppointmentTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
