package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the BusinessHours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessHours{}

// BusinessHours Business days and hours when the destination is open for deliveries.
type BusinessHours struct {
	// Day of the week.
	DayOfWeek *string `json:"DayOfWeek,omitempty"`
	// Time window during the day when the business is open.
	OpenIntervals []OpenInterval `json:"OpenIntervals,omitempty"`
}

// NewBusinessHours instantiates a new BusinessHours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessHours() *BusinessHours {
	this := BusinessHours{}
	return &this
}

// NewBusinessHoursWithDefaults instantiates a new BusinessHours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessHoursWithDefaults() *BusinessHours {
	this := BusinessHours{}
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *BusinessHours) GetDayOfWeek() string {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessHours) GetDayOfWeekOk() (*string, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *BusinessHours) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *BusinessHours) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetOpenIntervals returns the OpenIntervals field value if set, zero value otherwise.
func (o *BusinessHours) GetOpenIntervals() []OpenInterval {
	if o == nil || IsNil(o.OpenIntervals) {
		var ret []OpenInterval
		return ret
	}
	return o.OpenIntervals
}

// GetOpenIntervalsOk returns a tuple with the OpenIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessHours) GetOpenIntervalsOk() ([]OpenInterval, bool) {
	if o == nil || IsNil(o.OpenIntervals) {
		return nil, false
	}
	return o.OpenIntervals, true
}

// HasOpenIntervals returns a boolean if a field has been set.
func (o *BusinessHours) HasOpenIntervals() bool {
	if o != nil && !IsNil(o.OpenIntervals) {
		return true
	}

	return false
}

// SetOpenIntervals gets a reference to the given []OpenInterval and assigns it to the OpenIntervals field.
func (o *BusinessHours) SetOpenIntervals(v []OpenInterval) {
	o.OpenIntervals = v
}

func (o BusinessHours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DayOfWeek) {
		toSerialize["DayOfWeek"] = o.DayOfWeek
	}
	if !IsNil(o.OpenIntervals) {
		toSerialize["OpenIntervals"] = o.OpenIntervals
	}
	return toSerialize, nil
}

type NullableBusinessHours struct {
	value *BusinessHours
	isSet bool
}

func (v NullableBusinessHours) Get() *BusinessHours {
	return v.value
}

func (v *NullableBusinessHours) Set(val *BusinessHours) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessHours) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessHours(val *BusinessHours) *NullableBusinessHours {
	return &NullableBusinessHours{value: val, isSet: true}
}

func (v NullableBusinessHours) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBusinessHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
