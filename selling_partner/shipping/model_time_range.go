package shipping

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the TimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeRange{}

// TimeRange The time range.
type TimeRange struct {
	// The start date and time. This defaults to the current date and time.
	Start *time.Time `json:"start,omitempty"`
	// The end date and time. This must come after the value of start. This defaults to the next business day from the start.
	End *time.Time `json:"end,omitempty"`
}

// NewTimeRange instantiates a new TimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeRange() *TimeRange {
	this := TimeRange{}
	return &this
}

// NewTimeRangeWithDefaults instantiates a new TimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeRangeWithDefaults() *TimeRange {
	this := TimeRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *TimeRange) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRange) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *TimeRange) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *TimeRange) SetStart(v time.Time) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *TimeRange) GetEnd() time.Time {
	if o == nil || IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRange) GetEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *TimeRange) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *TimeRange) SetEnd(v time.Time) {
	o.End = &v
}

func (o TimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableTimeRange struct {
	value *TimeRange
	isSet bool
}

func (v NullableTimeRange) Get() *TimeRange {
	return v.value
}

func (v *NullableTimeRange) Set(val *TimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeRange(val *TimeRange) *NullableTimeRange {
	return &NullableTimeRange{value: val, isSet: true}
}

func (v NullableTimeRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
