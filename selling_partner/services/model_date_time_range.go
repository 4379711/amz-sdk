package services

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the DateTimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateTimeRange{}

// DateTimeRange A range of time.
type DateTimeRange struct {
	// The beginning of the time range. Must be in UTC in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	StartTime time.Time `json:"startTime"`
	// The end of the time range. Must be in UTC in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	EndTime time.Time `json:"endTime"`
}

type _DateTimeRange DateTimeRange

// NewDateTimeRange instantiates a new DateTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeRange(startTime time.Time, endTime time.Time) *DateTimeRange {
	this := DateTimeRange{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewDateTimeRangeWithDefaults instantiates a new DateTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeRangeWithDefaults() *DateTimeRange {
	this := DateTimeRange{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *DateTimeRange) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *DateTimeRange) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *DateTimeRange) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *DateTimeRange) SetEndTime(v time.Time) {
	o.EndTime = v
}

func (o DateTimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	return toSerialize, nil
}

type NullableDateTimeRange struct {
	value *DateTimeRange
	isSet bool
}

func (v NullableDateTimeRange) Get() *DateTimeRange {
	return v.value
}

func (v *NullableDateTimeRange) Set(val *DateTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeRange(val *DateTimeRange) *NullableDateTimeRange {
	return &NullableDateTimeRange{value: val, isSet: true}
}

func (v NullableDateTimeRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDateTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
