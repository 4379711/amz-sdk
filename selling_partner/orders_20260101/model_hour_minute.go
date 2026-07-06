package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the HourMinute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HourMinute{}

// HourMinute The time when the business opens or closes.
type HourMinute struct {
	// The hour when the business opens or closes, in 24-hour format (0-23).
	Hour *int32 `json:"hour,omitempty"`
	// The minute when the business opens or closes.
	Minute *int32 `json:"minute,omitempty"`
}

// NewHourMinute instantiates a new HourMinute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHourMinute() *HourMinute {
	this := HourMinute{}
	return &this
}

// NewHourMinuteWithDefaults instantiates a new HourMinute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHourMinuteWithDefaults() *HourMinute {
	this := HourMinute{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *HourMinute) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HourMinute) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *HourMinute) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *HourMinute) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *HourMinute) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HourMinute) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *HourMinute) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *HourMinute) SetMinute(v int32) {
	o.Minute = &v
}

func (o HourMinute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	return toSerialize, nil
}

type NullableHourMinute struct {
	value *HourMinute
	isSet bool
}

func (v NullableHourMinute) Get() *HourMinute {
	return v.value
}

func (v *NullableHourMinute) Set(val *HourMinute) {
	v.value = val
	v.isSet = true
}

func (v NullableHourMinute) IsSet() bool {
	return v.isSet
}

func (v *NullableHourMinute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHourMinute(val *HourMinute) *NullableHourMinute {
	return &NullableHourMinute{value: val, isSet: true}
}

func (v NullableHourMinute) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableHourMinute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
