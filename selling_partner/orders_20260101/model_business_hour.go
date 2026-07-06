package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the BusinessHour type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessHour{}

// BusinessHour Business days and hours when the destination is open for deliveries.
type BusinessHour struct {
	// Specific day of the week for which operating hours are being defined.
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
	// Collection of time windows during which the location is available for deliveries on the specified day.
	TimeWindows []TimeWindow `json:"timeWindows,omitempty"`
}

// NewBusinessHour instantiates a new BusinessHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessHour() *BusinessHour {
	this := BusinessHour{}
	return &this
}

// NewBusinessHourWithDefaults instantiates a new BusinessHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessHourWithDefaults() *BusinessHour {
	this := BusinessHour{}
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *BusinessHour) GetDayOfWeek() string {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessHour) GetDayOfWeekOk() (*string, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *BusinessHour) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *BusinessHour) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetTimeWindows returns the TimeWindows field value if set, zero value otherwise.
func (o *BusinessHour) GetTimeWindows() []TimeWindow {
	if o == nil || IsNil(o.TimeWindows) {
		var ret []TimeWindow
		return ret
	}
	return o.TimeWindows
}

// GetTimeWindowsOk returns a tuple with the TimeWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessHour) GetTimeWindowsOk() ([]TimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindows) {
		return nil, false
	}
	return o.TimeWindows, true
}

// HasTimeWindows returns a boolean if a field has been set.
func (o *BusinessHour) HasTimeWindows() bool {
	if o != nil && !IsNil(o.TimeWindows) {
		return true
	}

	return false
}

// SetTimeWindows gets a reference to the given []TimeWindow and assigns it to the TimeWindows field.
func (o *BusinessHour) SetTimeWindows(v []TimeWindow) {
	o.TimeWindows = v
}

func (o BusinessHour) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DayOfWeek) {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	if !IsNil(o.TimeWindows) {
		toSerialize["timeWindows"] = o.TimeWindows
	}
	return toSerialize, nil
}

type NullableBusinessHour struct {
	value *BusinessHour
	isSet bool
}

func (v NullableBusinessHour) Get() *BusinessHour {
	return v.value
}

func (v *NullableBusinessHour) Set(val *BusinessHour) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessHour) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessHour(val *BusinessHour) *NullableBusinessHour {
	return &NullableBusinessHour{value: val, isSet: true}
}

func (v NullableBusinessHour) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBusinessHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
