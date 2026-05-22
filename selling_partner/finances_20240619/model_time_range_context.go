package finances_20240619

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the TimeRangeContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeRangeContext{}

// TimeRangeContext Additional information related to time range for transaction.
type TimeRangeContext struct {
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	StartTime *time.Time `json:"startTime,omitempty"`
	// Fields with a schema type of date are in ISO 8601 date time format (for example GroupBeginDate).
	EndTime *time.Time `json:"endTime,omitempty"`
}

// NewTimeRangeContext instantiates a new TimeRangeContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeRangeContext() *TimeRangeContext {
	this := TimeRangeContext{}
	return &this
}

// NewTimeRangeContextWithDefaults instantiates a new TimeRangeContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeRangeContextWithDefaults() *TimeRangeContext {
	this := TimeRangeContext{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeRangeContext) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRangeContext) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeRangeContext) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TimeRangeContext) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeRangeContext) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRangeContext) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeRangeContext) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *TimeRangeContext) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o TimeRangeContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableTimeRangeContext struct {
	value *TimeRangeContext
	isSet bool
}

func (v NullableTimeRangeContext) Get() *TimeRangeContext {
	return v.value
}

func (v *NullableTimeRangeContext) Set(val *TimeRangeContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeRangeContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeRangeContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeRangeContext(val *TimeRangeContext) *NullableTimeRangeContext {
	return &NullableTimeRangeContext{value: val, isSet: true}
}

func (v NullableTimeRangeContext) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTimeRangeContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
