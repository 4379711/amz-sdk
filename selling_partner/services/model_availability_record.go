package services

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the AvailabilityRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityRecord{}

// AvailabilityRecord `AvailabilityRecord` to represent the capacity of a resource over a time range.
type AvailabilityRecord struct {
	// Denotes the time from when the resource is available in a day in ISO-8601 format.
	StartTime time.Time `json:"startTime"`
	// Denotes the time till when the resource is available in a day in ISO-8601 format.
	EndTime    time.Time   `json:"endTime"`
	Recurrence *Recurrence `json:"recurrence,omitempty"`
	// Signifies the capacity of a resource which is available.
	Capacity *int32 `json:"capacity,omitempty"`
}

type _AvailabilityRecord AvailabilityRecord

// NewAvailabilityRecord instantiates a new AvailabilityRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityRecord(startTime time.Time, endTime time.Time) *AvailabilityRecord {
	this := AvailabilityRecord{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewAvailabilityRecordWithDefaults instantiates a new AvailabilityRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityRecordWithDefaults() *AvailabilityRecord {
	this := AvailabilityRecord{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *AvailabilityRecord) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AvailabilityRecord) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *AvailabilityRecord) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *AvailabilityRecord) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *AvailabilityRecord) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *AvailabilityRecord) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *AvailabilityRecord) GetRecurrence() Recurrence {
	if o == nil || IsNil(o.Recurrence) {
		var ret Recurrence
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityRecord) GetRecurrenceOk() (*Recurrence, bool) {
	if o == nil || IsNil(o.Recurrence) {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *AvailabilityRecord) HasRecurrence() bool {
	if o != nil && !IsNil(o.Recurrence) {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given Recurrence and assigns it to the Recurrence field.
func (o *AvailabilityRecord) SetRecurrence(v Recurrence) {
	o.Recurrence = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *AvailabilityRecord) GetCapacity() int32 {
	if o == nil || IsNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityRecord) GetCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *AvailabilityRecord) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *AvailabilityRecord) SetCapacity(v int32) {
	o.Capacity = &v
}

func (o AvailabilityRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	if !IsNil(o.Recurrence) {
		toSerialize["recurrence"] = o.Recurrence
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	return toSerialize, nil
}

type NullableAvailabilityRecord struct {
	value *AvailabilityRecord
	isSet bool
}

func (v NullableAvailabilityRecord) Get() *AvailabilityRecord {
	return v.value
}

func (v *NullableAvailabilityRecord) Set(val *AvailabilityRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityRecord(val *AvailabilityRecord) *NullableAvailabilityRecord {
	return &NullableAvailabilityRecord{value: val, isSet: true}
}

func (v NullableAvailabilityRecord) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAvailabilityRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
