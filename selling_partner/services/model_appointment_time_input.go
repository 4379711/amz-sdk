package services

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the AppointmentTimeInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentTimeInput{}

// AppointmentTimeInput The input appointment time details.
type AppointmentTimeInput struct {
	// The date, time in UTC for the start time of an appointment in ISO 8601 format.
	StartTime time.Time `json:"startTime"`
	// The duration of an appointment in minutes.
	DurationInMinutes *int32 `json:"durationInMinutes,omitempty"`
}

type _AppointmentTimeInput AppointmentTimeInput

// NewAppointmentTimeInput instantiates a new AppointmentTimeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentTimeInput(startTime time.Time) *AppointmentTimeInput {
	this := AppointmentTimeInput{}
	this.StartTime = startTime
	return &this
}

// NewAppointmentTimeInputWithDefaults instantiates a new AppointmentTimeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentTimeInputWithDefaults() *AppointmentTimeInput {
	this := AppointmentTimeInput{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *AppointmentTimeInput) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AppointmentTimeInput) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *AppointmentTimeInput) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetDurationInMinutes returns the DurationInMinutes field value if set, zero value otherwise.
func (o *AppointmentTimeInput) GetDurationInMinutes() int32 {
	if o == nil || IsNil(o.DurationInMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationInMinutes
}

// GetDurationInMinutesOk returns a tuple with the DurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentTimeInput) GetDurationInMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInMinutes) {
		return nil, false
	}
	return o.DurationInMinutes, true
}

// HasDurationInMinutes returns a boolean if a field has been set.
func (o *AppointmentTimeInput) HasDurationInMinutes() bool {
	if o != nil && !IsNil(o.DurationInMinutes) {
		return true
	}

	return false
}

// SetDurationInMinutes gets a reference to the given int32 and assigns it to the DurationInMinutes field.
func (o *AppointmentTimeInput) SetDurationInMinutes(v int32) {
	o.DurationInMinutes = &v
}

func (o AppointmentTimeInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	if !IsNil(o.DurationInMinutes) {
		toSerialize["durationInMinutes"] = o.DurationInMinutes
	}
	return toSerialize, nil
}

type NullableAppointmentTimeInput struct {
	value *AppointmentTimeInput
	isSet bool
}

func (v NullableAppointmentTimeInput) Get() *AppointmentTimeInput {
	return v.value
}

func (v *NullableAppointmentTimeInput) Set(val *AppointmentTimeInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentTimeInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentTimeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentTimeInput(val *AppointmentTimeInput) *NullableAppointmentTimeInput {
	return &NullableAppointmentTimeInput{value: val, isSet: true}
}

func (v NullableAppointmentTimeInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAppointmentTimeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
