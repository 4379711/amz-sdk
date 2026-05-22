package services

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the FixedSlotCapacityQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FixedSlotCapacityQuery{}

// FixedSlotCapacityQuery Request schema for the `getFixedSlotCapacity` operation. This schema is used to define the time range, capacity types and slot duration which are being queried.
type FixedSlotCapacityQuery struct {
	// An array of capacity types which are being requested. Default value is `[SCHEDULED_CAPACITY]`.
	CapacityTypes []CapacityType `json:"capacityTypes,omitempty"`
	// Size in which slots are being requested. This value should be a multiple of 5 and fall in the range: 5 <= `slotDuration` <= 360.
	SlotDuration *float32 `json:"slotDuration,omitempty"`
	// Start date time from which the capacity slots are being requested in ISO 8601 format.
	StartDateTime time.Time `json:"startDateTime"`
	// End date time up to which the capacity slots are being requested in ISO 8601 format.
	EndDateTime time.Time `json:"endDateTime"`
}

type _FixedSlotCapacityQuery FixedSlotCapacityQuery

// NewFixedSlotCapacityQuery instantiates a new FixedSlotCapacityQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedSlotCapacityQuery(startDateTime time.Time, endDateTime time.Time) *FixedSlotCapacityQuery {
	this := FixedSlotCapacityQuery{}
	this.StartDateTime = startDateTime
	this.EndDateTime = endDateTime
	return &this
}

// NewFixedSlotCapacityQueryWithDefaults instantiates a new FixedSlotCapacityQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedSlotCapacityQueryWithDefaults() *FixedSlotCapacityQuery {
	this := FixedSlotCapacityQuery{}
	return &this
}

// GetCapacityTypes returns the CapacityTypes field value if set, zero value otherwise.
func (o *FixedSlotCapacityQuery) GetCapacityTypes() []CapacityType {
	if o == nil || IsNil(o.CapacityTypes) {
		var ret []CapacityType
		return ret
	}
	return o.CapacityTypes
}

// GetCapacityTypesOk returns a tuple with the CapacityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedSlotCapacityQuery) GetCapacityTypesOk() ([]CapacityType, bool) {
	if o == nil || IsNil(o.CapacityTypes) {
		return nil, false
	}
	return o.CapacityTypes, true
}

// HasCapacityTypes returns a boolean if a field has been set.
func (o *FixedSlotCapacityQuery) HasCapacityTypes() bool {
	if o != nil && !IsNil(o.CapacityTypes) {
		return true
	}

	return false
}

// SetCapacityTypes gets a reference to the given []CapacityType and assigns it to the CapacityTypes field.
func (o *FixedSlotCapacityQuery) SetCapacityTypes(v []CapacityType) {
	o.CapacityTypes = v
}

// GetSlotDuration returns the SlotDuration field value if set, zero value otherwise.
func (o *FixedSlotCapacityQuery) GetSlotDuration() float32 {
	if o == nil || IsNil(o.SlotDuration) {
		var ret float32
		return ret
	}
	return *o.SlotDuration
}

// GetSlotDurationOk returns a tuple with the SlotDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedSlotCapacityQuery) GetSlotDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.SlotDuration) {
		return nil, false
	}
	return o.SlotDuration, true
}

// HasSlotDuration returns a boolean if a field has been set.
func (o *FixedSlotCapacityQuery) HasSlotDuration() bool {
	if o != nil && !IsNil(o.SlotDuration) {
		return true
	}

	return false
}

// SetSlotDuration gets a reference to the given float32 and assigns it to the SlotDuration field.
func (o *FixedSlotCapacityQuery) SetSlotDuration(v float32) {
	o.SlotDuration = &v
}

// GetStartDateTime returns the StartDateTime field value
func (o *FixedSlotCapacityQuery) GetStartDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *FixedSlotCapacityQuery) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *FixedSlotCapacityQuery) SetStartDateTime(v time.Time) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value
func (o *FixedSlotCapacityQuery) GetEndDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value
// and a boolean to check if the value has been set.
func (o *FixedSlotCapacityQuery) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDateTime, true
}

// SetEndDateTime sets field value
func (o *FixedSlotCapacityQuery) SetEndDateTime(v time.Time) {
	o.EndDateTime = v
}

func (o FixedSlotCapacityQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CapacityTypes) {
		toSerialize["capacityTypes"] = o.CapacityTypes
	}
	if !IsNil(o.SlotDuration) {
		toSerialize["slotDuration"] = o.SlotDuration
	}
	toSerialize["startDateTime"] = o.StartDateTime
	toSerialize["endDateTime"] = o.EndDateTime
	return toSerialize, nil
}

type NullableFixedSlotCapacityQuery struct {
	value *FixedSlotCapacityQuery
	isSet bool
}

func (v NullableFixedSlotCapacityQuery) Get() *FixedSlotCapacityQuery {
	return v.value
}

func (v *NullableFixedSlotCapacityQuery) Set(val *FixedSlotCapacityQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedSlotCapacityQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedSlotCapacityQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedSlotCapacityQuery(val *FixedSlotCapacityQuery) *NullableFixedSlotCapacityQuery {
	return &NullableFixedSlotCapacityQuery{value: val, isSet: true}
}

func (v NullableFixedSlotCapacityQuery) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFixedSlotCapacityQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
