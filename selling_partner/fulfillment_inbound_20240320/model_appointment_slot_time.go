package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the AppointmentSlotTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentSlotTime{}

// AppointmentSlotTime An appointment slot time with start and end.
type AppointmentSlotTime struct {
	// The end timestamp of the appointment in UTC.
	EndTime time.Time `json:"endTime"`
	// The start timestamp of the appointment in UTC.
	StartTime time.Time `json:"startTime"`
}

type _AppointmentSlotTime AppointmentSlotTime

// NewAppointmentSlotTime instantiates a new AppointmentSlotTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentSlotTime(endTime time.Time, startTime time.Time) *AppointmentSlotTime {
	this := AppointmentSlotTime{}
	this.EndTime = endTime
	this.StartTime = startTime
	return &this
}

// NewAppointmentSlotTimeWithDefaults instantiates a new AppointmentSlotTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentSlotTimeWithDefaults() *AppointmentSlotTime {
	this := AppointmentSlotTime{}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *AppointmentSlotTime) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *AppointmentSlotTime) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *AppointmentSlotTime) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetStartTime returns the StartTime field value
func (o *AppointmentSlotTime) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AppointmentSlotTime) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *AppointmentSlotTime) SetStartTime(v time.Time) {
	o.StartTime = v
}

func (o AppointmentSlotTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endTime"] = o.EndTime
	toSerialize["startTime"] = o.StartTime
	return toSerialize, nil
}

type NullableAppointmentSlotTime struct {
	value *AppointmentSlotTime
	isSet bool
}

func (v NullableAppointmentSlotTime) Get() *AppointmentSlotTime {
	return v.value
}

func (v *NullableAppointmentSlotTime) Set(val *AppointmentSlotTime) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentSlotTime) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentSlotTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentSlotTime(val *AppointmentSlotTime) *NullableAppointmentSlotTime {
	return &NullableAppointmentSlotTime{value: val, isSet: true}
}

func (v NullableAppointmentSlotTime) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAppointmentSlotTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
