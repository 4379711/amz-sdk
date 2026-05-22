package services

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the RangeSlotCapacityQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeSlotCapacityQuery{}

// RangeSlotCapacityQuery Request schema for the `getRangeSlotCapacity` operation. This schema is used to define the time range and capacity types that are being queried.
type RangeSlotCapacityQuery struct {
	// An array of capacity types which are being requested. Default value is `[SCHEDULED_CAPACITY]`.
	CapacityTypes []CapacityType `json:"capacityTypes,omitempty"`
	// Start date time from which the capacity slots are being requested in ISO 8601 format.
	StartDateTime time.Time `json:"startDateTime"`
	// End date time up to which the capacity slots are being requested in ISO 8601 format.
	EndDateTime time.Time `json:"endDateTime"`
}

type _RangeSlotCapacityQuery RangeSlotCapacityQuery

// NewRangeSlotCapacityQuery instantiates a new RangeSlotCapacityQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeSlotCapacityQuery(startDateTime time.Time, endDateTime time.Time) *RangeSlotCapacityQuery {
	this := RangeSlotCapacityQuery{}
	this.StartDateTime = startDateTime
	this.EndDateTime = endDateTime
	return &this
}

// NewRangeSlotCapacityQueryWithDefaults instantiates a new RangeSlotCapacityQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeSlotCapacityQueryWithDefaults() *RangeSlotCapacityQuery {
	this := RangeSlotCapacityQuery{}
	return &this
}

// GetCapacityTypes returns the CapacityTypes field value if set, zero value otherwise.
func (o *RangeSlotCapacityQuery) GetCapacityTypes() []CapacityType {
	if o == nil || IsNil(o.CapacityTypes) {
		var ret []CapacityType
		return ret
	}
	return o.CapacityTypes
}

// GetCapacityTypesOk returns a tuple with the CapacityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeSlotCapacityQuery) GetCapacityTypesOk() ([]CapacityType, bool) {
	if o == nil || IsNil(o.CapacityTypes) {
		return nil, false
	}
	return o.CapacityTypes, true
}

// HasCapacityTypes returns a boolean if a field has been set.
func (o *RangeSlotCapacityQuery) HasCapacityTypes() bool {
	if o != nil && !IsNil(o.CapacityTypes) {
		return true
	}

	return false
}

// SetCapacityTypes gets a reference to the given []CapacityType and assigns it to the CapacityTypes field.
func (o *RangeSlotCapacityQuery) SetCapacityTypes(v []CapacityType) {
	o.CapacityTypes = v
}

// GetStartDateTime returns the StartDateTime field value
func (o *RangeSlotCapacityQuery) GetStartDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *RangeSlotCapacityQuery) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *RangeSlotCapacityQuery) SetStartDateTime(v time.Time) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value
func (o *RangeSlotCapacityQuery) GetEndDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value
// and a boolean to check if the value has been set.
func (o *RangeSlotCapacityQuery) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDateTime, true
}

// SetEndDateTime sets field value
func (o *RangeSlotCapacityQuery) SetEndDateTime(v time.Time) {
	o.EndDateTime = v
}

func (o RangeSlotCapacityQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CapacityTypes) {
		toSerialize["capacityTypes"] = o.CapacityTypes
	}
	toSerialize["startDateTime"] = o.StartDateTime
	toSerialize["endDateTime"] = o.EndDateTime
	return toSerialize, nil
}

type NullableRangeSlotCapacityQuery struct {
	value *RangeSlotCapacityQuery
	isSet bool
}

func (v NullableRangeSlotCapacityQuery) Get() *RangeSlotCapacityQuery {
	return v.value
}

func (v *NullableRangeSlotCapacityQuery) Set(val *RangeSlotCapacityQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeSlotCapacityQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeSlotCapacityQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeSlotCapacityQuery(val *RangeSlotCapacityQuery) *NullableRangeSlotCapacityQuery {
	return &NullableRangeSlotCapacityQuery{value: val, isSet: true}
}

func (v NullableRangeSlotCapacityQuery) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRangeSlotCapacityQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
