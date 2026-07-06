package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the TimeWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeWindow{}

// TimeWindow Specific time interval defining the start and end times.
type TimeWindow struct {
	StartTime *HourMinute `json:"startTime,omitempty"`
	EndTime   *HourMinute `json:"endTime,omitempty"`
}

// NewTimeWindow instantiates a new TimeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWindow() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// NewTimeWindowWithDefaults instantiates a new TimeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWindowWithDefaults() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeWindow) GetStartTime() HourMinute {
	if o == nil || IsNil(o.StartTime) {
		var ret HourMinute
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetStartTimeOk() (*HourMinute, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeWindow) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given HourMinute and assigns it to the StartTime field.
func (o *TimeWindow) SetStartTime(v HourMinute) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeWindow) GetEndTime() HourMinute {
	if o == nil || IsNil(o.EndTime) {
		var ret HourMinute
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetEndTimeOk() (*HourMinute, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeWindow) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given HourMinute and assigns it to the EndTime field.
func (o *TimeWindow) SetEndTime(v HourMinute) {
	o.EndTime = &v
}

func (o TimeWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableTimeWindow struct {
	value *TimeWindow
	isSet bool
}

func (v NullableTimeWindow) Get() *TimeWindow {
	return v.value
}

func (v *NullableTimeWindow) Set(val *TimeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWindow(val *TimeWindow) *NullableTimeWindow {
	return &NullableTimeWindow{value: val, isSet: true}
}

func (v NullableTimeWindow) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTimeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
