package services

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the AppointmentSlotReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentSlotReport{}

// AppointmentSlotReport Availability information as per the service context queried.
type AppointmentSlotReport struct {
	// Defines the type of slots.
	SchedulingType *string `json:"schedulingType,omitempty"`
	// Start Time from which the appointment slots are generated in ISO 8601 format.
	StartTime *time.Time `json:"startTime,omitempty"`
	// End Time up to which the appointment slots are generated in ISO 8601 format.
	EndTime *time.Time `json:"endTime,omitempty"`
	// A list of time windows along with associated capacity in which the service can be performed.
	AppointmentSlots []AppointmentSlot `json:"appointmentSlots,omitempty"`
}

// NewAppointmentSlotReport instantiates a new AppointmentSlotReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentSlotReport() *AppointmentSlotReport {
	this := AppointmentSlotReport{}
	return &this
}

// NewAppointmentSlotReportWithDefaults instantiates a new AppointmentSlotReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentSlotReportWithDefaults() *AppointmentSlotReport {
	this := AppointmentSlotReport{}
	return &this
}

// GetSchedulingType returns the SchedulingType field value if set, zero value otherwise.
func (o *AppointmentSlotReport) GetSchedulingType() string {
	if o == nil || IsNil(o.SchedulingType) {
		var ret string
		return ret
	}
	return *o.SchedulingType
}

// GetSchedulingTypeOk returns a tuple with the SchedulingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSlotReport) GetSchedulingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SchedulingType) {
		return nil, false
	}
	return o.SchedulingType, true
}

// HasSchedulingType returns a boolean if a field has been set.
func (o *AppointmentSlotReport) HasSchedulingType() bool {
	if o != nil && !IsNil(o.SchedulingType) {
		return true
	}

	return false
}

// SetSchedulingType gets a reference to the given string and assigns it to the SchedulingType field.
func (o *AppointmentSlotReport) SetSchedulingType(v string) {
	o.SchedulingType = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AppointmentSlotReport) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSlotReport) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AppointmentSlotReport) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *AppointmentSlotReport) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AppointmentSlotReport) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSlotReport) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AppointmentSlotReport) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *AppointmentSlotReport) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetAppointmentSlots returns the AppointmentSlots field value if set, zero value otherwise.
func (o *AppointmentSlotReport) GetAppointmentSlots() []AppointmentSlot {
	if o == nil || IsNil(o.AppointmentSlots) {
		var ret []AppointmentSlot
		return ret
	}
	return o.AppointmentSlots
}

// GetAppointmentSlotsOk returns a tuple with the AppointmentSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSlotReport) GetAppointmentSlotsOk() ([]AppointmentSlot, bool) {
	if o == nil || IsNil(o.AppointmentSlots) {
		return nil, false
	}
	return o.AppointmentSlots, true
}

// HasAppointmentSlots returns a boolean if a field has been set.
func (o *AppointmentSlotReport) HasAppointmentSlots() bool {
	if o != nil && !IsNil(o.AppointmentSlots) {
		return true
	}

	return false
}

// SetAppointmentSlots gets a reference to the given []AppointmentSlot and assigns it to the AppointmentSlots field.
func (o *AppointmentSlotReport) SetAppointmentSlots(v []AppointmentSlot) {
	o.AppointmentSlots = v
}

func (o AppointmentSlotReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchedulingType) {
		toSerialize["schedulingType"] = o.SchedulingType
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.AppointmentSlots) {
		toSerialize["appointmentSlots"] = o.AppointmentSlots
	}
	return toSerialize, nil
}

type NullableAppointmentSlotReport struct {
	value *AppointmentSlotReport
	isSet bool
}

func (v NullableAppointmentSlotReport) Get() *AppointmentSlotReport {
	return v.value
}

func (v *NullableAppointmentSlotReport) Set(val *AppointmentSlotReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentSlotReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentSlotReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentSlotReport(val *AppointmentSlotReport) *NullableAppointmentSlotReport {
	return &NullableAppointmentSlotReport{value: val, isSet: true}
}

func (v NullableAppointmentSlotReport) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAppointmentSlotReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
