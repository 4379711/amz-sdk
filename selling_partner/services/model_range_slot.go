package services

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the RangeSlot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeSlot{}

// RangeSlot Capacity slots represented in a format similar to availability rules.
type RangeSlot struct {
	// Start date time of slot in ISO 8601 format with precision of seconds.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// End date time of slot in ISO 8601 format with precision of seconds.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Capacity of the slot.
	Capacity *int32 `json:"capacity,omitempty"`
}

// NewRangeSlot instantiates a new RangeSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeSlot() *RangeSlot {
	this := RangeSlot{}
	return &this
}

// NewRangeSlotWithDefaults instantiates a new RangeSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeSlotWithDefaults() *RangeSlot {
	this := RangeSlot{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *RangeSlot) GetStartDateTime() time.Time {
	if o == nil || IsNil(o.StartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeSlot) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *RangeSlot) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *RangeSlot) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *RangeSlot) GetEndDateTime() time.Time {
	if o == nil || IsNil(o.EndDateTime) {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeSlot) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDateTime) {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *RangeSlot) HasEndDateTime() bool {
	if o != nil && !IsNil(o.EndDateTime) {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *RangeSlot) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *RangeSlot) GetCapacity() int32 {
	if o == nil || IsNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeSlot) GetCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *RangeSlot) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *RangeSlot) SetCapacity(v int32) {
	o.Capacity = &v
}

func (o RangeSlot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !IsNil(o.EndDateTime) {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	return toSerialize, nil
}

type NullableRangeSlot struct {
	value *RangeSlot
	isSet bool
}

func (v NullableRangeSlot) Get() *RangeSlot {
	return v.value
}

func (v *NullableRangeSlot) Set(val *RangeSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeSlot(val *RangeSlot) *NullableRangeSlot {
	return &NullableRangeSlot{value: val, isSet: true}
}

func (v NullableRangeSlot) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRangeSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
