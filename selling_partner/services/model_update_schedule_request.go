package services

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateScheduleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateScheduleRequest{}

// UpdateScheduleRequest Request schema for the `updateSchedule` operation.
type UpdateScheduleRequest struct {
	// List of `AvailabilityRecord`s to represent the capacity of a resource over a time range.
	Schedules []AvailabilityRecord `json:"schedules"`
}

type _UpdateScheduleRequest UpdateScheduleRequest

// NewUpdateScheduleRequest instantiates a new UpdateScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateScheduleRequest(schedules []AvailabilityRecord) *UpdateScheduleRequest {
	this := UpdateScheduleRequest{}
	this.Schedules = schedules
	return &this
}

// NewUpdateScheduleRequestWithDefaults instantiates a new UpdateScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScheduleRequestWithDefaults() *UpdateScheduleRequest {
	this := UpdateScheduleRequest{}
	return &this
}

// GetSchedules returns the Schedules field value
func (o *UpdateScheduleRequest) GetSchedules() []AvailabilityRecord {
	if o == nil {
		var ret []AvailabilityRecord
		return ret
	}

	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value
// and a boolean to check if the value has been set.
func (o *UpdateScheduleRequest) GetSchedulesOk() ([]AvailabilityRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedules, true
}

// SetSchedules sets field value
func (o *UpdateScheduleRequest) SetSchedules(v []AvailabilityRecord) {
	o.Schedules = v
}

func (o UpdateScheduleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schedules"] = o.Schedules
	return toSerialize, nil
}

type NullableUpdateScheduleRequest struct {
	value *UpdateScheduleRequest
	isSet bool
}

func (v NullableUpdateScheduleRequest) Get() *UpdateScheduleRequest {
	return v.value
}

func (v *NullableUpdateScheduleRequest) Set(val *UpdateScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScheduleRequest(val *UpdateScheduleRequest) *NullableUpdateScheduleRequest {
	return &NullableUpdateScheduleRequest{value: val, isSet: true}
}

func (v NullableUpdateScheduleRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
