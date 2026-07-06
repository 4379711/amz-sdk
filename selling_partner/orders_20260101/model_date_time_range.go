package orders_20260101

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the DateTimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateTimeRange{}

// DateTimeRange A time period with start and end boundaries.
type DateTimeRange struct {
	// The beginning of the time period, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	EarliestDateTime *time.Time `json:"earliestDateTime,omitempty"`
	// The end of the time period, in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	LatestDateTime *time.Time `json:"latestDateTime,omitempty"`
}

// NewDateTimeRange instantiates a new DateTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeRange() *DateTimeRange {
	this := DateTimeRange{}
	return &this
}

// NewDateTimeRangeWithDefaults instantiates a new DateTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeRangeWithDefaults() *DateTimeRange {
	this := DateTimeRange{}
	return &this
}

// GetEarliestDateTime returns the EarliestDateTime field value if set, zero value otherwise.
func (o *DateTimeRange) GetEarliestDateTime() time.Time {
	if o == nil || IsNil(o.EarliestDateTime) {
		var ret time.Time
		return ret
	}
	return *o.EarliestDateTime
}

// GetEarliestDateTimeOk returns a tuple with the EarliestDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetEarliestDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestDateTime) {
		return nil, false
	}
	return o.EarliestDateTime, true
}

// HasEarliestDateTime returns a boolean if a field has been set.
func (o *DateTimeRange) HasEarliestDateTime() bool {
	if o != nil && !IsNil(o.EarliestDateTime) {
		return true
	}

	return false
}

// SetEarliestDateTime gets a reference to the given time.Time and assigns it to the EarliestDateTime field.
func (o *DateTimeRange) SetEarliestDateTime(v time.Time) {
	o.EarliestDateTime = &v
}

// GetLatestDateTime returns the LatestDateTime field value if set, zero value otherwise.
func (o *DateTimeRange) GetLatestDateTime() time.Time {
	if o == nil || IsNil(o.LatestDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LatestDateTime
}

// GetLatestDateTimeOk returns a tuple with the LatestDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetLatestDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LatestDateTime) {
		return nil, false
	}
	return o.LatestDateTime, true
}

// HasLatestDateTime returns a boolean if a field has been set.
func (o *DateTimeRange) HasLatestDateTime() bool {
	if o != nil && !IsNil(o.LatestDateTime) {
		return true
	}

	return false
}

// SetLatestDateTime gets a reference to the given time.Time and assigns it to the LatestDateTime field.
func (o *DateTimeRange) SetLatestDateTime(v time.Time) {
	o.LatestDateTime = &v
}

func (o DateTimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EarliestDateTime) {
		toSerialize["earliestDateTime"] = o.EarliestDateTime
	}
	if !IsNil(o.LatestDateTime) {
		toSerialize["latestDateTime"] = o.LatestDateTime
	}
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
