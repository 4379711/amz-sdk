package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the DateRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateRange{}

// DateRange The time range within which something (for example, a delivery) will occur.
type DateRange struct {
	// Date timestamp
	Earliest time.Time `json:"earliest"`
	// Date timestamp
	Latest time.Time `json:"latest"`
}

type _DateRange DateRange

// NewDateRange instantiates a new DateRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateRange(earliest time.Time, latest time.Time) *DateRange {
	this := DateRange{}
	this.Earliest = earliest
	this.Latest = latest
	return &this
}

// NewDateRangeWithDefaults instantiates a new DateRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateRangeWithDefaults() *DateRange {
	this := DateRange{}
	return &this
}

// GetEarliest returns the Earliest field value
func (o *DateRange) GetEarliest() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Earliest
}

// GetEarliestOk returns a tuple with the Earliest field value
// and a boolean to check if the value has been set.
func (o *DateRange) GetEarliestOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Earliest, true
}

// SetEarliest sets field value
func (o *DateRange) SetEarliest(v time.Time) {
	o.Earliest = v
}

// GetLatest returns the Latest field value
func (o *DateRange) GetLatest() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Latest
}

// GetLatestOk returns a tuple with the Latest field value
// and a boolean to check if the value has been set.
func (o *DateRange) GetLatestOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latest, true
}

// SetLatest sets field value
func (o *DateRange) SetLatest(v time.Time) {
	o.Latest = v
}

func (o DateRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["earliest"] = o.Earliest
	toSerialize["latest"] = o.Latest
	return toSerialize, nil
}

type NullableDateRange struct {
	value *DateRange
	isSet bool
}

func (v NullableDateRange) Get() *DateRange {
	return v.value
}

func (v *NullableDateRange) Set(val *DateRange) {
	v.value = val
	v.isSet = true
}

func (v NullableDateRange) IsSet() bool {
	return v.isSet
}

func (v *NullableDateRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateRange(val *DateRange) *NullableDateRange {
	return &NullableDateRange{value: val, isSet: true}
}

func (v NullableDateRange) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDateRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
